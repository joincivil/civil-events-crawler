// Package main contains commands to run
package main

import (
	"context"
	"flag"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"

	elog "github.com/ethereum/go-ethereum/log"
	log "github.com/golang/glog"

	"github.com/joincivil/civil-events-crawler/pkg/eventcollector"
	"github.com/joincivil/civil-events-crawler/pkg/generated/handlerlist"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"
	"github.com/joincivil/civil-events-crawler/pkg/pubsub"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

	cconfig "github.com/joincivil/go-common/pkg/config"
	cerrors "github.com/joincivil/go-common/pkg/errors"
)

const (
	pprofPort = ":9090"
)

func initErrorReporter(config *utils.CrawlerConfig) (cerrors.ErrorReporter, error) {
	errRepConfig := &cerrors.MetaErrorReporterConfig{
		StackDriverProjectID:      "civil-media",
		StackDriverServiceName:    "crawler",
		StackDriverServiceVersion: "1.0",
		SentryDSN:                 config.SentryDsn,
		SentryDebug:               false,
		SentryEnv:                 config.SentryEnv,
		SentryLoggerName:          "crawler_logger",
		SentryRelease:             "1.0",
		SentrySampleRate:          1.0,
	}
	reporter, err := cerrors.NewMetaErrorReporter(errRepConfig)
	if err != nil {
		log.Errorf("Error creating meta reporter: %v", err)
		return nil, err
	}
	if reporter == nil {
		log.Infof("Enabling null error reporter")
		return &cerrors.NullErrorReporter{}, nil
	}
	return reporter, nil
}

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
	// Create version_data table
	err = persister.CreateVersionTable(&config.VersionNumber)
	if err != nil {
		log.Errorf("Error creating tables, stopping...; err: %v", err)
		os.Exit(1)
	}
	// Create event table
	err = persister.CreateEventTable()
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
	err = persister.PopulateBlockDataFromDB(postgres.EventTableBaseName)
	if err != nil {
		log.Errorf("Error populating event last occurrence block data: err: %v", err)
	}
	return persister
}

func cleanup(eventCol *eventcollector.EventCollector, killChan chan<- bool) {
	log.Info("Stopping crawler...")
	err := eventCol.StopCollection(false)
	if err != nil {
		log.Errorf("Error stopping collection: err: %v", err)
	}
	if killChan != nil {
		close(killChan)
	}
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

func enableGoEtherumLogging() {
	glog := elog.NewGlogHandler(elog.StreamHandler(os.Stderr, elog.TerminalFormat(false)))
	glog.Verbosity(elog.Lvl(elog.LvlDebug)) // nolint: unconvert
	glog.Vmodule("")                        // nolint: errcheck, gosec
	elog.Root().SetHandler(glog)
}

// For profiling running services
func startupPprofServices() {
	r := chi.NewRouter()

	// Register pprof handlers
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/debug/pprof/block", pprof.Handler("block"))

	err := http.ListenAndServe(pprofPort, r)
	if err != nil {
		log.Errorf("Error starting up pprof endpoints: %v", err)
	}
}

func startUp(config *utils.CrawlerConfig, errRep cerrors.ErrorReporter) error {
	killChan := make(chan bool)

	httpClient, err := utils.SetupHTTPEthClient(config.EthAPIURL)
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

	// Create a single persister to be passed up into the different persister
	// interfaces
	persister := persister(config)

	eventCol := eventcollector.NewEventCollector(
		&eventcollector.Config{
			Chain:               httpClient,
			HTTPClient:          httpClient,
			WsEthURL:            config.EthWsAPIURL,
			ErrRep:              errRep,
			Filterers:           contractFilterers(config),
			Watchers:            contractWatchers(config),
			RetrieverPersister:  persister.(model.RetrieverMetaDataPersister),
			ListenerPersister:   persister.(model.ListenerMetaDataPersister),
			EventDataPersister:  persister.(model.EventDataPersister),
			Triggers:            eventTriggers(config),
			StartBlock:          config.EthStartBlock,
			CrawlerPubSub:       crawlerPubSub(config),
			PollingEnabled:      config.PollingEnabled,
			PollingIntervalSecs: config.PollingIntervalSecs,
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

	errRep, err := initErrorReporter(config)
	if err != nil {
		log.Errorf("Error init error reporting: err: %+v\n", err)
		os.Exit(2)
	}

	if config.PprofEnable {
		go startupPprofServices()
		log.Infof("Enabling pprof endpoints at localhost%v", pprofPort)
	}

	err = startUp(config, errRep)
	if err != nil {
		log.Errorf("Crawler error: err: %+v\n", err)
		errRep.Error(err, nil)
		// XXX(PN): Ensure the error gets sent before we die
		time.Sleep(3 * time.Second)
		os.Exit(2)
	}
	log.Info("Crawler stopped")
}
