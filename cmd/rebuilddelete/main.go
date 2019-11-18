// Package main contains logic to delete old versions of tables
package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/golang/glog"

	"github.com/joincivil/civil-events-crawler/pkg/persistence"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

func main() {
	config := &utils.RebuildConfig{}
	flag.Usage = func() {
		config.OutputUsage()
		os.Exit(0)
	}
	flag.Parse()

	err := config.PopulateFromEnv()
	if err != nil {
		config.OutputUsage()
		log.Errorf("Invalid crawler config: err: %v\n", err)
		os.Exit(2)
	}
	persister, err := persistence.NewPostgresPersister(
		config.PersisterPostgresAddress,
		config.PersisterPostgresPort,
		config.PersisterPostgresUser,
		config.PersisterPostgresPw,
		config.PersisterPostgresDbname,
		nil,
		nil,
		nil,
	)
	if err != nil {
		log.Errorf("Error connecting to Postgresql, stopping...; err: %v", err)
		os.Exit(1)
	}

	versions, err := persister.OldVersions(persistence.CrawlerServiceName)
	if err != nil {
		log.Errorf("Error getting versions, stopping...; err: %v", err)
		os.Exit(1)
	}

	for _, version := range versions {
		eventTableName := fmt.Sprintf("%s_%s", postgres.EventTableBaseName, version)
		log.Infof("Attempting to delete table %v", eventTableName)
		err := persister.DropTable(eventTableName)
		if err != nil {
			log.Errorf("Error deleting %v table, stopping...; err: %v", eventTableName, err)
		}
		log.Infof("Successfully deleted table %v", eventTableName)
		err = persister.UpdateExistenceFalseForVersionTable(postgres.VersionTableName, version, persistence.CrawlerServiceName)
		if err != nil {
			log.Errorf("Error updating exists field for %v table, stopping...; err: %v", eventTableName, err)
		}
	}

	// NOTE(IS): Not deleting versions from version table so we can keep track.
	log.Info("Rebuild completed")
}
