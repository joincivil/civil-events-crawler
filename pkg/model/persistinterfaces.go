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

// RetrieveEventsCriteria contains the retrieval criteria for a RetrieveEvents
// query.
type RetrieveEventsCriteria struct {
	Hash            string   `db:"hash"`
	ContractAddress string   `db:"contract_address"`
	Offset          int      `db:"offset"`
	Count           int      `db:"count"`
	ExcludeHashes   []string `db:"exclude_hashes"`
	// Reverse reverses by id in DB
	Reverse   bool   `db:"reverse"`
	FromTs    int64  `db:"fromts"`
	BeforeTs  int64  `db:"beforets"`
	EventType string `db:"eventtype"`
}

// EventDataPersister handles storing the received Event data.
type EventDataPersister interface {
	// SaveEvents stores a list of Event(s)
	SaveEvents(events []*Event) error

	// RetrieveEvents retrieves the Events from the persistence layer based
	// on date in which it was received
	RetrieveEvents(criteria *RetrieveEventsCriteria) ([]*Event, error)
}
