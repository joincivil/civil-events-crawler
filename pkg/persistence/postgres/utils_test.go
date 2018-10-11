package postgres_test

import (
	"strings"
	"testing"

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
