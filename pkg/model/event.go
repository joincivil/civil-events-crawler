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

	// The payload data of the event.
	Payload *CivilEventPayload
}

// CivilEventPayload represents the payload of the event.  This is derived
// from the data in the generated event type from the sol/abi code.
type CivilEventPayload struct {

	// Internal map for field to CivilPayloadValue
	payload map[string]*CivilEventPayloadValue
}

// KeysValues returns the map of keys to values
func (p *CivilEventPayload) KeysValues() map[string]*CivilEventPayloadValue {
	return p.payload
}

// Keys returns all the field/key names in the payload
func (p *CivilEventPayload) Keys() []string {
	keys := make([]string, len(p.payload))
	index := 0
	for k := range p.payload {
		keys[index] = k
		index++
	}
	return keys
}

// Values returns all the values in the payload
func (p *CivilEventPayload) Values() []*CivilEventPayloadValue {
	values := make([]*CivilEventPayloadValue, len(p.payload))
	index := 0
	for _, v := range p.payload {
		values[index] = v
		index++
	}
	return values
}

// Value returns the value of the key.
func (p *CivilEventPayload) Value(key string) (val *CivilEventPayloadValue, exists bool) {
	val, exists = p.payload[key]
	return val, exists
}

// CivilEventPayloadValue represents a value for a payload key.
type CivilEventPayloadValue struct {
	value interface{}
}

// Value returns the raw payload value.
func (v *CivilEventPayloadValue) Value() interface{} {
	return v.value
}

// ToBool attempts to return the value as a bool
func (v *CivilEventPayloadValue) ToBool() (val bool, ok bool) {
	val, ok = v.value.(bool)
	return val, ok
}

// ToInt attempts to return the value as an int
func (v *CivilEventPayloadValue) ToInt() (val int, ok bool) {
	val, ok = v.value.(int)
	return val, ok
}

// ToFloat64 attempts to return the value as an float64
func (v *CivilEventPayloadValue) ToFloat64() (val float64, ok bool) {
	val, ok = v.value.(float64)
	return val, ok
}

// ToString attempts to return the value as an string
func (v *CivilEventPayloadValue) ToString() (val string, ok bool) {
	val, ok = v.value.(string)
	return val, ok
}
