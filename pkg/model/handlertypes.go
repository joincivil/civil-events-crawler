// Package model contains the general data models and interfaces for the Civil crawler.
package model // import "github.com/joincivil/civil-events-crawler/pkg/model"

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
)

// ContractWatchers defines an interface that starts up a particular set of listeners watcher loops.
type ContractWatchers interface {
	ContractName() string
	StartWatchers(client bind.ContractBackend, eventRecvChan chan CivilEvent) ([]event.Subscription, error)
}

// ContractFilterers defines an interface for event retrievers.
type ContractFilterers interface {
	ContractName() string
	StartFilterers(client bind.ContractBackend, pastEvents *[]CivilEvent) error
	EventNames() []string
	UpdateStartBlock(eventName string, startBlock int)
}
