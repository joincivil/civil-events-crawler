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
	// TODO(PN): When postgres persister is ready, replace NullPersister here.
	if config.PersisterType == utils.PersisterTypePostgresql {
		return &persistence.NullPersister{}
	}
	// Default to the NullPersister
	return &persistence.NullPersister{}
}

func setupKillNotify(eventCol *eventcollector.CivilEventCollector) {
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
		header, err := client.HeaderByNumber(context.TODO(), nil)
		if err == nil {
			log.Infof("Latest block number is: %v", header.Number)
		}
	}

	log.Info("Setting up event collector")
	eventCol := eventcollector.NewCivilEventCollector(
		client,
		contractFilterers(config),
		contractWatchers(config),
		listenerMetaDataPersister(config),
		retrieverMetaDataPersister(config),
		eventDataPersister(config),
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
