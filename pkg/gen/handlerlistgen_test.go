// Package gen_test contains tests for the gen package
package gen_test

import (
	"bytes"
	"testing"

	"github.com/joincivil/civil-events-crawler/pkg/gen"
)

func TestGenerateEventHandlerLists(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateEventHandlerLists(buf, "handlerlist")
	if err != nil {
		t.Errorf("Error generating event handler lists: err: %v", err)
	}
}
