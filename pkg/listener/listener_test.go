// Package listener_test contains tests for the listener package
package listener_test

import (
	"errors"
	"math/big"
	"runtime"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	ctime "github.com/joincivil/go-common/pkg/time"

	cutils "github.com/joincivil/civil-events-crawler/pkg/contractutils"
	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/generated/watcher"
	"github.com/joincivil/civil-events-crawler/pkg/listener"
	"github.com/joincivil/civil-events-crawler/pkg/model"
)

func TestListener(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	watchers := []model.ContractWatchers{
		watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr),
		watcher.NewNewsroomContractWatchers(contracts.NewsroomAddr),
	}
	listener := setupListener(t, contracts.Client, watchers)
	defer listener.Stop()
}

func TestListenerStop(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	watchers := []model.ContractWatchers{
		watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr),
		watcher.NewNewsroomContractWatchers(contracts.NewsroomAddr),
	}
	listener := setupListener(t, contracts.Client, watchers)

	// Simple test is if each watcher loop goroutine is shut down by Stop()
	initialNumRoutines := runtime.NumGoroutine()
	listener.Stop()
	if initialNumRoutines <= runtime.NumGoroutine() {
		t.Errorf("Number of goroutines has not gone down since listener.Stop")
	}
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
	eventRecvChan chan *model.Event) ([]event.Subscription, error) {
	return nil, errors.New("This is an error starting watchers")
}

func TestListenerEmptyWatchers(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	watchers := []model.ContractWatchers{
		&testErrorWatcher{},
	}
	listener := listener.NewEventListener(contracts.Client, watchers)
	if listener == nil {
		t.Fatal("Listener should not be nil")
	}
	err = listener.Start()
	if err == nil {
		t.Errorf("Listener should have failed with no watchers: %v", err)
	}
}

func TestListenerErrorStartWatchers(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	watchers := []model.ContractWatchers{}
	listener := listener.NewEventListener(contracts.Client, watchers)
	if listener == nil {
		t.Fatal("Listener should not be nil")
	}
	err = listener.Start()
	if err == nil {
		t.Errorf("Listener should have failed with error from StartWatchers: %v", err)
	}
}

// TestListenerEventChan mainly tests the EventRecvChan to ensure it can
// pass along a Event object
func TestListenerEventChan(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: err: %v", err)
	}
	watchers := []model.ContractWatchers{
		watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr),
		watcher.NewNewsroomContractWatchers(contracts.NewsroomAddr),
	}
	listener := setupListener(t, contracts.Client, watchers)
	defer listener.Stop()
	quitChan := make(chan interface{})
	eventRecv := make(chan bool)

	go func(quit <-chan interface{}, recv chan<- bool) {
		for {
			select {
			case event := <-listener.EventRecvChan:
				if event.EventType() != "Application" {
					t.Errorf("Eventtype is not correct: %v", event.EventType())
				}
				recv <- true
			case <-quit:
				return
			}
		}
	}(quitChan, eventRecv)

	tempPayload := &contract.CivilTCRContractApplication{
		ListingAddress: contracts.CivilTcrAddr,
		Deposit:        big.NewInt(1000),
		AppEndDate:     big.NewInt(1653860896),
		Data:           "DATA",
		Applicant:      contracts.CivilTcrAddr,
		Raw: types.Log{
			Address:     contracts.CivilTcrAddr,
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: 8888888,
		},
	}
	newEvent, _ := model.NewEventFromContractEvent("Application", "CivilTCRContract", contracts.CivilTcrAddr, tempPayload,
		ctime.CurrentEpochSecsInInt64(), model.Filterer)
	listener.EventRecvChan <- newEvent

	select {
	case <-eventRecv:
		close(quitChan)
	case <-time.After(5 * time.Second):
		t.Errorf("Event not received, should have been received by goroutine")
		close(quitChan)
	}
}

// TestListenerContractEvents tests event fired from a call to Apply()
// on a simulated TCR on a simulated backend. Tests two events so ensure
// we are handling two different events on the same channel.
func TestListenerContractEvents(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: err: %v", err)
	}
	watchers := []model.ContractWatchers{
		watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr),
		watcher.NewNewsroomContractWatchers(contracts.NewsroomAddr),
	}
	listener := setupListener(t, contracts.Client, watchers)
	defer listener.Stop()

	fireContractEventsSequence(t, contracts, listener)
}

func TestListenerAddWatchers(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: err: %v", err)
	}
	watchers := []model.ContractWatchers{
		watcher.NewNewsroomContractWatchers(contracts.NewsroomAddr),
	}
	listener := setupListener(t, contracts.Client, watchers)
	defer listener.Stop()

	watchersToAdd := watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr)
	listener.AddWatchers(watchersToAdd)

	fireContractEventsSequence(t, contracts, listener)
}

func TestListenerRemoveWatchersNoMoreWatchers(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: err: %v", err)
	}
	w := watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr)
	watchers := []model.ContractWatchers{
		w,
		watcher.NewNewsroomContractWatchers(contracts.NewsroomAddr),
	}
	listener := setupListener(t, contracts.Client, watchers)
	defer listener.Stop()

	// Remove the newsroom watchers and see if we still receive the TCR events.
	listener.RemoveWatchers(w)

	quitChan := make(chan interface{})
	eventRecv := make(chan bool)

	setupEventRecvLoop(t, listener, quitChan, eventRecv)
	expectedNumEvents := makeTheCalls(t, contracts)
	checkNotRecvEvents(t, quitChan, eventRecv, expectedNumEvents)
}

func TestListenerRemoveWatchers(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: err: %v", err)
	}
	w := watcher.NewNewsroomContractWatchers(contracts.NewsroomAddr)
	watchers := []model.ContractWatchers{
		w,
		watcher.NewCivilTCRContractWatchers(contracts.CivilTcrAddr),
	}
	listener := setupListener(t, contracts.Client, watchers)
	defer listener.Stop()

	// Remove the newsroom watchers and see if we still receive the TCR events.
	listener.RemoveWatchers(w)
	fireContractEventsSequence(t, contracts, listener)
}

func fireContractEventsSequence(t *testing.T, contracts *cutils.AllTestContracts,
	listener *listener.EventListener) {
	quitChan := make(chan interface{})
	eventRecv := make(chan bool)

	setupEventRecvLoop(t, listener, quitChan, eventRecv)
	expectedNumEvents := makeTheCalls(t, contracts)
	checkRecvEvents(t, quitChan, eventRecv, expectedNumEvents)
}

func setupEventRecvLoop(t *testing.T, listener *listener.EventListener,
	quitChan chan interface{}, eventRecv chan bool) {
	go func(quit <-chan interface{}, recv chan<- bool) {
		for {
			select {
			case event := <-listener.EventRecvChan:
				if event.EventType() != "Application" && event.EventType() != "Withdrawal" &&
					event.EventType() != "Deposit" {
					t.Errorf("EventType is not correct: eventType: %v", event.EventType())
				}
				recv <- true
			case <-quit:
				return
			}
		}
	}(quitChan, eventRecv)
}

func makeTheCalls(t *testing.T, contracts *cutils.AllTestContracts) int {
	expectedNumEvents := 3

	_, err := contracts.CivilTcrContract.Apply(contracts.Auth, contracts.NewsroomAddr, big.NewInt(400), "")
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

	return expectedNumEvents
}

func checkRecvEvents(t *testing.T, quitChan chan interface{}, eventRecv chan bool,
	expectedNumEvents int) {
	numEvents := 0
Loop:
	for {
		select {
		case <-eventRecv:
			numEvents++
			if numEvents == expectedNumEvents {
				close(quitChan)
				break Loop
			}
		case <-time.After(10 * time.Second):
			t.Errorf("Not all events were received")
			close(quitChan)
			break Loop
		}
	}
}

func checkNotRecvEvents(t *testing.T, quitChan chan interface{}, eventRecv chan bool,
	expectedNumEvents int) {
Loop:
	for {
		select {
		case <-eventRecv:
			t.Errorf("Should not have received events")
			close(quitChan)
			break Loop
		case <-time.After(5 * time.Second):
			close(quitChan)
			break Loop
		}
	}
}

func setupListener(t *testing.T, client bind.ContractBackend, watchers []model.ContractWatchers) *listener.EventListener {
	listener := listener.NewEventListener(client, watchers)
	if listener == nil {
		t.Fatal("Listener should not be nil")
	}
	err := listener.Start()
	if err != nil {
		t.Errorf("Listener should have started with no errors: %v", err)
	}
	return listener
}
