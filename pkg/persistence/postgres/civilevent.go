package postgres // import "github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"math/big"
	"strings"
)

// EventTableSchema returns the query to create this table
func EventTableSchema() string {
	schema := `
		CREATE TABLE IF NOT EXISTS event(
			id SERIAL PRIMARY KEY,
			event_type TEXT,
			hash TEXT UNIQUE,
			contract_address TEXT,
			contract_name TEXT,
			timestamp BIGINT,
			payload JSONB,
			log_payload JSONB
		);
	`
	return schema
}

// EventTableIndices returns the query to create indices for this table
func EventTableIndices() string {
	indexCreationQuery := `
		CREATE INDEX IF NOT EXISTS event_event_type_idx ON event (event_type);
		CREATE INDEX IF NOT EXISTS event_contract_address_idx ON event (contract_address);
		CREATE INDEX IF NOT EXISTS event_timestamp_idx ON event (timestamp);
	`
	return indexCreationQuery
}

// CivilEvent is the model for event table in DB
type CivilEvent struct {
	EventType       string       `db:"event_type"`
	EventHash       string       `db:"hash"`
	ContractAddress string       `db:"contract_address"`
	ContractName    string       `db:"contract_name"`
	Timestamp       int          `db:"timestamp"`
	EventPayload    JsonbPayload `db:"payload"`
	LogPayload      JsonbPayload `db:"log_payload"`
}

// NewCivilEvent constructs a civil event for DB from a model.civilevent
// Rename this to NewDBEventFromCivilEvent
func NewCivilEvent(civilEvent *model.CivilEvent) (*CivilEvent, error) {
	dbCivilEvent := &CivilEvent{}
	dbCivilEvent.EventType = civilEvent.EventType()
	dbCivilEvent.EventHash = civilEvent.Hash()
	dbCivilEvent.ContractName = civilEvent.ContractName()
	dbCivilEvent.ContractAddress = civilEvent.ContractAddress().Hex()
	dbCivilEvent.Timestamp = civilEvent.Timestamp()
	dbCivilEvent.EventPayload = make(JsonbPayload)
	dbCivilEvent.LogPayload = make(JsonbPayload)
	err := dbCivilEvent.parseEventPayload(civilEvent)
	if err != nil {
		return nil, err
	}
	return dbCivilEvent, nil
}

// parseEventPayload() parses and converts payloads from civilevent to store in DB:
func (c *CivilEvent) parseEventPayload(civilEvent *model.CivilEvent) error {
	err := c.EventDataToDB(civilEvent.EventPayload())
	if err != nil {
		return err
	}
	c.EventLogDataToDB(civilEvent.LogPayload())
	return nil
}

// EventDataToDB converts event types so they can be stored in the DB
func (c *CivilEvent) EventDataToDB(civilEvent map[string]interface{}) error {
	abi, err := model.AbiJSON(c.ContractName)
	if err != nil {
		return fmt.Errorf("Error getting abi from contract name: %v", err)
	}
	events, err := ReturnEventsFromABI(abi, c.EventType)
	if err != nil {
		return err
	}
	eventPayload := make(JsonbPayload)

	for _, input := range events.Inputs {
		eventFieldName := strings.Title(input.Name)
		eventField := civilEvent[eventFieldName]
		switch input.Type.String() {
		case "address":
			eventPayload[eventFieldName] = eventField.(common.Address).Hex()
		case "uint256":
			// NOTE: converting all *big.Int to int64. assuming for now that numbers will fall into int64 range.
			eventPayload[eventFieldName] = eventField.(*big.Int).Int64()
		case "string":
			eventPayload[eventFieldName] = eventField.(string)
		case "default":
			return fmt.Errorf("unsupported type")

		}
	}
	c.EventPayload = eventPayload
	return nil
}

// DBToEventData converts the db event model to a model.CivilEvent
// NOTE: because this is stored in DB as a map[string]interface{}, Postgres converts some fields, see notes in function.
func (c *CivilEvent) DBToEventData() (*model.CivilEvent, error) {
	civilEvent := &model.CivilEvent{}
	abi, err := model.AbiJSON(c.ContractName)
	if err != nil {
		return civilEvent, fmt.Errorf("Error getting abi from contract name: %v", err)
	}
	eventPayload := make(map[string]interface{})
	events, err := ReturnEventsFromABI(abi, c.EventType)
	if err != nil {
		return civilEvent, err
	}

	for _, input := range events.Inputs {
		eventFieldName := strings.Title(input.Name)
		eventField, ok := c.EventPayload[eventFieldName]
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
			// NOTE: Ints are stored in DB as float64
			num, numOk := eventField.(float64)
			if !numOk {
				return civilEvent, errors.New("Cannot cast DB int to float64")
			}
			eventPayload[eventFieldName] = big.NewInt(int64(num))
		case "string":
			str, stringOk := eventField.(string)
			if !stringOk {
				return civilEvent, errors.New("Cannot cast DB string val to string")
			}
			eventPayload[eventFieldName] = str
		default:
			return civilEvent, fmt.Errorf("unsupported type in %v field encountered in %v event", eventFieldName, c.EventHash)
		}
	}
	logPayload := c.DBToEventLogData()
	contractAddress := common.HexToAddress(c.ContractAddress)
	civilEvent, err = model.NewCivilEvent(c.EventType, c.ContractName, contractAddress, c.Timestamp, eventPayload,
		logPayload)

	return civilEvent, err

}

// EventLogDataToDB converts the raw log data to Postgresql types
func (c *CivilEvent) EventLogDataToDB(payload *types.Log) {
	c.LogPayload["Address"] = payload.Address.Hex()

	topics := make([]string, len(payload.Topics))
	for _, topic := range payload.Topics {
		topics = append(topics, topic.Hex())
	}
	c.LogPayload["Topics"] = topics

	c.LogPayload["Data"] = payload.Data

	c.LogPayload["BlockNumber"] = payload.BlockNumber
	c.LogPayload["TxHash"] = payload.TxHash.Hex()
	c.LogPayload["TxIndex"] = payload.TxIndex
	c.LogPayload["BlockHash"] = payload.BlockHash.Hex()
	c.LogPayload["Index"] = payload.Index
	c.LogPayload["Removed"] = payload.Removed

}

// DBToEventLogData converts the DB raw log payload back to types.Log
// NOTE: because this is stored in DB as a map[string]interface{}, Postgres converts some fields, see notes in function.
func (c *CivilEvent) DBToEventLogData() *types.Log {
	log := &types.Log{}
	log.Address = common.HexToAddress(c.LogPayload["Address"].(string))

	topics := c.LogPayload["Topics"].([]interface{})
	newTopics := make([]common.Hash, len(topics))
	for i, topic := range topics {
		topicString := topic.(string)
		newTopics[i] = common.HexToHash(topicString)
	}
	log.Topics = newTopics

	// NOTE: Data is stored in DB as a string
	log.Data = []byte(c.LogPayload["Data"].(string))

	// NOTE: BlockNumber is stored in DB as float64
	log.BlockNumber = uint64(c.LogPayload["BlockNumber"].(float64))

	log.TxHash = common.HexToHash(c.LogPayload["TxHash"].(string))

	// NOTE: TxIndex is stored in DB as float64
	log.TxIndex = uint(c.LogPayload["TxIndex"].(float64))

	log.BlockHash = common.HexToHash(c.LogPayload["BlockHash"].(string))

	// NOTE: Index is stored in DB as float64
	log.Index = uint(c.LogPayload["Index"].(float64))

	log.Removed = c.LogPayload["Removed"].(bool)
	return log
}
