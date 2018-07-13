// Package model contains the general data models and interfaces for the Civil crawler.
package model // import "github.com/joincivil/civil-events-crawler/pkg/model"

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/fatih/structs"
)

// NewCivilEventFromContractEvent creates a new civil event after converting eventData to interface{}
func NewCivilEventFromContractEvent(eventType string, contractName string, contractAddress common.Address, eventData interface{},
	timestamp int) (*CivilEvent, error) {
	event := &CivilEvent{}

	payload := NewCivilEventPayload(eventData)

	logPayload, err := extractRawFieldFromEvent(payload)
	if err != nil {
		return event, err
	}
	// convert eventData to map[string]interface{}
	eventPayload, err := extractFieldsFromEvent(payload, eventData, eventType, contractName)
	if err != nil {
		return event, err
	}
	event, err = NewCivilEvent(eventType, contractName, contractAddress, timestamp, eventPayload, logPayload)
	return event, err
}

// NewCivilEvent is a convenience function to create a new CivilEvent
func NewCivilEvent(eventType string, contractName string, contractAddress common.Address, timestamp int,
	eventPayload map[string]interface{}, logPayload *types.Log) (*CivilEvent, error) {
	event := &CivilEvent{}
	event.eventType = eventType
	event.contractName = contractName
	event.contractAddress = contractAddress
	event.eventPayload = eventPayload
	event.logPayload = logPayload
	event.timestamp = timestamp
	event.eventHash = event.hashEvent()
	return event, nil
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

	// contractName is the name of the contract
	contractName string

	// timestamp is the time this event was created.
	timestamp int

	// event payload that doesn't include the "Raw" field
	eventPayload map[string]interface{}

	// "Raw" types.log field from event
	logPayload *types.Log
}

func extractFieldsFromEvent(payload *CivilEventPayload, eventData interface{}, eventType string, contractName string) (map[string]interface{}, error) {
	eventPayload := make(map[string]interface{}, len(payload.data.Fields()))

	_abi, err := AbiJSON(contractName)
	if err != nil {
		return eventPayload, err
	}

	// Trim the eventType clean
	eventType = strings.Trim(eventType, " _")

	// Some contracts have an underscore prefix on their events. Handle both
	// non-underscore/underscore cases here.
	events, ok := _abi.Events[eventType]
	if !ok {
		events, ok = _abi.Events[fmt.Sprintf("_%s", eventType)]
		if !ok {
			return eventPayload, fmt.Errorf("No event type %v in contract %v", eventType, contractName)
		}
	}

	for _, input := range events.Inputs {
		eventFieldName := strings.Title(input.Name)
		eventField, ok := payload.Value(eventFieldName)
		if !ok {
			return eventPayload, errors.New("Can't get event name in civil event")
		}
		switch input.Type.String() {
		case "address":
			addressVal, ok := eventField.Address()
			if !ok {
				return eventPayload, errors.New("Could not convert to common.address type")
			}
			eventPayload[eventFieldName] = addressVal

		case "uint256":
			bigintVal, ok := eventField.BigInt()
			if !ok {
				return eventPayload, errors.New("Could not convert to big.int")
			}
			eventPayload[eventFieldName] = bigintVal
		case "string":
			stringVal, ok := eventField.String()
			if !ok {
				return eventPayload, errors.New("Could not convert to string")
			}
			eventPayload[eventFieldName] = stringVal
		default:
			return eventPayload, fmt.Errorf("unsupported type encountered when parsing %v field for %v event %v",
				input.Type.String(), contractName, eventType)
		}
	}

	return eventPayload, nil
}

// AbiJSON returns parsed abi of this particular contract, maybe this can go in contracttypes.go
func AbiJSON(contractName string) (abi.ABI, error) {
	contractType, ok := NameToContractTypes.GetFromContractName(contractName)
	if !ok {
		return abi.ABI{}, errors.New("Contract Name does not exist")
	}
	contractSpecs, ok := ContractTypeToSpecs.Get(contractType)
	if !ok {
		return abi.ABI{}, errors.New("Invalid contract type")
	}
	_abi, err := abi.JSON(strings.NewReader(contractSpecs.AbiStr()))
	if err != nil {
		return abi.ABI{}, errors.New("Cannot parse abi string")
	}
	return _abi, nil
}

func extractRawFieldFromEvent(payload *CivilEventPayload) (*types.Log, error) {
	rawPayload, ok := payload.Value("Raw")
	if !ok {
		return &types.Log{}, errors.New("Can't get raw value for event")
	}
	logPayload, ok := rawPayload.Log()
	if !ok {
		return &types.Log{}, errors.New("Can't get log field of raw value for event")
	}
	return logPayload, nil
}

// hashEvent returns a hash for event using contractAddress, eventType, and log index
// NOTE: Should we hash more parameters here?
func (e *CivilEvent) hashEvent() string {
	logIndex := int(e.logPayload.Index)
	txHash := e.logPayload.TxHash.Hex()
	eventBytes, _ := rlp.EncodeToBytes([]interface{}{e.contractAddress.Hex(), e.eventType, // nolint: gas
		strconv.Itoa(logIndex), txHash})
	h := crypto.Keccak256Hash(eventBytes)
	return h.Hex()
}

// Hash returns the hash of the CivilEvent
func (e *CivilEvent) Hash() string {
	return e.eventHash
}

// EventType returns the eventType for the CivilEvent
func (e *CivilEvent) EventType() string {
	return e.eventType
}

// ContractAddress returns the contractAddress for the CivilEvent
func (e *CivilEvent) ContractAddress() common.Address {
	return e.contractAddress
}

// Timestamp returns the timestamp for the CivilEvent
func (e *CivilEvent) Timestamp() int {
	return e.timestamp
}

// EventPayload returns the event payload for the CivilEvent
func (e *CivilEvent) EventPayload() map[string]interface{} {
	return e.eventPayload
}

// ContractName returns the contract name
func (e *CivilEvent) ContractName() string {
	return e.contractName
}

// LogPayload returns "Raw" types.log field of civil event
func (e *CivilEvent) LogPayload() *types.Log {
	return e.logPayload
}

// BlockNumber gets the block number from the CivilEvent Payload
func (e *CivilEvent) BlockNumber() uint64 {
	return e.logPayload.BlockNumber
}

// BlockHash gets the block hash from the CivilEvent Payload
func (e *CivilEvent) BlockHash() common.Hash {
	return e.logPayload.BlockHash
}

// LogPayloadToString is a string representation of some fields of log
func (e *CivilEvent) LogPayloadToString() string {
	log := e.logPayload
	return fmt.Sprintf(
		"log: addr: %v, blknum: %v, ind: %v, rem: %v",
		log.Address.Hex(),
		log.BlockNumber,
		log.Index,
		log.Removed,
	)
}

// CivilEventPayload represents the data from a Civil contract event
type CivilEventPayload struct {

	// data is a Struct from the structs package. Just makes it easier
	// to handle access for any kind of event struct.
	data *structs.Struct
}

// NewCivilEventPayload creates a new civil event payload
func NewCivilEventPayload(eventData interface{}) *CivilEventPayload {
	payload := &CivilEventPayload{
		data: structs.New(eventData),
	}
	return payload
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

// ToString returns a string representation for the payload
func (p *CivilEventPayload) ToString() string {
	strs := []string{}
	for _, key := range p.Keys() {
		var str string
		val, _ := p.Value(key)
		if v, ok := val.Address(); ok {
			str = fmt.Sprintf("%v: %v", key, v.Hex())
		} else if v, ok := val.Log(); ok {
			str = fmt.Sprintf(
				"%v: addr: %v, blknum: %v, ind: %v, rem: %v",
				key,
				v.Address.Hex(),
				v.BlockNumber,
				v.Index,
				v.Removed,
			)
		} else if v, ok := val.BigInt(); ok {
			str = fmt.Sprintf("%v: %v", key, v)
		} else if v, ok := val.String(); ok {
			str = fmt.Sprintf("%v: %v", key, v)
		} else if v, ok := val.Int64(); ok {
			str = fmt.Sprintf("%v: %v", key, v)
		}
		strs = append(strs, str)
	}
	return strings.Join(strs, "\n")
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
func (v *CivilEventPayloadValue) Address() (common.Address, bool) {
	val, ok := v.value.Value().(common.Address)
	return val, ok
}

// Log returns the value as types.Log
// Returns bool as false if unable to assert value as type types.Log
func (v *CivilEventPayloadValue) Log() (*types.Log, bool) {
	val, ok := v.value.Value().(types.Log)
	return &val, ok
}
