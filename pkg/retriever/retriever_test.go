// Package retriever_test contains tests for retriever package
package retriever_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cutils "github.com/joincivil/civil-events-crawler/pkg/contractutils"
	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/generated/filterer"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"math/big"
	"testing"
)

// Using rinkeby for now
// TODO(IS) change to simulated backend, write more tests
const (
	testTCRAddress = "0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d"
)

// TestEventCollection tests that events are being collected,
func TestEventCollection(t *testing.T) {
	client, err := cutils.SetupRinkebyClient()
	if err != nil {
		t.Errorf("Error connecting to rinkeby: %v", err)
	}
	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(common.HexToAddress(testTCRAddress)),
	}
	retrieve := retriever.NewCivilEventRetriever(client, filterers)
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
	client, err := cutils.SetupRinkebyClient()
	if err != nil {
		t.Errorf("Error connecting to rinkeby: %v", err)
	}
	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(common.HexToAddress(testTCRAddress)),
	}
	retrieve := retriever.NewCivilEventRetriever(client, filterers)
	model1, _ := model.NewCivilEventFromContractEvent("ApplicationWhitelisted", "CivilTCRContract", common.HexToAddress(testTCRAddress),
		testEvent1, utils.CurrentEpochSecsInInt())
	model2, _ := model.NewCivilEventFromContractEvent("Application", "CivilTCRContract", common.HexToAddress(testTCRAddress), testEvent2,
		utils.CurrentEpochSecsInInt())
	if err != nil {
		t.Errorf("Error connecting to rinkeby: %v", err)
	}
	retrieve.PastEvents = append(retrieve.PastEvents, *model1, *model2)
	err = retrieve.SortEventsByBlock()
	if err != nil {
		t.Error("Sorting didn't happen")
	}
}

// Check last events. TODO: make better test here w simulated backend
func TestLastEvents(t *testing.T) {
	client, err := cutils.SetupRinkebyClient()
	if err != nil {
		t.Errorf("Error connecting to rinkeby: %v", err)
	}
	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(common.HexToAddress(testTCRAddress)),
	}
	retrieve := retriever.NewCivilEventRetriever(client, filterers)
	if len(filterers[0].LastEvents()) != 0 {
		t.Error("LastEvents should be empty")
	}
	err = retrieve.Retrieve()
	if err != nil {
		t.Errorf("Error retrieving events: %v", err)
	}
	if len(filterers[0].LastEvents()) == 0 {
		t.Error("LastEvents should not be empty")
	}
}
