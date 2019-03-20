package eventcollector

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/golang/glog"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/go-common/pkg/eth"
	"github.com/pkg/errors"
)

// UpdateEventTimesFromBlockHeaders sets the event time from the block header
// for a slice of Events
func UpdateEventTimesFromBlockHeaders(events []*model.Event, retryChain *eth.RetryChainReader,
	headerCache *eth.BlockHeaderCache) error {
	for _, event := range events {
		err := UpdateEventTimeFromBlockHeader(event, retryChain, headerCache)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateEventTimeFromBlockHeader sets the event time from the block header for an Event
func UpdateEventTimeFromBlockHeader(event *model.Event, retryChain *eth.RetryChainReader,
	headerCache *eth.BlockHeaderCache) error {
	var header *types.Header
	var err error

	inCache := false
	if headerCache == nil {
		headerCache = eth.NewBlockHeaderCache(blockHeaderExpirySecs)
	} else {
		header = headerCache.HeaderByBlockNumber(event.BlockNumber())
		if header != nil {
			inCache = true
		}
	}
	if !inCache {
		blockNum := big.NewInt(0)
		blockNum.SetUint64(event.BlockNumber())

		log.Infof(
			"updateEventTimeFromBlockHeader: calling headerbynumber: %v, %v",
			event.BlockNumber(),
			blockNum.Int64(),
		) // Debug, remove later
		header, err = retryChain.HeaderByNumberWithRetry(event.BlockNumber(), 10, 500)

		log.Infof(
			"updateEventTimeFromBlockHeader: done calling headerbynumber: %v",
			header.TxHash.Hex(),
		) // Debug, remove later

		headerCache.AddHeader(event.BlockNumber(), header)
	}
	if err != nil {
		return errors.Wrap(err, "error update event time")
	}

	event.SetTimestamp(header.Time.Int64())
	return nil
}
