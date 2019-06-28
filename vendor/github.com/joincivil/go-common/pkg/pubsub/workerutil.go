package pubsub

import (
	"errors"
)

func initPubSub(projectID string) (*GooglePubSub, error) {
	// If no project ID, quit
	if projectID == "" {
		return nil, errors.New("Need PubSubProjectID")
	}

	ps, err := NewGooglePubSub(projectID)
	if err != nil {
		return nil, err
	}
	return ps, err
}

func initPubSubSubscribers(ps *GooglePubSub, topicName string, subName string) error {
	// If no crawl topic name, quit
	if topicName == "" {
		return errors.New("Pubsub topic name should be specified")
	}
	// If no subscription name, quit
	if subName == "" {
		return errors.New("Pubsub subscription name should be specified")
	}

	return ps.StartSubscribersWithConfig(
		SubscribeConfig{
			Name:    subName,
			AutoAck: false,
		},
	)
}
