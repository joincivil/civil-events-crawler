package postgres_test

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"
)

type testStruct struct {
	field1        string `db:"field1"` // nolint: megacheck
	field2        int    `db:"field2"` // nolint: megacheck
	field4Ignored string `db:"-"`      // nolint: megacheck
	field4        int    `db:"field4"` // nolint: megacheck
	field5Ignored bool   `db:"-"`      // nolint: megacheck
}

func TestStructFieldsForQuery(t *testing.T) {
	fieldNames, fieldNamesColon := postgres.StructFieldsForQuery(testStruct{}, true)

	splitFields := strings.Split(fieldNames, ",")
	if len(splitFields) != 3 {
		t.Errorf("Should have only returned 3 fields, not %v", len(splitFields))
	}
	for _, fieldName := range splitFields {
		fieldName = strings.TrimSpace(fieldName)
		if fieldName == "-" {
			t.Error("Should have not had '-' field in fields")
			break
		}
	}

	splitFieldsColon := strings.Split(fieldNamesColon, ",")
	if len(splitFieldsColon) != 3 {
		t.Errorf("Should have only returned 3 fields, not %v", len(splitFieldsColon))
	}
	for _, fieldName := range splitFieldsColon {
		fieldName = strings.TrimSpace(fieldName)
		if fieldName == ":-" {
			t.Error("Should have not had ':-' field in fields")
			break
		}
	}
}

func TestInsertIntoDBQueryString(t *testing.T) {
	insertString := postgres.InsertIntoDBQueryString("testtable", testStruct{})

	if !strings.Contains(insertString, "INSERT INTO testtable") {
		t.Errorf("Should have had INSERT INTO statement")
	}
	if !strings.Contains(insertString, "field1, field2, field4") {
		t.Errorf("Should have had INSERT INTO fields")
	}
	if !strings.Contains(insertString, "VALUES(:field1, :field2, :field4)") {
		t.Errorf("Should have had correct VALUES statement")
	}
}

func TestJsonPayload(t *testing.T) {
	payload := &postgres.JsonbPayload{
		"key1": "value1",
		"key2": 10,
		"key3": "value3",
		"key4": common.HexToAddress("0x98c8cf45bd844627e84e1c506ca87cc9436317d0"),
	}

	_, err := payload.Value()
	if err != nil {
		t.Errorf("Should not have received error retrieving Value: err: %v", err)
	}

	data := `{
		"field1": "value1",
		"field2": "value2",
		"field4": "value4"
	}`
	badData := `{
		/"field1": "value1",
		/"field2": "value2",
		"field4": "value4"
	}`
	badDataNotMap := `"invaliddata"`
	err = payload.Scan(data)
	if err == nil {
		t.Errorf("Should have received error scanning: err: %v", err)
	}
	err = payload.Scan([]byte(data))
	if err != nil {
		t.Errorf("Should not have received error scanning: err: %v", err)
	}
	err = payload.Scan([]byte(badData))
	if err == nil {
		t.Errorf("Should have received error scanning bad data: err: %v", err)
	}
	err = payload.Scan([]byte(badDataNotMap))
	if err == nil {
		t.Errorf("Should have received error scanning bad data not map: err: %v", err)
	}
}
