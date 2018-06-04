// Package listener contains all the components for the events listener, which
// streams a list of future events.
package listener // import "github.com/joincivil/civil-events-crawler/pkg/listener"

//go:generate sh -c "mkdir -p ../generated/watcher"
//go:generate sh -c "go run ../../cmd/watchergen/main.go civiltcr watcher > ../generated/watcher/civiltcr.go"
//go:generate sh -c "go run ../../cmd/watchergen/main.go newsroom watcher > ../generated/watcher/newsroom.go"

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/joincivil/civil-events-crawler/pkg/generated/watcher"
	"github.com/joincivil/civil-events-crawler/pkg/model"
)

// NewCivilEventListener creates a new CivilEventListener given the address
// of the contract to listen to.
func NewCivilEventListener(client bind.ContractBackend, contractAddress string) *CivilEventListener {
	address := common.HexToAddress(contractAddress)
	listener := &CivilEventListener{
		Client:             client,
		ContractAddress:    address,
		ContractAddressStr: contractAddress,
		EventRecvChan:      make(chan model.CivilEvent),
	}
	return listener
}

// CivilEventListener handles the listener stream for Civil-specific events
type CivilEventListener struct {

	// Client is a ethereum backend from go-ethereum
	Client bind.ContractBackend

	// ContractAddress is the Address type for the contract to watch
	ContractAddress common.Address

	// ContractAddressStr is the string repr for the address of the contract
	ContractAddressStr string

	// EventRecvChan is the channel to send and receive CivilEvents
	EventRecvChan chan model.CivilEvent

	watcherSubs []event.Subscription
}

// Start starts up the event listener
func (l *CivilEventListener) Start() error {
	// StartCivilTCRContractWatchers is generated
	subs, err := watcher.StartCivilTCRContractWatchers(
		l.Client,
		l.ContractAddress,
		l.EventRecvChan,
	)
	if err != nil {
		return err
	}
	l.watcherSubs = subs

	// StartNewsroomContractWatchers is generated
	subs, err = watcher.StartNewsroomContractWatchers(
		l.Client,
		l.ContractAddress,
		l.EventRecvChan,
	)
	if err != nil {
		return err
	}
	l.watcherSubs = append(l.watcherSubs, subs...)

	if len(subs) <= 0 {
		return errors.New("No watchers have been started")
	}
	l.watcherSubs = subs
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
