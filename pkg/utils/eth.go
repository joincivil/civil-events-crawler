package utils

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/golang/glog"
)

type blockHeaderCacheItem struct {
	addedTs int64
	header  *types.Header
}

// NewBlockHeaderCache is a convenience function to init a BlockHeaderCache
func NewBlockHeaderCache(expirySecs int64) *BlockHeaderCache {
	return &BlockHeaderCache{
		cache:            map[uint64]*blockHeaderCacheItem{},
		expiryPeriodSecs: expirySecs,
	}
}

// BlockHeaderCache stores a memory map of blockNumber to the types.Block.  Used
// as a cache to prevent multiple access via the ethclient
type BlockHeaderCache struct {
	cache            map[uint64]*blockHeaderCacheItem
	expiryPeriodSecs int64
}

// HeaderByBlockNumber returns a types.Block if it exists in the map given
// the block number. Returns nil if not found in map.
// If block was added more than expiryPeriodSecs ago, will return nil.
func (b *BlockHeaderCache) HeaderByBlockNumber(num uint64) *types.Header {
	blockItem, ok := b.cache[num]
	if !ok {
		return nil
	}
	if CurrentEpochSecsInInt64()-blockItem.addedTs > b.expiryPeriodSecs {
		delete(b.cache, num)
		return nil
	}
	return blockItem.header
}

// AddHeader adds a types.Block to the map with the key being block number.
func (b *BlockHeaderCache) AddHeader(num uint64, header *types.Header) {
	b.cache[num] = &blockHeaderCacheItem{
		addedTs: CurrentEpochSecsInInt64(),
		header:  header,
	}
}

// NormalizeEthAddress takes a string address to normalize the
// case of the ethereum address when it is a string.
// Runs through common.Address.Hex().
func NormalizeEthAddress(addr string) string {
	address := common.HexToAddress(addr)
	return address.Hex()
}

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
