package postgres // import "github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
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
	eventType = strings.Trim(eventType, " _")
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
