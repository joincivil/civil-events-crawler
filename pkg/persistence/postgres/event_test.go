package postgres_test

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/joincivil/go-common/pkg/generated/contract"
	ctime "github.com/joincivil/go-common/pkg/time"

	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"
)

var (
	contractAddress  = "0x77e5aaBddb760FBa989A1C4B2CDd4aA8Fa3d311d"
	pContractAddress = "0x4f4b97A4FaeBf2BD835a10C479E61faccEf8755D"
	testAddress      = "0xDFe273082089bB7f70Ee36Eebcde64832FE97E55"
	testEvent        = &contract.CivilTCRContractApplication{
		ListingAddress: common.HexToAddress(testAddress),
		Deposit:        big.NewInt(1000),
		AppEndDate:     big.NewInt(1653860896),
		Data:           "DATA",
		Applicant:      common.HexToAddress(testAddress),
		Raw: types.Log{
			Address: common.HexToAddress(testAddress),
			Topics: []common.Hash{
				common.HexToHash("0x09cd8dcaf170a50a26316b5fe0727dd9fb9581a688d65e758b16a1650da65c0b"),
				common.HexToHash("0x0000000000000000000000002652c60cf04bbf6bb6cc8a5e6f1c18143729d440"),
				common.HexToHash("0x00000000000000000000000025bf9a1595d6f6c70e6848b60cba2063e4d9e552"),
			},
			Data:        []byte("thisisadatastring"),
			BlockNumber: 8888888,
			TxHash:      common.Hash{},
			TxIndex:     2,
			BlockHash:   common.Hash{},
			Index:       2,
			Removed:     false,
		},
	}

	fakeVersion = "123"

	testEvent2 = &contract.ParameterizerContractProposalExpired{
		PropID: [32]byte{0x00, 0x01},
		Raw: types.Log{
			Address: common.HexToAddress(testAddress),
			Topics: []common.Hash{
				common.HexToHash("0x09cd8dcaf170a50a26316b5fe0727dd9fb9581a688d65e758b16a1650da65c0b"),
				common.HexToHash("0x0000000000000000000000002652c60cf04bbf6bb6cc8a5e6f1c18143729d440"),
				common.HexToHash("0x00000000000000000000000025bf9a1595d6f6c70e6848b60cba2063e4d9e552"),
			},
			Data:        []byte("thisisadatastring"),
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
	return model.NewEventFromContractEvent(
		"Application",
		"CivilTCRContract",
		common.HexToAddress(contractAddress),
		testEvent,
		ctime.CurrentEpochSecsInInt64(),
		model.Filterer,
	)
}

func setupDBEvent(t *testing.T) *postgres.Event {
	civilEvent, err := setupEvent()
	if err != nil {
		t.Errorf("setupEvent should have succeeded: err: %v", err)
	}
	dbEvent, err := setupDBEventFromEvent(civilEvent)
	if err != nil {
		t.Errorf("setupDBEventFromEvent should have succeeded: err: %v", err)
	}
	return dbEvent
}

func TestDBEventSetup(t *testing.T) {
	dbEvent := setupDBEvent(t)

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

func TestDBToEventData(t *testing.T) {
	dbEvent := setupDBEvent(t)
	contractName := dbEvent.ContractName

	_, err := dbEvent.DBToEventData()
	if err != nil {
		t.Errorf("Should have not received error when converted to model event: err: %v", err)
	}

	dbEvent.ContractName = ""
	_, err = dbEvent.DBToEventData()
	if err == nil {
		t.Errorf("Should have received error with no contract name: err: %v", err)
	}

	dbEvent.ContractName = contractName
	dbEvent.EventType = ""
	_, err = dbEvent.DBToEventData()
	if err == nil {
		t.Errorf("Should have received error with no event type: err: %v", err)
	}
}

func TestDBToEventLogData(t *testing.T) {
	dbEvent := setupDBEvent(t)

	tLog := dbEvent.DBToEventLogData()
	if tLog == nil {
		t.Errorf("Should have received valid type.Log")
	}

	dbEvent = setupDBEvent(t)
	dbEvent.LogPayload["Topics"] = []common.Hash{
		common.HexToHash("0x98c8cf45bd844627e84e1c506ca87cc9436317d0"),
	}
	tLog = dbEvent.DBToEventLogData()
	if tLog == nil {
		t.Errorf("Should have received type.Log")
	}
	if tLog.Topics != nil {
		t.Errorf("Should not have received topics")
	}

	dbEvent = setupDBEvent(t)
	dbEvent.LogPayload["Data"] = []int8{1, 2, 3}
	tLog = dbEvent.DBToEventLogData()
	if tLog == nil {
		t.Errorf("Should have received non-nil type.Log")
	}
	if tLog.Data != nil {
		t.Errorf("Should not have received data")
	}

	dbEvent = setupDBEvent(t)
	dbEvent.LogPayload["BlockNumber"] = uint64(64)
	tLog = dbEvent.DBToEventLogData()
	if tLog == nil {
		t.Errorf("Should have received type.Log")
	}
	if tLog.BlockNumber != 0 {
		t.Errorf("Should not have received block number")
	}

	dbEvent = setupDBEvent(t)
	dbEvent.LogPayload["TxIndex"] = uint(64)
	tLog = dbEvent.DBToEventLogData()
	if tLog == nil {
		t.Errorf("Should have received type.Log")
	}
	if tLog.TxIndex != 0 {
		t.Errorf("Should not have received tx index")
	}

}

func TestEventDataToDB(t *testing.T) {
	modelEvent, err := setupEvent()
	if err != nil {
		t.Errorf("Should not have received error setting up event: err: %v", err)
	}
	t.Logf("event payload = %v, %v", modelEvent.EventPayload(), modelEvent.ContractName())
	dbEvent := &postgres.Event{
		ContractName: modelEvent.ContractName(),
		EventType:    modelEvent.EventType(),
	}
	err = dbEvent.EventDataToDB(modelEvent.EventPayload())
	if err != nil {
		t.Errorf("Should not have received error converting event data: err: %v", err)
	}

	dbEvent = &postgres.Event{
		EventType: modelEvent.EventType(),
	}
	err = dbEvent.EventDataToDB(modelEvent.EventPayload())
	if err == nil {
		t.Errorf("Should have received error with no contract name in event: err: %v", err)
	}

	dbEvent = &postgres.Event{
		ContractName: modelEvent.ContractName(),
	}
	err = dbEvent.EventDataToDB(modelEvent.EventPayload())
	if err == nil {
		t.Errorf("Should have received error with no event name in event: err: %v", err)
	}
}

func TestInt64Overflow(t *testing.T) {
	deposit := new(big.Int)
	deposit.SetString("100000000000000000000", 10)
	testEvent.Deposit = deposit

	dbEvent := setupDBEvent(t)
	depositFloat, _ := new(big.Float).SetInt(deposit).Float64()
	if dbEvent.EventPayload["Deposit"] != depositFloat {
		t.Errorf("Wrong value, %v, %v", dbEvent.EventPayload["Deposit"], depositFloat)
	}
}

func TestCreateTableQuery(t *testing.T) {
	query := postgres.CreateEventTableQuery(fakeVersion)
	if query == "" {
		t.Errorf("Should have returned a value for query")
	}
	if !strings.Contains(query, "CREATE TABLE") {
		t.Errorf("Should have returned CREATE TABLE values")
	}
}

func TestCreateEventTableIndices(t *testing.T) {
	query := postgres.CreateEventTableIndices("event")
	if query == "" {
		t.Errorf("Should have returned a value for query")
	}
	if !strings.Contains(query, "CREATE INDEX IF NOT EXISTS") {
		t.Errorf("Should have returned CREATE TABLE values")
	}

}

func TestByte32Conversions(t *testing.T) {
	mEvent, err := model.NewEventFromContractEvent(
		"ProposalExpired",
		"ParameterizerContract",
		common.HexToAddress(pContractAddress),
		testEvent2,
		ctime.CurrentEpochSecsInInt64(),
		model.Filterer,
	)
	if err != nil {
		t.Errorf("error creating new event, err: %v", err)
	}
	dbEvent, err := setupDBEventFromEvent(mEvent)
	if err != nil {
		t.Errorf("setupDBEventFromEvent should have succeeded: err: %v", err)
	}
	// fmt.Println(dbEvent)
	newDBEvent, err := dbEvent.DBToEventData()
	if err != nil {
		t.Errorf("Should have not received error when converted to model event: err: %v", err)
	}
	if newDBEvent.EventPayload()["PropID"] == nil {
		t.Error("Should have PropID field in event payload")
	}
}
