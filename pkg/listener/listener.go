// Package listener contains all the components for the events listener, which
// streams a list of future events.
package listener // import "github.com/joincivil/civil-events-crawler/pkg/listener"

import (
	"errors"
	log "github.com/golang/glog"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/joincivil/civil-events-crawler/pkg/model"
)

const (
	eventRecvChanBufferSize = 1
)

// NewCivilEventListener creates a new CivilEventListener given the address
// of the contract to listen to.
func NewCivilEventListener(client bind.ContractBackend, watchers []model.ContractWatchers) *CivilEventListener {
	listener := &CivilEventListener{
		EventRecvChan: make(chan model.CivilEvent, eventRecvChanBufferSize),
		client:        client,
		watchers:      watchers,
		active:        false,
	}
	return listener
}

// CivilEventListener handles the listener stream for Civil-specific events
type CivilEventListener struct {

	// client is a ethereum backend from go-ethereum
	client bind.ContractBackend

	// EventRecvChan is the channel to send and receive CivilEvents
	EventRecvChan chan model.CivilEvent

	watchers []model.ContractWatchers

	active bool

	mutex sync.Mutex
}

// Start starts up the event listener and watchers
func (l *CivilEventListener) Start() error {
	defer l.mutex.Unlock()
	l.mutex.Lock()
	hasSubs := false
	for _, watchers := range l.watchers {
		newSubs, err := watchers.StartWatchers(
			l.client,
			l.EventRecvChan,
		)
		if err != nil {
			log.Errorf("Error starting watchers for %v at %v: err: %v",
				watchers.ContractName(), watchers.ContractAddress(), err)
		}
		if len(newSubs) > 0 {
			hasSubs = true
		}
	}

	if !hasSubs {
		return errors.New("No watchers have been started")
	}

	l.active = true
	return nil
}

// AddWatchers will add watchersto the listener. If the listener is already
// started, add to the list of watchers, start up with the watcher, and add it
// to the list of subscriptions in the listener.
// If the listener is not already started, will just be added to the list of watchers.
func (l *CivilEventListener) AddWatchers(w model.ContractWatchers) error {
	defer l.mutex.Unlock()
	l.mutex.Lock()
	l.watchers = append(l.watchers, w)
	if l.active {
		_, err := w.StartWatchers(
			l.client,
			l.EventRecvChan,
		)
		if err != nil {
			log.Errorf("Error starting watchers for %v at %v: err: %v",
				w.ContractName(), w.ContractAddress(), err)
			return err
		}
	}
	return nil
}

// RemoveWatchers will remove given watcher from the listener. If the listener is already
// started, stop the watcher, removes the subscription, and removes from watcher list.
// If the listener is not already started, will just be removed from the list of watchers.
func (l *CivilEventListener) RemoveWatchers(w model.ContractWatchers) error {
	defer l.mutex.Unlock()
	l.mutex.Lock()
	if l.watchers != nil && len(l.watchers) > 0 {
		for index, ew := range l.watchers {
			if w.ContractAddress() == ew.ContractAddress() &&
				w.ContractName() == ew.ContractName() {
				if l.active {
					_ = ew.StopWatchers()
				}
				// Delete the item in the watchers list.
				copy(l.watchers[index:], l.watchers[index+1:])
				l.watchers[len(l.watchers)-1] = nil
				l.watchers = l.watchers[:len(l.watchers)-1]
				return nil
			}
		}
	}
	return nil
}

// Stop shuts down the event listener and performs clean up
func (l *CivilEventListener) Stop() error {
	defer l.mutex.Unlock()
	l.mutex.Lock()
	if l.watchers != nil && len(l.watchers) > 0 {
		for _, w := range l.watchers {
			_ = w.StopWatchers()
		}
	}
	l.active = false
	return nil
}
