// +build integration

// This is an integration test file for postgrespersister. Postgres needs to be running.
// Run this using go test -tags=integration
// Run benchmark test using go test -tags=integration -bench=.

// NOTE(IS): This only tests with civiltcr and newsroom contract events because we only have
// event wrappers around those contracts so far.
package persistence

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"math/big"
	"reflect"
	"strings"
	"testing"
)

const (
	eventTestTableName = "event_test"
	postgresPort       = 5432
	postgresDBName     = "civil_crawler"
	postgresUser       = "docker"
	postgresPswd       = "docker"
	postgresHost       = "localhost"
)

var (
	contractAddress      = "0x77e5aaBddb760FBa989A1C4B2CDd4aA8Fa3d311d"
	testAddress          = "0xDFe273082089bB7f70Ee36Eebcde64832FE97E55"
	testApplicationEvent = &contract.CivilTCRContractApplication{
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
	testApplicationWhitelistedEvent = &contract.CivilTCRContractApplicationWhitelisted{
		ListingAddress: common.HexToAddress(testAddress),
		Raw: types.Log{
			Address:     common.HexToAddress(testAddress),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: 8888886,
			TxHash:      common.Hash{},
			TxIndex:     3,
			BlockHash:   common.Hash{},
			Index:       2,
			Removed:     false,
		},
	}
	testChallengeEvent = &contract.CivilTCRContractChallenge{
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
			BlockNumber: 8888887,
			TxHash:      common.Hash{},
			TxIndex:     4,
			BlockHash:   common.Hash{},
			Index:       2,
			Removed:     false,
		},
	}
	testNwsrmNameChangedEvent = &contract.NewsroomContractNameChanged{
		NewName: "test newsroom",
		Raw: types.Log{
			Address:     common.HexToAddress(testAddress),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: 8888889,
			TxHash:      common.Hash{},
			TxIndex:     4,
			BlockHash:   common.Hash{},
			Index:       2,
			Removed:     false,
		},
	}
)

/*
Helpers for tests
*/
// TODO(IS) create a more realistic raw.log payload?
// Sets up an Application event and if rand=true, generates a random hash for transaction hash so that the hash in DB is unique.
func setupApplicationEvent(rand bool) (*model.Event, error) {
	if rand {
		randString, _ := randomHex(32)
		testApplicationEvent.Raw.TxHash = common.HexToHash(randString)
	}
	return model.NewEventFromContractEvent("Application", "CivilTCRContract", common.HexToAddress(contractAddress),
		testApplicationEvent, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
}

// Sets up an ApplicationWhitelisted event and if rand=true, generates a random hash for transaction hash so that the hash in DB is unique.
func setupApplicationWhitelistedEvent(rand bool) (*model.Event, error) {
	if rand {
		randString, _ := randomHex(32)
		testApplicationWhitelistedEvent.Raw.TxHash = common.HexToHash(randString)
	}
	return model.NewEventFromContractEvent("ApplicationWhitelisted", "CivilTCRContract", common.HexToAddress(contractAddress),
		testApplicationWhitelistedEvent, utils.CurrentEpochNanoSecsInInt64(), model.Watcher)
}

// Sets up an Challenge event and if rand=true, generates a random hash for transaction hash so that the hash in DB is unique.
func setupChallengeEvent(rand bool) (*model.Event, error) {
	if rand {
		randString, _ := randomHex(32)
		testChallengeEvent.Raw.TxHash = common.HexToHash(randString)
	}
	return model.NewEventFromContractEvent("Challenge", "CivilTCRContract", common.HexToAddress(contractAddress),
		testChallengeEvent, utils.CurrentEpochNanoSecsInInt64(), model.Filterer)
}

// Sets up a Newsroom Name Changed event and if rand=true, generates a random hash for transaction hash so that the hash in DB is unique.
func setupNewsroomNameChanged(rand bool) (*model.Event, error) {
	if rand {
		randString, _ := randomHex(32)
		testNwsrmNameChangedEvent.Raw.TxHash = common.HexToHash(randString)
	}
	return model.NewEventFromContractEvent("NameChanged", "NewsroomContract", common.HexToAddress(contractAddress),
		testNwsrmNameChangedEvent, utils.CurrentEpochNanoSecsInInt64(), model.Watcher)
}

// specify fields for testing purposes
func setupApplicationEventWithParams(rand bool, contractAddress string, timestamp int64) (*model.Event, error) {
	if rand {
		randString, _ := randomHex(32)
		testApplicationEvent.Raw.TxHash = common.HexToHash(randString)
	}
	return model.NewEventFromContractEvent("Application", "CivilTCRContract", common.HexToAddress(contractAddress),
		testApplicationEvent, timestamp, model.Filterer)
}

// Sets up all the above events and returns a list of test events
func setupEvents(rand bool) ([]*model.Event, error) {
	appEvent, err := setupApplicationEvent(rand)
	if err != nil {
		return nil, fmt.Errorf("Cannot setup Application event: %v", err)
	}
	appWhitelisted, err := setupApplicationWhitelistedEvent(rand)
	if err != nil {
		return nil, fmt.Errorf("Cannot setup ApplicationWhitelisted event: %v", err)
	}
	challenge, err := setupChallengeEvent(rand)
	if err != nil {
		return nil, fmt.Errorf("Cannot setup Challenge event: %v", err)
	}
	nameChanged, err := setupNewsroomNameChanged(rand)
	if err != nil {
		return nil, fmt.Errorf("Cannot setup NameChanged event: %v", err)
	}
	return []*model.Event{appEvent, appWhitelisted, challenge, nameChanged}, nil
}

// random hex string generation
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// Change block number of events for tests
func changeBlockData(blockNo int, event *model.Event) {
	event.LogPayload().BlockNumber = uint64(blockNo)
}

func setupDBConnection() (*PostgresPersister, error) {
	postgresPersister, err := NewPostgresPersister(postgresHost, postgresPort, postgresUser, postgresPswd, postgresDBName)
	return postgresPersister, err
}

// allEventsFromTable gets all events from table and is used for testing
func allEventsFromTable(persister *PostgresPersister, tableName string) ([]postgres.Event, error) {
	dbEvent := []postgres.Event{}
	queryString := fmt.Sprintf("SELECT event_type, hash, contract_address, contract_name, timestamp, payload, log_payload "+
		"FROM %s;", tableName)
	err := persister.db.Select(&dbEvent, queryString)
	return dbEvent, err
}

func setupTestTable() (*PostgresPersister, error) {
	persister, err := setupDBConnection()
	if err != nil {
		return persister, fmt.Errorf("Error connecting to DB: %v", err)
	}
	createTableQuery := postgres.CreateEventTableQueryString(eventTestTableName)
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

/*
Tests for postgres setup
*/

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

	indexCreationQuery := postgres.CreateEventTableIndicesString(eventTestTableName)
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

// Try to save the an event with the same payload hash twice. This should not work
func TestDuplicateEvents(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)
	// create 2 events w same payload hash
	event1, err := setupApplicationEvent(false)
	if err != nil {
		t.Errorf("Cannot setup Application event: %v", err)
	}
	event2, err := setupApplicationEvent(false)
	if err != nil {
		t.Errorf("Cannot setup Application event: %v", err)
	}
	if event1.Hash() != event2.Hash() {
		t.Errorf("Hashes for events should be equal, but they are %v, %v", event1.Hash(), event2.Hash())
	}
	civilEventsFromContract := []*model.Event{event1, event2}
	// save to database, catch the error
	err = persister.saveEventsToTable(civilEventsFromContract, eventTestTableName)
	if !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		t.Errorf("Error for duplicate key value should have been thrown")
	}

	err = deleteTestTable(persister)
	if err != nil {
		t.Error(err)
	}
}

// Test saving the same list of events twice to ensure that start block for filterers is being updated w persistence
func TestStartBlockUpdate(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)
	events, err := setupEvents(false)
	if err != nil {
		t.Errorf("Couldn't setup civilEvent from contract %v", err)
	}
	err = persister.saveEventsToTable(events, eventTestTableName)
	if err != nil {
		fmt.Errorf("Error saving events to table, %v", err)
	}
	err = persister.saveEventsToTable(events, eventTestTableName)
	if err != nil {
		fmt.Errorf("Error saving same events to table, %v", err)
	}

	err = deleteTestTable(persister)
	if err != nil {
		t.Error(err)
	}
}

/*
event table tests
*/

// TODO(IS): test that all queries can be performed w the instance of db, i.e. connection pools are being returned
func TestMultipleQueries(t *testing.T) {

}

func TestSaveEvents(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)

	events, err := setupEvents(true)
	if err != nil {
		t.Errorf("Couldn't setup civilEvent from contract %v", err)
	}

	// save each type of test event to table
	err = persister.saveEventsToTable(events, eventTestTableName)
	if err != nil {
		t.Errorf("Cannot save event to events_test table: %v", err)
	}

	civilEventsDB, err := allEventsFromTable(persister, eventTestTableName)
	if err != nil {
		t.Errorf("error querying event from events_test table: %v", err)
	}
	if len(civilEventsDB) != 4 {
		t.Errorf("expected there to be 4 events in table but there are %v events", len(civilEventsDB))
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
	numEventTypes := 4
	civilEventsFromContract := make([]*model.Event, 0)
	for i := 1; i <= numEvents; i++ {
		events, err := setupEvents(true)
		if err != nil {
			b.Errorf("Couldn't setup civilEvent from contract %v", err)
		}
		civilEventsFromContract = append(civilEventsFromContract, events...)
	}

	err = persister.saveEventsToTable(civilEventsFromContract, eventTestTableName)
	if err != nil {
		b.Errorf("Cannot save event to event_test table: %v", err)
	}
	var numRows int
	err = persister.db.QueryRow(`SELECT COUNT(*) FROM
                                        event_test`).Scan(&numRows)
	if numRows != numEvents*numEventTypes {
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
	events, err := setupEvents(true)
	if err != nil {
		t.Errorf("Couldn't setup civilEvents from contracts %v", err)
	}

	err = persister.UpdateLastBlockData(events)
	if err != nil {
		t.Errorf("Couldn't update last block data: %v", err)
	}
	for _, event := range events {
		if persister.LastBlockNumber(event.EventType(), event.ContractAddress()) != event.LogPayload().BlockNumber {
			t.Error("Blocknumber was not updated correctly in persistence")
		}
		if persister.LastBlockHash(event.EventType(), event.ContractAddress()) != event.LogPayload().BlockHash {
			t.Error("Blockhash was not updated correctly in persistence")
		}
	}
	// should actually change block no here?
	// now save 4 new events, see if persistence is updated
	events2, err := setupEvents(true)
	if err != nil {
		t.Errorf("Couldn't setup civilEvents from contracts %v", err)
	}

	err = persister.UpdateLastBlockData(events)
	if err != nil {
		t.Errorf("Couldn't update last block data: %v", err)
	}
	for _, event := range events2 {
		if persister.LastBlockNumber(event.EventType(), event.ContractAddress()) != event.LogPayload().BlockNumber {
			t.Error("Blocknumber was not updated correctly in persistence")
		}
		if persister.LastBlockHash(event.EventType(), event.ContractAddress()) != event.LogPayload().BlockHash {
			t.Error("Blockhash was not updated correctly in persistence")
		}
	}

}

func TestLatestEventsQuery(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)

	testEvents, err := setupEvents(true)
	if err != nil {
		t.Errorf("Couldn't setup event %v", err)
	}
	// create more events that are at a later time
	testEventsLatest, err := setupEvents(true)
	if err != nil {
		t.Errorf("Couldn't setup event %v", err)
	}
	testEvents = append(testEvents, testEventsLatest...)
	numEvents := len(testEventsLatest)
	// save events
	err = persister.saveEventsToTable(testEvents, eventTestTableName)
	if err != nil {
		t.Errorf("Cannot save event to event_test table: %v", err)
	}

	// retrieve events
	dbEvents, err := persister.getLatestEvents(eventTestTableName)
	if err != nil {
		t.Errorf("Error retrieving events: %v", err)
	}
	if len(dbEvents) != numEvents {
		t.Errorf("Query should have only returned %v events but there are %v", numEvents, len(dbEvents))
	}

	// check fields that would change are equal
	for i, event := range testEventsLatest {
		latestTimestamp := event.Timestamp()
		queryTimestamp := dbEvents[i].Timestamp
		if latestTimestamp != queryTimestamp {
			t.Errorf("Query didn't pull the latest event for contract %v and event type %v", event.ContractName(), event.EventType())
		}

	}
	err = deleteTestTable(persister)
	if err != nil {
		t.Error(err)
	}
}

func TestPopulateBlockDataFromDB(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)

	// create test events, fill test block data
	numEvents := 3
	civilEventsFromContract := make([]*model.Event, 0)
	blockNo := 8888888
	correctEventToBlockData := make(map[common.Address]map[string]PersisterBlockData)
	for i := 1; i <= numEvents; i++ {
		events, err := setupEvents(true)
		if err != nil {
			t.Errorf("Couldn't setup civilEvent from contract %v", err)
		}
		for _, event := range events {
			changeBlockData(blockNo, event)
			contractAddress := event.ContractAddress()
			eventType := event.EventType()

			blockData := PersisterBlockData{event.LogPayload().BlockNumber, event.LogPayload().BlockHash}
			if correctEventToBlockData[contractAddress] == nil {
				correctEventToBlockData[contractAddress] = make(map[string]PersisterBlockData)
			}
			correctEventToBlockData[contractAddress][eventType] = blockData
		}
		blockNo++
		civilEventsFromContract = append(civilEventsFromContract, events...)
	}

	// add events with a different contract address
	contractAddress, _ = randomHex(42)

	for i := 1; i <= numEvents; i++ {
		events, err := setupEvents(true)
		if err != nil {
			t.Errorf("Couldn't setup civilEvent from contract %v", err)
		}
		for _, event := range events {
			changeBlockData(blockNo, event)
			contractAddress := event.ContractAddress()
			eventType := event.EventType()
			blockData := PersisterBlockData{event.LogPayload().BlockNumber, event.LogPayload().BlockHash}
			if correctEventToBlockData[contractAddress] == nil {
				correctEventToBlockData[contractAddress] = make(map[string]PersisterBlockData)
			}
			correctEventToBlockData[contractAddress][eventType] = blockData
		}
		blockNo++
		civilEventsFromContract = append(civilEventsFromContract, events...)
	}

	// save events to table
	err = persister.saveEventsToTable(civilEventsFromContract, eventTestTableName)
	if err != nil {
		t.Errorf("Cannot save event to event_test table: %v", err)
	}

	// populate persistence
	err = persister.PopulateBlockDataFromDB(eventTestTableName)
	if err != nil {
		t.Errorf("Cannot fill persistence, %v", err)
	}

	if !reflect.DeepEqual(correctEventToBlockData, persister.eventToBlockData) {
		t.Errorf("eventToBlockData was not populated correctly. it should be %v but is %v", correctEventToBlockData,
			persister.eventToBlockData)
	}

	err = deleteTestTable(persister)
	if err != nil {
		t.Error(err)
	}
}

// Two events with the same timestamp
// what even happens when 2 events have the same timestamp and you try to fill persistence
func TestSameTimestampEvents(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)

	// setup 2 challenge events (same contract address) with the same timestamp..
	// 2 cases
	// 1. events have the same timestamp and block number (this would rarely happen), but same block number will be taken
	// 2. events have the same timestamp and different block number. take the higher block number (I don't think this would ever happen)

	// case 1
	numEvents := 2
	civilEventsFromContract := make([]*model.Event, numEvents)
	timestamp := utils.CurrentEpochNanoSecsInInt64()
	contractAddress, _ := randomHex(42)

	for i := 0; i < numEvents; i++ {
		event, err := setupApplicationEventWithParams(true, contractAddress, timestamp)
		if err != nil {
			t.Errorf("Couldn't setup civilEvent from contract %v", err)
		}
		civilEventsFromContract[i] = event
	}

	// save events to table
	err = persister.saveEventsToTable(civilEventsFromContract, eventTestTableName)
	if err != nil {
		t.Errorf("Cannot save event to event_test table: %v", err)
	}

	// populate persistence
	err = persister.PopulateBlockDataFromDB(eventTestTableName)
	if err != nil {
		t.Errorf("Cannot fill persistence, %v", err)
	}

	// check peristence
	if persister.eventToBlockData[common.HexToAddress(contractAddress)]["Application"].BlockNumber !=
		civilEventsFromContract[0].LogPayload().BlockNumber {
		t.Errorf("Block numbers are not equal: %v, %v", persister.eventToBlockData[common.HexToAddress(contractAddress)]["Application"].BlockNumber,
			civilEventsFromContract[0].LogPayload().BlockNumber)
	}

	// case 2
	civilEventsFromContract2 := make([]*model.Event, numEvents)
	timestamp2 := utils.CurrentEpochNanoSecsInInt64()
	blockNo := 8888888

	for i := 0; i < numEvents; i++ {
		event, err := setupApplicationEventWithParams(true, contractAddress, timestamp2)
		if err != nil {
			t.Errorf("Couldn't setup civilEvent from contract %v", err)
		}
		blockNo++
		changeBlockData(blockNo, event)
		civilEventsFromContract2[i] = event
	}

	// save events to table
	err = persister.saveEventsToTable(civilEventsFromContract2, eventTestTableName)
	if err != nil {
		t.Errorf("Cannot save event to event_test table: %v", err)
	}

	// populate persistence
	err = persister.PopulateBlockDataFromDB(eventTestTableName)
	if err != nil {
		t.Errorf("Cannot fill persistence, %v", err)
	}

	// check peristence
	if civilEventsFromContract2[1].LogPayload().BlockNumber !=
		persister.eventToBlockData[common.HexToAddress(contractAddress)]["Application"].BlockNumber {
		t.Errorf("Block number is not what it should be")
	}

	err = deleteTestTable(persister)
	if err != nil {
		t.Error(err)
	}
}

// This conversion needs to be here, bc we need the actual event after being saved in DB.
func TestDBToEvent(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)

	appEvent, err := setupApplicationEvent(true)
	if err != nil {
		t.Errorf("setupEvent should have succeeded: err: %v", err)
	}

	// Get this event from DB
	events := []*model.Event{appEvent}
	err = persister.saveEventsToTable(events, eventTestTableName)
	if err != nil {
		t.Errorf("Cannot save event to event_test table: %v", err)
	}

	appEventDB, err := allEventsFromTable(persister, eventTestTableName)

	dbEvent := appEventDB[0]

	appEventFromDB, err := dbEvent.DBToEventData()
	if err != nil {
		t.Errorf("Could not convert db event back to civilevent: err: %v", err)
	}

	// deep equal doesn't work bc of nested slices, etc. so just compare each element
	if appEvent.ContractAddress() != appEventFromDB.ContractAddress() {
		t.Errorf("ContractAddress not equal: %v %v", appEvent.ContractAddress(), appEventFromDB.ContractAddress())
	}
	if appEvent.ContractName() != appEventFromDB.ContractName() {
		t.Errorf("ContractName not equal: %v %v", appEvent.ContractName(), appEventFromDB.ContractName())
	}
	if appEvent.Hash() != appEventFromDB.Hash() {
		t.Errorf("Hash not equal: %v %v", appEvent.Hash(), appEventFromDB.Hash())
	}
	if appEvent.EventType() != appEventFromDB.EventType() {
		t.Errorf("EventType not equal: %v %v", appEvent.EventType(), appEventFromDB.EventType())
	}
	if appEvent.Timestamp() != appEventFromDB.Timestamp() {
		t.Errorf("Timestamp not equal: %v %v", appEvent.Timestamp(), appEventFromDB.Timestamp())
	}

	// EventPayload
	if !reflect.DeepEqual(appEventFromDB.EventPayload(), appEvent.EventPayload()) {
		t.Errorf("EventPayloads not equal: %v %v", appEvent.EventPayload(), appEventFromDB.EventPayload())
	}

	// LogPayload
	civilLogPayload := appEvent.LogPayload()
	civilLogFromDBPayload := appEventFromDB.LogPayload()

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
	err = deleteTestTable(persister)
	if err != nil {
		t.Error(err)
	}

}

func TestRetrieveEvents(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)
	civilEventsFromContract, err := setupEvents(true)
	if err != nil {
		t.Errorf("Couldn't setup event %v", err)
	}

	err = persister.saveEventsToTable(civilEventsFromContract, eventTestTableName)
	if err != nil {
		t.Errorf("Should not have seen error when saving events to table: %v", err)
	}

	events, err := persister.retrieveEventsFromTable(eventTestTableName, &model.RetrieveEventsCriteria{
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
	if events[0].EventType() != civilEventsFromContract[0].EventType() {
		t.Errorf("Should have seen the type of the oldest event: err: %v", err)
	}

	events, err = persister.retrieveEventsFromTable(eventTestTableName, &model.RetrieveEventsCriteria{
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

	events, err = persister.retrieveEventsFromTable(eventTestTableName, &model.RetrieveEventsCriteria{
		Offset:  0,
		Count:   5,
		Reverse: false,
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 4 {
		t.Errorf("Should have seen only 4 events: %v", len(events))
	}

	events, err = persister.retrieveEventsFromTable(eventTestTableName, &model.RetrieveEventsCriteria{
		Offset:  1,
		Count:   5,
		Reverse: false,
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 3 {
		t.Errorf("Should have seen only 3 events: %v", len(events))
	}

	events, err = persister.retrieveEventsFromTable(eventTestTableName, &model.RetrieveEventsCriteria{
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
	if events[0].EventType() != civilEventsFromContract[3].EventType() {
		t.Errorf("Should have seen the type of the most recent event: err: %v", err)
	}

	events, err = persister.retrieveEventsFromTable(eventTestTableName, &model.RetrieveEventsCriteria{
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

	events, err = persister.retrieveEventsFromTable(eventTestTableName, &model.RetrieveEventsCriteria{
		Offset: 0,
		Count:  10,
		FromTs: civilEventsFromContract[1].Timestamp(),
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 2 {
		t.Errorf("Should have seen 2 events: %v", len(events))
	}
	if events[0].EventType() != "Challenge" {
		t.Errorf("Should have seen the type challenge")
	}

	events, err = persister.retrieveEventsFromTable(eventTestTableName, &model.RetrieveEventsCriteria{
		Offset:   0,
		Count:    10,
		BeforeTs: civilEventsFromContract[1].Timestamp(),
	})
	if err != nil {
		t.Errorf("Should not have received error when retrieving events: err: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("Should have seen only 1: %v", len(events))
	}
	if events[0].EventType() != "Application" {
		t.Errorf("Should have seen the type application")
	}

	err = deleteTestTable(persister)
	if err != nil {
		t.Error(err)
	}
}
