package postgres_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"math/big"
	"testing"
)

var (
	contractAddress = "0x77e5aaBddb760FBa989A1C4B2CDd4aA8Fa3d311d"
	testAddress     = "0xDFe273082089bB7f70Ee36Eebcde64832FE97E55"
	testEvent       = &contract.CivilTCRContractApplication{
		ListingAddress: common.HexToAddress(testAddress),
		Deposit:        big.NewInt(1000),
		AppEndDate:     big.NewInt(1653860896),
		Data:           "DATA",
		Applicant:      common.HexToAddress(testAddress),
		Raw: types.Log{
			Address:     common.HexToAddress(testAddress),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: 8888888,
			TxHash:      common.Hash{},
			TxIndex:     2,
			BlockHash:   common.Hash{},
			Index:       2,
			Removed:     false,
		},
	}
)

func setupDBEventFromEvent(civilEvent *model.Event) (*postgres.Event, error) {
	return postgres.NewDbEventFromEvent(civilEvent)
}

func setupEvent() (*model.Event, error) {
	return model.NewEventFromContractEvent("Application", "CivilTCRContract", common.HexToAddress(contractAddress),
		testEvent, utils.CurrentEpochSecsInInt64(), model.Filterer)
}

func setupDBEvent() (*postgres.Event, error) {
	civilEvent, err := setupEvent()
	if err != nil {
		return &postgres.Event{}, fmt.Errorf("setupEvent should have succeeded: err: %v", err)
	}
	dbEvent, err := setupDBEventFromEvent(civilEvent)
	if err != nil {
		return dbEvent, fmt.Errorf("setupDBEventFromEvent should have succeeded: err: %v", err)
	}
	return dbEvent, nil
}

func TestDBEventSetup(t *testing.T) {
	dbEvent, err := setupDBEvent()
	if err != nil {
		t.Errorf("%v", err)
	}
	if dbEvent == nil {
		t.Error("postgres Event should not be nil")
	}
	if dbEvent.EventType != "Application" {
		t.Errorf("EventType wasn't set correctly: %v", dbEvent.EventType)
	}
	if len(dbEvent.EventHash) == 0 {
		t.Errorf("EventHash wasn't set correctly: %v", dbEvent.EventHash)
	}
	if contractAddress != dbEvent.ContractAddress {
		t.Errorf("ContractAddress wasn't set correctly: %v", dbEvent.ContractAddress)
	}
	if "CivilTCRContract" != dbEvent.ContractName {
		t.Errorf("ContractName wasn't set correctly: %v", dbEvent.ContractName)
	}
	if dbEvent.Timestamp <= 0 {
		t.Errorf("Timestamp was not init correctly: %v", dbEvent.Timestamp)
	}
	if dbEvent.RetrievalMethod != 0 {
		t.Errorf("Retrieval method should be 0 for filterer but it is %v", dbEvent.RetrievalMethod)
	}
	if len(dbEvent.EventPayload) != 5 {
		t.Errorf("EventPayload was not init correctly: %v", dbEvent.EventPayload)
	}
	if len(dbEvent.LogPayload) != 9 {
		t.Errorf("EventPayload was not init correctly: %v", dbEvent.EventPayload)
	}

}
