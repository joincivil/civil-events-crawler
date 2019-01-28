// +build integration

// This is an integration test file for pubsub.go. Pubsub simulator needs to be running.
// Run this using go test -tags=integration
package pubsub_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/pubsub"
	"github.com/joincivil/go-common/pkg/generated/contract"
	ctime "github.com/joincivil/go-common/pkg/time"
	"math/big"
	"os"
	"testing"
	"time"
)

const (
	topicName = "testTopic"
	subName   = "testSubscription"
	projectID = "civil-media"
)

func setupCrawlerPubSub(t *testing.T) *pubsub.CrawlerPubSub {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8042")
	ps, err := pubsub.NewCrawlerPubSub(projectID, topicName)
	if err != nil {
		t.Errorf("Error starting pubsub %v", err)
	}
	return ps
}

func returnTestEvent(t *testing.T) *model.Event {
	contractAddress := "0x77e5aaBddb760FBa989A1C4B2CDd4aA8Fa3d311d"
	testAddress := "0xDFe273082089bB7f70Ee36Eebcde64832FE97E55"
	testEvent := &contract.CivilTCRContractApplication{
		ListingAddress: common.HexToAddress(testAddress),
		Deposit:        big.NewInt(1000),
		AppEndDate:     big.NewInt(1653860896),
		Data:           "DATA",
		Applicant:      common.HexToAddress(testAddress),
		Raw: types.Log{
			Address: common.HexToAddress(testAddress),
			Topics: []common.Hash{
				common.HexToHash("0x09cd8dcaf170a50a26316b5fe0727dd9fb9581a688d65e758b16a1650da65c0b"),
				common.HexToHash("0x0000000000000000000000002652c60cf04bbf6bb6cc8a5e6f1c18143729d440"),
				common.HexToHash("0x00000000000000000000000025bf9a1595d6f6c70e6848b60cba2063e4d9e552"),
			},
			Data:        []byte("thisisadatastring"),
			BlockNumber: 8888888,
			TxHash:      common.Hash{},
			TxIndex:     2,
			BlockHash:   common.Hash{},
			Index:       2,
			Removed:     false,
		},
	}
	event, err := model.NewEventFromContractEvent(
		"Application",
		"CivilTCRContract",
		common.HexToAddress(contractAddress),
		testEvent,
		ctime.CurrentEpochSecsInInt64(),
		model.Filterer,
	)
	if err != nil {
		t.Errorf("Error creating new event %v", err)
	}
	return event
}

func TestBuildMessage(t *testing.T) {
	cps := setupCrawlerPubSub(t)
	event := returnTestEvent(t)
	message, err := cps.BuildMessage(true, event.ContractAddress().Hex())
	if err != nil {
		t.Errorf("Error building message for pubsub %v", err)
	}
	if message.Payload != fmt.Sprintf("{\"newsroomException\":true,\"contractAddress\":\"%s\"}", event.ContractAddress().Hex()) {
		t.Errorf("Message payload contents are wrong %v", message.Payload)
	}
	if message.Topic != "testTopic" {
		t.Errorf("Message topic name is wrong %v", message.Topic)
	}
}

func TestPublishMessages(t *testing.T) {
	cps := setupCrawlerPubSub(t)
	te, err := cps.GooglePubsub.TopicExists(topicName)
	if err != nil {
		t.Errorf("Error checking if topic exists %v", err)
	}
	if te {
		err := cps.GooglePubsub.DeleteTopic(topicName)
		if err != nil {
			t.Errorf("Should have deleted existing topic")
		}
	}
	err = cps.GooglePubsub.CreateTopic(topicName)
	if err != nil {
		t.Errorf("Should have created a topic")
	}
	event := returnTestEvent(t)
	err = cps.StartPublishers()
	if err != nil {
		t.Errorf("Error starting publishers %v", err)
	}

	se, err := cps.GooglePubsub.SubscriptionExists(subName)
	if err != nil {
		t.Errorf("Error checking if subscription exists %v", err)
	}
	if se {
		err = cps.GooglePubsub.DeleteSubscription(subName)
		if err != nil {
			t.Errorf("Should have deleted existing subscription")
		}
	}

	err = cps.GooglePubsub.CreateSubscription(topicName, subName)
	if err != nil {
		t.Errorf("Error creating subscription %v", err)
	}
	err = cps.GooglePubsub.StartSubscribers(subName)
	if err != nil {
		t.Errorf("Error starting subscribers %v", err)
	}

	resultIDs := []string{}
	resultChan := make(chan string)

	go func() {
		for {
			select {
			case msg, ok := <-cps.GooglePubsub.SubscribeChan:
				if !ok {
					t.Errorf("Sending on closed channel")
				}
				resultChan <- string(msg.ID)
			}
		}
	}()

	go func() {
		time.Sleep(10)
		cps.PublishTriggerMessage()
		cps.PublishNewsroomExceptionMessage(event.ContractAddress().Hex())

	}()

Loop:
	for {
		select {
		case id, ok := <-resultChan:
			if !ok {
				t.Error("Sending on closed channel")
			}
			resultIDs = append(resultIDs, id)
		case <-time.After(20 * time.Second):
			break Loop
		}
	}

	if len(resultIDs) != 2 {
		t.Errorf("Should have seen two messages but only saw %v", resultIDs)
	}

}
