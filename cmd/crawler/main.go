// Package main contains commands to run
package main

import (
	"context"
	"flag"
	log "github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joincivil/civil-events-crawler/pkg/eventcollector"
	"github.com/joincivil/civil-events-crawler/pkg/generated/handlerlist"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

func contractFilterers(config *utils.CrawlerConfig) []model.ContractFilterers {
	return handlerlist.ContractFilterers(config.ContractAddressObjs)
}

func contractWatchers(config *utils.CrawlerConfig) []model.ContractWatchers {
	return handlerlist.ContractWatchers(config.ContractAddressObjs)
}

func eventTriggers(config *utils.CrawlerConfig) []eventcollector.Trigger {
	return []eventcollector.Trigger{
		&eventcollector.AddNewsroomWatchersTrigger{},
		&eventcollector.RemoveNewsroomWatchersTrigger{},
	}
}

func listenerMetaDataPersister(config *utils.CrawlerConfig) model.ListenerMetaDataPersister {
	p := persister(config)
	return p.(model.ListenerMetaDataPersister)
}

func retrieverMetaDataPersister(config *utils.CrawlerConfig) model.RetrieverMetaDataPersister {
	p := persister(config)
	return p.(model.RetrieverMetaDataPersister)
}

func eventDataPersister(config *utils.CrawlerConfig) model.EventDataPersister {
	p := persister(config)
	return p.(model.EventDataPersister)
}

func persister(config *utils.CrawlerConfig) interface{} {
	if config.PersisterType == utils.PersisterTypePostgresql {
		return postgresPersister(config)
	}
	// Default to the NullPersister
	return &persistence.NullPersister{}
}

func postgresPersister(config *utils.CrawlerConfig) *persistence.PostgresPersister {
	persister, err := persistence.NewPostgresPersister(
		config.PersisterPostgresAddress,
		config.PersisterPostgresPort,
		config.PersisterPostgresUser,
		config.PersisterPostgresPw,
		config.PersisterPostgresDbname,
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
	// Populate persistence with latest block data from events table
	err = persister.PopulateBlockDataFromDB("event")
	if err != nil {
		log.Errorf("Error populating persistence from Postgresql, stopping...; err: %v", err)
	}
	return persister
}

func setupKillNotify(eventCol *eventcollector.EventCollector) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		err := eventCol.StopCollection()
		if err != nil {
			log.Errorf("Error stopping collection: err: %v", err)
		}
		log.Info("Crawler stopped")
		os.Exit(1)
	}()
}

func startUp(config *utils.CrawlerConfig) error {
	log.Info("Setting up ethclient")
	client, err := ethclient.Dial(config.EthAPIURL)
	if err != nil {
		return err
	}

	if log.V(2) {
		header, logErr := client.HeaderByNumber(context.TODO(), nil)
		if logErr == nil {
			log.Infof("Latest block number is: %v", header.Number)
		}
		log.Infof("Starting to filter at block number %v", config.EthStartBlock)
	}

	log.Info("Setting up event collector")
	eventCol := eventcollector.NewEventCollector(
		client,
		client,
		contractFilterers(config),
		contractWatchers(config),
		listenerMetaDataPersister(config),
		retrieverMetaDataPersister(config),
		eventDataPersister(config),
		eventTriggers(config),
		config.EthStartBlock,
	)

	setupKillNotify(eventCol)
	defer func() {
		err := eventCol.StopCollection()
		if err != nil {
			log.Errorf("Error stopping collection: err: %v", err)
		}
	}()

	log.Info("Crawler starting...")
	// Will block here while handling collection
	return eventCol.StartCollection()
}

func main() {
	config := &utils.CrawlerConfig{}
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

	err = startUp(config)
	if err != nil {
		log.Errorf("Crawler error: err: %v\n", err)
		os.Exit(2)
	}
	log.Info("Crawler stopped")
}
