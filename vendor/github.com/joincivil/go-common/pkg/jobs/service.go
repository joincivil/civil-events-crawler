package jobs

import (
	"math/rand"

	log "github.com/golang/glog"
)

// JobService interface defines what is needed to retrieve and persist jobs
type JobService interface {
	GetJob(id string) (*Job, error)
	StartJob(id string, work func(updates chan<- string)) (*Job, error)
	StopSubscription(receipt *Subscription) error
}

// InMemoryJobService is an implementation of JobService that only stays in memory
type InMemoryJobService struct {
	jobs map[string]*Job
}

// NewInMemoryJobService builds a new InMemoryJobService
func NewInMemoryJobService() *InMemoryJobService {
	return &InMemoryJobService{
		jobs: map[string]*Job{},
	}
}

// StartJob starts a new job
func (s *InMemoryJobService) StartJob(id string, work func(updates chan<- string)) (*Job, error) {

	job := s.jobs[id]
	if job != nil {
		return nil, ErrJobAlreadyExists
	}

	log.Infof("Starting job with ID %v", id)
	job = NewJob(id, work)
	s.jobs[id] = job

	job.Start()

	go func() {
		job.WaitForFinish()
		log.Infof("Job complete (%v), cleaning up", id)
		delete(s.jobs, id)
	}()

	return job, nil
}

// GetJob retrieves a Subscription for the given ID
func (s *InMemoryJobService) GetJob(id string) (*Job, error) {
	job := s.jobs[id]
	if job == nil {
		return nil, ErrJobDoesNotExist
	}

	return job, nil
}

// StartSubscription creates a subscription for a job
func (s *InMemoryJobService) StartSubscription(jobID string) (*Subscription, error) {
	job := s.jobs[jobID]
	if job == nil {
		return nil, ErrJobDoesNotExist
	}

	return job.Subscribe(), nil
}

// StopSubscription cancels a subscription for a job
func (s *InMemoryJobService) StopSubscription(receipt *Subscription) error {
	job := s.jobs[receipt.JobID]
	job.Unsubscribe(receipt)

	return nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
