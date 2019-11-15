// Package helpers contains various common helper functions.
package helpers

import (
	log "github.com/golang/glog"

	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"
	"google.golang.org/appengine/log"

	cconfig "github.com/joincivil/go-common/pkg/config"
)

// EventPersister is a helper function to return the correct event persister based on
// the given configuration
func EventPersister(config cconfig.PersisterConfig) (model.EventDataPersister, error) {
	if config.PersistType() == cconfig.PersisterTypePostgresql {
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

		// Pass nil to crawler persistence so it uses the latest version
		err = persister.CreateVersionTable(nil)
		if err != nil {
			return nil, err
		}
		// Create event table
		err = persister.CreateEventTable()
		if err != nil {
			return nil, err
		}
		// Attempts to create all the necessary indices on the tables
		err = persister.CreateIndices()
		if err != nil {
			return nil, err
		}

		// Populate persistence with latest block data from events table
		err = persister.PopulateBlockDataFromDB(postgres.EventTableBaseName)
		if err != nil {
			log.Errorf("Error populating event last occurrence block data: err: %v", err)
		}

		return persister, nil
	}

	nullPersister := &persistence.NullPersister{}
	return nullPersister, nil
}
