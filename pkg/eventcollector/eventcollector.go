// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	// Client is the ethereum client from go-ethereum
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
	err := c.setupRetriever()
	if err != nil {
		return err
	}
	pastEvents := c.retrieve.PastEvents
	err = c.eventDataPersister.SaveEvents(pastEvents)
	if err != nil {
		return err
	}
	err = c.endRetrieving()
	if err != nil {
		return err
	}

	err = c.setupListener()
	if err != nil {
		return err
	}
	defer c.listen.Stop()

	eventRecv := make(chan bool)
	// errors channel to catch persistence errors
	errors := make(chan error)

	go func(recv chan<- bool, errors chan<- error) {
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
				recv <- true
			}
		}
	}(eventRecv, errors)

	// If we get an error in persistence, we should break from the loop and return error.
Loop:
	for {
		select {
		case <-eventRecv:
			continue
		case err = <-errors:
			break Loop
		}
	}

	return err
}

// StopCollection is for stopping the listener
func (c *CivilEventCollector) StopCollection() error {
	err := c.listen.Stop()
	return err
}

// UpdateStartingBlocks updates starting blocks for retriever based on persistence
func (c *CivilEventCollector) updateStartingBlocks() {
	for _, filter := range c.filterers {
		contractAddress := filter.ContractAddress()
		eventTypes := filter.EventTypes()
		for _, eventType := range eventTypes {
			lastBlock := c.retrieverPersister.LastBlockNumber(eventType, contractAddress)
			filter.UpdateStartBlock(eventType, lastBlock)
		}
	}
}

func (c *CivilEventCollector) setupRetriever() error {
	c.updateStartingBlocks()
	c.retrieve = retriever.NewCivilEventRetriever(c.client, c.filterers)
	err := c.retrieve.Retrieve()
	if err != nil {
		return err
	}
	err = c.retrieve.SortEventsByBlock()
	return err
}

func (c *CivilEventCollector) setupListener() error {
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
func (c *CivilEventCollector) endRetrieving() error {
	var err error
	for _, filter := range c.filterers {
		err = c.retrieverPersister.UpdateLastBlockData(filter.LastEvents())
		return err
	}
	return err
}
