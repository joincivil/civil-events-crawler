// Package retriever_test contains tests for retriever package
package retriever_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/golang/glog"
	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
	"math/big"
	"testing"
)

// Using rinkeby for now
// TODO(IS) change to simulated backend, write more tests
const (
	rinkebyAddress = "https://rinkeby.infura.io"
	testTCRAddress = "0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d"
	startBlock     = 2335623
)

func setupRinkebyClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(rinkebyAddress)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}
	return client, nil
}

// TestEventCollection tests that events are being collected,
func TestEventCollection(t *testing.T) {
	client, err := setupRinkebyClient()
	if err != nil {
		t.Errorf("Error connecting to rinkeby: %v", err)
	}
	retrieve := retriever.NewCivilEventRetriever(client, testTCRAddress, startBlock)
	err = retrieve.Retrieve()
	if err != nil {
		t.Errorf("Error retrieving events: %v", err)
	}
	pastEvents := retrieve.PastEvents
	if len(pastEvents) == 0 {
		t.Error("No events collected")
	}
}

// TestSorting tests that sorting is happening by block number
func TestSorting(t *testing.T) {
	testEvent1 := &contract.CivilTCRContractApplicationWhitelisted{
		ListingAddress: common.HexToAddress(testTCRAddress),
		Raw: types.Log{
			Address:     common.HexToAddress(testTCRAddress),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: 8888888,
		},
	}
	testEvent2 := &contract.CivilTCRContractApplication{
		ListingAddress: common.HexToAddress(testTCRAddress),
		Deposit:        big.NewInt(1000),
		AppEndDate:     big.NewInt(1653860896),
		Data:           "DATA",
		Applicant:      common.HexToAddress(testTCRAddress),
		Raw: types.Log{
			Address:     common.HexToAddress(testTCRAddress),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: 8888886,
		},
	}
	client, err := setupRinkebyClient()
	if err != nil {
		t.Errorf("Error connecting to rinkeby: %v", err)
	}
	retrieve := retriever.NewCivilEventRetriever(client, testTCRAddress, startBlock)
	model1 := model.NewCivilEvent("ApplicationWhitelisted", testEvent1)
	model2 := model.NewCivilEvent("Application", testEvent2)
	retrieve.PastEvents = append(retrieve.PastEvents, *model1, *model2)
	retrieve.SortEvents()
	if retrieve.PastEvents[0].EventType != "Application" {
		t.Error("Sorting didn't happen")
	}

}
