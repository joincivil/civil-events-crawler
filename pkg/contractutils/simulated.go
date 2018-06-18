// Package contractutils contains utilities related to smart contracts
// and testing smart contracts
package contractutils

import (
	// log "github.com/golang/glog"
	"context"
	"reflect"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/event"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

// newSimulatedBackendWithGasLimit adjusts the underlying unexported fields of simulated.Backend in go-ethereum
// to allow updates to the gasLimit for the Genesis block.  This will allow us increase the block gas limit
// to test the events and the contracts.
// NOTE(PN): FOR TESTING EVENTS ONLY!  Will break on updates to the go-ethereum repo.
func newSimulatedBackendWithGasLimit(alloc core.GenesisAlloc, gasLimit uint64) *backends.SimulatedBackend {
	backend := backends.NewSimulatedBackend(alloc)
	database := ethdb.NewMemDatabase()
	genesis := core.Genesis{Config: params.AllEthashProtocolChanges, Alloc: alloc, GasLimit: gasLimit}
	genesis.MustCommit(database)

	blockchain, _ := core.NewBlockChain(database, nil, genesis.Config, ethash.NewFaker(), vm.Config{}) // nolint: gas
	events := filters.NewEventSystem(new(event.TypeMux), &filterBackend{database, blockchain}, false)  //nolint: megacheck, staticcheck

	be := reflect.ValueOf(backend)
	field := be.Elem().FieldByName("database")
	newVal := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem() //nolint: gas
	newVal.Set(reflect.ValueOf(database))

	field = be.Elem().FieldByName("blockchain")
	newVal = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem() //nolint: gas
	newVal.Set(reflect.ValueOf(blockchain))

	field = be.Elem().FieldByName("config")
	newVal = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem() //nolint: gas
	newVal.Set(reflect.ValueOf(genesis.Config))

	field = be.Elem().FieldByName("events")
	newVal = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem() //nolint: gas
	newVal.Set(reflect.ValueOf(events))

	return backend
}

// filterBackend is a copy of the filterBackend from simulated.Backend in go-ethereum
// NOTE(PN): FOR TESTING THE EVENTS ONLY!  Will break on updates to the go-ethereum repo.
type filterBackend struct {
	db ethdb.Database
	bc *core.BlockChain
}

func (fb *filterBackend) ChainDb() ethdb.Database       { return fb.db }
func (fb *filterBackend) EventMux() *event.TypeMux      { panic("not supported") } //nolint: megacheck, staticcheck
func (fb *filterBackend) BloomStatus() (uint64, uint64) { return 4096, 0 }

func (fb *filterBackend) GetLogs(ctx context.Context, hash common.Hash) ([][]*types.Log, error) {
	number := rawdb.ReadHeaderNumber(fb.db, hash)
	if number == nil {
		return nil, nil
	}
	receipts := rawdb.ReadReceipts(fb.db, hash, *number)
	if receipts == nil {
		return nil, nil
	}
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	return logs, nil
}

func (fb *filterBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	number := rawdb.ReadHeaderNumber(fb.db, hash)
	if number == nil {
		return nil, nil
	}
	return rawdb.ReadReceipts(fb.db, hash, *number), nil
}

func (fb *filterBackend) HeaderByNumber(ctx context.Context, block rpc.BlockNumber) (*types.Header, error) {
	if block == rpc.LatestBlockNumber {
		return fb.bc.CurrentHeader(), nil
	}
	return fb.bc.GetHeaderByNumber(uint64(block.Int64())), nil
}

func (fb *filterBackend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		<-quit
		return nil
	})
}

func (fb *filterBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return fb.bc.SubscribeChainEvent(ch)
}
func (fb *filterBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return fb.bc.SubscribeRemovedLogsEvent(ch)
}
func (fb *filterBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return fb.bc.SubscribeLogsEvent(ch)
}

func (fb *filterBackend) ServiceFilter(ctx context.Context, ms *bloombits.MatcherSession) {
	panic("not supported")
}
