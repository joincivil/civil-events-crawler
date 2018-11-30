// Package persistence contains components to interact with the DB
package persistence // import "github.com/joincivil/civil-events-crawler/pkg/persistence"
// NOTE(IS): Only tested with event wrappers around contracts implemented: civiltcr and newsroom

import (
	"bytes"
	"fmt"
	"strings"

	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"

	// driver for postgresql
	_ "github.com/lib/pq"

	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"
)

const (
	eventTableName = "event"
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
	pgPersister.eventToBlockData = make(map[common.Address]map[string]PersisterBlockData)
	return pgPersister, nil
}

// PostgresPersister holds the DB connection and persistence
type PostgresPersister struct {
	eventToBlockData map[common.Address]map[string]PersisterBlockData
	db               *sqlx.DB
}

// SaveEvents saves events to events table in DB
func (p *PostgresPersister) SaveEvents(events []*model.Event) error {
	return p.saveEventsToTable(events, eventTableName)
}

// RetrieveEvents retrieves the Events given an offset, count, and asc/dec bool
func (p *PostgresPersister) RetrieveEvents(criteria *model.RetrieveEventsCriteria) ([]*model.Event, error) {
	return p.retrieveEventsFromTable(eventTableName, criteria)
}

// LastBlockNumber returns the last block number seen by the persister
func (p *PostgresPersister) LastBlockNumber(eventType string, contractAddress common.Address) uint64 {
	if p.eventToBlockData[contractAddress] == nil {
		p.eventToBlockData[contractAddress] = make(map[string]PersisterBlockData)
	}
	return p.eventToBlockData[contractAddress][eventType].BlockNumber
}

// LastBlockHash returns the last block hash seen by the persister
func (p *PostgresPersister) LastBlockHash(eventType string, contractAddress common.Address) common.Hash {
	if p.eventToBlockData[contractAddress] == nil {
		p.eventToBlockData[contractAddress] = make(map[string]PersisterBlockData)
	}
	return p.eventToBlockData[contractAddress][eventType].BlockHash
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

// PopulateBlockDataFromDB will fill the persistence with data from the DB.
func (p *PostgresPersister) PopulateBlockDataFromDB(tableName string) error {
	events, err := p.getLatestEvents(tableName)
	if err != nil {
		return err
	}
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
		if p.eventToBlockData[contractAddress] == nil {
			p.eventToBlockData[contractAddress] = make(map[string]PersisterBlockData)
		}
		p.eventToBlockData[contractAddress][modelEvent.EventType()] = blockData
	}
	return nil
}

// CreateTables creates tables in DB if they don't exist
func (p *PostgresPersister) CreateTables() error {
	eventTableQuery := postgres.CreateEventTableQuery()
	_, err := p.db.Exec(eventTableQuery)
	return err
}

// CreateIndices creates the indices for DB if they don't exist
func (p *PostgresPersister) CreateIndices() error {
	indexQuery := postgres.CreateEventTableIndices()
	_, err := p.db.Exec(indexQuery)
	return err
}

func (p *PostgresPersister) saveEventsToTable(events []*model.Event, tableName string) error {
	queryString := postgres.InsertIntoDBQueryString(tableName, postgres.Event{})
	// There is no way to batch insert using sqlx, so doing a loop here
	for _, event := range events {
		err := p.saveEventToTable(queryString, event)
		if err != nil {
			return err
		}
		log.Infof("saveEventsToTable: saved: %v, %v", event.EventType(), event.TxHash().Hex()) // Debug, remove later
	}
	return nil
}

func (p *PostgresPersister) retrieveEventsFromTable(tableName string, criteria *model.RetrieveEventsCriteria) ([]*model.Event, error) {
	queryString := p.retrieveEventsQuery(tableName, criteria)
	dbEvents := []postgres.Event{}
	nstmt, err := p.db.PrepareNamed(queryString)
	if err != nil {
		return nil, fmt.Errorf("Error preparing query with sqlx: %v", err)
	}
	err = nstmt.Select(&dbEvents, criteria)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving events from table: %v", err)
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

func (p *PostgresPersister) retrieveEventsQuery(tableName string, criteria *model.RetrieveEventsCriteria) string {
	queryBuf := bytes.NewBufferString("SELECT ")
	fields, _ := postgres.StructFieldsForQuery(postgres.Event{}, false)
	queryBuf.WriteString(fields)    // nolint: gosec
	queryBuf.WriteString(" FROM ")  // nolint: gosec
	queryBuf.WriteString(tableName) // nolint: gosec
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
		queryBuf.WriteString(" ORDER BY id DESC") // nolint: gosec
	} else {
		queryBuf.WriteString(" ORDER BY id") // nolint: gosec
	}
	if criteria.Offset > 0 {
		queryBuf.WriteString(" OFFSET :offset") // nolint: gosec
	}
	if criteria.Count > 0 {
		queryBuf.WriteString(" LIMIT :count") // nolint: gosec
	}
	return queryBuf.String()
}

func (p *PostgresPersister) saveEventToTable(query string, event *model.Event) error {
	dbEvent, err := postgres.NewDbEventFromEvent(event)
	if err != nil {
		return err
	}
	_, err = p.db.NamedExec(query, dbEvent)
	if err != nil {
		return fmt.Errorf("Error saving event to table: err %v: event: %T", err, dbEvent.LogPayload["Data"])
	}
	log.Infof("saveEventToTable: done") // Debug, remove later
	return nil
}

func (p *PostgresPersister) getLatestEvents(tableName string) ([]postgres.Event, error) {
	queryString := p.retrieveLatestEventsQueryString(tableName)
	events := []postgres.Event{}
	err := p.db.Select(&events, queryString)
	return events, err
}

// retrieveLatestEventsQueryString queries for the events with the latest timestamp given an event type and contract address
func (p *PostgresPersister) retrieveLatestEventsQueryString(tableName string) string {
	// Query for the latest timestamp.
	return fmt.Sprintf( // nolint: gosec
		`SELECT e.event_type, e.log_payload, e.payload, e.hash, e.contract_address, e.contract_name, max_e.timestamp
		FROM (SELECT event_type, contract_address, MAX(timestamp) AS timestamp FROM %s GROUP BY event_type, contract_address) max_e
		JOIN %s e ON e.event_type = max_e.event_type AND e.timestamp = max_e.timestamp AND e.contract_address = max_e.contract_address;
        `,
		tableName,
		tableName,
	)
}

func (p *PostgresPersister) parseEventToPersist(event *model.Event) error {
	eventType := event.EventType()
	contractAddress := event.ContractAddress()
	blockNumber := event.BlockNumber()
	blockHash := event.BlockHash()
	if p.eventToBlockData[contractAddress] == nil {
		p.eventToBlockData[contractAddress] = make(map[string]PersisterBlockData)
	}
	p.eventToBlockData[contractAddress][eventType] = PersisterBlockData{blockNumber, blockHash}
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
