// Package helpers contains various common helper functions.
package helpers

import (
	log "github.com/golang/glog"
	"github.com/jmoiron/sqlx"

	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

	cconfig "github.com/joincivil/go-common/pkg/config"
	"github.com/joincivil/go-common/pkg/strings"
)

// EventPersister is a helper function to return the correct event persister based on
// the given configuration
func EventPersister(config cconfig.PersisterConfig) (model.EventDataPersister, error) {
	if config.PersistType() == cconfig.PersisterTypePostgresql {
		return postgresPersister(config)
	}

	nullPersister := &persistence.NullPersister{}
	return nullPersister, nil
}

// EventPersisterFromSqlx is a helper function to return the correct event persister from
// a given sqlx.DB
func EventPersisterFromSqlx(db *sqlx.DB, config cconfig.PersisterConfig) (model.EventDataPersister, error) {
	persister, err := persistence.NewPostgresPersisterFromSqlx(db)
	if err != nil {
		return nil, err
	}

	err = initTablesAndData(persister, config)
	if err != nil {
		return nil, err
	}

	return persister, nil
}

func postgresPersister(config cconfig.PersisterConfig) (*persistence.PostgresPersister, error) {
	persister, err := persistence.NewPostgresPersister(
		config.Address(),
		config.Port(),
		config.User(),
		config.Password(),
		config.Dbname(),
		config.PoolMaxConns(),
		config.PoolMaxIdleConns(),
		config.PoolConnLifetimeSecs(),
	)
	if err != nil {
		return nil, err
	}

	err = initTablesAndData(persister, config)
	if err != nil {
		return nil, err
	}

	return persister, nil
}

func initTablesAndData(persister *persistence.PostgresPersister, config cconfig.PersisterConfig) error {
	// Pass nil to crawler persistence so it uses the latest version
	err := persister.CreateVersionTable(strings.StrToPtr(config.DataVersion()))
	if err != nil {
		return err
	}
	// Create event table
	err = persister.CreateEventTable()
	if err != nil {
		return err
	}
	// Attempts to create all the necessary indices on the tables
	err = persister.CreateIndices()
	if err != nil {
		return err
	}
	// Populate persistence with latest block data from events table
	err = persister.PopulateBlockDataFromDB(postgres.EventTableBaseName)
	if err != nil {
		log.Errorf("Error populating event last occurrence block data: err: %v", err)
	}

	return nil
}
