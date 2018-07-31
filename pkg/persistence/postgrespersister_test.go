// +build integration

// This is an integration test file for postgrespersister. Postgres needs to be running.
// Run this using go test -tags=integration
// Run benchmark test using go test -tags=integration -bench=.
package persistence

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"math/big"
	"reflect"
	"testing"
	"time"
)

const (
	postgresPort   = 5432
	postgresDBName = "civil_crawler"
	postgresUser   = "docker"
	postgresPswd   = "docker"
	postgresHost   = "localhost"
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
	testEvent2 = &contract.CivilTCRContractApplicationWhitelisted{
		ListingAddress: common.HexToAddress(testAddress),
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
	ChallengeID   *big.Int
	Data          string
	CommitEndDate *big.Int
	RevealEndDate *big.Int
	Challenger    common.Address
	testEvent3    = &contract.CivilTCRContractChallenge{
		ListingAddress: common.HexToAddress(testAddress),
		ChallengeID:    big.NewInt(8),
		Data:           "DATA",
		CommitEndDate:  big.NewInt(1653860896),
		RevealEndDate:  big.NewInt(1653860896),
		Challenger:     common.HexToAddress(testAddress),
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

// Sets up an Application event and generates a random hash for address so that the hash in DB is unique.
func setupEvent(rand bool) (*model.Event, error) {
	if rand {
		randString, _ := randomHex(32)
		testEvent.Raw.TxHash = common.HexToHash(randString)
	}
	return model.NewEventFromContractEvent("Application", "CivilTCRContract", common.HexToAddress(contractAddress),
		testEvent, utils.CurrentEpochSecsInInt()-(60*60*3))
}

// Sets up an ApplicationWhitelisted event and generates a random hash for address so that the hash in DB is unique.
func setupEvent2(rand bool) (*model.Event, error) {
	if rand {
		randString, _ := randomHex(32)
		testEvent2.Raw.TxHash = common.HexToHash(randString)
	}
	return model.NewEventFromContractEvent("ApplicationWhitelisted", "CivilTCRContract", common.HexToAddress(contractAddress),
		testEvent2, utils.CurrentEpochSecsInInt()-(60*60*2))
}

// Sets up an Challenge event and generates a random hash for address so that the hash in DB is unique.
func setupEvent3(rand bool) (*model.Event, error) {
	if rand {
		randString, _ := randomHex(32)
		testEvent2.Raw.TxHash = common.HexToHash(randString)
	}
	return model.NewEventFromContractEvent("Challenge", "CivilTCRContract", common.HexToAddress(contractAddress),
		testEvent3, utils.CurrentEpochSecsInInt()-(60*60))
}

// random hex string generation
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func setupDBConnection() (*PostgresPersister, error) {
	postgresPersister, err := NewPostgresPersister(postgresHost, postgresPort, postgresUser, postgresPswd, postgresDBName)
	return postgresPersister, err
}

func setupTestTable() (*PostgresPersister, error) {
	persister, err := setupDBConnection()
	if err != nil {
		return persister, fmt.Errorf("Error connecting to DB: %v", err)
	}
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS event_test(
			id SERIAL PRIMARY KEY,
			event_type TEXT,
			hash TEXT UNIQUE,
			contract_address TEXT,
			contract_name TEXT,
			timestamp INT,
			payload JSONB,
			log_payload JSONB
		);
	`
	_, err = persister.db.Query(createTableQuery)
	if err != nil {
		return persister, fmt.Errorf("Couldn't create test table %v", err)
	}

	return persister, nil
}

func deleteTestTable(persister *PostgresPersister) error {
	_, err := persister.db.Query("DROP TABLE event_test;")
	if err != nil {
		return fmt.Errorf("Couldn't delete test table %v", err)
	}
	return nil
}

func TestDBConnection(t *testing.T) {
	persister, err := setupDBConnection()
	if err != nil {
		t.Errorf("Error connecting to DB: %v", err)
	}
	var result int
	err = persister.db.QueryRow("SELECT 1;").Scan(&result)
	if err != nil {
		t.Errorf("Error querying DB: %v", err)
	}
	if result != 1 {
		t.Errorf("Wrong result from DB")
	}
}

func TestTableSetup(t *testing.T) {
	// run function to create tables, and test table exists
	persister, err := setupDBConnection()
	if err != nil {
		t.Errorf("Error connecting to DB: %v", err)
	}
	err = persister.CreateTables()
	if err != nil {
		t.Errorf("Error creating/checking for tables: %v", err)
	}
	// check table exists
	var exists bool
	err = persister.db.QueryRow(`SELECT EXISTS ( SELECT 1
                                        FROM   information_schema.tables 
                                        WHERE  table_schema = 'public'
                                        AND    table_name = 'event'
                                        );`).Scan(&exists)
	if err != nil {
		t.Errorf("Couldn't get table")
	}
	if !exists {
		t.Errorf("event table does not exist")
	}

}

func TestIndexCreation(t *testing.T) {
	persister, err := setupDBConnection()
	if err != nil {
		t.Errorf("Error connecting to DB: %v", err)
	}
	err = persister.CreateTables()
	if err != nil {
		t.Errorf("Error creating/checking for tables: %v", err)
	}
	err = persister.CreateIndices()
	if err != nil {
		t.Errorf("Error creating indices: %v", err)
	}
}

func TestIndexCreationTestTable(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)

	indexCreationQuery := `
		CREATE INDEX IF NOT EXISTS event_event_type_idx ON event_test (event_type);
		CREATE INDEX IF NOT EXISTS event_contract_address_idx ON event_test (contract_address);
		CREATE INDEX IF NOT EXISTS event_timestamp_idx ON event_test (timestamp);
	`
	_, err = persister.db.Query(indexCreationQuery)
	if err != nil {
		t.Errorf("Error creating indices in test table: %v", err)
	}
	// run query twice to ensure indices won't throw an error
	_, err = persister.db.Query(indexCreationQuery)
	if err != nil {
		t.Errorf("Error creating indices in test table: %v", err)
	}

}

func TestSaveToEventTestTable(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)
	event, err := setupEvent(true)
	if err != nil {
		t.Errorf("Couldn't setup civilEvent from contract %v", err)
	}

	civilEventsFromContract := []*model.Event{event}
	err = persister.saveEventsToTable(civilEventsFromContract, "event_test")
	if err != nil {
		t.Errorf("Cannot save event to events_test table: %v", err)
	}

	civilEventDB, err := persister.GetAllEvents("event_test")
	if err != nil {
		t.Errorf("error querying event from events_test table: %v", err)
	}
	if len(civilEventDB) != 1 {
		t.Errorf("expected there to be only 1 event in table but there is %v events", len(civilEventDB))
	}
	err = deleteTestTable(persister)
	if err != nil {
		t.Error(err)
	}
}

// TODO (IS): fix this test. this isn't true benchmark of saving events bc of the hashing function on creation
// of each event
func BenchmarkSavingManyEventsToEventTestTable(b *testing.B) {
	persister, err := setupTestTable()
	if err != nil {
		b.Error(err)
	}
	defer deleteTestTable(persister)

	numEvents := 100
	civilEventsFromContract := make([]*model.Event, 0)
	for i := 1; i <= numEvents; i++ {
		event, err := setupEvent(true)
		if err != nil {
			b.Errorf("Couldn't setup civilEvent from contract %v", err)
		}
		civilEventsFromContract = append(civilEventsFromContract, event)
	}

	err = persister.saveEventsToTable(civilEventsFromContract, "event_test")
	if err != nil {
		b.Errorf("Cannot save event to event_test table: %v", err)
	}
	var numRows int
	err = persister.db.QueryRow(`SELECT COUNT(*) FROM
                                        event_test`).Scan(&numRows)
	if numRows != numEvents {
		b.Errorf("Number of rows in event_test table should be %v but it is %v", numEvents, numRows)
	}
	err = deleteTestTable(persister)
	if err != nil {
		b.Error(err)
	}
}

func TestPersistenceUpdate(t *testing.T) {
	//Check that persistence is being updated
	persister, err := setupDBConnection()
	if err != nil {
		t.Errorf("Error connecting to DB: %v", err)
	}
	event, err := setupEvent(true)
	if err != nil {
		t.Errorf("Couldn't setup civilEvent from contract %v", err)
	}
	civilEventsFromContract := []*model.Event{event}
	err = persister.UpdateLastBlockData(civilEventsFromContract)
	if err != nil {
		t.Errorf("Couldn't update last block data: %v", err)
	}

	if persister.LastBlockNumber(event.EventType(), event.ContractAddress()) != testEvent.Raw.BlockNumber {
		t.Error("Blocknumber was not updated correctly in persistence")
	}
	if persister.LastBlockHash(event.EventType(), event.ContractAddress()) != testEvent.Raw.BlockHash {
		t.Error("Blockhash was not updated correctly in persistence")
	}
}

func TestLatestEventsQuery(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)
	var latestTimestamp int
	civilEventsFromContract := make([]*model.Event, 0)

	event, err := setupEvent(true)
	if err != nil {
		t.Errorf("Couldn't setup Application civilEvent from contract %v", err)
	}
	event2, err := setupEvent2(true)
	if err != nil {
		t.Errorf("Couldn't setup ApplicationWhitelisted civilEvent from contract %v", err)
	}
	civilEventsFromContract = append(civilEventsFromContract, event, event2)

	err = persister.saveEventsToTable(civilEventsFromContract, "event_test")
	if err != nil {
		t.Errorf("Cannot save event to event_test table: %v", err)
	}

	events, err := persister.getLatestEvents("event_test")
	if err != nil {
		t.Errorf("Error retrieving events: %v", err)
	}
	if len(events) != 2 {
		t.Errorf("Query should have only returned 2 events but there are %v", len(events))
	}

	latestTimestamp = event.Timestamp()
	queryTimestamp := events[0].Timestamp
	if queryTimestamp != latestTimestamp {
		t.Errorf("Query didn't pull the latest event for contract and event type for %v", events[0].EventType)
	}

	latestTimestamp = event2.Timestamp()
	queryTimestamp2 := events[1].Timestamp
	if queryTimestamp2 != latestTimestamp {
		t.Errorf("Query didn't pull the latest event for contract and event type %v", events[1].EventType)
	}

}

func TestPersistenceUpdateFromDB(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)
	numEvents := 3
	civilEventsFromContract := make([]*model.Event, 0)
	for i := 1; i <= numEvents; i++ {
		event, err := setupEvent(true)
		if err != nil {
			t.Errorf("Couldn't setup civilEvent from contract %v", err)
		}
		civilEventsFromContract = append(civilEventsFromContract, event)
		time.Sleep(1 * time.Second)
	}

	err = persister.saveEventsToTable(civilEventsFromContract, "event_test")
	if err != nil {
		t.Errorf("Cannot save event to event_test table: %v", err)
	}
	err = persister.PopulateBlockDataFromDB("event_test")
	if err != nil {
		t.Errorf("Cannot fill persistence, %v", err)
	}

	blockNumber := persister.eventToLastBlockNumber[common.HexToAddress(contractAddress)]["Application"].BlockNumber
	correctBlockNumber := testEvent.Raw.BlockNumber
	if blockNumber != correctBlockNumber {
		t.Errorf("Block number should be %v but is %v", correctBlockNumber, blockNumber)
	}

	blockHash := persister.eventToLastBlockNumber[common.HexToAddress(contractAddress)]["Application"].BlockHash
	correctBlockHash := testEvent.Raw.BlockHash
	if blockHash != correctBlockHash {
		t.Errorf("Block number should be %v but is %v", correctBlockHash, blockHash)
	}

}

// This conversion needs to be here, bc we need the actual event after being saved in DB.
func TestDBToEvent(t *testing.T) {
	civilEvent, err := setupEvent(true)
	if err != nil {
		t.Errorf("setupEvent should have succeeded: err: %v", err)
	}
	// Get this event from DB
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)
	civilEventsFromContract := []*model.Event{civilEvent}
	err = persister.saveEventsToTable(civilEventsFromContract, "event_test")
	if err != nil {
		t.Errorf("Cannot save event to event_test table: %v", err)
	}

	civilEventDB, err := persister.GetAllEvents("event_test")

	dbEvent := civilEventDB[0]

	civilEventFromDB, err := dbEvent.DBToEventData()
	if err != nil {
		t.Errorf("Could not convert db event back to civilevent: err: %v", err)
	}

	// deep equal doesn't work bc of nested slices, etc. so just compare each element
	if civilEvent.ContractAddress() != civilEventFromDB.ContractAddress() {
		t.Errorf("ContractAddress not equal: %v %v", civilEvent.ContractAddress(), civilEventFromDB.ContractAddress())
	}
	if civilEvent.ContractName() != civilEventFromDB.ContractName() {
		t.Errorf("ContractName not equal: %v %v", civilEvent.ContractName(), civilEventFromDB.ContractName())
	}
	if civilEvent.Hash() != civilEventFromDB.Hash() {
		t.Errorf("Hash not equal: %v %v", civilEvent.Hash(), civilEventFromDB.Hash())
	}
	if civilEvent.EventType() != civilEventFromDB.EventType() {
		t.Errorf("EventType not equal: %v %v", civilEvent.EventType(), civilEventFromDB.EventType())
	}
	if civilEvent.Timestamp() != civilEventFromDB.Timestamp() {
		t.Errorf("Timestamp not equal: %v %v", civilEvent.Timestamp(), civilEventFromDB.Timestamp())
	}

	// EventPayload
	if !reflect.DeepEqual(civilEventFromDB.EventPayload(), civilEvent.EventPayload()) {
		t.Errorf("EventPayloads not equal: %v %v", civilEvent.EventPayload(), civilEventFromDB.EventPayload())
	}

	// LogPayload
	civilLogPayload := civilEvent.LogPayload()
	civilLogFromDBPayload := civilEventFromDB.LogPayload()

	if civilLogPayload.Address != civilLogFromDBPayload.Address {
		t.Errorf("Address in Log not equal: %v %v", civilLogPayload.Address, civilLogFromDBPayload.Address)
	}
	if !reflect.DeepEqual(civilLogPayload.Topics, civilLogFromDBPayload.Topics) {
		t.Errorf("Topics in Log not equal: %v %v", civilLogPayload.Topics, civilLogFromDBPayload.Topics)
	}
	if !reflect.DeepEqual(civilLogPayload.Data, civilLogFromDBPayload.Data) {
		t.Errorf("Data in Log not equal: %v %v", civilLogPayload.Data, civilLogFromDBPayload.Data)
	}
	if civilLogPayload.BlockNumber != civilLogFromDBPayload.BlockNumber {
		t.Errorf("BlockNumber in Log not equal: %v %v", civilLogPayload.BlockNumber, civilLogFromDBPayload.BlockNumber)
	}
	if civilLogPayload.TxHash != civilLogFromDBPayload.TxHash {
		t.Errorf("TxHash in Log not equal: %v %v", civilLogPayload.TxHash, civilLogFromDBPayload.TxHash)
	}
	if civilLogPayload.TxIndex != civilLogFromDBPayload.TxIndex {
		t.Errorf("TxIndex in Log not equal: %v %v", civilLogPayload.TxIndex, civilLogFromDBPayload.TxIndex)
	}
	if civilLogPayload.BlockHash != civilLogFromDBPayload.BlockHash {
		t.Errorf("BlockHash in Log not equal: %v %v", civilLogPayload.BlockHash, civilLogFromDBPayload.BlockHash)
	}
	if civilLogPayload.Index != civilLogFromDBPayload.Index {
		t.Errorf("Index in Log not equal: %v %v", civilLogPayload.Index, civilLogFromDBPayload.Index)
	}
	if civilLogPayload.Removed != civilLogFromDBPayload.Removed {
		t.Errorf("Removed in Log not equal: %v %v", civilLogPayload.Removed, civilLogFromDBPayload.Removed)
	}

}

func TestRetrieveEvents(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)
	civilEventsFromContract := []*model.Event{}
	event, err := setupEvent(true)
	if err != nil {
		t.Errorf("Should not have received error from setting up event from contract %v", err)
	}
	event2, err := setupEvent2(true)
	if err != nil {
		t.Errorf("Should not have received error from setting up event from contract %v", err)
	}
	event3, err := setupEvent3(true)
	if err != nil {
		t.Errorf("Should not have received error from setting up event from contract %v", err)
	}
	civilEventsFromContract = append(civilEventsFromContract, event, event2, event3)

	err = persister.saveEventsToTable(civilEventsFromContract, "event_test")
	if err != nil {
		t.Errorf("Should not have seen error when saving events to table: %v", err)
	}

	events, err := persister.retrieveEvents("event_test", &model.RetrieveEventsCriteria{
		Offset:  0,
		Count:   1,
		Reverse: false,
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("Should have seen only 1 event: %v", len(events))
	}
	if events[0].EventType() != event.EventType() {
		t.Errorf("Should have seen the type of the oldest event: err: %v", err)
	}

	events, err = persister.retrieveEvents("event_test", &model.RetrieveEventsCriteria{
		Offset:  0,
		Count:   3,
		Reverse: false,
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 3 {
		t.Errorf("Should have seen only 3 event: %v", len(events))
	}

	events, err = persister.retrieveEvents("event_test", &model.RetrieveEventsCriteria{
		Offset:  0,
		Count:   5,
		Reverse: false,
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 3 {
		t.Errorf("Should have seen only 3 events: %v", len(events))
	}

	events, err = persister.retrieveEvents("event_test", &model.RetrieveEventsCriteria{
		Offset:  1,
		Count:   5,
		Reverse: false,
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 2 {
		t.Errorf("Should have seen only 2 events: %v", len(events))
	}

	events, err = persister.retrieveEvents("event_test", &model.RetrieveEventsCriteria{
		Offset:  0,
		Count:   3,
		Reverse: true,
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 3 {
		t.Errorf("Should have seen only 2 event: %v", len(events))
	}
	if events[0].EventType() != event3.EventType() {
		t.Errorf("Should have seen the type of the most recent event: err: %v", err)
	}

	events, err = persister.retrieveEvents("event_test", &model.RetrieveEventsCriteria{
		Offset:    0,
		Count:     10,
		EventType: "Application",
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("Should have seen only 1 event: %v", len(events))
	}
	if events[0].EventType() != "Application" {
		t.Errorf("Should have seen the type application")
	}

	events, err = persister.retrieveEvents("event_test", &model.RetrieveEventsCriteria{
		Offset: 0,
		Count:  10,
		FromTs: event2.Timestamp(),
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("Should have seen only 1 event: %v", len(events))
	}
	if events[0].EventType() != "Challenge" {
		t.Errorf("Should have seen the type challenge")
	}

	events, err = persister.retrieveEvents("event_test", &model.RetrieveEventsCriteria{
		Offset:   0,
		Count:    10,
		BeforeTs: event2.Timestamp(),
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("Should have seen only 1 event: %v", len(events))
	}
	if events[0].EventType() != "Application" {
		t.Errorf("Should have seen the type application")
	}
}
