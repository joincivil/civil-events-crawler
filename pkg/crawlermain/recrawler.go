package crawlermain

import (
	"context"

	log "github.com/golang/glog"

	"github.com/joincivil/civil-events-crawler/pkg/eventcollector"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

	"github.com/joincivil/go-common/pkg/eth"
)

// StartUpRecrawler runs the recrawler based on the config
// The recrawler pulls events from Civil epoch, checks the DB for existing events,
// and adds them if they are missing. Can be configured to update existing events
// if their log meta data has changed.
func StartUpRecrawler(config *utils.RecrawlerConfig) error {
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

	// Update start block from the configured start
	filterers := contractFilterers(config.ContractAddressObjs)
	for _, filterer := range filterers {
		for _, etype := range filterer.EventTypes() {
			filterer.UpdateStartBlock(etype, config.EthStartBlock+1)
		}
	}

	epersister := eventDataPersister(config)

	eventRetriever := retriever.NewEventRetriever(httpClient, filterers)
	err = eventRetriever.Retrieve()
	if err != nil {
		log.Errorf("Error retrieving events: err: %v", err)
	}

	evs := eventRetriever.PastEvents

	eventsToAppend := []*model.Event{}
	var headerCache *eth.BlockHeaderCache
	retryChain := &eth.RetryChainReader{ChainReader: httpClient}
	eventFound := false

	for _, ev := range evs {
		// Retrieve all events that share the same txHash
		criteria := &model.RetrieveEventsCriteria{
			TxHash: ev.TxHash().Hex(),
		}
		log.Infof("looking for txhash: %v", ev.TxHash().Hex())
		events, err := epersister.RetrieveEvents(criteria)
		if err != nil {
			log.Errorf("error retrieving events: %v", err)
			continue
		}

		payload := ev.EventPayload()

		if len(events) > 0 {
			for _, event := range events {
				// If not same event type, skip since there could be multiple events in the same tx
				if event.EventType() != ev.EventType() {
					log.Infof("OTHER: %v, %v, %v, %v", event.Hash(), event.EventType(),
						event.Timestamp(), event.LogPayloadToString())
					continue
				}

				// Event types are the same, so check to see if the values are matching in the event payload
				valsMatch := true
				for key, val := range payload {
					epayload := event.EventPayload()
					eval, ok := epayload[key]
					// If not found in payload, skip
					if !ok {
						log.Infof("found matching event type, but no field found")
						valsMatch = false
						break
					}

					log.Infof("%v == %v, %T == %T", val, eval, val, eval)
					eq, eqerr := utils.IsInterfaceEqual(val, eval)
					if eqerr != nil {
						log.Errorf("error comparing interfaces: err: %v", eqerr)
						continue
					}
					// If vals not equal, skip
					if !eq {
						log.Infof("found matching event type, but payload vals don't match, %v != %v", val, eval)
						valsMatch = false
						break
					}
				}

				if !valsMatch {
					log.Info("vals don't match, so skipping")
					continue
				}

				// If the event payload match, then does the log payload data match?
				// If so, that is the one we are looking for.
				if ev.TxIndex() == event.TxIndex() ||
					ev.LogIndex() == event.LogIndex() ||
					ev.BlockNumber() == event.BlockNumber() ||
					ev.BlockHash() == event.BlockHash() {
					log.Infof("FOUND: %v, %v, %v, %v", event.Hash(), event.EventType(),
						event.Timestamp(), event.LogPayloadToString())
					eventFound = true
					continue
				}

				// If the log data is not the same, then we should update?
				if ev.TxIndex() != event.TxIndex() ||
					ev.LogIndex() != event.LogIndex() ||
					ev.BlockNumber() != event.BlockNumber() ||
					ev.BlockHash() != event.BlockHash() {
					log.Infof("NOTMATCH: %v, %v, %v, %v", event.Hash(), event.EventType(), event.Timestamp(), event.LogPayloadToString())
				}
			}
			log.Infof("done checking db events: did we find it = %v\n\n", eventFound)
		}

		if eventFound {
			continue
		}

		err = eventcollector.UpdateEventTimeFromBlockHeader(ev, retryChain, headerCache)
		if err != nil {
			log.Errorf("error updating event time from block header: %v", err)
			continue
		}

		eventsToAppend = append(eventsToAppend, ev)
	}

	for _, event := range eventsToAppend {
		log.Infof("Event to add: %v, %v, %v, %v", event.EventType(), event.Timestamp(), event.LogPayloadToString(), event.Timestamp())
	}

	if config.WetRun && len(eventsToAppend) > 0 {
		log.Infof("Wet run, would be saving to the db")
		// err = epersister.SaveEvents(eventsToAppend)
		// if err != nil {
		// 	return fmt.Errorf("Error saving events: err: %v", err)
		// }
	} else {
		log.Infof("WET_RUN=false or no events, did not save")
	}

	return nil
}
