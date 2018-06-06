// Package gen_test contains tests for the gen package
package gen_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/joincivil/civil-events-crawler/pkg/gen"
)

func TestGenerateWatchersFromTemplate(t *testing.T) {
	watchEvent1 := &gen.WatchEvent{
		WatchMethod:    "Application",
		WatchEventName: "_Application",
		EventType:      "CivilTCRContractApplication",
		ParamValues: []*gen.WatchEventMethodParam{
			{Type: "common.Address"},
			{Type: "common.Address"},
		},
	}
	watchEvent2 := &gen.WatchEvent{
		WatchMethod:    "ApplicationRemoved",
		WatchEventName: "_ApplicationRemoved",
		EventType:      "CivilTCRContractApplicationRemoved",
		ParamValues: []*gen.WatchEventMethodParam{
			{Type: "common.Address"},
		},
	}
	testWatchers := &gen.ContractWatchers{
		PackageName:         "contractwatchers",
		ContractImportPath:  "github.com/joincivil/civil-events-crawler/pkg/generated/contract",
		ContractTypePackage: "contract",
		ContractTypeName:    "CivilTCRContract",
		GenTime:             time.Now().UTC(),
		WatchEvents: []*gen.WatchEvent{
			watchEvent1,
			watchEvent2,
		},
	}
	buf := &bytes.Buffer{}
	err := gen.GenerateWatchersFromTemplate(buf, testWatchers, true)
	if err != nil {
		t.Errorf("Error generating watchers: err: %v", err)
	}

	// TODO(PN): Some basic checks, need more here.
	code := buf.String()
	if !strings.Contains(code, "func startWatchApplication") {
		t.Error("Did not see expected startWatchApplication in the generated code")
	}
	if !strings.Contains(code, "func startWatchApplicationRemoved") {
		t.Error("Did not see expected startWatchApplicationRemoved in the generated code")
	}
	if !strings.Contains(code, "_contract *contract.CivilTCRContract") {
		t.Error("Did not see expected contract.CivilTCRContract in the generated code")
	}
}

func TestGenerateWatchersForCivilTcr(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateCivilWatchers(buf, gen.CivilTcrContractType, "contractwatchers")
	if err != nil {
		t.Errorf("Error generating watchers for the Civil TCR contract: err: %v", err)
	}

	// TODO(PN): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (w *CivilTCRContractWatchers) StartCivilTCRContractWatchers(") {
		t.Error("Did not see expected StartCivilTCRContractWatchers in the generated code")
	}
	if !strings.Contains(code, "func startWatchApplication") {
		t.Error("Did not see expected startWatchApplication in the generated code")
	}
	if !strings.Contains(code, "func startWatchApplicationRemoved") {
		t.Error("Did not see expected startWatchApplicationRemoved in the generated code")
	}
	if !strings.Contains(code, "_contract *contract.CivilTCRContract") {
		t.Error("Did not see expected tcr.CivilTCRContract in the generated code")
	}
}

func TestGenerateWatchersForNewsroom(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateCivilWatchers(buf, gen.NewsroomContractType, "contractwatchers")
	if err != nil {
		t.Errorf("Error generating watchers for the Newsroom contract: err: %v", err)
	}

	// TODO(PN): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (w *NewsroomContractWatchers) StartNewsroomContractWatchers(") {
		t.Error("Did not see expected StartNewsroomContractWatchers in the generated code")
	}
	if !strings.Contains(code, "func startWatchRevisionUpdated") {
		t.Error("Did not see expected startWatchRevisionUpdated in the generated code")
	}
	if !strings.Contains(code, "_contract *contract.NewsroomContract") {
		t.Error("Did not see expected newsroom.NewsroomContract in the generated code")
	}
}
