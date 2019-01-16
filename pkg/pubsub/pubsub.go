package pubsub

import (
	"encoding/json"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	cpubsub "github.com/joincivil/go-common/pkg/pubsub"
)

// CrawlerPubSub handles logic for crawler pubsub
type CrawlerPubSub struct {
	GooglePubsub *cpubsub.GooglePubSub
	Topic        string
}

// NewCrawlerPubSub creates a new CrawlerPubSub
func NewCrawlerPubSub(projectID string, topic string) (*CrawlerPubSub, error) {
	pubsub, err := cpubsub.NewGooglePubSub(projectID)
	if err != nil {
		return nil, err
	}
	return &CrawlerPubSub{GooglePubsub: pubsub, Topic: topic}, nil
}

// CrawlerPubSubMessage is the message sent from the crawler
type CrawlerPubSubMessage struct {
	Timestamp int64  `json:"timestamp"`
	Hash      string `json:"hash"`
}

// StartPublishers starts the publishers
func (c *CrawlerPubSub) StartPublishers() error {
	return c.GooglePubsub.StartPublishers()
}

// StopPublishers stops the publishers
func (c *CrawlerPubSub) StopPublishers() error {
	return c.GooglePubsub.StopPublishers()
}

// BuildMessage builds a message for the publisher
func (c *CrawlerPubSub) BuildMessage(timestamp int64, hash string) (*cpubsub.GooglePubSubMsg, error) {
	msg := CrawlerPubSubMessage{
		Timestamp: timestamp,
		Hash:      hash,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return &cpubsub.GooglePubSubMsg{Topic: c.Topic, Payload: string(msgBytes)}, nil
}

//BuildFilteringFinishedMessage is the message sent when filtering is done. It has no timestamp or hash
// because there are a bunch of events associated with this.
func (c *CrawlerPubSub) BuildFilteringFinishedMessage() (*cpubsub.GooglePubSubMsg, error) {
	return c.BuildMessage(0, "")
}

// PublishFilteringFinishedMessage sends a message to pubsub that filtering has finished
func (c *CrawlerPubSub) PublishFilteringFinishedMessage() error {
	msg, err := c.BuildFilteringFinishedMessage()
	if err != nil {
		return err
	}
	return c.GooglePubsub.Publish(msg)
}

// PublishWatchedEventMessage sends a message that an event has been watched for
func (c *CrawlerPubSub) PublishWatchedEventMessage(event *model.Event) error {
	msg, err := c.BuildMessage(event.Timestamp(), event.Hash())
	if err != nil {
		return err
	}
	return c.GooglePubsub.Publish(msg)
}
