// Package model contains the general data models and interfaces.
package model // import "github.com/joincivil/civil-events-crawler/pkg/model"

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// ContractWatchers defines an interface that starts up a particular set of listeners watcher loops.
type ContractWatchers interface {
	ContractName() string
	ContractAddress() common.Address
	StartWatchers(client bind.ContractBackend, eventRecvChan chan *Event) ([]event.Subscription, error)
	StopWatchers() error
}

// ContractFilterers defines an interface that starts up a particular set of filterers
type ContractFilterers interface {
	ContractName() string
	ContractAddress() common.Address
	EventTypes() []string
	UpdateStartBlock(eventType string, startBlock uint64)
	LastEvents() []*Event
	StartFilterers(client bind.ContractBackend, pastEvents []*Event) (error, []*Event)
}
