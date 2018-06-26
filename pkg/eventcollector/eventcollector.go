// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	"errors"
	"fmt"
	log "github.com/golang/glog"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/joincivil/civil-events-crawler/pkg/listener"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
)

// NewCivilEventCollector creates a new civil event collector
func NewCivilEventCollector(client bind.ContractBackend, filterers []model.ContractFilterers,
	watchers []model.ContractWatchers, retrieverPersister model.RetrieverMetaDataPersister,
	listenerPersister model.ListenerMetaDataPersister, eventDataPersister model.EventDataPersister,
	triggers []Trigger) *CivilEventCollector {

	eventcollector := &CivilEventCollector{
		client:             client,
		filterers:          filterers,
		watchers:           watchers,
		retrieverPersister: retrieverPersister,
		listenerPersister:  listenerPersister,
		eventDataPersister: eventDataPersister,
		triggers:           triggers,
	}
	return eventcollector
}

// CivilEventCollector handles logic for getting historical and live events
type CivilEventCollector struct {
	client bind.ContractBackend

	triggers []Trigger

	filterers []model.ContractFilterers

	watchers []model.ContractWatchers

	retrieverPersister model.RetrieverMetaDataPersister

	listenerPersister model.ListenerMetaDataPersister

	eventDataPersister model.EventDataPersister

	listen *listener.CivilEventListener

	retrieve *retriever.CivilEventRetriever

	// quitChan is created in StartCollection() and stops the goroutine listening for events.
	quitChan chan interface{}

	mutex sync.Mutex
}

// StartCollection contains logic to run retriever and listener.
func (c *CivilEventCollector) StartCollection() error {
	err := c.retrieveEvents()
	if err != nil {
		return err
	}
	pastEvents := c.retrieve.PastEvents
	err = c.eventDataPersister.SaveEvents(pastEvents)
	if err != nil {
		return err
	}
	err = c.persistRetrieverLastBlockData()
	if err != nil {
		return err
	}

	err = c.startListener()
	if err != nil {
		return err
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
						event.Payload().ToString(),
					)
				}
				// Save event to persister
				err = c.eventDataPersister.SaveEvents([]model.CivilEvent{event})
				if err != nil {
					errors <- err
					return
				}
				// Update last block in persistence in case of error
				err = c.listenerPersister.UpdateLastBlockData([]model.CivilEvent{event})
				if err != nil {
					errors <- err
					return
				}
				// Call event triggers
				err = c.callTriggers(&event)
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
		return err
	case <-c.quitChan:
		return nil
	}
}

// StopCollection is for stopping the listener
func (c *CivilEventCollector) StopCollection() error {
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
func (c *CivilEventCollector) AddWatchers(w model.ContractWatchers) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	return c.listen.AddWatchers(w)
}

// RemoveWatchers will remove given watcher from the embedded listener.
func (c *CivilEventCollector) RemoveWatchers(w model.ContractWatchers) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	return c.listen.RemoveWatchers(w)
}

// UpdateStartingBlocks updates starting blocks for retriever based on persistence
func (c *CivilEventCollector) updateRetrieverStartingBlocks() {
	for _, filter := range c.filterers {
		contractAddress := filter.ContractAddress()
		eventTypes := filter.EventTypes()
		for _, eventType := range eventTypes {
			lastBlock := c.retrieverPersister.LastBlockNumber(eventType, contractAddress)
			filter.UpdateStartBlock(eventType, lastBlock)
		}
	}
}

func (c *CivilEventCollector) retrieveEvents() error {
	c.updateRetrieverStartingBlocks()
	c.retrieve = retriever.NewCivilEventRetriever(c.client, c.filterers)
	err := c.retrieve.Retrieve()
	if err != nil {
		return err
	}
	err = c.retrieve.SortEventsByBlock()
	return err
}

func (c *CivilEventCollector) startListener() error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	c.listen = listener.NewCivilEventListener(c.client, c.watchers)
	if c.listen == nil {
		return errors.New("Listener should not be nil")
	}
	err := c.listen.Start()
	if err != nil {
		return fmt.Errorf("Listener should have started with no errors: %v", err)
	}
	return nil
}

func (c *CivilEventCollector) callTriggers(event *model.CivilEvent) error {
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
func (c *CivilEventCollector) persistRetrieverLastBlockData() error {
	var err error
	for _, filter := range c.filterers {
		err = c.retrieverPersister.UpdateLastBlockData(filter.LastEvents())
	}
	return err
}
