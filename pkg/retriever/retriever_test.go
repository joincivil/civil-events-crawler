// Package retriever_test contains tests for retriever package
package retriever_test

import (
	// "fmt"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core/types"
	cutils "github.com/joincivil/civil-events-crawler/pkg/contractutils"
	commongen "github.com/joincivil/civil-events-crawler/pkg/generated/common"
	"github.com/joincivil/civil-events-crawler/pkg/generated/filterer"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
	// "github.com/joincivil/civil-events-crawler/pkg/utils"
	"math/big"
	"sort"
	"testing"
)

const (
	testTCRAddress = "0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d"
)

func TestFilterersEventList(t *testing.T) {
	tcrFilterers := filterer.NewCivilTCRContractFilterers(common.HexToAddress(testTCRAddress))
	internalEvents := tcrFilterers.EventTypes()
	events := commongen.EventTypesCivilTCRContract()
	for index, event := range events {
		if event != internalEvents[index] {
			t.Error("Internal events list should match external events list")
		}
	}
}

// Setup some TCR Contract Events for testing
func setupTCRContractEvents(t *testing.T, contracts *cutils.AllTestContracts) int {
	numEvents := 3
	_, err := contracts.CivilTcrContract.Apply(contracts.Auth, contracts.NewsroomAddr, big.NewInt(5000), "")
	if err != nil {
		t.Fatalf("Application failed, error: %v", err)
	}
	contracts.Client.Commit()

	// NOTE (IS): why does gas estimation fail here?
	// this fail with "failed to estimate gas needed: gas required exceeds allowance or always failing transaction"
	// c, err := contracts.CivilTcrContract.Challenge(contracts.Auth, contracts.NewsroomAddr, "")
	// fmt.Println(c)
	// if err != nil {
	// 	t.Fatalf("Challenge failed, error: %v", err)
	// }
	// contracts.Client.Commit()

	_, err = contracts.CivilTcrContract.Deposit(contracts.Auth, contracts.NewsroomAddr, big.NewInt(5))
	if err != nil {
		t.Fatalf("Deposit failed, error: %v", err)
	}
	contracts.Client.Commit()

	_, err = contracts.CivilTcrContract.Withdraw(contracts.Auth, contracts.NewsroomAddr, big.NewInt(5))
	if err != nil {
		t.Fatalf("Withdraw failed, error: %v", err)
	}
	contracts.Client.Commit()

	return numEvents
}

// Setup some Newsroom Events for testing
func setupNewsroomContractEvents(t *testing.T, contracts *cutils.AllTestContracts) int {
	numEvents := 1
	_, err := contracts.NewsroomContract.SetName(contracts.Auth, "hey")
	if err != nil {
		t.Fatalf("NameChanged failed, error: %v", err)
	}
	contracts.Client.Commit()
	return numEvents
}

// Sets up Retriever for retrieving tests
func setupTestRetriever(t *testing.T) (*cutils.AllTestContracts, *retriever.EventRetriever, []model.ContractFilterers) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}

	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(contracts.CivilTcrAddr),
		filterer.NewNewsroomContractFilterers(contracts.NewsroomAddr),
	}
	retriever := retriever.NewEventRetriever(contracts.Client, filterers)
	return contracts, retriever, filterers
}

func setupAllEvents(t *testing.T, contracts *cutils.AllTestContracts) int {
	expectedNumTCREvents := setupTCRContractEvents(t, contracts)
	// NOTE(IS): Setting up Newsroom contract emits 3 events
	expectedNumNewsroomEvents := 3
	expectedNumNewsroomEvents += setupNewsroomContractEvents(t, contracts)
	numEvents := expectedNumTCREvents + expectedNumNewsroomEvents
	return numEvents
}

// TestFilterers tests that past events are being retrieved upon Retrieve() call
func TestRetrieveMethod(t *testing.T) {
	contracts, retriever, _ := setupTestRetriever(t)

	numEvents := setupAllEvents(t, contracts)
	err := retriever.Retrieve()
	if err != nil {
		t.Errorf("Error retrieving events: %v", err)
	}

	pastEvents := retriever.PastEvents

	if len(pastEvents) != numEvents {
		t.Errorf("Should have collected %v events but collected %v", numEvents, len(pastEvents))
	}
}

// TestSorting tests that sorting by block number works
func TestSorting(t *testing.T) {
	contracts, retriever, _ := setupTestRetriever(t)

	numEvents := setupAllEvents(t, contracts)
	err := retriever.Retrieve()
	if err != nil {
		t.Errorf("Error retrieving events: %v", err)
	}

	pastEvents := retriever.PastEvents

	if len(pastEvents) != numEvents {
		t.Errorf("Should have collected %v events but collected %v", numEvents, len(pastEvents))
	}

	err = retriever.SortEventsByBlock()
	if err != nil {
		t.Errorf("Error sorting events: %v", err)
	}
	sortedEvents := retriever.PastEvents

	blockNumbers := make([]int, numEvents)
	for idx, event := range sortedEvents {
		blockNumbers[idx] = int(event.BlockNumber())
	}
	if !sort.IntsAreSorted(blockNumbers) {
		t.Error("Events are not sorted")
	}

}

func TestLastEvents(t *testing.T) {

	contracts, retriever, filterers := setupTestRetriever(t)

	for _, filterer := range filterers {
		if len(filterer.LastEvents()) != 0 {
			t.Errorf("LastEvents should be empty, but is %v", len(filterer.LastEvents()))
		}
	}

	numEvents := setupAllEvents(t, contracts)
	err := retriever.Retrieve()
	if err != nil {
		t.Errorf("Error retrieving events: %v", err)
	}
	pastEvents := retriever.PastEvents

	// Check that past events is not empty
	if len(pastEvents) != numEvents {
		t.Errorf("Should have collected %v events but collected %v", numEvents, len(pastEvents))
	}

	for _, filterer := range filterers {
		if len(filterer.LastEvents()) == 0 {
			t.Errorf("LastEvents should be empty, but is %v", len(filterer.LastEvents()))
		}
	}
}
