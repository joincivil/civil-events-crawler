package eth

import (
	"context"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/golang/glog"
)

// RetryChainReader is a ChainReader that includes some version of ChainReader
// functions that are wrapped by a retry mechanism
type RetryChainReader struct {
	ethereum.ChainReader
}

// HeaderByNumberWithRetry is a version of HeaderByNumber that has a retry
// mechanism
func (r *RetryChainReader) HeaderByNumberWithRetry(blockNumber uint64, maxAttempts int,
	baseWaitMs int) (*types.Header, error) {
	blockNum := big.NewInt(0)
	blockNum.SetUint64(blockNumber)

	attempt := 1
	var header *types.Header
	var err error
	for {
		header, err = r.HeaderByNumber(context.Background(), blockNum)
		if err != nil {
			if err != ethereum.NotFound {
				return nil, err
			}

			log.Infof(
				"block not found, sleep/attempt again, waiting %v ms...",
				baseWaitMs*attempt,
			)

			// Take a break and see if the block can be found
			time.Sleep(time.Duration(baseWaitMs) * time.Duration(attempt) * time.Millisecond)
			if attempt > maxAttempts {
				return nil, err
			}

			attempt++
			continue
		}
		break
	}
	return header, nil
}
