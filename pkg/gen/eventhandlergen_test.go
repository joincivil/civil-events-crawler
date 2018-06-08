// Package gen_test contains tests for the gen package
package gen_test

import (
	"bytes"
	"github.com/joincivil/civil-events-crawler/pkg/gen"
	"strings"
	"testing"
	"time"
)

func TestGenerateEventHandlersFromTemplate(t *testing.T) {
	Event1 := &gen.EventHandler{
		EventMethod: "Application",
		EventName:   "_Application",
		EventType:   "CivilTCRContractApplication",
		ParamValues: []*gen.EventHandlerMethodParam{
			{Type: "common.Address"},
			{Type: "common.Address"},
		},
	}
	Event2 := &gen.EventHandler{
		EventMethod: "ApplicationRemoved",
		EventName:   "_ApplicationRemoved",
		EventType:   "CivilTCRContractApplicationRemoved",
		ParamValues: []*gen.EventHandlerMethodParam{
			{Type: "common.Address"},
		},
	}
	testWatchers := &gen.ContractData{
		PackageName:         "watcher",
		ContractImportPath:  "github.com/joincivil/civil-events-crawler/pkg/generated/contract",
		ContractTypePackage: "contract",
		ContractTypeName:    "CivilTCRContract",
		GenTime:             time.Now().UTC(),
		EventHandlers: []*gen.EventHandler{
			Event1,
			Event2,
		},
	}
	testRetrievers := &gen.ContractData{
		PackageName:         "retrieve",
		ContractImportPath:  "github.com/joincivil/civil-events-crawler/pkg/generated/contract",
		ContractTypePackage: "contract",
		ContractTypeName:    "CivilTCRContract",
		GenTime:             time.Now().UTC(),
		EventHandlers: []*gen.EventHandler{
			Event1,
			Event2,
		},
	}
	bufWatcher := &bytes.Buffer{}
	bufRetriever := &bytes.Buffer{}
	err := gen.GenerateEventHandlersFromTemplate(bufWatcher, testWatchers, true, "watcher")
	if err != nil {
		t.Errorf("Error generating watchers: err: %v", err)
	}
	err = gen.GenerateEventHandlersFromTemplate(bufRetriever, testRetrievers, true, "retrieve")
	if err != nil {
		t.Errorf("Error generating retrievers: err: %v", err)
	}

	// TODO(PN or IS): Some basic checks, need more here.
	watcherCode := bufWatcher.String()
	retrieverCode := bufRetriever.String()

	if !strings.Contains(watcherCode, "func startWatchApplication") {
		t.Error("Did not see expected startWatchApplication in the generated watcher code")
	}
	if !strings.Contains(watcherCode, "func startWatchApplicationRemoved") {
		t.Error("Did not see expected startWatchApplicationRemoved in the generated watcher code")
	}
	if !strings.Contains(watcherCode, "_contract *contract.CivilTCRContract") {
		t.Error("Did not see expected contract.CivilTCRContract in the generated watcher code")
	}
	if !strings.Contains(retrieverCode, "func (r *CivilTCRContractRetrievers) RetrieveCivilTCRContractEvents") {
		t.Error("Did not see expected RetrieveCivilTCRContractEvents in the generated retriever code")
	}
	if !strings.Contains(retrieverCode, "func RetrieveApplication") {
		t.Error("Did not see expected RetrieveApplication in the generated retriever code")
	}
	if !strings.Contains(retrieverCode, "_contract *contract.CivilTCRContract") {
		t.Error("Did not see expected contract.CivilTCRContract in the generated retriever code")
	}
}

func TestGenerateWatchersForCivilTcr(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateCivilEventHandlers(buf, gen.CivilTcrContractType, "watcher")
	if err != nil {
		t.Errorf("Error generating watchers for the Civil TCR contract: err: %v", err)
	}

	// TODO(PN): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (w *CivilTCRContractWatchers) StartCivilTCRContractWatchers") {
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
	err := gen.GenerateCivilEventHandlers(buf, gen.NewsroomContractType, "watcher")
	if err != nil {
		t.Errorf("Error generating watchers for the Newsroom contract: err: %v", err)
	}

	// TODO(PN): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (w *NewsroomContractWatchers) StartNewsroomContractWatchers") {
		t.Error("Did not see expected StartNewsroomContractWatchers in the generated code")
	}
	if !strings.Contains(code, "func startWatchRevisionUpdated") {
		t.Error("Did not see expected startWatchRevisionUpdated in the generated code")
	}
	if !strings.Contains(code, "_contract *contract.NewsroomContract") {
		t.Error("Did not see expected newsroom.NewsroomContract in the generated code")
	}
}

func TestGenerateRetrieversForCivilTcr(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateCivilEventHandlers(buf, gen.CivilTcrContractType, "retrieve")
	if err != nil {
		t.Errorf("Error generating retrievers for the Civil TCR contract: err: %v", err)
	}

	// TODO(IS): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (r *CivilTCRContractRetrievers) RetrieveCivilTCRContractEvents") {
		t.Error("Did not see expected RetrieveCivilTCRContractEvents in the generated code")
	}
	if !strings.Contains(code, "func RetrieveApplication") {
		t.Error("Did not see expected RetrieveApplication in the generated code")
	}
	if !strings.Contains(code, "func RetrieveApplicationRemoved") {
		t.Error("Did not see expected RetrieveApplicationRemoved in the generated code")
	}
	if !strings.Contains(code, "_contract *contract.CivilTCRContract") {
		t.Error("Did not see expected tcr.CivilTCRContract in the generated code")
	}
}

func TestGenerateRetrieversForNewsroom(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateCivilEventHandlers(buf, gen.NewsroomContractType, "retrieve")
	if err != nil {
		t.Errorf("Error generating retrievers for the Newsroom contract: err: %v", err)
	}

	// TODO(IS): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (r *NewsroomContractRetrievers) RetrieveNewsroomContractEvents") {
		t.Error("Did not see expected RetrieveNewsroomContractEvents in the generated code")
	}
	if !strings.Contains(code, "func RetrieveRevisionUpdated") {
		t.Error("Did not see expected RetrieveRevisionUpdated in the generated code")
	}
	if !strings.Contains(code, "_contract *contract.NewsroomContract") {
		t.Error("Did not see expected newsroom.NewsroomContract in the generated code")
	}
}
