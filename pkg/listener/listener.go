// Package listener contains all the components for the events listener, which
// streams a list of future events.
package listener // import "github.com/joincivil/civil-events-crawler/pkg/listener"

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	log "github.com/golang/glog"
	"github.com/joincivil/civil-events-crawler/pkg/generated/tcr"
	"github.com/joincivil/civil-events-crawler/pkg/model"
)

// NewCivilEventListener creates a new CivilEventListener given the address
// of the contract to listen to.
func NewCivilEventListener(client *ethclient.Client, contractAddress string) *CivilEventListener {
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

	// Client is the ethereum client from go-ethereum
	Client *ethclient.Client

	// ContractAddress is the Address type for the contract to watch
	ContractAddress common.Address

	// ContractAddressStr is the string repr for the address of the contract
	ContractAddressStr string

	// EventRecvChan is the channel to send and receive CivilEvents
	EventRecvChan chan model.CivilEvent

	watcherSubs []*event.Subscription
}

// Start starts up the event listener
func (l *CivilEventListener) Start() error {
	civilTCR, err := tcr.NewCivilTCRContract(l.ContractAddress, l.Client)
	if err != nil {
		log.Errorf("Error initializing TCR: %v", err)
		return err
	}
	subs, err := l.startWatchers(civilTCR)
	if err != nil {
		return err
	}
	if len(subs) <= 0 {
		return errors.New("No watchers have been started")
	}
	l.watcherSubs = subs
	return nil
}

// Stop shuts down the event listener and performs clean up
func (l *CivilEventListener) Stop() error {
	for _, sub := range l.watcherSubs {
		(*sub).Unsubscribe()
	}
	return nil
}

func (l *CivilEventListener) startWatchers(civilTCR *tcr.CivilTCRContract) ([]*event.Subscription, error) {
	subs := []*event.Subscription{}
	sub, err := startWatchApplication(l.EventRecvChan, civilTCR)
	if err != nil {
		return nil, fmt.Errorf("Error starting _Application watch: %v", err)
	}
	subs = append(subs, &sub)

	sub, err = startWatchApplicationRemoved(l.EventRecvChan, civilTCR)
	if err != nil {
		return nil, fmt.Errorf("Error starting _ApplicationRemoved watch: %v", err)
	}
	subs = append(subs, &sub)

	sub, err = startWatchApplicationWhitelisted(l.EventRecvChan, civilTCR)
	if err != nil {
		return nil, fmt.Errorf("Error starting _ApplicationWhitelisted watch: %v", err)
	}
	subs = append(subs, &sub)

	return subs, nil
}
