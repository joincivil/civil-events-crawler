// Package retriever contains all components for events retriever, which
// gets all past events
package retriever // import "github.com/joincivil/civil-events-crawler/pkg/retriever"

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/golang/glog"
	"github.com/joincivil/civil-events-crawler/pkg/generated/tcr"
	"github.com/joincivil/civil-events-crawler/pkg/model"
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
	civilTCR, err := tcr.NewCivilTCRContract(r.ContractAddress, r.Client)
	if err != nil {
		log.Errorf("Error initializing TCR: %v", err)
		return err
	}

	var opts = &bind.FilterOpts{
		Start: r.StartBlock,
	}

	err = RetrieveApplication(opts, civilTCR, &r.PastEvents)
	if err != nil {
		return fmt.Errorf("Error getting past _Application events: %v", err)
	}

	err = RetrieveApplicationRemoved(opts, civilTCR, &r.PastEvents)
	if err != nil {
		return fmt.Errorf("Error getting past _ApplicationRemoved events: %v", err)
	}

	err = RetrieveApplicationWhitelisted(opts, civilTCR, &r.PastEvents)
	if err != nil {
		return fmt.Errorf("Error getting past _ApplicationWhitelisted events: %v", err)
	}

	err = RetrieveChallenge(opts, civilTCR, &r.PastEvents)
	if err != nil {
		return fmt.Errorf("Error getting past _Challenge events: %v", err)
	}

	return nil
}
