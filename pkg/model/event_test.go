// Package model_test contains the tests for the model package
package model_test

import (
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"testing"
	"time"
)

func TestCivilEventPayload(t *testing.T) {

	now := time.Now()
	tsfloat := float64(now.UnixNano()) / float64(1000*1000*1000)
	ts := int(tsfloat)

	event := &model.CivilEvent{}
	event.EventType = "SomeCivilEvent"
	event.Timestamp = ts

	payload := &model.CivilEventPayload{}
	payloadMap := &map[string]*CivilEventPayloadValue{}
}
