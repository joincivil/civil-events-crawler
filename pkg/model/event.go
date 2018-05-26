// Package model contains the general data models for the Civil crawler.
package model // import "github.com/joincivil/civil-events-crawler/pkg/model"

// CivilEvent represents a single Civil smart contract event log item.
// Represents any event type from the sol/abi generated code and creates
// a single type to handle.
type CivilEvent struct {

	// EventType is the type of event. i.e. challenges, appeals, etc
	EventType string

	// Timestamp is the time this event was created.
	Timestamp int

	// Payload is the data from the raw event.
	Payload *CivilEventPayload
}

// CivilEventPayload represents the payload of the event.  This is derived
// from the data in the generated event type from the sol/abi code.
type CivilEventPayload struct {

	// PayloadMap stores the field to CivilPayloadValues
	PayloadMap *map[string]*CivilEventPayloadValue
}

// Keys returns all the field/key names in the payload
func (p *CivilEventPayload) Keys() []string {
	keys := make([]string, len(*p.PayloadMap))
	index := 0
	for k := range *p.PayloadMap {
		keys[index] = k
		index++
	}
	return keys
}

// Values returns all the values in the payload
func (p *CivilEventPayload) Values() []*CivilEventPayloadValue {
	values := make([]*CivilEventPayloadValue, len(*p.PayloadMap))
	index := 0
	for _, v := range *p.PayloadMap {
		values[index] = v
		index++
	}
	return values
}

// Value returns the value of the key.
func (p *CivilEventPayload) Value(key string) (val *CivilEventPayloadValue, exists bool) {
	val, exists = (*p.PayloadMap)[key]
	return val, exists
}

// CivilEventPayloadValue represents a value for a payload key.
type CivilEventPayloadValue struct {
	Value interface{}
}

// ToBool attempts to return the value as a bool
func (v *CivilEventPayloadValue) ToBool() (val bool, ok bool) {
	val, ok = v.Value.(bool)
	return val, ok
}

// ToInt attempts to return the value as an int
func (v *CivilEventPayloadValue) ToInt() (val int, ok bool) {
	val, ok = v.Value.(int)
	return val, ok
}

// ToFloat64 attempts to return the value as an float64
func (v *CivilEventPayloadValue) ToFloat64() (val float64, ok bool) {
	val, ok = v.Value.(float64)
	return val, ok
}

// ToString attempts to return the value as an string
func (v *CivilEventPayloadValue) ToString() (val string, ok bool) {
	val, ok = v.Value.(string)
	return val, ok
}
