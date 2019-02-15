package jobs

import (
	"errors"
	"sync"
)

var (
	// ErrJobDoesNotExist is thrown when retrieving a job that does not exist
	ErrJobDoesNotExist = errors.New("Job does not exist")
	// ErrJobAlreadyExists is thrown when trying to submit a job with the same ID as one previously submitted
	ErrJobAlreadyExists = errors.New("Job with this ID already exists")

	statusRunning     = "running"
	statusComplete    = "complete"
	statusInitialized = "initialized"
)

// Subscription receives job updates
type Subscription struct {
	JobID        string
	SubscriberID string
	Updates      chan string
}

// Job is a reference to a running job
type Job struct {
	ID        string
	status    string
	observers map[string]*Subscription
	updates   chan string
	work      func(chan<- string)
	mu        sync.Mutex
}

// NewJob creates a new job instance
func NewJob(jobID string, work func(chan<- string)) *Job {
	job := &Job{
		ID:        jobID,
		observers: map[string]*Subscription{},
		work:      work,
		updates:   make(chan string),
		status:    statusInitialized,
	}
	return job
}

// GetStatus returns the status of the job
func (j *Job) GetStatus() string {
	j.mu.Lock()
	defer j.mu.Unlock()
	return j.status
}

func (j *Job) setStatus(status string) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.status = status
}

// Start begins working on a job
func (j *Job) Start() {

	// start doing the work
	go func() {
		j.setStatus(statusRunning)
		j.work(j.updates)
		j.setStatus(statusComplete)
		close(j.updates)
	}()

	go func() {
		// copy all updates to all subscribers
		for msg := range j.updates {
			j.mu.Lock()
			for _, subscription := range j.observers {
				subscription.Updates <- msg
			}
			j.mu.Unlock()
		}
		// work is complete, so close all of the observers
		j.mu.Lock()
		for _, observer := range j.observers {
			close(observer.Updates)
			delete(j.observers, observer.SubscriberID)
		}
		j.mu.Unlock()
	}()

}

// Subscribe creates a channel that provides updates to the work
func (j *Job) Subscribe() *Subscription {
	id := randString(10)
	sub := &Subscription{
		JobID:        j.ID,
		SubscriberID: id,
		Updates:      make(chan string),
	}

	j.mu.Lock()
	j.observers[id] = sub
	j.mu.Unlock()
	return sub
}

// Unsubscribe removes a subscription from the broadcast
func (j *Job) Unsubscribe(receipt *Subscription) {
	subscription := j.observers[receipt.SubscriberID]
	close(subscription.Updates)
	j.mu.Lock()
	delete(j.observers, receipt.SubscriberID)
	j.mu.Unlock()
}

// WaitForFinish subscribes to job updates, blocks and returns when the update
// channel is closed, indicating the job is completed
func (j *Job) WaitForFinish() {
	s := j.Subscribe()
	for range s.Updates {
		// Waiting here for the channel to close
	}
}
