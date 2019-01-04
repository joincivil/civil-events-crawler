package pubsub

import (
	"context"
	"errors"
	"os"
	"sync"
	"time"

	log "github.com/golang/glog"

	"cloud.google.com/go/pubsub"
)

const (
	googleCredsEnvVarName    = "GOOGLE_APPLICATION_CREDENTIALS"
	googleEmulatorEnvVarName = "PUBSUB_EMULATOR_HOST"
)

// NewPubSubClient creates a new pubsub client via Google project ID
func NewPubSubClient(projectID string) (*pubsub.Client, *context.Context, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, nil, err
	}
	return client, &ctx, nil
}

// NewGooglePubSub returns a new GooglePubSub struct
func NewGooglePubSub(projectID string) (*GooglePubSub, error) {
	googleCredsEnvVar := os.Getenv(googleCredsEnvVarName)
	emulatorEnvVar := os.Getenv(googleEmulatorEnvVarName)
	if googleCredsEnvVar == "" && emulatorEnvVar == "" {
		return nil, errors.New("Required Google envvars do not appear to be set")
	}
	client, ctx, err := NewPubSubClient(projectID)
	if err != nil {
		return nil, err
	}
	return &GooglePubSub{
		projectID: projectID,
		client:    client,
		ctx:       *ctx,
	}, nil
}

// GooglePubSubMsg represents a messages to be published
type GooglePubSubMsg struct {
	Topic   string
	Payload string
}

// GooglePubSub is a wrapper around handling for Google Pub/Sub.  Manages
// pooled goroutines of publishers and subscribers.
type GooglePubSub struct {
	projectID string
	client    *pubsub.Client
	ctx       context.Context

	publishChan       chan *GooglePubSubMsg
	publishKill       chan bool
	publishStarted    bool
	numRunningPublish int
	publishMutex      sync.Mutex

	SubscribeChan          chan *pubsub.Message
	subscribeConfig        SubscribeConfig
	subscribeContext       context.Context
	subscribeContextCancel context.CancelFunc
	subscribeStarted       bool
	subscribeMutex         sync.Mutex
	numRunningSubscribe    int
}

// SubscribeConfig is a config for the wrapper around a Google Pubsub Subscription
type SubscribeConfig struct {
	Name    string
	AutoAck bool
}

// StartPublishers starts up a pool of PubSub publishers.
func (g *GooglePubSub) StartPublishers() error {
	g.publishMutex.Lock()
	defer g.publishMutex.Unlock()
	g.publishChan = make(chan *GooglePubSubMsg)
	g.publishKill = make(chan bool)
	g.publishStarted = false
	g.numRunningPublish = 0

	// Keeping it to 1 goroutine for now as the pubsub lib has some concurrency
	// built in.
	// multiplier := 1
	// numRoutines := runtime.NumCPU() * multiplier
	numRoutines := 1

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < numRoutines; i++ {
			wg.Add(1)
			go g.publisher(&wg)
			log.Info("Publisher started")
		}
		g.publishMutex.Lock()
		g.publishStarted = true
		g.publishMutex.Unlock()
		log.Infof("All publishers started: num: %v", numRoutines)

		wg.Wait()
		log.Info("All publishers stopped")
	}()
	return nil
}

// CreateTopic creates a new pubsub topic
func (g *GooglePubSub) CreateTopic(topicName string) error {
	_, err := g.client.CreateTopic(g.ctx, topicName)
	if err != nil {
		log.Errorf("Failed to create topic: %v", err)
		return err
	}
	return err
}

// DeleteTopic deletes an existing pubsub topic
func (g *GooglePubSub) DeleteTopic(topicName string) error {
	topic := g.client.Topic(topicName)
	err := topic.Delete(g.ctx)
	if err != nil {
		log.Errorf("Failed to delete topic: %v, err: %v", topicName, err)
		return err
	}
	return nil
}

// TopicExists checks the existence of a topic
func (g *GooglePubSub) TopicExists(topicName string) (bool, error) {
	topic := g.client.Topic(topicName)
	ok, err := topic.Exists(g.ctx)
	if err != nil {
		log.Errorf("Failed on existence check of topic: %v, err: %v", topicName, err)
		return false, err
	}
	return ok, nil
}

// Publish publishes the given message to the pubsub.
func (g *GooglePubSub) Publish(msg *GooglePubSubMsg) error {
	g.publishChan <- msg
	return nil
}

// PublishersStarted returns true if the publishers are running, false if not.
func (g *GooglePubSub) PublishersStarted() bool {
	g.publishMutex.Lock()
	pubStarted := g.publishStarted
	g.publishMutex.Unlock()
	return pubStarted
}

// NumPublishersRunning return the number of publishers goroutines running.
func (g *GooglePubSub) NumPublishersRunning() int {
	g.publishMutex.Lock()
	numPubs := g.numRunningPublish
	g.publishMutex.Unlock()
	return numPubs
}

// StopPublishers will stop the publisher goroutines
func (g *GooglePubSub) StopPublishers() error {
	for {
		g.publishMutex.Lock()
		if g.numRunningPublish == 0 {
			g.publishMutex.Unlock()
			break
		}
		g.publishMutex.Unlock()
		g.publishKill <- true
		time.Sleep(1 * time.Second)
	}
	close(g.publishKill)
	close(g.publishChan)
	g.publishStarted = false
	return nil
}

// StartSubscribersWithConfig starts up a pool of PubSub publishers.
func (g *GooglePubSub) StartSubscribersWithConfig(config SubscribeConfig) error {
	ok, err := g.SubscriptionExists(config.Name)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("Subscription doesn't exist")
	}

	ctx := context.Background()
	g.subscribeConfig = config
	g.subscribeContext, g.subscribeContextCancel = context.WithCancel(ctx)
	g.SubscribeChan = make(chan *pubsub.Message)
	g.subscribeStarted = false
	g.numRunningSubscribe = 0

	// Keeping it to 1 goroutine for now as the pubsub lib has some concurrency
	// built in.
	// multiplier := 1
	// numRoutines := runtime.NumCPU() * multiplier
	numRoutines := 1

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < numRoutines; i++ {
			wg.Add(1)
			go g.subscriber(&wg)
			log.Info("Subscriber started")
		}
		g.subscribeMutex.Lock()
		g.subscribeStarted = true
		g.subscribeMutex.Unlock()
		log.Infof("All subscribers started: num: %v", numRoutines)

		wg.Wait()
		log.Info("All subscribers stopped")
	}()
	return nil
}

// StartSubscribers starts up a pool of PubSub publishers using a default config
func (g *GooglePubSub) StartSubscribers(subscriptionName string) error {
	return g.StartSubscribersWithConfig(
		SubscribeConfig{
			Name:    subscriptionName,
			AutoAck: true,
		},
	)
}

// CreateSubscriptionWithConfig creates a new subscription
func (g *GooglePubSub) CreateSubscriptionWithConfig(topicName string, subName string, config pubsub.SubscriptionConfig) error {
	topic := g.client.Topic(topicName)
	config.Topic = topic
	_, err := g.client.CreateSubscription(g.ctx, subName, config)
	if err != nil {
		log.Errorf("Failed to create subscription: %v", err)
		return err
	}
	return err
}

// CreateSubscription creates a new subscription with a default config
func (g *GooglePubSub) CreateSubscription(topicName string, subName string) error {
	subConfig := pubsub.SubscriptionConfig{
		AckDeadline: 10 * time.Second,
	}
	return g.CreateSubscriptionWithConfig(topicName, subName, subConfig)
}

// DeleteSubscription deletes an existing subscription
func (g *GooglePubSub) DeleteSubscription(subName string) error {
	sub := g.client.Subscription(subName)
	err := sub.Delete(g.ctx)
	if err != nil {
		log.Errorf("Failed to delete subscription: %v", err)
		return err
	}
	return nil
}

// SubscriptionExists checks for existence of an existing subscription
func (g *GooglePubSub) SubscriptionExists(subName string) (bool, error) {
	sub := g.client.Subscription(subName)
	ok, err := sub.Exists(g.ctx)
	if err != nil {
		log.Errorf("Failed to check existence of subscription: %v", err)
		return false, err
	}
	return ok, nil
}

// SubscribersStarted returns true if the subscribers are running, false if not.
func (g *GooglePubSub) SubscribersStarted() bool {
	g.subscribeMutex.Lock()
	subStarted := g.subscribeStarted
	g.subscribeMutex.Unlock()
	return subStarted
}

// NumSubscribersRunning return the number of subscriber goroutines running.
func (g *GooglePubSub) NumSubscribersRunning() int {
	g.subscribeMutex.Lock()
	numSubs := g.numRunningSubscribe
	g.subscribeMutex.Unlock()
	return numSubs
}

// SubscriptionName returns the name of the subscription to track.
func (g *GooglePubSub) SubscriptionName() string {
	return g.subscribeConfig.Name
}

// StopSubscribers will stop the subscriber goroutines
func (g *GooglePubSub) StopSubscribers() error {
	g.subscribeContextCancel()
	close(g.SubscribeChan)
	g.subscribeMutex.Lock()
	g.subscribeStarted = false
	g.subscribeMutex.Unlock()
	return nil
}

// publisher is a func meant to be run in a goroutine as a part of a pool
// of goroutines handling publishing to a pubsub.
func (g *GooglePubSub) publisher(wg *sync.WaitGroup) {
	defer wg.Done()

	topicNameToTopic := map[string]*pubsub.Topic{}
	g.publishMutex.Lock()
	g.numRunningPublish++
	g.publishMutex.Unlock()

Loop:
	for {
		select {
		case msg := <-g.publishChan:
			log.Infof("Message received for topic %v\n", msg.Topic)
			topic, ok := topicNameToTopic[msg.Topic]
			if !ok {
				topic = g.client.Topic(msg.Topic)
				topicNameToTopic[msg.Topic] = topic
			}

			payload := []byte(msg.Payload)
			res := topic.Publish(g.ctx, &pubsub.Message{Data: payload})

			id, err := res.Get(g.ctx)
			if err != nil {
				log.Errorf("Failed to publish: err: %v", err)
				continue
			}
			log.Infof("Message published to topic %v: id %v\n", msg.Topic, id)

		case <-g.publishKill:
			log.Info("Stopping publisher")
			break Loop
		}
	}

	// Stop the topics
	for _, topic := range topicNameToTopic {
		topic.Stop()
	}
	g.publishMutex.Lock()
	g.numRunningPublish--
	g.publishMutex.Unlock()
}

func (g *GooglePubSub) subscriber(wg *sync.WaitGroup) {
	defer wg.Done()

	g.subscribeMutex.Lock()
	g.numRunningSubscribe++
	g.subscribeMutex.Unlock()
	sub := g.client.Subscription(g.subscribeConfig.Name)
	err := sub.Receive(g.subscribeContext, func(ctx context.Context, msg *pubsub.Message) {
		log.Infof("Got message: %v: %v\n", msg.ID, msg)
		g.SubscribeChan <- msg
		if g.subscribeConfig.AutoAck {
			msg.Ack()
		}
	})
	g.subscribeMutex.Lock()
	g.numRunningSubscribe--
	g.subscribeMutex.Unlock()
	if err != nil {
		log.Errorf("Error with subscription: %v\n", err)
		return
	}
	log.Info("Subscriber stopped")
}
