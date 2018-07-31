// Package persistence contains components to interact with the DB
package persistence // import "github.com/joincivil/civil-events-crawler/pkg/persistence"

import (
	"bytes"
	"fmt"
	log "github.com/golang/glog"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	// driver for postgresql
	_ "github.com/lib/pq"

	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"
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
	return pgPersister, nil
}

// PostgresPersister holds the DB connection and persistence
type PostgresPersister struct {
	eventToLastBlockNumber map[common.Address]map[string]PersisterBlockData
	db                     *sqlx.DB
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
func (p *PostgresPersister) UpdateLastBlockData(events []*model.Event) error {
	for _, event := range events {
		err := p.parseEventToPersist(event)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveEvents saves events to events table in DB
func (p *PostgresPersister) SaveEvents(events []*model.Event) error {
	return p.saveEventsToTable(events, "event")
}

// RetrieveEvents retrieves the Events given an offset, count, and asc/dec bool
func (p *PostgresPersister) RetrieveEvents(criteria *model.RetrieveEventsCriteria) (
	[]*model.Event, error) {
	return p.retrieveEvents("event", criteria)
}

// CreateTables creates tables in DB if they don't exist
func (p *PostgresPersister) CreateTables() error {
	// TODO(IS): per PN's advice have some logic to determine which models need to be part
	// of this DB to run create for those set of models.
	schema := postgres.EventTableSchema()
	_, err := p.db.Exec(schema)
	return err
}

// CreateIndices creates the indices for DB if they don't exist
func (p *PostgresPersister) CreateIndices() error {
	indexQuery := postgres.EventTableIndices()
	_, err := p.db.Exec(indexQuery)
	return err
}

func (p *PostgresPersister) saveEventsToTable(events []*model.Event, tableName string) error {
	var err error
	for _, event := range events {
		dbEvent, dbEventErr := postgres.NewDbEventFromEvent(event)
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

func (p *PostgresPersister) saveEvent(dbEvent *postgres.Event, query string) error {
	_, err := p.db.NamedExec(query, dbEvent)
	return err
}

func (p *PostgresPersister) getInsertEventQueryString(tableName string) string {
	return fmt.Sprintf("INSERT INTO %s (event_type, hash, contract_address, contract_name, timestamp, payload, log_payload)"+
		" VALUES (:event_type, :hash, :contract_address, :contract_name, :timestamp, :payload, :log_payload);",
		tableName)
}

// GetAllEvents gets all events from table
func (p *PostgresPersister) GetAllEvents(tableName string) ([]postgres.Event, error) {
	dbEvent := []postgres.Event{}
	queryString := fmt.Sprintf("SELECT event_type, hash, contract_address, contract_name, timestamp, payload, log_payload "+
		"FROM %s;", tableName)
	err := p.db.Select(&dbEvent, queryString)
	return dbEvent, err
}

func (p *PostgresPersister) retrieveEvents(tableName string, criteria *model.RetrieveEventsCriteria) (
	[]*model.Event, error) {
	dbEvents, err := p.retrieveDbEvents(tableName, criteria)
	if err != nil {
		return nil, err
	}
	events := make([]*model.Event, len(dbEvents))
	for index, dbEvent := range dbEvents {
		modelEvent, err := dbEvent.DBToEventData()
		if err != nil {
			log.Errorf("Error converting db to event data: err: %v", err)
			continue
		}
		events[index] = modelEvent
	}
	return events, nil
}

func (p *PostgresPersister) retrieveDbEvents(tableName string, criteria *model.RetrieveEventsCriteria) (
	[]postgres.Event, error) {
	queryBuf := bytes.NewBufferString("SELECT")
	queryBuf.WriteString(
		" event_type, hash, contract_address, contract_name, timestamp, payload, log_payload",
	) // nolint: gosec
	queryBuf.WriteString(fmt.Sprintf(" FROM %v", tableName)) // nolint: gosec

	if criteria.FromTs > 0 {
		queryBuf.WriteString(" WHERE timestamp > :fromts") // nolint: gosec
	}
	if criteria.BeforeTs > 0 {
		p.addWhereAnd(queryBuf)
		queryBuf.WriteString(" timestamp < :beforets") // nolint: gosec
	}
	if criteria.EventType != "" {
		p.addWhereAnd(queryBuf)
		queryBuf.WriteString(" event_type = :eventtype") // nolint: gosec
	}
	if criteria.Reverse {
		queryBuf.WriteString(" ORDER BY timestamp DESC") // nolint: gosec
	}
	if criteria.Offset > 0 {
		queryBuf.WriteString(" OFFSET :offset") // nolint: gosec
	}
	if criteria.Count > 0 {
		queryBuf.WriteString(" LIMIT :count") // nolint: gosec
	}

	dbEvent := []postgres.Event{}
	nstmt, err := p.db.PrepareNamed(queryBuf.String())
	if err != nil {
		return nil, err
	}
	err = nstmt.Select(&dbEvent, criteria)
	return dbEvent, err
}

// PopulateBlockDataFromDB will fill the persistence with data from the DB.
func (p *PostgresPersister) PopulateBlockDataFromDB(tableName string) error {
	events, err := p.getLatestEvents(tableName)
	if err != nil {
		return err
	}
	// TODO: if you have 2 events w the same timestamp, just take the max block number
	for _, event := range events {
		modelEvent, err := event.DBToEventData()
		if err != nil {
			return err
		}
		blockData := PersisterBlockData{}
		logPayload := modelEvent.LogPayload()
		blockData.BlockNumber = logPayload.BlockNumber
		blockData.BlockHash = logPayload.BlockHash
		contractAddress := modelEvent.ContractAddress()
		p.eventToLastBlockNumber[contractAddress] = make(map[string]PersisterBlockData)
		p.eventToLastBlockNumber[contractAddress][modelEvent.EventType()] = blockData
	}
	return nil
}

func (p *PostgresPersister) getLatestEvents(tableName string) ([]postgres.Event, error) {
	query := p.retrieveLatestEventsQueryString(tableName)
	events := []postgres.Event{}
	err := p.db.Select(&events, query)
	return events, err
}

func (p *PostgresPersister) retrieveLatestEventsQueryString(tableName string) string {
	// Query for the latest timestamp.
	return fmt.Sprintf("SELECT e.event_type, e.log_payload, e.payload, e.hash, e.contract_address, e.contract_name, max_e.timestamp "+
		"FROM (SELECT event_type, contract_address, MAX(timestamp) AS timestamp FROM %s GROUP BY event_type, contract_address) max_e "+
		"JOIN %s e ON e.event_type = max_e.event_type AND e.timestamp = max_e.timestamp AND e.contract_address = max_e.contract_address;", tableName, tableName)
}

func (p *PostgresPersister) parseEventToPersist(event *model.Event) error {
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

func (p *PostgresPersister) addWhereAnd(buf *bytes.Buffer) {
	if !strings.Contains(buf.String(), "WHERE") {
		buf.WriteString(" WHERE") // nolint: gosec
	} else {
		buf.WriteString(" AND") // nolint: gosec
	}
}

// PersisterBlockData is the data about block stored for persistence
type PersisterBlockData struct {
	BlockNumber uint64
	BlockHash   common.Hash
}
