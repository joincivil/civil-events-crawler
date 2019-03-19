package utils

import (
	"sync"
	"time"
)

const (
	completeTimeoutSecs = 5
)

// WatcherSubscription is represents a subscription to a stream of watcher events.
// This is a modification/custom version of the Subscription struct in the go-ethereum package,
// so much credit to the authors there.
// https://github.com/ethereum/go-ethereum/blob/master/event/subscription.go
type WatcherSubscription interface {
	Err() <-chan error
	Unsubscribe()
}

// NewWatcherSubscription runs a producer function as a watcher subscription in a new
// goroutine. The channel given to the producer is closed when Unsubscribe is called.  Returns
// error returned by producer to the subscription error channel.
// This is a modification/custom version of the NewSubscription struct in the go-ethereum package,
// so much credit to the authors there.
// https://github.com/ethereum/go-ethereum/blob/master/event/subscription.go
// Differs in how it handles completion of the producer function with separate channel to prevent
// deadlocks.
func NewWatcherSubscription(watcherName string, producer func(<-chan struct{}) error) WatcherSubscription {
	s := &funcSub{quit: make(chan struct{}), err: make(chan error, 1), complete: make(chan struct{})}
	s.watcherName = watcherName
	go func() {
		defer close(s.err)
		err := producer(s.quit)
		s.mu.Lock()
		defer s.mu.Unlock()
		if !s.unsubscribed {
			if err != nil {
				s.err <- err
			}
			s.unsubscribed = true
		}
		close(s.complete)
		s.muComp.Lock()
		s.complete = nil
		s.muComp.Unlock()
	}()
	return s
}

type funcSub struct {
	quit         chan struct{}
	complete     chan struct{}
	err          chan error
	mu           sync.Mutex
	muComp       sync.Mutex
	unsubscribed bool
	watcherName  string
}

func (s *funcSub) Unsubscribe() {
	s.mu.Lock()
	if s.unsubscribed {
		s.mu.Unlock()
		return
	}
	s.unsubscribed = true
	close(s.quit)
	s.mu.Unlock()

	// Wait for producer shutdown.
	s.muComp.Lock()
	if s.complete != nil {
		select {
		case <-s.complete:
		case <-time.After(time.Second * time.Duration(completeTimeoutSecs)):
		}
	}
	s.muComp.Unlock()
}

func (s *funcSub) Err() <-chan error {
	return s.err
}
