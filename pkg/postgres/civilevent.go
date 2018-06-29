package postgres // import "github.com/joincivil/civil-events-crawler/pkg/postgres"

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/generated/eventdef"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"reflect"
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
// TODO: struct tags giving me errors
type CivilEvent struct {
	eventType       string          `db: "event_type"`
	eventHash       string          `db: "hash"`
	contractAddress string          `db: "contract_address"`
	contractName    string          `db: "contract_name"`
	timestamp       int             `db: "int"`
	eventPayload    eventPayloadMap `db: "payload`
	logPayload      eventPayloadMap `db: "log_payload"`
}

// NewCivilEvent constructs a civil event for DB from a model.civilevent
func NewCivilEvent(civilEvent model.CivilEvent) (*CivilEvent, error) {
	dbCivilEvent := &CivilEvent{}
	dbCivilEvent.eventType = civilEvent.EventType()
	dbCivilEvent.eventHash = civilEvent.Hash()
	dbCivilEvent.contractName = civilEvent.ContractName()
	dbCivilEvent.contractAddress = civilEvent.ContractAddress().Hex()
	dbCivilEvent.timestamp = civilEvent.Timestamp()
	err := dbCivilEvent.parseEventPayload(civilEvent)
	if err != nil {
		return nil, err
	}
	return dbCivilEvent, nil
}

// parseEventPayload() parses and converts payloads from civilevent to store in DB:
func (c *CivilEvent) parseEventPayload(civilEvent model.CivilEvent) error {
	_abi, err := civilEvent.AbiJSON()
	if err != nil {
		return err
	}
	err = c.convertEventDataToDB(civilEvent.EventPayload(), _abi)
	if err != nil {
		return err
	}
	c.convertEventLogDataToDB(civilEvent.LogPayload())
	return nil
}

//convertEventToDB() converts event types so they can be stored in the DB
func (c *CivilEvent) convertEventDataToDB(civilEvent map[string]interface{}, _abi abi.ABI) error {
	// loop through abi, and for each val in map, have a way to convert it
	for _, input := range _abi.Events["_"+c.eventType].Inputs {
		eventFieldName := strings.Title(input.Name)
		eventField := civilEvent[eventFieldName]
		switch input.Type.String() {
		case "address":
			c.eventPayload[eventFieldName] = eventField.(common.Address).Hex()
		case "uint256":
			// how to convert big int?
			c.eventPayload[eventFieldName] = eventField
		case "string":
			c.eventPayload[eventFieldName] = eventField.(string)
		case "default":
			return fmt.Errorf("unsupported type")

		}
	}
	return nil
}

// convertDBToEventData converts the db event model to a model.CivilEvent
// NOTE: this only supports conversions for civil event types
func (c *CivilEvent) convertDBToEventData() (*model.CivilEvent, error) {

	
	// // need to create a new instance of the civil event from contracts
	// eventData := contractMapping.Map[c.contractName+c.EventType]
	// for _, input := range _abi.Events[c.eventType].Inputs {
	// 	dbField := c.eventPayload[input.Name]
	// 	switch input.Type.String() {
	// 	// this is the type in the ABI definition
	// 	case "address":
	// 		//do stuff for address conversion here
	// 		model.CivilEventPayloadValue(common.HexToAddress(dbField))
	// 	case "uint256":
	// 		//do stuff for uint256. they convert this to a pointer to bigint

	// 	}
	// }

	// this line needs to be a struct to align with the way constructor
	// creates the map[string]interface{}
	eventData = interface{}
	return &model.NewCivilEvent(c.eventType, c.contractName, eventData, c.timestamp)

}

// convertLogToDB() converts the "Raw"
func (c *CivilEvent) convertEventLogDataToDB(payload *types.Log) {
	c.logPayload["Address"] = payload.Address.Hex()

	topics := make([]string, len(payload.Topics))
	for _, topic := range payload.Topics {
		topics = append(topics, topic.Hex())
	}
	c.logPayload["Topics"] = topics

	c.logPayload["Data"] = common.BytesToHash(payload.Data).Hex()
	c.logPayload["BlockNumber"] = payload.BlockNumber
	c.logPayload["TxHash"] = payload.TxHash.Hex()
	c.logPayload["TxIndex"] = payload.TxIndex
	c.logPayload["BlockHash"] = payload.BlockHash.Hex()
	c.logPayload["Index"] = payload.Index
	c.logPayload["Removed"] = payload.Removed

}

// convertDBToLog() converts the DB payload back to "Raw"
func (c *CivilEvent) convertDBToLog() *types.Log {
	log := &types.Log{}
	log.Address = common.HexToAddress(c.logPayload["Address"].(string))

	logTopics := c.logPayload["Topics"].([]string)
	topics := make([]common.Hash, len(logTopics))
	for _, topic := range logTopics {
		topics = append(topics, common.HexToHash(topic))
	}
	log.Topics = topics

	log.Data = common.HexToHash(c.logPayload["Data"].(string)).Bytes()
	log.BlockNumber = c.logPayload["BlockNumber"].(uint64)
	log.TxHash = common.HexToHash(c.logPayload["TxHash"].(string))
	log.TxIndex = c.logPayload["TxIndex"].(uint)
	log.BlockHash = common.HexToHash(c.logPayload["BlockHash"].(string))
	log.Index = c.logPayload["Index"].(uint)
	log.Removed = c.logPayload["Removed"].(bool)
	return log
}

// //convertDBToEvent() converts queries to DB back to these types
// func (c *CivilEvent) convertDBToEventDataOld(newsroomMapping *eventdef.NewsroomContractEventNameToStruct,
// 	contractMapping *eventdef.CivilTCRContractEventNameToStruct) (*model.CivilEvent, error) {
// 	// reconstruct eventdata here:
// 	var eventData interface{}
// 	switch c.contractName {
// 	case "CivilTCRContract":
// 		eventData = contractMapping.Map[c.contractName+c.eventType]
// 	case "Newsroom":
// 		eventData = newsroomMapping.Map[c.contractName+c.eventType]
// 	}
// 	newEvent := reflect.ValueOf(eventData).Elem()
// 	dbEvent := reflect.ValueOf(c.eventPayload).Elem()

// 	for i := 0; i < newEvent.NumField(); i++ {
// 		f := newEvent.Field(i)
// 		fDB := dbEvent.Field(i)

// 		varName := newEvent.Type().Field(i).Name
// 		varType := newEvent.Type().Field(i).Type

// 		if varName == "Raw" {
// 			f.Set(reflect.ValueOf(c.convertDBToLog()))
// 		}

// 		switch varType.String() {
// 		case "common.Address":
// 			stringVersion := reflect.ValueOf(fDB.Interface()).String()
// 			addressVersion := common.HexToAddress(stringVersion)
// 			f.Set(reflect.ValueOf(addressVersion))
// 			// TODO:
// 		// case "*big.int":
// 		// 	// make this a pointer instead of just bigint
// 		default:
// 			f.Set(reflect.ValueOf(f))
// 		}

// 	}
// 	// if you do this below, the timestamp will be different
// 	return model.NewCivilEvent(c.eventType, c.contractName, common.HexToAddress(c.contractAddress), eventData)
// }
