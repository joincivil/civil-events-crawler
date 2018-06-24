// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	"errors"
	"github.com/joincivil/civil-events-crawler/pkg/generated/watcher"
	"github.com/joincivil/civil-events-crawler/pkg/model"
)

// Civil specific triggers

// AddNewsroomWatchersTrigger is a Trigger that starts up a new watcher for a newly
// created newsroom.
type AddNewsroomWatchersTrigger struct{}

// Description returns the description of the trigger
func (n *AddNewsroomWatchersTrigger) Description() string {
	return "Start watching for newsroom events on TCR whitelisting events"
}

// ShouldRun returns true or false on whether this trigger should be run
func (n *AddNewsroomWatchersTrigger) ShouldRun(collector *CivilEventCollector,
	event *model.CivilEvent) bool {
	return event.EventType() == "_ApplicationWhiteListed"
}

// Run returns the triggered code
func (n *AddNewsroomWatchersTrigger) Run(collector *CivilEventCollector,
	event *model.CivilEvent) error {
	if !n.ShouldRun(collector, event) {
		return errors.New("AddNewsroomWatchersTrigger should not run")
	}
	addrValue, ok := event.Payload().Value("ListingAddress")
	if !ok {
		return errors.New("No newsroom address found")
	}
	newsroomAddr, ok := addrValue.Address()
	if !ok {
		return errors.New("Invalid address value")
	}
	return collector.AddWatchers(
		watcher.NewNewsroomContractWatchers(*newsroomAddr),
	)
}

// RemoveNewsroomWatchersTrigger is a Trigger that removes a new watcher for a
// newsroom to be delisted.
type RemoveNewsroomWatchersTrigger struct{}

// Description returns the description of the trigger
func (n *RemoveNewsroomWatchersTrigger) Description() string {
	return "Remove watching of newsroom events"
}

// ShouldRun returns true or false on whether this trigger should be run
func (n *RemoveNewsroomWatchersTrigger) ShouldRun(collector *CivilEventCollector,
	event *model.CivilEvent) bool {
	return event.EventType() == "_ListingRemoved"
}

// Run returns the triggered code
func (n *RemoveNewsroomWatchersTrigger) Run(collector *CivilEventCollector,
	event *model.CivilEvent) error {
	if !n.ShouldRun(collector, event) {
		return errors.New("RemoveNewsroomWatchersTriggershould not run")
	}
	addrValue, ok := event.Payload().Value("ListingAddress")
	if !ok {
		return errors.New("No newsroom address found")
	}
	newsroomAddr, ok := addrValue.Address()
	if !ok {
		return errors.New("Invalid address value")
	}
	return collector.RemoveWatchers(
		watcher.NewNewsroomContractWatchers(*newsroomAddr),
	)
}
