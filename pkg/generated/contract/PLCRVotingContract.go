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

// PLCRVotingContractABI is the input ABI used to generate the binding from.
const PLCRVotingContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_POLL_NONCE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"voteTokenBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pollMap\",\"outputs\":[{\"name\":\"commitEndDate\",\"type\":\"uint256\"},{\"name\":\"revealEndDate\",\"type\":\"uint256\"},{\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"name\":\"votesFor\",\"type\":\"uint256\"},{\"name\":\"votesAgainst\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pollNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_VoteCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesFor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesAgainst\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"choice\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_VoteRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commitEndDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealEndDate\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"_PollCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_VotingRightsGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_VotingRightsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_TokensRescued\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_numTokens\",\"type\":\"uint256\"}],\"name\":\"requestVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_numTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"},{\"name\":\"_numTokens\",\"type\":\"uint256\"},{\"name\":\"_prevPollID\",\"type\":\"uint256\"}],\"name\":\"commitVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prevID\",\"type\":\"uint256\"},{\"name\":\"_nextID\",\"type\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_numTokens\",\"type\":\"uint256\"}],\"name\":\"validPosition\",\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"getNumPassingTokens\",\"outputs\":[{\"name\":\"correctVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"getNumLosingTokens\",\"outputs\":[{\"name\":\"correctVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteQuorum\",\"type\":\"uint256\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"}],\"name\":\"startPoll\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"isPassed\",\"outputs\":[{\"name\":\"passed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getTotalNumberOfTokensForWinningOption\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getTotalNumberOfTokensForLosingOption\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"pollEnded\",\"outputs\":[{\"name\":\"ended\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"commitPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"revealPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"didCommit\",\"outputs\":[{\"name\":\"committed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"didReveal\",\"outputs\":[{\"name\":\"revealed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"pollExists\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getCommitHash\",\"outputs\":[{\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getNumTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getLastNode\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getLockedTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_numTokens\",\"type\":\"uint256\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getInsertPointForNumTokens\",\"outputs\":[{\"name\":\"prevNode\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_terminationDate\",\"type\":\"uint256\"}],\"name\":\"isExpired\",\"outputs\":[{\"name\":\"expired\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"attrUUID\",\"outputs\":[{\"name\":\"UUID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PLCRVotingContractBin is the compiled bytecode used for deploying new contracts.
const PLCRVotingContractBin = `0x608060405234801561001057600080fd5b50604051602080612a298339810180604052810190808051906020019092919050505080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600080819055505061299e8061008b6000396000f300608060405260043610610180576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063053e71a6146101855780632173a10f146101c65780632c052031146101f157806332ed3d601461025c5780633b930294146102b1578063427fa1d214610308578063441c77c01461035f57806349403183146103a45780636148fed5146103e95780636afa97a8146104465780636b2d95d4146104b15780636cbf9c5e146105085780637f97e83614610557578063819b0293146105bc57806388d21ff31461063557806397508f361461067a57806397603560146106a5578063a1103f37146106d2578063a25236fe1461073b578063a4439dc514610768578063aa7ca464146107ad578063b11d8bb814610812578063b43bd06914610853578063d1382092146108be578063d901402b1461091f578063d9548e5314610988578063e7b1d43c146109cd578063e8cfa3f0146109fa578063ee68483014610a3b578063fc0c546a14610a80575b600080fd5b34801561019157600080fd5b506101b060048036038101908080359060200190929190505050610ad7565b6040518082815260200191505060405180910390f35b3480156101d257600080fd5b506101db610b38565b6040518082815260200191505060405180910390f35b3480156101fd57600080fd5b50610246600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050610b3d565b6040518082815260200191505060405180910390f35b34801561026857600080fd5b5061029b600480360381019080803590602001909291908035906020019092919080359060200190929190505050610d7c565b6040518082815260200191505060405180910390f35b3480156102bd57600080fd5b506102f2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e99565b6040518082815260200191505060405180910390f35b34801561031457600080fd5b50610349600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610eb1565b6040518082815260200191505060405180910390f35b34801561036b57600080fd5b5061038a60048036038101908080359060200190929190505050610fa9565b604051808215151515815260200191505060405180910390f35b3480156103b057600080fd5b506103cf60048036038101908080359060200190929190505050610ff8565b604051808215151515815260200191505060405180910390f35b3480156103f557600080fd5b506104146004803603810190808035906020019092919050505061108b565b604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390f35b34801561045257600080fd5b5061049b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506110c1565b6040518082815260200191505060405180910390f35b3480156104bd57600080fd5b506104f2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111bb565b6040518082815260200191505060405180910390f35b34801561051457600080fd5b5061055560048036038101908080359060200190929190803560001916906020019092919080359060200190929190803590602001909291905050506111d6565b005b34801561056357600080fd5b506105a2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506118af565b604051808215151515815260200191505060405180910390f35b3480156105c857600080fd5b5061061b6004803603810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061192e565b604051808215151515815260200191505060405180910390f35b34801561064157600080fd5b5061066060048036038101908080359060200190929190505050611972565b604051808215151515815260200191505060405180910390f35b34801561068657600080fd5b5061068f61198d565b6040518082815260200191505060405180910390f35b3480156106b157600080fd5b506106d060048036038101908080359060200190929190505050611993565b005b3480156106de57600080fd5b5061071d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611bc8565b60405180826000191660001916815260200191505060405180910390f35b34801561074757600080fd5b5061076660048036038101908080359060200190929190505050611c2b565b005b34801561077457600080fd5b5061079360048036038101908080359060200190929190505050611f0e565b604051808215151515815260200191505060405180910390f35b3480156107b957600080fd5b506107f8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611f4b565b604051808215151515815260200191505060405180910390f35b34801561081e57600080fd5b50610851600480360381019080803590602001909291908035906020019092919080359060200190929190505050611fca565b005b34801561085f57600080fd5b506108a8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050612320565b6040518082815260200191505060405180910390f35b3480156108ca57600080fd5b50610909600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061241a565b6040518082815260200191505060405180910390f35b34801561092b57600080fd5b5061096a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061251f565b60405180826000191660001916815260200191505060405180910390f35b34801561099457600080fd5b506109b360048036038101908080359060200190929190505050612627565b604051808215151515815260200191505060405180910390f35b3480156109d957600080fd5b506109f860048036038101908080359060200190929190505050612633565b005b348015610a0657600080fd5b50610a256004803603810190808035906020019092919050505061284a565b6040518082815260200191505060405180910390f35b348015610a4757600080fd5b50610a66600480360381019080803590602001909291905050506128ab565b604051808215151515815260200191505060405180910390f35b348015610a8c57600080fd5b50610a956128e7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000610ae2826128ab565b1515610aed57600080fd5b610af682610ff8565b15610b195760016000838152602001908152602001600020600301549050610b33565b600160008381526020019081526020016000206004015490505b919050565b600081565b6000806000610b4b86610eb1565b9150610b57868361241a565b90505b600082141515610d6f57610b6e868361241a565b90508481111515610c7a5783821415610c7257600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002073__DLL___________________________________6330fe0a0a9091846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060206040518083038186803b158015610c3457600080fd5b505af4158015610c48573d6000803e3d6000fd5b505050506040513d6020811015610c5e57600080fd5b810190808051906020019092919050505091505b819250610d73565b600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002073__DLL___________________________________6330fe0a0a9091846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060206040518083038186803b158015610d2d57600080fd5b505af4158015610d41573d6000803e3d6000fd5b505050506040513d6020811015610d5757600080fd5b81019080805190602001909291905050509150610b5a565b8192505b50509392505050565b6000806000600160005401600081905550610da0854261290d90919063ffffffff16565b9150610db5848361290d90919063ffffffff16565b905060a060405190810160405280838152602001828152602001878152602001600081526020016000815250600160008054815260200190815260200160002060008201518160000155602082015181600101556040820151816002015560608201518160030155608082015181600401559050503373ffffffffffffffffffffffffffffffffffffffff166000547f404f1f1c229d9eb2a949e7584da6ffde9d059ef2169f487ca815434cce0640d088858560405180848152602001838152602001828152602001935050505060405180910390a3600054925050509392505050565b60026020528060005260406000206000915090505481565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002073__DLL___________________________________6330fe0a0a909160006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060206040518083038186803b158015610f6757600080fd5b505af4158015610f7b573d6000803e3d6000fd5b505050506040513d6020811015610f9157600080fd5b81019080805190602001909291905050509050919050565b6000610fb482611972565b1515610fbf57600080fd5b610fde6001600084815260200190815260200160002060010154612627565b158015610ff15750610fef82611f0e565b155b9050919050565b6000611002612942565b61100b836128ab565b151561101657600080fd5b6001600084815260200190815260200160002060a060405190810160405290816000820154815260200160018201548152602001600282015481526020016003820154815260200160048201548152505090508060800151816060015101816040015102816060015160640211915050919050565b60016020528060005260406000206000915090508060000154908060010154908060020154908060030154908060040154905085565b6000806000806110d0866128ab565b15156110db57600080fd5b6001600087815260200190815260200160002060060160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561114757600080fd5b61115086610ff8565b61115b57600161115e565b60005b60ff169250828560405180838152602001828152602001925050506040518091039020915061118d878761251f565b9050806000191682600019161415156111a557600080fd5b6111af878761241a565b93505050509392505050565b60006111cf826111ca84610eb1565b61241a565b9050919050565b6000806111e286611f0e565b15156111ed57600080fd5b83600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015151561123b57600080fd5b6000861415151561124b57600080fd5b60008314806113445750600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002073__DLL___________________________________63366a5ba29091856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060206040518083038186803b15801561130857600080fd5b505af415801561131c573d6000803e3d6000fd5b505050506040513d602081101561133257600080fd5b81019080805190602001909291905050505b151561134f57600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002073__DLL___________________________________6307d29ac99091856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060206040518083038186803b15801561140257600080fd5b505af4158015611416573d6000803e3d6000fd5b505050506040513d602081101561142c57600080fd5b8101908080519060200190929190505050915085821461144c578161153b565b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002073__DLL___________________________________6307d29ac99091886040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060206040518083038186803b1580156114ff57600080fd5b505af4158015611513573d6000803e3d6000fd5b505050506040513d602081101561152957600080fd5b81019080805190602001909291905050505b91506115498383338761192e565b151561155457600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002073__DLL___________________________________639735c51b90918589866040518563ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018085815260200184815260200183815260200182815260200194505050505060006040518083038186803b15801561161757600080fd5b505af415801561162b573d6000803e3d6000fd5b505050506116393387611bc8565b9050600473__AttributeStore________________________63977aa031909183876040518463ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180848152602001836000191660001916815260200180602001838152602001828103825260098152602001807f6e756d546f6b656e73000000000000000000000000000000000000000000000081525060200194505050505060006040518083038186803b1580156116fa57600080fd5b505af415801561170e573d6000803e3d6000fd5b50505050600473__AttributeStore________________________63977aa03190918388600190046040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808481526020018360001916600019168152602001806020018381526020018281038252600a8152602001807f636f6d6d6974486173680000000000000000000000000000000000000000000081525060200194505050505060006040518083038186803b1580156117d557600080fd5b505af41580156117e9573d6000803e3d6000fd5b50505050600180600088815260200190815260200160002060050160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff16867fea7979e4280d7e6bffc1c7d83a1ac99f16d02ecc14465ce41016226783b663d7866040518082815260200191505060405180910390a3505050505050565b60006118ba82611972565b15156118c557600080fd5b6001600083815260200190815260200160002060050160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600080600061193d858861241a565b841015915061194c858761241a565b8411158061195a5750600086145b90508180156119665750805b92505050949350505050565b600080821415801561198657506000548211155b9050919050565b60005481565b6119b26001600083815260200190815260200160002060010154612627565b15156119bd57600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002073__DLL___________________________________63366a5ba29091836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060206040518083038186803b158015611a7057600080fd5b505af4158015611a84573d6000803e3d6000fd5b505050506040513d6020811015611a9a57600080fd5b81019080805190602001909291905050501515611ab657600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002073__DLL___________________________________636d900ed09091836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060006040518083038186803b158015611b6957600080fd5b505af4158015611b7d573d6000803e3d6000fd5b505050503373ffffffffffffffffffffffffffffffffffffffff16817f402507661c8c8cb90e0a796450b8bdd28b6c516f05279c0cd29e84c344e1699a60405160405180910390a350565b60008282604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401828152602001925050506040518091039020905092915050565b80600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015611ce957600080fd5b505af1158015611cfd573d6000803e3d6000fd5b505050506040513d6020811015611d1357600080fd5b810190808051906020019092919050505010151515611d3157600080fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015611e7757600080fd5b505af1158015611e8b573d6000803e3d6000fd5b505050506040513d6020811015611ea157600080fd5b81019080805190602001909291905050501515611ebd57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff167ff7aaf024511d9982df8cd0d437c71c30106e6848cd1ba3d288d7a9c0e276aeda826040518082815260200191505060405180910390a250565b6000611f1982611972565b1515611f2457600080fd5b611f436001600084815260200190815260200160002060000154612627565b159050919050565b6000611f5682611972565b1515611f6157600080fd5b6001600083815260200190815260200160002060060160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000611fd584610fa9565b1515611fe057600080fd5b6001600085815260200190815260200160002060050160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561204c57600080fd5b6001600085815260200190815260200160002060060160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515156120b957600080fd5b6120c3338561251f565b60001916838360405180838152602001828152602001925050506040518091039020600019161415156120f557600080fd5b6120ff338561241a565b9050600183141561213357806001600086815260200190815260200160002060030160008282540192505081905550612158565b8060016000868152602001908152602001600020600401600082825401925050819055505b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002073__DLL___________________________________636d900ed09091866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060006040518083038186803b15801561220b57600080fd5b505af415801561221f573d6000803e3d6000fd5b50505050600180600086815260200190815260200160002060060160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff1683857ff42c78852433ace4bdcb44f6e80c8daae529e2d999c88cf6bf8f77b1e2890fdd84600160008a815260200190815260200160002060030154600160008b81526020019081526020016000206004015460405180848152602001838152602001828152602001935050505060405180910390a450505050565b60008060008061232f866128ab565b151561233a57600080fd5b6001600087815260200190815260200160002060060160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156123a657600080fd5b6123af86610ff8565b6123ba5760006123bd565b60015b60ff16925082856040518083815260200182815260200192505050604051809103902091506123ec878761251f565b90508060001916826000191614151561240457600080fd5b61240e878761241a565b93505050509392505050565b6000600473__AttributeStore________________________6350389f5c90916124448686611bc8565b6040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180838152602001826000191660001916815260200180602001828103825260098152602001807f6e756d546f6b656e730000000000000000000000000000000000000000000000815250602001935050505060206040518083038186803b1580156124dc57600080fd5b505af41580156124f0573d6000803e3d6000fd5b505050506040513d602081101561250657600080fd5b8101908080519060200190929190505050905092915050565b6000600473__AttributeStore________________________6350389f5c90916125498686611bc8565b6040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018260001916600019168152602001806020018281038252600a8152602001807f636f6d6d69744861736800000000000000000000000000000000000000000000815250602001935050505060206040518083038186803b1580156125e157600080fd5b505af41580156125f5573d6000803e3d6000fd5b505050506040513d602081101561260b57600080fd5b8101908080519060200190929190505050600102905092915050565b60008142119050919050565b600061268f612641336111bb565b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461292990919063ffffffff16565b90508181101515156126a057600080fd5b81600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540392505081905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156127b257600080fd5b505af11580156127c6573d6000803e3d6000fd5b505050506040513d60208110156127dc57600080fd5b810190808051906020019092919050505015156127f857600080fd5b3373ffffffffffffffffffffffffffffffffffffffff167ffaeb7dbb9992397d26ea1944efd40c80b40f702faf69b46c67ad10aba68ccb79836040518082815260200191505060405180910390a25050565b6000612855826128ab565b151561286057600080fd5b61286982610ff8565b1561288c57600160008381526020019081526020016000206004015490506128a6565b600160008381526020019081526020016000206003015490505b919050565b60006128b682611972565b15156128c157600080fd5b6128e06001600084815260200190815260200160002060010154612627565b9050919050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000818301905082811015151561292057fe5b80905092915050565b600082821115151561293757fe5b818303905092915050565b60a060405190810160405280600081526020016000815260200160008152602001600081526020016000815250905600a165627a7a723058205ac8e3cc95b29c2dc5ed648de7dfe6bf3923d68ff4344db368e0c72ae3d578a60029`

// DeployPLCRVotingContract deploys a new Ethereum contract, binding an instance of PLCRVotingContract to it.
func DeployPLCRVotingContract(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddr common.Address) (common.Address, *types.Transaction, *PLCRVotingContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PLCRVotingContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PLCRVotingContractBin), backend, _tokenAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PLCRVotingContract{PLCRVotingContractCaller: PLCRVotingContractCaller{contract: contract}, PLCRVotingContractTransactor: PLCRVotingContractTransactor{contract: contract}, PLCRVotingContractFilterer: PLCRVotingContractFilterer{contract: contract}}, nil
}

// PLCRVotingContract is an auto generated Go binding around an Ethereum contract.
type PLCRVotingContract struct {
	PLCRVotingContractCaller     // Read-only binding to the contract
	PLCRVotingContractTransactor // Write-only binding to the contract
	PLCRVotingContractFilterer   // Log filterer for contract events
}

// PLCRVotingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PLCRVotingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PLCRVotingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PLCRVotingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PLCRVotingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PLCRVotingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PLCRVotingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PLCRVotingContractSession struct {
	Contract     *PLCRVotingContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PLCRVotingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PLCRVotingContractCallerSession struct {
	Contract *PLCRVotingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PLCRVotingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PLCRVotingContractTransactorSession struct {
	Contract     *PLCRVotingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PLCRVotingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PLCRVotingContractRaw struct {
	Contract *PLCRVotingContract // Generic contract binding to access the raw methods on
}

// PLCRVotingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PLCRVotingContractCallerRaw struct {
	Contract *PLCRVotingContractCaller // Generic read-only contract binding to access the raw methods on
}

// PLCRVotingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PLCRVotingContractTransactorRaw struct {
	Contract *PLCRVotingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPLCRVotingContract creates a new instance of PLCRVotingContract, bound to a specific deployed contract.
func NewPLCRVotingContract(address common.Address, backend bind.ContractBackend) (*PLCRVotingContract, error) {
	contract, err := bindPLCRVotingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingContract{PLCRVotingContractCaller: PLCRVotingContractCaller{contract: contract}, PLCRVotingContractTransactor: PLCRVotingContractTransactor{contract: contract}, PLCRVotingContractFilterer: PLCRVotingContractFilterer{contract: contract}}, nil
}

// NewPLCRVotingContractCaller creates a new read-only instance of PLCRVotingContract, bound to a specific deployed contract.
func NewPLCRVotingContractCaller(address common.Address, caller bind.ContractCaller) (*PLCRVotingContractCaller, error) {
	contract, err := bindPLCRVotingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingContractCaller{contract: contract}, nil
}

// NewPLCRVotingContractTransactor creates a new write-only instance of PLCRVotingContract, bound to a specific deployed contract.
func NewPLCRVotingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PLCRVotingContractTransactor, error) {
	contract, err := bindPLCRVotingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingContractTransactor{contract: contract}, nil
}

// NewPLCRVotingContractFilterer creates a new log filterer instance of PLCRVotingContract, bound to a specific deployed contract.
func NewPLCRVotingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PLCRVotingContractFilterer, error) {
	contract, err := bindPLCRVotingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingContractFilterer{contract: contract}, nil
}

// bindPLCRVotingContract binds a generic wrapper to an already deployed contract.
func bindPLCRVotingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PLCRVotingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PLCRVotingContract *PLCRVotingContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PLCRVotingContract.Contract.PLCRVotingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PLCRVotingContract *PLCRVotingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.PLCRVotingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PLCRVotingContract *PLCRVotingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.PLCRVotingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PLCRVotingContract *PLCRVotingContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PLCRVotingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PLCRVotingContract *PLCRVotingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PLCRVotingContract *PLCRVotingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.contract.Transact(opts, method, params...)
}

// INITIALPOLLNONCE is a free data retrieval call binding the contract method 0x2173a10f.
//
// Solidity: function INITIAL_POLL_NONCE() constant returns(uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) INITIALPOLLNONCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "INITIAL_POLL_NONCE")
	return *ret0, err
}

// INITIALPOLLNONCE is a free data retrieval call binding the contract method 0x2173a10f.
//
// Solidity: function INITIAL_POLL_NONCE() constant returns(uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) INITIALPOLLNONCE() (*big.Int, error) {
	return _PLCRVotingContract.Contract.INITIALPOLLNONCE(&_PLCRVotingContract.CallOpts)
}

// INITIALPOLLNONCE is a free data retrieval call binding the contract method 0x2173a10f.
//
// Solidity: function INITIAL_POLL_NONCE() constant returns(uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) INITIALPOLLNONCE() (*big.Int, error) {
	return _PLCRVotingContract.Contract.INITIALPOLLNONCE(&_PLCRVotingContract.CallOpts)
}

// AttrUUID is a free data retrieval call binding the contract method 0xa1103f37.
//
// Solidity: function attrUUID(_user address, _pollID uint256) constant returns(UUID bytes32)
func (_PLCRVotingContract *PLCRVotingContractCaller) AttrUUID(opts *bind.CallOpts, _user common.Address, _pollID *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "attrUUID", _user, _pollID)
	return *ret0, err
}

// AttrUUID is a free data retrieval call binding the contract method 0xa1103f37.
//
// Solidity: function attrUUID(_user address, _pollID uint256) constant returns(UUID bytes32)
func (_PLCRVotingContract *PLCRVotingContractSession) AttrUUID(_user common.Address, _pollID *big.Int) ([32]byte, error) {
	return _PLCRVotingContract.Contract.AttrUUID(&_PLCRVotingContract.CallOpts, _user, _pollID)
}

// AttrUUID is a free data retrieval call binding the contract method 0xa1103f37.
//
// Solidity: function attrUUID(_user address, _pollID uint256) constant returns(UUID bytes32)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) AttrUUID(_user common.Address, _pollID *big.Int) ([32]byte, error) {
	return _PLCRVotingContract.Contract.AttrUUID(&_PLCRVotingContract.CallOpts, _user, _pollID)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(_pollID uint256) constant returns(active bool)
func (_PLCRVotingContract *PLCRVotingContractCaller) CommitPeriodActive(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "commitPeriodActive", _pollID)
	return *ret0, err
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(_pollID uint256) constant returns(active bool)
func (_PLCRVotingContract *PLCRVotingContractSession) CommitPeriodActive(_pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.CommitPeriodActive(&_PLCRVotingContract.CallOpts, _pollID)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(_pollID uint256) constant returns(active bool)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) CommitPeriodActive(_pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.CommitPeriodActive(&_PLCRVotingContract.CallOpts, _pollID)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(_voter address, _pollID uint256) constant returns(committed bool)
func (_PLCRVotingContract *PLCRVotingContractCaller) DidCommit(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "didCommit", _voter, _pollID)
	return *ret0, err
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(_voter address, _pollID uint256) constant returns(committed bool)
func (_PLCRVotingContract *PLCRVotingContractSession) DidCommit(_voter common.Address, _pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.DidCommit(&_PLCRVotingContract.CallOpts, _voter, _pollID)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(_voter address, _pollID uint256) constant returns(committed bool)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) DidCommit(_voter common.Address, _pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.DidCommit(&_PLCRVotingContract.CallOpts, _voter, _pollID)
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(_voter address, _pollID uint256) constant returns(revealed bool)
func (_PLCRVotingContract *PLCRVotingContractCaller) DidReveal(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "didReveal", _voter, _pollID)
	return *ret0, err
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(_voter address, _pollID uint256) constant returns(revealed bool)
func (_PLCRVotingContract *PLCRVotingContractSession) DidReveal(_voter common.Address, _pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.DidReveal(&_PLCRVotingContract.CallOpts, _voter, _pollID)
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(_voter address, _pollID uint256) constant returns(revealed bool)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) DidReveal(_voter common.Address, _pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.DidReveal(&_PLCRVotingContract.CallOpts, _voter, _pollID)
}

// GetCommitHash is a free data retrieval call binding the contract method 0xd901402b.
//
// Solidity: function getCommitHash(_voter address, _pollID uint256) constant returns(commitHash bytes32)
func (_PLCRVotingContract *PLCRVotingContractCaller) GetCommitHash(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "getCommitHash", _voter, _pollID)
	return *ret0, err
}

// GetCommitHash is a free data retrieval call binding the contract method 0xd901402b.
//
// Solidity: function getCommitHash(_voter address, _pollID uint256) constant returns(commitHash bytes32)
func (_PLCRVotingContract *PLCRVotingContractSession) GetCommitHash(_voter common.Address, _pollID *big.Int) ([32]byte, error) {
	return _PLCRVotingContract.Contract.GetCommitHash(&_PLCRVotingContract.CallOpts, _voter, _pollID)
}

// GetCommitHash is a free data retrieval call binding the contract method 0xd901402b.
//
// Solidity: function getCommitHash(_voter address, _pollID uint256) constant returns(commitHash bytes32)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) GetCommitHash(_voter common.Address, _pollID *big.Int) ([32]byte, error) {
	return _PLCRVotingContract.Contract.GetCommitHash(&_PLCRVotingContract.CallOpts, _voter, _pollID)
}

// GetInsertPointForNumTokens is a free data retrieval call binding the contract method 0x2c052031.
//
// Solidity: function getInsertPointForNumTokens(_voter address, _numTokens uint256, _pollID uint256) constant returns(prevNode uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) GetInsertPointForNumTokens(opts *bind.CallOpts, _voter common.Address, _numTokens *big.Int, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "getInsertPointForNumTokens", _voter, _numTokens, _pollID)
	return *ret0, err
}

// GetInsertPointForNumTokens is a free data retrieval call binding the contract method 0x2c052031.
//
// Solidity: function getInsertPointForNumTokens(_voter address, _numTokens uint256, _pollID uint256) constant returns(prevNode uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) GetInsertPointForNumTokens(_voter common.Address, _numTokens *big.Int, _pollID *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetInsertPointForNumTokens(&_PLCRVotingContract.CallOpts, _voter, _numTokens, _pollID)
}

// GetInsertPointForNumTokens is a free data retrieval call binding the contract method 0x2c052031.
//
// Solidity: function getInsertPointForNumTokens(_voter address, _numTokens uint256, _pollID uint256) constant returns(prevNode uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) GetInsertPointForNumTokens(_voter common.Address, _numTokens *big.Int, _pollID *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetInsertPointForNumTokens(&_PLCRVotingContract.CallOpts, _voter, _numTokens, _pollID)
}

// GetLastNode is a free data retrieval call binding the contract method 0x427fa1d2.
//
// Solidity: function getLastNode(_voter address) constant returns(pollID uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) GetLastNode(opts *bind.CallOpts, _voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "getLastNode", _voter)
	return *ret0, err
}

// GetLastNode is a free data retrieval call binding the contract method 0x427fa1d2.
//
// Solidity: function getLastNode(_voter address) constant returns(pollID uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) GetLastNode(_voter common.Address) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetLastNode(&_PLCRVotingContract.CallOpts, _voter)
}

// GetLastNode is a free data retrieval call binding the contract method 0x427fa1d2.
//
// Solidity: function getLastNode(_voter address) constant returns(pollID uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) GetLastNode(_voter common.Address) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetLastNode(&_PLCRVotingContract.CallOpts, _voter)
}

// GetLockedTokens is a free data retrieval call binding the contract method 0x6b2d95d4.
//
// Solidity: function getLockedTokens(_voter address) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) GetLockedTokens(opts *bind.CallOpts, _voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "getLockedTokens", _voter)
	return *ret0, err
}

// GetLockedTokens is a free data retrieval call binding the contract method 0x6b2d95d4.
//
// Solidity: function getLockedTokens(_voter address) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) GetLockedTokens(_voter common.Address) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetLockedTokens(&_PLCRVotingContract.CallOpts, _voter)
}

// GetLockedTokens is a free data retrieval call binding the contract method 0x6b2d95d4.
//
// Solidity: function getLockedTokens(_voter address) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) GetLockedTokens(_voter common.Address) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetLockedTokens(&_PLCRVotingContract.CallOpts, _voter)
}

// GetNumLosingTokens is a free data retrieval call binding the contract method 0x6afa97a8.
//
// Solidity: function getNumLosingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) GetNumLosingTokens(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "getNumLosingTokens", _voter, _pollID, _salt)
	return *ret0, err
}

// GetNumLosingTokens is a free data retrieval call binding the contract method 0x6afa97a8.
//
// Solidity: function getNumLosingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) GetNumLosingTokens(_voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetNumLosingTokens(&_PLCRVotingContract.CallOpts, _voter, _pollID, _salt)
}

// GetNumLosingTokens is a free data retrieval call binding the contract method 0x6afa97a8.
//
// Solidity: function getNumLosingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) GetNumLosingTokens(_voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetNumLosingTokens(&_PLCRVotingContract.CallOpts, _voter, _pollID, _salt)
}

// GetNumPassingTokens is a free data retrieval call binding the contract method 0xb43bd069.
//
// Solidity: function getNumPassingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) GetNumPassingTokens(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "getNumPassingTokens", _voter, _pollID, _salt)
	return *ret0, err
}

// GetNumPassingTokens is a free data retrieval call binding the contract method 0xb43bd069.
//
// Solidity: function getNumPassingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) GetNumPassingTokens(_voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetNumPassingTokens(&_PLCRVotingContract.CallOpts, _voter, _pollID, _salt)
}

// GetNumPassingTokens is a free data retrieval call binding the contract method 0xb43bd069.
//
// Solidity: function getNumPassingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) GetNumPassingTokens(_voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetNumPassingTokens(&_PLCRVotingContract.CallOpts, _voter, _pollID, _salt)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xd1382092.
//
// Solidity: function getNumTokens(_voter address, _pollID uint256) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) GetNumTokens(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "getNumTokens", _voter, _pollID)
	return *ret0, err
}

// GetNumTokens is a free data retrieval call binding the contract method 0xd1382092.
//
// Solidity: function getNumTokens(_voter address, _pollID uint256) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) GetNumTokens(_voter common.Address, _pollID *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetNumTokens(&_PLCRVotingContract.CallOpts, _voter, _pollID)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xd1382092.
//
// Solidity: function getNumTokens(_voter address, _pollID uint256) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) GetNumTokens(_voter common.Address, _pollID *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetNumTokens(&_PLCRVotingContract.CallOpts, _voter, _pollID)
}

// GetTotalNumberOfTokensForLosingOption is a free data retrieval call binding the contract method 0xe8cfa3f0.
//
// Solidity: function getTotalNumberOfTokensForLosingOption(_pollID uint256) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) GetTotalNumberOfTokensForLosingOption(opts *bind.CallOpts, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "getTotalNumberOfTokensForLosingOption", _pollID)
	return *ret0, err
}

// GetTotalNumberOfTokensForLosingOption is a free data retrieval call binding the contract method 0xe8cfa3f0.
//
// Solidity: function getTotalNumberOfTokensForLosingOption(_pollID uint256) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) GetTotalNumberOfTokensForLosingOption(_pollID *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetTotalNumberOfTokensForLosingOption(&_PLCRVotingContract.CallOpts, _pollID)
}

// GetTotalNumberOfTokensForLosingOption is a free data retrieval call binding the contract method 0xe8cfa3f0.
//
// Solidity: function getTotalNumberOfTokensForLosingOption(_pollID uint256) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) GetTotalNumberOfTokensForLosingOption(_pollID *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetTotalNumberOfTokensForLosingOption(&_PLCRVotingContract.CallOpts, _pollID)
}

// GetTotalNumberOfTokensForWinningOption is a free data retrieval call binding the contract method 0x053e71a6.
//
// Solidity: function getTotalNumberOfTokensForWinningOption(_pollID uint256) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) GetTotalNumberOfTokensForWinningOption(opts *bind.CallOpts, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "getTotalNumberOfTokensForWinningOption", _pollID)
	return *ret0, err
}

// GetTotalNumberOfTokensForWinningOption is a free data retrieval call binding the contract method 0x053e71a6.
//
// Solidity: function getTotalNumberOfTokensForWinningOption(_pollID uint256) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) GetTotalNumberOfTokensForWinningOption(_pollID *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetTotalNumberOfTokensForWinningOption(&_PLCRVotingContract.CallOpts, _pollID)
}

// GetTotalNumberOfTokensForWinningOption is a free data retrieval call binding the contract method 0x053e71a6.
//
// Solidity: function getTotalNumberOfTokensForWinningOption(_pollID uint256) constant returns(numTokens uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) GetTotalNumberOfTokensForWinningOption(_pollID *big.Int) (*big.Int, error) {
	return _PLCRVotingContract.Contract.GetTotalNumberOfTokensForWinningOption(&_PLCRVotingContract.CallOpts, _pollID)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_PLCRVotingContract *PLCRVotingContractCaller) IsExpired(opts *bind.CallOpts, _terminationDate *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "isExpired", _terminationDate)
	return *ret0, err
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_PLCRVotingContract *PLCRVotingContractSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.IsExpired(&_PLCRVotingContract.CallOpts, _terminationDate)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.IsExpired(&_PLCRVotingContract.CallOpts, _terminationDate)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(_pollID uint256) constant returns(passed bool)
func (_PLCRVotingContract *PLCRVotingContractCaller) IsPassed(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "isPassed", _pollID)
	return *ret0, err
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(_pollID uint256) constant returns(passed bool)
func (_PLCRVotingContract *PLCRVotingContractSession) IsPassed(_pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.IsPassed(&_PLCRVotingContract.CallOpts, _pollID)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(_pollID uint256) constant returns(passed bool)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) IsPassed(_pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.IsPassed(&_PLCRVotingContract.CallOpts, _pollID)
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(_pollID uint256) constant returns(ended bool)
func (_PLCRVotingContract *PLCRVotingContractCaller) PollEnded(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "pollEnded", _pollID)
	return *ret0, err
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(_pollID uint256) constant returns(ended bool)
func (_PLCRVotingContract *PLCRVotingContractSession) PollEnded(_pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.PollEnded(&_PLCRVotingContract.CallOpts, _pollID)
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(_pollID uint256) constant returns(ended bool)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) PollEnded(_pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.PollEnded(&_PLCRVotingContract.CallOpts, _pollID)
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(_pollID uint256) constant returns(exists bool)
func (_PLCRVotingContract *PLCRVotingContractCaller) PollExists(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "pollExists", _pollID)
	return *ret0, err
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(_pollID uint256) constant returns(exists bool)
func (_PLCRVotingContract *PLCRVotingContractSession) PollExists(_pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.PollExists(&_PLCRVotingContract.CallOpts, _pollID)
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(_pollID uint256) constant returns(exists bool)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) PollExists(_pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.PollExists(&_PLCRVotingContract.CallOpts, _pollID)
}

// PollMap is a free data retrieval call binding the contract method 0x6148fed5.
//
// Solidity: function pollMap( uint256) constant returns(commitEndDate uint256, revealEndDate uint256, voteQuorum uint256, votesFor uint256, votesAgainst uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) PollMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CommitEndDate *big.Int
	RevealEndDate *big.Int
	VoteQuorum    *big.Int
	VotesFor      *big.Int
	VotesAgainst  *big.Int
}, error) {
	ret := new(struct {
		CommitEndDate *big.Int
		RevealEndDate *big.Int
		VoteQuorum    *big.Int
		VotesFor      *big.Int
		VotesAgainst  *big.Int
	})
	out := ret
	err := _PLCRVotingContract.contract.Call(opts, out, "pollMap", arg0)
	return *ret, err
}

// PollMap is a free data retrieval call binding the contract method 0x6148fed5.
//
// Solidity: function pollMap( uint256) constant returns(commitEndDate uint256, revealEndDate uint256, voteQuorum uint256, votesFor uint256, votesAgainst uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) PollMap(arg0 *big.Int) (struct {
	CommitEndDate *big.Int
	RevealEndDate *big.Int
	VoteQuorum    *big.Int
	VotesFor      *big.Int
	VotesAgainst  *big.Int
}, error) {
	return _PLCRVotingContract.Contract.PollMap(&_PLCRVotingContract.CallOpts, arg0)
}

// PollMap is a free data retrieval call binding the contract method 0x6148fed5.
//
// Solidity: function pollMap( uint256) constant returns(commitEndDate uint256, revealEndDate uint256, voteQuorum uint256, votesFor uint256, votesAgainst uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) PollMap(arg0 *big.Int) (struct {
	CommitEndDate *big.Int
	RevealEndDate *big.Int
	VoteQuorum    *big.Int
	VotesFor      *big.Int
	VotesAgainst  *big.Int
}, error) {
	return _PLCRVotingContract.Contract.PollMap(&_PLCRVotingContract.CallOpts, arg0)
}

// PollNonce is a free data retrieval call binding the contract method 0x97508f36.
//
// Solidity: function pollNonce() constant returns(uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) PollNonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "pollNonce")
	return *ret0, err
}

// PollNonce is a free data retrieval call binding the contract method 0x97508f36.
//
// Solidity: function pollNonce() constant returns(uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) PollNonce() (*big.Int, error) {
	return _PLCRVotingContract.Contract.PollNonce(&_PLCRVotingContract.CallOpts)
}

// PollNonce is a free data retrieval call binding the contract method 0x97508f36.
//
// Solidity: function pollNonce() constant returns(uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) PollNonce() (*big.Int, error) {
	return _PLCRVotingContract.Contract.PollNonce(&_PLCRVotingContract.CallOpts)
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(_pollID uint256) constant returns(active bool)
func (_PLCRVotingContract *PLCRVotingContractCaller) RevealPeriodActive(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "revealPeriodActive", _pollID)
	return *ret0, err
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(_pollID uint256) constant returns(active bool)
func (_PLCRVotingContract *PLCRVotingContractSession) RevealPeriodActive(_pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.RevealPeriodActive(&_PLCRVotingContract.CallOpts, _pollID)
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(_pollID uint256) constant returns(active bool)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) RevealPeriodActive(_pollID *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.RevealPeriodActive(&_PLCRVotingContract.CallOpts, _pollID)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PLCRVotingContract *PLCRVotingContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PLCRVotingContract *PLCRVotingContractSession) Token() (common.Address, error) {
	return _PLCRVotingContract.Contract.Token(&_PLCRVotingContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) Token() (common.Address, error) {
	return _PLCRVotingContract.Contract.Token(&_PLCRVotingContract.CallOpts)
}

// ValidPosition is a free data retrieval call binding the contract method 0x819b0293.
//
// Solidity: function validPosition(_prevID uint256, _nextID uint256, _voter address, _numTokens uint256) constant returns(valid bool)
func (_PLCRVotingContract *PLCRVotingContractCaller) ValidPosition(opts *bind.CallOpts, _prevID *big.Int, _nextID *big.Int, _voter common.Address, _numTokens *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "validPosition", _prevID, _nextID, _voter, _numTokens)
	return *ret0, err
}

// ValidPosition is a free data retrieval call binding the contract method 0x819b0293.
//
// Solidity: function validPosition(_prevID uint256, _nextID uint256, _voter address, _numTokens uint256) constant returns(valid bool)
func (_PLCRVotingContract *PLCRVotingContractSession) ValidPosition(_prevID *big.Int, _nextID *big.Int, _voter common.Address, _numTokens *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.ValidPosition(&_PLCRVotingContract.CallOpts, _prevID, _nextID, _voter, _numTokens)
}

// ValidPosition is a free data retrieval call binding the contract method 0x819b0293.
//
// Solidity: function validPosition(_prevID uint256, _nextID uint256, _voter address, _numTokens uint256) constant returns(valid bool)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) ValidPosition(_prevID *big.Int, _nextID *big.Int, _voter common.Address, _numTokens *big.Int) (bool, error) {
	return _PLCRVotingContract.Contract.ValidPosition(&_PLCRVotingContract.CallOpts, _prevID, _nextID, _voter, _numTokens)
}

// VoteTokenBalance is a free data retrieval call binding the contract method 0x3b930294.
//
// Solidity: function voteTokenBalance( address) constant returns(uint256)
func (_PLCRVotingContract *PLCRVotingContractCaller) VoteTokenBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVotingContract.contract.Call(opts, out, "voteTokenBalance", arg0)
	return *ret0, err
}

// VoteTokenBalance is a free data retrieval call binding the contract method 0x3b930294.
//
// Solidity: function voteTokenBalance( address) constant returns(uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) VoteTokenBalance(arg0 common.Address) (*big.Int, error) {
	return _PLCRVotingContract.Contract.VoteTokenBalance(&_PLCRVotingContract.CallOpts, arg0)
}

// VoteTokenBalance is a free data retrieval call binding the contract method 0x3b930294.
//
// Solidity: function voteTokenBalance( address) constant returns(uint256)
func (_PLCRVotingContract *PLCRVotingContractCallerSession) VoteTokenBalance(arg0 common.Address) (*big.Int, error) {
	return _PLCRVotingContract.Contract.VoteTokenBalance(&_PLCRVotingContract.CallOpts, arg0)
}

// CommitVote is a paid mutator transaction binding the contract method 0x6cbf9c5e.
//
// Solidity: function commitVote(_pollID uint256, _secretHash bytes32, _numTokens uint256, _prevPollID uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactor) CommitVote(opts *bind.TransactOpts, _pollID *big.Int, _secretHash [32]byte, _numTokens *big.Int, _prevPollID *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.contract.Transact(opts, "commitVote", _pollID, _secretHash, _numTokens, _prevPollID)
}

// CommitVote is a paid mutator transaction binding the contract method 0x6cbf9c5e.
//
// Solidity: function commitVote(_pollID uint256, _secretHash bytes32, _numTokens uint256, _prevPollID uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractSession) CommitVote(_pollID *big.Int, _secretHash [32]byte, _numTokens *big.Int, _prevPollID *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.CommitVote(&_PLCRVotingContract.TransactOpts, _pollID, _secretHash, _numTokens, _prevPollID)
}

// CommitVote is a paid mutator transaction binding the contract method 0x6cbf9c5e.
//
// Solidity: function commitVote(_pollID uint256, _secretHash bytes32, _numTokens uint256, _prevPollID uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactorSession) CommitVote(_pollID *big.Int, _secretHash [32]byte, _numTokens *big.Int, _prevPollID *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.CommitVote(&_PLCRVotingContract.TransactOpts, _pollID, _secretHash, _numTokens, _prevPollID)
}

// RequestVotingRights is a paid mutator transaction binding the contract method 0xa25236fe.
//
// Solidity: function requestVotingRights(_numTokens uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactor) RequestVotingRights(opts *bind.TransactOpts, _numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.contract.Transact(opts, "requestVotingRights", _numTokens)
}

// RequestVotingRights is a paid mutator transaction binding the contract method 0xa25236fe.
//
// Solidity: function requestVotingRights(_numTokens uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractSession) RequestVotingRights(_numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.RequestVotingRights(&_PLCRVotingContract.TransactOpts, _numTokens)
}

// RequestVotingRights is a paid mutator transaction binding the contract method 0xa25236fe.
//
// Solidity: function requestVotingRights(_numTokens uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactorSession) RequestVotingRights(_numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.RequestVotingRights(&_PLCRVotingContract.TransactOpts, _numTokens)
}

// RescueTokens is a paid mutator transaction binding the contract method 0x97603560.
//
// Solidity: function rescueTokens(_pollID uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactor) RescueTokens(opts *bind.TransactOpts, _pollID *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.contract.Transact(opts, "rescueTokens", _pollID)
}

// RescueTokens is a paid mutator transaction binding the contract method 0x97603560.
//
// Solidity: function rescueTokens(_pollID uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractSession) RescueTokens(_pollID *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.RescueTokens(&_PLCRVotingContract.TransactOpts, _pollID)
}

// RescueTokens is a paid mutator transaction binding the contract method 0x97603560.
//
// Solidity: function rescueTokens(_pollID uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactorSession) RescueTokens(_pollID *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.RescueTokens(&_PLCRVotingContract.TransactOpts, _pollID)
}

// RevealVote is a paid mutator transaction binding the contract method 0xb11d8bb8.
//
// Solidity: function revealVote(_pollID uint256, _voteOption uint256, _salt uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactor) RevealVote(opts *bind.TransactOpts, _pollID *big.Int, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.contract.Transact(opts, "revealVote", _pollID, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0xb11d8bb8.
//
// Solidity: function revealVote(_pollID uint256, _voteOption uint256, _salt uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractSession) RevealVote(_pollID *big.Int, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.RevealVote(&_PLCRVotingContract.TransactOpts, _pollID, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0xb11d8bb8.
//
// Solidity: function revealVote(_pollID uint256, _voteOption uint256, _salt uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactorSession) RevealVote(_pollID *big.Int, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.RevealVote(&_PLCRVotingContract.TransactOpts, _pollID, _voteOption, _salt)
}

// StartPoll is a paid mutator transaction binding the contract method 0x32ed3d60.
//
// Solidity: function startPoll(_voteQuorum uint256, _commitDuration uint256, _revealDuration uint256) returns(pollID uint256)
func (_PLCRVotingContract *PLCRVotingContractTransactor) StartPoll(opts *bind.TransactOpts, _voteQuorum *big.Int, _commitDuration *big.Int, _revealDuration *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.contract.Transact(opts, "startPoll", _voteQuorum, _commitDuration, _revealDuration)
}

// StartPoll is a paid mutator transaction binding the contract method 0x32ed3d60.
//
// Solidity: function startPoll(_voteQuorum uint256, _commitDuration uint256, _revealDuration uint256) returns(pollID uint256)
func (_PLCRVotingContract *PLCRVotingContractSession) StartPoll(_voteQuorum *big.Int, _commitDuration *big.Int, _revealDuration *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.StartPoll(&_PLCRVotingContract.TransactOpts, _voteQuorum, _commitDuration, _revealDuration)
}

// StartPoll is a paid mutator transaction binding the contract method 0x32ed3d60.
//
// Solidity: function startPoll(_voteQuorum uint256, _commitDuration uint256, _revealDuration uint256) returns(pollID uint256)
func (_PLCRVotingContract *PLCRVotingContractTransactorSession) StartPoll(_voteQuorum *big.Int, _commitDuration *big.Int, _revealDuration *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.StartPoll(&_PLCRVotingContract.TransactOpts, _voteQuorum, _commitDuration, _revealDuration)
}

// WithdrawVotingRights is a paid mutator transaction binding the contract method 0xe7b1d43c.
//
// Solidity: function withdrawVotingRights(_numTokens uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactor) WithdrawVotingRights(opts *bind.TransactOpts, _numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.contract.Transact(opts, "withdrawVotingRights", _numTokens)
}

// WithdrawVotingRights is a paid mutator transaction binding the contract method 0xe7b1d43c.
//
// Solidity: function withdrawVotingRights(_numTokens uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractSession) WithdrawVotingRights(_numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.WithdrawVotingRights(&_PLCRVotingContract.TransactOpts, _numTokens)
}

// WithdrawVotingRights is a paid mutator transaction binding the contract method 0xe7b1d43c.
//
// Solidity: function withdrawVotingRights(_numTokens uint256) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactorSession) WithdrawVotingRights(_numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.WithdrawVotingRights(&_PLCRVotingContract.TransactOpts, _numTokens)
}

// PLCRVotingContractPollCreatedIterator is returned from FilterPollCreated and is used to iterate over the raw logs and unpacked data for PollCreated events raised by the PLCRVotingContract contract.
type PLCRVotingContractPollCreatedIterator struct {
	Event *PLCRVotingContractPollCreated // Event containing the contract specifics and raw log

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
func (it *PLCRVotingContractPollCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingContractPollCreated)
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
		it.Event = new(PLCRVotingContractPollCreated)
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
func (it *PLCRVotingContractPollCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingContractPollCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingContractPollCreated represents a PollCreated event raised by the PLCRVotingContract contract.
type PLCRVotingContractPollCreated struct {
	VoteQuorum    *big.Int
	CommitEndDate *big.Int
	RevealEndDate *big.Int
	PollID        *big.Int
	Creator       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPollCreated is a free log retrieval operation binding the contract event 0x404f1f1c229d9eb2a949e7584da6ffde9d059ef2169f487ca815434cce0640d0.
//
// Solidity: event _PollCreated(voteQuorum uint256, commitEndDate uint256, revealEndDate uint256, pollID indexed uint256, creator indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) FilterPollCreated(opts *bind.FilterOpts, pollID []*big.Int, creator []common.Address) (*PLCRVotingContractPollCreatedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.FilterLogs(opts, "_PollCreated", pollIDRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingContractPollCreatedIterator{contract: _PLCRVotingContract.contract, event: "_PollCreated", logs: logs, sub: sub}, nil
}

// WatchPollCreated is a free log subscription operation binding the contract event 0x404f1f1c229d9eb2a949e7584da6ffde9d059ef2169f487ca815434cce0640d0.
//
// Solidity: event _PollCreated(voteQuorum uint256, commitEndDate uint256, revealEndDate uint256, pollID indexed uint256, creator indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) WatchPollCreated(opts *bind.WatchOpts, sink chan<- *PLCRVotingContractPollCreated, pollID []*big.Int, creator []common.Address) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.WatchLogs(opts, "_PollCreated", pollIDRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingContractPollCreated)
				if err := _PLCRVotingContract.contract.UnpackLog(event, "_PollCreated", log); err != nil {
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

// PLCRVotingContractTokensRescuedIterator is returned from FilterTokensRescued and is used to iterate over the raw logs and unpacked data for TokensRescued events raised by the PLCRVotingContract contract.
type PLCRVotingContractTokensRescuedIterator struct {
	Event *PLCRVotingContractTokensRescued // Event containing the contract specifics and raw log

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
func (it *PLCRVotingContractTokensRescuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingContractTokensRescued)
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
		it.Event = new(PLCRVotingContractTokensRescued)
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
func (it *PLCRVotingContractTokensRescuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingContractTokensRescuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingContractTokensRescued represents a TokensRescued event raised by the PLCRVotingContract contract.
type PLCRVotingContractTokensRescued struct {
	PollID *big.Int
	Voter  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensRescued is a free log retrieval operation binding the contract event 0x402507661c8c8cb90e0a796450b8bdd28b6c516f05279c0cd29e84c344e1699a.
//
// Solidity: event _TokensRescued(pollID indexed uint256, voter indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) FilterTokensRescued(opts *bind.FilterOpts, pollID []*big.Int, voter []common.Address) (*PLCRVotingContractTokensRescuedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.FilterLogs(opts, "_TokensRescued", pollIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingContractTokensRescuedIterator{contract: _PLCRVotingContract.contract, event: "_TokensRescued", logs: logs, sub: sub}, nil
}

// WatchTokensRescued is a free log subscription operation binding the contract event 0x402507661c8c8cb90e0a796450b8bdd28b6c516f05279c0cd29e84c344e1699a.
//
// Solidity: event _TokensRescued(pollID indexed uint256, voter indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) WatchTokensRescued(opts *bind.WatchOpts, sink chan<- *PLCRVotingContractTokensRescued, pollID []*big.Int, voter []common.Address) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.WatchLogs(opts, "_TokensRescued", pollIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingContractTokensRescued)
				if err := _PLCRVotingContract.contract.UnpackLog(event, "_TokensRescued", log); err != nil {
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

// PLCRVotingContractVoteCommittedIterator is returned from FilterVoteCommitted and is used to iterate over the raw logs and unpacked data for VoteCommitted events raised by the PLCRVotingContract contract.
type PLCRVotingContractVoteCommittedIterator struct {
	Event *PLCRVotingContractVoteCommitted // Event containing the contract specifics and raw log

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
func (it *PLCRVotingContractVoteCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingContractVoteCommitted)
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
		it.Event = new(PLCRVotingContractVoteCommitted)
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
func (it *PLCRVotingContractVoteCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingContractVoteCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingContractVoteCommitted represents a VoteCommitted event raised by the PLCRVotingContract contract.
type PLCRVotingContractVoteCommitted struct {
	PollID    *big.Int
	NumTokens *big.Int
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoteCommitted is a free log retrieval operation binding the contract event 0xea7979e4280d7e6bffc1c7d83a1ac99f16d02ecc14465ce41016226783b663d7.
//
// Solidity: event _VoteCommitted(pollID indexed uint256, numTokens uint256, voter indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) FilterVoteCommitted(opts *bind.FilterOpts, pollID []*big.Int, voter []common.Address) (*PLCRVotingContractVoteCommittedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.FilterLogs(opts, "_VoteCommitted", pollIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingContractVoteCommittedIterator{contract: _PLCRVotingContract.contract, event: "_VoteCommitted", logs: logs, sub: sub}, nil
}

// WatchVoteCommitted is a free log subscription operation binding the contract event 0xea7979e4280d7e6bffc1c7d83a1ac99f16d02ecc14465ce41016226783b663d7.
//
// Solidity: event _VoteCommitted(pollID indexed uint256, numTokens uint256, voter indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) WatchVoteCommitted(opts *bind.WatchOpts, sink chan<- *PLCRVotingContractVoteCommitted, pollID []*big.Int, voter []common.Address) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.WatchLogs(opts, "_VoteCommitted", pollIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingContractVoteCommitted)
				if err := _PLCRVotingContract.contract.UnpackLog(event, "_VoteCommitted", log); err != nil {
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

// PLCRVotingContractVoteRevealedIterator is returned from FilterVoteRevealed and is used to iterate over the raw logs and unpacked data for VoteRevealed events raised by the PLCRVotingContract contract.
type PLCRVotingContractVoteRevealedIterator struct {
	Event *PLCRVotingContractVoteRevealed // Event containing the contract specifics and raw log

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
func (it *PLCRVotingContractVoteRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingContractVoteRevealed)
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
		it.Event = new(PLCRVotingContractVoteRevealed)
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
func (it *PLCRVotingContractVoteRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingContractVoteRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingContractVoteRevealed represents a VoteRevealed event raised by the PLCRVotingContract contract.
type PLCRVotingContractVoteRevealed struct {
	PollID       *big.Int
	NumTokens    *big.Int
	VotesFor     *big.Int
	VotesAgainst *big.Int
	Choice       *big.Int
	Voter        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVoteRevealed is a free log retrieval operation binding the contract event 0xf42c78852433ace4bdcb44f6e80c8daae529e2d999c88cf6bf8f77b1e2890fdd.
//
// Solidity: event _VoteRevealed(pollID indexed uint256, numTokens uint256, votesFor uint256, votesAgainst uint256, choice indexed uint256, voter indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) FilterVoteRevealed(opts *bind.FilterOpts, pollID []*big.Int, choice []*big.Int, voter []common.Address) (*PLCRVotingContractVoteRevealedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	var choiceRule []interface{}
	for _, choiceItem := range choice {
		choiceRule = append(choiceRule, choiceItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.FilterLogs(opts, "_VoteRevealed", pollIDRule, choiceRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingContractVoteRevealedIterator{contract: _PLCRVotingContract.contract, event: "_VoteRevealed", logs: logs, sub: sub}, nil
}

// WatchVoteRevealed is a free log subscription operation binding the contract event 0xf42c78852433ace4bdcb44f6e80c8daae529e2d999c88cf6bf8f77b1e2890fdd.
//
// Solidity: event _VoteRevealed(pollID indexed uint256, numTokens uint256, votesFor uint256, votesAgainst uint256, choice indexed uint256, voter indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) WatchVoteRevealed(opts *bind.WatchOpts, sink chan<- *PLCRVotingContractVoteRevealed, pollID []*big.Int, choice []*big.Int, voter []common.Address) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	var choiceRule []interface{}
	for _, choiceItem := range choice {
		choiceRule = append(choiceRule, choiceItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.WatchLogs(opts, "_VoteRevealed", pollIDRule, choiceRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingContractVoteRevealed)
				if err := _PLCRVotingContract.contract.UnpackLog(event, "_VoteRevealed", log); err != nil {
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

// PLCRVotingContractVotingRightsGrantedIterator is returned from FilterVotingRightsGranted and is used to iterate over the raw logs and unpacked data for VotingRightsGranted events raised by the PLCRVotingContract contract.
type PLCRVotingContractVotingRightsGrantedIterator struct {
	Event *PLCRVotingContractVotingRightsGranted // Event containing the contract specifics and raw log

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
func (it *PLCRVotingContractVotingRightsGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingContractVotingRightsGranted)
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
		it.Event = new(PLCRVotingContractVotingRightsGranted)
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
func (it *PLCRVotingContractVotingRightsGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingContractVotingRightsGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingContractVotingRightsGranted represents a VotingRightsGranted event raised by the PLCRVotingContract contract.
type PLCRVotingContractVotingRightsGranted struct {
	NumTokens *big.Int
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVotingRightsGranted is a free log retrieval operation binding the contract event 0xf7aaf024511d9982df8cd0d437c71c30106e6848cd1ba3d288d7a9c0e276aeda.
//
// Solidity: event _VotingRightsGranted(numTokens uint256, voter indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) FilterVotingRightsGranted(opts *bind.FilterOpts, voter []common.Address) (*PLCRVotingContractVotingRightsGrantedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.FilterLogs(opts, "_VotingRightsGranted", voterRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingContractVotingRightsGrantedIterator{contract: _PLCRVotingContract.contract, event: "_VotingRightsGranted", logs: logs, sub: sub}, nil
}

// WatchVotingRightsGranted is a free log subscription operation binding the contract event 0xf7aaf024511d9982df8cd0d437c71c30106e6848cd1ba3d288d7a9c0e276aeda.
//
// Solidity: event _VotingRightsGranted(numTokens uint256, voter indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) WatchVotingRightsGranted(opts *bind.WatchOpts, sink chan<- *PLCRVotingContractVotingRightsGranted, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.WatchLogs(opts, "_VotingRightsGranted", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingContractVotingRightsGranted)
				if err := _PLCRVotingContract.contract.UnpackLog(event, "_VotingRightsGranted", log); err != nil {
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

// PLCRVotingContractVotingRightsWithdrawnIterator is returned from FilterVotingRightsWithdrawn and is used to iterate over the raw logs and unpacked data for VotingRightsWithdrawn events raised by the PLCRVotingContract contract.
type PLCRVotingContractVotingRightsWithdrawnIterator struct {
	Event *PLCRVotingContractVotingRightsWithdrawn // Event containing the contract specifics and raw log

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
func (it *PLCRVotingContractVotingRightsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingContractVotingRightsWithdrawn)
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
		it.Event = new(PLCRVotingContractVotingRightsWithdrawn)
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
func (it *PLCRVotingContractVotingRightsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingContractVotingRightsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingContractVotingRightsWithdrawn represents a VotingRightsWithdrawn event raised by the PLCRVotingContract contract.
type PLCRVotingContractVotingRightsWithdrawn struct {
	NumTokens *big.Int
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVotingRightsWithdrawn is a free log retrieval operation binding the contract event 0xfaeb7dbb9992397d26ea1944efd40c80b40f702faf69b46c67ad10aba68ccb79.
//
// Solidity: event _VotingRightsWithdrawn(numTokens uint256, voter indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) FilterVotingRightsWithdrawn(opts *bind.FilterOpts, voter []common.Address) (*PLCRVotingContractVotingRightsWithdrawnIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.FilterLogs(opts, "_VotingRightsWithdrawn", voterRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingContractVotingRightsWithdrawnIterator{contract: _PLCRVotingContract.contract, event: "_VotingRightsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchVotingRightsWithdrawn is a free log subscription operation binding the contract event 0xfaeb7dbb9992397d26ea1944efd40c80b40f702faf69b46c67ad10aba68ccb79.
//
// Solidity: event _VotingRightsWithdrawn(numTokens uint256, voter indexed address)
func (_PLCRVotingContract *PLCRVotingContractFilterer) WatchVotingRightsWithdrawn(opts *bind.WatchOpts, sink chan<- *PLCRVotingContractVotingRightsWithdrawn, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVotingContract.contract.WatchLogs(opts, "_VotingRightsWithdrawn", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingContractVotingRightsWithdrawn)
				if err := _PLCRVotingContract.contract.UnpackLog(event, "_VotingRightsWithdrawn", log); err != nil {
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
