// Package retriever contains all components for events retriever, which
// gets all past events
package retriever // import "github.com/joincivil/civil-events-crawler/pkg/retriever"

import (
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

	// filterers contains a list of ContractFilterers
	filterers []model.ContractFilterers
}

// Retrieve gets all events from StartBlock until now
func (r *CivilEventRetriever) Retrieve() error {
	for _, filter := range r.filterers {
		err, pastEvents := filter.StartFilterers(r.client, r.PastEvents)
		if err != nil {
			return err
		}
		r.PastEvents = pastEvents

	}
	return nil
}

// GetBlockNumber is a helper function to get the block number of an event
func (r *CivilEventRetriever) GetBlockNumber(event model.CivilEvent) (uint64, error) {
	payload := event.Payload()
	eventHash := event.Hash()
	rawPayload, ok := payload.Value("Raw")
	if !ok {
		return uint64(0), fmt.Errorf("Can't get raw value for %v", eventHash)
	}
	rawPayloadLog, ok := rawPayload.Log()
	if !ok {
		return uint64(0), fmt.Errorf("Can't get log field of raw value for %v", eventHash)
	}
	return rawPayloadLog.BlockNumber, nil
}

// SortEventsByBlock sorts events in PastEvents by block number
// NOTE(IS): This is not optimal, but for now checking that values exist outside of sort
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
