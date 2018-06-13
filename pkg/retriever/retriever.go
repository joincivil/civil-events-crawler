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

// NewCivilEventRetriever creates a CivilEventRetriever given a ContractFilterer
// connection to client and startBlock.
// Because contract name is now abstracted, we need startBlockToEvent to be a map from
// filterer as well
func NewCivilEventRetriever(client bind.ContractBackend, // startBlockToEvent [string]uint64,
	filterers []model.ContractFilterers) *CivilEventRetriever {
	retriever := &CivilEventRetriever{
		client:     client,
		PastEvents: make([]model.CivilEvent, 0),
		// StartBlock: startBlockToEvent,
		filterers: filterers,
	}
	return retriever
}

// CivilEventRetriever handles the iterator returned from retrieving past events
type CivilEventRetriever struct {

	// Client is the ethereum client from go-ethereum
	client bind.ContractBackend

	// PastEvents is a slice that holds all past CivilEvents requested
	PastEvents []model.CivilEvent

	// // StartBlock is the block number from where PastEvents were scraped from
	// StartBlockToEvent map[string]uint64

	filterers []model.ContractFilterers
}

// Retrieve gets all events from StartBlock until now
func (r *CivilEventRetriever) Retrieve() error {
	var err error

	for _, filterer := range r.filterers {
		err = filterer.StartFilterers(
			r.client,
			&r.PastEvents,
			// r.StartBlockToEvent,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// SortEventsByBlock sorts events in PastEvents by block number
func (r *CivilEventRetriever) SortEventsByBlock() bool {
	pastEvents := r.PastEvents
	var errVar string
	sort.Slice(pastEvents, func(i, j int) bool {
		payload1 := pastEvents[i].Payload()
		eventHash1 := pastEvents[i].Hash()
		payload2 := pastEvents[j].Payload()
		eventHash2 := pastEvents[j].Hash()

		rawPayload1, ok := payload1.Value("Raw")
		if !ok {
			errVar = fmt.Sprintf("Can't get raw value for %v", eventHash1)
			log.Error(errVar)
		}
		rawPayload2, ok := payload2.Value("Raw")
		if !ok {
			errVar = fmt.Sprintf("Can't get raw value for %v", eventHash2)
			log.Error(errVar)
		}
		rawPayloadLog1, ok := rawPayload1.Log()
		if !ok {
			errVar = fmt.Sprintf("Can't get raw value for %v", eventHash1)
			log.Error(errVar)
		}
		rawPayloadLog2, ok := rawPayload2.Log()
		if !ok {
			errVar = fmt.Sprintf("Can't get raw value for %v", eventHash2)
			log.Error(errVar)
		}
		// BlockNumber is 0 when the value isn't there.
		blockNumber1 := rawPayloadLog1.BlockNumber
		if blockNumber1 == 0 {
			errVar = fmt.Sprintf("Can't get block number for %v", eventHash1)
			log.Error(errVar)
		}
		blockNumber2 := rawPayloadLog2.BlockNumber
		if blockNumber2 == 0 {
			errVar = fmt.Sprintf("Can't get block number for %v", eventHash2)
			log.Error(errVar)
		}
		return rawPayloadLog1.BlockNumber < rawPayloadLog2.BlockNumber
	})
	return errVar == ""
}
