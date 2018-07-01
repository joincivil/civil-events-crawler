// Package postgres contains interface for postgres DB.
package postgres // import "github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

import (
	"fmt"
	"github.com/jmoiron/sqlx"
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
	return &Interface{
		db: db,
	}, nil
}

// Interface is the interface to postgres
type Interface struct {
	db *sqlx.DB
}

// CreateTables() creates tables in DB if they don't exist
// TODO (IS): move this to postgres/civilevent.go
func (p *Interface) CreateTables() error {
	// TODO: per PN's advice have some logic to determine which models need to be part
	// of this DB to run create for those set of models.
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
func (p *Interface) SaveToEventsTable(events []*model.CivilEvent) error {
	var err error
	for _, event := range events {
		dbEvent, err := NewCivilEvent(event)
		if err != nil {
			return err
		}
		err = p.saveEvent(dbEvent)
		if err != nil {
			return err
		}
	}
	return err
}

func (p *Interface) saveEvent(dbEvent *CivilEvent) error {
	query := `
		INSERT INTO events(event_type, hash, contract_address, contract_name, timestamp, payload, log_payload) 
		VALUES(:event_type, :hash, :contract_address, :contract_name, :timestamp, :payload, log_payload)`
	_, err := p.db.NamedExec(query, dbEvent)
	return err
}
