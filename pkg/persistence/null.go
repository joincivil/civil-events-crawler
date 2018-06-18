// Package persistence are implementations of the model/persisttypes interfaces
package persistence

import (
	"github.com/ethereum/go-ethereum/common"
	log "github.com/golang/glog"

	"github.com/joincivil/civil-events-crawler/pkg/model"
)

// NullPersister is a persister that does not save any values and always returns
// defaults for interface methods. Handy for testing and for crawler single use scenarios.
// Implements the ListenerMetaDataPersister, RetrieverMetaDataRetriever, and EventsDataPersister.
// Acts as the 'none' configured PersisterType.
type NullPersister struct{}

// LastBlockNumber returns the last block number seen by the listener for
// an event type and contract address
func (n *NullPersister) LastBlockNumber(eventType string, contractAddress common.Address) uint64 {
	return uint64(0)
}

// LastBlockHash returns the last block hash seen by the listener for an
// event type and contract address
func (n *NullPersister) LastBlockHash(eventType string, contractAddress common.Address) common.Hash {
	return common.Hash{}
}

// UpdateLastBlockData should update the last block data from the CivilEvent(s)
func (n *NullPersister) UpdateLastBlockData(events []model.CivilEvent) error {
	// Only log this if INFO log level -v=2
	if log.V(2) {
		for _, event := range events {
			rawVal, _ := event.Payload().Value("Raw")
			eventLog, _ := rawVal.Log()
			log.Infof(
				"NullPersister: UpdatedLastBlockData: event: type: %v, addr: %v, blknum: %v\n",
				event.EventType(),
				event.ContractAddress().Hex(),
				eventLog.BlockNumber,
			)
		}
	}
	return nil
}

// SaveEvents stores a list of CivilEvent(s)
func (n *NullPersister) SaveEvents(events []model.CivilEvent) error {
	// Only log this if INFO log level -v=2
	if log.V(2) {
		for _, event := range events {
			rawVal, _ := event.Payload().Value("Raw")
			eventLog, _ := rawVal.Log()
			log.Infof(
				"NullPersister: SaveEvents: event: type: %v, addr: %v, blknum: %v\n",
				event.EventType(),
				event.ContractAddress().Hex(),
				eventLog.BlockNumber,
			)
		}
	}
	return nil
}

// RetrieveEvents retrieves the CivilEvents from the persistence layer based
// on date in which it was received
func (n *NullPersister) RetrieveEvents(offset uint, count uint, reverse bool) ([]*model.CivilEvent, error) {
	return []*model.CivilEvent{}, nil
}
