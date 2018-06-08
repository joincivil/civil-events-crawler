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

// GovernmentContractABI is the input ABI used to generate the binding from.
const GovernmentContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"constitutionURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"constitutionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appellate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"params\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governmentController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"appellateAddr\",\"type\":\"address\"},{\"name\":\"governmentControllerAddr\",\"type\":\"address\"},{\"name\":\"appealFeeAmount\",\"type\":\"uint256\"},{\"name\":\"requestAppealLength\",\"type\":\"uint256\"},{\"name\":\"judgeAppealLength\",\"type\":\"uint256\"},{\"name\":\"appealSupermajorityPercentage\",\"type\":\"uint256\"},{\"name\":\"constHash\",\"type\":\"bytes32\"},{\"name\":\"constURI\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAppellate\",\"type\":\"address\"}],\"name\":\"AppellateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ParameterSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAppellate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGovernmentController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAppellate\",\"type\":\"address\"}],\"name\":\"setAppellate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GovernmentContractBin is the compiled bytecode used for deploying new contracts.
const GovernmentContractBin = `0x608060405234801561001057600080fd5b50604051610d21380380610d218339810180604052810190808051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080518201929190505050876000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555086600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101486040805190810160405280601081526020017f7265717565737441707065616c4c656e0000000000000000000000000000000081525086610261640100000000026401000000009004565b6101966040805190810160405280600e81526020017f6a7564676541707065616c4c656e00000000000000000000000000000000000081525085610261640100000000026401000000009004565b6101e46040805190810160405280600981526020017f61707065616c466565000000000000000000000000000000000000000000000081525087610261640100000000026401000000009004565b6102326040805190810160405280601481526020017f61707065616c566f746550657263656e7461676500000000000000000000000081525084610261640100000000026401000000009004565b8160028160001916905550806003908051906020019061025392919061038b565b505050505050505050610430565b8060046000846040518082805190602001908083835b60208310151561029c5780518252602082019150602081019050602083039250610277565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916600019168152602001908152602001600020819055507fc6eeddce4b4af1253d6c284283a236f83c854ba2163447903920eeec1842903b82826040518080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561034c578082015181840152602081019050610331565b50505050905090810190601f1680156103795780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106103cc57805160ff19168380011785556103fa565b828001600101855582156103fa579182015b828111156103f95782518255916020019190600101906103de565b5b509050610407919061040b565b5090565b61042d91905b80821115610429576000816000905550600101610411565b5090565b90565b6108e28061043f6000396000f3006080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806355122425146100a957806356e1fb88146100ec5780635793b9cf14610143578063693ec85e1461019a5780638a42ebe9146102175780638ca7f51c1461028a578063c7d93fd41461031a578063d5fd9e661461034d578063dc6ab527146103a4578063f2a2129b146103e9575b600080fd5b3480156100b557600080fd5b506100ea600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610440565b005b3480156100f857600080fd5b50610101610542565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561014f57600080fd5b5061015861056b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101a657600080fd5b50610201600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610595565b6040518082815260200191505060405180910390f35b34801561022357600080fd5b50610288600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192908035906020019092919050505061061c565b005b34801561029657600080fd5b5061029f610685565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102df5780820151818401526020810190506102c4565b50505050905090810190601f16801561030c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561032657600080fd5b5061032f610723565b60405180826000191660001916815260200191505060405180910390f35b34801561035957600080fd5b50610362610729565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103b057600080fd5b506103d3600480360381019080803560001916906020019092919050505061074e565b6040518082815260200191505060405180910390f35b3480156103f557600080fd5b506103fe610766565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561049c57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fded0dafe9941e82c316d4ed230b29b07f3e2f3a3064cb0c71dab14b5bcef7b2581604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600060046000836040518082805190602001908083835b6020831015156105d157805182526020820191506020810190506020830392506105ac565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916600019168152602001908152602001600020549050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561067757600080fd5b610681828261078c565b5050565b60038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561071b5780601f106106f05761010080835404028352916020019161071b565b820191906000526020600020905b8154815290600101906020018083116106fe57829003601f168201915b505050505081565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8060046000846040518082805190602001908083835b6020831015156107c757805182526020820191506020810190506020830392506107a2565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916600019168152602001908152602001600020819055507fc6eeddce4b4af1253d6c284283a236f83c854ba2163447903920eeec1842903b82826040518080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561087757808201518184015260208101905061085c565b50505050905090810190601f1680156108a45780820380516001836020036101000a031916815260200191505b50935050505060405180910390a150505600a165627a7a72305820933f344b0a12f9193a02f57cb57caa7100ec280469be277d34e19102065125b30029`

// DeployGovernmentContract deploys a new Ethereum contract, binding an instance of GovernmentContract to it.
func DeployGovernmentContract(auth *bind.TransactOpts, backend bind.ContractBackend, appellateAddr common.Address, governmentControllerAddr common.Address, appealFeeAmount *big.Int, requestAppealLength *big.Int, judgeAppealLength *big.Int, appealSupermajorityPercentage *big.Int, constHash [32]byte, constURI string) (common.Address, *types.Transaction, *GovernmentContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernmentContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernmentContractBin), backend, appellateAddr, governmentControllerAddr, appealFeeAmount, requestAppealLength, judgeAppealLength, appealSupermajorityPercentage, constHash, constURI)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernmentContract{GovernmentContractCaller: GovernmentContractCaller{contract: contract}, GovernmentContractTransactor: GovernmentContractTransactor{contract: contract}, GovernmentContractFilterer: GovernmentContractFilterer{contract: contract}}, nil
}

// GovernmentContract is an auto generated Go binding around an Ethereum contract.
type GovernmentContract struct {
	GovernmentContractCaller     // Read-only binding to the contract
	GovernmentContractTransactor // Write-only binding to the contract
	GovernmentContractFilterer   // Log filterer for contract events
}

// GovernmentContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernmentContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernmentContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernmentContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernmentContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernmentContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernmentContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernmentContractSession struct {
	Contract     *GovernmentContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GovernmentContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernmentContractCallerSession struct {
	Contract *GovernmentContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GovernmentContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernmentContractTransactorSession struct {
	Contract     *GovernmentContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GovernmentContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernmentContractRaw struct {
	Contract *GovernmentContract // Generic contract binding to access the raw methods on
}

// GovernmentContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernmentContractCallerRaw struct {
	Contract *GovernmentContractCaller // Generic read-only contract binding to access the raw methods on
}

// GovernmentContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernmentContractTransactorRaw struct {
	Contract *GovernmentContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernmentContract creates a new instance of GovernmentContract, bound to a specific deployed contract.
func NewGovernmentContract(address common.Address, backend bind.ContractBackend) (*GovernmentContract, error) {
	contract, err := bindGovernmentContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernmentContract{GovernmentContractCaller: GovernmentContractCaller{contract: contract}, GovernmentContractTransactor: GovernmentContractTransactor{contract: contract}, GovernmentContractFilterer: GovernmentContractFilterer{contract: contract}}, nil
}

// NewGovernmentContractCaller creates a new read-only instance of GovernmentContract, bound to a specific deployed contract.
func NewGovernmentContractCaller(address common.Address, caller bind.ContractCaller) (*GovernmentContractCaller, error) {
	contract, err := bindGovernmentContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernmentContractCaller{contract: contract}, nil
}

// NewGovernmentContractTransactor creates a new write-only instance of GovernmentContract, bound to a specific deployed contract.
func NewGovernmentContractTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernmentContractTransactor, error) {
	contract, err := bindGovernmentContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernmentContractTransactor{contract: contract}, nil
}

// NewGovernmentContractFilterer creates a new log filterer instance of GovernmentContract, bound to a specific deployed contract.
func NewGovernmentContractFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernmentContractFilterer, error) {
	contract, err := bindGovernmentContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernmentContractFilterer{contract: contract}, nil
}

// bindGovernmentContract binds a generic wrapper to an already deployed contract.
func bindGovernmentContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernmentContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernmentContract *GovernmentContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernmentContract.Contract.GovernmentContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernmentContract *GovernmentContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernmentContract.Contract.GovernmentContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernmentContract *GovernmentContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernmentContract.Contract.GovernmentContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernmentContract *GovernmentContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernmentContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernmentContract *GovernmentContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernmentContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernmentContract *GovernmentContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernmentContract.Contract.contract.Transact(opts, method, params...)
}

// Appellate is a free data retrieval call binding the contract method 0xd5fd9e66.
//
// Solidity: function appellate() constant returns(address)
func (_GovernmentContract *GovernmentContractCaller) Appellate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "appellate")
	return *ret0, err
}

// Appellate is a free data retrieval call binding the contract method 0xd5fd9e66.
//
// Solidity: function appellate() constant returns(address)
func (_GovernmentContract *GovernmentContractSession) Appellate() (common.Address, error) {
	return _GovernmentContract.Contract.Appellate(&_GovernmentContract.CallOpts)
}

// Appellate is a free data retrieval call binding the contract method 0xd5fd9e66.
//
// Solidity: function appellate() constant returns(address)
func (_GovernmentContract *GovernmentContractCallerSession) Appellate() (common.Address, error) {
	return _GovernmentContract.Contract.Appellate(&_GovernmentContract.CallOpts)
}

// ConstitutionHash is a free data retrieval call binding the contract method 0xc7d93fd4.
//
// Solidity: function constitutionHash() constant returns(bytes32)
func (_GovernmentContract *GovernmentContractCaller) ConstitutionHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "constitutionHash")
	return *ret0, err
}

// ConstitutionHash is a free data retrieval call binding the contract method 0xc7d93fd4.
//
// Solidity: function constitutionHash() constant returns(bytes32)
func (_GovernmentContract *GovernmentContractSession) ConstitutionHash() ([32]byte, error) {
	return _GovernmentContract.Contract.ConstitutionHash(&_GovernmentContract.CallOpts)
}

// ConstitutionHash is a free data retrieval call binding the contract method 0xc7d93fd4.
//
// Solidity: function constitutionHash() constant returns(bytes32)
func (_GovernmentContract *GovernmentContractCallerSession) ConstitutionHash() ([32]byte, error) {
	return _GovernmentContract.Contract.ConstitutionHash(&_GovernmentContract.CallOpts)
}

// ConstitutionURI is a free data retrieval call binding the contract method 0x8ca7f51c.
//
// Solidity: function constitutionURI() constant returns(string)
func (_GovernmentContract *GovernmentContractCaller) ConstitutionURI(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "constitutionURI")
	return *ret0, err
}

// ConstitutionURI is a free data retrieval call binding the contract method 0x8ca7f51c.
//
// Solidity: function constitutionURI() constant returns(string)
func (_GovernmentContract *GovernmentContractSession) ConstitutionURI() (string, error) {
	return _GovernmentContract.Contract.ConstitutionURI(&_GovernmentContract.CallOpts)
}

// ConstitutionURI is a free data retrieval call binding the contract method 0x8ca7f51c.
//
// Solidity: function constitutionURI() constant returns(string)
func (_GovernmentContract *GovernmentContractCallerSession) ConstitutionURI() (string, error) {
	return _GovernmentContract.Contract.ConstitutionURI(&_GovernmentContract.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(value uint256)
func (_GovernmentContract *GovernmentContractCaller) Get(opts *bind.CallOpts, name string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "get", name)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(value uint256)
func (_GovernmentContract *GovernmentContractSession) Get(name string) (*big.Int, error) {
	return _GovernmentContract.Contract.Get(&_GovernmentContract.CallOpts, name)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(value uint256)
func (_GovernmentContract *GovernmentContractCallerSession) Get(name string) (*big.Int, error) {
	return _GovernmentContract.Contract.Get(&_GovernmentContract.CallOpts, name)
}

// GetAppellate is a free data retrieval call binding the contract method 0x56e1fb88.
//
// Solidity: function getAppellate() constant returns(address)
func (_GovernmentContract *GovernmentContractCaller) GetAppellate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "getAppellate")
	return *ret0, err
}

// GetAppellate is a free data retrieval call binding the contract method 0x56e1fb88.
//
// Solidity: function getAppellate() constant returns(address)
func (_GovernmentContract *GovernmentContractSession) GetAppellate() (common.Address, error) {
	return _GovernmentContract.Contract.GetAppellate(&_GovernmentContract.CallOpts)
}

// GetAppellate is a free data retrieval call binding the contract method 0x56e1fb88.
//
// Solidity: function getAppellate() constant returns(address)
func (_GovernmentContract *GovernmentContractCallerSession) GetAppellate() (common.Address, error) {
	return _GovernmentContract.Contract.GetAppellate(&_GovernmentContract.CallOpts)
}

// GetGovernmentController is a free data retrieval call binding the contract method 0x5793b9cf.
//
// Solidity: function getGovernmentController() constant returns(address)
func (_GovernmentContract *GovernmentContractCaller) GetGovernmentController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "getGovernmentController")
	return *ret0, err
}

// GetGovernmentController is a free data retrieval call binding the contract method 0x5793b9cf.
//
// Solidity: function getGovernmentController() constant returns(address)
func (_GovernmentContract *GovernmentContractSession) GetGovernmentController() (common.Address, error) {
	return _GovernmentContract.Contract.GetGovernmentController(&_GovernmentContract.CallOpts)
}

// GetGovernmentController is a free data retrieval call binding the contract method 0x5793b9cf.
//
// Solidity: function getGovernmentController() constant returns(address)
func (_GovernmentContract *GovernmentContractCallerSession) GetGovernmentController() (common.Address, error) {
	return _GovernmentContract.Contract.GetGovernmentController(&_GovernmentContract.CallOpts)
}

// GovernmentController is a free data retrieval call binding the contract method 0xf2a2129b.
//
// Solidity: function governmentController() constant returns(address)
func (_GovernmentContract *GovernmentContractCaller) GovernmentController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "governmentController")
	return *ret0, err
}

// GovernmentController is a free data retrieval call binding the contract method 0xf2a2129b.
//
// Solidity: function governmentController() constant returns(address)
func (_GovernmentContract *GovernmentContractSession) GovernmentController() (common.Address, error) {
	return _GovernmentContract.Contract.GovernmentController(&_GovernmentContract.CallOpts)
}

// GovernmentController is a free data retrieval call binding the contract method 0xf2a2129b.
//
// Solidity: function governmentController() constant returns(address)
func (_GovernmentContract *GovernmentContractCallerSession) GovernmentController() (common.Address, error) {
	return _GovernmentContract.Contract.GovernmentController(&_GovernmentContract.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xdc6ab527.
//
// Solidity: function params( bytes32) constant returns(uint256)
func (_GovernmentContract *GovernmentContractCaller) Params(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "params", arg0)
	return *ret0, err
}

// Params is a free data retrieval call binding the contract method 0xdc6ab527.
//
// Solidity: function params( bytes32) constant returns(uint256)
func (_GovernmentContract *GovernmentContractSession) Params(arg0 [32]byte) (*big.Int, error) {
	return _GovernmentContract.Contract.Params(&_GovernmentContract.CallOpts, arg0)
}

// Params is a free data retrieval call binding the contract method 0xdc6ab527.
//
// Solidity: function params( bytes32) constant returns(uint256)
func (_GovernmentContract *GovernmentContractCallerSession) Params(arg0 [32]byte) (*big.Int, error) {
	return _GovernmentContract.Contract.Params(&_GovernmentContract.CallOpts, arg0)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(name string, value uint256) returns()
func (_GovernmentContract *GovernmentContractTransactor) Set(opts *bind.TransactOpts, name string, value *big.Int) (*types.Transaction, error) {
	return _GovernmentContract.contract.Transact(opts, "set", name, value)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(name string, value uint256) returns()
func (_GovernmentContract *GovernmentContractSession) Set(name string, value *big.Int) (*types.Transaction, error) {
	return _GovernmentContract.Contract.Set(&_GovernmentContract.TransactOpts, name, value)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(name string, value uint256) returns()
func (_GovernmentContract *GovernmentContractTransactorSession) Set(name string, value *big.Int) (*types.Transaction, error) {
	return _GovernmentContract.Contract.Set(&_GovernmentContract.TransactOpts, name, value)
}

// SetAppellate is a paid mutator transaction binding the contract method 0x55122425.
//
// Solidity: function setAppellate(newAppellate address) returns()
func (_GovernmentContract *GovernmentContractTransactor) SetAppellate(opts *bind.TransactOpts, newAppellate common.Address) (*types.Transaction, error) {
	return _GovernmentContract.contract.Transact(opts, "setAppellate", newAppellate)
}

// SetAppellate is a paid mutator transaction binding the contract method 0x55122425.
//
// Solidity: function setAppellate(newAppellate address) returns()
func (_GovernmentContract *GovernmentContractSession) SetAppellate(newAppellate common.Address) (*types.Transaction, error) {
	return _GovernmentContract.Contract.SetAppellate(&_GovernmentContract.TransactOpts, newAppellate)
}

// SetAppellate is a paid mutator transaction binding the contract method 0x55122425.
//
// Solidity: function setAppellate(newAppellate address) returns()
func (_GovernmentContract *GovernmentContractTransactorSession) SetAppellate(newAppellate common.Address) (*types.Transaction, error) {
	return _GovernmentContract.Contract.SetAppellate(&_GovernmentContract.TransactOpts, newAppellate)
}

// GovernmentContractAppellateSetIterator is returned from FilterAppellateSet and is used to iterate over the raw logs and unpacked data for AppellateSet events raised by the GovernmentContract contract.
type GovernmentContractAppellateSetIterator struct {
	Event *GovernmentContractAppellateSet // Event containing the contract specifics and raw log

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
func (it *GovernmentContractAppellateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernmentContractAppellateSet)
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
		it.Event = new(GovernmentContractAppellateSet)
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
func (it *GovernmentContractAppellateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernmentContractAppellateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernmentContractAppellateSet represents a AppellateSet event raised by the GovernmentContract contract.
type GovernmentContractAppellateSet struct {
	NewAppellate common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAppellateSet is a free log retrieval operation binding the contract event 0xded0dafe9941e82c316d4ed230b29b07f3e2f3a3064cb0c71dab14b5bcef7b25.
//
// Solidity: event AppellateSet(newAppellate address)
func (_GovernmentContract *GovernmentContractFilterer) FilterAppellateSet(opts *bind.FilterOpts) (*GovernmentContractAppellateSetIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "AppellateSet")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractAppellateSetIterator{contract: _GovernmentContract.contract, event: "AppellateSet", logs: logs, sub: sub}, nil
}

// WatchAppellateSet is a free log subscription operation binding the contract event 0xded0dafe9941e82c316d4ed230b29b07f3e2f3a3064cb0c71dab14b5bcef7b25.
//
// Solidity: event AppellateSet(newAppellate address)
func (_GovernmentContract *GovernmentContractFilterer) WatchAppellateSet(opts *bind.WatchOpts, sink chan<- *GovernmentContractAppellateSet) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "AppellateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernmentContractAppellateSet)
				if err := _GovernmentContract.contract.UnpackLog(event, "AppellateSet", log); err != nil {
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

// GovernmentContractParameterSetIterator is returned from FilterParameterSet and is used to iterate over the raw logs and unpacked data for ParameterSet events raised by the GovernmentContract contract.
type GovernmentContractParameterSetIterator struct {
	Event *GovernmentContractParameterSet // Event containing the contract specifics and raw log

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
func (it *GovernmentContractParameterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernmentContractParameterSet)
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
		it.Event = new(GovernmentContractParameterSet)
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
func (it *GovernmentContractParameterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernmentContractParameterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernmentContractParameterSet represents a ParameterSet event raised by the GovernmentContract contract.
type GovernmentContractParameterSet struct {
	Name  string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParameterSet is a free log retrieval operation binding the contract event 0xc6eeddce4b4af1253d6c284283a236f83c854ba2163447903920eeec1842903b.
//
// Solidity: event ParameterSet(name string, value uint256)
func (_GovernmentContract *GovernmentContractFilterer) FilterParameterSet(opts *bind.FilterOpts) (*GovernmentContractParameterSetIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "ParameterSet")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractParameterSetIterator{contract: _GovernmentContract.contract, event: "ParameterSet", logs: logs, sub: sub}, nil
}

// WatchParameterSet is a free log subscription operation binding the contract event 0xc6eeddce4b4af1253d6c284283a236f83c854ba2163447903920eeec1842903b.
//
// Solidity: event ParameterSet(name string, value uint256)
func (_GovernmentContract *GovernmentContractFilterer) WatchParameterSet(opts *bind.WatchOpts, sink chan<- *GovernmentContractParameterSet) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "ParameterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernmentContractParameterSet)
				if err := _GovernmentContract.contract.UnpackLog(event, "ParameterSet", log); err != nil {
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
