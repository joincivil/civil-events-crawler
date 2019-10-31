// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2019-10-31 19:47:31.594458 +0000 UTC
package filterer

import (
	log "github.com/golang/glog"
	"runtime"
	"sync"

	"github.com/Jeffail/tunny"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	commongen "github.com/joincivil/civil-events-crawler/pkg/generated/common"
	"github.com/joincivil/civil-events-crawler/pkg/model"

	"github.com/joincivil/go-common/pkg/generated/contract"
	ctime "github.com/joincivil/go-common/pkg/time"

	"math/big"
)

func NewMultiSigWalletContractFilterers(contractAddress common.Address) *MultiSigWalletContractFilterers {
	contractFilterers := &MultiSigWalletContractFilterers{
		contractAddress:   contractAddress,
		eventTypes:        commongen.EventTypesMultiSigWalletContract(),
		eventToStartBlock: make(map[string]uint64),
		lastEvents:        make([]*model.Event, 0),
	}
	for _, eventType := range contractFilterers.eventTypes {
		contractFilterers.eventToStartBlock[eventType] = 0
	}
	return contractFilterers
}

type MultiSigWalletContractFilterers struct {
	contractAddress   common.Address
	contract          *contract.MultiSigWalletContract
	eventTypes        []string
	eventToStartBlock map[string]uint64
	lastEvents        []*model.Event
	lastEventsMutex   sync.Mutex
	pastEventsMutex   sync.Mutex
}

func (f *MultiSigWalletContractFilterers) ContractName() string {
	return "MultiSigWalletContract"
}

func (f *MultiSigWalletContractFilterers) ContractAddress() common.Address {
	return f.contractAddress
}

func (f *MultiSigWalletContractFilterers) StartFilterers(client bind.ContractBackend, pastEvents []*model.Event) (error, []*model.Event) {
	return f.StartMultiSigWalletContractFilterers(client, pastEvents)
}

func (f *MultiSigWalletContractFilterers) EventTypes() []string {
	return f.eventTypes
}

func (f *MultiSigWalletContractFilterers) UpdateStartBlock(eventType string, startBlock uint64) {
	f.eventToStartBlock[eventType] = startBlock
}

func (f *MultiSigWalletContractFilterers) LastEvents() []*model.Event {
	return f.lastEvents
}

// StartMultiSigWalletContractFilterers retrieves events for MultiSigWalletContract
func (f *MultiSigWalletContractFilterers) StartMultiSigWalletContractFilterers(client bind.ContractBackend, pastEvents []*model.Event) (error, []*model.Event) {
	contract, err := contract.NewMultiSigWalletContract(f.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartMultiSigWalletContract: err: %v", err)
		return err, pastEvents
	}
	f.contract = contract

	workerMultiplier := 1
	numWorkers := runtime.NumCPU() * workerMultiplier
	log.Infof("Num of workers: %v", numWorkers)
	pool := tunny.NewFunc(numWorkers, func(payload interface{}) interface{} {
		f := payload.(func())
		f()
		return nil
	})
	defer pool.Close()

	wg := sync.WaitGroup{}
	resultsChan := make(chan []*model.Event)
	done := make(chan bool)

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["Confirmation"]
			e, pevents := f.startFilterConfirmation(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving Confirmation: err: %v", e)
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

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["Deposit"]
			e, pevents := f.startFilterDeposit(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving Deposit: err: %v", e)
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

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["Execution"]
			e, pevents := f.startFilterExecution(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving Execution: err: %v", e)
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

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["ExecutionFailure"]
			e, pevents := f.startFilterExecutionFailure(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving ExecutionFailure: err: %v", e)
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

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["OwnerAddition"]
			e, pevents := f.startFilterOwnerAddition(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving OwnerAddition: err: %v", e)
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

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["OwnerRemoval"]
			e, pevents := f.startFilterOwnerRemoval(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving OwnerRemoval: err: %v", e)
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

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["RequirementChange"]
			e, pevents := f.startFilterRequirementChange(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving RequirementChange: err: %v", e)
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

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["Revocation"]
			e, pevents := f.startFilterRevocation(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving Revocation: err: %v", e)
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

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["Submission"]
			e, pevents := f.startFilterSubmission(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving Submission: err: %v", e)
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

	go func() {
		wg.Wait()
		done <- true
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
	log.Infof("Total events found: %v", len(pastEvents))
	return nil, pastEvents
}

func (f *MultiSigWalletContractFilterers) startFilterConfirmation(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for Confirmation for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.MultiSigWalletContractConfirmationIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterConfirmation(
			opts,
			[]common.Address{},
			[]*big.Int{},
		)
		if err == nil {
			log.Infof("Successful filter: Confirmation for contract %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: Confirmation for contract %v: err: %v", f.contractAddress.Hex(), err)
			return err, pastEvents
		}
		log.Infof("Retrying filter: Confirmation for contract %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Confirmation", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("Confirmation events added: %v", numEventsAdded)
	return nil, pastEvents
}

func (f *MultiSigWalletContractFilterers) startFilterDeposit(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for Deposit for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.MultiSigWalletContractDepositIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterDeposit(
			opts,
			[]common.Address{},
		)
		if err == nil {
			log.Infof("Successful filter: Deposit for contract %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: Deposit for contract %v: err: %v", f.contractAddress.Hex(), err)
			return err, pastEvents
		}
		log.Infof("Retrying filter: Deposit for contract %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Deposit", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("Deposit events added: %v", numEventsAdded)
	return nil, pastEvents
}

func (f *MultiSigWalletContractFilterers) startFilterExecution(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for Execution for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.MultiSigWalletContractExecutionIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterExecution(
			opts,
			[]*big.Int{},
		)
		if err == nil {
			log.Infof("Successful filter: Execution for contract %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: Execution for contract %v: err: %v", f.contractAddress.Hex(), err)
			return err, pastEvents
		}
		log.Infof("Retrying filter: Execution for contract %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Execution", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("Execution events added: %v", numEventsAdded)
	return nil, pastEvents
}

func (f *MultiSigWalletContractFilterers) startFilterExecutionFailure(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for ExecutionFailure for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.MultiSigWalletContractExecutionFailureIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterExecutionFailure(
			opts,
			[]*big.Int{},
		)
		if err == nil {
			log.Infof("Successful filter: ExecutionFailure for contract %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: ExecutionFailure for contract %v: err: %v", f.contractAddress.Hex(), err)
			return err, pastEvents
		}
		log.Infof("Retrying filter: ExecutionFailure for contract %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ExecutionFailure", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("ExecutionFailure events added: %v", numEventsAdded)
	return nil, pastEvents
}

func (f *MultiSigWalletContractFilterers) startFilterOwnerAddition(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for OwnerAddition for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.MultiSigWalletContractOwnerAdditionIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterOwnerAddition(
			opts,
			[]common.Address{},
		)
		if err == nil {
			log.Infof("Successful filter: OwnerAddition for contract %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: OwnerAddition for contract %v: err: %v", f.contractAddress.Hex(), err)
			return err, pastEvents
		}
		log.Infof("Retrying filter: OwnerAddition for contract %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("OwnerAddition", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("OwnerAddition events added: %v", numEventsAdded)
	return nil, pastEvents
}

func (f *MultiSigWalletContractFilterers) startFilterOwnerRemoval(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for OwnerRemoval for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.MultiSigWalletContractOwnerRemovalIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterOwnerRemoval(
			opts,
			[]common.Address{},
		)
		if err == nil {
			log.Infof("Successful filter: OwnerRemoval for contract %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: OwnerRemoval for contract %v: err: %v", f.contractAddress.Hex(), err)
			return err, pastEvents
		}
		log.Infof("Retrying filter: OwnerRemoval for contract %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("OwnerRemoval", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("OwnerRemoval events added: %v", numEventsAdded)
	return nil, pastEvents
}

func (f *MultiSigWalletContractFilterers) startFilterRequirementChange(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for RequirementChange for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.MultiSigWalletContractRequirementChangeIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterRequirementChange(
			opts,
		)
		if err == nil {
			log.Infof("Successful filter: RequirementChange for contract %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: RequirementChange for contract %v: err: %v", f.contractAddress.Hex(), err)
			return err, pastEvents
		}
		log.Infof("Retrying filter: RequirementChange for contract %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("RequirementChange", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("RequirementChange events added: %v", numEventsAdded)
	return nil, pastEvents
}

func (f *MultiSigWalletContractFilterers) startFilterRevocation(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for Revocation for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.MultiSigWalletContractRevocationIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterRevocation(
			opts,
			[]common.Address{},
			[]*big.Int{},
		)
		if err == nil {
			log.Infof("Successful filter: Revocation for contract %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: Revocation for contract %v: err: %v", f.contractAddress.Hex(), err)
			return err, pastEvents
		}
		log.Infof("Retrying filter: Revocation for contract %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Revocation", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("Revocation events added: %v", numEventsAdded)
	return nil, pastEvents
}

func (f *MultiSigWalletContractFilterers) startFilterSubmission(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for Submission for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.MultiSigWalletContractSubmissionIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterSubmission(
			opts,
			[]*big.Int{},
		)
		if err == nil {
			log.Infof("Successful filter: Submission for contract %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: Submission for contract %v: err: %v", f.contractAddress.Hex(), err)
			return err, pastEvents
		}
		log.Infof("Retrying filter: Submission for contract %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Submission", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("Submission events added: %v", numEventsAdded)
	return nil, pastEvents
}
