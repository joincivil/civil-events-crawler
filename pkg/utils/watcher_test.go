package utils

import (
	"errors"
	"testing"
	"time"
)

func TestWatcherSubscriptionUnsubFirst(t *testing.T) {
	func1 := NewWatcherSubscription(func(quit <-chan struct{}) error {
		time.Sleep(time.Second * 3)
		return nil
	})

	fs := func1.(*funcSub)
	fs.mu.Lock()
	if fs.unsubscribed {
		t.Errorf("Should not have been set to unsubscribed")
	}
	fs.mu.Unlock()

	func1.Unsubscribe()
	fs.mu.Lock()
	if !fs.unsubscribed {
		t.Errorf("Should have been set to unsubscribed")
	}
	fs.mu.Unlock()
}

func TestWatcherSubscriptionUnsubAfter(t *testing.T) {
	func1 := NewWatcherSubscription(func(quit <-chan struct{}) error {
		return nil
	})

	fs := func1.(*funcSub)
	fs.mu.Lock()
	if fs.unsubscribed {
		t.Errorf("Should not have been set to unsubscribed")
	}
	fs.mu.Unlock()

	time.Sleep(time.Second * 3)

	func1.Unsubscribe()
	fs.mu.Lock()
	if !fs.unsubscribed {
		t.Errorf("Should have been set to unsubscribed")
	}
	fs.mu.Unlock()
}

func TestWatcherSubscriptionError(t *testing.T) {
	func1 := NewWatcherSubscription(func(quit <-chan struct{}) error {
		time.Sleep(2 * time.Second)
		return errors.New("this is an error")
	})

	select {
	case <-func1.Err():
	case <-time.After(5 * time.Second):
		t.Errorf("should have received an error")
	}
}
