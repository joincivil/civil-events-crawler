// +build integration

package utils_test

import (
	// "fmt"
	"testing"
	"time"

	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

func TestCreateDeleteTopic(t *testing.T) {
	ps, err := utils.NewGooglePubSub("civil-media")
	if err != nil {
		t.Fatalf("Should not have failed to create a new pubsub obj: err: %v", err)
	}
	err = ps.CreateTopic("test-topic")
	if err != nil {
		t.Errorf("Should have created a topic")
	}
	err = ps.CreateTopic("test-topic")
	if err == nil {
		t.Errorf("Should have prevented creation of the same topic")
	}
	err = ps.DeleteTopic("test-topic")
	if err != nil {
		t.Errorf("Should have deleted the test-topic")
	}
	err = ps.CreateTopic("test-topic")
	if err != nil {
		t.Errorf("Should have created a topic after it was deleted")
	}
	err = ps.DeleteTopic("test-topic")
	if err != nil {
		t.Errorf("Should have deleted the test-topic")
	}
	ok, err := ps.TopicExists("test-topic")
	if err != nil {
		t.Errorf("Should have successfully checked tpic existence")
	}
	if ok {
		t.Errorf("Should have not found the topic")
	}
}

func TestCreateDeleteSubscription(t *testing.T) {
	ps, err := utils.NewGooglePubSub("civil-media")
	if err != nil {
		t.Fatalf("Should not have failed to create a new pubsub obj: err: %v", err)
	}
	err = ps.CreateSubscription("test-topic", "test-subscription")
	if err == nil {
		t.Errorf("Should not have created a subscription since there is no topic")
	}
	err = ps.CreateTopic("test-topic")
	if err != nil {
		t.Errorf("Should have created a topic")
	}
	err = ps.CreateSubscription("test-topic", "test-subscription")
	if err != nil {
		t.Errorf("Should have created a subscription")
	}
	err = ps.CreateSubscription("test-topic", "test-subscription")
	if err == nil {
		t.Errorf("Should not prevented the creation of a the same subscription")
	}
	err = ps.DeleteSubscription("test-subscription")
	if err != nil {
		t.Errorf("Should have deleted the subscription")
	}
	err = ps.DeleteTopic("test-topic")
	if err != nil {
		t.Errorf("Should have deleted the test-topic")
	}
}

func TestStartStopPubSubPublishers(t *testing.T) {
	ps, err := utils.NewGooglePubSub("civil-media")
	if err != nil {
		t.Fatalf("Should not have failed to create a new pubsub obj: err: %v", err)
	}
	err = ps.StartPublishers()
	if err != nil {
		t.Fatalf("Should have started publishers up: err: %v", err)
	}
	time.Sleep(2 * time.Second)
	err = ps.StopPublishers()
	if err != nil {
		t.Fatalf("Should have stopped publishers: err: %v", err)
	}
}

func TestStartStopPubSubSubscribers(t *testing.T) {
	ps, err := utils.NewGooglePubSub("civil-media")
	if err != nil {
		t.Fatalf("Should not have failed to create a new pubsub obj: err: %v", err)
	}
	err = ps.StartSubscribers("test-subscription")
	if err == nil {
		t.Fatalf("Should not have started up since subscription doesn't exist: err: %v", err)
	}
	err = ps.CreateTopic("test-topic")
	if err != nil {
		t.Errorf("Should have created a topic")
	}
	err = ps.CreateSubscription("test-topic", "test-subscription")
	if err != nil {
		t.Errorf("Should not prevented the creation of a the same subscription")
	}
	err = ps.StartSubscribers("test-subscription")
	if err != nil {
		t.Fatalf("Should have started up subscription: err: %v", err)
	}
	time.Sleep(2 * time.Second)
	err = ps.StopSubscribers()
	if err != nil {
		t.Fatalf("Should have stopped subscribers: err: %v", err)
	}
	err = ps.DeleteSubscription("test-subscription")
	if err != nil {
		t.Errorf("Should have deleted the subscription")
	}
	err = ps.DeleteTopic("test-topic")
	if err != nil {
		t.Errorf("Should have deleted the test-topic")
	}
}

func TestPubSubPublishersPublish(t *testing.T) {
	ps, err := utils.NewGooglePubSub("civil-media")
	if err != nil {
		t.Fatalf("Should not have failed to create a new pubsub obj: err: %v", err)
	}
	err = ps.CreateTopic("test-topic")
	if err != nil {
		t.Errorf("Should have created a topic")
	}
	err = ps.CreateSubscription("test-topic", "test-subscription")
	if err != nil {
		t.Errorf("Should not prevented the creation of a the same subscription")
	}

	err = ps.StartSubscribers("test-subscription")
	if err != nil {
		t.Fatalf("Should have started up subscription: err: %v", err)
	}
	err = ps.StartPublishers()
	if err != nil {
		t.Fatalf("Should have started publishers up: err: %v", err)
	}

	numResults := 0
	resultChan := make(chan bool)

	go func() {
		select {
		case <-ps.SubscribeChan:
			resultChan <- true
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ps.Publish(&utils.GooglePubSubMsg{
			Topic:   "test-topic",
			Payload: "payloadvalue",
		})
	}()

	select {
	case <-resultChan:
		numResults++
	}

	if numResults == 0 {
		t.Errorf("Should have received a messages from pub sub")
	}

	err = ps.StopPublishers()
	if err != nil {
		t.Fatalf("Should have stopped publishers: err: %v", err)
	}
	err = ps.StopSubscribers()
	if err != nil {
		t.Fatalf("Should have stopped publishers: err: %v", err)
	}
	err = ps.DeleteSubscription("test-subscription")
	if err != nil {
		t.Errorf("Should have deleted the subscription")
	}
	err = ps.DeleteTopic("test-topic")
	if err != nil {
		t.Errorf("Should have deleted the test-topic")
	}
}
