package postgres // import "github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/model"

	cpostgres "github.com/joincivil/go-common/pkg/persistence/postgres"
)

const (
	// EventTableBaseName is the type of table this code defines
	EventTableBaseName = "event"
)

// CreateEventTableQuery returns the query to create this table
func CreateEventTableQuery(tableName string) string {
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

// CreateEventTableIndices returns the query to create this table
func CreateEventTableIndices(tableName string) string {
	queryString := fmt.Sprintf(`
		CREATE INDEX IF NOT EXISTS event_event_type_idx ON %s (event_type);
		CREATE INDEX IF NOT EXISTS event_contract_address_idx ON %s (contract_address);
		CREATE INDEX IF NOT EXISTS event_timestamp_idx ON %s (timestamp);
	`, tableName, tableName, tableName)
	return queryString
}

// Event is the model for events table in DB
type Event struct {
	EventType       string                 `db:"event_type"`
	EventHash       string                 `db:"hash"`
	ContractAddress string                 `db:"contract_address"`
	ContractName    string                 `db:"contract_name"`
	Timestamp       int64                  `db:"timestamp"`
	RetrievalMethod int                    `db:"retrieval_method"`
	EventPayload    cpostgres.JsonbPayload `db:"payload"`
	LogPayload      cpostgres.JsonbPayload `db:"log_payload"`
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
	dbEvent.EventPayload = make(cpostgres.JsonbPayload)
	dbEvent.LogPayload = make(cpostgres.JsonbPayload)
	err := dbEvent.parseEventPayload(event)
	if err != nil {
		return nil, err
	}
	return dbEvent, nil
}

// EventDataToDB converts event data payload types so they can be stored in the DB
func (c *Event) EventDataToDB(eventData map[string]interface{}) error {
	abi, err := model.AbiJSON(c.ContractName)
	if err != nil {
		return fmt.Errorf("Error getting ABI from contract name: %v", err)
	}
	abiEvent, err := model.ReturnEventFromABI(abi, c.EventType)
	if err != nil {
		return fmt.Errorf("Error parsing ABI to get events, err: %v", err)
	}
	eventPayload := make(cpostgres.JsonbPayload)

	for _, input := range abiEvent.Inputs {
		eventFieldName := strings.Title(input.Name)
		eventField := eventData[eventFieldName]
		switch input.Type.String() {
		case "address":
			eventPayload[eventFieldName] = eventField.(common.Address).Hex()
		case "uint256":
			// NOTE(IS): Store all values as float64 to avoid overflow for int64
			i := eventField.(*big.Int)
			f := new(big.Float).SetInt(i)
			val, _ := f.Float64()
			eventPayload[eventFieldName] = val
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
	abiEvent, err := model.ReturnEventFromABI(abi, c.EventType)
	if err != nil {
		return event, err
	}

	for _, input := range abiEvent.Inputs {
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
			// NOTE: Ints are converted to float64 and stored in DB as float64
			num, numOk := eventField.(float64)
			if !numOk {
				return event, errors.New("Cannot cast DB float to float64")
			}
			bigInt := new(big.Int)
			bigInt.SetString(strconv.FormatFloat(num, 'f', -1, 64), 10)
			eventPayload[eventFieldName] = bigInt
		case "string":
			str, stringOk := eventField.(string)
			if !stringOk {
				return event, errors.New("Cannot cast DB string val to string")
			}
			eventPayload[eventFieldName] = str
		default:
			return event, fmt.Errorf("unsupported type in %v field encountered in %v event",
				eventFieldName, c.EventHash)
		}
	}

	logPayload := c.DBToEventLogData()
	contractAddress := common.HexToAddress(c.ContractAddress)
	event, err = model.NewEvent(
		c.EventType,
		c.ContractName,
		contractAddress,
		c.Timestamp,
		model.RetrievalMethod(c.RetrievalMethod),
		eventPayload,
		logPayload,
	)

	return event, err

}

// EventLogDataToDB explicitly converts the raw log data to Postgresql types
func (c *Event) EventLogDataToDB(payload *types.Log) {
	c.LogPayload["Address"] = payload.Address.Hex()

	topics := make([]interface{}, len(payload.Topics))

	for index, topic := range payload.Topics {
		topics[index] = topic.Hex()
	}

	c.LogPayload["Topics"] = topics
	// Store data field as a base64 encoded string to retain valid []byte.
	c.LogPayload["Data"] = base64.StdEncoding.EncodeToString(payload.Data)
	c.LogPayload["BlockNumber"] = float64(payload.BlockNumber)
	c.LogPayload["TxHash"] = payload.TxHash.Hex()
	c.LogPayload["TxIndex"] = float64(payload.TxIndex)
	c.LogPayload["BlockHash"] = payload.BlockHash.Hex()
	c.LogPayload["Index"] = float64(payload.Index)
	c.LogPayload["Removed"] = payload.Removed

}

// DBToEventLogData converts the DB raw log payload back to types.Log
// NOTE(IS): because jsonb payloads are stored in DB as a map[string]interface{},
// Postgres converts some fields, see notes in function.
func (c *Event) DBToEventLogData() *types.Log {
	tlog := &types.Log{}
	tlog.Address = common.HexToAddress(c.LogPayload["Address"].(string))

	tlog.Topics = c.typeInferTopics(c.LogPayload["Topics"])
	// Data is a base64 encoded string from []byte. Convert this back to []byte.
	tlog.Data = c.typeInferData(c.LogPayload["Data"])
	// BlockNumber is stored in DB as float64
	tlog.BlockNumber = c.typeInferBlockNumber(c.LogPayload["BlockNumber"])
	tlog.TxHash = common.HexToHash(c.LogPayload["TxHash"].(string))
	// TxIndex is stored in DB as float64
	tlog.TxIndex = c.typeInferIndex(c.LogPayload["TxIndex"])
	tlog.BlockHash = common.HexToHash(c.LogPayload["BlockHash"].(string))
	// Index is stored in DB as float64
	tlog.Index = c.typeInferIndex(c.LogPayload["Index"])
	tlog.Removed = c.LogPayload["Removed"].(bool)

	return tlog
}

func (c *Event) typeInferIndex(txIndexInterface interface{}) uint {
	var returnTxIndex uint
	switch val := txIndexInterface.(type) {
	case float64:
		returnTxIndex = uint(val)
	default:
		log.Errorf("DB Index type infer expected as float64, instead is %T", val)
	}
	return returnTxIndex
}

func (c *Event) typeInferBlockNumber(blockInterface interface{}) uint64 {
	var returnBlockNumber uint64
	switch val := blockInterface.(type) {
	case float64:
		returnBlockNumber = uint64(val)
	default:
		log.Errorf("DB Block number type infer expected as float64, instead is %T", val)
	}
	return returnBlockNumber
}

func (c *Event) typeInferData(dataInterface interface{}) []byte {
	var returnData []byte
	switch data := dataInterface.(type) {
	// Decode base64 encoded string back to []byte
	case string:
		bys, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			log.Errorf("Error decoding base64 data field: err: %v", err)
		} else {
			returnData = bys
		}
	default:
		log.Errorf("DB Data type infer expected as []byte, instead is %T", data)
	}
	return returnData
}

func (c *Event) typeInferTopics(topicsInterface interface{}) []common.Hash {
	var theTopics []common.Hash
	switch topics := topicsInterface.(type) {
	case []interface{}:
		theTopics = make([]common.Hash, len(topics))
		for i, topicInterface := range topics {
			switch topic := topicInterface.(type) {
			case string:
				theTopics[i] = common.HexToHash(topic)
			default:
				log.Errorf("DB Topic type infer expected as string, instead is %T", topic)
			}
		}
	default:
		log.Errorf("DB Topics type expected as []interface{}, instead is %T", topics)
	}
	return theTopics
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
