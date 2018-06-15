// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/joincivil/civil-events-crawler/pkg/listener"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
)

// NewCivilEventCollector creates a new civil event collector
func NewCivilEventCollector(client bind.ContractBackend, filterers []model.ContractFilterers,
	watchers []model.ContractWatchers, retrieverPersister model.RetrieverMetaDataPersister,
	listenerPersister model.ListenerMetaDataPersister, listen listener.CivilEventListener) *CivilEventCollector {

	eventcollector := &CivilEventCollector{
		client:             client,
		filterers:          filterers,
		watchers:           watchers,
		retrieverPersister: retrieverPersister,
		listenerPersister:  listenerPersister,
		listen:             listen,
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

	listen listener.CivilEventListener
}

// StartCollection contains logic to run retriever and listener.
func (c *CivilEventCollector) StartCollection() error {
	c.updateStartingBlocks()
	retrieve := retriever.NewCivilEventRetriever(c.client, c.filterers)
	retrieve.Retrieve()
	retrieve.SortEventsByBlock()
	// Here should update where the retrieving left off
	err := c.listen.Start()
	if err != nil {
		return err
	}
	return nil
}

// StopCollection stops listener
func (c *CivilEventCollector) StopCollection() error {
	// TODO: update for last blocks of each event type and put all events together
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
