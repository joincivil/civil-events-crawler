// Package time_test contains tests for the eth utils
package utils_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

func TestBlockHeaderCache(t *testing.T) {
	expiryCacheSecs := 1
	cache := utils.NewBlockHeaderCache(int64(expiryCacheSecs))
	header := cache.HeaderByBlockNumber(uint64(1))
	if header != nil {
		t.Error("Should have failed to retrieve any headers")
	}

	ts := big.NewInt(utils.CurrentEpochSecsInInt64())
	header1 := &types.Header{
		Time: ts,
	}

	cache.AddHeader(uint64(1), header1)

	header = cache.HeaderByBlockNumber(uint64(1))
	if header == nil {
		t.Error("Should have retrieved a header")
	}
	if header.Time != ts {
		t.Error("Should have been the same time as the added Header")
	}

	time.Sleep(time.Duration(expiryCacheSecs+1) * time.Second)

	header = cache.HeaderByBlockNumber(uint64(1))
	if header != nil {
		t.Error("Should have not retrieved a header after duration")
	}
}

func TestNormalizeEthAddress(t *testing.T) {
	addr1 := "0x39eD84CE90Bc48DD76C4760DD0F90997Ba274F9d"
	addr2 := "0x39ed84ce90bc48dd76c4760dd0f90997ba274f9d"

	normalized1 := utils.NormalizeEthAddress(addr1)
	normalized2 := utils.NormalizeEthAddress(addr2)

	if normalized1 == "" {
		t.Errorf("Should have converted address correctly")
	}
	if normalized2 == "" {
		t.Errorf("Should have converted address correctly")
	}
	if normalized1 != normalized2 {
		t.Errorf("Addresses should have matched")
	}
}
