// Package persistence contains components to interact with the DB
package persistence // import "github.com/joincivil/civil-events-crawler/pkg/persistence"

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"
	// driver for postgresql
	_ "github.com/lib/pq"
)

// NewPostgresPersister creates a new postgres persister
func NewPostgresPersister(host string, port int, user string, password string, dbname string) (*PostgresPersister, error) {
	pgPersister := &PostgresPersister{}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return pgPersister, fmt.Errorf("Error connecting to sqlx: %v", err)
	}
	pgPersister.db = db
	pgPersister.eventToLastBlockNumber = make(map[common.Address]map[string]PersisterBlockData)
	// Batch query here for latest block numbers to fill persistence
	err = pgPersister.fillPersistence("events")
	if err != nil {
		return pgPersister, fmt.Errorf("Error with filling persistence from DB: %v", err)
	}
	return pgPersister, nil
}

// PostgresPersister holds the DB connection and persistence
type PostgresPersister struct {
	eventToLastBlockNumber map[common.Address]map[string]PersisterBlockData
	db                     *sqlx.DB
}

// CreateTables creates tables in DB if they don't exist
func (p *PostgresPersister) CreateTables() error {
	// TODO(IS): per PN's advice have some logic to determine which models need to be part
	// of this DB to run create for those set of models.
	schema := postgres.EventsTableSchema()
	_, err := p.db.Exec(schema)
	return err
}

// SaveEvents saves events to events table in DB
func (p *PostgresPersister) SaveEvents(events []*model.CivilEvent) error {
	return p.saveEventsToTable(events, "events")
}

func (p *PostgresPersister) saveEventsToTable(events []*model.CivilEvent, tableName string) error {
	var err error
	for _, event := range events {
		dbEvent, dbEventErr := postgres.NewCivilEvent(event)
		if dbEventErr != nil {
			return dbEventErr
		}
		query := p.getInsertEventQueryString(tableName)
		err = p.saveEvent(dbEvent, query)
		if err != nil {
			return err
		}
	}
	return err
}

func (p *PostgresPersister) saveEvent(dbEvent *postgres.CivilEvent, query string) error {
	_, err := p.db.NamedExec(query, dbEvent)
	return err
}

func (p *PostgresPersister) getInsertEventQueryString(tableName string) string {
	return fmt.Sprintf("INSERT INTO %s (event_type, hash, contract_address, contract_name, timestamp, payload, log_payload)"+
		" VALUES (:event_type, :hash, :contract_address, :contract_name, :timestamp, :payload, :log_payload)",
		tableName)
}

// GetEvents gets all events from table
// NOTE: this function gets all events from table for now.
func (p *PostgresPersister) GetEvents(tableName string) ([]postgres.CivilEvent, error) {
	civilEventDB := []postgres.CivilEvent{}
	queryString := fmt.Sprintf("SELECT event_type, hash, contract_address, contract_name, timestamp, payload, log_payload "+
		"FROM %s;", tableName)
	err := p.db.Select(&civilEventDB, queryString)
	return civilEventDB, err
}

// fillPersistence will fill the persistence with data from the DB.
func (p *PostgresPersister) fillPersistence(tableName string) error {
	events, err := p.getLatestEvents(tableName)
	if err != nil {
		return err
	}
	// TODO: if you have 2 events w the same timestamp, just take the max block number
	for _, event := range events {
		civilEvent, err := event.DBToEventData()
		if err != nil {
			return err
		}
		blockData := PersisterBlockData{}
		logPayload := civilEvent.LogPayload()
		blockData.BlockNumber = logPayload.BlockNumber
		blockData.BlockHash = logPayload.BlockHash
		contractAddress := civilEvent.ContractAddress()
		p.eventToLastBlockNumber[contractAddress] = make(map[string]PersisterBlockData)
		p.eventToLastBlockNumber[contractAddress][civilEvent.EventType()] = blockData
	}
	return nil
}

func (p *PostgresPersister) getLatestEvents(tableName string) ([]postgres.CivilEvent, error) {
	query := p.retrieveLatestEventsQueryString(tableName)
	events := []postgres.CivilEvent{}
	err := p.db.Select(&events, query)
	return events, err
}

func (p *PostgresPersister) retrieveLatestEventsQueryString(tableName string) string {
	// Query for the latest timestamp.
	return fmt.Sprintf("SELECT e.event_type, e.log_payload, e.payload, e.hash, e.contract_address, e.contract_name, max_e.timestamp "+
		"FROM (SELECT event_type, contract_address, MAX(timestamp) AS timestamp FROM %s GROUP BY event_type, contract_address) max_e "+
		"JOIN %s e ON e.event_type = max_e.event_type AND e.timestamp = max_e.timestamp;", tableName, tableName)
}

// LastBlockNumber returns the last block number seen by the persister
func (p *PostgresPersister) LastBlockNumber(eventType string, contractAddress common.Address) uint64 {
	if p.eventToLastBlockNumber[contractAddress] == nil {
		p.eventToLastBlockNumber[contractAddress] = make(map[string]PersisterBlockData)
	}
	return p.eventToLastBlockNumber[contractAddress][eventType].BlockNumber
}

// LastBlockHash returns the last block hash seen by the persister
func (p *PostgresPersister) LastBlockHash(eventType string, contractAddress common.Address) common.Hash {
	if p.eventToLastBlockNumber[contractAddress] == nil {
		p.eventToLastBlockNumber[contractAddress] = make(map[string]PersisterBlockData)
	}
	return p.eventToLastBlockNumber[contractAddress][eventType].BlockHash
}

// UpdateLastBlockData updates the last block number seen by the persister
func (p *PostgresPersister) UpdateLastBlockData(events []*model.CivilEvent) error {
	for _, event := range events {
		err := p.parseEventToPersist(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PostgresPersister) parseEventToPersist(event *model.CivilEvent) error {
	eventType := event.EventType()
	contractAddress := event.ContractAddress()
	blockNumber := event.BlockNumber()
	blockHash := event.BlockHash()
	if p.eventToLastBlockNumber[contractAddress] == nil {
		p.eventToLastBlockNumber[contractAddress] = make(map[string]PersisterBlockData)
	}
	p.eventToLastBlockNumber[contractAddress][eventType] = PersisterBlockData{blockNumber, blockHash}
	return nil
}

// PersisterBlockData is the data about block stored for persistence
type PersisterBlockData struct {
	BlockNumber uint64
	BlockHash   common.Hash
}
