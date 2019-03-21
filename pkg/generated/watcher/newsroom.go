// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at 2019-03-20 17:54:53.475742 +0000 UTC
package watcher

import (
	// "fmt"
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"
	log "github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

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
	errors          chan error
	contractAddress common.Address
	contract        *contract.NewsroomContract
	activeSubs      []utils.WatcherSubscription
}

func (w *NewsroomContractWatchers) ContractAddress() common.Address {
	return w.contractAddress
}

func (w *NewsroomContractWatchers) ContractName() string {
	return "NewsroomContract"
}

func (w *NewsroomContractWatchers) cancelFunc(cancelFn context.CancelFunc, killCancel <-chan bool) {
}

func (w *NewsroomContractWatchers) StopWatchers(unsub bool) error {
	if unsub {
		for _, sub := range w.activeSubs {
			sub.Unsubscribe()
		}
	}
	w.activeSubs = nil
	return nil
}

func (w *NewsroomContractWatchers) StartWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event, errs chan error) ([]utils.WatcherSubscription, error) {
	return w.StartNewsroomContractWatchers(client, eventRecvChan, errs)
}

// StartNewsroomContractWatchers starts up the event watchers for NewsroomContract
func (w *NewsroomContractWatchers) StartNewsroomContractWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event, errs chan error) ([]utils.WatcherSubscription, error) {
	w.errors = errs
	contract, err := contract.NewNewsroomContract(w.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartNewsroomContract: err: %v", err)
		return nil, errors.Wrap(err, "error initializing StartNewsroomContract")
	}
	w.contract = contract

	var sub utils.WatcherSubscription
	subs := []utils.WatcherSubscription{}

	sub, err = w.startWatchContentPublished(eventRecvChan)
	if err != nil {
		return nil, errors.WithMessage(err, "error starting startContentPublished")
	}
	subs = append(subs, sub)

	sub, err = w.startWatchNameChanged(eventRecvChan)
	if err != nil {
		return nil, errors.WithMessage(err, "error starting startNameChanged")
	}
	subs = append(subs, sub)

	sub, err = w.startWatchOwnershipRenounced(eventRecvChan)
	if err != nil {
		return nil, errors.WithMessage(err, "error starting startOwnershipRenounced")
	}
	subs = append(subs, sub)

	sub, err = w.startWatchOwnershipTransferred(eventRecvChan)
	if err != nil {
		return nil, errors.WithMessage(err, "error starting startOwnershipTransferred")
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRevisionSigned(eventRecvChan)
	if err != nil {
		return nil, errors.WithMessage(err, "error starting startRevisionSigned")
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRevisionUpdated(eventRecvChan)
	if err != nil {
		return nil, errors.WithMessage(err, "error starting startRevisionUpdated")
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRoleAdded(eventRecvChan)
	if err != nil {
		return nil, errors.WithMessage(err, "error starting startRoleAdded")
	}
	subs = append(subs, sub)

	sub, err = w.startWatchRoleRemoved(eventRecvChan)
	if err != nil {
		return nil, errors.WithMessage(err, "error starting startRoleRemoved")
	}
	subs = append(subs, sub)

	w.activeSubs = subs
	return subs, nil
}

func (w *NewsroomContractWatchers) startWatchContentPublished(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	preemptiveTimeoutSecs := 60 * 30
	return utils.NewWatcherSubscription("WatchContentPublished", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractContentPublished, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan bool)
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan bool) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchContentPublished start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractContentPublished)
			log.Infof("startupFn: Starting WatchContentPublished")
			sub, err := w.contract.WatchContentPublished(
				opts,
				recvChan,
				[]common.Address{},
				[]*big.Int{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchContentPublished")
					sub.Unsubscribe()
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchContentPublished")
			}
			log.Infof("startupFn: WatchContentPublished started")
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchContentPublished: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchContentPublished for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 30 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(preemptiveTimeoutSecs)):
				log.Infof("Premptive restart of ContentPublished")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting ContentPublished: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old ContentPublished")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart ContentPublished")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchContentPublished: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchContentPublished")
				}
				modelEvent, err := model.NewEventFromContractEvent("ContentPublished", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchContentPublished: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchContentPublished")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchContentPublished, fatal (a): %v", err)
					err = errors.Wrap(err, "error with WatchContentPublished")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchContentPublished (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchContentPublished, fatal (b): %v", err)
				err = errors.Wrap(err, "error with WatchContentPublished")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchContentPublished")
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchNameChanged(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	preemptiveTimeoutSecs := 60 * 30
	return utils.NewWatcherSubscription("WatchNameChanged", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractNameChanged, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan bool)
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan bool) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchNameChanged start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractNameChanged)
			log.Infof("startupFn: Starting WatchNameChanged")
			sub, err := w.contract.WatchNameChanged(
				opts,
				recvChan,
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchNameChanged")
					sub.Unsubscribe()
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchNameChanged")
			}
			log.Infof("startupFn: WatchNameChanged started")
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchNameChanged: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchNameChanged for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 30 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(preemptiveTimeoutSecs)):
				log.Infof("Premptive restart of NameChanged")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting NameChanged: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old NameChanged")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart NameChanged")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchNameChanged: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchNameChanged")
				}
				modelEvent, err := model.NewEventFromContractEvent("NameChanged", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchNameChanged: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchNameChanged")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchNameChanged, fatal (a): %v", err)
					err = errors.Wrap(err, "error with WatchNameChanged")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchNameChanged (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchNameChanged, fatal (b): %v", err)
				err = errors.Wrap(err, "error with WatchNameChanged")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchNameChanged")
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchOwnershipRenounced(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	preemptiveTimeoutSecs := 60 * 30
	return utils.NewWatcherSubscription("WatchOwnershipRenounced", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractOwnershipRenounced, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan bool)
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan bool) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchOwnershipRenounced start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractOwnershipRenounced)
			log.Infof("startupFn: Starting WatchOwnershipRenounced")
			sub, err := w.contract.WatchOwnershipRenounced(
				opts,
				recvChan,
				[]common.Address{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchOwnershipRenounced")
					sub.Unsubscribe()
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchOwnershipRenounced")
			}
			log.Infof("startupFn: WatchOwnershipRenounced started")
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchOwnershipRenounced: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchOwnershipRenounced for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 30 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(preemptiveTimeoutSecs)):
				log.Infof("Premptive restart of OwnershipRenounced")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting OwnershipRenounced: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old OwnershipRenounced")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart OwnershipRenounced")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchOwnershipRenounced: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchOwnershipRenounced")
				}
				modelEvent, err := model.NewEventFromContractEvent("OwnershipRenounced", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchOwnershipRenounced: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchOwnershipRenounced")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchOwnershipRenounced, fatal (a): %v", err)
					err = errors.Wrap(err, "error with WatchOwnershipRenounced")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchOwnershipRenounced (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchOwnershipRenounced, fatal (b): %v", err)
				err = errors.Wrap(err, "error with WatchOwnershipRenounced")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchOwnershipRenounced")
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchOwnershipTransferred(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	preemptiveTimeoutSecs := 60 * 30
	return utils.NewWatcherSubscription("WatchOwnershipTransferred", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractOwnershipTransferred, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan bool)
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan bool) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchOwnershipTransferred start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractOwnershipTransferred)
			log.Infof("startupFn: Starting WatchOwnershipTransferred")
			sub, err := w.contract.WatchOwnershipTransferred(
				opts,
				recvChan,
				[]common.Address{},
				[]common.Address{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchOwnershipTransferred")
					sub.Unsubscribe()
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchOwnershipTransferred")
			}
			log.Infof("startupFn: WatchOwnershipTransferred started")
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchOwnershipTransferred: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchOwnershipTransferred for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 30 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(preemptiveTimeoutSecs)):
				log.Infof("Premptive restart of OwnershipTransferred")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting OwnershipTransferred: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old OwnershipTransferred")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart OwnershipTransferred")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchOwnershipTransferred: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchOwnershipTransferred")
				}
				modelEvent, err := model.NewEventFromContractEvent("OwnershipTransferred", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchOwnershipTransferred: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchOwnershipTransferred")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchOwnershipTransferred, fatal (a): %v", err)
					err = errors.Wrap(err, "error with WatchOwnershipTransferred")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchOwnershipTransferred (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchOwnershipTransferred, fatal (b): %v", err)
				err = errors.Wrap(err, "error with WatchOwnershipTransferred")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchOwnershipTransferred")
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRevisionSigned(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	preemptiveTimeoutSecs := 60 * 30
	return utils.NewWatcherSubscription("WatchRevisionSigned", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractRevisionSigned, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan bool)
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan bool) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchRevisionSigned start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractRevisionSigned)
			log.Infof("startupFn: Starting WatchRevisionSigned")
			sub, err := w.contract.WatchRevisionSigned(
				opts,
				recvChan,
				[]*big.Int{},
				[]*big.Int{},
				[]common.Address{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchRevisionSigned")
					sub.Unsubscribe()
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchRevisionSigned")
			}
			log.Infof("startupFn: WatchRevisionSigned started")
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRevisionSigned: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRevisionSigned for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 30 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(preemptiveTimeoutSecs)):
				log.Infof("Premptive restart of RevisionSigned")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting RevisionSigned: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old RevisionSigned")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart RevisionSigned")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchRevisionSigned: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchRevisionSigned")
				}
				modelEvent, err := model.NewEventFromContractEvent("RevisionSigned", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchRevisionSigned: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchRevisionSigned")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchRevisionSigned, fatal (a): %v", err)
					err = errors.Wrap(err, "error with WatchRevisionSigned")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchRevisionSigned (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRevisionSigned, fatal (b): %v", err)
				err = errors.Wrap(err, "error with WatchRevisionSigned")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchRevisionSigned")
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRevisionUpdated(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	preemptiveTimeoutSecs := 60 * 30
	return utils.NewWatcherSubscription("WatchRevisionUpdated", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractRevisionUpdated, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan bool)
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan bool) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchRevisionUpdated start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractRevisionUpdated)
			log.Infof("startupFn: Starting WatchRevisionUpdated")
			sub, err := w.contract.WatchRevisionUpdated(
				opts,
				recvChan,
				[]common.Address{},
				[]*big.Int{},
				[]*big.Int{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchRevisionUpdated")
					sub.Unsubscribe()
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchRevisionUpdated")
			}
			log.Infof("startupFn: WatchRevisionUpdated started")
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRevisionUpdated: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRevisionUpdated for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 30 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(preemptiveTimeoutSecs)):
				log.Infof("Premptive restart of RevisionUpdated")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting RevisionUpdated: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old RevisionUpdated")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart RevisionUpdated")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchRevisionUpdated: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchRevisionUpdated")
				}
				modelEvent, err := model.NewEventFromContractEvent("RevisionUpdated", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchRevisionUpdated: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchRevisionUpdated")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchRevisionUpdated, fatal (a): %v", err)
					err = errors.Wrap(err, "error with WatchRevisionUpdated")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchRevisionUpdated (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRevisionUpdated, fatal (b): %v", err)
				err = errors.Wrap(err, "error with WatchRevisionUpdated")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchRevisionUpdated")
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRoleAdded(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	preemptiveTimeoutSecs := 60 * 30
	return utils.NewWatcherSubscription("WatchRoleAdded", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractRoleAdded, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan bool)
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan bool) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchRoleAdded start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractRoleAdded)
			log.Infof("startupFn: Starting WatchRoleAdded")
			sub, err := w.contract.WatchRoleAdded(
				opts,
				recvChan,
				[]common.Address{},
				[]common.Address{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchRoleAdded")
					sub.Unsubscribe()
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchRoleAdded")
			}
			log.Infof("startupFn: WatchRoleAdded started")
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRoleAdded: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRoleAdded for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 30 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(preemptiveTimeoutSecs)):
				log.Infof("Premptive restart of RoleAdded")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting RoleAdded: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old RoleAdded")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart RoleAdded")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchRoleAdded: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchRoleAdded")
				}
				modelEvent, err := model.NewEventFromContractEvent("RoleAdded", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchRoleAdded: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchRoleAdded")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchRoleAdded, fatal (a): %v", err)
					err = errors.Wrap(err, "error with WatchRoleAdded")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchRoleAdded (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRoleAdded, fatal (b): %v", err)
				err = errors.Wrap(err, "error with WatchRoleAdded")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchRoleAdded")
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRoleRemoved(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	preemptiveTimeoutSecs := 60 * 30
	return utils.NewWatcherSubscription("WatchRoleRemoved", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractRoleRemoved, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan bool)
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan bool) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchRoleRemoved start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractRoleRemoved)
			log.Infof("startupFn: Starting WatchRoleRemoved")
			sub, err := w.contract.WatchRoleRemoved(
				opts,
				recvChan,
				[]common.Address{},
				[]common.Address{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchRoleRemoved")
					sub.Unsubscribe()
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchRoleRemoved")
			}
			log.Infof("startupFn: WatchRoleRemoved started")
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRoleRemoved: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchRoleRemoved for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 30 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(preemptiveTimeoutSecs)):
				log.Infof("Premptive restart of RoleRemoved")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting RoleRemoved: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old RoleRemoved")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart RoleRemoved")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchRoleRemoved: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchRoleRemoved")
				}
				modelEvent, err := model.NewEventFromContractEvent("RoleRemoved", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchRoleRemoved: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchRoleRemoved")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchRoleRemoved, fatal (a): %v", err)
					err = errors.Wrap(err, "error with WatchRoleRemoved")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchRoleRemoved (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRoleRemoved, fatal (b): %v", err)
				err = errors.Wrap(err, "error with WatchRoleRemoved")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchRoleRemoved")
				return nil
			}
		}
	}), nil
}
