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

// NoOpTokenControllerContractABI is the input ABI used to generate the binding from.
const NoOpTokenControllerContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"SUCCESS_CODE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SUCCESS_MESSAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"detectTransferRestriction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"restrictionCode\",\"type\":\"uint8\"}],\"name\":\"messageForTransferRestriction\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NoOpTokenControllerContractBin is the compiled bytecode used for deploying new contracts.
const NoOpTokenControllerContractBin = `0x608060405234801561001057600080fd5b50610216806100206000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630e969a0581146100665780637f4ab1dd14610091578063d4ce141514610121578063e7984d1714610158575b600080fd5b34801561007257600080fd5b5061007b61016d565b6040805160ff9092168252519081900360200190f35b34801561009d57600080fd5b506100ac60ff60043516610172565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100e65781810151838201526020016100ce565b50505050905090810190601f1680156101135780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561012d57600080fd5b5061007b73ffffffffffffffffffffffffffffffffffffffff600435811690602435166044356101aa565b34801561016457600080fd5b506100ac6101b3565b600081565b5060408051808201909152600781527f5355434345535300000000000000000000000000000000000000000000000000602082015290565b60009392505050565b60408051808201909152600781527f53554343455353000000000000000000000000000000000000000000000000006020820152815600a165627a7a723058208c703e993ff1ffb3530dbb00e97983c979a151d69d4afb9055385bc59af53e1f0029`

// DeployNoOpTokenControllerContract deploys a new Ethereum contract, binding an instance of NoOpTokenControllerContract to it.
func DeployNoOpTokenControllerContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NoOpTokenControllerContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NoOpTokenControllerContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NoOpTokenControllerContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NoOpTokenControllerContract{NoOpTokenControllerContractCaller: NoOpTokenControllerContractCaller{contract: contract}, NoOpTokenControllerContractTransactor: NoOpTokenControllerContractTransactor{contract: contract}, NoOpTokenControllerContractFilterer: NoOpTokenControllerContractFilterer{contract: contract}}, nil
}

// NoOpTokenControllerContract is an auto generated Go binding around an Ethereum contract.
type NoOpTokenControllerContract struct {
	NoOpTokenControllerContractCaller     // Read-only binding to the contract
	NoOpTokenControllerContractTransactor // Write-only binding to the contract
	NoOpTokenControllerContractFilterer   // Log filterer for contract events
}

// NoOpTokenControllerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type NoOpTokenControllerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOpTokenControllerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NoOpTokenControllerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOpTokenControllerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NoOpTokenControllerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOpTokenControllerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NoOpTokenControllerContractSession struct {
	Contract     *NoOpTokenControllerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NoOpTokenControllerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NoOpTokenControllerContractCallerSession struct {
	Contract *NoOpTokenControllerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// NoOpTokenControllerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NoOpTokenControllerContractTransactorSession struct {
	Contract     *NoOpTokenControllerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// NoOpTokenControllerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type NoOpTokenControllerContractRaw struct {
	Contract *NoOpTokenControllerContract // Generic contract binding to access the raw methods on
}

// NoOpTokenControllerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NoOpTokenControllerContractCallerRaw struct {
	Contract *NoOpTokenControllerContractCaller // Generic read-only contract binding to access the raw methods on
}

// NoOpTokenControllerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NoOpTokenControllerContractTransactorRaw struct {
	Contract *NoOpTokenControllerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNoOpTokenControllerContract creates a new instance of NoOpTokenControllerContract, bound to a specific deployed contract.
func NewNoOpTokenControllerContract(address common.Address, backend bind.ContractBackend) (*NoOpTokenControllerContract, error) {
	contract, err := bindNoOpTokenControllerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoOpTokenControllerContract{NoOpTokenControllerContractCaller: NoOpTokenControllerContractCaller{contract: contract}, NoOpTokenControllerContractTransactor: NoOpTokenControllerContractTransactor{contract: contract}, NoOpTokenControllerContractFilterer: NoOpTokenControllerContractFilterer{contract: contract}}, nil
}

// NewNoOpTokenControllerContractCaller creates a new read-only instance of NoOpTokenControllerContract, bound to a specific deployed contract.
func NewNoOpTokenControllerContractCaller(address common.Address, caller bind.ContractCaller) (*NoOpTokenControllerContractCaller, error) {
	contract, err := bindNoOpTokenControllerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpTokenControllerContractCaller{contract: contract}, nil
}

// NewNoOpTokenControllerContractTransactor creates a new write-only instance of NoOpTokenControllerContract, bound to a specific deployed contract.
func NewNoOpTokenControllerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*NoOpTokenControllerContractTransactor, error) {
	contract, err := bindNoOpTokenControllerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpTokenControllerContractTransactor{contract: contract}, nil
}

// NewNoOpTokenControllerContractFilterer creates a new log filterer instance of NoOpTokenControllerContract, bound to a specific deployed contract.
func NewNoOpTokenControllerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*NoOpTokenControllerContractFilterer, error) {
	contract, err := bindNoOpTokenControllerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoOpTokenControllerContractFilterer{contract: contract}, nil
}

// bindNoOpTokenControllerContract binds a generic wrapper to an already deployed contract.
func bindNoOpTokenControllerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NoOpTokenControllerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoOpTokenControllerContract *NoOpTokenControllerContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NoOpTokenControllerContract.Contract.NoOpTokenControllerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoOpTokenControllerContract *NoOpTokenControllerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpTokenControllerContract.Contract.NoOpTokenControllerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoOpTokenControllerContract *NoOpTokenControllerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpTokenControllerContract.Contract.NoOpTokenControllerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoOpTokenControllerContract *NoOpTokenControllerContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NoOpTokenControllerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoOpTokenControllerContract *NoOpTokenControllerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpTokenControllerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoOpTokenControllerContract *NoOpTokenControllerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpTokenControllerContract.Contract.contract.Transact(opts, method, params...)
}

// SUCCESSCODE is a free data retrieval call binding the contract method 0x0e969a05.
//
// Solidity: function SUCCESS_CODE() constant returns(uint8)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractCaller) SUCCESSCODE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _NoOpTokenControllerContract.contract.Call(opts, out, "SUCCESS_CODE")
	return *ret0, err
}

// SUCCESSCODE is a free data retrieval call binding the contract method 0x0e969a05.
//
// Solidity: function SUCCESS_CODE() constant returns(uint8)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractSession) SUCCESSCODE() (uint8, error) {
	return _NoOpTokenControllerContract.Contract.SUCCESSCODE(&_NoOpTokenControllerContract.CallOpts)
}

// SUCCESSCODE is a free data retrieval call binding the contract method 0x0e969a05.
//
// Solidity: function SUCCESS_CODE() constant returns(uint8)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractCallerSession) SUCCESSCODE() (uint8, error) {
	return _NoOpTokenControllerContract.Contract.SUCCESSCODE(&_NoOpTokenControllerContract.CallOpts)
}

// SUCCESSMESSAGE is a free data retrieval call binding the contract method 0xe7984d17.
//
// Solidity: function SUCCESS_MESSAGE() constant returns(string)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractCaller) SUCCESSMESSAGE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NoOpTokenControllerContract.contract.Call(opts, out, "SUCCESS_MESSAGE")
	return *ret0, err
}

// SUCCESSMESSAGE is a free data retrieval call binding the contract method 0xe7984d17.
//
// Solidity: function SUCCESS_MESSAGE() constant returns(string)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractSession) SUCCESSMESSAGE() (string, error) {
	return _NoOpTokenControllerContract.Contract.SUCCESSMESSAGE(&_NoOpTokenControllerContract.CallOpts)
}

// SUCCESSMESSAGE is a free data retrieval call binding the contract method 0xe7984d17.
//
// Solidity: function SUCCESS_MESSAGE() constant returns(string)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractCallerSession) SUCCESSMESSAGE() (string, error) {
	return _NoOpTokenControllerContract.Contract.SUCCESSMESSAGE(&_NoOpTokenControllerContract.CallOpts)
}

// DetectTransferRestriction is a free data retrieval call binding the contract method 0xd4ce1415.
//
// Solidity: function detectTransferRestriction(from address, to address, value uint256) constant returns(uint8)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractCaller) DetectTransferRestriction(opts *bind.CallOpts, from common.Address, to common.Address, value *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _NoOpTokenControllerContract.contract.Call(opts, out, "detectTransferRestriction", from, to, value)
	return *ret0, err
}

// DetectTransferRestriction is a free data retrieval call binding the contract method 0xd4ce1415.
//
// Solidity: function detectTransferRestriction(from address, to address, value uint256) constant returns(uint8)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractSession) DetectTransferRestriction(from common.Address, to common.Address, value *big.Int) (uint8, error) {
	return _NoOpTokenControllerContract.Contract.DetectTransferRestriction(&_NoOpTokenControllerContract.CallOpts, from, to, value)
}

// DetectTransferRestriction is a free data retrieval call binding the contract method 0xd4ce1415.
//
// Solidity: function detectTransferRestriction(from address, to address, value uint256) constant returns(uint8)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractCallerSession) DetectTransferRestriction(from common.Address, to common.Address, value *big.Int) (uint8, error) {
	return _NoOpTokenControllerContract.Contract.DetectTransferRestriction(&_NoOpTokenControllerContract.CallOpts, from, to, value)
}

// MessageForTransferRestriction is a free data retrieval call binding the contract method 0x7f4ab1dd.
//
// Solidity: function messageForTransferRestriction(restrictionCode uint8) constant returns(string)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractCaller) MessageForTransferRestriction(opts *bind.CallOpts, restrictionCode uint8) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NoOpTokenControllerContract.contract.Call(opts, out, "messageForTransferRestriction", restrictionCode)
	return *ret0, err
}

// MessageForTransferRestriction is a free data retrieval call binding the contract method 0x7f4ab1dd.
//
// Solidity: function messageForTransferRestriction(restrictionCode uint8) constant returns(string)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractSession) MessageForTransferRestriction(restrictionCode uint8) (string, error) {
	return _NoOpTokenControllerContract.Contract.MessageForTransferRestriction(&_NoOpTokenControllerContract.CallOpts, restrictionCode)
}

// MessageForTransferRestriction is a free data retrieval call binding the contract method 0x7f4ab1dd.
//
// Solidity: function messageForTransferRestriction(restrictionCode uint8) constant returns(string)
func (_NoOpTokenControllerContract *NoOpTokenControllerContractCallerSession) MessageForTransferRestriction(restrictionCode uint8) (string, error) {
	return _NoOpTokenControllerContract.Contract.MessageForTransferRestriction(&_NoOpTokenControllerContract.CallOpts, restrictionCode)
}
