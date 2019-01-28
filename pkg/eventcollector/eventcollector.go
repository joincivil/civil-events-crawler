// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	"errors"
	"fmt"
	"math/big"
	"runtime"
	"sync"

	log "github.com/golang/glog"

	"github.com/Jeffail/tunny"

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

	"github.com/joincivil/go-common/pkg/eth"
)

const (
	blockHeaderExpirySecs = 60 * 5 // 5 mins
)

// Config contains the configuration dependencies for the EventCollector
type Config struct {
	Chain              ethereum.ChainReader
	RetryChain         eth.RetryChainReader
	Client             bind.ContractBackend
	Filterers          []model.ContractFilterers
	Watchers           []model.ContractWatchers
	RetrieverPersister model.RetrieverMetaDataPersister
	ListenerPersister  model.ListenerMetaDataPersister
	EventDataPersister model.EventDataPersister
	Triggers           []Trigger
	StartBlock         uint64
	DisableListener    bool
	CrawlerPubSub      *pubsub.CrawlerPubSub
}

// NewEventCollector creates a new event collector
func NewEventCollector(config *Config) *EventCollector {
	eventcollector := &EventCollector{
		chain:              config.Chain,
		retryChain:         eth.RetryChainReader{ChainReader: config.Chain},
		client:             config.Client,
		filterers:          config.Filterers,
		watchers:           config.Watchers,
		retrieverPersister: config.RetrieverPersister,
		listenerPersister:  config.ListenerPersister,
		eventDataPersister: config.EventDataPersister,
		triggers:           config.Triggers,
		startBlock:         config.StartBlock,
		startChan:          make(chan bool),
		quitChan:           make(chan bool),
		disableListener:    config.DisableListener,
		crawlerPubSub:      config.CrawlerPubSub,
	}
	return eventcollector
}

// EventCollector handles logic for getting historical and live events
type EventCollector struct {
	chain ethereum.ChainReader

	retryChain eth.RetryChainReader

	client bind.ContractBackend

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

	disableListener bool

	crawlerPubSub *pubsub.CrawlerPubSub
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
		return fmt.Errorf("Error updating date for event: err: %v", err)
	}
	log.Infof("handleEvent: updated event time from block header: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

	err = c.eventDataPersister.SaveEvents([]*model.Event{event})
	if err != nil {
		return fmt.Errorf("Error saving events: err: %v", err)
	}
	log.Infof("handleEvent: events saved: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

	if c.crawlerPubSub != nil {
		err = c.crawlerPubSub.PublishTriggerMessage()
		if err != nil {
			return fmt.Errorf("Error sending message for event %v to pubsub: %v", event.Hash(), err)
		}
	}

	// Update last block in persistence in case of error
	err = c.listenerPersister.UpdateLastBlockData([]*model.Event{event})
	if err != nil {
		return fmt.Errorf("Error updating last block: err: %v", err)
	}
	log.Infof("handleEvent: updated last block data: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

	// Call event triggers
	err = c.callTriggers(event)
	if err != nil {
		return fmt.Errorf("Error calling triggers: err: %v", err)
	}
	log.Infof("handleEvent: triggers called: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

	// We need to get past newsroom events for the newsroom contract of a newly added watcher
	if event.EventType() == "Application" {
		newsroomAddr := event.EventPayload()["ListingAddress"].(common.Address)
		newsroomEvents, err := c.FilterAddedNewsroomContract(newsroomAddr)
		if err != nil {
			return fmt.Errorf("Error filtering new newsroom contract: err: %v", err)
		}
		log.Infof("Found %v newsroom events for address %v after filtering: hsh: %v",
			len(newsroomEvents), newsroomAddr.Hex(), hash) // Debug, remove later
		err = c.eventDataPersister.SaveEvents(newsroomEvents)
		if err != nil {
			return fmt.Errorf("Error saving events for application %v", err)
		}
		log.Infof("Saved newsroom events at address %v, hsh: %v", newsroomAddr.Hex(), hash) //Debug, remove later

		if c.crawlerPubSub != nil {
			err := c.crawlerPubSub.PublishNewsroomExceptionMessage(newsroomAddr.Hex())
			if err != nil {
				return fmt.Errorf("Error sending message for event %v to pubsub: %v", event.Hash(), err)
			}
		}
	}

	log.Infof("handleEvent: done: %v, tx: %v, hsh: %v", eventType, txHash.Hex(), hash) // Debug, remove later
	return nil
}

// FilterAddedNewsroomContract runs a filterer on the newly watched newsroom contract to ensure we have all events.
func (c *EventCollector) FilterAddedNewsroomContract(newsroomAddr common.Address) ([]*model.Event, error) {
	nwsrmFilterer := filterer.NewNewsroomContractFilterers(newsroomAddr)
	c.updateFiltererStartingBlock(nwsrmFilterer)
	retrieve := retriever.NewEventRetriever(c.client, []model.ContractFilterers{nwsrmFilterer})
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
	err := c.runRetriever()
	if err != nil {
		c.sendStartSignal()
		return err
	}

	if c.disableListener {
		log.Infof("Listener is disabled, not starting")
		c.sendStartSignal()
		return nil
	}

	if c.crawlerPubSub != nil {
		err = c.crawlerPubSub.PublishTriggerMessage()
		if err != nil {
			return err
		}
	}

	err = c.startListener()
	if err != nil {
		c.sendStartSignal()
		return fmt.Errorf("Error starting listener: err: %v", err)
	}

	c.sendStartSignal()

	select {
	case err = <-c.errorsChan:
		log.Errorf("Received error on the collector errors chan: err: %v", err)
		return fmt.Errorf("Error during event handling: err: %v", err)
	case <-c.quitChan:
		return nil
	}
}

func (c *EventCollector) runRetriever() error {
	err := c.retrieveEvents(c.filterers)
	if err != nil {
		return fmt.Errorf("Error retrieving events: err: %v", err)
	}
	pastEvents := c.retrieve.PastEvents

	// Check pastEvents for any new newsrooms to track
	additionalEvents, err := c.CheckRetrievedEventsForNewsroom(pastEvents)
	if err != nil {
		return fmt.Errorf("Error checking newsroom events during filterer, err: %v", err)
	}
	if len(additionalEvents) > 0 {
		pastEvents = append(pastEvents, additionalEvents...)
	}
	err = c.retrieve.SortEventsByBlock(pastEvents)
	if err != nil {
		return fmt.Errorf("Error sorting retrieved events: %v", err)
	}
	err = c.updateEventTimesFromBlockHeaders(pastEvents)
	if err != nil {
		return fmt.Errorf("Error updating dates for events: err: %v", err)
	}
	err = c.eventDataPersister.SaveEvents(pastEvents)
	if err != nil {
		return fmt.Errorf("Error persisting events: err: %v", err)
	}
	err = c.persistRetrieverLastBlockData()
	if err != nil {
		return fmt.Errorf("Error persisting last block data: err: %v", err)
	}
	return nil
}

func (c *EventCollector) startListener() error {
	defer c.mutex.Unlock()
	c.mutex.Lock()

	c.listen = listener.NewEventListener(c.client, c.watchers)
	if c.listen == nil {
		return errors.New("Listener should not be nil")
	}

	err := c.listen.Start()
	if err != nil {
		return fmt.Errorf("Listener should have started with no errors: %v", err)
	}

	go func(quit <-chan bool, errors chan<- error) {
		multiplier := 1
		numCPUs := runtime.NumCPU() * multiplier
		pool := tunny.NewFunc(numCPUs, c.handleEvent)
		defer pool.Close()

		for {
			select {
			case event := <-c.listen.EventRecvChan:
				if log.V(2) {
					log.Infof(
						"event received: %v, %v, %v, \n%v",
						event.EventType(),
						event.Hash(),
						event.Timestamp(),
						event.LogPayloadToString(),
					)
				}
				go func(e *model.Event, errs chan<- error) {
					log.Infof(
						"startListener func: pool.Process start: queued: %v, queue_size: %v",
						pool.QueueLength(),
						pool.GetSize(),
					) // Debug, remove later
					result := pool.Process(
						handleEventInputs{
							event:  e,
							errors: errs,
						},
					)
					if result != nil {
						err := result.(error)
						if err != nil {
							log.Errorf(
								"startListener go func: Error processing: %v, %v, %v, %v, \n%v",
								err,
								event.EventType(),
								event.Hash(),
								event.Timestamp(),
								event.LogPayloadToString(),
							)
						}
					}
					if log.V(2) {
						log.Infof(
							"startListener func: pool.Process done: %v, %v, %v, \n%v",
							event.EventType(),
							event.Hash(),
							event.Timestamp(),
							event.LogPayloadToString(),
						) // Debug, remove later
					}
				}(event, errors)
				log.Infof("startListener: started process from pool") // Debug, remove later
			case <-quit:
				log.Infof("startListener: quit") // Debug, remove later
				return
			}
		}
	}(c.quitChan, c.errorsChan)

	return nil
}

// StopCollection is for stopping the listener
func (c *EventCollector) StopCollection() error {
	var err error
	if c.crawlerPubSub != nil {
		err = c.crawlerPubSub.StopPublishers()
		if err != nil {
			return err
		}
	}

	if c.listen != nil {
		err = c.listen.Stop()
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

func (c *EventCollector) sendStartSignal() {
	if c.startChan != nil {
		close(c.startChan)
	}
}

// UpdateStartingBlocks updates starting blocks for retriever based on persistence
func (c *EventCollector) updateRetrieverStartingBlocks() {
	for _, filter := range c.filterers {
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
		log.Infof(
			"updateEventTimeFromBlockHeader: done calling headerbynumber: %v",
			header.TxHash.Hex(),
		) // Debug, remove later

		c.headerCache.AddHeader(event.BlockNumber(), header)
	}
	if err != nil {
		return err
	}
	event.SetTimestamp(header.Time.Int64())
	return nil
}

func (c *EventCollector) retrieveEvents(filterers []model.ContractFilterers) error {
	c.updateRetrieverStartingBlocks()
	c.retrieve = retriever.NewEventRetriever(c.client, filterers)
	return c.retrieve.Retrieve()
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
			}
			if _, ok := existingWatcherNewsroomAddr[newsroomAddr]; !ok {
				newWatcher := watcher.NewNewsroomContractWatchers(newsroomAddr)
				watchersToAdd[newsroomAddr] = newWatcher
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
			return additionalEvents, fmt.Errorf("Error retrieving new Newsroom events: err: %v", err)
		}
		additionalEvents = append(additionalEvents, c.retrieve.PastEvents...)
	}
	return additionalEvents, nil
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
