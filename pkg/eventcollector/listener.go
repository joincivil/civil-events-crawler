package eventcollector

import (
	"runtime"
	"time"

	"github.com/Jeffail/tunny"
	"github.com/davecgh/go-spew/spew"
	log "github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	specs "github.com/joincivil/civil-events-crawler/pkg/contractspecs"
	"github.com/joincivil/civil-events-crawler/pkg/listener"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

// AddWatchers will add watchers to the embedded listener.
func (c *EventCollector) AddWatchers(w model.ContractWatchers) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	c.watchers = append(c.watchers, w)
	// If has active listener, add and run watchers
	if c.listen != nil {
		return c.listen.AddWatchers(w)
	}
	return nil
}

// RemoveWatchers will remove given watcher from the embedded listener.
func (c *EventCollector) RemoveWatchers(w model.ContractWatchers) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	if c.watchers != nil && len(c.watchers) > 0 {
		for index, ew := range c.watchers {
			if w.ContractAddress() == ew.ContractAddress() &&
				w.ContractName() == ew.ContractName() {
				// Delete the item in the watchers list.
				copy(c.watchers[index:], c.watchers[index+1:])
				c.watchers[len(c.filterers)-1] = nil
				c.watchers = c.watchers[:len(c.filterers)-1]
				return nil
			}
		}
	}
	return c.listen.RemoveWatchers(w)
}

func (c *EventCollector) getExistingNewsroomWatchers() map[common.Address]bool {
	existingNewsroomAddr := map[common.Address]bool{}
	c.mutex.Lock()
	for _, existing := range c.watchers {
		specs, _ := specs.ContractTypeToSpecs.Get(specs.NewsroomContractType)
		if existing.ContractName() == specs.Name() {
			existingNewsroomAddr[existing.ContractAddress()] = true
		}
	}
	c.mutex.Unlock()
	return existingNewsroomAddr
}

func (c *EventCollector) isListenerEnabled() bool {
	if (c.wsEthURL == "" && c.wsClient == nil) || c.pollingEnabled {
		return false
	}
	return true
}

func (c *EventCollector) initWsClient() (bind.ContractBackend, chan struct{}, error) {
	var killChan chan struct{}
	if c.wsClient != nil {
		return c.wsClient, nil, nil
	}
	killChan = make(chan struct{})
	ethclient, err := utils.SetupWebsocketEthClient(c.wsEthURL, killChan, wsPingDelaySecs)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "unable to setup ws client")
	}
	wsClient := bind.ContractBackend(ethclient)
	return wsClient, killChan, nil
}

func (c *EventCollector) startListenerLoop() (*listener.EventListener, error) {
	defer c.mutex.Unlock()
	c.mutex.Lock()

	wsClient, killChan, err := c.initWsClient()
	if err != nil {
		return nil, errors.WithMessage(err, "startupListenerLoop: Unable to create new websocket client")
	}

	listener := listener.NewEventListener(wsClient, c.watchers)
	if listener == nil {
		return nil, errors.New("startupListenerLoop: Listener should not be nil")
	}

	_, err = listener.Start()
	if err != nil {
		return nil, errors.WithMessage(err, "startupListenerLoop: Listener should have started with no errors")
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

	go func(shutdown <-chan struct{}, quit <-chan struct{}, errs chan<- error) {
		multiplier := 1
		numCPUs := runtime.NumCPU() * multiplier
		pool := tunny.NewFunc(numCPUs, c.handleEvent)
		defer pool.Close()

		for {
			select {
			case event := <-listener.EventRecvChan:
				c.processEventsFromEventRecvChan(pool, event, errs)

			case err := <-listener.Errors:
				// Any errors from the watchers
				log.Infof("startupListenerLoop: watcher error chan: %v", err)
				// make sure we send these errors to the main errors chan
				errs <- errors.WithMessage(err, "startupListenerLoop: listener.Errors")
				c.flushEvents(listener.EventRecvChan, pool, errs)
				c.flushErrors(listener.Errors)
				cleanupFn()
				return

			case <-quit:
				log.Infof("startupListenerLoop: Quit event recv loop")
				c.flushEvents(listener.EventRecvChan, pool, errs)
				c.flushErrors(listener.Errors)
				cleanupFn()
				return

			case <-shutdown:
				log.Infof("startupListenerLoop: Shutdown event recv loop")
				cleanupFn()
				return
			}
		}
	}(c.shutdownChan, c.listenerStopChan, c.listenerErrChan)

	return listener, nil
}

func (c *EventCollector) flushEvents(evts chan *model.Event, pool *tunny.Pool,
	errs chan<- error) {
	for {
		select {
		case e := <-evts:
			log.Infof("Flushing event")
			c.processEventsFromEventRecvChan(pool, e, errs)

		case <-time.After(time.Second * time.Duration(60)):
			return
		}
	}
}

func (c *EventCollector) flushErrors(errs chan error) {
	// Flush errors chan since we are only handling the first we receive
	// and killing this consumer. It is possible a number of watchers
	// receive the same error and push it this channel. Without flushing
	// the channel may block indefinitely.
	for {
		select {
		case e := <-errs:
			if e != nil {
				log.Infof("Flushed error: %v", e)
			}
		case <-time.After(time.Second * time.Duration(30)):
			return
		}
	}
}

func (c *EventCollector) processEventsFromEventRecvChan(pool *tunny.Pool,
	event *model.Event, errs chan<- error) {
	if log.V(2) {
		log.Infof(
			"startupListener: Recv loop event received: %v",
			spew.Sprintf("%#+v", event),
		)
	} else {
		log.Infof(
			"startupListenerLoop: Recv loop event received: eventType: %v, hash: %v, ts: %v",
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
					"startupListenerLoop: pool.Process Error processing, recovering: err: %v: %v",
					err,
					spew.Sdump(event),
				)
				c.errRep.Error(err, nil)
			}
		}
		if log.V(2) {
			log.Infof(
				"startupListenerLoop: pool.Process done: %v",
				spew.Sprintf("%#+v", event),
			)
		} else {
			log.Infof(
				"startupListenerLoop: pool.Process done: %v, %v, %v",
				event.EventType(),
				event.Hash(),
				event.Timestamp(),
			)
		}
	}(event, errs)
}
