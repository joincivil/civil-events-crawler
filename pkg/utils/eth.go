package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
