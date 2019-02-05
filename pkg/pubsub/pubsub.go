package pubsub

import (
	"encoding/json"
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
	NewsroomException bool   `json:"newsroomException"`
	ContractAddress   string `json:"contractAddress"`
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
func (c *CrawlerPubSub) BuildMessage(newsroomException bool,
	contractAddress string) (*cpubsub.GooglePubSubMsg, error) {
	msg := CrawlerPubSubMessage{
		NewsroomException: newsroomException,
		ContractAddress:   contractAddress,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return &cpubsub.GooglePubSubMsg{Topic: c.Topic, Payload: string(msgBytes)}, nil
}

// PublishProcessorTriggerMessage publishes an empty message to process events
func (c *CrawlerPubSub) PublishProcessorTriggerMessage() error {
	msg, err := c.BuildMessage(false, "")
	if err != nil {
		return err
	}
	return c.GooglePubsub.Publish(msg)
}

// PublishNewsroomExceptionMessage sends a message to pubsub to get all past events for a newsroom
func (c *CrawlerPubSub) PublishNewsroomExceptionMessage(contractAddress string) error {
	msg, err := c.BuildMessage(true, contractAddress)
	if err != nil {
		return err
	}
	return c.GooglePubsub.Publish(msg)
}
