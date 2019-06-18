// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	log "github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/Jeffail/tunny"
	"github.com/lib/pq"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/joincivil/civil-events-crawler/pkg/generated/filterer"
	"github.com/joincivil/civil-events-crawler/pkg/generated/watcher"
	"github.com/joincivil/civil-events-crawler/pkg/listener"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/pubsub"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

	cerrors "github.com/joincivil/go-common/pkg/errors"
	"github.com/joincivil/go-common/pkg/eth"
)

const (
	blockHeaderExpirySecs = 60 * 5 // 5 mins

	wsPingDelaySecs = 10 // 10 secs

	defaultPollingIntervalSecs = 60 // 60 secs
)

// Config contains the configuration dependencies for the EventCollector
// NOTES ON WSCLIENT
// - If both WsEthURL / WsClient is set, WsClient will be initially used. If there
//   is a failure, the WS connection will be recreated using the WsEthURL.
// - If WsClient is nil, will use WsEthURL to create the ws connection.
// - If WsEthURL is not set, but WsClient is, it will continue to try and use
//   the given connection even on error (at your own risk).
// - If both are nil, will disable the event listeners.
type Config struct {
	Chain               ethereum.ChainReader
	HTTPClient          bind.ContractBackend
	WsClient            bind.ContractBackend
	WsEthURL            string
	ErrRep              cerrors.ErrorReporter
	Filterers           []model.ContractFilterers
	Watchers            []model.ContractWatchers
	RetrieverPersister  model.RetrieverMetaDataPersister
	ListenerPersister   model.ListenerMetaDataPersister
	EventDataPersister  model.EventDataPersister
	Triggers            []Trigger
	StartBlock          uint64
	CrawlerPubSub       *pubsub.CrawlerPubSub
	PollingEnabled      bool
	PollingIntervalSecs int
}

// NewEventCollector creates a new event collector
func NewEventCollector(config *Config) *EventCollector {
	eventcollector := &EventCollector{
		chain:               config.Chain,
		retryChain:          eth.RetryChainReader{ChainReader: config.Chain},
		httpClient:          config.HTTPClient,
		wsClient:            config.WsClient,
		wsEthURL:            config.WsEthURL,
		errRep:              config.ErrRep,
		filterers:           config.Filterers,
		watchers:            config.Watchers,
		retrieverPersister:  config.RetrieverPersister,
		listenerPersister:   config.ListenerPersister,
		eventDataPersister:  config.EventDataPersister,
		triggers:            config.Triggers,
		startBlock:          config.StartBlock,
		startChan:           make(chan bool),
		quitChan:            make(chan bool),
		errorsChan:          make(chan error),
		crawlerPubSub:       config.CrawlerPubSub,
		pollingEnabled:      config.PollingEnabled,
		pollingIntervalSecs: config.PollingIntervalSecs,
	}
	return eventcollector
}

// EventCollector handles logic for getting historical and live events
type EventCollector struct {
	chain ethereum.ChainReader

	retryChain eth.RetryChainReader

	httpClient bind.ContractBackend

	wsClient bind.ContractBackend

	wsEthURL string

	errRep cerrors.ErrorReporter

	triggers []Trigger

	filterers []model.ContractFilterers

	watchers []model.ContractWatchers

	retrieverPersister model.RetrieverMetaDataPersister

	listenerPersister model.ListenerMetaDataPersister

	eventDataPersister model.EventDataPersister

	listen *listener.EventListener

	retrieve *retriever.EventRetriever

	startBlock uint64

	// startChan is closed when the event collector has started
	startChan chan bool

	// quitChan is closed to stop goroutines and the event collector
	quitChan chan bool

	errorsChan chan error

	mutex sync.Mutex

	headerCache *eth.BlockHeaderCache

	crawlerPubSub *pubsub.CrawlerPubSub

	pollingEnabled bool

	pollingIntervalSecs int
}

// FilterAddedNewsroomContract runs a filterer on the newly watched newsroom contract to ensure we have all events.
func (c *EventCollector) FilterAddedNewsroomContract(newsroomAddr common.Address) ([]*model.Event, error) {
	nwsrmFilterer := filterer.NewNewsroomContractFilterers(newsroomAddr)
	c.updateFiltererStartingBlock(nwsrmFilterer)
	retrieve := retriever.NewEventRetriever(c.httpClient, []model.ContractFilterers{nwsrmFilterer})
	err := retrieve.Retrieve()
	if err != nil {
		return nil, err
	}
	nwsrmEvents := retrieve.PastEvents
	return nwsrmEvents, nil
}

// StartChan returns the channel will send a "event collector started" signal
func (c *EventCollector) StartChan() chan bool {
	return c.startChan
}

// StartCollection contains logic to run retriever and listener.
func (c *EventCollector) StartCollection() error {
	var once sync.Once
	for {
		err := c.runRetriever()
		if err != nil {
			once.Do(func() { c.sendStartSignal() })
			if !c.isAllowedErrRetriever(err) {
				log.Errorf("Error running retriever: err: %v", err)
				return errors.WithMessage(err, "error running retriever in startcol")
			}
			log.Errorf("Error running retriever, recovering: err: %v", err)
			c.errRep.Error(err, nil)
		}

		if c.crawlerPubSub != nil {
			err = c.crawlerPubSub.PublishProcessorTriggerMessage()
			if err != nil {
				log.Errorf("Error publishing trigger message: err: %v", err)
				c.errRep.Error(err, nil)
			}
		}

		if c.pollingEnabled {
			once.Do(func() { c.sendStartSignal() })
			log.Infof("Polling enabled, waiting...")
			select {
			case <-time.After(time.Duration(c.pollingIntSecs()) * time.Second):
				// Loop back to call runRetriever
				continue
			case <-c.quitChan:
				return nil
			}
		}

		if !c.isListenerEnabled() {
			log.Infof("Listener is disabled, not starting")
			once.Do(func() { c.sendStartSignal() })
			return nil
		}

		err = c.startListener()
		if err != nil {
			once.Do(func() { c.sendStartSignal() })
			log.Errorf("Error starting listener: err: %v", err)
			return errors.WithMessage(err, "error starting listener in startcol")
		}

		once.Do(func() { c.sendStartSignal() })

		select {
		// All errors from the watchers and startListener loop get funnelled here.
		case err = <-c.errorsChan:
			log.Errorf("Received error on chan, restarting collector: err: %v", err)
			c.errRep.Error(err, nil)
			// nil this out to kill start signal, not needed at this point
			c.startChan = nil
			// Stop the old listener before starting new one
			if err := c.listen.Stop(true); err != nil {
				log.Errorf("Error stopping listener: err: %v", err)
				c.errRep.Error(err, nil)
			}

		case <-c.quitChan:
			log.Infof("Received quitChan signal, stopping collection")
			return nil
		}
	}
}

// StopCollection is for stopping the listener
func (c *EventCollector) StopCollection(unsubWatchers bool) error {
	var err error
	if c.crawlerPubSub != nil {
		err = c.crawlerPubSub.StopPublishers()
		if err != nil {
			return err
		}
	}
	if c.listen != nil {
		err = c.listen.Stop(unsubWatchers)
	}
	if c.quitChan != nil {
		close(c.quitChan)
	}
	return err
}

// AddWatchers will add watchers to the embedded listener.
func (c *EventCollector) AddWatchers(w model.ContractWatchers) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	return c.listen.AddWatchers(w)
}

// RemoveWatchers will remove given watcher from the embedded listener.
func (c *EventCollector) RemoveWatchers(w model.ContractWatchers) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	return c.listen.RemoveWatchers(w)
}

// CheckRetrievedEventsForNewsroom checks pastEvents for TCR events that may include new newsroom events,
// creates new Newsroom filterers and watchers upon valid events, filters for events, and then returns those events
func (c *EventCollector) CheckRetrievedEventsForNewsroom(pastEvents []*model.Event) ([]*model.Event, error) {
	existingFiltererNewsroomAddr := c.getExistingNewsroomFilterers()
	existingWatcherNewsroomAddr := c.getExistingNewsroomWatchers()
	watchersToAdd := map[common.Address]*watcher.NewsroomContractWatchers{}
	additionalNewsroomFilterers := []model.ContractFilterers{}
	additionalEvents := []*model.Event{}

	for _, event := range pastEvents {
		// NOTE(IS): We should track events from "Application" so we don't miss other events.
		if event.EventType() == "Application" {
			newsroomAddr, ok := event.EventPayload()["ListingAddress"].(common.Address)
			if !ok {
				return additionalEvents, fmt.Errorf("Cannot get newsroomAddr from eventpayload")
			}
			if _, ok := existingFiltererNewsroomAddr[newsroomAddr]; !ok {
				log.Infof("Adding Newsroom filterer for %v", newsroomAddr.Hex())
				newFilterer := filterer.NewNewsroomContractFilterers(newsroomAddr)
				additionalNewsroomFilterers = append(additionalNewsroomFilterers, newFilterer)
				existingFiltererNewsroomAddr[newsroomAddr] = true
			}
			if _, ok := existingWatcherNewsroomAddr[newsroomAddr]; !ok {
				newWatcher := watcher.NewNewsroomContractWatchers(newsroomAddr)
				watchersToAdd[newsroomAddr] = newWatcher
				existingWatcherNewsroomAddr[newsroomAddr] = true
			}
		}
		if event.EventType() == "ListingRemoved" {
			newsroomAddr, ok := event.EventPayload()["ListingAddress"].(common.Address)
			if !ok {
				return additionalEvents, fmt.Errorf("Cannot get newsroomAddr from eventpayload")
			}
			watchersToAdd[newsroomAddr] = nil
		}
	}
	for addr, watcher := range watchersToAdd {
		if watcher != nil {
			log.Infof("Adding Newsroom watcher for %v", addr.Hex())
			c.watchers = append(c.watchers, watcher)
		} else {
			log.Infof("Not adding %v to watchers because it was removed.", addr.Hex())
		}

	}

	if len(additionalNewsroomFilterers) > 0 {
		// NOTE(IS): This overwrites the previous retriever with the new filterers
		// TODO(IS): Better solution for this
		err := c.retrieveEvents(additionalNewsroomFilterers)
		if err != nil {
			return additionalEvents, errors.WithMessage(err, "error retrieving new Newsroom events")
		}
		additionalEvents = append(additionalEvents, c.retrieve.PastEvents...)
	}
	return additionalEvents, nil
}

func (c *EventCollector) runRetriever() error {
	err := c.retrieveEvents(c.filterers)
	if err != nil {
		return errors.WithMessage(err, "error retrieving events")
	}
	pastEvents := c.retrieve.PastEvents

	// Check pastEvents for any new newsrooms to track
	additionalEvents, err := c.CheckRetrievedEventsForNewsroom(pastEvents)
	if err != nil {
		return errors.WithMessage(err, "error checking newsroom events during filterer")
	}

	if len(additionalEvents) > 0 {
		pastEvents = append(pastEvents, additionalEvents...)
	}

	err = c.retrieve.SortEventsByBlock(pastEvents)
	if err != nil {
		return errors.WithMessage(err, "error sorting retrieved events")
	}

	err = c.updateEventTimesFromBlockHeaders(pastEvents)
	if err != nil {
		return errors.WithMessage(err, "error updating dates for events")
	}

	err = c.eventDataPersister.SaveEvents(pastEvents)
	if err != nil {
		return errors.WithMessage(err, "error persisting events")
	}

	err = c.persistRetrieverLastBlockData()
	if err != nil {
		return errors.WithMessage(err, "error persisting last block data")
	}
	return nil
}

func (c *EventCollector) startListener() error {
	defer c.mutex.Unlock()
	c.mutex.Lock()

	wsClient, killChan, err := c.initWsClient()
	if err != nil {
		return errors.WithMessage(err, "startupListener: Unable to create new websocket client")
	}

	c.listen = listener.NewEventListener(wsClient, c.watchers)
	if c.listen == nil {
		return errors.New("startupListener: Listener should not be nil")
	}

	err = c.listen.Start()
	if err != nil {
		return errors.WithMessage(err, "startupListener: Listener should have started with no errors")
	}

	cleanupFn := func() {
		if killChan != nil {
			close(killChan)
		}
		wsClient = nil
		if c.wsEthURL != "" {
			// if we have the ws url, nil out the set wsClient so we
			// potentially create a new one
			c.wsClient = nil
		}
	}

	go func(quit <-chan bool, errs chan<- error) {
		multiplier := 1
		numCPUs := runtime.NumCPU() * multiplier
		pool := tunny.NewFunc(numCPUs, c.handleEvent)
		defer pool.Close()

		for {
			select {
			case event := <-c.listen.EventRecvChan:
				if log.V(2) {
					log.Infof(
						"startupListener: Recv loop event received: %v",
						spew.Sprintf("%#+v", event),
					)
				} else {
					log.Infof(
						"startupListener: Recv loop event received: eventType: %v, hash: %v, ts: %v",
						event.EventType(),
						event.Hash(),
						event.Timestamp(),
					)
				}
				go func(e *model.Event, errs chan<- error) {
					result := pool.Process(
						handleEventInputs{
							event:  e,
							errors: errs,
						},
					)
					// Handler error from event processing
					if result != nil {
						err := result.(error)
						if err != nil {
							log.Errorf(
								"startupListener: pool.Process Error processing, recovering: err: %v: %v",
								err,
								spew.Sdump(event),
							)
							c.errRep.Error(err, nil)
						}
					}
					if log.V(2) {
						log.Infof(
							"startupListener: pool.Process done: %v",
							spew.Sprintf("%#+v", event),
						)
					} else {
						log.Infof(
							"startupListener: pool.Process done: %v, %v, %v",
							event.EventType(),
							event.Hash(),
							event.Timestamp(),
						)
					}
				}(event, errs)

			case err := <-c.listen.Errors:
				// Any errors from the watchers
				log.Infof("startupListener: watcher error chan: %v", err)
				cleanupFn()
				// make sure we send these errors to the main errors chan
				c.errorsChan <- errors.WithMessage(err, "startListener: c.listen.Errors")

				// Flush errors chan since we are only handling the first we receive
				// and killing this consumer. It is possible a number of watchers
				// receive the same error and push it this channel. Without flushing
				// the channel may block indefinitely.
				go func() {
					for {
						select {
						case e := <-c.listen.Errors:
							if e != nil {
								log.Infof("Flushed error: %v", e)
							}
						case <-time.After(time.Second * time.Duration(60*5)):
							return
						}
					}
				}()
				return

			case <-quit:
				log.Infof("startupListener: Quit event recv loop")
				cleanupFn()
				return
			}
		}
	}(c.quitChan, c.errorsChan)

	return nil
}

type handleEventInputs struct {
	event  *model.Event
	errors chan<- error
}

// handleEvent is the func used for the goroutine pool that handles
// incoming events fromt the watchers
func (c *EventCollector) handleEvent(payload interface{}) interface{} {
	inputs := payload.(handleEventInputs)
	eventType := inputs.event.EventType() // Debug, remove later
	hash := inputs.event.Hash()           // Debug, remove later
	txHash := inputs.event.TxHash()       // Debug, remove later
	log.Infof("handleEvent: handling event: %v, tx: %v, hsh: %v", eventType,
		txHash.Hex(), hash) // Debug, remove later
	event := inputs.event

	err := c.updateEventTimeFromBlockHeader(event)
	if err != nil {
		return errors.WithMessage(err, "error updating date for event")
	}
	log.Infof("handleEvent: updated event time from block header: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

	err = c.eventDataPersister.SaveEvents([]*model.Event{event})
	if err != nil {
		return errors.WithMessage(err, "error saving events")
	}
	log.Infof("handleEvent: events saved: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

	if c.crawlerPubSub != nil {
		err = c.crawlerPubSub.PublishProcessorTriggerMessage()
		if err != nil {
			return errors.WithMessagef(err, "error sending message for event %v to pubsub", event.Hash())
		}
	}

	// Update last block in persistence in case of error
	err = c.listenerPersister.UpdateLastBlockData([]*model.Event{event})
	if err != nil {
		return errors.WithMessage(err, "error updating last block")
	}
	log.Infof("handleEvent: updated last block data: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

	// Call event triggers
	err = c.callTriggers(event)
	if err != nil {
		return errors.WithMessage(err, "error calling triggers")
	}
	log.Infof("handleEvent: triggers called: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

	// We need to get past newsroom events for the newsroom contract of a newly added watcher
	if event.EventType() == "Application" {
		newsroomAddr := event.EventPayload()["ListingAddress"].(common.Address)
		newsroomEvents, err := c.FilterAddedNewsroomContract(newsroomAddr)
		if err != nil {
			return errors.WithMessage(err, "error filtering new newsroom contract")
		}
		log.Infof("Found %v newsroom events for address %v after filtering: hsh: %v",
			len(newsroomEvents), newsroomAddr.Hex(), hash) // Debug, remove later
		err = c.eventDataPersister.SaveEvents(newsroomEvents)
		if err != nil {
			return errors.WithMessage(err, "error saving events for application")
		}
		log.Infof("Saved newsroom events at address %v, hsh: %v", newsroomAddr.Hex(), hash) //Debug, remove later

		if c.crawlerPubSub != nil {
			err := c.crawlerPubSub.PublishNewsroomExceptionMessage(newsroomAddr.Hex())
			if err != nil {
				return errors.WithMessagef(err, "error sending message for event %v to pubsub", event.Hash())
			}
		}
	}

	log.Infof("handleEvent: done: %v, tx: %v, hsh: %v", eventType, txHash.Hex(), hash) // Debug, remove later
	return nil
}

func (c *EventCollector) sendStartSignal() {
	if c.startChan != nil {
		close(c.startChan)
	}
}

func (c *EventCollector) isListenerEnabled() bool {
	if (c.wsEthURL == "" && c.wsClient == nil) || c.pollingEnabled {
		return false
	}
	return true
}

func (c *EventCollector) initWsClient() (bind.ContractBackend, chan bool, error) {
	var killChan chan bool
	if c.wsClient != nil {
		return c.wsClient, nil, nil
	}
	killChan = make(chan bool)
	ethclient, err := utils.SetupWebsocketEthClient(c.wsEthURL, killChan, wsPingDelaySecs)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "unable to setup ws client")
	}
	wsClient := bind.ContractBackend(ethclient)
	return wsClient, killChan, nil
}

func (c *EventCollector) pollingIntSecs() int {
	intSecs := c.pollingIntervalSecs
	if intSecs == 0 {
		intSecs = defaultPollingIntervalSecs
	}
	return intSecs
}

// UpdateStartingBlocks updates starting blocks for retriever based on persistence
func (c *EventCollector) updateRetrieverStartingBlocks(filterers []model.ContractFilterers) {
	for _, filter := range filterers {
		c.updateFiltererStartingBlock(filter)
	}
}

func (c *EventCollector) updateFiltererStartingBlock(filter model.ContractFilterers) {
	contractAddress := filter.ContractAddress()
	eventTypes := filter.EventTypes()
	for _, eventType := range eventTypes {
		lastBlock := c.retrieverPersister.LastBlockNumber(eventType, contractAddress)
		// If lastBlock is 0, assume it has never been set, so set to default
		// start block value.
		if lastBlock == 0 {
			lastBlock = c.startBlock
		}
		// NOTE (IS): Starting at lastBlock+1. There could be a scenario where this could miss the rest of events in prev block?
		filter.UpdateStartBlock(eventType, lastBlock+1)
	}
}

func (c *EventCollector) updateEventTimesFromBlockHeaders(events []*model.Event) error {
	for _, event := range events {
		err := c.updateEventTimeFromBlockHeader(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *EventCollector) updateEventTimeFromBlockHeader(event *model.Event) error {
	var header *types.Header
	var err error

	inCache := false
	if c.headerCache == nil {
		c.headerCache = eth.NewBlockHeaderCache(blockHeaderExpirySecs)
	} else {
		header = c.headerCache.HeaderByBlockNumber(event.BlockNumber())
		if header != nil {
			inCache = true
		}
	}
	if !inCache {
		blockNum := big.NewInt(0)
		blockNum.SetUint64(event.BlockNumber())

		log.Infof(
			"updateEventTimeFromBlockHeader: calling headerbynumber: %v, %v",
			event.BlockNumber(),
			blockNum.Int64(),
		) // Debug, remove later

		header, err = c.retryChain.HeaderByNumberWithRetry(event.BlockNumber(), 10, 500)
		if err == nil && header != nil {
			log.Infof(
				"updateEventTimeFromBlockHeader: done calling headerbynumber: %v",
				header.TxHash.Hex(),
			) // Debug, remove later
		}

		c.headerCache.AddHeader(event.BlockNumber(), header)
	}
	if err != nil {
		return errors.Wrap(err, "error update event time")
	}
	event.SetTimestamp(int64(header.Time))
	return nil
}

func (c *EventCollector) retrieveEvents(filterers []model.ContractFilterers) error {
	c.updateRetrieverStartingBlocks(filterers)
	c.retrieve = retriever.NewEventRetriever(c.httpClient, filterers)
	return c.retrieve.Retrieve()
}

func (c *EventCollector) getExistingNewsroomFilterers() map[common.Address]bool {
	existingNewsroomAddr := map[common.Address]bool{}
	for _, existing := range c.filterers {
		specs, _ := model.ContractTypeToSpecs.Get(model.NewsroomContractType)
		if existing.ContractName() == specs.Name() {
			existingNewsroomAddr[existing.ContractAddress()] = true
		}
	}
	return existingNewsroomAddr
}

func (c *EventCollector) getExistingNewsroomWatchers() map[common.Address]bool {
	existingNewsroomAddr := map[common.Address]bool{}
	for _, existing := range c.watchers {
		specs, _ := model.ContractTypeToSpecs.Get(model.NewsroomContractType)
		if existing.ContractName() == specs.Name() {
			existingNewsroomAddr[existing.ContractAddress()] = true
		}
	}
	return existingNewsroomAddr
}

func (c *EventCollector) callTriggers(event *model.Event) error {
	var err error
	for _, trigger := range c.triggers {
		if trigger.ShouldRun(c, event) {
			err = trigger.Run(c, event)
		}
	}
	return err
}

// persistRetrieverLastBlockData saves the last seen events for each filter to
// persistence. Returns the last error seen when updating the block data.
func (c *EventCollector) persistRetrieverLastBlockData() error {
	var err error
	for _, filter := range c.filterers {
		err = c.retrieverPersister.UpdateLastBlockData(filter.LastEvents())
	}
	return err
}

// isAllowedErrRetriever returns if an error should be ignored or not in the
// filterers. This is used in the eventcollector to ensure we only fail on
// particular errors and recover on others.
// ex. if an event hash already exists, we ignore, since that would be "correct" as
// sometimes we may receive the same event and do not want to save it again.
func (c *EventCollector) isAllowedErrRetriever(err error) bool {
	switch causeErr := errors.Cause(err).(type) {
	case *pq.Error:
		log.Infof("*pq error code %v: %v, constraint: %v, msg: %v", causeErr.Code,
			causeErr.Code.Name(), causeErr.Constraint, causeErr.Message)
		return true
	case pq.Error:
		log.Infof("pq error code %v: %v, constraint: %v, msg: %v", causeErr.Code,
			causeErr.Code.Name(), causeErr.Constraint, causeErr.Message)
		return true
	default:
		log.Infof("not allowed error type: %T", causeErr)
	}
	return false
}
