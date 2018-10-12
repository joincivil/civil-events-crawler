// Package gen_test contains tests for the gen package
package gen_test

import (
	"bytes"
	"github.com/joincivil/civil-events-crawler/pkg/gen"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"strings"
	"testing"
	"time"
)

func TestGenerateEventHandlersFromTemplate(t *testing.T) {
	Event1 := &gen.EventHandlerTmplData{
		EventMethod: "Application",
		EventName:   "_Application",
		EventType:   "CivilTCRContractApplication",
		ParamValues: []*gen.EventHandlerMethodParam{
			{Type: "common.Address"},
			{Type: "common.Address"},
		},
	}
	Event2 := &gen.EventHandlerTmplData{
		EventMethod: "ApplicationRemoved",
		EventName:   "_ApplicationRemoved",
		EventType:   "CivilTCRContractApplicationRemoved",
		ParamValues: []*gen.EventHandlerMethodParam{
			{Type: "common.Address"},
		},
	}
	testWatchers := &gen.EventHandlerContractTmplData{
		PackageName:         "watcher",
		ContractImportPath:  "github.com/joincivil/civil-events-crawler/pkg/generated/contract",
		ContractTypePackage: "contract",
		ContractTypeName:    "CivilTCRContract",
		GenTime:             time.Now().UTC(),
		EventHandlers: []*gen.EventHandlerTmplData{
			Event1,
			Event2,
		},
	}
	testFilterers := &gen.EventHandlerContractTmplData{
		PackageName:         "retrieve",
		ContractImportPath:  "github.com/joincivil/civil-events-crawler/pkg/generated/contract",
		ContractTypePackage: "contract",
		ContractTypeName:    "CivilTCRContract",
		GenTime:             time.Now().UTC(),
		EventHandlers: []*gen.EventHandlerTmplData{
			Event1,
			Event2,
		},
	}
	bufWatcher := &bytes.Buffer{}
	bufFilterer := &bytes.Buffer{}
	err := gen.GenerateEventHandlersFromTemplate(bufWatcher, testWatchers, true, "watcher")
	if err != nil {
		t.Errorf("Error generating watchers: err: %v", err)
	}
	err = gen.GenerateEventHandlersFromTemplate(bufFilterer, testFilterers, true, "filterer")
	if err != nil {
		t.Errorf("Error generating filterers: err: %v", err)
	}

	// TODO(PN or IS): Some basic checks, need more here.
	watcherCode := bufWatcher.String()
	filtererCode := bufFilterer.String()

	if !strings.Contains(watcherCode, "func (w *CivilTCRContractWatchers) StartWatchers") {
		t.Error("Did not see expected Startwatchers in the generated watcher code")
	}
	if !strings.Contains(watcherCode, "func (w *CivilTCRContractWatchers) startWatchApplication") {
		t.Error("Did not see expected startWatchApplication in the generated watcher code")
	}
	if !strings.Contains(watcherCode, "func (w *CivilTCRContractWatchers) startWatchApplicationRemoved") {
		t.Error("Did not see expected startWatchApplicationRemoved in the generated watcher code")
	}
	if !strings.Contains(filtererCode, "func (f *CivilTCRContractFilterers) StartFilterers") {
		t.Error("Did not see expected StartFilterers in the generated Filterer code")
	}
	if !strings.Contains(filtererCode, "func (f *CivilTCRContractFilterers) startFilterApplication") {
		t.Error("Did not see expected startFilterApplication in the generated filterer code")
	}

}

func TestGenerateWatchersForCivilTcr(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateEventHandlers(buf, model.CivilTcrContractType, "watcher", "watcher")
	if err != nil {
		t.Errorf("Error generating watchers for the Civil TCR contract: err: %v", err)
	}

	// TODO(PN): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (w *CivilTCRContractWatchers) StartWatchers") {
		t.Error("Did not see expected StartWatchers in the generated watcher code")
	}
	if !strings.Contains(code, "func (w *CivilTCRContractWatchers) StartCivilTCRContractWatchers") {
		t.Error("Did not see expected StartCivilTCRContractWatchers in the generated code")
	}
	if !strings.Contains(code, "func (w *CivilTCRContractWatchers) startWatchApplication") {
		t.Error("Did not see expected startWatchApplication in the generated code")
	}
	if !strings.Contains(code, "func (w *CivilTCRContractWatchers) startWatchApplicationRemoved") {
		t.Error("Did not see expected startWatchApplicationRemoved in the generated code")
	}
}

func TestGenerateWatchersForNewsroom(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateEventHandlers(buf, model.NewsroomContractType, "watcher", "watcher")
	if err != nil {
		t.Errorf("Error generating watchers for the Newsroom contract: err: %v", err)
	}

	// TODO(PN): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (w *NewsroomContractWatchers) StartWatchers") {
		t.Error("Did not see expected StartWatchers in the generated watcher code")
	}
	if !strings.Contains(code, "func (w *NewsroomContractWatchers) StartNewsroomContractWatchers") {
		t.Error("Did not see expected StartNewsroomContractWatchers in the generated code")
	}
	if !strings.Contains(code, "func (w *NewsroomContractWatchers) startWatchRevisionUpdated") {
		t.Error("Did not see expected startWatchRevisionUpdated in the generated code")
	}
}

func TestGenerateWatchersForPLCRVoting(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateEventHandlers(buf, model.PLCRVotingContractType, "watcher", "watcher")
	if err != nil {
		t.Errorf("Error generating watchers for the PLCRVoting contract: err: %v", err)
	}

	// TODO(PN): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (w *PLCRVotingContractWatchers) StartWatchers") {
		t.Error("Did not see expected StartWatchers in the generated watcher code")
	}
	if !strings.Contains(code, "func (w *PLCRVotingContractWatchers) StartPLCRVotingContractWatchers") {
		t.Error("Did not see expected StartNewsroomContractWatchers in the generated code")
	}
	if !strings.Contains(code, "func (w *PLCRVotingContractWatchers) startWatchPollCreated") {
		t.Error("Did not see expected startWatchRevisionUpdated in the generated code")
	}
}

func TestGenerateFilterersForCivilTcr(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateEventHandlers(buf, model.CivilTcrContractType, "filterer", "filterer")
	if err != nil {
		t.Errorf("Error generating filterers for the Civil TCR contract: err: %v", err)
	}

	// TODO(IS): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (f *CivilTCRContractFilterers) StartFilterers") {
		t.Error("Did not see expected StartFilterers in the generated code")
	}
	if !strings.Contains(code, "func (f *CivilTCRContractFilterers) startFilterApplication") {
		t.Error("Did not see expected startFilterApplication in the generated code")
	}
	if !strings.Contains(code, "func (f *CivilTCRContractFilterers) startFilterApplicationRemoved") {
		t.Error("Did not see expected startFilterApplicationRemoved in the generated code")
	}
}

func TestGenerateFilterersForNewsroom(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateEventHandlers(buf, model.NewsroomContractType, "filterer", "filterer")
	if err != nil {
		t.Errorf("Error generating Filterers for the Newsroom contract: err: %v", err)
	}

	// TODO(IS): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (f *NewsroomContractFilterers) StartFilterers") {
		t.Error("Did not see expected StartFilterers in the generated code")
	}
	if !strings.Contains(code, "func (f *NewsroomContractFilterers) startFilterRevisionUpdated") {
		t.Error("Did not see expected startFilterRevisionUpdated in the generated code")
	}
}

func TestGenerateFilterersForPLCRVoting(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateEventHandlers(buf, model.PLCRVotingContractType, "filterer", "filterer")
	if err != nil {
		t.Errorf("Error generating filterers for the PLCRVoting contract: err: %v", err)
	}

	// TODO(IS): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (f *PLCRVotingContractFilterers) StartFilterers") {
		t.Error("Did not see expected StartFilterers in the generated code")
	}
	if !strings.Contains(code, "func (f *PLCRVotingContractFilterers) startFilterPollCreated") {
		t.Error("Did not see expected startFilterPollCreated in the generated code")
	}
	if !strings.Contains(code, "func (f *PLCRVotingContractFilterers) startFilterVotingRightsWithdrawn") {
		t.Error("Did not see expected startFilterVotingRightsWithdrawn in the generated code")
	}
}
