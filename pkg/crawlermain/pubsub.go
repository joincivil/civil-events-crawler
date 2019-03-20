package crawlermain

import (
	"os"

	log "github.com/golang/glog"

	"github.com/joincivil/civil-events-crawler/pkg/pubsub"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

func crawlerPubSub(config *utils.CrawlerConfig) *pubsub.CrawlerPubSub {
	if config.PubSubProjectID == "" || config.PubSubTopicName == "" {
		return nil
	}

	pubsub, err := pubsub.NewCrawlerPubSub(config.PubSubProjectID, config.PubSubTopicName)
	if err != nil {
		log.Errorf("Error initializing pubsub, stopping...; err: %v, %v, %v", err, config.PubSubProjectID, config.PubSubTopicName)
		os.Exit(1)
	}
	topicExists, err := pubsub.GooglePubsub.TopicExists(config.PubSubTopicName)
	if err != nil {
		log.Errorf("Error checking for existence of topic: err: %v", err)
		os.Exit(1)
	}
	if !topicExists {
		log.Errorf("Topic: %v does not exist", config.PubSubTopicName)
		os.Exit(1)
	}
	err = pubsub.StartPublishers()
	if err != nil {
		log.Errorf("Error starting publishers, stopping...; err: %v", err)
		os.Exit(1)
	}
	return pubsub
}
