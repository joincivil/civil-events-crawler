package postgres // import "github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"math/big"
	"strings"
)

// eventPayloadMap is the jsonb payload
type eventPayloadMap map[string]interface{}

func (ep *eventPayloadMap) Value() (driver.Value, error) {
	return json.Marshal(ep)
}

func (ep *eventPayloadMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}
	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*ep, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}

// CivilEvent is the model for events table in DB
// TODO(IS): struct tag giving errors
type CivilEvent struct {
	eventType       string          // `db: "event_type"`
	eventHash       string          // `db: "hash"`
	contractAddress string          // `db: "contract_address"`
	contractName    string          // `db: "contract_name"`
	timestamp       int             // `db: "int"`
	eventPayload    eventPayloadMap // `db: "payload`
	logPayload      eventPayloadMap // `db: "log_payload"`
}

// NewCivilEvent constructs a civil event for DB from a model.civilevent
func NewCivilEvent(civilEvent *model.CivilEvent) (*CivilEvent, error) {
	dbCivilEvent := &CivilEvent{}
	dbCivilEvent.eventType = civilEvent.EventType()
	dbCivilEvent.eventHash = civilEvent.Hash()
	dbCivilEvent.contractName = civilEvent.ContractName()
	dbCivilEvent.contractAddress = civilEvent.ContractAddress().Hex()
	dbCivilEvent.timestamp = civilEvent.Timestamp()
	dbCivilEvent.eventPayload = make(map[string]interface{})
	dbCivilEvent.logPayload = make(map[string]interface{})
	err := dbCivilEvent.parseEventPayload(civilEvent)
	if err != nil {
		return nil, err
	}
	return dbCivilEvent, nil
}

// parseEventPayload() parses and converts payloads from civilevent to store in DB:
func (c *CivilEvent) parseEventPayload(civilEvent *model.CivilEvent) error {
	_abi, err := model.AbiJSON(c.contractName)
	if err != nil {
		return err
	}
	err = c.EventDataToDB(civilEvent.EventPayload(), _abi)
	if err != nil {
		return err
	}
	c.EventLogDataToDB(civilEvent.LogPayload())
	return nil
}

//EventDataToDB converts event types so they can be stored in the DB
func (c *CivilEvent) EventDataToDB(civilEvent map[string]interface{}, _abi abi.ABI) error {
	// loop through abi, and for each val in map, have a way to convert it
	for _, input := range _abi.Events["_"+c.eventType].Inputs {
		eventFieldName := strings.Title(input.Name)
		eventField := civilEvent[eventFieldName]
		switch input.Type.String() {
		case "address":
			c.eventPayload[eventFieldName] = eventField.(common.Address).Hex()
		case "uint256":
			// NOTE: converting all *big.Int to int64. assuming for now that numbers will fall into int64 range.
			c.eventPayload[eventFieldName] = eventField.(*big.Int).Int64()
		case "string":
			c.eventPayload[eventFieldName] = eventField.(string)
		case "default":
			return fmt.Errorf("unsupported type")

		}
	}
	return nil
}

// DBToEventData converts the db event model to a model.CivilEvent
func (c *CivilEvent) DBToEventData() (*model.CivilEvent, error) {
	civilEvent := &model.CivilEvent{}
	_abi, err := model.AbiJSON(c.contractName)
	if err != nil {
		return civilEvent, err
	}
	eventPayload := make(map[string]interface{})

	for _, input := range _abi.Events["_"+c.eventType].Inputs {
		eventFieldName := strings.Title(input.Name)
		eventField, ok := c.eventPayload[eventFieldName]
		if !ok {
			return civilEvent, fmt.Errorf("Cannot get %v field of DB CivilEvent", eventFieldName)
		}
		switch input.Type.String() {
		case "address":
			address, addressOk := eventField.(string)
			if !addressOk {
				return civilEvent, errors.New("Cannot cast DB contract address to string")
			}
			eventPayload[eventFieldName] = common.HexToAddress(address)
		case "uint256":
			num, numOk := eventField.(int64)
			if !numOk {
				return civilEvent, errors.New("Cannot cast DB int to int64")
			}
			eventPayload[eventFieldName] = big.NewInt(num)
		case "string":
			str, stringOk := eventField.(string)
			if !stringOk {
				return civilEvent, errors.New("Cannot cast DB string val to string")
			}
			eventPayload[eventFieldName] = str
		default:
			return civilEvent, fmt.Errorf("unsupported type in %v field encountered in %v event", eventFieldName, c.eventHash)
		}
	}

	logPayload := c.DBToLog()
	contractAddress := common.HexToAddress(c.contractAddress)
	civilEvent, err = model.NewCivilEvent(c.eventType, c.contractName, contractAddress, c.timestamp, eventPayload,
		logPayload)

	return civilEvent, err

}

// EventLogDataToDB converts the "Raw"
func (c *CivilEvent) EventLogDataToDB(payload *types.Log) {
	c.logPayload["Address"] = payload.Address.Hex()

	topics := make([]string, len(payload.Topics))
	for _, topic := range payload.Topics {
		topics = append(topics, topic.Hex())
	}
	c.logPayload["Topics"] = topics

	c.logPayload["Data"] = payload.Data //common.BytesToHash(payload.Data).Hex()
	c.logPayload["BlockNumber"] = payload.BlockNumber
	c.logPayload["TxHash"] = payload.TxHash.Hex()
	c.logPayload["TxIndex"] = payload.TxIndex
	c.logPayload["BlockHash"] = payload.BlockHash.Hex()
	c.logPayload["Index"] = payload.Index
	c.logPayload["Removed"] = payload.Removed

}

// DBToLog converts the DB payload back to "Raw"
func (c *CivilEvent) DBToLog() *types.Log {
	log := &types.Log{}
	log.Address = common.HexToAddress(c.logPayload["Address"].(string))

	logTopics := c.logPayload["Topics"].([]string)
	topics := make([]common.Hash, len(logTopics))
	for _, topic := range logTopics {
		topics = append(topics, common.HexToHash(topic))
	}
	log.Topics = topics

	log.Data = c.logPayload["Data"].([]byte) // common.HexToHash(c.logPayload["Data"].(string)).Bytes()

	log.BlockNumber = c.logPayload["BlockNumber"].(uint64)
	log.TxHash = common.HexToHash(c.logPayload["TxHash"].(string))
	log.TxIndex = c.logPayload["TxIndex"].(uint)
	log.BlockHash = common.HexToHash(c.logPayload["BlockHash"].(string))
	log.Index = c.logPayload["Index"].(uint)
	log.Removed = c.logPayload["Removed"].(bool)
	return log
}

// EventType returns the event type of the Postgres CivilEvent
func (c *CivilEvent) EventType() string {
	return c.eventType
}

// EventHash returns the event hash of the Postgres CivilEvent
func (c *CivilEvent) EventHash() string {
	return c.eventHash
}

// ContractAddress returns the contract address of the Postgres CivilEvent
func (c *CivilEvent) ContractAddress() string {
	return c.contractAddress
}

// ContractName returns the contract name of the Postgres CivilEvent
func (c *CivilEvent) ContractName() string {
	return c.contractName
}

// Timestamp returns the timestamp of the Postgres CivilEvent
func (c *CivilEvent) Timestamp() int {
	return c.timestamp
}

// EventPayload returns the event payload of the Postgres CivilEvent
func (c *CivilEvent) EventPayload() map[string]interface{} {
	return c.eventPayload
}

// LogPayload returns the event payload of the Postgres CivilEvent
func (c *CivilEvent) LogPayload() map[string]interface{} {
	return c.logPayload
}
