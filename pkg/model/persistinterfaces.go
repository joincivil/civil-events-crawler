// Package model contains the general data models and interfaces.
package model // import "github.com/joincivil/civil-events-crawler/pkg/model"

import (
	"github.com/ethereum/go-ethereum/common"
)

// ListenerMetaDataPersister handles storing any metadata related to running
// the listener.
type ListenerMetaDataPersister interface {
	// LastBlockNumber returns the last block number seen by the listener for
	// an event type and contract address
	LastBlockNumber(eventType string, contractAddress common.Address) uint64

	// LastBlockHash returns the last block hash seen by the listener for an
	// event type and contract address
	LastBlockHash(eventType string, contractAddress common.Address) common.Hash

	// UpdateLastBlockData should update the last block data from the Event(s)
	UpdateLastBlockData(events []*Event) error
}

// RetrieverMetaDataPersister handles storing any metadata related to running
// the retriever.
type RetrieverMetaDataPersister interface {
	// LastBlockNumber returns the last block number seen by the retriever for
	// an event type and contract address
	LastBlockNumber(eventType string, contractAddress common.Address) uint64

	// LastBlockHash returns the last block hash seen by the retriever for an event
	// type and contract address
	LastBlockHash(eventType string, contractAddress common.Address) common.Hash

	// UpdateLastBlockData should update the last block Number from the Event(s)
	UpdateLastBlockData(events []*Event) error
}

// EventDataPersister handles storing the received Event data.
type EventDataPersister interface {
	// SaveEvents stores a list of Event(s)
	SaveEvents(events []*Event) error

	// RetrieveEvents retrieves the Events from the persistence layer based
	// on date in which it was received
	// TODO: We will not query the events table for this, still to implement
	// RetrieveEvents(offset uint, count uint, reverse bool) ([]*Event, error)
}
