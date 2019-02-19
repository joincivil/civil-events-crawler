// Code generated by 'gen/handlerlistgen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// File was generated at 2019-02-19 22:10:51.44607 +0000 UTC
package handlerlist

import (
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/common"

	"github.com/joincivil/civil-events-crawler/pkg/generated/filterer"
	"github.com/joincivil/civil-events-crawler/pkg/generated/watcher"
	"github.com/joincivil/civil-events-crawler/pkg/model"
)

func ContractFilterers(nameToAddrs map[string][]common.Address) []model.ContractFilterers {
	filters := []model.ContractFilterers{}

	var addrs []common.Address
	var addr common.Address
	var ok bool

	addrs, ok = nameToAddrs["civiltcr"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewCivilTCRContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added CivilTCRContract filterer")
		}
	}

	addrs, ok = nameToAddrs["newsroom"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewNewsroomContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added NewsroomContract filterer")
		}
	}

	addrs, ok = nameToAddrs["civilplcrvoting"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewCivilPLCRVotingContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added CivilPLCRVotingContract filterer")
		}
	}

	addrs, ok = nameToAddrs["cvltoken"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewCVLTokenContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added CVLTokenContract filterer")
		}
	}

	return filters
}

func ContractWatchers(nameToAddrs map[string][]common.Address) []model.ContractWatchers {
	watchers := []model.ContractWatchers{}

	var addrs []common.Address
	var addr common.Address
	var ok bool

	addrs, ok = nameToAddrs["civiltcr"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewCivilTCRContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added CivilTCRContract watcher")
		}
	}

	addrs, ok = nameToAddrs["newsroom"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewNewsroomContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added NewsroomContract watcher")
		}
	}

	addrs, ok = nameToAddrs["civilplcrvoting"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewCivilPLCRVotingContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added CivilPLCRVotingContract watcher")
		}
	}

	addrs, ok = nameToAddrs["cvltoken"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewCVLTokenContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added CVLTokenContract watcher")
		}
	}

	return watchers
}
