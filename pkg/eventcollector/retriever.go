package eventcollector

import (
	"time"

	log "github.com/golang/glog"
	"github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"

	specs "github.com/joincivil/civil-events-crawler/pkg/contractspecs"
	"github.com/joincivil/civil-events-crawler/pkg/generated/filterer"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
)

const (
	pqUniqueViolationCode = "23505"
	pqHashKeyConstraint   = "event_hash_key"
)

// AddFilterers will add filterer to the embedded retriever.
func (c *EventCollector) AddFilterers(w model.ContractFilterers) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	c.filterers = append(c.filterers, w)
	return nil
}

// RemoveFilterers will remove given filterer from the embedded retriever.
func (c *EventCollector) RemoveFilterers(w model.ContractFilterers) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	if c.filterers != nil && len(c.filterers) > 0 {
		for index, ew := range c.filterers {
			if w.ContractAddress() == ew.ContractAddress() &&
				w.ContractName() == ew.ContractName() {
				// Delete the item in the filterers list.
				copy(c.filterers[index:], c.filterers[index+1:])
				c.filterers[len(c.filterers)-1] = nil
				c.filterers = c.filterers[:len(c.filterers)-1]
				return nil
			}
		}
	}
	return nil
}

// FilterAddedNewsroomContract runs a filterer on the newly watched newsroom contract to ensure we have all events.
func (c *EventCollector) FilterAddedNewsroomContract(newsroomAddr common.Address) ([]*model.Event, error) {
	nwsrmFilterer := filterer.NewNewsroomContractFilterers(newsroomAddr)
	c.updateFiltererStartingBlock(nwsrmFilterer)
	retrieve := retriever.NewEventRetriever(c.httpClient, []model.ContractFilterers{nwsrmFilterer})
	err := retrieve.Retrieve(false)
	if err != nil {
		return nil, err
	}
	nwsrmEvents := retrieve.PastEvents
	return nwsrmEvents, nil
}

func (c *EventCollector) getExistingNewsroomFilterers() map[common.Address]bool {
	existingNewsroomAddr := map[common.Address]bool{}
	c.mutex.Lock()
	for _, existing := range c.filterers {
		specs, _ := specs.ContractTypeToSpecs.Get(specs.NewsroomContractType)
		if existing.ContractName() == specs.Name() {
			existingNewsroomAddr[existing.ContractAddress()] = true
		}
	}
	c.mutex.Unlock()
	return existingNewsroomAddr
}

// FilterAddedMultiSigContract runs a filterer on the newly watched multisig contract to ensure we have all events.
func (c *EventCollector) FilterAddedMultiSigContract(multiSigAddr common.Address) ([]*model.Event, error) {
	multiSigFilterer := filterer.NewMultiSigWalletContractFilterers(multiSigAddr)
	c.updateFiltererStartingBlock(multiSigFilterer)
	retrieve := retriever.NewEventRetriever(c.httpClient, []model.ContractFilterers{multiSigFilterer})
	err := retrieve.Retrieve(false)
	if err != nil {
		return nil, err
	}
	multiSigEvents := retrieve.PastEvents
	return multiSigEvents, nil
}

func (c *EventCollector) getExistingMultiSigFilterers() map[common.Address]bool {
	existingMultiSigAddr := map[common.Address]bool{}
	c.mutex.Lock()
	for _, existing := range c.filterers {
		specs, _ := specs.ContractTypeToSpecs.Get(specs.MultiSigWalletContractType)
		if existing.ContractName() == specs.Name() {
			existingMultiSigAddr[existing.ContractAddress()] = true
		}
	}
	c.mutex.Unlock()
	return existingMultiSigAddr
}

func (c *EventCollector) pollingIntSecs() int {
	intSecs := c.pollingIntervalSecs
	if intSecs == 0 {
		intSecs = defaultPollingIntervalSecs
	}
	return intSecs
}

func (c *EventCollector) retrieveEvents(filterers []model.ContractFilterers,
	nonSubOnly bool) (*retriever.EventRetriever, error) {
	c.updateRetrieverStartingBlocks(filterers)
	r := retriever.NewEventRetriever(c.httpClient, filterers)
	err := r.Retrieve(nonSubOnly)
	if err != nil {
		return nil, errors.Wrap(err, "retrieveEvents.Retrieve")
	}
	return r, nil
}

// isAllowedErrRetriever returns if an error should be ignored or not in the
// filterers. This is used in the eventcollector to ensure we only fail on
// particular errors and recover on others.
// ex. if an event hash already exists, we ignore, since that would be "correct" as
// sometimes we may receive the same event and do not want to save it again.
func (c *EventCollector) isAllowedErrRetriever(err error) bool {
	switch causeErr := errors.Cause(err).(type) {
	case *pq.Error:
		// log.Infof("*pq error code %v: %v, constraint: %v, msg: %v", causeErr.Code,
		// 	causeErr.Code.Name(), causeErr.Constraint, causeErr.Message)
		if causeErr.Code == pqUniqueViolationCode &&
			causeErr.Constraint == pqHashKeyConstraint {
			return true
		}
	case pq.Error:
		// log.Infof("pq error code %v: %v, constraint: %v, msg: %v", causeErr.Code,
		// 	causeErr.Code.Name(), causeErr.Constraint, causeErr.Message)
		if causeErr.Code == pqUniqueViolationCode &&
			causeErr.Constraint == pqHashKeyConstraint {
			return true
		}
	default:
		log.Infof("not allowed error type: %T", causeErr)
	}
	return false
}

// persistRetrieverLastBlockData saves the last seen events for each filter to
// persistence. Returns the last error seen when updating the block data.
func (c *EventCollector) persistRetrieverLastBlockData() error {
	var err error
	for _, filter := range c.filterers {
		err = c.retrieverPersister.UpdateLastBlockData(filter.LastEvents())
	}
	return err
}

// UpdateStartingBlocks updates starting blocks for retriever based on persistence
func (c *EventCollector) updateRetrieverStartingBlocks(filterers []model.ContractFilterers) {
	for _, filter := range filterers {
		c.updateFiltererStartingBlock(filter)
	}
}

func (c *EventCollector) updateFiltererStartingBlock(filter model.ContractFilterers) {
	contractAddress := filter.ContractAddress()
	eventTypes := filter.EventTypes()
	for _, eventType := range eventTypes {
		lastBlock := c.retrieverPersister.LastBlockNumber(eventType, contractAddress)
		// If lastBlock is 0, assume it has never been set, so set to default
		// start block value.
		if lastBlock == 0 {
			lastBlock = c.startBlock
		}
		// NOTE (IS): Starting at lastBlock+1. There could be a scenario where this could miss the rest of events in prev block?
		filter.UpdateStartBlock(eventType, lastBlock+1)
	}
}

func (c *EventCollector) runRetrieverLoop(pollingOnly bool) error {
	// Always initially retrieve all the enabled events
	err := c.runRetriever(false)
	if err != nil {
		if !c.isAllowedErrRetriever(err) {
			log.Errorf("Error running retriever: err: %v", err)
			return errors.Wrap(err, "error running initial retriever in loop")
		}
		log.Errorf("Error running retriever, recovering: err: %v", err)
		c.errRep.Error(err, nil)
	}

	// Startup the main polling loop.  If polling only, keep retrieving all events.  If
	// not polling only, only retrieve non-subscribed/watched events.
	go func() {
		for {
			log.Infof("Waiting for next retriever run")
			select {
			case <-c.retrieverStopChan:
				return
			case <-c.shutdownChan:
				return
			case <-time.After(time.Duration(c.pollingIntSecs()) * time.Second):
			}

			err := c.runRetriever(!pollingOnly)
			if err != nil {
				if !c.isAllowedErrRetriever(err) {
					log.Errorf("Error running retriever: err: %v", errors.Cause(err))
					c.errRep.Error(errors.Cause(err), nil)
					c.retrieverErrChan <- errors.Wrap(err, "error running retriever in startcol")
				} else {
					log.Errorf("Recoverable error running retriever: err: %v", errors.Cause(err))
					c.errRep.Error(errors.Cause(err), nil)
				}
			}

			if c.crawlerPubSub != nil {
				err = c.crawlerPubSub.PublishProcessorTriggerMessage()
				if err != nil {
					log.Errorf("Error publishing trigger message: err: %v", err)
					c.errRep.Error(errors.Cause(err), nil)
				}
			}
		}
	}()

	return nil
}

func (c *EventCollector) runRetriever(nonSubOnly bool) error {
	log.Infof("runRetriever: retrieving events")
	ret, err := c.retrieveEvents(c.filterers, nonSubOnly)
	if err != nil {
		return errors.WithMessage(err, "error retrieving events")
	}
	pastEvents := ret.PastEvents

	log.Infof("runRetriever: checking events for newsrooms")
	additionalEvents, err := c.CheckRetrievedEventsForNewsroom(pastEvents)
	if err != nil {
		return errors.WithMessage(err, "error checking newsroom events during filterer")
	}

	if len(additionalEvents) > 0 {
		pastEvents = append(pastEvents, additionalEvents...)
	}

	log.Infof("runRetriever: sorting events by block")
	err = ret.SortEventsByBlock(pastEvents)
	if err != nil {
		return errors.WithMessage(err, "error sorting retrieved events")
	}

	log.Infof("runRetriever: updating event times from headers")
	err = c.updateEventTimesFromBlockHeaders(pastEvents)
	if err != nil {
		return errors.WithMessage(err, "error updating dates for events")
	}

	log.Infof("runRetriever: saving events")
	errs := c.eventDataPersister.SaveEvents(pastEvents)
	if len(errs) > 0 {
		for _, err := range errs {
			if !c.isAllowedErrRetriever(err) {
				return errors.WithMessage(err, "error persisting events")
			}
			log.Errorf("Error persisting events, recovering: err: %v", err)
			c.errRep.Error(err, nil)
		}
	}

	log.Infof("runRetriever: persist retriever last block data")
	err = c.persistRetrieverLastBlockData()
	if err != nil {
		return errors.WithMessage(err, "error persisting last block data")
	}

	log.Infof("runRetriever: done")
	return nil
}
