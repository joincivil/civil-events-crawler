// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// EIP20ContractABI is the input ABI used to generate the binding from.
const EIP20ContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialAmount\",\"type\":\"uint256\"},{\"name\":\"_tokenName\",\"type\":\"string\"},{\"name\":\"_decimalUnits\",\"type\":\"uint8\"},{\"name\":\"_tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EIP20ContractBin is the compiled bytecode used for deploying new contracts.
const EIP20ContractBin = `0x608060405234801561001057600080fd5b506040516106f83803806106f883398101604090815281516020808401518385015160608601513360009081526004855295862085905594849055908501805193959094919391019161006891600191860190610095565b506002805460ff191660ff8416179055805161008b906003906020840190610095565b5050505050610130565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100d657805160ff1916838001178555610103565b82800160010185558215610103579182015b828111156101035782518255916020019190600101906100e8565b5061010f929150610113565b5090565b61012d91905b8082111561010f5760008155600101610119565b90565b6105b98061013f6000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461009d578063095ea7b31461012757806318160ddd1461015f57806323b872dd14610186578063313ce567146101b057806370a08231146101db57806395d89b41146101fc578063a9059cbb14610211578063dd62ed3e14610235575b600080fd5b3480156100a957600080fd5b506100b261025c565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100ec5781810151838201526020016100d4565b50505050905090810190601f1680156101195780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561013357600080fd5b5061014b600160a060020a03600435166024356102e9565b604080519115158252519081900360200190f35b34801561016b57600080fd5b5061017461034f565b60408051918252519081900360200190f35b34801561019257600080fd5b5061014b600160a060020a0360043581169060243516604435610355565b3480156101bc57600080fd5b506101c5610459565b6040805160ff9092168252519081900360200190f35b3480156101e757600080fd5b50610174600160a060020a0360043516610462565b34801561020857600080fd5b506100b261047d565b34801561021d57600080fd5b5061014b600160a060020a03600435166024356104d8565b34801561024157600080fd5b50610174600160a060020a0360043581169060243516610562565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102e15780601f106102b6576101008083540402835291602001916102e1565b820191906000526020600020905b8154815290600101906020018083116102c457829003601f168201915b505050505081565b336000818152600560209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b600160a060020a0383166000818152600560209081526040808320338452825280832054938352600490915281205490919083118015906103965750828110155b15156103a157600080fd5b600160a060020a038085166000908152600460205260408082208054870190559187168152208054849003905560001981101561040357600160a060020a03851660009081526005602090815260408083203384529091529020805484900390555b83600160a060020a031685600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a3506001949350505050565b60025460ff1681565b600160a060020a031660009081526004602052604090205490565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102e15780601f106102b6576101008083540402835291602001916102e1565b336000908152600460205260408120548211156104f457600080fd5b33600081815260046020908152604080832080548790039055600160a060020a03871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a350600192915050565b600160a060020a039182166000908152600560209081526040808320939094168252919091522054905600a165627a7a7230582009e03231304b1a3493c50f88ded5a37fbc32b612b1938db9f0f2e3ec9c4e0e840029`

// DeployEIP20Contract deploys a new Ethereum contract, binding an instance of EIP20Contract to it.
func DeployEIP20Contract(auth *bind.TransactOpts, backend bind.ContractBackend, _initialAmount *big.Int, _tokenName string, _decimalUnits uint8, _tokenSymbol string) (common.Address, *types.Transaction, *EIP20Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP20ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EIP20ContractBin), backend, _initialAmount, _tokenName, _decimalUnits, _tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EIP20Contract{EIP20ContractCaller: EIP20ContractCaller{contract: contract}, EIP20ContractTransactor: EIP20ContractTransactor{contract: contract}, EIP20ContractFilterer: EIP20ContractFilterer{contract: contract}}, nil
}

// EIP20Contract is an auto generated Go binding around an Ethereum contract.
type EIP20Contract struct {
	EIP20ContractCaller     // Read-only binding to the contract
	EIP20ContractTransactor // Write-only binding to the contract
	EIP20ContractFilterer   // Log filterer for contract events
}

// EIP20ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type EIP20ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP20ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP20ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP20ContractSession struct {
	Contract     *EIP20Contract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP20ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP20ContractCallerSession struct {
	Contract *EIP20ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EIP20ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP20ContractTransactorSession struct {
	Contract     *EIP20ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EIP20ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type EIP20ContractRaw struct {
	Contract *EIP20Contract // Generic contract binding to access the raw methods on
}

// EIP20ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP20ContractCallerRaw struct {
	Contract *EIP20ContractCaller // Generic read-only contract binding to access the raw methods on
}

// EIP20ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP20ContractTransactorRaw struct {
	Contract *EIP20ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP20Contract creates a new instance of EIP20Contract, bound to a specific deployed contract.
func NewEIP20Contract(address common.Address, backend bind.ContractBackend) (*EIP20Contract, error) {
	contract, err := bindEIP20Contract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP20Contract{EIP20ContractCaller: EIP20ContractCaller{contract: contract}, EIP20ContractTransactor: EIP20ContractTransactor{contract: contract}, EIP20ContractFilterer: EIP20ContractFilterer{contract: contract}}, nil
}

// NewEIP20ContractCaller creates a new read-only instance of EIP20Contract, bound to a specific deployed contract.
func NewEIP20ContractCaller(address common.Address, caller bind.ContractCaller) (*EIP20ContractCaller, error) {
	contract, err := bindEIP20Contract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP20ContractCaller{contract: contract}, nil
}

// NewEIP20ContractTransactor creates a new write-only instance of EIP20Contract, bound to a specific deployed contract.
func NewEIP20ContractTransactor(address common.Address, transactor bind.ContractTransactor) (*EIP20ContractTransactor, error) {
	contract, err := bindEIP20Contract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP20ContractTransactor{contract: contract}, nil
}

// NewEIP20ContractFilterer creates a new log filterer instance of EIP20Contract, bound to a specific deployed contract.
func NewEIP20ContractFilterer(address common.Address, filterer bind.ContractFilterer) (*EIP20ContractFilterer, error) {
	contract, err := bindEIP20Contract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP20ContractFilterer{contract: contract}, nil
}

// bindEIP20Contract binds a generic wrapper to an already deployed contract.
func bindEIP20Contract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP20ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP20Contract *EIP20ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EIP20Contract.Contract.EIP20ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP20Contract *EIP20ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP20Contract.Contract.EIP20ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP20Contract *EIP20ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP20Contract.Contract.EIP20ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP20Contract *EIP20ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EIP20Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP20Contract *EIP20ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP20Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP20Contract *EIP20ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP20Contract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_EIP20Contract *EIP20ContractCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EIP20Contract.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_EIP20Contract *EIP20ContractSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EIP20Contract.Contract.Allowance(&_EIP20Contract.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_EIP20Contract *EIP20ContractCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EIP20Contract.Contract.Allowance(&_EIP20Contract.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_EIP20Contract *EIP20ContractCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EIP20Contract.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_EIP20Contract *EIP20ContractSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _EIP20Contract.Contract.BalanceOf(&_EIP20Contract.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_EIP20Contract *EIP20ContractCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _EIP20Contract.Contract.BalanceOf(&_EIP20Contract.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_EIP20Contract *EIP20ContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _EIP20Contract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_EIP20Contract *EIP20ContractSession) Decimals() (uint8, error) {
	return _EIP20Contract.Contract.Decimals(&_EIP20Contract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_EIP20Contract *EIP20ContractCallerSession) Decimals() (uint8, error) {
	return _EIP20Contract.Contract.Decimals(&_EIP20Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EIP20Contract *EIP20ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EIP20Contract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EIP20Contract *EIP20ContractSession) Name() (string, error) {
	return _EIP20Contract.Contract.Name(&_EIP20Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EIP20Contract *EIP20ContractCallerSession) Name() (string, error) {
	return _EIP20Contract.Contract.Name(&_EIP20Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_EIP20Contract *EIP20ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EIP20Contract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_EIP20Contract *EIP20ContractSession) Symbol() (string, error) {
	return _EIP20Contract.Contract.Symbol(&_EIP20Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_EIP20Contract *EIP20ContractCallerSession) Symbol() (string, error) {
	return _EIP20Contract.Contract.Symbol(&_EIP20Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EIP20Contract *EIP20ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EIP20Contract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EIP20Contract *EIP20ContractSession) TotalSupply() (*big.Int, error) {
	return _EIP20Contract.Contract.TotalSupply(&_EIP20Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EIP20Contract *EIP20ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _EIP20Contract.Contract.TotalSupply(&_EIP20Contract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_EIP20Contract *EIP20ContractTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20Contract.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_EIP20Contract *EIP20ContractSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20Contract.Contract.Approve(&_EIP20Contract.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_EIP20Contract *EIP20ContractTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20Contract.Contract.Approve(&_EIP20Contract.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_EIP20Contract *EIP20ContractTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20Contract.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_EIP20Contract *EIP20ContractSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20Contract.Contract.Transfer(&_EIP20Contract.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_EIP20Contract *EIP20ContractTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20Contract.Contract.Transfer(&_EIP20Contract.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_EIP20Contract *EIP20ContractTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20Contract.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_EIP20Contract *EIP20ContractSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20Contract.Contract.TransferFrom(&_EIP20Contract.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_EIP20Contract *EIP20ContractTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20Contract.Contract.TransferFrom(&_EIP20Contract.TransactOpts, _from, _to, _value)
}

// EIP20ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EIP20Contract contract.
type EIP20ContractApprovalIterator struct {
	Event *EIP20ContractApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EIP20ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EIP20ContractApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EIP20ContractApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EIP20ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EIP20ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EIP20ContractApproval represents a Approval event raised by the EIP20Contract contract.
type EIP20ContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_EIP20Contract *EIP20ContractFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*EIP20ContractApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _EIP20Contract.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &EIP20ContractApprovalIterator{contract: _EIP20Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_EIP20Contract *EIP20ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EIP20ContractApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _EIP20Contract.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EIP20ContractApproval)
				if err := _EIP20Contract.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EIP20ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EIP20Contract contract.
type EIP20ContractTransferIterator struct {
	Event *EIP20ContractTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EIP20ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EIP20ContractTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EIP20ContractTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EIP20ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EIP20ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EIP20ContractTransfer represents a Transfer event raised by the EIP20Contract contract.
type EIP20ContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_EIP20Contract *EIP20ContractFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*EIP20ContractTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _EIP20Contract.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &EIP20ContractTransferIterator{contract: _EIP20Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_EIP20Contract *EIP20ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EIP20ContractTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _EIP20Contract.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EIP20ContractTransfer)
				if err := _EIP20Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
