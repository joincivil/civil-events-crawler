// Package listener contains all the components for the events listener, which
// streams a list of future events.
package listener // import "github.com/joincivil/civil-events-crawler/pkg/listener"

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	log "github.com/golang/glog"
	"github.com/joincivil/civil-events-crawler/pkg/generated/tcr"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	// "github.com/joincivil/civil-events-crawler/pkg/utils"
)

func startWatchApplication(eventRecvChan chan model.CivilEvent, civilTCR *tcr.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *tcr.CivilTCRContractApplication)
	sub, err := civilTCR.WatchApplication(
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

func startWatchApplicationRemoved(eventRecvChan chan model.CivilEvent, civilTCR *tcr.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *tcr.CivilTCRContractApplicationRemoved)
	sub, err := civilTCR.WatchApplicationRemoved(
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

func startWatchApplicationWhitelisted(eventRecvChan chan model.CivilEvent, civilTCR *tcr.CivilTCRContract) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
	recvChan := make(chan *tcr.CivilTCRContractApplicationWhitelisted)
	sub, err := civilTCR.WatchApplicationWhitelisted(
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
