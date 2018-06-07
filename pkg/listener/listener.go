// Package listener contains all the components for the events listener, which
// streams a list of future events.
package listener // import "github.com/joincivil/civil-events-crawler/pkg/listener"

import (
	"errors"
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/joincivil/civil-events-crawler/pkg/model"
)

// NewCivilEventListener creates a new CivilEventListener given the address
// of the contract to listen to.
func NewCivilEventListener(client bind.ContractBackend, contractAddress string,
	watchers []model.ContractWatchers) *CivilEventListener {
	address := common.HexToAddress(contractAddress)
	listener := &CivilEventListener{
		ContractAddress:    address,
		ContractAddressStr: contractAddress,
		EventRecvChan:      make(chan model.CivilEvent),
		client:             client,
		watchers:           watchers,
	}
	return listener
}

// CivilEventListener handles the listener stream for Civil-specific events
type CivilEventListener struct {

	// client is a ethereum backend from go-ethereum
	client bind.ContractBackend

	// ContractAddress is the Address type for the contract to watch
	ContractAddress common.Address

	// ContractAddressStr is the string repr for the address of the contract
	ContractAddressStr string

	// EventRecvChan is the channel to send and receive CivilEvents
	EventRecvChan chan model.CivilEvent

	watchers []model.ContractWatchers

	watcherSubs []event.Subscription
}

// Start starts up the event listener and watchers
func (l *CivilEventListener) Start() error {
	l.watcherSubs = []event.Subscription{}

	var err error
	var subs []event.Subscription
	for _, watcher := range l.watchers {
		subs, err = watcher.StartWatchers(
			l.client,
			l.ContractAddress,
			l.EventRecvChan,
		)
		if err != nil {
			log.Errorf("Error starting watchers for %v: err: %v",
				watcher.ContractName(), err)
		}
		l.watcherSubs = append(l.watcherSubs, subs...)
	}

	if len(subs) <= 0 {
		return errors.New("No watchers have been started")
	}

	return nil
}

// Stop shuts down the event listener and performs clean up
func (l *CivilEventListener) Stop() error {
	for _, sub := range l.watcherSubs {
		sub.Unsubscribe()
	}
	l.watcherSubs = nil
	return nil
}
