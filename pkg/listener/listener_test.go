// Package listener_test contains tests for the listener package
package listener_test

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joincivil/civil-events-crawler/pkg/listener"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"testing"
	"time"
)

const (
	rinkebyAddress    = "wss://rinkeby.infura.io/ws"
	rinkebyTCRAddress = "0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d"
)

func setupRinkebyClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(rinkebyAddress)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func setupListener(t *testing.T) *listener.CivilEventListener {
	client, err := setupRinkebyClient()
	if err != nil {
		t.Fatalf("Unable to connect to rinkeby: %v", err)
	}
	listener := listener.NewCivilEventListener(client, rinkebyTCRAddress)
	if listener == nil {
		t.Fatal("Listener should not be nil")
	}
	defer listener.Stop()
	err = listener.Start()
	if err != nil {
		t.Errorf("Listener should have started with no errors: %v", err)
	}
	return listener
}

func TestCivilListener(t *testing.T) {
	listener := setupListener(t)
	defer listener.Stop()
}

// TestCivilListenerRoutine mainly tests the EventRecvChan to ensure it can
// pass along a CivilEvent object
func TestCivilListenerRoutine(t *testing.T) {
	listener := setupListener(t)
	defer listener.Stop()
	quitChan := make(chan interface{})
	eventRecv := make(chan bool)

	go func(quit <-chan interface{}, recv chan<- bool) {
		for {
			select {
			case event := <-listener.EventRecvChan:
				if event.EventType != "_Application" {
					t.Errorf("Eventtype is not correct: %v", event.EventType)
				}
				recv <- true
			case <-quit:
				return
			}
		}
	}(quitChan, eventRecv)

	newEvent := &model.CivilEvent{
		EventType: "_Application",
		Timestamp: utils.CurrentEpochSecsInInt(),
		Payload:   &model.CivilEventPayload{},
	}
	listener.EventRecvChan <- *newEvent

	select {
	case <-eventRecv:
		close(quitChan)
	case <-time.After(5 * time.Second):
		t.Errorf("Event not received, should have been received by goroutine")
		close(quitChan)
	}
}
