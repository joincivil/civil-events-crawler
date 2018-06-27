// Package eventcollector contains business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

import (
	"github.com/joincivil/civil-events-crawler/pkg/model"
)

// Trigger is an interface for triggering code to run upon receiving
// an event type
type Trigger interface {
	// Description returns the description of the trigger
	Description() string

	// ShouldRun returns true or false on whether this trigger should be run
	ShouldRun(collector *CivilEventCollector, event *model.CivilEvent) bool

	// Run returns the triggered code
	Run(collector *CivilEventCollector, event *model.CivilEvent) error
}
