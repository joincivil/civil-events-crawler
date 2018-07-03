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
		return nil, err
	}
	pgPersister.db = db
	pgPersister.eventToLastBlockNumber = make(map[common.Address]map[string]PersisterBlockData)
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
	err := p.db.Select(&civilEventDB, "SELECT event_type, hash, contract_address, contract_name, timestamp, payload, log_payload FROM events_test;")
	return civilEventDB, err
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

// UpdateLastBlockData updates the last block number seen by the retriever
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
