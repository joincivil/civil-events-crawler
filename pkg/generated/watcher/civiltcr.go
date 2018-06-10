// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at 2018-06-10 03:46:10.32133596 +0000 UTC
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

func NewCivilTCRContractWatchers(contractAddress common.Address) *CivilTCRContractWatchers {
	return &CivilTCRContractWatchers{
		contractAddress: contractAddress,
	}
}

type CivilTCRContractWatchers struct {
	contractAddress common.Address
}

func (w *CivilTCRContractWatchers) ContractName() string {
	return "CivilTCRContract"
}

func (w *CivilTCRContractWatchers) StartWatchers(client bind.ContractBackend,
	eventRecvChan chan model.CivilEvent) ([]event.Subscription, error) {
	return w.StartCivilTCRContractWatchers(client, eventRecvChan)
}

// StartCivilTCRContractWatchers starts up the event watchers for CivilTCRContract
func (w *CivilTCRContractWatchers) StartCivilTCRContractWatchers(client bind.ContractBackend,
	eventRecvChan chan model.CivilEvent) ([]event.Subscription, error) {
	contract, err := contract.NewCivilTCRContract(w.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartCivilTCRContract: err: %v", err)
		return nil, err
	}

	var sub event.Subscription
	subs := []event.Subscription{}

	sub, err = startWatchAppealGranted(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startAppealGranted: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchAppealRequested(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startAppealRequested: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchApplication(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startApplication: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchApplicationRemoved(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startApplicationRemoved: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchApplicationWhitelisted(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startApplicationWhitelisted: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchChallenge(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startChallenge: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchChallengeFailed(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startChallengeFailed: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchChallengeSucceeded(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startChallengeSucceeded: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchDeposit(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startDeposit: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchFailedChallengeOverturned(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startFailedChallengeOverturned: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchGovernmentTransfered(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startGovernmentTransfered: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchGrantedAppealChallenged(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startGrantedAppealChallenged: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchGrantedAppealConfirmed(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startGrantedAppealConfirmed: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchGrantedAppealOverturned(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startGrantedAppealOverturned: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchListingRemoved(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startListingRemoved: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchListingWithdrawn(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startListingWithdrawn: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchRewardClaimed(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startRewardClaimed: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchSuccessfulChallengeOverturned(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startSuccessfulChallengeOverturned: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchTouchAndRemoved(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startTouchAndRemoved: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = startWatchWithdrawal(eventRecvChan, contract)
	if err != nil {
		return nil, fmt.Errorf("Error starting startWithdrawal: err: %v", err)
	}
	subs = append(subs, sub)

	return subs, nil
}

func startWatchAppealGranted(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractAppealGranted)
	sub, err := _contract.WatchAppealGranted(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error starting WatchAppealGranted: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_AppealGranted", event)
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

func startWatchAppealRequested(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractAppealRequested)
	sub, err := _contract.WatchAppealRequested(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error starting WatchAppealRequested: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_AppealRequested", event)
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

func startWatchApplication(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractApplication)
	sub, err := _contract.WatchApplication(
		opts,
		recvChan,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchApplication: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_Application", event)
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

func startWatchApplicationRemoved(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractApplicationRemoved)
	sub, err := _contract.WatchApplicationRemoved(
		opts,
		recvChan,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchApplicationRemoved: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_ApplicationRemoved", event)
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

func startWatchApplicationWhitelisted(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractApplicationWhitelisted)
	sub, err := _contract.WatchApplicationWhitelisted(
		opts,
		recvChan,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchApplicationWhitelisted: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_ApplicationWhitelisted", event)
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

func startWatchChallenge(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractChallenge)
	sub, err := _contract.WatchChallenge(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchChallenge: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_Challenge", event)
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

func startWatchChallengeFailed(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractChallengeFailed)
	sub, err := _contract.WatchChallengeFailed(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error starting WatchChallengeFailed: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_ChallengeFailed", event)
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

func startWatchChallengeSucceeded(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractChallengeSucceeded)
	sub, err := _contract.WatchChallengeSucceeded(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error starting WatchChallengeSucceeded: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_ChallengeSucceeded", event)
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

func startWatchDeposit(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractDeposit)
	sub, err := _contract.WatchDeposit(
		opts,
		recvChan,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchDeposit: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_Deposit", event)
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

func startWatchFailedChallengeOverturned(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractFailedChallengeOverturned)
	sub, err := _contract.WatchFailedChallengeOverturned(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error starting WatchFailedChallengeOverturned: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_FailedChallengeOverturned", event)
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

func startWatchGovernmentTransfered(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractGovernmentTransfered)
	sub, err := _contract.WatchGovernmentTransfered(
		opts,
		recvChan,
	)
	if err != nil {
		log.Errorf("Error starting WatchGovernmentTransfered: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_GovernmentTransfered", event)
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

func startWatchGrantedAppealChallenged(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractGrantedAppealChallenged)
	sub, err := _contract.WatchGrantedAppealChallenged(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error starting WatchGrantedAppealChallenged: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_GrantedAppealChallenged", event)
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

func startWatchGrantedAppealConfirmed(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractGrantedAppealConfirmed)
	sub, err := _contract.WatchGrantedAppealConfirmed(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error starting WatchGrantedAppealConfirmed: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_GrantedAppealConfirmed", event)
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

func startWatchGrantedAppealOverturned(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractGrantedAppealOverturned)
	sub, err := _contract.WatchGrantedAppealOverturned(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error starting WatchGrantedAppealOverturned: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_GrantedAppealOverturned", event)
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

func startWatchListingRemoved(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractListingRemoved)
	sub, err := _contract.WatchListingRemoved(
		opts,
		recvChan,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchListingRemoved: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_ListingRemoved", event)
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

func startWatchListingWithdrawn(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractListingWithdrawn)
	sub, err := _contract.WatchListingWithdrawn(
		opts,
		recvChan,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchListingWithdrawn: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_ListingWithdrawn", event)
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

func startWatchRewardClaimed(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractRewardClaimed)
	sub, err := _contract.WatchRewardClaimed(
		opts,
		recvChan,
		[]*big.Int{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchRewardClaimed: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_RewardClaimed", event)
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

func startWatchSuccessfulChallengeOverturned(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractSuccessfulChallengeOverturned)
	sub, err := _contract.WatchSuccessfulChallengeOverturned(
		opts,
		recvChan,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error starting WatchSuccessfulChallengeOverturned: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_SuccessfulChallengeOverturned", event)
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

func startWatchTouchAndRemoved(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractTouchAndRemoved)
	sub, err := _contract.WatchTouchAndRemoved(
		opts,
		recvChan,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchTouchAndRemoved: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_TouchAndRemoved", event)
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

func startWatchWithdrawal(eventRecvChan chan model.CivilEvent, _contract *contract.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *contract.CivilTCRContractWithdrawal)
	sub, err := _contract.WatchWithdrawal(
		opts,
		recvChan,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error starting WatchWithdrawal: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("_Withdrawal", event)
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
