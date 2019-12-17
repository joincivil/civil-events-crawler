package eventcollector

import (
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/go-common/pkg/eth"
)

func (c *EventCollector) updateEventTimesFromBlockHeaders(events []*model.Event) error {
	for _, event := range events {
		err := c.updateEventTimeFromBlockHeader(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *EventCollector) updateEventTimeFromBlockHeader(event *model.Event) error {
	var header *types.Header
	var err error

	inCache := false
	if c.headerCache == nil {
		c.headerCache = eth.NewBlockHeaderCache(blockHeaderExpirySecs)
	} else {
		header = c.headerCache.HeaderByBlockNumber(event.BlockNumber())
		if header != nil {
			inCache = true
		}
	}
	if !inCache {
		blockNum := big.NewInt(0)
		blockNum.SetUint64(event.BlockNumber())
		header, err = c.retryChain.HeaderByNumberWithRetry(event.BlockNumber(), 10, 500)
		c.headerCache.AddHeader(event.BlockNumber(), header)
	}
	if err != nil {
		return errors.Wrap(err, "error update event time")
	}
	event.SetTimestamp(int64(header.Time))
	return nil
}
