// Package postgres contains interface for postgres DB.
package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joincivil/civil-events-crawler/pkg/generated/eventdef"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	// driver for postgresql
	_ "github.com/lib/pq"
)

// NewPostgresInterface initializes DB
func NewPostgresInterface(host string, port int, user string, password string, dbname string) (*Interface, error) {
	psqlInfo := fmt.Sprintf("host=%s, port=%d, user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	tcrMapping := eventdef.NewCivilTCRContractEventNameToStruct()
	newsroomMapping := eventdef.NewNewsroomContractEventNameToStruct()
	return &Interface{
		db:               db,
		CivilTCREventDef: tcrMapping,
		NewsroomEventDef: newsroomMapping,
	}, nil
}

// Interface is the interface to postgres
type Interface struct {
	db               *sqlx.DB
	CivilTCREventDef *eventdef.CivilTCRContractEventNameToStruct
	NewsroomEventDef *eventdef.NewsroomContractEventNameToStruct
}

// createTables() creates tables in DB if they don't exist
func (p *Interface) createTables() error {
	schema := `
		CREATE TABLE IF NOT EXISTS events(
			id SERIAL PRIMARY KEY,
			hash TEXT UNIQUE,
			event_type TEXT,
			contract_address TEXT,
			contract_name TEXT,
			timestamp INT,
			payload JSONB,
			log_payload JSONB
		);
	`
	_, err := p.db.Exec(schema)
	return err
}

// SaveToEventsTable saves events to DB
func (p *Interface) SaveToEventsTable(events []model.CivilEvent) error {
	// TODO: change to sqlx and change type
	// for _, event := range events {
	// 	// convert event to a type that can be saved by postgres
	// 	eventPostgres := NewCivilEvent(event)
	// 	// you will use sqlx so can just insert the struct directly
	// 	sqlStatementSaveEvents := `
	//            INSERT INTO events (event_type, event_hash, contract_address, timestamp, payload)
	//            VALUES ($1, $2, $3, $4, $5)`
	// 	_, err := p.db.Exec(sqlStatementSaveEvents)
	// }

	return nil
}
