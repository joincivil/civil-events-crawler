package eth

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/go-common/pkg/time"
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
	if time.CurrentEpochSecsInInt64()-blockItem.addedTs > b.expiryPeriodSecs {
		delete(b.cache, num)
		return nil
	}
	return blockItem.header
}

// AddHeader adds a types.Block to the map with the key being block number.
func (b *BlockHeaderCache) AddHeader(num uint64, header *types.Header) {
	b.cache[num] = &blockHeaderCacheItem{
		addedTs: time.CurrentEpochSecsInInt64(),
		header:  header,
	}
}
