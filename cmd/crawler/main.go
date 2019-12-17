// Package main contains commands to run
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/profiler"
	"github.com/go-chi/chi"

	elog "github.com/ethereum/go-ethereum/log"
	log "github.com/golang/glog"

	"github.com/joincivil/civil-events-crawler/pkg/eventcollector"
	"github.com/joincivil/civil-events-crawler/pkg/generated/handlerlist"
	"github.com/joincivil/civil-events-crawler/pkg/helpers"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/pubsub"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

	cerrors "github.com/joincivil/go-common/pkg/errors"
)

const (
	pprofPort = ":9090"

	defaultCrawlerServiceName    = "crawler"
	defaultCrawlerServiceVersion = "0.1.0"
)

func initErrorReporter(config *utils.CrawlerConfig) (cerrors.ErrorReporter, error) {
	if config.StackdriverProjectID == "" {
		log.Errorf("Stackdriver project ID required for use, returning null reporter")
		return &cerrors.NullErrorReporter{}, nil
	}

	// Ensure we have some valid defaults set if values not set in config
	// In practical terms, these should always be set by the environment if possible
	if config.StackdriverServiceName == "" {
		config.StackdriverServiceName = defaultCrawlerServiceName
	}
	if config.StackdriverServiceVersion == "" {
		config.StackdriverServiceVersion = defaultCrawlerServiceVersion
	}
	if config.SentryLoggerName == "" {
		config.SentryLoggerName = fmt.Sprintf("%v_logger", defaultCrawlerServiceName)
	}
	if config.SentryRelease == "" {
		config.SentryRelease = defaultCrawlerServiceVersion
	}
	errRepConfig := &cerrors.MetaErrorReporterConfig{
		StackDriverProjectID:      config.StackdriverProjectID,
		StackDriverServiceName:    config.StackdriverServiceName,
		StackDriverServiceVersion: config.StackdriverServiceVersion,
		SentryDSN:                 config.SentryDsn,
		SentryDebug:               false,
		SentryEnv:                 config.SentryEnv,
		SentryLoggerName:          config.SentryLoggerName,
		SentryRelease:             config.SentryRelease,
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

func cleanup(eventCol *eventcollector.EventCollector, killChan chan<- struct{}) {
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

func setupKillNotify(eventCol *eventcollector.EventCollector, killChan chan<- struct{}) {
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
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

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

func startupCloudProfiler(config *utils.CrawlerConfig) {
	if config.CloudProfileProjectID == "" {
		log.Errorf("Cloud profiler project ID required for use, disabling")
		return
	}

	// Put some sane defaults here, but should be set by the config
	if config.CloudProfileServiceName == "" {
		config.CloudProfileServiceName = defaultCrawlerServiceName
	}
	if config.CloudProfileServiceVersion == "" {
		config.CloudProfileServiceVersion = defaultCrawlerServiceVersion
	}
	err := profiler.Start(profiler.Config{
		ProjectID:      config.CloudProfileProjectID,
		Service:        config.CloudProfileServiceName,
		ServiceVersion: config.CloudProfileServiceVersion,
		DebugLogging:   true,
	})
	if err != nil {
		log.Errorf("Error starting up cloud profiler: %v", err)
		return
	}
	log.Infof("Enabling cloud profiler")
}

func startUp(config *utils.CrawlerConfig, errRep cerrors.ErrorReporter) error {
	killChan := make(chan struct{})

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
	persister, err := helpers.EventPersister(config)
	if err != nil {
		return err
	}

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
			PreemptSecs:         config.PreemptSecs,
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

	// Starts up the cloud profiler, if it is enabled in the config
	startupCloudProfiler(config)

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
