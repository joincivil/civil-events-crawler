// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at 2019-03-03 23:56:08.978547 +0000 UTC
package watcher

import (
	"fmt"
	log "github.com/golang/glog"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/joincivil/civil-events-crawler/pkg/model"

	"github.com/joincivil/go-common/pkg/generated/contract"
	ctime "github.com/joincivil/go-common/pkg/time"

	"math/big"
)

func NewNewsroomContractWatchers(contractAddress common.Address) *NewsroomContractWatchers {
	return &NewsroomContractWatchers{
		contractAddress: contractAddress,
	}
}

type NewsroomContractWatchers struct {
	contractAddress common.Address
	contract        *contract.NewsroomContract
	activeSubs      []event.Subscription
}

func (w *NewsroomContractWatchers) ContractAddress() common.Address {
	return w.contractAddress
}

func (w *NewsroomContractWatchers) ContractName() string {
	return "NewsroomContract"
}

func (w *NewsroomContractWatchers) StopWatchers() error {
	for _, sub := range w.activeSubs {
		sub.Unsubscribe()
	}
	w.activeSubs = nil
	return nil
}

func (w *NewsroomContractWatchers) StartWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event) ([]event.Subscription, error) {
	return w.StartNewsroomContractWatchers(client, eventRecvChan)
}

// StartNewsroomContractWatchers starts up the event watchers for NewsroomContract
func (w *NewsroomContractWatchers) StartNewsroomContractWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event) ([]event.Subscription, error) {
	contract, err := contract.NewNewsroomContract(w.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartNewsroomContract: err: %v", err)
		return nil, err
	}
	w.contract = contract

	var sub event.Subscription
	subs := []event.Subscription{}

	sub, err = w.startWatchContentPublished(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startContentPublished: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchNameChanged(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startNameChanged: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchOwnershipRenounced(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startOwnershipRenounced: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchOwnershipTransferred(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startOwnershipTransferred: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRevisionSigned(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRevisionSigned: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRevisionUpdated(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRevisionUpdated: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRoleAdded(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRoleAdded: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRoleRemoved(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRoleRemoved: err: %v", err)
	}
	subs = append(subs, sub)

	w.activeSubs = subs
	return subs, nil
}

func (w *NewsroomContractWatchers) startWatchContentPublished(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractContentPublished, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractContentPublished)
				log.Infof("startupFn: Starting WatchContentPublished")
				sub, err := w.contract.WatchContentPublished(
					opts,
					recvChan,
					[]common.Address{},
					[]*big.Int{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchContentPublished")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchContentPublished: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchContentPublished: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchContentPublished for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("WATCHER: Premptive restart of ContentPublished")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error starting ContentPublished: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("WATCHER: Done preemptive restart ContentPublished: oldsub: %v, new sub: %v", oldSub, sub)
			case event := <-recvChan:
				log.Errorf("Received event on WatchContentPublished: %v", event)
				modelEvent, err := model.NewEventFromContractEvent("ContentPublished", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchContentPublished, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchContentPublished, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchContentPublished, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchContentPublished, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchNameChanged(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractNameChanged, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractNameChanged)
				log.Infof("startupFn: Starting WatchNameChanged")
				sub, err := w.contract.WatchNameChanged(
					opts,
					recvChan,
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchNameChanged")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchNameChanged: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchNameChanged: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchNameChanged for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("WATCHER: Premptive restart of NameChanged")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error starting NameChanged: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("WATCHER: Done preemptive restart NameChanged: oldsub: %v, new sub: %v", oldSub, sub)
			case event := <-recvChan:
				log.Errorf("Received event on WatchNameChanged: %v", event)
				modelEvent, err := model.NewEventFromContractEvent("NameChanged", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchNameChanged, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchNameChanged, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchNameChanged, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchNameChanged, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchOwnershipRenounced(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractOwnershipRenounced, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractOwnershipRenounced)
				log.Infof("startupFn: Starting WatchOwnershipRenounced")
				sub, err := w.contract.WatchOwnershipRenounced(
					opts,
					recvChan,
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchOwnershipRenounced")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchOwnershipRenounced: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchOwnershipRenounced: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchOwnershipRenounced for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("WATCHER: Premptive restart of OwnershipRenounced")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error starting OwnershipRenounced: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("WATCHER: Done preemptive restart OwnershipRenounced: oldsub: %v, new sub: %v", oldSub, sub)
			case event := <-recvChan:
				log.Errorf("Received event on WatchOwnershipRenounced: %v", event)
				modelEvent, err := model.NewEventFromContractEvent("OwnershipRenounced", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchOwnershipRenounced, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchOwnershipRenounced, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchOwnershipRenounced, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchOwnershipRenounced, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchOwnershipTransferred(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractOwnershipTransferred, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractOwnershipTransferred)
				log.Infof("startupFn: Starting WatchOwnershipTransferred")
				sub, err := w.contract.WatchOwnershipTransferred(
					opts,
					recvChan,
					[]common.Address{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchOwnershipTransferred")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchOwnershipTransferred: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchOwnershipTransferred: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchOwnershipTransferred for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("WATCHER: Premptive restart of OwnershipTransferred")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error starting OwnershipTransferred: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("WATCHER: Done preemptive restart OwnershipTransferred: oldsub: %v, new sub: %v", oldSub, sub)
			case event := <-recvChan:
				log.Errorf("Received event on WatchOwnershipTransferred: %v", event)
				modelEvent, err := model.NewEventFromContractEvent("OwnershipTransferred", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchOwnershipTransferred, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchOwnershipTransferred, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchOwnershipTransferred, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchOwnershipTransferred, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRevisionSigned(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractRevisionSigned, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractRevisionSigned)
				log.Infof("startupFn: Starting WatchRevisionSigned")
				sub, err := w.contract.WatchRevisionSigned(
					opts,
					recvChan,
					[]*big.Int{},
					[]*big.Int{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchRevisionSigned")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchRevisionSigned: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRevisionSigned: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRevisionSigned for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("WATCHER: Premptive restart of RevisionSigned")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error starting RevisionSigned: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("WATCHER: Done preemptive restart RevisionSigned: oldsub: %v, new sub: %v", oldSub, sub)
			case event := <-recvChan:
				log.Errorf("Received event on WatchRevisionSigned: %v", event)
				modelEvent, err := model.NewEventFromContractEvent("RevisionSigned", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchRevisionSigned, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchRevisionSigned, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRevisionSigned, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchRevisionSigned, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRevisionUpdated(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractRevisionUpdated, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractRevisionUpdated)
				log.Infof("startupFn: Starting WatchRevisionUpdated")
				sub, err := w.contract.WatchRevisionUpdated(
					opts,
					recvChan,
					[]common.Address{},
					[]*big.Int{},
					[]*big.Int{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchRevisionUpdated")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchRevisionUpdated: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRevisionUpdated: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRevisionUpdated for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("WATCHER: Premptive restart of RevisionUpdated")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error starting RevisionUpdated: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("WATCHER: Done preemptive restart RevisionUpdated: oldsub: %v, new sub: %v", oldSub, sub)
			case event := <-recvChan:
				log.Errorf("Received event on WatchRevisionUpdated: %v", event)
				modelEvent, err := model.NewEventFromContractEvent("RevisionUpdated", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchRevisionUpdated, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchRevisionUpdated, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRevisionUpdated, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchRevisionUpdated, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRoleAdded(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractRoleAdded, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractRoleAdded)
				log.Infof("startupFn: Starting WatchRoleAdded")
				sub, err := w.contract.WatchRoleAdded(
					opts,
					recvChan,
					[]common.Address{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchRoleAdded")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchRoleAdded: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRoleAdded: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRoleAdded for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("WATCHER: Premptive restart of RoleAdded")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error starting RoleAdded: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("WATCHER: Done preemptive restart RoleAdded: oldsub: %v, new sub: %v", oldSub, sub)
			case event := <-recvChan:
				log.Errorf("Received event on WatchRoleAdded: %v", event)
				modelEvent, err := model.NewEventFromContractEvent("RoleAdded", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchRoleAdded, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchRoleAdded, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRoleAdded, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchRoleAdded, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRoleRemoved(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.NewsroomContractRoleRemoved, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.NewsroomContractRoleRemoved)
				log.Infof("startupFn: Starting WatchRoleRemoved")
				sub, err := w.contract.WatchRoleRemoved(
					opts,
					recvChan,
					[]common.Address{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchRoleRemoved")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchRoleRemoved: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRoleRemoved: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRoleRemoved for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("WATCHER: Premptive restart of RoleRemoved")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error starting RoleRemoved: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("WATCHER: Done preemptive restart RoleRemoved: oldsub: %v, new sub: %v", oldSub, sub)
			case event := <-recvChan:
				log.Errorf("Received event on WatchRoleRemoved: %v", event)
				modelEvent, err := model.NewEventFromContractEvent("RoleRemoved", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchRoleRemoved, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchRoleRemoved, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRoleRemoved, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error restarting WatchRoleRemoved, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}
