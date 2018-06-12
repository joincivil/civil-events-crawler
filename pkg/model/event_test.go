// Package model_test contains the tests for the model package
package model_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"math/big"
	"reflect"
	"testing"
)

var (
	contractAddress = "0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d"
	testAddress     = "0xdfe273082089bb7f70ee36eebcde64832fe97e55"
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
		},
	}
	testEvent2 = &contract.CivilTCRContractApplicationWhitelisted{
		ListingAddress: common.HexToAddress(testAddress),
		Raw: types.Log{
			Address:     common.HexToAddress(testAddress),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: 8888888,
		},
	}
)

func setupCivilEvent() *model.CivilEvent {
	event := model.NewCivilEvent("_Application", common.HexToAddress(contractAddress),
		testEvent)
	return event
}

func TestCivilEventSetup(t *testing.T) {
	event := setupCivilEvent()
	if event == nil {
		t.Errorf("Civil event was not initialized correctly")
	}
	if event.EventType != "_Application" {
		t.Errorf("EventType was not init correctly: %v", event.EventType)
	}
	if event.Timestamp <= 0 {
		t.Errorf("Timestamp was not init correctly: %v", event.Timestamp)
	}
	if event.Payload == nil {
		t.Errorf("Payload was not init correctly: %v", event.Payload)
	}
}

func TestCivilEventPayload(t *testing.T) {
	event := setupCivilEvent()

	datafields := event.Payload.Keys()
	if len(datafields) != 6 {
		t.Errorf("Payload does not have all the fields: %v", datafields)
	}
}

func TestCivilEventPayloadValues(t *testing.T) {
	event := setupCivilEvent()

	_, ok := event.Payload.Value("NonexistentKey")
	if ok {
		t.Errorf("Non-existent key should not return value")
	}

	value, ok := event.Payload.Value("ListingAddress")
	if !ok {
		t.Errorf("ListingAddress not found")
	}
	val := value.Val()
	_, ok = val.(common.Address)
	if !ok {
		t.Errorf("ListingAddress Val() call should have accepted type assert")
	}

	address, ok := value.Address()
	if !ok {
		t.Errorf("ListingAddress cannot be the type common.Address")
	}
	testAddress := common.HexToAddress(testAddress)
	if address.Hex() != testAddress.Hex() {
		t.Errorf("ListingAddress not == original: %v", address.Hex())
	}
	if value.Kind() != reflect.Array {
		t.Errorf("ListingAddress not an array kind: %v", value.Kind())
	}
	_, ok = value.BigInt()
	if ok {
		t.Errorf("ListingAddress should fail on type assert to big.Int")
	}
	_, ok = value.Int64()
	if ok {
		t.Errorf("ListingAddress should fail on type assert to int64")
	}

	value, ok = event.Payload.Value("Deposit")
	if !ok {
		t.Errorf("Deposit not found")
	}
	depositInt, ok := value.BigInt()
	if !ok {
		t.Errorf("Deposit cannot be the type big.Int")
	}
	if depositInt.Int64() != big.NewInt(1000).Int64() {
		t.Errorf("Deposit not == original: %v", depositInt.Int64())
	}
	depositInt64, ok := value.Int64()
	if !ok {
		t.Errorf("Deposit cannot be the type int64")
	}
	if depositInt64 != int64(1000) {
		t.Errorf("Deposit not == original: %v", depositInt64)
	}
	if value.Kind() != reflect.Ptr {
		t.Errorf("Deposit not an ptr kind: %v", value.Kind())
	}

	value, ok = event.Payload.Value("Data")
	if !ok {
		t.Errorf("Data not found")
	}

	dataStr, ok := value.String()
	if !ok {
		t.Errorf("Data cannot be the type string")
	}
	if dataStr != "DATA" {
		t.Errorf("Data not == original: %v", dataStr)
	}

	value, ok = event.Payload.Value("Raw")
	if !ok {
		t.Errorf("Raw not found")
	}
	_, ok = value.Log()
	if !ok {
		t.Errorf("Raw log cannot be the type types.Log")
	}
}

// Test that these 2 event hashes are not equal
func TestCivilEventHash(t *testing.T) {
	civilEvent1 := model.NewCivilEvent("Application", common.HexToAddress(contractAddress),
		testEvent)
	civilEvent2 := model.NewCivilEvent("ApplicationWhitelisted", common.HexToAddress(contractAddress),
		testEvent2)
	if civilEvent2.EventHash == civilEvent1.EventHash {
		t.Error("These events should have different hashes but they are the same")
	}
}
