// Package eventcollector_test contains tests for the eventcollector
package eventcollector_test

import (
	"errors"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	cutils "github.com/joincivil/civil-events-crawler/pkg/contractutils"
	"github.com/joincivil/civil-events-crawler/pkg/eventcollector"
	"github.com/joincivil/civil-events-crawler/pkg/generated/filterer"
	"github.com/joincivil/civil-events-crawler/pkg/generated/watcher"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence"
)

type testTrigger struct {
	shouldRun bool
	runErr    error
}

func (t *testTrigger) Description() string {
	return "testtrigger"
}

func (t *testTrigger) ShouldRun(collector *eventcollector.CivilEventCollector, event *model.CivilEvent) bool {
	return t.shouldRun
}

func (t *testTrigger) Run(collector *eventcollector.CivilEventCollector, event *model.CivilEvent) error {
	return t.runErr
}

type lastBlockData struct {
	eventType       string
	contractAddress string
	lastBlockNumber uint64
}

type testPersister struct {
	lastBlock            lastBlockData
	events               []model.CivilEvent
	m                    sync.Mutex
	updateLastBlockError error
	saveEventsError      error
}

func (n *testPersister) LastBlockNumber(eventType string, contractAddress common.Address) uint64 {
	return n.lastBlock.lastBlockNumber
}

func (n *testPersister) LastBlockHash(eventType string, contractAddress common.Address) common.Hash {
	return common.Hash{}
}

func (n *testPersister) UpdateLastBlockData(events []model.CivilEvent) error {
	if len(events) == 0 {
		return n.updateLastBlockError
	}
	event := events[0]
	n.lastBlock.eventType = event.EventType()
	n.lastBlock.contractAddress = event.ContractAddress().Hex()
	rawLog := event.LogPayload()
	n.lastBlock.lastBlockNumber = rawLog.BlockNumber
	return n.updateLastBlockError
}

func (n *testPersister) SaveEvents(events []model.CivilEvent) error {
	n.m.Lock()
	defer n.m.Unlock()
	if n.events == nil {
		n.events = []model.CivilEvent{}
	}
	n.events = append(n.events, events...)
	return n.saveEventsError
}

func (n *testPersister) RetrieveEvents(offset uint, count uint, reverse bool) ([]*model.CivilEvent, error) {
	n.m.Lock()
	defer n.m.Unlock()

	events := []*model.CivilEvent{}
	for index := range n.events {
		events = append(events, &n.events[index])
	}
	return events, nil
}

type testErrorWatcher struct{}

func (t *testErrorWatcher) ContractName() string {
	return "TestErrorContract"
}

func (t *testErrorWatcher) ContractAddress() common.Address {
	return common.HexToAddress("0xf86a8a467666c752fa99bb7ca954d269ab6136bf")
}

func (t *testErrorWatcher) StopWatchers() error {
	return nil
}

func (t *testErrorWatcher) StartWatchers(client bind.ContractBackend,
	eventRecvChan chan model.CivilEvent) ([]event.Subscription, error) {
	return nil, errors.New("This is an error starting watchers")
}

func collectionStart(collector *eventcollector.CivilEventCollector, t *testing.T,
	errChan chan error) {
	err := collector.StartCollection()
	if err != nil {
		t.Errorf("Error starting collection: err: %v", err)
		errChan <- err
	}
}

func setupTestCollector(contracts *cutils.AllTestContracts) *eventcollector.CivilEventCollector {
	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(contracts.CivilTcrAddr),
		filterer.NewNewsroomContractFilterers(contracts.NewsroomAddr),
	}
	watchers := []model.ContractWatchers{
		watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr),
		watcher.NewNewsroomContractWatchers(contracts.NewsroomAddr),
	}
	persister := &persistence.NullPersister{}
	triggers := []eventcollector.Trigger{
		&eventcollector.AddNewsroomWatchersTrigger{},
		&eventcollector.RemoveNewsroomWatchersTrigger{},
		&testTrigger{shouldRun: true, runErr: nil},
	}

	collector := eventcollector.NewCivilEventCollector(
		contracts.Client,
		filterers,
		watchers,
		persister,
		persister,
		persister,
		triggers,
	)
	return collector
}

func setupTestCollectorTestPersister(contracts *cutils.AllTestContracts) (*eventcollector.CivilEventCollector, *testPersister) {
	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(contracts.CivilTcrAddr),
		filterer.NewNewsroomContractFilterers(contracts.NewsroomAddr),
	}
	watchers := []model.ContractWatchers{
		watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr),
		watcher.NewNewsroomContractWatchers(contracts.NewsroomAddr),
	}
	persister := &testPersister{}
	triggers := []eventcollector.Trigger{
		&eventcollector.AddNewsroomWatchersTrigger{},
		&eventcollector.RemoveNewsroomWatchersTrigger{},
		&testTrigger{shouldRun: true, runErr: nil},
	}
	collector := eventcollector.NewCivilEventCollector(
		contracts.Client,
		filterers,
		watchers,
		persister,
		persister,
		persister,
		triggers,
	)
	return collector, persister
}

func setupTestCollectorTestPersisterBadSaveEvents(contracts *cutils.AllTestContracts) (*eventcollector.CivilEventCollector, *testPersister) {
	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(contracts.CivilTcrAddr),
		filterer.NewNewsroomContractFilterers(contracts.NewsroomAddr),
	}
	watchers := []model.ContractWatchers{
		watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr),
		watcher.NewNewsroomContractWatchers(contracts.NewsroomAddr),
	}
	badPersister := &testPersister{saveEventsError: errors.New("Bad save")}
	goodPersister := &testPersister{}
	triggers := []eventcollector.Trigger{
		&eventcollector.AddNewsroomWatchersTrigger{},
		&eventcollector.RemoveNewsroomWatchersTrigger{},
		&testTrigger{shouldRun: true, runErr: nil},
	}
	collector := eventcollector.NewCivilEventCollector(
		contracts.Client,
		filterers,
		watchers,
		goodPersister,
		goodPersister,
		badPersister,
		triggers,
	)
	return collector, badPersister
}

func setupTestCollectorTestPersisterBadUpdateBlockData(contracts *cutils.AllTestContracts) (*eventcollector.CivilEventCollector, *testPersister) {
	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(contracts.CivilTcrAddr),
		filterer.NewNewsroomContractFilterers(contracts.NewsroomAddr),
	}
	watchers := []model.ContractWatchers{
		watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr),
		watcher.NewNewsroomContractWatchers(contracts.NewsroomAddr),
	}
	badPersister := &testPersister{updateLastBlockError: errors.New("Bad update")}
	goodPersister := &testPersister{}
	triggers := []eventcollector.Trigger{
		&eventcollector.AddNewsroomWatchersTrigger{},
		&eventcollector.RemoveNewsroomWatchersTrigger{},
		&testTrigger{shouldRun: true, runErr: nil},
	}
	collector := eventcollector.NewCivilEventCollector(
		contracts.Client,
		filterers,
		watchers,
		goodPersister,
		goodPersister,
		badPersister,
		triggers,
	)
	return collector, badPersister
}

func setupTestCollectorBadWatcher(contracts *cutils.AllTestContracts) *eventcollector.CivilEventCollector {
	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(contracts.CivilTcrAddr),
		filterer.NewNewsroomContractFilterers(contracts.NewsroomAddr),
	}
	watchers := []model.ContractWatchers{
		&testErrorWatcher{},
	}
	persister := &persistence.NullPersister{}
	triggers := []eventcollector.Trigger{
		&eventcollector.AddNewsroomWatchersTrigger{},
		&eventcollector.RemoveNewsroomWatchersTrigger{},
	}

	collector := eventcollector.NewCivilEventCollector(
		contracts.Client,
		filterers,
		watchers,
		persister,
		persister,
		persister,
		triggers,
	)
	return collector
}

func TestNewCivilEventCollector(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	collector := setupTestCollector(contracts)

	errChan := make(chan error)
	go collectionStart(collector, t, errChan)

	select {
	case err := <-errChan:
		t.Errorf("Should not have received error on start collection: err: %v", err)
	case <-time.After(5 * time.Second):
	}
}

func TestNewCivilEventCollectorBadWatcher(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	collector := setupTestCollectorBadWatcher(contracts)
	err = collector.StartCollection()
	if err == nil {
		t.Errorf("Should have received error on start collection: err: %v", err)
	}
}

func TestCivilEventCollectorStopCollection(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	collector := setupTestCollector(contracts)

	errChan := make(chan error)
	go collectionStart(collector, t, errChan)

	err = collector.StopCollection()
	if err != nil {
		t.Errorf("Should not have returned an error when stopping collection: err: %v", err)
	}
}

func TestCivilEventCollectorAddRemoveWatchers(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	collector := setupTestCollector(contracts)

	errChan := make(chan error)
	go collectionStart(collector, t, errChan)

	time.Sleep(3 * time.Second)

	err = collector.RemoveWatchers(
		watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr),
	)
	if err != nil {
		t.Errorf("Should not have returned an error when removing watcher: err: %v", err)
	}

	err = collector.AddWatchers(
		watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr),
	)
	if err != nil {
		t.Errorf("Should not have returned an error when adding watcher: err: %v", err)
	}
}

func TestCivilEventCollectorCollection(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	collector, persister := setupTestCollectorTestPersister(contracts)

	errChan := make(chan error)
	go collectionStart(collector, t, errChan)

	_, err = contracts.CivilTcrContract.Apply(contracts.Auth, contracts.NewsroomAddr, big.NewInt(400), "")
	if err != nil {
		t.Fatalf("Application failed: err: %v", err)
	}

	contracts.Client.Commit()

	_, err = contracts.CivilTcrContract.Withdraw(contracts.Auth, contracts.NewsroomAddr, big.NewInt(50))
	if err != nil {
		t.Fatalf("Withdrawal failed: err: %v", err)
	}

	contracts.Client.Commit()

	_, err = contracts.CivilTcrContract.Deposit(contracts.Auth, contracts.NewsroomAddr, big.NewInt(50))
	if err != nil {
		t.Fatalf("Deposit failed: err: %v", err)
	}

	contracts.Client.Commit()

	time.Sleep(3 * time.Second)

	events, _ := persister.RetrieveEvents(0, 10, false)

	if len(events) == 0 {
		t.Error("Should have seen some events in the persister")
	}
	if len(events) != 6 {
		t.Error("Should have seen 6 events in the persister")
	}
}

func TestNewCivilEventCollectorBadEventSave(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	collector, _ := setupTestCollectorTestPersisterBadSaveEvents(contracts)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		e := collector.StartCollection()
		if e == nil {
			t.Errorf("Should have returned an error from start collection: err: %v", e)
		}
	}()

	_, err = contracts.CivilTcrContract.Apply(contracts.Auth, contracts.NewsroomAddr, big.NewInt(400), "")
	if err != nil {
		t.Errorf("Application failed: err: %v", err)
	}
	contracts.Client.Commit()
	wg.Wait()
}

func TestNewCivilEventCollectorBadUpdateBlockData(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	collector, _ := setupTestCollectorTestPersisterBadUpdateBlockData(contracts)

	go func() {
		e := collector.StartCollection()
		if e == nil {
			t.Errorf("Should have returned an error from start collection: err: %v", e)
		}
	}()

	_, err = contracts.CivilTcrContract.Apply(contracts.Auth, contracts.NewsroomAddr, big.NewInt(400), "")
	if err != nil {
		t.Errorf("Application failed: err: %v", err)
	}
	contracts.Client.Commit()
}
