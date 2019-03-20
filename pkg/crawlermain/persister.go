package crawlermain

import (
	"os"

	log "github.com/golang/glog"
	cconfig "github.com/joincivil/go-common/pkg/config"

	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence"
)

// PersisterConfig defines an interface for a config struct to pass into these helpers
type PersisterConfig interface {
	Address() string
	Port() int
	Type() cconfig.PersisterType
	Username() string
	Password() string
	Dbname() string
}

func listenerMetaDataPersister(config PersisterConfig) model.ListenerMetaDataPersister {
	p := persister(config)
	return p.(model.ListenerMetaDataPersister)
}

func retrieverMetaDataPersister(config PersisterConfig) model.RetrieverMetaDataPersister {
	p := persister(config)
	return p.(model.RetrieverMetaDataPersister)
}

func eventDataPersister(config PersisterConfig) model.EventDataPersister {
	p := persister(config)
	return p.(model.EventDataPersister)
}

func persister(config PersisterConfig) interface{} {
	if config.Type() == cconfig.PersisterTypePostgresql {
		return postgresPersister(config)
	}
	// Default to the NullPersister
	return &persistence.NullPersister{}
}

func postgresPersister(config PersisterConfig) *persistence.PostgresPersister {
	persister, err := persistence.NewPostgresPersister(
		config.Address(),
		config.Port(),
		config.Username(),
		config.Password(),
		config.Dbname(),
	)
	if err != nil {
		log.Errorf("Error connecting to Postgresql, stopping...; err: %v", err)
		os.Exit(1)
	}
	// Attempts to create all the necessary tables here
	err = persister.CreateTables()
	if err != nil {
		log.Errorf("Error creating tables, stopping...; err: %v", err)
		os.Exit(1)
	}
	// Attempts to create all the necessary indices on the tables
	err = persister.CreateIndices()
	if err != nil {
		log.Errorf("Error creating indices, stopping...; err: %v", err)
		os.Exit(1)
	}
	// Populate persistence with latest block data from events table
	err = persister.PopulateBlockDataFromDB("event")
	if err != nil {
		log.Errorf("Error populating persistence from Postgresql, stopping...; err: %v", err)
	}
	return persister
}
