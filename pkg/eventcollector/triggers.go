// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	"errors"

	log "github.com/golang/glog"

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
	newsroomAddr := event.LogPayload().Address
	err := collector.AddWatchers(
		watcher.NewNewsroomContractWatchers(newsroomAddr),
	)
	if err != nil {
		return err
	}
	log.Infof("Adding watchers for newsroom at address: %v", newsroomAddr)
	return nil
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
	newsroomAddr := event.LogPayload().Address
	err := collector.RemoveWatchers(
		watcher.NewNewsroomContractWatchers(newsroomAddr),
	)
	if err != nil {
		return err
	}
	log.Infof("Removing watchers for newsroom at address: %v", newsroomAddr)
	return nil
}
