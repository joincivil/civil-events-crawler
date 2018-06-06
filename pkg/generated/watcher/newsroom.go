// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at 2018-06-06 19:39:48.648471449 +0000 UTC
package watcher

import (
	"fmt"
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"

	"math/big"
)

type NewsroomContractWatchers struct{}

func (w *NewsroomContractWatchers) ContractName() string {
	return "NewsroomContract"
}

func (w *NewsroomContractWatchers) StartWatchers(client bind.ContractBackend, contractAddress common.Address,
	eventRecvChan chan model.CivilEvent) ([]event.Subscription, error) {
	return w.StartNewsroomContractWatchers(client, contractAddress, eventRecvChan)
}

// StartNewsroomContractWatchers starts up the event watchers for NewsroomContract
func (w *NewsroomContractWatchers) StartNewsroomContractWatchers(client bind.ContractBackend,
	contractAddress common.Address, eventRecvChan chan model.CivilEvent) ([]event.Subscription, error) {
	contract, err := contract.NewNewsroomContract(contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartNewsroomContract: err: %v", err)
		return nil, err
	}

	var sub event.Subscription
	subs := []event.Subscription{}

	sub, err = startWatchContentPublished(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startContentPublished: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchNameChanged(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startNameChanged: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchOwnershipTransferred(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startOwnershipTransferred: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchRevisionSigned(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRevisionSigned: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchRevisionUpdated(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRevisionUpdated: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchRoleAdded(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRoleAdded: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchRoleRemoved(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRoleRemoved: err: %v", err)
	}
	subs = append(subs, sub)

	return subs, nil
}

func startWatchContentPublished(eventRecvChan chan model.CivilEvent, _contract *contract.NewsroomContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.NewsroomContractContentPublished)
	sub, err := _contract.WatchContentPublished(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error starting WatchContentPublished: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("ContentPublished", event)
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func startWatchNameChanged(eventRecvChan chan model.CivilEvent, _contract *contract.NewsroomContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.NewsroomContractNameChanged)
	sub, err := _contract.WatchNameChanged(
		opts,
		recvChan,
	)
	if err != nil {
		log.Errorf("Error starting WatchNameChanged: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("NameChanged", event)
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func startWatchOwnershipTransferred(eventRecvChan chan model.CivilEvent, _contract *contract.NewsroomContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.NewsroomContractOwnershipTransferred)
	sub, err := _contract.WatchOwnershipTransferred(
		opts,
		recvChan,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchOwnershipTransferred: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("OwnershipTransferred", event)
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func startWatchRevisionSigned(eventRecvChan chan model.CivilEvent, _contract *contract.NewsroomContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.NewsroomContractRevisionSigned)
	sub, err := _contract.WatchRevisionSigned(
		opts,
		recvChan,
		[]*big.Int{},
		[]*big.Int{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchRevisionSigned: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("RevisionSigned", event)
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func startWatchRevisionUpdated(eventRecvChan chan model.CivilEvent, _contract *contract.NewsroomContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.NewsroomContractRevisionUpdated)
	sub, err := _contract.WatchRevisionUpdated(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error starting WatchRevisionUpdated: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("RevisionUpdated", event)
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func startWatchRoleAdded(eventRecvChan chan model.CivilEvent, _contract *contract.NewsroomContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.NewsroomContractRoleAdded)
	sub, err := _contract.WatchRoleAdded(
		opts,
		recvChan,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchRoleAdded: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("RoleAdded", event)
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func startWatchRoleRemoved(eventRecvChan chan model.CivilEvent, _contract *contract.NewsroomContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.NewsroomContractRoleRemoved)
	sub, err := _contract.WatchRoleRemoved(
		opts,
		recvChan,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchRoleRemoved: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("RoleRemoved", event)
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
