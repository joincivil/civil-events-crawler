package crawlermain

import (
	"context"

	log "github.com/golang/glog"

	"github.com/joincivil/civil-events-crawler/pkg/eventcollector"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

	cerrors "github.com/joincivil/go-common/pkg/errors"
)

// StartUpCrawler fires up the main crawler process
func StartUpCrawler(config *utils.CrawlerConfig, errRep cerrors.ErrorReporter) error {
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

	eventCol := eventcollector.NewEventCollector(
		&eventcollector.Config{
			Chain:               httpClient,
			HTTPClient:          httpClient,
			WsEthURL:            config.EthWsAPIURL,
			ErrRep:              errRep,
			Filterers:           contractFilterers(config.ContractAddressObjs),
			Watchers:            contractWatchers(config.ContractAddressObjs),
			RetrieverPersister:  retrieverMetaDataPersister(config),
			ListenerPersister:   listenerMetaDataPersister(config),
			EventDataPersister:  eventDataPersister(config),
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
