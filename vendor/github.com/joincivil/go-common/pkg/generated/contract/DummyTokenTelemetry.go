// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DummyTokenTelemetryContractABI is the input ABI used to generate the binding from.
const DummyTokenTelemetryContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"onRequestVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DummyTokenTelemetryContractBin is the compiled bytecode used for deploying new contracts.
const DummyTokenTelemetryContractBin = `0x608060405234801561001057600080fd5b5060a38061001f6000396000f300608060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637252487381146043575b600080fd5b348015604e57600080fd5b50607173ffffffffffffffffffffffffffffffffffffffff600435166024356073565b005b50505600a165627a7a723058207ad0f6acebe5f8699f63ec66e456db01d1e29bed089ebf940cec987824831bf60029`

// DeployDummyTokenTelemetryContract deploys a new Ethereum contract, binding an instance of DummyTokenTelemetryContract to it.
func DeployDummyTokenTelemetryContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DummyTokenTelemetryContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DummyTokenTelemetryContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DummyTokenTelemetryContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DummyTokenTelemetryContract{DummyTokenTelemetryContractCaller: DummyTokenTelemetryContractCaller{contract: contract}, DummyTokenTelemetryContractTransactor: DummyTokenTelemetryContractTransactor{contract: contract}, DummyTokenTelemetryContractFilterer: DummyTokenTelemetryContractFilterer{contract: contract}}, nil
}

// DummyTokenTelemetryContract is an auto generated Go binding around an Ethereum contract.
type DummyTokenTelemetryContract struct {
	DummyTokenTelemetryContractCaller     // Read-only binding to the contract
	DummyTokenTelemetryContractTransactor // Write-only binding to the contract
	DummyTokenTelemetryContractFilterer   // Log filterer for contract events
}

// DummyTokenTelemetryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DummyTokenTelemetryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyTokenTelemetryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DummyTokenTelemetryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyTokenTelemetryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DummyTokenTelemetryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyTokenTelemetryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DummyTokenTelemetryContractSession struct {
	Contract     *DummyTokenTelemetryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DummyTokenTelemetryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DummyTokenTelemetryContractCallerSession struct {
	Contract *DummyTokenTelemetryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// DummyTokenTelemetryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DummyTokenTelemetryContractTransactorSession struct {
	Contract     *DummyTokenTelemetryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// DummyTokenTelemetryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DummyTokenTelemetryContractRaw struct {
	Contract *DummyTokenTelemetryContract // Generic contract binding to access the raw methods on
}

// DummyTokenTelemetryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DummyTokenTelemetryContractCallerRaw struct {
	Contract *DummyTokenTelemetryContractCaller // Generic read-only contract binding to access the raw methods on
}

// DummyTokenTelemetryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DummyTokenTelemetryContractTransactorRaw struct {
	Contract *DummyTokenTelemetryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDummyTokenTelemetryContract creates a new instance of DummyTokenTelemetryContract, bound to a specific deployed contract.
func NewDummyTokenTelemetryContract(address common.Address, backend bind.ContractBackend) (*DummyTokenTelemetryContract, error) {
	contract, err := bindDummyTokenTelemetryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DummyTokenTelemetryContract{DummyTokenTelemetryContractCaller: DummyTokenTelemetryContractCaller{contract: contract}, DummyTokenTelemetryContractTransactor: DummyTokenTelemetryContractTransactor{contract: contract}, DummyTokenTelemetryContractFilterer: DummyTokenTelemetryContractFilterer{contract: contract}}, nil
}

// NewDummyTokenTelemetryContractCaller creates a new read-only instance of DummyTokenTelemetryContract, bound to a specific deployed contract.
func NewDummyTokenTelemetryContractCaller(address common.Address, caller bind.ContractCaller) (*DummyTokenTelemetryContractCaller, error) {
	contract, err := bindDummyTokenTelemetryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DummyTokenTelemetryContractCaller{contract: contract}, nil
}

// NewDummyTokenTelemetryContractTransactor creates a new write-only instance of DummyTokenTelemetryContract, bound to a specific deployed contract.
func NewDummyTokenTelemetryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DummyTokenTelemetryContractTransactor, error) {
	contract, err := bindDummyTokenTelemetryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DummyTokenTelemetryContractTransactor{contract: contract}, nil
}

// NewDummyTokenTelemetryContractFilterer creates a new log filterer instance of DummyTokenTelemetryContract, bound to a specific deployed contract.
func NewDummyTokenTelemetryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DummyTokenTelemetryContractFilterer, error) {
	contract, err := bindDummyTokenTelemetryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DummyTokenTelemetryContractFilterer{contract: contract}, nil
}

// bindDummyTokenTelemetryContract binds a generic wrapper to an already deployed contract.
func bindDummyTokenTelemetryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DummyTokenTelemetryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyTokenTelemetryContract *DummyTokenTelemetryContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DummyTokenTelemetryContract.Contract.DummyTokenTelemetryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyTokenTelemetryContract *DummyTokenTelemetryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyTokenTelemetryContract.Contract.DummyTokenTelemetryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyTokenTelemetryContract *DummyTokenTelemetryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyTokenTelemetryContract.Contract.DummyTokenTelemetryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyTokenTelemetryContract *DummyTokenTelemetryContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DummyTokenTelemetryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyTokenTelemetryContract *DummyTokenTelemetryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyTokenTelemetryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyTokenTelemetryContract *DummyTokenTelemetryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyTokenTelemetryContract.Contract.contract.Transact(opts, method, params...)
}

// OnRequestVotingRights is a paid mutator transaction binding the contract method 0x72524873.
//
// Solidity: function onRequestVotingRights(user address, tokenAmount uint256) returns()
func (_DummyTokenTelemetryContract *DummyTokenTelemetryContractTransactor) OnRequestVotingRights(opts *bind.TransactOpts, user common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _DummyTokenTelemetryContract.contract.Transact(opts, "onRequestVotingRights", user, tokenAmount)
}

// OnRequestVotingRights is a paid mutator transaction binding the contract method 0x72524873.
//
// Solidity: function onRequestVotingRights(user address, tokenAmount uint256) returns()
func (_DummyTokenTelemetryContract *DummyTokenTelemetryContractSession) OnRequestVotingRights(user common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _DummyTokenTelemetryContract.Contract.OnRequestVotingRights(&_DummyTokenTelemetryContract.TransactOpts, user, tokenAmount)
}

// OnRequestVotingRights is a paid mutator transaction binding the contract method 0x72524873.
//
// Solidity: function onRequestVotingRights(user address, tokenAmount uint256) returns()
func (_DummyTokenTelemetryContract *DummyTokenTelemetryContractTransactorSession) OnRequestVotingRights(user common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _DummyTokenTelemetryContract.Contract.OnRequestVotingRights(&_DummyTokenTelemetryContract.TransactOpts, user, tokenAmount)
}
