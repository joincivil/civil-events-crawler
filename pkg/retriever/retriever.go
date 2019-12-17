// Package retriever contains all components for events retriever, which
// gets all past events
package retriever // import "github.com/joincivil/civil-events-crawler/pkg/retriever"

import (
	"runtime"
	"sort"
	"sync"
	"time"

	log "github.com/golang/glog"

	"github.com/Jeffail/tunny"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/joincivil/civil-events-crawler/pkg/model"
)

const (
	workerMultiplier = 2
)

// NewEventRetriever creates a EventRetriever given a list of ContractFilterers and
// connection to client
func NewEventRetriever(client bind.ContractBackend, filterers []model.ContractFilterers) *EventRetriever {
	retriever := &EventRetriever{
		client:     client,
		PastEvents: make([]*model.Event, 0),
		filterers:  filterers,
	}
	return retriever
}

// EventRetriever handles the iterator returned from retrieving past events
type EventRetriever struct {

	// Client is the ethereum client from go-ethereum
	client bind.ContractBackend

	// PastEvents is a slice that holds all past Events requested
	PastEvents []*model.Event

	// filterers contains a list of ContractFilterers
	filterers []model.ContractFilterers

	// Mutex to lock writes/reads to PastEvents
	pastEventsMutex sync.Mutex

	// Mutex to lock access to the filterers map
	mutex sync.Mutex
}

// Retrieve gets all events from StartBlock until now
// If nonSubOnly is true, retrieve only event types that are not subscribed to by
// the watchers.  This is useful when polling for events alongside setting up watchers
// for events.
func (r *EventRetriever) Retrieve(nonSubOnly bool) error {
	start := time.Now()
	numWorkers := runtime.NumCPU() * workerMultiplier

	// Worker pool to run the filterers
	pool := tunny.NewFunc(numWorkers, func(payload interface{}) interface{} {
		f := payload.(func())
		f()
		return nil
	})
	defer pool.Close()

	wg := sync.WaitGroup{}
	for _, filter := range r.filterers {
		wg.Add(1)

		go func(filt model.ContractFilterers) {
			filtererFunc := func() {
				log.Infof(
					"Starting filterer: %v, %v",
					filt.ContractName(),
					filt.ContractAddress().Hex(),
				)

				pastEvents := []*model.Event{}
				pastEvents, err := filt.StartFilterers(r.client, pastEvents, nonSubOnly)
				if err != nil {
					log.Errorf("Error retrieving filterer events: err: %v", err)
					return
				}

				r.pastEventsMutex.Lock()
				r.PastEvents = append(r.PastEvents, pastEvents...)
				r.pastEventsMutex.Unlock()
			}

			pool.Process(filtererFunc)
			wg.Done()
			log.Infof(
				"Completed filterer: %v, %v",
				filt.ContractName(),
				filt.ContractAddress().Hex(),
			)
		}(filter)
	}

	wg.Wait()
	log.Infof("All %v filterers have run, took %v", len(r.filterers), time.Since(start))
	return nil
}

// AddFilterers add filterers to the retriever
func (r *EventRetriever) AddFilterers(w model.ContractFilterers) error {
	defer r.mutex.Unlock()
	r.mutex.Lock()
	r.filterers = append(r.filterers, w)
	return nil
}

// RemoveFilterers remove given filterers from the retriever
func (r *EventRetriever) RemoveFilterers(w model.ContractFilterers) error {
	defer r.mutex.Unlock()
	r.mutex.Lock()
	if r.filterers != nil && len(r.filterers) > 0 {
		for index, ew := range r.filterers {
			if w.ContractAddress() == ew.ContractAddress() &&
				w.ContractName() == ew.ContractName() {
				// Delete the item in the filterers list.
				copy(r.filterers[index:], r.filterers[index+1:])
				r.filterers[len(r.filterers)-1] = nil
				r.filterers = r.filterers[:len(r.filterers)-1]
				return nil
			}
		}
	}
	return nil
}

// SortEventsByBlock sorts events in PastEvents by block number
// NOTE(IS): This is not optimal, but for now checking that values exist outside of sort
// Pass in nil if you want to sort retriever.PastEvents
func (r *EventRetriever) SortEventsByBlock(events []*model.Event) error {
	if events == nil {
		events = r.PastEvents
	}
	sort.Slice(events, func(i, j int) bool {
		blockNumber1 := events[i].BlockNumber()
		blockNumber2 := events[j].BlockNumber()
		return blockNumber1 < blockNumber2
	})
	return nil
}
