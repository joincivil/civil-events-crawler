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
func NewCivilEvent(eventType string, contractAddress common.Address, eventData interface{}) *CivilEvent {
	event := &CivilEvent{}
	event.eventType = eventType
	event.contractAddress = contractAddress
	event.timestamp = utils.CurrentEpochSecsInInt()
	event.payload = &CivilEventPayload{
		data: structs.New(eventData),
	}
	event.eventHash = event.hashEvent()
	return event
}

// CivilEvent represents a single Civil smart contract event log item.
// Represents any event type from the sol/abi generated code and creates
// a single type to handle in the listener/retriever.
type CivilEvent struct {

	// eventHash is the hash of event
	eventHash string

	// eventType is the type of event. i.e. _Challenge, _Appeal, _Application.
	eventType string

	// contractAddress of the contract emitting the event
	contractAddress common.Address

	// timestamp is the time this event was created.
	timestamp int

	// payload is the data from the raw event.
	payload *CivilEventPayload
}

// hashEvent returns a hash for event using contractAddress, eventType and timestamp
func (e *CivilEvent) hashEvent() string {
	eventBytes, _ := rlp.EncodeToBytes([]interface{}{e.contractAddress.Hex(), e.eventType,
		strconv.Itoa(e.timestamp)})
	h := crypto.Keccak256Hash(eventBytes)
	return h.Hex()
}

// GetHash returns the hash of the CivilEvent
func (e *CivilEvent) GetHash() string {
	return e.eventHash
}

// GetEventType returns the eventType for the CivilEvent
func (e *CivilEvent) GetEventType() string {
	return e.eventType
}

//GetContractAddress returns the contractAddress for the CivilEvent
func (e *CivilEvent) GetContractAddress() common.Address {
	return e.contractAddress
}

//GetTimestamp returns the timestamp for the CivilEvent
func (e *CivilEvent) GetTimestamp() int {
	return e.timestamp
}

//GetPayload returns the event payload for the CivilEvent
func (e *CivilEvent) GetPayload() *CivilEventPayload {
	return e.payload
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
