// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	log "github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joincivil/civil-events-crawler/pkg/generated/watcher"
	"github.com/joincivil/civil-events-crawler/pkg/model"
)

// Civil specific triggers

// AddNewsroomWatchersTrigger is a Trigger that starts up a new watcher for a newly
// created newsroom.
type AddNewsroomWatchersTrigger struct{}

// Description returns the description of the trigger
func (n *AddNewsroomWatchersTrigger) Description() string {
	return "Start watching for newsroom events as soon as newsroom is created in factory"
}

// ShouldRun returns true or false on whether this trigger should be run
func (n *AddNewsroomWatchersTrigger) ShouldRun(collector *EventCollector,
	event *model.Event) bool {
	return event.ContractName() == "NewsroomFactory" && event.EventType() == "ContractInstantiation"
}

// Run returns the triggered code
func (n *AddNewsroomWatchersTrigger) Run(collector *EventCollector,
	event *model.Event) error {
	if !n.ShouldRun(collector, event) {
		return errors.New("AddNewsroomWatchersTrigger should not run")
	}
	newsroomAddr := event.EventPayload()["Instantiation"].(common.Address)
	err := collector.AddWatchers(
		watcher.NewNewsroomContractWatchers(newsroomAddr),
	)
	if err != nil {
		return errors.WithMessage(err, "error adding watchers")
	}
	log.Infof("Adding watchers for newsroom at address: %v", newsroomAddr.Hex())
	return nil
}

// AddMultisigWatchersTrigger is a Trigger that starts up a new watcher for a newly
// created multisig.
type AddMultisigWatchersTrigger struct{}

// Description returns the description of the trigger
func (n *AddMultisigWatchersTrigger) Description() string {
	return "Start watching for multisig events as soon as multisig is created in factory"
}

// ShouldRun returns true or false on whether this trigger should be run
func (n *AddMultisigWatchersTrigger) ShouldRun(collector *EventCollector,
	event *model.Event) bool {
	return event.ContractName() == "MultisigFactory" && event.EventType() == "ContractInstantiation"
}

// Run returns the triggered code
func (n *AddMultisigWatchersTrigger) Run(collector *EventCollector,
	event *model.Event) error {
	if !n.ShouldRun(collector, event) {
		return errors.New("AddMultisigWatchersTrigger should not run")
	}
	multisigAddr := event.EventPayload()["Instantiation"].(common.Address)
	err := collector.AddWatchers(
		watcher.NewMultiSigWalletContractWatchers(multisigAddr),
	)
	if err != nil {
		return errors.WithMessage(err, "error adding watchers")
	}
	log.Infof("Adding watchers for multisig at address: %v", multisigAddr.Hex())
	return nil
}
