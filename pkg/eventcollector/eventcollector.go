// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	"fmt"
	"sync"
	"time"

	log "github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/joincivil/civil-events-crawler/pkg/generated/filterer"
	"github.com/joincivil/civil-events-crawler/pkg/generated/watcher"
	"github.com/joincivil/civil-events-crawler/pkg/listener"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/pubsub"

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
	PreemptSecs         *int
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
		startChan:           make(chan struct{}),
		shutdownChan:        make(chan struct{}),
		listenerStopChan:    make(chan struct{}),
		listenerErrChan:     make(chan error),
		retrieverStopChan:   make(chan struct{}),
		retrieverErrChan:    make(chan error),
		crawlerPubSub:       config.CrawlerPubSub,
		pollingEnabled:      config.PollingEnabled,
		pollingIntervalSecs: config.PollingIntervalSecs,
		preemptSecs:         config.PreemptSecs,
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

	// Most recent listener
	listen *listener.EventListener

	startBlock uint64

	// startChan is closed when the event collector has started
	startChan chan struct{}

	// shutdownChan is closed to stop all processes of the event collector
	shutdownChan chan struct{}

	retrieverStopChan chan struct{}

	retrieverErrChan chan error

	listenerStopChan chan struct{}

	listenerErrChan chan error

	mutex sync.Mutex

	headerCache *eth.BlockHeaderCache

	crawlerPubSub *pubsub.CrawlerPubSub

	pollingEnabled bool

	pollingIntervalSecs int

	preemptSecs *int
}

// StartChan returns the channel will send a "event collector started" signal
func (c *EventCollector) StartChan() chan struct{} {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	return c.startChan
}

// StartCollection contains logic to run retriever and listener.
func (c *EventCollector) StartCollection() error {
	defer c.resetCollector()
	var once sync.Once
	var err error
	var l *listener.EventListener
	pd := c.preemptDuration()

	// This is the main loop
	for {
		c.retrieverStopChan = make(chan struct{})
		err = c.runRetrieverLoop(c.pollingEnabled)
		if err != nil {
			once.Do(func() { c.sendStartSignal() })
			return errors.Wrap(err, "startcollection.runRetrieverLoop")
		}

		if !c.pollingEnabled && c.isListenerEnabled() {
			c.listenerStopChan = make(chan struct{})
			l, err = c.startListenerLoop()
			if err != nil {
				once.Do(func() { c.sendStartSignal() })
				log.Errorf("Error starting listener: err: %v", err)
				return errors.Wrap(err, "error starting listener in startcol")
			}
			c.mutex.Lock()
			c.listen = l
			c.mutex.Unlock()
			log.Infof("Listeners enabled")
			once.Do(func() { c.sendStartSignal() })
		}

		if c.pollingEnabled {
			once.Do(func() { c.sendStartSignal() })
			log.Infof("Polling via retrievers only")
		}

		// All errors from the filterers and retriever loop get funnelled here.
		// If any errors occur, reset the collector and restart
		select {
		case <-time.After(pd):
			log.Infof("Premptive restart...")
			c.resetCollector()

		case err = <-c.retrieverErrChan:
			err = errors.Cause(err)
			log.Errorf("Received error from retriever, resetting collector: err: %v", err)
			c.errRep.Error(err, nil)
			c.resetCollector()

		// All errors from the watchers and startListener loop get funnelled here.
		case err = <-c.listenerErrChan:
			err = errors.Cause(err)
			log.Errorf("Received error from listener, resetting collector: err: %v", err)
			c.errRep.Error(err, nil)
			c.resetCollector()

		case <-c.shutdownChan:
			log.Infof("Received shutdownChan signal, stopping collection")
			return nil
		}

		log.Infof("Restarting collection loop")
	}
}

// StopCollection is for stopping the listener
func (c *EventCollector) StopCollection(unsubWatchers bool) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
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
	if c.shutdownChan != nil {
		close(c.shutdownChan)
	}
	return err
}

// CheckRetrievedEventsForNewsroom checks pastEvents for TCR events that may
// include new newsroom events, creates new Newsroom filterers and watchers
// upon valid events, filters for events, and then returns those events.
func (c *EventCollector) CheckRetrievedEventsForNewsroom(pastEvents []*model.Event) (
	[]*model.Event, error) {
	existingFiltererNewsroomAddr := c.getExistingNewsroomFilterers()
	existingWatcherNewsroomAddr := c.getExistingNewsroomWatchers()

	watchersToAdd := map[common.Address]model.ContractWatchers{}
	filterersToAdd := map[common.Address]model.ContractFilterers{}
	eventsToAdd := []*model.Event{}

	for _, event := range pastEvents {

		// NOTE(IS): We should track events from "Application" so we don't miss other events.
		if event.EventType() == "Application" {
			newsroomAddr, ok := event.EventPayload()["ListingAddress"].(common.Address)
			if !ok {
				return eventsToAdd, fmt.Errorf("Cannot get newsroomAddr from eventpayload")
			}

			if _, ok := existingFiltererNewsroomAddr[newsroomAddr]; !ok {
				newFilterer := filterer.NewNewsroomContractFilterers(newsroomAddr)
				filterersToAdd[newsroomAddr] = newFilterer
				existingFiltererNewsroomAddr[newsroomAddr] = true
			}

			if _, ok := existingWatcherNewsroomAddr[newsroomAddr]; !ok {
				// log.Infof("Adding Newsroom watcher for %v", newsroomAddr.Hex())
				newWatcher := watcher.NewNewsroomContractWatchers(newsroomAddr)
				watchersToAdd[newsroomAddr] = newWatcher
				existingWatcherNewsroomAddr[newsroomAddr] = true
			}
		}

		if event.EventType() == "ListingRemoved" {
			newsroomAddr, ok := event.EventPayload()["ListingAddress"].(common.Address)
			if !ok {
				return eventsToAdd, fmt.Errorf("Cannot get newsroomAddr from eventpayload")
			}
			watchersToAdd[newsroomAddr] = nil
		}
	}

	var e error
	for addr, watcher := range watchersToAdd {
		if watcher != nil {
			log.Infof("Adding new newsroom watcher for %v", addr.Hex())
			// Add new watcher to the list of watchers in the collector
			// and listener
			e = c.AddWatchers(watcher)
			if e != nil {
				log.Errorf("Error adding new filterer: err: %v", e)
			}
		} else {
			log.Infof("Not adding %v to watchers because it was removed.", addr.Hex())
		}

	}

	if len(filterersToAdd) > 0 {
		newFilts := make([]model.ContractFilterers, len(filterersToAdd))
		ind := 0
		for addr, filt := range filterersToAdd {
			log.Infof("Adding new newsroom filterer for %v", addr.Hex())
			newFilts[ind] = filt
			ind++
			// Add new filterer to the list of filterers in the collector
			// and retriever
			e = c.AddFilterers(filt)
			if e != nil {
				log.Errorf("Error adding new filterer: err: %v", e)
			}
		}

		// Retrieve events for the new filters only
		r, err := c.retrieveEvents(newFilts, false)
		if err != nil {
			return eventsToAdd, errors.WithMessage(err, "error retrieving new Newsroom events")
		}
		log.Infof("Finished retrieving events for new newsrooms")
		eventsToAdd = append(eventsToAdd, r.PastEvents...)
	}

	return eventsToAdd, nil
}

func (c *EventCollector) resetCollector() {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	// Stop the currently running retriever loop
	close(c.retrieverStopChan)
	// nil this out to kill start signal, not needed at this point
	c.startChan = nil
	// Stop the old listener before starting new one
	if c.listen != nil {
		if err := c.listen.Stop(true); err != nil {
			log.Errorf("Error stopping listener: err: %v", err)
			c.errRep.Error(err, nil)
		}
	}
	// Stop the currently running listener loop
	close(c.listenerStopChan)
	c.listen = nil
}

func (c *EventCollector) preemptDuration() time.Duration {
	if c.preemptSecs == nil || *c.preemptSecs == 0 {
		// Don't preempt == a very long time
		log.Infof("Disabling preempt")
		return time.Hour * 24 * 365 * 10
	}
	log.Infof("Preempt delay secs: %v", *c.preemptSecs)
	return time.Duration(*c.preemptSecs) * time.Second
}

func (c *EventCollector) sendStartSignal() {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	if c.startChan != nil {
		close(c.startChan)
	}
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
