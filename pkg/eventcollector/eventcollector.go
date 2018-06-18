// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	log "github.com/golang/glog"
	"github.com/joincivil/civil-events-crawler/pkg/listener"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
)

// NewCivilEventCollector creates a new civil event collector
func NewCivilEventCollector(client bind.ContractBackend, filterers []model.ContractFilterers,
	watchers []model.ContractWatchers, retrieverPersister model.RetrieverMetaDataPersister,
	listenerPersister model.ListenerMetaDataPersister, eventDataPersister model.EventDataPersister) *CivilEventCollector {

	eventcollector := &CivilEventCollector{
		client:             client,
		filterers:          filterers,
		watchers:           watchers,
		retrieverPersister: retrieverPersister,
		listenerPersister:  listenerPersister,
	}
	return eventcollector
}

// CivilEventCollector handles logic for getting historical and live events
type CivilEventCollector struct {
	client bind.ContractBackend

	filterers []model.ContractFilterers

	watchers []model.ContractWatchers

	retrieverPersister model.RetrieverMetaDataPersister

	listenerPersister model.ListenerMetaDataPersister

	eventDataPersister model.EventDataPersister

	listen *listener.CivilEventListener

	retrieve *retriever.CivilEventRetriever
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
	defer c.stopListener()

	quitChan := make(chan interface{})
	// errors channel to catch persistence errors
	errors := make(chan error)

	go func(quit <-chan interface{}, errors chan<- error) {
		for {
			select {
			case event := <-c.listen.EventRecvChan:
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
			case <-quit:
				return
			}
		}
	}(quitChan, errors)

	for err = range errors {
		close(quitChan)
		return err
	}
	return err
}

// StopCollection is for stopping the listener
func (c *CivilEventCollector) StopCollection() error {
	err := c.listen.Stop()
	return err
}

func (c *CivilEventCollector) stopListener() {
	err := c.listen.Stop()
	if err != nil {
		log.Errorf("Error stopping listener, %v", err)
	}
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

// endRetrieving saves the last seen events for each filter to persistence
func (c *CivilEventCollector) persistRetrieverLastBlockData() error {
	var err error
	for _, filter := range c.filterers {
		err = c.retrieverPersister.UpdateLastBlockData(filter.LastEvents())
		return err
	}
	return err
}
