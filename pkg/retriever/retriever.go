// Package retriever contains all components for events retriever, which
// gets all past events
package retriever // import "github.com/joincivil/civil-events-crawler/pkg/retriever"

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"sort"
)

// NewCivilEventRetriever creates a CivilEventRetriever given a list of ContractFilterers and
// connection to client
func NewCivilEventRetriever(client bind.ContractBackend, filterers []model.ContractFilterers) *CivilEventRetriever {
	retriever := &CivilEventRetriever{
		client:     client,
		PastEvents: make([]model.CivilEvent, 0),
		filterers:  filterers,
	}
	return retriever
}

// CivilEventRetriever handles the iterator returned from retrieving past events
type CivilEventRetriever struct {

	// Client is the ethereum client from go-ethereum
	client bind.ContractBackend

	// PastEvents is a slice that holds all past CivilEvents requested
	PastEvents []model.CivilEvent

	filterers []model.ContractFilterers
}

// Retrieve gets all events from StartBlock until now
func (r *CivilEventRetriever) Retrieve() error {
	for _, filter := range r.filterers {
		err := filter.StartFilterers(r.client, &r.PastEvents)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetBlockNumber is a helper function to get the block number of an event
func (r *CivilEventRetriever) GetBlockNumber(event model.CivilEvent) (uint64, error) {
	payload := event.Payload()
	eventHash := event.Hash()
	// NOTE (IS): IMO the following error handling is not necessary. they will be already thrown if
	// the hash can't be created in the event creation.
	rawPayload, ok := payload.Value("Raw")
	if !ok {
		err := fmt.Sprintf("Can't get raw value for %v", eventHash)
		return uint64(0), errors.New(err)
	}
	rawPayloadLog, ok := rawPayload.Log()
	if !ok {
		err := fmt.Sprintf("Can't get log field of raw value for %v", eventHash)
		return uint64(0), errors.New(err)
	}
	return rawPayloadLog.BlockNumber, nil
}

// SortEventsByBlock sorts events in PastEvents by block number
// NOTE(IS): This is not optimal, but for now checking that values exist outside of sort
// Also, see note on L51.
func (r *CivilEventRetriever) SortEventsByBlock() error {
	pastEvents := r.PastEvents
	for _, event := range pastEvents {
		_, err := r.GetBlockNumber(event)
		if err != nil {
			return err
		}
	}
	sort.Slice(pastEvents, func(i, j int) bool {
		blockNumber1, _ := r.GetBlockNumber(pastEvents[i])
		blockNumber2, _ := r.GetBlockNumber(pastEvents[j])
		return blockNumber1 < blockNumber2
	})
	return nil
}
