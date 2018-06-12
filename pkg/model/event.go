// Package model contains the general data models and interfaces for the Civil crawler.
package model // import "github.com/joincivil/civil-events-crawler/pkg/model"

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/fatih/structs"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"math/big"
	"reflect"
	"strconv"
)

// NewCivilEvent is a convenience function to create a new CivilEvent
// NOTE (IS): I think fields should be static if we want a hash to be totally proper
func NewCivilEvent(eventType string, contractAddress common.Address, eventData interface{}) *CivilEvent {
	event := &CivilEvent{}
	event.EventType = eventType
	event.ContractAddress = contractAddress
	event.Timestamp = utils.CurrentEpochSecsInInt()
	event.Payload = &CivilEventPayload{
		data: structs.New(eventData),
	}
	// Assuming here that once a CivilEvent is created, fields will not be changed
	event.EventHash = event.HashEvent()
	return event
}

// CivilEvent represents a single Civil smart contract event log item.
// Represents any event type from the sol/abi generated code and creates
// a single type to handle in the listener/retriever.
type CivilEvent struct {

	//EventHash is the hash of event
	EventHash string

	// EventType is the type of event. i.e. _Challenge, _Appeal, _Application.
	EventType string

	// Address of the contract emitting the event
	ContractAddress common.Address

	// Timestamp is the time this event was created.
	Timestamp int

	// Payload is the data from the raw event.
	Payload *CivilEventPayload
}

// HashEvent creates a hash for event based on data
// Hash is of ContractAddress, EventType and Timestamp
// Not good practice to save this as a mutable field bc the hash could be
// changed...
func (e *CivilEvent) HashEvent() string {
	eventBytes, _ := rlp.EncodeToBytes([]interface{}{e.ContractAddress.Hex(), e.EventType,
		strconv.Itoa(e.Timestamp)})
	h := crypto.Keccak256Hash(eventBytes)
	return h.Hex()
}

// CivilEventPayload represents the data from a Civil contract event
type CivilEventPayload struct {

	// data is a Struct from the structs package. Just makes it easier
	// to handle access for any kind of event struct.
	data *structs.Struct
}

// Keys retrieves all the available key names in the event payload
func (p *CivilEventPayload) Keys() []string {
	keyFields := p.data.Fields()
	keys := make([]string, len(keyFields))
	for ind, field := range keyFields {
		keys[ind] = field.Name()
	}
	return keys
}

// Value returns the CivilEventPayloadValue of the given key
func (p *CivilEventPayload) Value(key string) (*CivilEventPayloadValue, bool) {
	field, ok := p.data.FieldOk(key)
	if !ok {
		return nil, ok
	}
	return &CivilEventPayloadValue{value: field}, ok
}

// CivilEventPayloadValue represents a single value for a key in the payload
type CivilEventPayloadValue struct {
	value *structs.Field
}

// Kind returns the value's basic type as described with reflect.Kind
func (v *CivilEventPayloadValue) Kind() reflect.Kind {
	return v.value.Kind()
}

// Val returns the value as an unknown type interface{}
func (v *CivilEventPayloadValue) Val() interface{} {
	return v.value.Value()
}

// String returns the value as a string
// Returns bool as false if unable to assert value as type string
func (v *CivilEventPayloadValue) String() (string, bool) {
	val, ok := v.value.Value().(string)
	return val, ok
}

// Int64 returns the value as a int64.
// Returns bool as false if unable to assert value as type int64
func (v *CivilEventPayloadValue) Int64() (int64, bool) {
	val, ok := v.BigInt()
	if !ok {
		return 0, ok
	}
	return val.Int64(), ok
}

// BigInt returns the value as a big.Int
// Returns bool as false if unable to assert value as type big.Int
func (v *CivilEventPayloadValue) BigInt() (*big.Int, bool) {
	val, ok := v.value.Value().(*big.Int)
	return val, ok
}

// Address returns the value as common.Address
// Returns bool as false if unable to assert value as type common.Address
func (v *CivilEventPayloadValue) Address() (*common.Address, bool) {
	val, ok := v.value.Value().(common.Address)
	return &val, ok
}

// Log returns the value as types.Log
// Returns bool as false if unable to assert value as type types.Log
func (v *CivilEventPayloadValue) Log() (*types.Log, bool) {
	val, ok := v.value.Value().(types.Log)
	return &val, ok
}
