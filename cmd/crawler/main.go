// Package main contains commands to run
package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	elog "github.com/ethereum/go-ethereum/log"
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joincivil/civil-events-crawler/pkg/eventcollector"
	"github.com/joincivil/civil-events-crawler/pkg/generated/handlerlist"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence"
	"github.com/joincivil/civil-events-crawler/pkg/pubsub"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

	cconfig "github.com/joincivil/go-common/pkg/config"
)

// const (
// 	websocketPingDelaySecs = 10 // 10 secs
// )

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

func crawlerPubSub(config *utils.CrawlerConfig) *pubsub.CrawlerPubSub {
	if config.PubSubProjectID == "" || config.PubSubTopicName == "" {
		return nil
	}

	pubsub, err := pubsub.NewCrawlerPubSub(config.PubSubProjectID, config.PubSubTopicName)
	if err != nil {
		log.Errorf("Error initializing pubsub, stopping...; err: %v, %v, %v", err, config.PubSubProjectID, config.PubSubTopicName)
		os.Exit(1)
	}
	topicExists, err := pubsub.GooglePubsub.TopicExists(config.PubSubTopicName)
	if err != nil {
		log.Errorf("Error checking for existence of topic: err: %v", err)
		os.Exit(1)
	}
	if !topicExists {
		log.Errorf("Topic: %v does not exist", config.PubSubTopicName)
		os.Exit(1)
	}
	err = pubsub.StartPublishers()
	if err != nil {
		log.Errorf("Error starting publishers, stopping...; err: %v", err)
		os.Exit(1)
	}
	return pubsub
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

func setupHTTPEthClient(config *utils.CrawlerConfig) (*ethclient.Client, error) {
	if isWebsocketURL(config.EthAPIURL) {
		return nil, fmt.Errorf(
			"Fatal: Valid HTTP eth client URL required: configured url: %v",
			config.EthAPIURL,
		)
	}

	client, err := ethclient.Dial(config.EthAPIURL)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func setupWebsocketEthClient(config *utils.CrawlerConfig, killChan <-chan bool) (*ethclient.Client, error) {
	if config.EthWsAPIURL == "" {
		return nil, nil
	}

	if !isWebsocketURL(config.EthWsAPIURL) {
		return nil, nil
	}

	client, err := ethclient.Dial(config.EthWsAPIURL)
	if err != nil {
		return nil, err
	}

	// go eth.WebsocketPing(client, killChan, websocketPingDelaySecs)

	return client, nil
}

func enableGoEtherumLogging() {
	glog := elog.NewGlogHandler(elog.StreamHandler(os.Stderr, elog.TerminalFormat(false)))
	glog.Verbosity(elog.Lvl(elog.LvlDebug)) // nolint: unconvert
	glog.Vmodule("")                        // nolint: errcheck, gosec
	elog.Root().SetHandler(glog)
}

func startUp(config *utils.CrawlerConfig) error {
	killChan := make(chan bool)

	httpClient, err := setupHTTPEthClient(config)
	if err != nil {
		return err
	}

	wsClient, err := setupWebsocketEthClient(config, killChan)
	if err != nil {
		return err
	}

	log.Infof("Starting to filter at block number %v", config.EthStartBlock)
	if log.V(2) {
		header, logErr := httpClient.HeaderByNumber(context.TODO(), nil)
		if logErr == nil {
			log.Infof("Latest block number is: %v", header.Number)
		}
		// If v info level logging, include the ethereum lib logging
		enableGoEtherumLogging()
	}

	eventCol := eventcollector.NewEventCollector(
		&eventcollector.Config{
			Chain:              httpClient,
			HTTPClient:         httpClient,
			WsClient:           wsClient,
			Filterers:          contractFilterers(config),
			Watchers:           contractWatchers(config),
			RetrieverPersister: retrieverMetaDataPersister(config),
			ListenerPersister:  listenerMetaDataPersister(config),
			EventDataPersister: eventDataPersister(config),
			Triggers:           eventTriggers(config),
			StartBlock:         config.EthStartBlock,
			CrawlerPubSub:      crawlerPubSub(config),
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
