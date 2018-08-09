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

const (
	eventTableName = "event"
)

// CreateEventTableQuery returns the query to create the event table
func CreateEventTableQuery() string {
	return CreateEventTableQueryString(eventTableName)
}

// CreateEventTableQueryString returns the query to create this table
func CreateEventTableQueryString(tableName string) string {
	queryString := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s(
            id SERIAL PRIMARY KEY,
            event_type TEXT,
            hash TEXT UNIQUE,
            contract_address TEXT,
            contract_name TEXT,
            timestamp BIGINT,
            retrieval_method SMALLINT,
            payload JSONB,
            log_payload JSONB
        );
    `, tableName)
	return queryString
}

// EventTableIndices returns the query to create indices for this table
func EventTableIndices() string {
	return CreateEventTableIndicesString(eventTableName)
}

// CreateEventTableIndicesString returns the query to create this table
func CreateEventTableIndicesString(tableName string) string {
	queryString := fmt.Sprintf(`
		CREATE INDEX IF NOT EXISTS event_event_type_idx ON %s (event_type);
		CREATE INDEX IF NOT EXISTS event_contract_address_idx ON %s (contract_address);
		CREATE INDEX IF NOT EXISTS event_timestamp_idx ON %s (timestamp);
	`, tableName, tableName, tableName)
	return queryString
}

// Event is the model for events table in DB
type Event struct {
	EventType       string       `db:"event_type"`
	EventHash       string       `db:"hash"`
	ContractAddress string       `db:"contract_address"`
	ContractName    string       `db:"contract_name"`
	Timestamp       int64        `db:"timestamp"`
	RetrievalMethod int          `db:"retrieval_method"`
	EventPayload    JsonbPayload `db:"payload"`
	LogPayload      JsonbPayload `db:"log_payload"`
}

// NewDbEventFromEvent constructs an event for DB from a model.event
func NewDbEventFromEvent(event *model.Event) (*Event, error) {
	dbEvent := &Event{}
	dbEvent.EventType = event.EventType()
	dbEvent.EventHash = event.Hash()
	dbEvent.ContractName = event.ContractName()
	dbEvent.ContractAddress = event.ContractAddress().Hex()
	dbEvent.Timestamp = event.Timestamp()
	dbEvent.RetrievalMethod = int(event.RetrievalMethod())
	dbEvent.EventPayload = make(JsonbPayload)
	dbEvent.LogPayload = make(JsonbPayload)
	err := dbEvent.parseEventPayload(event)
	if err != nil {
		return nil, err
	}
	return dbEvent, nil
}

// EventDataToDB converts event types so they can be stored in the DB
func (c *Event) EventDataToDB(event map[string]interface{}) error {
	abi, err := model.AbiJSON(c.ContractName)
	if err != nil {
		return fmt.Errorf("Error getting ABI from contract name: %v", err)
	}
	events, err := model.ReturnEventsFromABI(abi, c.EventType)
	if err != nil {
		return fmt.Errorf("Error parsing ABI to get events, err: %v", err)
	}
	eventPayload := make(JsonbPayload)

	for _, input := range events.Inputs {
		eventFieldName := strings.Title(input.Name)
		eventField := event[eventFieldName]
		switch input.Type.String() {
		case "address":
			eventPayload[eventFieldName] = eventField.(common.Address).Hex()
		case "uint256":
			// NOTE(IS): converting all *big.Int to int64. assuming for now that numbers will fall into int64 range.
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

// DBToEventData converts the db event model to a model.
// NOTE(IS): because jsonb payloads are stored in DB as a map[string]interface{}, Postgres converts some fields, see notes in function.
func (c *Event) DBToEventData() (*model.Event, error) {
	event := &model.Event{}
	abi, err := model.AbiJSON(c.ContractName)
	if err != nil {
		return event, fmt.Errorf("Error getting abi from contract name: %v", err)
	}
	eventPayload := make(map[string]interface{})
	events, err := model.ReturnEventsFromABI(abi, c.EventType)
	if err != nil {
		return event, err
	}

	for _, input := range events.Inputs {
		eventFieldName := strings.Title(input.Name)
		eventField, ok := c.EventPayload[eventFieldName]
		if !ok {
			return event, fmt.Errorf("Cannot get %v field of DB Event", eventFieldName)
		}
		switch input.Type.String() {
		case "address":
			address, addressOk := eventField.(string)
			if !addressOk {
				return event, errors.New("Cannot cast DB contract address to string")
			}
			eventPayload[eventFieldName] = common.HexToAddress(address)
		case "uint256":
			// NOTE: Ints are stored in DB as float64
			num, numOk := eventField.(float64)
			if !numOk {
				return event, errors.New("Cannot cast DB int to float64")
			}
			eventPayload[eventFieldName] = big.NewInt(int64(num))
		case "string":
			str, stringOk := eventField.(string)
			if !stringOk {
				return event, errors.New("Cannot cast DB string val to string")
			}
			eventPayload[eventFieldName] = str
		default:
			return event, fmt.Errorf("unsupported type in %v field encountered in %v event", eventFieldName, c.EventHash)
		}
	}
	logPayload := c.DBToEventLogData()
	contractAddress := common.HexToAddress(c.ContractAddress)
	event, err = model.NewEvent(c.EventType, c.ContractName, contractAddress, c.Timestamp, model.RetrievalMethod(c.RetrievalMethod), eventPayload,
		logPayload)

	return event, err

}

// EventLogDataToDB converts the raw log data to Postgresql types
func (c *Event) EventLogDataToDB(payload *types.Log) {
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
// NOTE(IS): because jsonb payloads are stored in DB as a map[string]interface{}, Postgres converts some fields, see notes in function.
func (c *Event) DBToEventLogData() *types.Log {
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

// parseEventPayload() parses and converts payloads from event to store in DB
func (c *Event) parseEventPayload(event *model.Event) error {
	err := c.EventDataToDB(event.EventPayload())
	if err != nil {
		return err
	}
	c.EventLogDataToDB(event.LogPayload())
	return nil
}
