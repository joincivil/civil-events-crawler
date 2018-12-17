// Package model_test contains the tests for the model package
package model_test

import (
	// "fmt"
	"math/big"
	"reflect"
	"strings"
	"testing"

	ctime "github.com/joincivil/go-common/pkg/time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"
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
			Address: common.HexToAddress(testAddress),
			Topics: []common.Hash{
				common.HexToHash("0x98c8cf45bd844627e84e1c506ca87cc9436317d0"),
			},
			Data:        []byte("data"),
			BlockNumber: 8888888,
			Index:       2,
			TxIndex:     3,
			TxHash:      common.HexToHash("0x98c8cf45bd844627e84e1c506ca87cc9436317d0"),
			Removed:     true,
		},
	}
	testEvent2 = &contract.CivilTCRContractApplicationWhitelisted{
		ListingAddress: common.HexToAddress(testAddress),
		Raw: types.Log{
			Address: common.HexToAddress(testAddress),
			Topics: []common.Hash{
				common.HexToHash("0x98c8cf45bd844627e84e1c506ca87cc9436317d0"),
			},
			Data:        []byte("data"),
			BlockNumber: 8888888,
			Index:       1,
			TxIndex:     4,
			TxHash:      common.HexToHash("0x98c8cf45bd844627e84e1c506ca87cc9436317d0"),
			Removed:     false,
		},
	}
)

func setupEvent() (*model.Event, error) {
	return model.NewEventFromContractEvent(
		"Application",
		"CivilTCRContract",
		common.HexToAddress(contractAddress),
		testEvent,
		ctime.CurrentEpochSecsInInt64(),
		model.Filterer,
	)
}

func TestEventSetup(t *testing.T) {
	event, err := setupEvent()
	if err != nil {
		t.Errorf("setupEvent should have succeeded: err: %v", err)
	}
	if event == nil {
		t.Errorf("Event was not initialized correctly")
	}
	if event.EventType() != "Application" {
		t.Errorf("EventType was not init correctly: %v", event.EventType())
	}
	if strings.ToLower(event.ContractAddress().Hex()) != strings.ToLower(contractAddress) {
		t.Errorf("ContractAddress was not init correctly: %v", event.ContractAddress())
	}
	if event.Timestamp() <= 0 {
		t.Errorf("Timestamp was not init correctly: %v", event.Timestamp())
	}
	if event.EventPayload() == nil {
		t.Errorf("Payload was not init correctly: %v", event.EventPayload())
	}
}

func TestEventPayload(t *testing.T) {
	event, _ := setupEvent()
	payload := event.EventPayload()
	if len(payload) != 5 {
		t.Errorf("Payload does not have all the fields: %v", payload)
	}
}

type testStructNoRaw struct {
	name string
}

func TestEventPayloadNoRaw(t *testing.T) {
	noRawTestEvent := &testStructNoRaw{
		name: "name",
	}
	_, err := model.NewEventFromContractEvent("Application", "CivilTCRContract", common.HexToAddress(contractAddress),
		noRawTestEvent, ctime.CurrentEpochSecsInInt64(), model.Watcher)
	if err == nil {
		t.Errorf("Event creation should have failed with no raw event to create hash: err: %v", err)
	}
}

type testStructNotLog struct {
	name string
	Raw  string
}

func TestEventPayloadNotLog(t *testing.T) {
	notLogTestEvent := &testStructNotLog{
		name: "name",
		Raw:  "name",
	}
	_, err := model.NewEventFromContractEvent("Application", "CivilTCRContract", common.HexToAddress(contractAddress),
		notLogTestEvent, ctime.CurrentEpochSecsInInt64(), model.Watcher)
	if err == nil {
		t.Errorf("Event creation should have failed with no Log found: err: %v", err)
	}
}

func TestEventPayloadValues(t *testing.T) {
	event, _ := setupEvent()
	payload := event.EventPayload()
	if len(payload) != 5 {
		t.Errorf("Wrong number of fields in eventPayload field %v", payload)
	}

	value := payload["ListingAddress"]
	if value != testEvent.ListingAddress {
		t.Errorf("ListingAddress not converted correctly %v", value)
	}

	value = payload["Deposit"]
	if value != testEvent.Deposit {
		t.Errorf("Deposit not converted correctly %v", value)
	}

	value = payload["AppEndDate"]
	if value != testEvent.AppEndDate {
		t.Errorf("AppEndDate not converted correctly %v", value)
	}

	value = payload["Data"]
	if value != testEvent.Data {
		t.Errorf("Data not converted correctly %v", value)
	}

	value = payload["Applicant"]
	if value != testEvent.Applicant {
		t.Errorf("Applicant not converted correctly %v", value)
	}
}

// Test that these 2 event hashes are not equal
func TestEventHashDifferent(t *testing.T) {
	civilEvent1, _ := model.NewEventFromContractEvent("Application", "CivilTCRContract", common.HexToAddress(contractAddress), testEvent,
		ctime.CurrentEpochSecsInInt64(), model.Watcher)
	civilEvent2, _ := model.NewEventFromContractEvent("ApplicationWhitelisted", "CivilTCRContract", common.HexToAddress(contractAddress),
		testEvent2, ctime.CurrentEpochSecsInInt64(), model.Watcher)
	if civilEvent2.Hash() == civilEvent1.Hash() {
		t.Error("These events should have different hashes")
	}
}

// Test that hash created on the same event are the same
func TestEventHashSame(t *testing.T) {
	timestamp := ctime.CurrentEpochSecsInInt64()
	civilEvent, _ := model.NewEventFromContractEvent("Application", "CivilTCRContract", common.HexToAddress(contractAddress), testEvent,
		timestamp, model.Watcher)
	civilEventDup, _ := model.NewEventFromContractEvent("Application", "CivilTCRContract", common.HexToAddress(contractAddress),
		testEvent, timestamp, model.Watcher)
	if civilEvent.Hash() != civilEventDup.Hash() {
		t.Error("These events should have the same hashes")
	}
}

func TestEventPayloadStructValues(t *testing.T) {
	payload := model.NewEventPayload(testEvent)

	_, ok := payload.Value("NonexistentKey")
	if ok {
		t.Errorf("Non-existent key should not return value")
	}

	toStr := payload.ToString()
	if toStr == "" {
		t.Errorf("ToString is returning an empty string")
	}

	value, ok := payload.Value("ListingAddress")
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

	value, ok = payload.Value("Deposit")
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

	value, ok = payload.Value("Data")
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

	value, ok = payload.Value("Raw")
	if !ok {
		t.Errorf("Raw not found")
	}
	_, ok = value.Log()
	if !ok {
		t.Errorf("Raw log cannot be the type types.Log")
	}

}

func TestEventLogToString(t *testing.T) {
	event, err := setupEvent()
	if err != nil {
		t.Errorf("setupEvent should have succeeded: err: %v", err)
	}
	logString := event.LogPayloadToString()
	logStringTrue := "log: addr: 0xDFe273082089bB7f70Ee36Eebcde64832FE97E55, blknum: 8888888, txhash: 0x00000000000000000000000098c8cf45bd844627e84e1c506ca87cc9436317d0, txidx: 3, blkhash: 0x0000000000000000000000000000000000000000000000000000000000000000, idx: 2, rem: true"
	if logString != logStringTrue {
		t.Errorf("logString is not what it should be %v", logString)
	}
}

func TestEventLogPayloadImmutability(t *testing.T) {
	event, err := setupEvent()
	if err != nil {
		t.Errorf("setupEvent should have succeeded: err: %v", err)
	}
	logPayload := event.LogPayload()
	newBlockNo := uint64(32)
	logPayload.BlockNumber = newBlockNo
	if newBlockNo == event.LogPayload().BlockNumber {
		t.Errorf("these should not be equal")
	}
}

func TestEventGettersMutators(t *testing.T) {
	event, err := setupEvent()
	if err != nil {
		t.Errorf("setupEvent should have succeeded: err: %v", err)
	}

	if event.Hash() == "" {
		t.Errorf("hash should have returned a valid hash: %v", event.Hash())
	}

	if event.EventType() != "Application" {
		t.Errorf("hash should have returned the valid event: %v", event.EventType())
	}

	normalizedAddress := common.HexToAddress(contractAddress)
	if event.ContractAddress().Hex() != normalizedAddress.Hex() {
		t.Errorf("contract address should have matched: %v, %v", event.ContractAddress().Hex(), normalizedAddress.Hex())
	}

	if event.Timestamp() == 0 {
		t.Errorf("timestamp should have been > 0: %v", event.Timestamp())
	}

	event.SetTimestamp(1024)
	if event.Timestamp() != 1024 {
		t.Errorf("timestamp should have been 1024: %v", event.Timestamp())
	}

	if event.RetrievalMethod() != model.Filterer {
		t.Errorf("retrieval method should have been filterer: %v", event.RetrievalMethod())
	}

	if event.ContractName() != "CivilTCRContract" {
		t.Errorf("contract name should have been civiltcrcontract: %v", event.ContractName())
	}

	if event.LogTopics()[0].Hex() != common.HexToHash("0x98c8cf45bd844627e84e1c506ca87cc9436317d0").Hex() {
		t.Errorf("should have had correct topic: %v", event.LogTopics()[0].Hex())
	}

	if string(event.LogData()) != "data" {
		t.Errorf("should have had correct data: %v", event.LogData())
	}

	if event.BlockNumber() != uint64(8888888) {
		t.Errorf("should have had correct blocknumber: %v", event.BlockNumber())
	}

	if event.TxIndex() != 3 {
		t.Errorf("should have had correct txindex: %v", event.TxIndex())
	}

	if event.TxHash().Hex() != common.HexToHash("0x98c8cf45bd844627e84e1c506ca87cc9436317d0").Hex() {
		t.Errorf("should have had correct txHash: %v", event.TxHash())
	}

	if event.BlockHash().Hex() == "" {
		t.Errorf("should have had correct block hash: %v", event.BlockHash().Hex())
	}

	if !event.LogRemoved() {
		t.Errorf("should have been the correct log removed bool: %v", event.LogRemoved())
	}

	if event.LogIndex() != 2 {
		t.Errorf("should have been the correct log index: %v", event.LogIndex())
	}
}
