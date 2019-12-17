package eventcollector

import (
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/go-common/pkg/eth"
)

type handleEventInputs struct {
	event  *model.Event
	errors chan<- error
}

// handleEvent is the func used for the goroutine pool that handles
// incoming events from the watchers
func (c *EventCollector) handleEvent(payload interface{}) interface{} {
	inputs := payload.(handleEventInputs)
	event := inputs.event

	err := c.updateEventTimeFromBlockHeader(event)
	if err != nil {
		return errors.WithMessage(err, "error updating date for event")
	}

	errs := c.eventDataPersister.SaveEvents([]*model.Event{event})
	if len(errs) > 0 {
		return errors.WithMessage(errs[0], "error saving events")
	}

	if c.crawlerPubSub != nil {
		err = c.crawlerPubSub.PublishProcessorTriggerMessage()
		if err != nil {
			return errors.WithMessagef(err, "error sending message for event %v to pubsub", event.Hash())
		}
	}

	// Update last block in persistence in case of error
	err = c.listenerPersister.UpdateLastBlockData([]*model.Event{event})
	if err != nil {
		return errors.WithMessage(err, "error updating last block")
	}

	// Call event triggers
	err = c.callTriggers(event)
	if err != nil {
		return errors.WithMessage(err, "error calling triggers")
	}

	// We need to get past newsroom events for the newsroom contract of a newly added watcher
	if event.EventType() == "Application" {
		newsroomAddr := event.EventPayload()["ListingAddress"].(common.Address)
		newsroomEvents, err := c.FilterAddedNewsroomContract(newsroomAddr)
		if err != nil {
			return errors.WithMessage(err, "error filtering new newsroom contract")
		}

		errs := c.eventDataPersister.SaveEvents(newsroomEvents)
		if len(errs) > 0 {
			return errors.WithMessage(errs[0], "error saving events for application")
		}

		if c.crawlerPubSub != nil {
			err := c.crawlerPubSub.PublishNewsroomExceptionMessage(newsroomAddr.Hex())
			if err != nil {
				return errors.WithMessagef(err, "error sending message for event %v to pubsub", event.Hash())
			}
		}
	}

	return nil
}

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
