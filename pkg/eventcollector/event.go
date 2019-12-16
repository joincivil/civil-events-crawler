package eventcollector

import (
	"math/big"

	log "github.com/golang/glog"
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
	eventType := inputs.event.EventType() // Debug, remove later
	hash := inputs.event.Hash()           // Debug, remove later
	txHash := inputs.event.TxHash()       // Debug, remove later
	log.Infof("handleEvent: handling event: %v, tx: %v, hsh: %v", eventType,
		txHash.Hex(), hash) // Debug, remove later
	event := inputs.event

	err := c.updateEventTimeFromBlockHeader(event)
	if err != nil {
		return errors.WithMessage(err, "error updating date for event")
	}
	log.Infof("handleEvent: updated event time from block header: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

	err = c.eventDataPersister.SaveEvents([]*model.Event{event})
	if err != nil {
		return errors.WithMessage(err, "error saving events")
	}
	log.Infof("handleEvent: events saved: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

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
	log.Infof("handleEvent: updated last block data: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

	// Call event triggers
	err = c.callTriggers(event)
	if err != nil {
		return errors.WithMessage(err, "error calling triggers")
	}
	log.Infof("handleEvent: triggers called: %v, tx: %v, hsh: %v",
		eventType, txHash.Hex(), hash) // Debug, remove later

	// We need to get past newsroom events for the newsroom contract of a newly added watcher
	if event.EventType() == "Application" {
		newsroomAddr := event.EventPayload()["ListingAddress"].(common.Address)
		newsroomEvents, err := c.FilterAddedNewsroomContract(newsroomAddr)
		if err != nil {
			return errors.WithMessage(err, "error filtering new newsroom contract")
		}
		log.Infof("Found %v newsroom events for address %v after filtering: hsh: %v",
			len(newsroomEvents), newsroomAddr.Hex(), hash) // Debug, remove later
		err = c.eventDataPersister.SaveEvents(newsroomEvents)
		if err != nil {
			return errors.WithMessage(err, "error saving events for application")
		}
		log.Infof("Saved newsroom events at address %v, hsh: %v", newsroomAddr.Hex(), hash) //Debug, remove later

		if c.crawlerPubSub != nil {
			err := c.crawlerPubSub.PublishNewsroomExceptionMessage(newsroomAddr.Hex())
			if err != nil {
				return errors.WithMessagef(err, "error sending message for event %v to pubsub", event.Hash())
			}
		}
	}

	log.Infof("handleEvent: done: %v, tx: %v, hsh: %v", eventType, txHash.Hex(), hash) // Debug, remove later
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

		log.Infof(
			"updateEventTimeFromBlockHeader: calling headerbynumber: %v, %v",
			event.BlockNumber(),
			blockNum.Int64(),
		) // Debug, remove later

		header, err = c.retryChain.HeaderByNumberWithRetry(event.BlockNumber(), 10, 500)
		if err == nil && header != nil {
			log.Infof(
				"updateEventTimeFromBlockHeader: done calling headerbynumber: %v",
				header.TxHash.Hex(),
			) // Debug, remove later
		}

		c.headerCache.AddHeader(event.BlockNumber(), header)
	}
	if err != nil {
		return errors.Wrap(err, "error update event time")
	}
	event.SetTimestamp(int64(header.Time))
	return nil
}
