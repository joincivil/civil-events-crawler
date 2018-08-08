package postgres // import "github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"reflect"
)

// JsonbPayload is the jsonb payload
type JsonbPayload map[string]interface{}

// Value is the value interface implemented for the sql driver
func (jp JsonbPayload) Value() (driver.Value, error) {
	return json.Marshal(jp)
}

// Scan is the scan interface implemented for the sql driver
func (jp *JsonbPayload) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}
	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*jp, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}

// ReturnEventsFromABI returns abi.Event struct from the ABI
func ReturnEventsFromABI(_abi abi.ABI, eventType string) (abi.Event, error) {
	// Trim the eventType clean
	events := abi.Event{}
	ok := false
	// Some contracts have an underscore prefix on their events. Handle both
	// non-underscore/underscore cases here.
	events, ok = _abi.Events[eventType]
	if !ok {
		events, ok = _abi.Events[fmt.Sprintf("_%s", eventType)]
		if !ok {
			return events, fmt.Errorf("No event type %v in contract", eventType)
		}
	}
	return events, nil
}

// StructFieldsForQuery is a generic Insert statement for any table
// TODO(IS): gosec linting errors for bytes.buffer use here. is it inefficient?
func StructFieldsForQuery(exampleStruct interface{}, colon bool) (string, string) {
	var fields bytes.Buffer
	var fieldsWithColon bytes.Buffer
	valStruct := reflect.ValueOf(exampleStruct)
	typeOf := valStruct.Type()
	for i := 0; i < valStruct.NumField(); i++ {
		dbFieldName := typeOf.Field(i).Tag.Get("db")
		fields.WriteString(dbFieldName) // nolint: gosec
		if colon {
			fieldsWithColon.WriteString(":")         // nolint: gosec
			fieldsWithColon.WriteString(dbFieldName) // nolint: gosec
		}
		if i+1 < valStruct.NumField() {
			fields.WriteString(", ") // nolint: gosec
			if colon {
				fieldsWithColon.WriteString(", ") // nolint: gosec
			}
		}
	}
	return fields.String(), fieldsWithColon.String()
}

// InsertIntoDBQueryString creates the query to insert a given struct into a given table
func InsertIntoDBQueryString(tableName string, dbModelStruct interface{}) string {
	fieldNames, fieldNamesColon := StructFieldsForQuery(dbModelStruct, true)
	queryString := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s);", tableName, fieldNames, fieldNamesColon) // nolint: gosec
	return queryString
}
