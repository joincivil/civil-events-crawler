// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2019-12-16 23:06:44.902149 +0000 UTC
package filterer

import (
	log "github.com/golang/glog"
	"runtime"
	"sync"

	"github.com/Jeffail/tunny"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	specs "github.com/joincivil/civil-events-crawler/pkg/contractspecs"
	commongen "github.com/joincivil/civil-events-crawler/pkg/generated/common"
	"github.com/joincivil/civil-events-crawler/pkg/model"

	"github.com/joincivil/go-common/pkg/generated/contract"
	ctime "github.com/joincivil/go-common/pkg/time"

	"math/big"
)

func NewNewsroomContractFilterers(contractAddress common.Address) *NewsroomContractFilterers {
	contractFilterers := &NewsroomContractFilterers{
		contractAddress:   contractAddress,
		eventTypes:        commongen.EventTypesNewsroomContract(),
		eventToStartBlock: make(map[string]uint64),
		lastEvents:        make([]*model.Event, 0),
	}
	for _, eventType := range contractFilterers.eventTypes {
		contractFilterers.eventToStartBlock[eventType] = 0
	}
	return contractFilterers
}

type NewsroomContractFilterers struct {
	contractAddress   common.Address
	contract          *contract.NewsroomContract
	eventTypes        []string
	eventToStartBlock map[string]uint64
	lastEvents        []*model.Event
	lastEventsMutex   sync.Mutex
	pastEventsMutex   sync.Mutex
}

func (f *NewsroomContractFilterers) ContractName() string {
	return "NewsroomContract"
}

func (f *NewsroomContractFilterers) ContractAddress() common.Address {
	return f.contractAddress
}

func (f *NewsroomContractFilterers) StartFilterers(client bind.ContractBackend,
	pastEvents []*model.Event, nonSubOnly bool) ([]*model.Event, error) {
	return f.StartNewsroomContractFilterers(client, pastEvents, nonSubOnly)
}

func (f *NewsroomContractFilterers) EventTypes() []string {
	return f.eventTypes
}

func (f *NewsroomContractFilterers) UpdateStartBlock(eventType string, startBlock uint64) {
	f.eventToStartBlock[eventType] = startBlock
}

func (f *NewsroomContractFilterers) LastEvents() []*model.Event {
	return f.lastEvents
}

// StartNewsroomContractFilterers retrieves events for NewsroomContract
func (f *NewsroomContractFilterers) StartNewsroomContractFilterers(client bind.ContractBackend,
	pastEvents []*model.Event, nonSubOnly bool) ([]*model.Event, error) {
	contract, err := contract.NewNewsroomContract(f.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartNewsroomContract: err: %v", err)
		return pastEvents, err
	}
	f.contract = contract

	workerMultiplier := 1
	numWorkers := runtime.NumCPU() * workerMultiplier
	log.Infof("Filter worker #: %v", numWorkers)
	pool := tunny.NewFunc(numWorkers, func(payload interface{}) interface{} {
		f := payload.(func())
		f()
		return nil
	})
	defer pool.Close()

	wg := sync.WaitGroup{}
	resultsChan := make(chan []*model.Event)
	done := make(chan struct{})
	filtsRun := 0

	if !specs.IsEventDisabled("NewsroomContract", "ContentPublished") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("NewsroomContract", "ContentPublished")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["ContentPublished"]
				pevents, e := f.startFilterContentPublished(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving ContentPublished: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("NewsroomContract", "NameChanged") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("NewsroomContract", "NameChanged")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["NameChanged"]
				pevents, e := f.startFilterNameChanged(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving NameChanged: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("NewsroomContract", "OwnershipRenounced") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("NewsroomContract", "OwnershipRenounced")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["OwnershipRenounced"]
				pevents, e := f.startFilterOwnershipRenounced(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving OwnershipRenounced: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("NewsroomContract", "OwnershipTransferred") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("NewsroomContract", "OwnershipTransferred")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["OwnershipTransferred"]
				pevents, e := f.startFilterOwnershipTransferred(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving OwnershipTransferred: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("NewsroomContract", "RevisionSigned") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("NewsroomContract", "RevisionSigned")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["RevisionSigned"]
				pevents, e := f.startFilterRevisionSigned(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving RevisionSigned: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("NewsroomContract", "RevisionUpdated") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("NewsroomContract", "RevisionUpdated")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["RevisionUpdated"]
				pevents, e := f.startFilterRevisionUpdated(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving RevisionUpdated: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("NewsroomContract", "RoleAdded") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("NewsroomContract", "RoleAdded")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["RoleAdded"]
				pevents, e := f.startFilterRoleAdded(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving RoleAdded: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("NewsroomContract", "RoleRemoved") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("NewsroomContract", "RoleRemoved")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["RoleRemoved"]
				pevents, e := f.startFilterRoleRemoved(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving RoleRemoved: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	go func() {
		wg.Wait()
		close(done)
		log.Info("Filtering routines complete")
	}()

Loop:
	for {
		select {
		case <-done:
			break Loop
		case pevents := <-resultsChan:
			f.pastEventsMutex.Lock()
			pastEvents = append(pastEvents, pevents...)
			f.pastEventsMutex.Unlock()
		}
	}
	log.Infof("Total filterers run: %v, events found: %v", filtsRun, len(pastEvents))
	return pastEvents, nil
}

func (f *NewsroomContractFilterers) startFilterContentPublished(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering NewsroomContract ContentPublished for %v at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.NewsroomContractContentPublishedIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterContentPublished(
			opts,
			[]common.Address{},
			[]*big.Int{},
		)
		if err == nil {
			log.Infof("Done filter: NewsroomContract ContentPublished for %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: NewsroomContract ContentPublished for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: NewsroomContract ContentPublished for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ContentPublished", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("NewsroomContract ContentPublished added: %v", numEventsAdded)
	return pastEvents, nil
}

func (f *NewsroomContractFilterers) startFilterNameChanged(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering NewsroomContract NameChanged for %v at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.NewsroomContractNameChangedIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterNameChanged(
			opts,
		)
		if err == nil {
			log.Infof("Done filter: NewsroomContract NameChanged for %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: NewsroomContract NameChanged for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: NewsroomContract NameChanged for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("NameChanged", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("NewsroomContract NameChanged added: %v", numEventsAdded)
	return pastEvents, nil
}

func (f *NewsroomContractFilterers) startFilterOwnershipRenounced(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering NewsroomContract OwnershipRenounced for %v at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.NewsroomContractOwnershipRenouncedIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterOwnershipRenounced(
			opts,
			[]common.Address{},
		)
		if err == nil {
			log.Infof("Done filter: NewsroomContract OwnershipRenounced for %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: NewsroomContract OwnershipRenounced for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: NewsroomContract OwnershipRenounced for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("OwnershipRenounced", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("NewsroomContract OwnershipRenounced added: %v", numEventsAdded)
	return pastEvents, nil
}

func (f *NewsroomContractFilterers) startFilterOwnershipTransferred(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering NewsroomContract OwnershipTransferred for %v at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.NewsroomContractOwnershipTransferredIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterOwnershipTransferred(
			opts,
			[]common.Address{},
			[]common.Address{},
		)
		if err == nil {
			log.Infof("Done filter: NewsroomContract OwnershipTransferred for %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: NewsroomContract OwnershipTransferred for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: NewsroomContract OwnershipTransferred for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("OwnershipTransferred", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("NewsroomContract OwnershipTransferred added: %v", numEventsAdded)
	return pastEvents, nil
}

func (f *NewsroomContractFilterers) startFilterRevisionSigned(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering NewsroomContract RevisionSigned for %v at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.NewsroomContractRevisionSignedIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterRevisionSigned(
			opts,
			[]*big.Int{},
			[]*big.Int{},
			[]common.Address{},
		)
		if err == nil {
			log.Infof("Done filter: NewsroomContract RevisionSigned for %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: NewsroomContract RevisionSigned for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: NewsroomContract RevisionSigned for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("RevisionSigned", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("NewsroomContract RevisionSigned added: %v", numEventsAdded)
	return pastEvents, nil
}

func (f *NewsroomContractFilterers) startFilterRevisionUpdated(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering NewsroomContract RevisionUpdated for %v at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.NewsroomContractRevisionUpdatedIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterRevisionUpdated(
			opts,
			[]common.Address{},
			[]*big.Int{},
			[]*big.Int{},
		)
		if err == nil {
			log.Infof("Done filter: NewsroomContract RevisionUpdated for %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: NewsroomContract RevisionUpdated for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: NewsroomContract RevisionUpdated for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("RevisionUpdated", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("NewsroomContract RevisionUpdated added: %v", numEventsAdded)
	return pastEvents, nil
}

func (f *NewsroomContractFilterers) startFilterRoleAdded(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering NewsroomContract RoleAdded for %v at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.NewsroomContractRoleAddedIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterRoleAdded(
			opts,
			[]common.Address{},
			[]common.Address{},
		)
		if err == nil {
			log.Infof("Done filter: NewsroomContract RoleAdded for %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: NewsroomContract RoleAdded for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: NewsroomContract RoleAdded for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("RoleAdded", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("NewsroomContract RoleAdded added: %v", numEventsAdded)
	return pastEvents, nil
}

func (f *NewsroomContractFilterers) startFilterRoleRemoved(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering NewsroomContract RoleRemoved for %v at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.NewsroomContractRoleRemovedIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterRoleRemoved(
			opts,
			[]common.Address{},
			[]common.Address{},
		)
		if err == nil {
			log.Infof("Done filter: NewsroomContract RoleRemoved for %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: NewsroomContract RoleRemoved for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: NewsroomContract RoleRemoved for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("RoleRemoved", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("NewsroomContract RoleRemoved added: %v", numEventsAdded)
	return pastEvents, nil
}
