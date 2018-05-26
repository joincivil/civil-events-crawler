// Package model_test contains the tests for the model package
package model_test

import (
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"testing"
	"time"
)

var testPayloadMap = &map[string]*model.CivilEventPayloadValue{
	"intfield":         &model.CivilEventPayloadValue{Value: 125},
	"intfieldempty":    &model.CivilEventPayloadValue{Value: 0},
	"stringfield":      &model.CivilEventPayloadValue{Value: "this is a str"},
	"stringfieldempty": &model.CivilEventPayloadValue{Value: ""},
	"floatfield":       &model.CivilEventPayloadValue{Value: float64(5.4)},
	"floatfieldempty":  &model.CivilEventPayloadValue{Value: float64(0.0)},
	"boolfieldfalse":   &model.CivilEventPayloadValue{Value: false},
	"boolfieldtrue":    &model.CivilEventPayloadValue{Value: true},
}

func stringInSlice(slice []string, str string) bool {
	for _, b := range slice {
		if b == str {
			return true
		}
	}
	return false
}

func setupCivilEvent() *model.CivilEvent {
	now := time.Now()
	tsfloat := float64(now.UnixNano()) / float64(1000*1000*1000)
	ts := int(tsfloat)

	event := &model.CivilEvent{}
	event.EventType = "SomeCivilEvent"
	event.Timestamp = ts

	payload := &model.CivilEventPayload{}
	payload.PayloadMap = testPayloadMap
	event.Payload = payload
	return event
}

func TestCivilEventSetup(t *testing.T) {
	event := setupCivilEvent()
	if event == nil {
		t.Error("Civil event was not initialized correctly")
	}
}

func TestCivilEventPayloadKeys(t *testing.T) {
	event := setupCivilEvent()
	keys := event.Payload.Keys()
	if keys == nil {
		t.Error("Keys results is empty")
	}
	if len(keys) != len(*testPayloadMap) {
		t.Error("Keys results is correct length")
	}
	for k, _ := range *testPayloadMap {
		if !stringInSlice(keys, k) {
			t.Errorf("%v not found in list of keys", k)
		}
	}
}

func TestCivilEventPayloadValues(t *testing.T) {
	event := setupCivilEvent()
	values := event.Payload.Values()
	if values == nil {
		t.Error("Values results is empty")
	}
	if len(values) != len(*testPayloadMap) {
		t.Error("Values results is correct length")
	}
	for _, v := range values {
		hasEqual := false
		for _, val := range *testPayloadMap {
			if v.Value == val.Value {
				hasEqual = true
				break
			}
		}
		if !hasEqual {
			t.Errorf("%v is not found in list of values", v.Value)
		}
	}
}

func TestCivilEventPayloadValueFunc(t *testing.T) {
	event := setupCivilEvent()
	for k, v := range *testPayloadMap {
		value, exists := event.Payload.Value(k)
		if !exists || value.Value != v.Value {
			t.Errorf("%v, %v is not correct", k, v)
		}
	}
	value, exists := event.Payload.Value("randomkeynotinmap")
	if exists {
		t.Error("Check for nonexistent key should return false for exists")
	}
	if value != nil {
		t.Error("Value should be nil for nonexistent key")
	}
}

func TestCivilEventPayloadValueBool(t *testing.T) {
	event := setupCivilEvent()
	val, exists := event.Payload.Value("boolfieldtrue")
	if !exists {
		t.Error("Bool value does not exist, should exist")
	}
	boolval, ok := val.ToBool()
	if !ok {
		t.Error("Bool type assertion did not work, should have worked")
	}
	if !boolval {
		t.Error("Bool value should be true, not false")
	}
	intval, ok := val.ToInt()
	if ok {
		t.Error("Int type assertion should not have not worked for bool value")
	}
	if intval != 0 {
		t.Error("Int value should be 0 since value is bool")
	}
}

func TestCivilEventPayloadValueInt(t *testing.T) {
	event := setupCivilEvent()
	intFieldName := "intfield"
	fieldVal := (*testPayloadMap)[intFieldName]
	correctVal := fieldVal.Value.(int)

	val, exists := event.Payload.Value(intFieldName)
	if !exists {
		t.Error("Int value does not exist, should exist")
	}
	intval, ok := val.ToInt()
	if !ok {
		t.Error("Int type assertion did not work, should have worked")
	}
	if intval != correctVal {
		t.Error("Int values does not match test data")
	}
	boolval, ok := val.ToBool()
	if ok {
		t.Error("Bool type assertion should not have not worked for int value")
	}
	if boolval {
		t.Error("Bool value should be false since value is int ")
	}
}

func TestCivilEventPayloadValueFloat64(t *testing.T) {
	event := setupCivilEvent()
	floatFieldName := "floatfield"
	fieldVal := (*testPayloadMap)[floatFieldName]
	correctVal := fieldVal.Value.(float64)

	val, exists := event.Payload.Value(floatFieldName)
	if !exists {
		t.Error("Float value does not exist, should exist")
	}
	floatval, ok := val.ToFloat64()
	if !ok {
		t.Error("Float type assertion did not work, should have worked")
	}
	if floatval != correctVal {
		t.Error("Float value does not match test data")
	}
	intval, ok := val.ToInt()
	if ok {
		t.Error("Int type assertion should not have not worked for float value")
	}
	if intval != 0 {
		t.Error("Int value should be 0 since value is float")
	}
}

func TestCivilEventPayloadValueString(t *testing.T) {
	event := setupCivilEvent()
	stringFieldName := "stringfield"
	fieldVal := (*testPayloadMap)[stringFieldName]
	correctVal := fieldVal.Value.(string)

	val, exists := event.Payload.Value(stringFieldName)
	if !exists {
		t.Error("String value does not exist, should exist")
	}
	strval, ok := val.ToString()
	if !ok {
		t.Error("String type assertion did not work, should have worked")
	}
	if strval != correctVal {
		t.Error("String values does not match test data")
	}
	intval, ok := val.ToInt()
	if ok {
		t.Error("Int type assertion should not have not worked for string value")
	}
	if intval != 0 {
		t.Error("Int value should be 0 since value is string")
	}
}
