// Package retriever contains all components for events retriever, which
// gets all past events
package retriever // import "github.com/joincivil/civil-events-crawler/pkg/retriever"

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/golang/glog"
	"github.com/joincivil/civil-events-crawler/pkg/generated/tcr"
	"github.com/joincivil/civil-events-crawler/pkg/model"
)

// RetrieveApplication gets _Application events
func RetrieveApplication(opts *bind.FilterOpts, civilTCR *tcr.CivilTCRContract, pastEvents *[]model.CivilEvent) error {
	itr, err := civilTCR.FilterApplication(opts, []common.Address{}, []common.Address{})
	if err != nil {
		log.Errorf("Error getting event _Application: %v", err)
		return err
	}
	nextApplication := itr.Next()
	for nextApplication {
		civilEvent := model.NewCivilEvent("_Application", itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextApplication = itr.Next()
	}
	return nil
}

// RetrieveApplicationRemoved gets _ApplicationRemoved events
func RetrieveApplicationRemoved(opts *bind.FilterOpts, civilTCR *tcr.CivilTCRContract, pastEvents *[]model.CivilEvent) error {
	itr, err := civilTCR.FilterApplicationRemoved(opts, []common.Address{})
	if err != nil {
		log.Errorf("Error getting event _ApplicationRemoved: %v", err)
		return err
	}

	nextApplication := itr.Next()
	for nextApplication {
		civilEvent := model.NewCivilEvent("_ApplicationRemoved", itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextApplication = itr.Next()
	}
	return nil
}

// RetrieveApplicationWhitelisted gets _ApplicationWhitelisted events
func RetrieveApplicationWhitelisted(opts *bind.FilterOpts, civilTCR *tcr.CivilTCRContract, pastEvents *[]model.CivilEvent) error {
	itr, err := civilTCR.FilterApplicationWhitelisted(opts, []common.Address{})
	if err != nil {
		log.Errorf("Error getting event _ApplicationWhitelisted: %v", err)
		return err
	}
	nextApplication := itr.Next()
	for nextApplication {
		civilEvent := model.NewCivilEvent("_ApplicationWhitelisted", itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextApplication = itr.Next()
	}
	return nil
}

// RetrieveChallenge gets _Challenge events
func RetrieveChallenge(opts *bind.FilterOpts, civilTCR *tcr.CivilTCRContract, pastEvents *[]model.CivilEvent) error {
	itr, err := civilTCR.FilterChallenge(opts, []common.Address{}, []common.Address{})
	if err != nil {
		log.Errorf("Error getting event _Challenge: %v", err)
		return err
	}
	nextApplication := itr.Next()
	for nextApplication {
		civilEvent := model.NewCivilEvent("_Challenge", itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextApplication = itr.Next()
	}
	return nil
}
