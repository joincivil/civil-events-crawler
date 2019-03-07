// Package persistence contains components to interact with the DB
package persistence // import "github.com/joincivil/civil-events-crawler/pkg/persistence"

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"

	// driver for postgresql
	_ "github.com/lib/pq"

	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

	cpersist "github.com/joincivil/go-common/pkg/persistence"
	cpostgres "github.com/joincivil/go-common/pkg/persistence/postgres"
	ctime "github.com/joincivil/go-common/pkg/time"
)

const (
	crawlerServiceName = "crawler"
	// Could make this configurable later if needed
	maxOpenConns    = 50
	maxIdleConns    = 10
	connMaxLifetime = time.Second * 1800 // 30 mins
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
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	pgPersister.eventToBlockData = make(map[common.Address]map[string]PersisterBlockData)
	return pgPersister, nil
}

// PostgresPersister holds the DB connection and persistence
type PostgresPersister struct {
	eventToBlockData map[common.Address]map[string]PersisterBlockData
	db               *sqlx.DB
	version          *string
}

// PersisterVersion returns and sets the latest version of this persistence
func (p *PostgresPersister) PersisterVersion() (*string, error) {
	return p.persisterVersionFromTable(postgres.VersionTableName)
}

// GetTableName formats tabletype with version of this persister to return the table name
func (p *PostgresPersister) GetTableName(tableType string) string {
	return fmt.Sprintf("%s_%s", tableType, *p.version)
}

// SaveVersion saves the version for this persistence
func (p *PostgresPersister) SaveVersion(versionNumber *string) error {
	err := p.saveVersionToTable(postgres.VersionTableName, versionNumber)
	if err != nil {
		return err
	}
	p.version = versionNumber
	return nil
}

// SaveEvents saves events to events table in DB
func (p *PostgresPersister) SaveEvents(events []*model.Event) error {
	eventTableName := p.GetTableName(postgres.EventTableBaseName)
	return p.saveEventsToTable(events, eventTableName)
}

// RetrieveEvents retrieves the Events given an offset, count, and asc/dec bool. Ordered by db id.
func (p *PostgresPersister) RetrieveEvents(criteria *model.RetrieveEventsCriteria) ([]*model.Event, error) {
	eventTableName := p.GetTableName(postgres.EventTableBaseName)
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
func (p *PostgresPersister) PopulateBlockDataFromDB(tableType string) error {
	tableName := p.GetTableName(tableType)
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

// CreateVersionTable creates the version table and sets the version with new or existing version
func (p *PostgresPersister) CreateVersionTable(version *string) error {
	versionTableQuery := postgres.CreateVersionTableQuery(postgres.VersionTableName)
	_, err := p.db.Exec(versionTableQuery)
	if err != nil {
		return err
	}

	currentVersion, err := p.PersisterVersion()
	if err != nil && version == nil {
		if err == cpersist.ErrPersisterNoResults {
			return fmt.Errorf("No version in version table, specify a version: err %v", err)
		}
		return err
	}
	if currentVersion == version {
		return nil
	}
	if version == nil {
		// NOTE(IS): Use existing version, but update timestamp
		version = currentVersion
	}
	p.version = version
	return p.SaveVersion(version)
}

// CreateEventTable creates event table
func (p *PostgresPersister) CreateEventTable() error {
	eventTableName := p.GetTableName(postgres.EventTableBaseName)
	eventTableQuery := postgres.CreateEventTableQuery(eventTableName)
	_, err := p.db.Exec(eventTableQuery)
	return err
}

// CreateIndices creates the indices for DB if they don't exist
func (p *PostgresPersister) CreateIndices() error {
	eventTableName := p.GetTableName(postgres.EventTableBaseName)
	indexQuery := postgres.CreateEventTableIndices(eventTableName)
	_, err := p.db.Exec(indexQuery)
	return err
}

func (p *PostgresPersister) persisterVersionFromTable(tableName string) (*string, error) {
	if p.version == nil {
		version, err := p.retrieveVersionFromTable(tableName)
		if err != nil {
			return nil, err
		}
		p.version = version
	}
	return p.version, nil
}

func (p *PostgresPersister) retrieveVersionFromTable(tableName string) (*string, error) {
	dbVersion := []postgres.Version{}
	queryString := fmt.Sprintf(`SELECT * FROM %s WHERE service_name=$1 ORDER BY last_updated_timestamp DESC LIMIT 1;`, tableName) // nolint: gosec
	err := p.db.Select(&dbVersion, queryString, crawlerServiceName)
	if err != nil {
		return nil, err
	}
	if len(dbVersion) == 0 {
		return nil, cpersist.ErrPersisterNoResults
	}
	return dbVersion[0].Version, nil
}

// saveVersionToTable saves the version
func (p *PostgresPersister) saveVersionToTable(tableName string, versionNumber *string) error {
	dbVersionStruct := postgres.Version{
		Version:           versionNumber,
		ServiceName:       crawlerServiceName,
		LastUpdatedDateTs: ctime.CurrentEpochSecsInInt64()}
	queryString := p.upsertVersionDataQueryString(tableName, dbVersionStruct)
	_, err := p.db.NamedExec(queryString, dbVersionStruct)
	if err != nil {
		return fmt.Errorf("Error saving version to table: %v", err)
	}
	return nil
}

func (p *PostgresPersister) upsertVersionDataQueryString(tableName string, dbModelStruct interface{}) string {
	onConflict := "version, service_name"
	timestampField := "last_updated_timestamp"
	timestampValue := ":last_updated_timestamp"
	fieldNames, fieldNamesColon := cpostgres.StructFieldsForQuery(dbModelStruct, true, "")
	queryString := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s) ON CONFLICT(%s) DO UPDATE SET %s=%s;",
		tableName, fieldNames, fieldNamesColon, onConflict, timestampField, timestampValue) // nolint: gosec
	return queryString
}

func (p *PostgresPersister) saveEventsToTable(events []*model.Event, tableName string) error {
	queryString := cpostgres.InsertIntoDBQueryString(tableName, postgres.Event{})
	// There is no way to batch insert using sqlx, so doing a loop here
	for _, event := range events {
		err := p.saveEventToTable(queryString, event)
		if err != nil {
			return fmt.Errorf("Error saving %v to db, err: %v", event.Hash(), err)
		}
		log.Infof("saveEventsToTable: saved: %v, %v", event.EventType(), event.Hash()) // Debug, remove later
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
	fields, _ := cpostgres.StructFieldsForQuery(postgres.Event{}, false, "e")
	queryBuf.WriteString(fields)    // nolint: gosec
	queryBuf.WriteString(" FROM ")  // nolint: gosec
	queryBuf.WriteString(tableName) // nolint: gosec
	queryBuf.WriteString(" AS e ")  // nolint: gosec
	if criteria.Hash != "" {
		// If querying by hash, we don't need any other criteria
		queryBuf.WriteString(" WHERE e.hash = :hash") // nolint: gosec
	} else {
		if criteria.FromTs > 0 {
			queryBuf.WriteString(" WHERE e.timestamp >= :fromts") // nolint: gosec
		} else if criteria.BeforeTs > 0 {
			p.addWhereAnd(queryBuf)
			queryBuf.WriteString(" e.timestamp < :beforets") // nolint: gosec
		}
		if criteria.EventType != "" {
			p.addWhereAnd(queryBuf)
			queryBuf.WriteString(" e.event_type = :eventtype") // nolint: gosec
		}
		if criteria.ContractAddress != "" {
			p.addWhereAnd(queryBuf)
			queryBuf.WriteString(" e.contract_address = :contract_address") // nolint: gosec
		}
		//TODO(IS): the following query DOES NOT WORK
		if len(criteria.ExcludeHashes) > 0 {
			sFields, _ := cpostgres.StructFieldsForQuery(postgres.Event{}, false, "")
			p.addWhereAnd(queryBuf)
			notInQuery := fmt.Sprintf(" NOT EXISTS (SELECT %v FROM %v WHERE e.hash IN ('%v'))", sFields,
				tableName, strings.Join(criteria.ExcludeHashes, "','"))
			queryBuf.WriteString(notInQuery) // nolint: gosec
		}
		if criteria.Reverse {
			queryBuf.WriteString(" ORDER BY e.id DESC") // nolint: gosec
		} else {
			queryBuf.WriteString(" ORDER BY e.id") // nolint: gosec
		}
		if criteria.Offset > 0 {
			queryBuf.WriteString(" OFFSET :offset") // nolint: gosec
		}
		if criteria.Count > 0 {
			queryBuf.WriteString(" LIMIT :count") // nolint: gosec
		}
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
