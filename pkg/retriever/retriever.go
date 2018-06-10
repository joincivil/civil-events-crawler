// Package retriever contains all components for events retriever, which
// gets all past events
package retriever // import "github.com/joincivil/civil-events-crawler/pkg/retriever"

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	log "github.com/golang/glog"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"sort"
)

// NewCivilEventRetriever creates a CivilEventRetriever given a contract address
// connection to client and startBlock. Logic should go in main script to
// check startblock of last event?
func NewCivilEventRetriever(client bind.ContractBackend, startBlock int,
	filterers []model.ContractFilterers) *CivilEventRetriever {
	retriever := &CivilEventRetriever{
		client:     client,
		PastEvents: make([]model.CivilEvent, 0),
		StartBlock: uint64(startBlock),
		filterers:  filterers,
	}
	return retriever
}

// CivilEventRetriever handles the iterator returned from retrieving past events
// TODO (IS): We should pass a different StartBlock for each type of event in case of failure.
type CivilEventRetriever struct {

	// Client is the ethereum client from go-ethereum
	client bind.ContractBackend

	// PastEvents is a slice that holds all past CivilEvents requested
	PastEvents []model.CivilEvent

	// StartBlock is the block number from where PastEvents were scraped from
	StartBlock uint64

	filterers []model.ContractFilterers
}

// Retrieve gets all events from StartBlock until now
func (r *CivilEventRetriever) Retrieve() error {
	var err error

	for _, filterer := range r.filterers {
		err = filterer.StartFilterers(
			r.client,
			&r.PastEvents,
			r.StartBlock,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// SortEventsByBlock sorts events in PastEvents by block number
// TODO (IS): Maybe we should have a hash for each event so we can reference
// the event that gave us an error.
func (r *CivilEventRetriever) SortEventsByBlock() bool {
	pastEvents := r.PastEvents
	var errVar string
	// TODO (IS): Revisit this error handling.
	// var errors []string
	sort.Slice(pastEvents, func(i, j int) bool {
		rawPayload1, ok := pastEvents[i].Payload.Value("Raw")
		if !ok {
			log.Error("Can't get raw value from event")
			errVar = fmt.Sprintf("Can't get raw value for %v", pastEvents[i].EventType)
			// errors = append(errors, errVar)
		}
		rawPayload2, ok := pastEvents[j].Payload.Value("Raw")
		if !ok {
			log.Error("Can't get raw value from event")
			errVar = fmt.Sprintf("Can't get raw value for %v", pastEvents[j].EventType)
			// errors = append(errors, errVar)
		}
		rawPayloadLog1, ok := rawPayload1.Log()
		if !ok {
			log.Error("Can't convert to Log")
			errVar = fmt.Sprintf("Can't get raw value for %v", pastEvents[j].EventType)
			// errors = append(errors, errVar)
		}
		rawPayloadLog2, ok := rawPayload2.Log()
		if !ok {
			log.Error("Can't get raw value from event")
			errVar = fmt.Sprintf("Can't get raw value for %v", pastEvents[j].EventType)
			// errors = append(errors, errVar)
		}
		// BlockNumber is 0 when the value isn't there.
		blockNumber1 := rawPayloadLog1.BlockNumber
		if blockNumber1 == 0 {
			log.Error("Can't get block number from event")
			errVar = fmt.Sprintf("Can't get block number for %v", pastEvents[i].EventType)
		}
		blockNumber2 := rawPayloadLog2.BlockNumber
		if blockNumber2 == 0 {
			log.Error("Can't get block number from event")
			errVar = fmt.Sprintf("Can't get block number for %v", pastEvents[j].EventType)
		}
		return rawPayloadLog1.BlockNumber < rawPayloadLog2.BlockNumber
	})
	if errVar != "" {
		return false
	}
	return true
}
