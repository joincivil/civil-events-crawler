package postgres

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/generated/eventdef"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"reflect"
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
	eventType       string          //`db: "event_type"`
	eventHash       string          //`db: "hash"`
	contractAddress string          //`db: "contract_address"`
	contractName    string          //`db: "contract_name"`
	timestamp       int             //`db: "int"`
	eventPayload    eventPayloadMap //`db: "payload`
	logPayload      eventPayloadMap //`db: "log_payload"`
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
	c.processEventData(civilEvent)
	c.processEventLog(civilEvent.Payload())
	return nil
}

//convertEventToDB() converts event types so they can be stored in the DB
func (c *CivilEvent) processEventData(civilEvent model.CivilEvent) {
	rawPayload := civilEvent.RawPayload()
	payload := civilEvent.Payload()
	val := reflect.ValueOf(rawPayload).Elem()
	for j := 0; j < val.NumField(); j++ {
		name := val.Type().Field(j).Name
		if name != "Raw" {
			eventValue, _ := payload.Value(name)
			switch val.Field(j).Type().String() {
			case "common.Address":
				address, _ := eventValue.Address()
				c.eventPayload[name] = address.Hex()
			case "*big.int":
				num, _ := eventValue.BigInt()
				c.eventPayload[name] = num
			default:
				c.eventPayload[name] = eventValue
			}
		}
	}
}

//convertDBToEvent() converts queries to DB back to these types
func (c *CivilEvent) convertDBToEvent(newsroomMapping *eventdef.NewsroomContractEventNameToStruct,
	contractMapping *eventdef.CivilTCRContractEventNameToStruct) (*model.CivilEvent, error) {
	// reconstruct eventdata here:
	var eventData interface{}
	switch c.contractName {
	case "CivilTCRContract":
		eventData = contractMapping.Map[c.contractName+c.eventType]
	case "Newsroom":
		eventData = newsroomMapping.Map[c.contractName+c.eventType]
	}
	newEvent := reflect.ValueOf(eventData).Elem()
	dbEvent := reflect.ValueOf(c.eventPayload).Elem()

	for i := 0; i < newEvent.NumField(); i++ {
		f := newEvent.Field(i)
		fDB := dbEvent.Field(i)

		varName := newEvent.Type().Field(i).Name
		varType := newEvent.Type().Field(i).Type

		if varName == "Raw" {
			f.Set(reflect.ValueOf(c.convertDBToLog()))
		}

		switch varType.String() {
		case "common.Address":
			stringVersion := reflect.ValueOf(fDB.Interface()).String()
			addressVersion := common.HexToAddress(stringVersion)
			f.Set(reflect.ValueOf(addressVersion))
			// TODO:
		// case "*big.int":
		// 	// make this a pointer instead of just bigint
		default:
			f.Set(reflect.ValueOf(f))
		}

	}

	return model.NewCivilEvent(c.eventType, c.contractName, common.HexToAddress(c.contractAddress), eventData)
}

// convertLogToDB() converts the "Raw"
func (c *CivilEvent) processEventLog(payload *model.CivilEventPayload) {
	rawPayload, _ := payload.Value("Raw")
	rawLogPayload, _ := rawPayload.Log()
	c.logPayload["Address"] = rawLogPayload.Address.Hex()

	topics := make([]string, len(rawLogPayload.Topics))
	for _, topic := range rawLogPayload.Topics {
		topics = append(topics, topic.Hex())
	}
	c.logPayload["Topics"] = topics

	c.logPayload["Data"] = common.BytesToHash(rawLogPayload.Data).Hex()
	c.logPayload["BlockNumber"] = rawLogPayload.BlockNumber
	c.logPayload["TxHash"] = rawLogPayload.TxHash.Hex()
	c.logPayload["TxIndex"] = rawLogPayload.TxIndex
	c.logPayload["BlockHash"] = rawLogPayload.BlockHash.Hex()
	c.logPayload["Index"] = rawLogPayload.Index
	c.logPayload["Removed"] = rawLogPayload.Removed

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
