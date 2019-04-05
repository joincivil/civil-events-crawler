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

// CVLTokenContractABI is the input ABI used to generate the binding from.
const CVLTokenContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialAmount\",\"type\":\"uint256\"},{\"name\":\"_tokenName\",\"type\":\"string\"},{\"name\":\"_decimalUnits\",\"type\":\"uint8\"},{\"name\":\"_tokenSymbol\",\"type\":\"string\"},{\"name\":\"_controller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"changeController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"detectTransferRestriction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"restrictionCode\",\"type\":\"uint8\"}],\"name\":\"messageForTransferRestriction\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CVLTokenContractBin is the compiled bytecode used for deploying new contracts.
const CVLTokenContractBin = `0x60806040523480156200001157600080fd5b50604051620011dc380380620011dc83398101604090815281516020808401519284015160608501516080860151948601805194969095929491019285918491869162000065916003919086019062000229565b5081516200007b90600490602085019062000229565b506005805460ff191660ff929092169190911761010060a860020a03191661010033021790555050600160a060020a03811615156200011b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f636f6e74726f6c6c6572206e6f742070726f7669646564000000000000000000604482015290519081900360640190fd5b60068054600160a060020a031916600160a060020a0383161790556200014b338664010000000062000156810204565b5050505050620002ce565b600160a060020a03821615156200016c57600080fd5b60025462000189908264010000000062000d416200021582021704565b600255600160a060020a038216600090815260208190526040902054620001bf908264010000000062000d416200021582021704565b600160a060020a0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b818101828110156200022357fe5b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200026c57805160ff19168380011785556200029c565b828001600101855582156200029c579182015b828111156200029c5782518255916020019190600101906200027f565b50620002aa929150620002ae565b5090565b620002cb91905b80821115620002aa5760008155600101620002b5565b90565b610efe80620002de6000396000f3006080604052600436106100fb5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610100578063095ea7b31461018a57806318160ddd146101c257806323b872dd146101e9578063313ce56714610213578063395093511461023e5780633cebb8231461026257806370a0823114610285578063715018a6146102a65780637f4ab1dd146102bb5780638da5cb5b146102d657806395d89b4114610307578063a457c2d71461031c578063a9059cbb14610340578063d4ce141514610364578063dd62ed3e1461038e578063f2fde38b146103b5578063f77c4791146103d6575b600080fd5b34801561010c57600080fd5b506101156103eb565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561014f578181015183820152602001610137565b50505050905090810190601f16801561017c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019657600080fd5b506101ae600160a060020a0360043516602435610481565b604080519115158252519081900360200190f35b3480156101ce57600080fd5b506101d76104ff565b60408051918252519081900360200190f35b3480156101f557600080fd5b506101ae600160a060020a0360043581169060243516604435610505565b34801561021f57600080fd5b5061022861061e565b6040805160ff9092168252519081900360200190f35b34801561024a57600080fd5b506101ae600160a060020a0360043516602435610627565b34801561026e57600080fd5b50610283600160a060020a03600435166106d7565b005b34801561029157600080fd5b506101d7600160a060020a03600435166107cd565b3480156102b257600080fd5b506102836107e8565b3480156102c757600080fd5b5061011560ff600435166108ab565b3480156102e257600080fd5b506102eb6109a1565b60408051600160a060020a039092168252519081900360200190f35b34801561031357600080fd5b506101156109b5565b34801561032857600080fd5b506101ae600160a060020a0360043516602435610a16565b34801561034c57600080fd5b506101ae600160a060020a0360043516602435610a61565b34801561037057600080fd5b50610228600160a060020a0360043581169060243516604435610b79565b34801561039a57600080fd5b506101d7600160a060020a0360043581169060243516610c27565b3480156103c157600080fd5b50610283600160a060020a0360043516610c52565b3480156103e257600080fd5b506102eb610cc5565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104775780601f1061044c57610100808354040283529160200191610477565b820191906000526020600020905b81548152906001019060200180831161045a57829003601f168201915b5050505050905090565b6000600160a060020a038316151561049857600080fd5b336000818152600160209081526040808320600160a060020a03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60025490565b600654604080517fd4ce1415000000000000000000000000000000000000000000000000000000008152600160a060020a0380871660048301528086166024830152604482018590529151600093879387938793919092169163d4ce141591606480830192602092919082900301818a87803b15801561058457600080fd5b505af1158015610598573d6000803e3d6000fd5b505050506040513d60208110156105ae57600080fd5b505160ff1615610608576040805160e560020a62461bcd02815260206004820152601960248201527f746f6b656e207472616e73666572207265737472696374656400000000000000604482015290519081900360640190fd5b610613878787610cd4565b979650505050505050565b60055460ff1690565b6000600160a060020a038316151561063e57600080fd5b336000908152600160209081526040808320600160a060020a0387168452909152902054610672908363ffffffff610d4116565b336000818152600160209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6005546101009004600160a060020a0316331461073e576040805160e560020a62461bcd02815260206004820152600960248201527f6e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a038116151561079e576040805160e560020a62461bcd02815260206004820152601760248201527f636f6e74726f6c6c6572206e6f742070726f7669646564000000000000000000604482015290519081900360640190fd5b6006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a031660009081526020819052604090205490565b6005546101009004600160a060020a0316331461084f576040805160e560020a62461bcd02815260206004820152600960248201527f6e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600554604051610100909104600160a060020a0316907ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482090600090a26005805474ffffffffffffffffffffffffffffffffffffffff0019169055565b600654604080517f7f4ab1dd00000000000000000000000000000000000000000000000000000000815260ff841660048201529051606092600160a060020a031691637f4ab1dd91602480830192600092919082900301818387803b15801561091357600080fd5b505af1158015610927573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561095057600080fd5b81019080805164010000000081111561096857600080fd5b8201602081018481111561097b57600080fd5b815164010000000081118282018710171561099557600080fd5b50909695505050505050565b6005546101009004600160a060020a031681565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104775780601f1061044c57610100808354040283529160200191610477565b6000600160a060020a0383161515610a2d57600080fd5b336000908152600160209081526040808320600160a060020a0387168452909152902054610672908363ffffffff610d5416565b600654604080517fd4ce14150000000000000000000000000000000000000000000000000000000081523360048201819052600160a060020a038087166024840152604483018690529251600094919387938793919091169163d4ce14159160648082019260209290919082900301818a87803b158015610ae157600080fd5b505af1158015610af5573d6000803e3d6000fd5b505050506040513d6020811015610b0b57600080fd5b505160ff1615610b65576040805160e560020a62461bcd02815260206004820152601960248201527f746f6b656e207472616e73666572207265737472696374656400000000000000604482015290519081900360640190fd5b610b6f8686610d66565b9695505050505050565b600654604080517fd4ce1415000000000000000000000000000000000000000000000000000000008152600160a060020a0386811660048301528581166024830152604482018590529151600093929092169163d4ce14159160648082019260209290919082900301818787803b158015610bf357600080fd5b505af1158015610c07573d6000803e3d6000fd5b505050506040513d6020811015610c1d57600080fd5b5051949350505050565b600160a060020a03918216600090815260016020908152604080832093909416825291909152205490565b6005546101009004600160a060020a03163314610cb9576040805160e560020a62461bcd02815260206004820152600960248201527f6e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610cc281610d7c565b50565b600654600160a060020a031681565b600160a060020a0383166000908152600160209081526040808320338452909152812054610d08908363ffffffff610d5416565b600160a060020a0385166000908152600160209081526040808320338452909152902055610d37848484610e05565b5060019392505050565b81810182811015610d4e57fe5b92915050565b600082821115610d6057fe5b50900390565b6000610d73338484610e05565b50600192915050565b600160a060020a0381161515610d9157600080fd5b600554604051600160a060020a0380841692610100900416907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a360058054600160a060020a039092166101000274ffffffffffffffffffffffffffffffffffffffff0019909216919091179055565b600160a060020a0382161515610e1a57600080fd5b600160a060020a038316600090815260208190526040902054610e43908263ffffffff610d5416565b600160a060020a038085166000908152602081905260408082209390935590841681522054610e78908263ffffffff610d4116565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35050505600a165627a7a7230582030a198df1f38432ea8f6adc56236b500ea7b354cb207803a028028c34d1c1c8a0029`

// DeployCVLTokenContract deploys a new Ethereum contract, binding an instance of CVLTokenContract to it.
func DeployCVLTokenContract(auth *bind.TransactOpts, backend bind.ContractBackend, _initialAmount *big.Int, _tokenName string, _decimalUnits uint8, _tokenSymbol string, _controller common.Address) (common.Address, *types.Transaction, *CVLTokenContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CVLTokenContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CVLTokenContractBin), backend, _initialAmount, _tokenName, _decimalUnits, _tokenSymbol, _controller)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CVLTokenContract{CVLTokenContractCaller: CVLTokenContractCaller{contract: contract}, CVLTokenContractTransactor: CVLTokenContractTransactor{contract: contract}, CVLTokenContractFilterer: CVLTokenContractFilterer{contract: contract}}, nil
}

// CVLTokenContract is an auto generated Go binding around an Ethereum contract.
type CVLTokenContract struct {
	CVLTokenContractCaller     // Read-only binding to the contract
	CVLTokenContractTransactor // Write-only binding to the contract
	CVLTokenContractFilterer   // Log filterer for contract events
}

// CVLTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CVLTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CVLTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CVLTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CVLTokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CVLTokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CVLTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CVLTokenContractSession struct {
	Contract     *CVLTokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CVLTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CVLTokenContractCallerSession struct {
	Contract *CVLTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CVLTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CVLTokenContractTransactorSession struct {
	Contract     *CVLTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CVLTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CVLTokenContractRaw struct {
	Contract *CVLTokenContract // Generic contract binding to access the raw methods on
}

// CVLTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CVLTokenContractCallerRaw struct {
	Contract *CVLTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// CVLTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CVLTokenContractTransactorRaw struct {
	Contract *CVLTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCVLTokenContract creates a new instance of CVLTokenContract, bound to a specific deployed contract.
func NewCVLTokenContract(address common.Address, backend bind.ContractBackend) (*CVLTokenContract, error) {
	contract, err := bindCVLTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CVLTokenContract{CVLTokenContractCaller: CVLTokenContractCaller{contract: contract}, CVLTokenContractTransactor: CVLTokenContractTransactor{contract: contract}, CVLTokenContractFilterer: CVLTokenContractFilterer{contract: contract}}, nil
}

// NewCVLTokenContractCaller creates a new read-only instance of CVLTokenContract, bound to a specific deployed contract.
func NewCVLTokenContractCaller(address common.Address, caller bind.ContractCaller) (*CVLTokenContractCaller, error) {
	contract, err := bindCVLTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CVLTokenContractCaller{contract: contract}, nil
}

// NewCVLTokenContractTransactor creates a new write-only instance of CVLTokenContract, bound to a specific deployed contract.
func NewCVLTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CVLTokenContractTransactor, error) {
	contract, err := bindCVLTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CVLTokenContractTransactor{contract: contract}, nil
}

// NewCVLTokenContractFilterer creates a new log filterer instance of CVLTokenContract, bound to a specific deployed contract.
func NewCVLTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CVLTokenContractFilterer, error) {
	contract, err := bindCVLTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CVLTokenContractFilterer{contract: contract}, nil
}

// bindCVLTokenContract binds a generic wrapper to an already deployed contract.
func bindCVLTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CVLTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CVLTokenContract *CVLTokenContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CVLTokenContract.Contract.CVLTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CVLTokenContract *CVLTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.CVLTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CVLTokenContract *CVLTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.CVLTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CVLTokenContract *CVLTokenContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CVLTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CVLTokenContract *CVLTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CVLTokenContract *CVLTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_CVLTokenContract *CVLTokenContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CVLTokenContract.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_CVLTokenContract *CVLTokenContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CVLTokenContract.Contract.Allowance(&_CVLTokenContract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_CVLTokenContract *CVLTokenContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CVLTokenContract.Contract.Allowance(&_CVLTokenContract.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_CVLTokenContract *CVLTokenContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CVLTokenContract.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_CVLTokenContract *CVLTokenContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CVLTokenContract.Contract.BalanceOf(&_CVLTokenContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_CVLTokenContract *CVLTokenContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CVLTokenContract.Contract.BalanceOf(&_CVLTokenContract.CallOpts, owner)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_CVLTokenContract *CVLTokenContractCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CVLTokenContract.contract.Call(opts, out, "controller")
	return *ret0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_CVLTokenContract *CVLTokenContractSession) Controller() (common.Address, error) {
	return _CVLTokenContract.Contract.Controller(&_CVLTokenContract.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_CVLTokenContract *CVLTokenContractCallerSession) Controller() (common.Address, error) {
	return _CVLTokenContract.Contract.Controller(&_CVLTokenContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_CVLTokenContract *CVLTokenContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CVLTokenContract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_CVLTokenContract *CVLTokenContractSession) Decimals() (uint8, error) {
	return _CVLTokenContract.Contract.Decimals(&_CVLTokenContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_CVLTokenContract *CVLTokenContractCallerSession) Decimals() (uint8, error) {
	return _CVLTokenContract.Contract.Decimals(&_CVLTokenContract.CallOpts)
}

// DetectTransferRestriction is a free data retrieval call binding the contract method 0xd4ce1415.
//
// Solidity: function detectTransferRestriction(from address, to address, value uint256) constant returns(uint8)
func (_CVLTokenContract *CVLTokenContractCaller) DetectTransferRestriction(opts *bind.CallOpts, from common.Address, to common.Address, value *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CVLTokenContract.contract.Call(opts, out, "detectTransferRestriction", from, to, value)
	return *ret0, err
}

// DetectTransferRestriction is a free data retrieval call binding the contract method 0xd4ce1415.
//
// Solidity: function detectTransferRestriction(from address, to address, value uint256) constant returns(uint8)
func (_CVLTokenContract *CVLTokenContractSession) DetectTransferRestriction(from common.Address, to common.Address, value *big.Int) (uint8, error) {
	return _CVLTokenContract.Contract.DetectTransferRestriction(&_CVLTokenContract.CallOpts, from, to, value)
}

// DetectTransferRestriction is a free data retrieval call binding the contract method 0xd4ce1415.
//
// Solidity: function detectTransferRestriction(from address, to address, value uint256) constant returns(uint8)
func (_CVLTokenContract *CVLTokenContractCallerSession) DetectTransferRestriction(from common.Address, to common.Address, value *big.Int) (uint8, error) {
	return _CVLTokenContract.Contract.DetectTransferRestriction(&_CVLTokenContract.CallOpts, from, to, value)
}

// MessageForTransferRestriction is a free data retrieval call binding the contract method 0x7f4ab1dd.
//
// Solidity: function messageForTransferRestriction(restrictionCode uint8) constant returns(string)
func (_CVLTokenContract *CVLTokenContractCaller) MessageForTransferRestriction(opts *bind.CallOpts, restrictionCode uint8) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CVLTokenContract.contract.Call(opts, out, "messageForTransferRestriction", restrictionCode)
	return *ret0, err
}

// MessageForTransferRestriction is a free data retrieval call binding the contract method 0x7f4ab1dd.
//
// Solidity: function messageForTransferRestriction(restrictionCode uint8) constant returns(string)
func (_CVLTokenContract *CVLTokenContractSession) MessageForTransferRestriction(restrictionCode uint8) (string, error) {
	return _CVLTokenContract.Contract.MessageForTransferRestriction(&_CVLTokenContract.CallOpts, restrictionCode)
}

// MessageForTransferRestriction is a free data retrieval call binding the contract method 0x7f4ab1dd.
//
// Solidity: function messageForTransferRestriction(restrictionCode uint8) constant returns(string)
func (_CVLTokenContract *CVLTokenContractCallerSession) MessageForTransferRestriction(restrictionCode uint8) (string, error) {
	return _CVLTokenContract.Contract.MessageForTransferRestriction(&_CVLTokenContract.CallOpts, restrictionCode)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CVLTokenContract *CVLTokenContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CVLTokenContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CVLTokenContract *CVLTokenContractSession) Name() (string, error) {
	return _CVLTokenContract.Contract.Name(&_CVLTokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CVLTokenContract *CVLTokenContractCallerSession) Name() (string, error) {
	return _CVLTokenContract.Contract.Name(&_CVLTokenContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CVLTokenContract *CVLTokenContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CVLTokenContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CVLTokenContract *CVLTokenContractSession) Owner() (common.Address, error) {
	return _CVLTokenContract.Contract.Owner(&_CVLTokenContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CVLTokenContract *CVLTokenContractCallerSession) Owner() (common.Address, error) {
	return _CVLTokenContract.Contract.Owner(&_CVLTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CVLTokenContract *CVLTokenContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CVLTokenContract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CVLTokenContract *CVLTokenContractSession) Symbol() (string, error) {
	return _CVLTokenContract.Contract.Symbol(&_CVLTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CVLTokenContract *CVLTokenContractCallerSession) Symbol() (string, error) {
	return _CVLTokenContract.Contract.Symbol(&_CVLTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CVLTokenContract *CVLTokenContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CVLTokenContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CVLTokenContract *CVLTokenContractSession) TotalSupply() (*big.Int, error) {
	return _CVLTokenContract.Contract.TotalSupply(&_CVLTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CVLTokenContract *CVLTokenContractCallerSession) TotalSupply() (*big.Int, error) {
	return _CVLTokenContract.Contract.TotalSupply(&_CVLTokenContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_CVLTokenContract *CVLTokenContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_CVLTokenContract *CVLTokenContractSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.Approve(&_CVLTokenContract.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_CVLTokenContract *CVLTokenContractTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.Approve(&_CVLTokenContract.TransactOpts, spender, value)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(_controller address) returns()
func (_CVLTokenContract *CVLTokenContractTransactor) ChangeController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _CVLTokenContract.contract.Transact(opts, "changeController", _controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(_controller address) returns()
func (_CVLTokenContract *CVLTokenContractSession) ChangeController(_controller common.Address) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.ChangeController(&_CVLTokenContract.TransactOpts, _controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(_controller address) returns()
func (_CVLTokenContract *CVLTokenContractTransactorSession) ChangeController(_controller common.Address) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.ChangeController(&_CVLTokenContract.TransactOpts, _controller)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_CVLTokenContract *CVLTokenContractTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_CVLTokenContract *CVLTokenContractSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.DecreaseAllowance(&_CVLTokenContract.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_CVLTokenContract *CVLTokenContractTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.DecreaseAllowance(&_CVLTokenContract.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_CVLTokenContract *CVLTokenContractTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_CVLTokenContract *CVLTokenContractSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.IncreaseAllowance(&_CVLTokenContract.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_CVLTokenContract *CVLTokenContractTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.IncreaseAllowance(&_CVLTokenContract.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CVLTokenContract *CVLTokenContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CVLTokenContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CVLTokenContract *CVLTokenContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _CVLTokenContract.Contract.RenounceOwnership(&_CVLTokenContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CVLTokenContract *CVLTokenContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CVLTokenContract.Contract.RenounceOwnership(&_CVLTokenContract.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(success bool)
func (_CVLTokenContract *CVLTokenContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(success bool)
func (_CVLTokenContract *CVLTokenContractSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.Transfer(&_CVLTokenContract.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(success bool)
func (_CVLTokenContract *CVLTokenContractTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.Transfer(&_CVLTokenContract.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(success bool)
func (_CVLTokenContract *CVLTokenContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(success bool)
func (_CVLTokenContract *CVLTokenContractSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.TransferFrom(&_CVLTokenContract.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(success bool)
func (_CVLTokenContract *CVLTokenContractTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.TransferFrom(&_CVLTokenContract.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_CVLTokenContract *CVLTokenContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _CVLTokenContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_CVLTokenContract *CVLTokenContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.TransferOwnership(&_CVLTokenContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_CVLTokenContract *CVLTokenContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _CVLTokenContract.Contract.TransferOwnership(&_CVLTokenContract.TransactOpts, _newOwner)
}

// CVLTokenContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CVLTokenContract contract.
type CVLTokenContractApprovalIterator struct {
	Event *CVLTokenContractApproval // Event containing the contract specifics and raw log

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
func (it *CVLTokenContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CVLTokenContractApproval)
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
		it.Event = new(CVLTokenContractApproval)
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
func (it *CVLTokenContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CVLTokenContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CVLTokenContractApproval represents a Approval event raised by the CVLTokenContract contract.
type CVLTokenContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_CVLTokenContract *CVLTokenContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CVLTokenContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CVLTokenContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CVLTokenContractApprovalIterator{contract: _CVLTokenContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_CVLTokenContract *CVLTokenContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CVLTokenContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CVLTokenContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CVLTokenContractApproval)
				if err := _CVLTokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// CVLTokenContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the CVLTokenContract contract.
type CVLTokenContractOwnershipRenouncedIterator struct {
	Event *CVLTokenContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *CVLTokenContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CVLTokenContractOwnershipRenounced)
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
		it.Event = new(CVLTokenContractOwnershipRenounced)
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
func (it *CVLTokenContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CVLTokenContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CVLTokenContractOwnershipRenounced represents a OwnershipRenounced event raised by the CVLTokenContract contract.
type CVLTokenContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_CVLTokenContract *CVLTokenContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*CVLTokenContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _CVLTokenContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CVLTokenContractOwnershipRenouncedIterator{contract: _CVLTokenContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_CVLTokenContract *CVLTokenContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *CVLTokenContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _CVLTokenContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CVLTokenContractOwnershipRenounced)
				if err := _CVLTokenContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// CVLTokenContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CVLTokenContract contract.
type CVLTokenContractOwnershipTransferredIterator struct {
	Event *CVLTokenContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CVLTokenContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CVLTokenContractOwnershipTransferred)
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
		it.Event = new(CVLTokenContractOwnershipTransferred)
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
func (it *CVLTokenContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CVLTokenContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CVLTokenContractOwnershipTransferred represents a OwnershipTransferred event raised by the CVLTokenContract contract.
type CVLTokenContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CVLTokenContract *CVLTokenContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CVLTokenContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CVLTokenContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CVLTokenContractOwnershipTransferredIterator{contract: _CVLTokenContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CVLTokenContract *CVLTokenContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CVLTokenContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CVLTokenContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CVLTokenContractOwnershipTransferred)
				if err := _CVLTokenContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// CVLTokenContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CVLTokenContract contract.
type CVLTokenContractTransferIterator struct {
	Event *CVLTokenContractTransfer // Event containing the contract specifics and raw log

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
func (it *CVLTokenContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CVLTokenContractTransfer)
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
		it.Event = new(CVLTokenContractTransfer)
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
func (it *CVLTokenContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CVLTokenContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CVLTokenContractTransfer represents a Transfer event raised by the CVLTokenContract contract.
type CVLTokenContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_CVLTokenContract *CVLTokenContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CVLTokenContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CVLTokenContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CVLTokenContractTransferIterator{contract: _CVLTokenContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_CVLTokenContract *CVLTokenContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CVLTokenContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CVLTokenContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CVLTokenContractTransfer)
				if err := _CVLTokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
