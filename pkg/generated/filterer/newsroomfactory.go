// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2019-12-19 21:31:50.993546 +0000 UTC
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
)

func NewNewsroomFactoryFilterers(contractAddress common.Address) *NewsroomFactoryFilterers {
	contractFilterers := &NewsroomFactoryFilterers{
		contractAddress:   contractAddress,
		eventTypes:        commongen.EventTypesNewsroomFactory(),
		eventToStartBlock: make(map[string]uint64),
		lastEvents:        make([]*model.Event, 0),
	}
	for _, eventType := range contractFilterers.eventTypes {
		contractFilterers.eventToStartBlock[eventType] = 0
	}
	return contractFilterers
}

type NewsroomFactoryFilterers struct {
	contractAddress   common.Address
	contract          *contract.NewsroomFactory
	eventTypes        []string
	eventToStartBlock map[string]uint64
	lastEvents        []*model.Event
	lastEventsMutex   sync.Mutex
	pastEventsMutex   sync.Mutex
}

func (f *NewsroomFactoryFilterers) ContractName() string {
	return "NewsroomFactory"
}

func (f *NewsroomFactoryFilterers) ContractAddress() common.Address {
	return f.contractAddress
}

func (f *NewsroomFactoryFilterers) StartFilterers(client bind.ContractBackend,
	pastEvents []*model.Event, nonSubOnly bool) ([]*model.Event, error) {
	return f.StartNewsroomFactoryFilterers(client, pastEvents, nonSubOnly)
}

func (f *NewsroomFactoryFilterers) EventTypes() []string {
	return f.eventTypes
}

func (f *NewsroomFactoryFilterers) UpdateStartBlock(eventType string, startBlock uint64) {
	f.eventToStartBlock[eventType] = startBlock
}

func (f *NewsroomFactoryFilterers) LastEvents() []*model.Event {
	return f.lastEvents
}

// StartNewsroomFactoryFilterers retrieves events for NewsroomFactory
func (f *NewsroomFactoryFilterers) StartNewsroomFactoryFilterers(client bind.ContractBackend,
	pastEvents []*model.Event, nonSubOnly bool) ([]*model.Event, error) {
	contract, err := contract.NewNewsroomFactory(f.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartNewsroomFactory: err: %v", err)
		return pastEvents, err
	}
	f.contract = contract

	workerMultiplier := 1
	numWorkers := runtime.NumCPU() * workerMultiplier
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

	if !specs.IsEventDisabled("NewsroomFactory", "ContractInstantiation") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("NewsroomFactory", "ContractInstantiation")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["ContractInstantiation"]
				pevents, e := f.startFilterContractInstantiation(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving ContractInstantiation: err: %v", e)
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

func (f *NewsroomFactoryFilterers) startFilterContractInstantiation(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering NewsroomFactory ContractInstantiation for %v at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.NewsroomFactoryContractInstantiationIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterContractInstantiation(
			opts,
		)
		if err == nil {
			log.Infof("Done filter: NewsroomFactory ContractInstantiation for %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: NewsroomFactory ContractInstantiation for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: NewsroomFactory ContractInstantiation for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ContractInstantiation", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("NewsroomFactory ContractInstantiation added: %v", numEventsAdded)
	return pastEvents, nil
}
