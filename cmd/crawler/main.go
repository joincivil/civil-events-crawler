// Package main contains commands to run
package main

import (
	"context"
	"flag"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joincivil/civil-events-crawler/pkg/eventcollector"
	"github.com/joincivil/civil-events-crawler/pkg/generated/handlerlist"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

	cconfig "github.com/joincivil/go-common/pkg/config"
)

const (
	websocketPingDelaySecs = 60 * 5 // 5 mins
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
	if config.PersisterType == cconfig.PersisterTypePostgresql {
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

func cleanup(eventCol *eventcollector.EventCollector, killChan chan<- bool) {
	log.Info("Stopping crawler...")
	err := eventCol.StopCollection()
	if err != nil {
		log.Errorf("Error stopping collection: err: %v", err)
	}
	close(killChan)
	log.Info("Crawler stopped")
}

func setupKillNotify(eventCol *eventcollector.EventCollector, killChan chan<- bool) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup(eventCol, killChan)
		os.Exit(1)
	}()
}

// websocketPing periodically makes a call over the websocket conn
// to the Eth node to ensure the connection stays alive.
// Since there is no built in facility to do pings with the go-eth lib,
// need to do this ourselves by making a eth_getBlockByNumber RPC call.
// NOTE(PN): Need to ensure the client passed in is a websocket client.
// XXX(PN): Need to replace this someday with something better.
func websocketPing(client *ethclient.Client, killChan <-chan bool) {
Loop:
	for {
		select {
		case <-time.After(websocketPingDelaySecs * time.Second):
			_, err := client.HeaderByNumber(context.TODO(), nil)
			// header, err := client.HeaderByNumber(context.TODO(), nil)
			if err != nil {
				log.Errorf("Ping header by number failed: err: %v", err)
			}
			// log.Infof("Ping success: block number: %v", header.Number)

		case <-killChan:
			log.Infof("Closing websocket ping")
			break Loop
		}
	}
}

func isWebsocketURL(rawurl string) bool {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Infof("Unable to parse URL: err: %v", err)
		return false
	}
	if u.Scheme == "ws" || u.Scheme == "wss" {
		return true
	}
	return false
}

func setupWebsocketPing(config *utils.CrawlerConfig, client *ethclient.Client,
	killChan <-chan bool) error {
	// If websocket connection, setup "ping"
	// otherwise, ignore this
	if isWebsocketURL(config.EthAPIURL) {
		go websocketPing(client, killChan)
	}
	return nil
}

func setupEthClient(config *utils.CrawlerConfig, killChan <-chan bool) (*ethclient.Client, error) {
	client, err := ethclient.Dial(config.EthAPIURL)
	if err != nil {
		return nil, err
	}
	err = setupWebsocketPing(config, client, killChan)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func startUp(config *utils.CrawlerConfig) error {
	killChan := make(chan bool)

	client, err := setupEthClient(config, killChan)
	if err != nil {
		return err
	}

	log.Infof("Starting to filter at block number %v", config.EthStartBlock)
	if log.V(2) {
		header, logErr := client.HeaderByNumber(context.TODO(), nil)
		if logErr == nil {
			log.Infof("Latest block number is: %v", header.Number)
		}
	}

	eventCol := eventcollector.NewEventCollector(
		&eventcollector.Config{
			Chain:              client,
			Client:             client,
			Filterers:          contractFilterers(config),
			Watchers:           contractWatchers(config),
			RetrieverPersister: retrieverMetaDataPersister(config),
			ListenerPersister:  listenerMetaDataPersister(config),
			EventDataPersister: eventDataPersister(config),
			Triggers:           eventTriggers(config),
			StartBlock:         config.EthStartBlock,
			DisableListener:    !isWebsocketURL(config.EthAPIURL),
		},
	)

	// Setup shutdown/cleanup hooks
	setupKillNotify(eventCol, killChan)
	defer func() {
		cleanup(eventCol, killChan)
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
