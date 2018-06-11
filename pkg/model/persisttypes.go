// Package model contains the general data models and interfaces for the Civil crawler.
package model // import "github.com/joincivil/civil-events-crawler/pkg/model"

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ListenerMetaDataPersister handles storing any metadata related to running
// the listener.
type ListenerMetaDataPersister interface {
	// LastBlockNumber returns the last block number seen by the listener
	LastBlockNumber() uint64

	// LastBlockHash returns the last block hash seen by the listener
	LastBlockHash() common.Hash

	// UpdateDataFromRawLog should update the last block data from the raw event
	// types.Log
	UpdateDataFromRawLog(rawLog types.Log) error
}

// RetrieverMetaDataPersister handles storing any metadata related to running
// the listener.
type RetrieverMetaDataPersister interface {
	// LastBlockNumber returns the last block number seen by the retriever
	LastBlockNumber() uint64

	// LastBlockHash returns the last block hash seen by the retriever
	LastBlockHash() common.Hash

	// UpdateDataFromRawLog should update the last block data from the raw event
	// types.Log
	UpdateDataFromRawLog(rawLog types.Log) error
}

// EventDataPersister handles storing the received CivilEvent data.
type EventDataPersister interface {
	// SaveEvents stores the event
	SaveEvent(event *CivilEvent) error

	// RetrieveEvents retrieves the CivilEvents from the persistence layer based
	// on date in which it was received
	RetrieveEvents(offset uint, count uint, reverse bool) ([]*CivilEvent, error)
}
