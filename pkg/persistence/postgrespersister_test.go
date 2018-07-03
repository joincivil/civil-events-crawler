// +build integration

// This is an integration test file for postgrespersister. Postgres needs to be running.
// Run this using go test -tags=integration.
package persistence

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"math/big"
	"testing"
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
)

func setupCivilEvent() (*model.CivilEvent, error) {
	return model.NewCivilEventFromContractEvent("Application", "CivilTCRContract", common.HexToAddress(contractAddress),
		testEvent, utils.CurrentEpochSecsInInt())
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
	_, err = persister.db.Query("CREATE TABLE events_test AS SELECT * FROM events WHERE 1=2;")
	if err != nil {
		return persister, fmt.Errorf("Couldn't create test table %v", err)
	}
	return persister, nil
}

func deleteTestTable(persister *PostgresPersister) error {
	_, err := persister.db.Query("DROP TABLE events_test;")
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
                                        AND    table_name = 'events'
                                        );`).Scan(&exists)
	if err != nil {
		t.Errorf("Couldn't get table")
	}
	if !exists {
		t.Errorf("events table does not exist")
	}

}

func TestSaveToEventsTestTable(t *testing.T) {
	persister, err := setupTestTable()
	if err != nil {
		t.Error(err)
	}
	defer deleteTestTable(persister)
	event, err := setupCivilEvent()
	if err != nil {
		t.Errorf("Couldn't setup civilEvent from contract %v", err)
	}

	civilEventsFromContract := []*model.CivilEvent{event}
	err = persister.saveEventsToTable(civilEventsFromContract, "events_test")
	if err != nil {
		t.Errorf("Cannot save event to events_test table: %v", err)
	}

	civilEventDB, err := persister.GetEvents("events_test")
	if err != nil {
		t.Errorf("error querying event from events_test table: %v", err)
	}
	_ = &civilEventDB[0]
	err = deleteTestTable(persister)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkTestSavingManyEventsToEventsTestTable(b *testing.B) {
	persister, err := setupTestTable()
	if err != nil {
		b.Error(err)
	}
	defer deleteTestTable(persister)
	event, err := setupCivilEvent()
	if err != nil {
		b.Errorf("Couldn't setup civilEvent from contract %v", err)
	}

	numEvents := 100
	civilEventsFromContract := make([]*model.CivilEvent, 0)
	for i := 1; i <= numEvents; i++ {
		civilEventsFromContract = append(civilEventsFromContract, event)
	}
	err = persister.saveEventsToTable(civilEventsFromContract, "events_test")
	if err != nil {
		b.Errorf("Cannot save event to events_test table: %v", err)
	}
	var numRows int
	err = persister.db.QueryRow(`SELECT COUNT(*) FROM
                                        events_test`).Scan(&numRows)
	if numRows != numEvents {
		b.Errorf("Number of rows in events_test table should be %v but it is %v", numEvents, numRows)
	}
	err = deleteTestTable(persister)
	if err != nil {
		b.Error(err)
	}
}

func TestPersistence(t *testing.T) {
	//Check that persistence is being updated
	persister, err := setupDBConnection()
	if err != nil {
		t.Errorf("Error connecting to DB: %v", err)
	}
	event, err := setupCivilEvent()
	if err != nil {
		t.Errorf("Couldn't setup civilEvent from contract %v", err)
	}
	civilEventsFromContract := []*model.CivilEvent{event}
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
