// Package persistence contains components to interact with the DB
package persistence // import "github.com/joincivil/civil-events-crawler/pkg/persistence"

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	log "github.com/golang/glog"
	"github.com/pkg/errors"

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
	//CrawlerServiceName is the name of the crawler service
	CrawlerServiceName = "crawler"
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
		return pgPersister, errors.Wrap(err, "error connecting to sqlx")
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

// OldVersions returns all versions except for the most recent one for this service
func (p *PostgresPersister) OldVersions(serviceName string) ([]string, error) {
	return p.oldVersionsFromTable(serviceName, postgres.VersionTableName)
}

// GetTableName formats tabletype with version of this persister to return the table name
func (p *PostgresPersister) GetTableName(tableType string) string {
	if p.version == nil || *p.version == "" {
		return tableType
	}
	return fmt.Sprintf("%s_%s", tableType, *p.version)
}

// DropTable drops the table with the specified tableName
func (p *PostgresPersister) DropTable(tableName string) error {
	_, err := p.db.Query(fmt.Sprintf("DROP TABLE IF EXISTS %v;", tableName)) // nolint: gosec
	return err
}

// UpdateExistenceFalseForVersionTable updates the tableName's exists field to false in the version table
func (p *PostgresPersister) UpdateExistenceFalseForVersionTable(tableName string, versionNumber string,
	serviceName string) error {
	if versionNumber == "" {
		return nil
	}
	dbVersionStruct := postgres.Version{
		Version:     &versionNumber,
		ServiceName: serviceName,
		Exists:      false}
	onConflict := fmt.Sprintf("%s, %s", postgres.VersionFieldName, postgres.ServiceFieldName)
	updatedFields := []string{postgres.ExistsFieldName}
	queryString := p.upsertVersionDataQueryString(tableName, dbVersionStruct, onConflict,
		updatedFields)
	_, err := p.db.NamedExec(queryString, dbVersionStruct)
	if err != nil {
		return fmt.Errorf("Error saving version to table: %v", err)
	}
	return nil
}

// SaveVersion saves the version for this persistence
func (p *PostgresPersister) SaveVersion(versionNumber *string) error {
	if versionNumber == nil || *versionNumber == "" {
		return nil
	}
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
			return errors.WithMessage(err, "error updating last block data")
		}
	}
	return nil
}

// PopulateBlockDataFromDB will determine the block data for latest occurrence of each event type
// and store it into an internal map. One of the purposes of this map determines at which block
// to start looking for each type of event.
func (p *PostgresPersister) PopulateBlockDataFromDB(tableType string) error {
	tableName := p.GetTableName(tableType)
	events, err := p.getLatestEvents(tableName)
	if err != nil {
		return err
	}
	var lasterr error
	for _, event := range events {
		modelEvent, err := event.DBToEventData()
		if err != nil {
			log.Errorf("Error converting db event to event: %v", err)
			lasterr = errors.WithMessage(err, "error converting db event to event")
			continue
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
	return lasterr
}

// CreateVersionTable creates the version table and sets the version with new or existing version
func (p *PostgresPersister) CreateVersionTable(version *string) error {
	versionTableQuery := postgres.CreateVersionTableQuery(postgres.VersionTableName)
	_, err := p.db.Exec(versionTableQuery)
	if err != nil {
		return err
	}

	// Check to see if there is an existing latest version
	currentVersion, err := p.PersisterVersion()
	if err != nil && err != cpersist.ErrPersisterNoResults {
		return err
	}

	// If no version found anywhere, don't use versioned tables
	if (currentVersion == nil || *currentVersion == "") && (version == nil || *version == "") {
		log.Infof("No version found, not using versioned tables")
		return nil
	}

	// If the incoming version is the same as the currentVersion, don't do anything
	if currentVersion != nil && version != nil && *currentVersion == *version {
		log.Infof("Using data version: %v", *version)
		return nil
	}

	// If version does not exist, but currentVersion does, use currentVersion
	if currentVersion != nil && (version == nil || *version == "") {
		// NOTE(IS): Use existing version, but update timestamp
		version = currentVersion
	}

	log.Infof("Updated data version: %v", *version)
	p.version = version
	return p.SaveVersion(version)
}

// CreateEventTable creates event table
func (p *PostgresPersister) CreateEventTable() error {
	eventTableName := p.GetTableName(postgres.EventTableBaseName)
	eventTableQuery := postgres.CreateEventTableQuery(eventTableName)
	_, err := p.db.Exec(eventTableQuery)
	return errors.Wrap(err, "error creating tables")
}

// CreateIndices creates the indices for DB if they don't exist
func (p *PostgresPersister) CreateIndices() error {
	eventTableName := p.GetTableName(postgres.EventTableBaseName)
	indexQuery := postgres.CreateEventTableIndices(eventTableName)
	_, err := p.db.Exec(indexQuery)
	return errors.Wrap(err, "error creating indexes")
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

func (p *PostgresPersister) oldVersionsFromTable(serviceName string, tableName string) ([]string, error) {
	dbVersions := []postgres.Version{}
	// nolint
	queryStringLargest := fmt.Sprintf(
		`SELECT MAX(last_updated_timestamp) FROM %s WHERE service_name='%s'`,
		tableName,
		serviceName,
	)
	// nolint
	queryString := fmt.Sprintf(
		`SELECT * FROM %s WHERE service_name=$1 AND exists=true AND last_updated_timestamp !=(%s)`, // nolint: gosec
		tableName,
		queryStringLargest,
	)
	err := p.db.Select(&dbVersions, queryString, serviceName)
	if err != nil {
		return []string{}, err
	}
	versions := []string{}
	for _, version := range dbVersions {
		versions = append(versions, *version.Version)
	}

	return versions, nil
}

func (p *PostgresPersister) retrieveVersionFromTable(tableName string) (*string, error) {
	dbVersion := []postgres.Version{}
	queryString := fmt.Sprintf(`SELECT * FROM %s WHERE service_name=$1 ORDER BY last_updated_timestamp DESC LIMIT 1;`, tableName) // nolint: gosec
	err := p.db.Select(&dbVersion, queryString, CrawlerServiceName)
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
		ServiceName:       CrawlerServiceName,
		LastUpdatedDateTs: ctime.CurrentEpochSecsInInt64(),
		Exists:            true}
	onConflict := fmt.Sprintf("%s, %s", postgres.VersionFieldName, postgres.ServiceFieldName)
	updateFields := []string{postgres.LastUpdatedTsFieldName, postgres.ExistsFieldName}
	queryString := p.upsertVersionDataQueryString(tableName, dbVersionStruct, onConflict,
		updateFields)
	_, err := p.db.NamedExec(queryString, dbVersionStruct)
	if err != nil {
		return fmt.Errorf("Error saving version to table: %v", err)
	}
	return nil
}

func (p *PostgresPersister) upsertVersionDataQueryString(tableName string, dbModelStruct interface{},
	onConflict string, updatedFields []string) string {
	var queryString strings.Builder
	fieldNames, fieldNamesColon := cpostgres.StructFieldsForQuery(dbModelStruct, true, "")
	// nolint
	queryString.WriteString(fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s) ON CONFLICT(%s) DO UPDATE SET ",
		tableName, fieldNames, fieldNamesColon, onConflict))
	for idx, field := range updatedFields {
		// nolint
		queryString.WriteString(fmt.Sprintf("%s=:%s", field, field))
		if idx+1 < len(updatedFields) {
			queryString.WriteString(", ") // nolint: gosec
		}
	}
	return queryString.String()
}

func (p *PostgresPersister) saveEventsToTable(events []*model.Event, tableName string) error {
	queryString := cpostgres.InsertIntoDBQueryString(tableName, postgres.Event{})
	var lasterr error
	// There is no way to batch insert using sqlx, so doing a loop here
	for _, event := range events {
		err := p.saveEventToTable(queryString, event)
		if err != nil {
			// We want to ensure we save as much as possible before returning error
			// so just return the last error we saw
			log.Errorf("saveEventsToTable: err saving %v to db: err: %v", event.Hash(), err)
			lasterr = errors.WithMessagef(err, "error saving %v to db", event.Hash())
			continue
		}
		log.Infof("saveEventsToTable: saved: %v, %v", event.EventType(), event.Hash()) // Debug, remove later
	}
	return lasterr
}

func (p *PostgresPersister) retrieveEventsFromTable(tableName string, criteria *model.RetrieveEventsCriteria) ([]*model.Event, error) {
	queryString := p.retrieveEventsQuery(tableName, criteria)
	dbEvents := []postgres.Event{}
	nstmt, err := p.db.PrepareNamed(queryString)
	if err != nil {
		return nil, errors.Wrap(err, "error preparing query with sqlx")
	}
	err = nstmt.Select(&dbEvents, criteria)
	if err != nil {
		return nil, errors.Wrap(err, "error retrieving events from table")
	}
	events := make([]*model.Event, len(dbEvents))
	var lasterr error
	for index, dbEvent := range dbEvents {
		modelEvent, err := dbEvent.DBToEventData()
		if err != nil {
			log.Errorf("Error converting db to event data: err: %v", err)
			lasterr = errors.WithMessage(err, "error converting db to event data")
			continue
		}
		events[index] = modelEvent
	}
	return events, lasterr
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
			// nolint
			notInQuery := fmt.Sprintf(" NOT EXISTS (SELECT %v FROM %v WHERE e.hash IN ('%v'))", sFields,
				tableName, strings.Join(criteria.ExcludeHashes, "','"))
			queryBuf.WriteString(notInQuery) // nolint: gosec
		}

		queryBuf.WriteString(" ORDER BY") // nolint: gosec
		// Using standard SQL here bc of issues using : with prepared statements
		if criteria.Reverse {
			queryBuf.WriteString(" CAST(e.log_payload->>'BlockNumber' as integer) DESC,") // nolint: gosec
			queryBuf.WriteString(" CAST(e.log_payload->>'Index' as integer) DESC")        // nolint: gosec
		} else {
			queryBuf.WriteString(" CAST(e.log_payload->>'BlockNumber' as integer),") // nolint: gosec
			queryBuf.WriteString(" CAST(e.log_payload->>'Index' as integer)")        // nolint: gosec
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
		return errors.Wrap(err, "error saving event to table")
	}
	log.Infof("saveEventToTable: done") // Debug, remove later
	return nil
}

func (p *PostgresPersister) getLatestEvents(tableName string) ([]postgres.Event, error) {
	queryString := p.retrieveLatestEventsQueryString(tableName)
	events := []postgres.Event{}
	err := p.db.Select(&events, queryString)
	if err != nil {
		return nil, errors.Wrap(err, "error getting latest events")
	}
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
