// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	"context"
	"errors"
	"fmt"
	log "github.com/golang/glog"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/joincivil/civil-events-crawler/pkg/listener"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

const (
	blockHeaderExpirySecs = 60 * 5 // 5 mins
)

// NewEventCollector creates a new event collector
func NewEventCollector(chain ethereum.ChainReader, client bind.ContractBackend,
	filterers []model.ContractFilterers, watchers []model.ContractWatchers,
	retrieverPersister model.RetrieverMetaDataPersister,
	listenerPersister model.ListenerMetaDataPersister,
	eventDataPersister model.EventDataPersister, triggers []Trigger,
	startBlock uint64) *EventCollector {
	eventcollector := &EventCollector{
		chain:              chain,
		client:             client,
		filterers:          filterers,
		watchers:           watchers,
		retrieverPersister: retrieverPersister,
		listenerPersister:  listenerPersister,
		eventDataPersister: eventDataPersister,
		triggers:           triggers,
		startBlock:         startBlock,
	}
	return eventcollector
}

// EventCollector handles logic for getting historical and live events
type EventCollector struct {
	chain ethereum.ChainReader

	client bind.ContractBackend

	triggers []Trigger

	filterers []model.ContractFilterers

	watchers []model.ContractWatchers

	retrieverPersister model.RetrieverMetaDataPersister

	listenerPersister model.ListenerMetaDataPersister

	eventDataPersister model.EventDataPersister

	listen *listener.EventListener

	retrieve *retriever.EventRetriever

	startBlock uint64

	// quitChan is created in StartCollection() and stops the goroutine listening for events.
	quitChan chan interface{}

	mutex sync.Mutex

	headerCache *utils.BlockHeaderCache
}

// StartCollection contains logic to run retriever and listener.
func (c *EventCollector) StartCollection() error {
	err := c.retrieveEvents()
	if err != nil {
		return fmt.Errorf("Error retrieving events: err: %v", err)
	}
	pastEvents := c.retrieve.PastEvents
	err = c.updateEventTimesFromBlockHeaders(pastEvents)
	if err != nil {
		return fmt.Errorf("Error updating dates for events: err: %v", err)
	}
	err = c.eventDataPersister.SaveEvents(pastEvents)
	if err != nil {
		return fmt.Errorf("Error persisting events: err: %v", err)
	}
	err = c.persistRetrieverLastBlockData()
	if err != nil {
		return fmt.Errorf("Error persisting last block data: err: %v", err)
	}

	err = c.startListener()
	if err != nil {
		return fmt.Errorf("Error starting listener: err: %v", err)
	}
	defer func() {
		err = c.StopCollection()
		if err != nil {
			log.Errorf("Error stopping collection: err: %v", err)
		}
	}()

	c.quitChan = make(chan interface{})
	// errors channel to catch persistence errors
	errorsChan := make(chan error)

	go func(quit <-chan interface{}, errors chan<- error) {
		for {
			select {
			case event := <-c.listen.EventRecvChan:
				if log.V(2) {
					log.Infof(
						"event received: %v, %v, %v, \n%v",
						event.EventType(),
						event.Hash(),
						event.Timestamp(),
						event.LogPayloadToString(),
					)
				}
				err = c.updateEventTimeFromBlockHeader(event)
				if err != nil {
					err = fmt.Errorf("Error updating date for event: err: %v", err)
					errors <- err
					return
				}
				// Save event to persister
				err = c.eventDataPersister.SaveEvents([]*model.Event{event})
				if err != nil {
					err = fmt.Errorf("Error saving events: err: %v", err)
					errors <- err
					return
				}
				// Update last block in persistence in case of error
				err = c.listenerPersister.UpdateLastBlockData([]*model.Event{event})
				if err != nil {
					err = fmt.Errorf("Error updating last block: err: %v", err)
					errors <- err
					return
				}
				// Call event triggers
				err = c.callTriggers(event)
				if err != nil {
					log.Errorf("Error calling triggers: err: %v", err)
				}
			case <-quit:
				return
			}
		}
	}(c.quitChan, errorsChan)

	select {
	case err = <-errorsChan:
		return fmt.Errorf("Error during event handling: err: %v", err)
	case <-c.quitChan:
		return nil
	}
}

// StopCollection is for stopping the listener
func (c *EventCollector) StopCollection() error {
	var err error
	if c.listen != nil {
		err = c.listen.Stop()
	}
	if c.quitChan != nil {
		close(c.quitChan)
	}
	return err
}

// AddWatchers will add watchers to the embedded listener.
func (c *EventCollector) AddWatchers(w model.ContractWatchers) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	return c.listen.AddWatchers(w)
}

// RemoveWatchers will remove given watcher from the embedded listener.
func (c *EventCollector) RemoveWatchers(w model.ContractWatchers) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	return c.listen.RemoveWatchers(w)
}

// UpdateStartingBlocks updates starting blocks for retriever based on persistence
func (c *EventCollector) updateRetrieverStartingBlocks() {
	for _, filter := range c.filterers {
		contractAddress := filter.ContractAddress()
		eventTypes := filter.EventTypes()
		for _, eventType := range eventTypes {
			lastBlock := c.retrieverPersister.LastBlockNumber(eventType, contractAddress)
			// If lastBlock is 0, assume it has never been set, so set to default
			// start block value.
			if lastBlock == 0 {
				lastBlock = c.startBlock
			}
			// NOTE (IS): Starting at lastBlock+1. There could be a scenario where this could miss the rest of events in prev block?
			filter.UpdateStartBlock(eventType, lastBlock+1)
		}
	}
}

func (c *EventCollector) updateEventTimesFromBlockHeaders(events []*model.Event) error {
	for _, event := range events {
		err := c.updateEventTimeFromBlockHeader(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *EventCollector) updateEventTimeFromBlockHeader(event *model.Event) error {
	var header *types.Header
	var err error

	inCache := false
	if c.headerCache == nil {
		c.headerCache = utils.NewBlockHeaderCache(blockHeaderExpirySecs)
	} else {
		header = c.headerCache.HeaderByBlockNumber(event.BlockNumber())
		if header != nil {
			inCache = true
		}
	}
	if !inCache {
		blockNum := big.NewInt(0)
		blockNum.SetUint64(event.BlockNumber())
		header, err = c.chain.HeaderByNumber(context.Background(), blockNum)
		if err != nil {
			return err
		}
		c.headerCache.AddHeader(event.BlockNumber(), header)
	}
	if err != nil {
		return err
	}
	event.SetTimestamp(header.Time.Int64())
	return nil
}

func (c *EventCollector) retrieveEvents() error {
	c.updateRetrieverStartingBlocks()
	c.retrieve = retriever.NewEventRetriever(c.client, c.filterers)
	err := c.retrieve.Retrieve()
	if err != nil {
		return err
	}
	err = c.retrieve.SortEventsByBlock()
	return err
}

func (c *EventCollector) startListener() error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	c.listen = listener.NewEventListener(c.client, c.watchers)
	if c.listen == nil {
		return errors.New("Listener should not be nil")
	}
	err := c.listen.Start()
	if err != nil {
		return fmt.Errorf("Listener should have started with no errors: %v", err)
	}
	return nil
}

func (c *EventCollector) callTriggers(event *model.Event) error {
	var err error
	for _, trigger := range c.triggers {
		if trigger.ShouldRun(c, event) {
			err = trigger.Run(c, event)
		}
	}
	return err
}

// persistRetrieverLastBlockData saves the last seen events for each filter to
// persistence. Returns the last error seen when updating the block data.
func (c *EventCollector) persistRetrieverLastBlockData() error {
	var err error
	for _, filter := range c.filterers {
		err = c.retrieverPersister.UpdateLastBlockData(filter.LastEvents())
	}
	return err
}
