// Package retriever contains all components for events retriever, which
// gets all past events
package retriever // import "github.com/joincivil/civil-events-crawler/pkg/retriever"

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joincivil/civil-events-crawler/pkg/generated/retrieve"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"sort"
)

// NewCivilEventRetriever creates a CivilEventRetriever given a contract address
// connection to client and startBlock. Logic should go in main script to
// check startblock of last event?
func NewCivilEventRetriever(client bind.ContractBackend, contractAddress string, startBlock int) *CivilEventRetriever {
	address := common.HexToAddress(contractAddress)
	retriever := &CivilEventRetriever{
		Client:             client,
		ContractAddress:    address,
		ContractAddressStr: contractAddress,
		PastEvents:         make([]model.CivilEvent, 0),
		StartBlock:         uint64(startBlock),
	}
	return retriever
}

// CivilEventRetriever handles the iterator returned from retrieving past events
type CivilEventRetriever struct {

	// Client is the ethereum client from go-ethereum
	Client bind.ContractBackend

	// ContractAddress is the Address type for the contract to watch
	ContractAddress common.Address

	// ContractAddressStr is the string repr for the address of the contract
	ContractAddressStr string

	// PastEvents is a slice that holds all past CivilEvents requested
	PastEvents []model.CivilEvent

	// StartBlock is the block number from where PastEvents were scraped from
	StartBlock uint64
}

// Retrieve gets all events from StartBlock until now
func (r *CivilEventRetriever) Retrieve() error {

	// RetrieveCivilTCRContractEvents is generated
	err := retrieve.RetrieveCivilTCRContractEvents(
		r.Client,
		r.ContractAddress,
		&r.PastEvents,
		r.StartBlock,
	)
	if err != nil {
		return err
	}

	err = retrieve.RetrieveNewsroomContractEvents(
		r.Client,
		r.ContractAddress,
		&r.PastEvents,
		r.StartBlock,
	)
	if err != nil {
		return err
	}

	return nil
}

// SortEvents sorts events in PastEvents by block number
func (r *CivilEventRetriever) SortEvents() error {
	pastEvents := r.PastEvents
	sort.Slice(pastEvents, func(i, j int) bool {
		rawPayload1, ok := pastEvents[i].Payload.Value("Raw")
		rawPayload2, ok := pastEvents[j].Payload.Value("Raw")
		if !ok {
			fmt.Println("Can't get raw value from event")
		}
		rawPayloadLog1, ok := rawPayload1.Log()
		rawPayloadLog2, ok := rawPayload2.Log()
		if !ok {
			fmt.Println("Can't get raw value from event")
		}
		return rawPayloadLog1.BlockNumber < rawPayloadLog2.BlockNumber
	})
	r.PastEvents = pastEvents
	return nil
}
