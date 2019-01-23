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

// CivilPLCRVotingContractABI is the input ABI used to generate the binding from.
const CivilPLCRVotingContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getTotalNumberOfTokensForWinningOption\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_POLL_NONCE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_numTokens\",\"type\":\"uint256\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getInsertPointForNumTokens\",\"outputs\":[{\"name\":\"prevNode\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteQuorum\",\"type\":\"uint256\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"}],\"name\":\"startPoll\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"voteTokenBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollIDs\",\"type\":\"uint256[]\"},{\"name\":\"_secretHashes\",\"type\":\"bytes32[]\"},{\"name\":\"_numsTokens\",\"type\":\"uint256[]\"},{\"name\":\"_prevPollIDs\",\"type\":\"uint256[]\"}],\"name\":\"commitVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"telemetry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getLastNode\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"revealPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"isPassed\",\"outputs\":[{\"name\":\"passed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pollMap\",\"outputs\":[{\"name\":\"commitEndDate\",\"type\":\"uint256\"},{\"name\":\"revealEndDate\",\"type\":\"uint256\"},{\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"name\":\"votesFor\",\"type\":\"uint256\"},{\"name\":\"votesAgainst\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getLockedTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"},{\"name\":\"_numTokens\",\"type\":\"uint256\"},{\"name\":\"_prevPollID\",\"type\":\"uint256\"}],\"name\":\"commitVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"didCommit\",\"outputs\":[{\"name\":\"committed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollIDs\",\"type\":\"uint256[]\"},{\"name\":\"_voteOptions\",\"type\":\"uint256[]\"},{\"name\":\"_salts\",\"type\":\"uint256[]\"}],\"name\":\"revealVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prevID\",\"type\":\"uint256\"},{\"name\":\"_nextID\",\"type\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_numTokens\",\"type\":\"uint256\"}],\"name\":\"validPosition\",\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"pollExists\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pollNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"attrUUID\",\"outputs\":[{\"name\":\"UUID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"commitPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"didReveal\",\"outputs\":[{\"name\":\"revealed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"getNumPassingTokens\",\"outputs\":[{\"name\":\"correctVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollIDs\",\"type\":\"uint256[]\"}],\"name\":\"rescueTokensInMultiplePolls\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getNumTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getCommitHash\",\"outputs\":[{\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_terminationDate\",\"type\":\"uint256\"}],\"name\":\"isExpired\",\"outputs\":[{\"name\":\"expired\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_numTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"pollEnded\",\"outputs\":[{\"name\":\"ended\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"telemetryAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_VoteCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesFor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesAgainst\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"choice\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"_VoteRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commitEndDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealEndDate\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"_PollCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_VotingRightsGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_VotingRightsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_TokensRescued\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_numTokens\",\"type\":\"uint256\"}],\"name\":\"requestVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"getNumLosingTokens\",\"outputs\":[{\"name\":\"correctVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getTotalNumberOfTokensForLosingOption\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CivilPLCRVotingContractBin is the compiled bytecode used for deploying new contracts.
const CivilPLCRVotingContractBin = `0x608060405234801561001057600080fd5b50604051604080611eb583398101604052805160209091015181600160a060020a038116151561003f57600080fd5b60058054600160a060020a031916600160a060020a03928316179055600080558116151561006c57600080fd5b60068054600160a060020a031916600160a060020a039290921691909117905550611e198061009c6000396000f3006080604052600436106101925763ffffffff60e060020a600035041663053e71a681146101975780632173a10f146101c15780632c052031146101d657806332ed3d60146101fd5780633b9302941461021b5780633ec36b991461023c57806340f41c1914610282578063427fa1d2146102b3578063441c77c0146102d457806349403183146103005780636148fed5146103185780636afa97a81461035b5780636b2d95d4146103825780636cbf9c5e146103a35780637f97e836146103c45780638090f92e146103e8578063819b02931461042057806388d21ff31461044a57806397508f36146104625780639760356014610477578063a1103f371461048f578063a25236fe146104b3578063a4439dc5146104cb578063aa7ca464146104e3578063b11d8bb814610507578063b43bd06914610525578063bb11ed7e1461054c578063d1382092146105a1578063d901402b146105c5578063d9548e53146105e9578063e7b1d43c14610601578063e8cfa3f014610619578063ee68483014610631578063fc0c546a14610649575b600080fd5b3480156101a357600080fd5b506101af60043561065e565b60408051918252519081900360200190f35b3480156101cd57600080fd5b506101af6106b3565b3480156101e257600080fd5b506101af600160a060020a03600435166024356044356106b8565b34801561020957600080fd5b506101af600435602435604435610881565b34801561022757600080fd5b506101af600160a060020a036004351661095d565b34801561024857600080fd5b50610280602460048035828101929082013591813580830192908201359160443580830192908201359160643591820191013561096f565b005b34801561028e57600080fd5b50610297610a0f565b60408051600160a060020a039092168252519081900360200190f35b3480156102bf57600080fd5b506101af600160a060020a0360043516610a1e565b3480156102e057600080fd5b506102ec600435610ad8565b604080519115158252519081900360200190f35b34801561030c57600080fd5b506102ec600435610b23565b34801561032457600080fd5b50610330600435610ba0565b6040805195865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561036757600080fd5b506101af600160a060020a0360043516602435604435610bce565b34801561038e57600080fd5b506101af600160a060020a0360043516610c85565b3480156103af57600080fd5b50610280600435602435604435606435610c99565b3480156103d057600080fd5b506102ec600160a060020a036004351660243561120b565b3480156103f457600080fd5b506102806024600480358281019290820135918135808301929082013591604435918201910135611250565b34801561042c57600080fd5b506102ec600435602435600160a060020a03604435166064356112cc565b34801561045657600080fd5b506102ec60043561130a565b34801561046e57600080fd5b506101af61131f565b34801561048357600080fd5b50610280600435611325565b34801561049b57600080fd5b506101af600160a060020a03600435166024356114c9565b3480156104bf57600080fd5b50610280600435611500565b3480156104d757600080fd5b506102ec60043561159a565b3480156104ef57600080fd5b506102ec600160a060020a03600435166024356115c8565b34801561051357600080fd5b5061028060043560243560443561160d565b34801561053157600080fd5b506101af600160a060020a0360043516602435604435611810565b34801561055857600080fd5b50604080516020600480358082013583810280860185019096528085526102809536959394602494938501929182918501908490808284375094975061189c9650505050505050565b3480156105ad57600080fd5b506101af600160a060020a03600435166024356118d4565b3480156105d157600080fd5b506101af600160a060020a03600435166024356119a6565b3480156105f557600080fd5b506102ec600435611a45565b34801561060d57600080fd5b50610280600435611a4a565b34801561062557600080fd5b506101af600435611b76565b34801561063d57600080fd5b506102ec600435611bca565b34801561065557600080fd5b50610297611bfc565b600061066982611bca565b151561067457600080fd5b61067d82610b23565b1561069a57506000818152600160205260409020600301546106ae565b506000818152600160205260409020600401545b919050565b600081565b60008060006106c686610a1e565b91506106d286836118d4565b90505b8115610874576106e586836118d4565b90508481116107b757838214156107af57600160a060020a03861660009081526003602090815260409182902082517f30fe0a0a000000000000000000000000000000000000000000000000000000008152600481019190915260248101859052915173__DLL___________________________________926330fe0a0a926044808301939192829003018186803b15801561078057600080fd5b505af4158015610794573d6000803e3d6000fd5b505050506040513d60208110156107aa57600080fd5b505191505b819250610878565b600160a060020a03861660009081526003602090815260409182902082517f30fe0a0a000000000000000000000000000000000000000000000000000000008152600481019190915260248101859052915173__DLL___________________________________926330fe0a0a926044808301939192829003018186803b15801561084157600080fd5b505af4158015610855573d6000803e3d6000fd5b505050506040513d602081101561086b57600080fd5b505191506106d5565b8192505b50509392505050565b600080546001018155808061089c428663ffffffff611c0b16565b91506108ae828563ffffffff611c0b16565b6040805160a08101825284815260208082018481528284018b8152600060608086018281526080870183815283548452600180885289852098518955955195880195909555925160028701559151600386015591516004909401939093555483518b8152918201879052818401859052925193945033937f404f1f1c229d9eb2a949e7584da6ffde9d059ef2169f487ca815434cce0640d0929181900390910190a35050600054949350505050565b60026020526000908152604090205481565b600087861461097d57600080fd5b87841461098957600080fd5b87821461099557600080fd5b5060005b87811015610a04576109fc8989838181106109b057fe5b9050602002013588888481811015156109c557fe5b602002919091013590508787858181106109db57fe5b9050602002013586868681811015156109f057fe5b90506020020135610c99565b600101610999565b505050505050505050565b600654600160a060020a031681565b600160a060020a038116600090815260036020908152604080832081517f30fe0a0a000000000000000000000000000000000000000000000000000000008152600481019190915260248101849052905173__DLL___________________________________926330fe0a0a9260448082019391829003018186803b158015610aa657600080fd5b505af4158015610aba573d6000803e3d6000fd5b505050506040513d6020811015610ad057600080fd5b505192915050565b6000610ae38261130a565b1515610aee57600080fd5b60008281526001602081905260409091200154610b0a90611a45565b158015610b1d5750610b1b8261159a565b155b92915050565b6000610b2d611dbd565b610b3683611bca565b1515610b4157600080fd5b5050600090815260016020818152604092839020835160a0810185528154815292810154918301919091526002810154928201839052600381015460608301819052600490910154608090920182905290810190910260649091021190565b6001602081905260009182526040909120805491810154600282015460038301546004909301549192909185565b600080600080610bdd86611bca565b1515610be857600080fd5b6000868152600160209081526040808320600160a060020a038b16845260060190915290205460ff161515610c1c57600080fd5b610c2586610b23565b610c30576001610c33565b60005b60ff1692508285604051808381526020018281526020019250505060405180910390209150610c6287876119a6565b9050818114610c7057600080fd5b610c7a87876118d4565b979650505050505050565b6000610b1d82610c9484610a1e565b6118d4565b6000806000610ca78761159a565b1515610cb257600080fd5b33600090815260026020526040902054851115610cf55733600090815260026020526040902054610cea90869063ffffffff611c1816565b9250610cf583611500565b33600090815260026020526040902054851115610d1157600080fd5b861515610d1d57600080fd5b851515610d2957600080fd5b831580610ddf57503360009081526003602090815260409182902082517f366a5ba2000000000000000000000000000000000000000000000000000000008152600481019190915260248101879052915173__DLL___________________________________9263366a5ba2926044808301939192829003018186803b158015610db257600080fd5b505af4158015610dc6573d6000803e3d6000fd5b505050506040513d6020811015610ddc57600080fd5b50515b1515610dea57600080fd5b3360009081526003602090815260409182902082517f07d29ac9000000000000000000000000000000000000000000000000000000008152600481019190915260248101879052915173__DLL___________________________________926307d29ac9926044808301939192829003018186803b158015610e6b57600080fd5b505af4158015610e7f573d6000803e3d6000fd5b505050506040513d6020811015610e9557600080fd5b5051915086821415610f51573360009081526003602090815260409182902082517f07d29ac90000000000000000000000000000000000000000000000000000000081526004810191909152602481018a9052915173__DLL___________________________________926307d29ac9926044808301939192829003018186803b158015610f2257600080fd5b505af4158015610f36573d6000803e3d6000fd5b505050506040513d6020811015610f4c57600080fd5b505191505b610f5d848333886112cc565b1515610f6857600080fd5b3360009081526003602052604080822081517f9735c51b000000000000000000000000000000000000000000000000000000008152600481019190915260248101879052604481018a905260648101859052905173__DLL___________________________________92639735c51b9260848082019391829003018186803b158015610ff357600080fd5b505af4158015611007573d6000803e3d6000fd5b5050505061101533886114c9565b604080517f977aa031000000000000000000000000000000000000000000000000000000008152600481810152602481018390526064810188905260806044820152600960848201527f6e756d546f6b656e73000000000000000000000000000000000000000000000060a4820152905191925073__AttributeStore________________________9163977aa0319160c480820192600092909190829003018186803b1580156110c557600080fd5b505af41580156110d9573d6000803e3d6000fd5b5050604080517f977aa03100000000000000000000000000000000000000000000000000000000815260048181015260248101859052606481018a905260806044820152600a60848201527f636f6d6d6974486173680000000000000000000000000000000000000000000060a4820152905173__AttributeStore________________________935063977aa031925060c4808301926000929190829003018186803b15801561118957600080fd5b505af415801561119d573d6000803e3d6000fd5b5050506000888152600160208181526040808420338086526005909101835293819020805460ff1916909317909255815189815291519293508a927fea7979e4280d7e6bffc1c7d83a1ac99f16d02ecc14465ce41016226783b663d79281900390910190a350505050505050565b60006112168261130a565b151561122157600080fd5b506000908152600160209081526040808320600160a060020a0394909416835260059093019052205460ff1690565b600085841461125e57600080fd5b85821461126a57600080fd5b5060005b858110156112c3576112bb87878381811061128557fe5b90506020020135868684818110151561129a57fe5b9050602002013585858581811015156112af57fe5b9050602002013561160d565b60010161126e565b50505050505050565b60008060006112db85886118d4565b84101591506112ea85876118d4565b841115806112f6575085155b9050818015610c7a57509695505050505050565b60008115801590610b1d575050600054101590565b60005481565b6000818152600160208190526040909120015461134190611a45565b151561134c57600080fd5b3360009081526003602090815260409182902082517f366a5ba2000000000000000000000000000000000000000000000000000000008152600481019190915260248101849052915173__DLL___________________________________9263366a5ba2926044808301939192829003018186803b1580156113cd57600080fd5b505af41580156113e1573d6000803e3d6000fd5b505050506040513d60208110156113f757600080fd5b5051151561140457600080fd5b3360009081526003602052604080822081517f6d900ed0000000000000000000000000000000000000000000000000000000008152600481019190915260248101849052905173__DLL___________________________________92636d900ed09260448082019391829003018186803b15801561148157600080fd5b505af4158015611495573d6000803e3d6000fd5b50506040513392508391507f402507661c8c8cb90e0a796450b8bdd28b6c516f05279c0cd29e84c344e1699a90600090a350565b604080516c01000000000000000000000000600160a060020a03851602815260148101839052905190819003603401902092915050565b61150981611c2a565b600654336000818152600260205260408082205481517f725248730000000000000000000000000000000000000000000000000000000081526004810194909452602484015251600160a060020a0390931692637252487392604480820193929182900301818387803b15801561157f57600080fd5b505af1158015611593573d6000803e3d6000fd5b5050505050565b60006115a58261130a565b15156115b057600080fd5b600082815260016020526040902054610b1b90611a45565b60006115d38261130a565b15156115de57600080fd5b506000908152600160209081526040808320600160a060020a0394909416835260069093019052205460ff1690565b600061161884610ad8565b151561162357600080fd5b600084815260016020908152604080832033845260050190915290205460ff16151561164e57600080fd5b600084815260016020908152604080832033845260060190915290205460ff161561167857600080fd5b61168233856119a6565b60408051858152602081018590528151908190039091019020146116a557600080fd5b6116af33856118d4565b905082600114156116d65760008481526001602052604090206003018054820190556116ee565b60008481526001602052604090206004018054820190555b3360009081526003602052604080822081517f6d900ed0000000000000000000000000000000000000000000000000000000008152600481019190915260248101879052905173__DLL___________________________________92636d900ed09260448082019391829003018186803b15801561176b57600080fd5b505af415801561177f573d6000803e3d6000fd5b505050600085815260016020818152604080842033808652600682018452828620805460ff191686179055948a9052928252600383015460049093015481518781529283019390935281810192909252606081018690529051919250859187917f9b19aaec524fad29c0ced9b9973a15e3045d7c3be156d71394ab40f0d5f119ff919081900360800190a450505050565b60008060008061181f86611bca565b151561182a57600080fd5b6000868152600160209081526040808320600160a060020a038b16845260060190915290205460ff16151561185e57600080fd5b61186786610b23565b611872576000610c33565b6040805160018082526020820188905282519182900390920190209093509150610c6287876119a6565b60005b81518110156118d0576118c882828151811015156118b957fe5b90602001906020020151611325565b60010161189f565b5050565b600073__AttributeStore________________________6350389f5c60046118fc86866114c9565b6040805160e060020a63ffffffff86160281526004810193909352602483019190915260606044830152600960648301527f6e756d546f6b656e73000000000000000000000000000000000000000000000060848301525160a4808301926020929190829003018186803b15801561197357600080fd5b505af4158015611987573d6000803e3d6000fd5b505050506040513d602081101561199d57600080fd5b50519392505050565b600073__AttributeStore________________________6350389f5c60046119ce86866114c9565b6040805160e060020a63ffffffff86160281526004810193909352602483019190915260606044830152600a60648301527f636f6d6d6974486173680000000000000000000000000000000000000000000060848301525160a4808301926020929190829003018186803b15801561197357600080fd5b421190565b6000611a74611a5833610c85565b336000908152600260205260409020549063ffffffff611c1816565b905081811015611a8357600080fd5b3360008181526002602090815260408083208054879003905560055481517fa9059cbb0000000000000000000000000000000000000000000000000000000081526004810195909552602485018790529051600160a060020a039091169363a9059cbb9360448083019493928390030190829087803b158015611b0557600080fd5b505af1158015611b19573d6000803e3d6000fd5b505050506040513d6020811015611b2f57600080fd5b50511515611b3c57600080fd5b60408051838152905133917ffaeb7dbb9992397d26ea1944efd40c80b40f702faf69b46c67ad10aba68ccb79919081900360200190a25050565b6000611b8182611bca565b1515611b8c57600080fd5b611b9582610b23565b15611bb257506000818152600160205260409020600401546106ae565b506000818152600160205260409020600301546106ae565b6000611bd58261130a565b1515611be057600080fd5b60008281526001602081905260409091200154610b1d90611a45565b600554600160a060020a031681565b81810182811015610b1d57fe5b600082821115611c2457fe5b50900390565b600554604080517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015290518392600160a060020a0316916370a082319160248083019260209291908290030181600087803b158015611c8f57600080fd5b505af1158015611ca3573d6000803e3d6000fd5b505050506040513d6020811015611cb957600080fd5b50511015611cc657600080fd5b33600081815260026020908152604080832080548601905560055481517f23b872dd0000000000000000000000000000000000000000000000000000000081526004810195909552306024860152604485018690529051600160a060020a03909116936323b872dd9360648083019493928390030190829087803b158015611d4d57600080fd5b505af1158015611d61573d6000803e3d6000fd5b505050506040513d6020811015611d7757600080fd5b50511515611d8457600080fd5b60408051828152905133917ff7aaf024511d9982df8cd0d437c71c30106e6848cd1ba3d288d7a9c0e276aeda919081900360200190a250565b60a060405190810160405280600081526020016000815260200160008152602001600081526020016000815250905600a165627a7a72305820de2c9039769dc173b477b6ba51655a1fde18db43b707914a50068712b9dce4950029`

// DeployCivilPLCRVotingContract deploys a new Ethereum contract, binding an instance of CivilPLCRVotingContract to it.
func DeployCivilPLCRVotingContract(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddr common.Address, telemetryAddr common.Address) (common.Address, *types.Transaction, *CivilPLCRVotingContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CivilPLCRVotingContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CivilPLCRVotingContractBin), backend, tokenAddr, telemetryAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CivilPLCRVotingContract{CivilPLCRVotingContractCaller: CivilPLCRVotingContractCaller{contract: contract}, CivilPLCRVotingContractTransactor: CivilPLCRVotingContractTransactor{contract: contract}, CivilPLCRVotingContractFilterer: CivilPLCRVotingContractFilterer{contract: contract}}, nil
}

// CivilPLCRVotingContract is an auto generated Go binding around an Ethereum contract.
type CivilPLCRVotingContract struct {
	CivilPLCRVotingContractCaller     // Read-only binding to the contract
	CivilPLCRVotingContractTransactor // Write-only binding to the contract
	CivilPLCRVotingContractFilterer   // Log filterer for contract events
}

// CivilPLCRVotingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CivilPLCRVotingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilPLCRVotingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CivilPLCRVotingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilPLCRVotingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CivilPLCRVotingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilPLCRVotingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CivilPLCRVotingContractSession struct {
	Contract     *CivilPLCRVotingContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CivilPLCRVotingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CivilPLCRVotingContractCallerSession struct {
	Contract *CivilPLCRVotingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// CivilPLCRVotingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CivilPLCRVotingContractTransactorSession struct {
	Contract     *CivilPLCRVotingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// CivilPLCRVotingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CivilPLCRVotingContractRaw struct {
	Contract *CivilPLCRVotingContract // Generic contract binding to access the raw methods on
}

// CivilPLCRVotingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CivilPLCRVotingContractCallerRaw struct {
	Contract *CivilPLCRVotingContractCaller // Generic read-only contract binding to access the raw methods on
}

// CivilPLCRVotingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CivilPLCRVotingContractTransactorRaw struct {
	Contract *CivilPLCRVotingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCivilPLCRVotingContract creates a new instance of CivilPLCRVotingContract, bound to a specific deployed contract.
func NewCivilPLCRVotingContract(address common.Address, backend bind.ContractBackend) (*CivilPLCRVotingContract, error) {
	contract, err := bindCivilPLCRVotingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CivilPLCRVotingContract{CivilPLCRVotingContractCaller: CivilPLCRVotingContractCaller{contract: contract}, CivilPLCRVotingContractTransactor: CivilPLCRVotingContractTransactor{contract: contract}, CivilPLCRVotingContractFilterer: CivilPLCRVotingContractFilterer{contract: contract}}, nil
}

// NewCivilPLCRVotingContractCaller creates a new read-only instance of CivilPLCRVotingContract, bound to a specific deployed contract.
func NewCivilPLCRVotingContractCaller(address common.Address, caller bind.ContractCaller) (*CivilPLCRVotingContractCaller, error) {
	contract, err := bindCivilPLCRVotingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CivilPLCRVotingContractCaller{contract: contract}, nil
}

// NewCivilPLCRVotingContractTransactor creates a new write-only instance of CivilPLCRVotingContract, bound to a specific deployed contract.
func NewCivilPLCRVotingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CivilPLCRVotingContractTransactor, error) {
	contract, err := bindCivilPLCRVotingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CivilPLCRVotingContractTransactor{contract: contract}, nil
}

// NewCivilPLCRVotingContractFilterer creates a new log filterer instance of CivilPLCRVotingContract, bound to a specific deployed contract.
func NewCivilPLCRVotingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CivilPLCRVotingContractFilterer, error) {
	contract, err := bindCivilPLCRVotingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CivilPLCRVotingContractFilterer{contract: contract}, nil
}

// bindCivilPLCRVotingContract binds a generic wrapper to an already deployed contract.
func bindCivilPLCRVotingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CivilPLCRVotingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CivilPLCRVotingContract *CivilPLCRVotingContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CivilPLCRVotingContract.Contract.CivilPLCRVotingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CivilPLCRVotingContract *CivilPLCRVotingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.CivilPLCRVotingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CivilPLCRVotingContract *CivilPLCRVotingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.CivilPLCRVotingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CivilPLCRVotingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.contract.Transact(opts, method, params...)
}

// INITIALPOLLNONCE is a free data retrieval call binding the contract method 0x2173a10f.
//
// Solidity: function INITIAL_POLL_NONCE() constant returns(uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) INITIALPOLLNONCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "INITIAL_POLL_NONCE")
	return *ret0, err
}

// INITIALPOLLNONCE is a free data retrieval call binding the contract method 0x2173a10f.
//
// Solidity: function INITIAL_POLL_NONCE() constant returns(uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) INITIALPOLLNONCE() (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.INITIALPOLLNONCE(&_CivilPLCRVotingContract.CallOpts)
}

// INITIALPOLLNONCE is a free data retrieval call binding the contract method 0x2173a10f.
//
// Solidity: function INITIAL_POLL_NONCE() constant returns(uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) INITIALPOLLNONCE() (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.INITIALPOLLNONCE(&_CivilPLCRVotingContract.CallOpts)
}

// AttrUUID is a free data retrieval call binding the contract method 0xa1103f37.
//
// Solidity: function attrUUID(_user address, _pollID uint256) constant returns(UUID bytes32)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) AttrUUID(opts *bind.CallOpts, _user common.Address, _pollID *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "attrUUID", _user, _pollID)
	return *ret0, err
}

// AttrUUID is a free data retrieval call binding the contract method 0xa1103f37.
//
// Solidity: function attrUUID(_user address, _pollID uint256) constant returns(UUID bytes32)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) AttrUUID(_user common.Address, _pollID *big.Int) ([32]byte, error) {
	return _CivilPLCRVotingContract.Contract.AttrUUID(&_CivilPLCRVotingContract.CallOpts, _user, _pollID)
}

// AttrUUID is a free data retrieval call binding the contract method 0xa1103f37.
//
// Solidity: function attrUUID(_user address, _pollID uint256) constant returns(UUID bytes32)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) AttrUUID(_user common.Address, _pollID *big.Int) ([32]byte, error) {
	return _CivilPLCRVotingContract.Contract.AttrUUID(&_CivilPLCRVotingContract.CallOpts, _user, _pollID)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(_pollID uint256) constant returns(active bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) CommitPeriodActive(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "commitPeriodActive", _pollID)
	return *ret0, err
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(_pollID uint256) constant returns(active bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) CommitPeriodActive(_pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.CommitPeriodActive(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(_pollID uint256) constant returns(active bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) CommitPeriodActive(_pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.CommitPeriodActive(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(_voter address, _pollID uint256) constant returns(committed bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) DidCommit(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "didCommit", _voter, _pollID)
	return *ret0, err
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(_voter address, _pollID uint256) constant returns(committed bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) DidCommit(_voter common.Address, _pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.DidCommit(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(_voter address, _pollID uint256) constant returns(committed bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) DidCommit(_voter common.Address, _pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.DidCommit(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID)
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(_voter address, _pollID uint256) constant returns(revealed bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) DidReveal(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "didReveal", _voter, _pollID)
	return *ret0, err
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(_voter address, _pollID uint256) constant returns(revealed bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) DidReveal(_voter common.Address, _pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.DidReveal(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID)
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(_voter address, _pollID uint256) constant returns(revealed bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) DidReveal(_voter common.Address, _pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.DidReveal(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID)
}

// GetCommitHash is a free data retrieval call binding the contract method 0xd901402b.
//
// Solidity: function getCommitHash(_voter address, _pollID uint256) constant returns(commitHash bytes32)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) GetCommitHash(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "getCommitHash", _voter, _pollID)
	return *ret0, err
}

// GetCommitHash is a free data retrieval call binding the contract method 0xd901402b.
//
// Solidity: function getCommitHash(_voter address, _pollID uint256) constant returns(commitHash bytes32)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) GetCommitHash(_voter common.Address, _pollID *big.Int) ([32]byte, error) {
	return _CivilPLCRVotingContract.Contract.GetCommitHash(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID)
}

// GetCommitHash is a free data retrieval call binding the contract method 0xd901402b.
//
// Solidity: function getCommitHash(_voter address, _pollID uint256) constant returns(commitHash bytes32)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) GetCommitHash(_voter common.Address, _pollID *big.Int) ([32]byte, error) {
	return _CivilPLCRVotingContract.Contract.GetCommitHash(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID)
}

// GetInsertPointForNumTokens is a free data retrieval call binding the contract method 0x2c052031.
//
// Solidity: function getInsertPointForNumTokens(_voter address, _numTokens uint256, _pollID uint256) constant returns(prevNode uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) GetInsertPointForNumTokens(opts *bind.CallOpts, _voter common.Address, _numTokens *big.Int, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "getInsertPointForNumTokens", _voter, _numTokens, _pollID)
	return *ret0, err
}

// GetInsertPointForNumTokens is a free data retrieval call binding the contract method 0x2c052031.
//
// Solidity: function getInsertPointForNumTokens(_voter address, _numTokens uint256, _pollID uint256) constant returns(prevNode uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) GetInsertPointForNumTokens(_voter common.Address, _numTokens *big.Int, _pollID *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetInsertPointForNumTokens(&_CivilPLCRVotingContract.CallOpts, _voter, _numTokens, _pollID)
}

// GetInsertPointForNumTokens is a free data retrieval call binding the contract method 0x2c052031.
//
// Solidity: function getInsertPointForNumTokens(_voter address, _numTokens uint256, _pollID uint256) constant returns(prevNode uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) GetInsertPointForNumTokens(_voter common.Address, _numTokens *big.Int, _pollID *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetInsertPointForNumTokens(&_CivilPLCRVotingContract.CallOpts, _voter, _numTokens, _pollID)
}

// GetLastNode is a free data retrieval call binding the contract method 0x427fa1d2.
//
// Solidity: function getLastNode(_voter address) constant returns(pollID uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) GetLastNode(opts *bind.CallOpts, _voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "getLastNode", _voter)
	return *ret0, err
}

// GetLastNode is a free data retrieval call binding the contract method 0x427fa1d2.
//
// Solidity: function getLastNode(_voter address) constant returns(pollID uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) GetLastNode(_voter common.Address) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetLastNode(&_CivilPLCRVotingContract.CallOpts, _voter)
}

// GetLastNode is a free data retrieval call binding the contract method 0x427fa1d2.
//
// Solidity: function getLastNode(_voter address) constant returns(pollID uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) GetLastNode(_voter common.Address) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetLastNode(&_CivilPLCRVotingContract.CallOpts, _voter)
}

// GetLockedTokens is a free data retrieval call binding the contract method 0x6b2d95d4.
//
// Solidity: function getLockedTokens(_voter address) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) GetLockedTokens(opts *bind.CallOpts, _voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "getLockedTokens", _voter)
	return *ret0, err
}

// GetLockedTokens is a free data retrieval call binding the contract method 0x6b2d95d4.
//
// Solidity: function getLockedTokens(_voter address) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) GetLockedTokens(_voter common.Address) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetLockedTokens(&_CivilPLCRVotingContract.CallOpts, _voter)
}

// GetLockedTokens is a free data retrieval call binding the contract method 0x6b2d95d4.
//
// Solidity: function getLockedTokens(_voter address) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) GetLockedTokens(_voter common.Address) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetLockedTokens(&_CivilPLCRVotingContract.CallOpts, _voter)
}

// GetNumLosingTokens is a free data retrieval call binding the contract method 0x6afa97a8.
//
// Solidity: function getNumLosingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) GetNumLosingTokens(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "getNumLosingTokens", _voter, _pollID, _salt)
	return *ret0, err
}

// GetNumLosingTokens is a free data retrieval call binding the contract method 0x6afa97a8.
//
// Solidity: function getNumLosingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) GetNumLosingTokens(_voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetNumLosingTokens(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID, _salt)
}

// GetNumLosingTokens is a free data retrieval call binding the contract method 0x6afa97a8.
//
// Solidity: function getNumLosingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) GetNumLosingTokens(_voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetNumLosingTokens(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID, _salt)
}

// GetNumPassingTokens is a free data retrieval call binding the contract method 0xb43bd069.
//
// Solidity: function getNumPassingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) GetNumPassingTokens(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "getNumPassingTokens", _voter, _pollID, _salt)
	return *ret0, err
}

// GetNumPassingTokens is a free data retrieval call binding the contract method 0xb43bd069.
//
// Solidity: function getNumPassingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) GetNumPassingTokens(_voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetNumPassingTokens(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID, _salt)
}

// GetNumPassingTokens is a free data retrieval call binding the contract method 0xb43bd069.
//
// Solidity: function getNumPassingTokens(_voter address, _pollID uint256, _salt uint256) constant returns(correctVotes uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) GetNumPassingTokens(_voter common.Address, _pollID *big.Int, _salt *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetNumPassingTokens(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID, _salt)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xd1382092.
//
// Solidity: function getNumTokens(_voter address, _pollID uint256) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) GetNumTokens(opts *bind.CallOpts, _voter common.Address, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "getNumTokens", _voter, _pollID)
	return *ret0, err
}

// GetNumTokens is a free data retrieval call binding the contract method 0xd1382092.
//
// Solidity: function getNumTokens(_voter address, _pollID uint256) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) GetNumTokens(_voter common.Address, _pollID *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetNumTokens(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xd1382092.
//
// Solidity: function getNumTokens(_voter address, _pollID uint256) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) GetNumTokens(_voter common.Address, _pollID *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetNumTokens(&_CivilPLCRVotingContract.CallOpts, _voter, _pollID)
}

// GetTotalNumberOfTokensForLosingOption is a free data retrieval call binding the contract method 0xe8cfa3f0.
//
// Solidity: function getTotalNumberOfTokensForLosingOption(_pollID uint256) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) GetTotalNumberOfTokensForLosingOption(opts *bind.CallOpts, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "getTotalNumberOfTokensForLosingOption", _pollID)
	return *ret0, err
}

// GetTotalNumberOfTokensForLosingOption is a free data retrieval call binding the contract method 0xe8cfa3f0.
//
// Solidity: function getTotalNumberOfTokensForLosingOption(_pollID uint256) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) GetTotalNumberOfTokensForLosingOption(_pollID *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetTotalNumberOfTokensForLosingOption(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// GetTotalNumberOfTokensForLosingOption is a free data retrieval call binding the contract method 0xe8cfa3f0.
//
// Solidity: function getTotalNumberOfTokensForLosingOption(_pollID uint256) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) GetTotalNumberOfTokensForLosingOption(_pollID *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetTotalNumberOfTokensForLosingOption(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// GetTotalNumberOfTokensForWinningOption is a free data retrieval call binding the contract method 0x053e71a6.
//
// Solidity: function getTotalNumberOfTokensForWinningOption(_pollID uint256) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) GetTotalNumberOfTokensForWinningOption(opts *bind.CallOpts, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "getTotalNumberOfTokensForWinningOption", _pollID)
	return *ret0, err
}

// GetTotalNumberOfTokensForWinningOption is a free data retrieval call binding the contract method 0x053e71a6.
//
// Solidity: function getTotalNumberOfTokensForWinningOption(_pollID uint256) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) GetTotalNumberOfTokensForWinningOption(_pollID *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetTotalNumberOfTokensForWinningOption(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// GetTotalNumberOfTokensForWinningOption is a free data retrieval call binding the contract method 0x053e71a6.
//
// Solidity: function getTotalNumberOfTokensForWinningOption(_pollID uint256) constant returns(numTokens uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) GetTotalNumberOfTokensForWinningOption(_pollID *big.Int) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.GetTotalNumberOfTokensForWinningOption(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) IsExpired(opts *bind.CallOpts, _terminationDate *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "isExpired", _terminationDate)
	return *ret0, err
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.IsExpired(&_CivilPLCRVotingContract.CallOpts, _terminationDate)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.IsExpired(&_CivilPLCRVotingContract.CallOpts, _terminationDate)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(_pollID uint256) constant returns(passed bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) IsPassed(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "isPassed", _pollID)
	return *ret0, err
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(_pollID uint256) constant returns(passed bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) IsPassed(_pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.IsPassed(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(_pollID uint256) constant returns(passed bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) IsPassed(_pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.IsPassed(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(_pollID uint256) constant returns(ended bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) PollEnded(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "pollEnded", _pollID)
	return *ret0, err
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(_pollID uint256) constant returns(ended bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) PollEnded(_pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.PollEnded(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(_pollID uint256) constant returns(ended bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) PollEnded(_pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.PollEnded(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(_pollID uint256) constant returns(exists bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) PollExists(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "pollExists", _pollID)
	return *ret0, err
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(_pollID uint256) constant returns(exists bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) PollExists(_pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.PollExists(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(_pollID uint256) constant returns(exists bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) PollExists(_pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.PollExists(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// PollMap is a free data retrieval call binding the contract method 0x6148fed5.
//
// Solidity: function pollMap( uint256) constant returns(commitEndDate uint256, revealEndDate uint256, voteQuorum uint256, votesFor uint256, votesAgainst uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) PollMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "pollMap", arg0)
	return *ret, err
}

// PollMap is a free data retrieval call binding the contract method 0x6148fed5.
//
// Solidity: function pollMap( uint256) constant returns(commitEndDate uint256, revealEndDate uint256, voteQuorum uint256, votesFor uint256, votesAgainst uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) PollMap(arg0 *big.Int) (struct {
	CommitEndDate *big.Int
	RevealEndDate *big.Int
	VoteQuorum    *big.Int
	VotesFor      *big.Int
	VotesAgainst  *big.Int
}, error) {
	return _CivilPLCRVotingContract.Contract.PollMap(&_CivilPLCRVotingContract.CallOpts, arg0)
}

// PollMap is a free data retrieval call binding the contract method 0x6148fed5.
//
// Solidity: function pollMap( uint256) constant returns(commitEndDate uint256, revealEndDate uint256, voteQuorum uint256, votesFor uint256, votesAgainst uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) PollMap(arg0 *big.Int) (struct {
	CommitEndDate *big.Int
	RevealEndDate *big.Int
	VoteQuorum    *big.Int
	VotesFor      *big.Int
	VotesAgainst  *big.Int
}, error) {
	return _CivilPLCRVotingContract.Contract.PollMap(&_CivilPLCRVotingContract.CallOpts, arg0)
}

// PollNonce is a free data retrieval call binding the contract method 0x97508f36.
//
// Solidity: function pollNonce() constant returns(uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) PollNonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "pollNonce")
	return *ret0, err
}

// PollNonce is a free data retrieval call binding the contract method 0x97508f36.
//
// Solidity: function pollNonce() constant returns(uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) PollNonce() (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.PollNonce(&_CivilPLCRVotingContract.CallOpts)
}

// PollNonce is a free data retrieval call binding the contract method 0x97508f36.
//
// Solidity: function pollNonce() constant returns(uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) PollNonce() (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.PollNonce(&_CivilPLCRVotingContract.CallOpts)
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(_pollID uint256) constant returns(active bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) RevealPeriodActive(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "revealPeriodActive", _pollID)
	return *ret0, err
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(_pollID uint256) constant returns(active bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) RevealPeriodActive(_pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.RevealPeriodActive(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(_pollID uint256) constant returns(active bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) RevealPeriodActive(_pollID *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.RevealPeriodActive(&_CivilPLCRVotingContract.CallOpts, _pollID)
}

// Telemetry is a free data retrieval call binding the contract method 0x40f41c19.
//
// Solidity: function telemetry() constant returns(address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) Telemetry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "telemetry")
	return *ret0, err
}

// Telemetry is a free data retrieval call binding the contract method 0x40f41c19.
//
// Solidity: function telemetry() constant returns(address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) Telemetry() (common.Address, error) {
	return _CivilPLCRVotingContract.Contract.Telemetry(&_CivilPLCRVotingContract.CallOpts)
}

// Telemetry is a free data retrieval call binding the contract method 0x40f41c19.
//
// Solidity: function telemetry() constant returns(address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) Telemetry() (common.Address, error) {
	return _CivilPLCRVotingContract.Contract.Telemetry(&_CivilPLCRVotingContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) Token() (common.Address, error) {
	return _CivilPLCRVotingContract.Contract.Token(&_CivilPLCRVotingContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) Token() (common.Address, error) {
	return _CivilPLCRVotingContract.Contract.Token(&_CivilPLCRVotingContract.CallOpts)
}

// ValidPosition is a free data retrieval call binding the contract method 0x819b0293.
//
// Solidity: function validPosition(_prevID uint256, _nextID uint256, _voter address, _numTokens uint256) constant returns(valid bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) ValidPosition(opts *bind.CallOpts, _prevID *big.Int, _nextID *big.Int, _voter common.Address, _numTokens *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "validPosition", _prevID, _nextID, _voter, _numTokens)
	return *ret0, err
}

// ValidPosition is a free data retrieval call binding the contract method 0x819b0293.
//
// Solidity: function validPosition(_prevID uint256, _nextID uint256, _voter address, _numTokens uint256) constant returns(valid bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) ValidPosition(_prevID *big.Int, _nextID *big.Int, _voter common.Address, _numTokens *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.ValidPosition(&_CivilPLCRVotingContract.CallOpts, _prevID, _nextID, _voter, _numTokens)
}

// ValidPosition is a free data retrieval call binding the contract method 0x819b0293.
//
// Solidity: function validPosition(_prevID uint256, _nextID uint256, _voter address, _numTokens uint256) constant returns(valid bool)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) ValidPosition(_prevID *big.Int, _nextID *big.Int, _voter common.Address, _numTokens *big.Int) (bool, error) {
	return _CivilPLCRVotingContract.Contract.ValidPosition(&_CivilPLCRVotingContract.CallOpts, _prevID, _nextID, _voter, _numTokens)
}

// VoteTokenBalance is a free data retrieval call binding the contract method 0x3b930294.
//
// Solidity: function voteTokenBalance( address) constant returns(uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCaller) VoteTokenBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilPLCRVotingContract.contract.Call(opts, out, "voteTokenBalance", arg0)
	return *ret0, err
}

// VoteTokenBalance is a free data retrieval call binding the contract method 0x3b930294.
//
// Solidity: function voteTokenBalance( address) constant returns(uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) VoteTokenBalance(arg0 common.Address) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.VoteTokenBalance(&_CivilPLCRVotingContract.CallOpts, arg0)
}

// VoteTokenBalance is a free data retrieval call binding the contract method 0x3b930294.
//
// Solidity: function voteTokenBalance( address) constant returns(uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractCallerSession) VoteTokenBalance(arg0 common.Address) (*big.Int, error) {
	return _CivilPLCRVotingContract.Contract.VoteTokenBalance(&_CivilPLCRVotingContract.CallOpts, arg0)
}

// CommitVote is a paid mutator transaction binding the contract method 0x6cbf9c5e.
//
// Solidity: function commitVote(_pollID uint256, _secretHash bytes32, _numTokens uint256, _prevPollID uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactor) CommitVote(opts *bind.TransactOpts, _pollID *big.Int, _secretHash [32]byte, _numTokens *big.Int, _prevPollID *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.contract.Transact(opts, "commitVote", _pollID, _secretHash, _numTokens, _prevPollID)
}

// CommitVote is a paid mutator transaction binding the contract method 0x6cbf9c5e.
//
// Solidity: function commitVote(_pollID uint256, _secretHash bytes32, _numTokens uint256, _prevPollID uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) CommitVote(_pollID *big.Int, _secretHash [32]byte, _numTokens *big.Int, _prevPollID *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.CommitVote(&_CivilPLCRVotingContract.TransactOpts, _pollID, _secretHash, _numTokens, _prevPollID)
}

// CommitVote is a paid mutator transaction binding the contract method 0x6cbf9c5e.
//
// Solidity: function commitVote(_pollID uint256, _secretHash bytes32, _numTokens uint256, _prevPollID uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactorSession) CommitVote(_pollID *big.Int, _secretHash [32]byte, _numTokens *big.Int, _prevPollID *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.CommitVote(&_CivilPLCRVotingContract.TransactOpts, _pollID, _secretHash, _numTokens, _prevPollID)
}

// CommitVotes is a paid mutator transaction binding the contract method 0x3ec36b99.
//
// Solidity: function commitVotes(_pollIDs uint256[], _secretHashes bytes32[], _numsTokens uint256[], _prevPollIDs uint256[]) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactor) CommitVotes(opts *bind.TransactOpts, _pollIDs []*big.Int, _secretHashes [][32]byte, _numsTokens []*big.Int, _prevPollIDs []*big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.contract.Transact(opts, "commitVotes", _pollIDs, _secretHashes, _numsTokens, _prevPollIDs)
}

// CommitVotes is a paid mutator transaction binding the contract method 0x3ec36b99.
//
// Solidity: function commitVotes(_pollIDs uint256[], _secretHashes bytes32[], _numsTokens uint256[], _prevPollIDs uint256[]) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) CommitVotes(_pollIDs []*big.Int, _secretHashes [][32]byte, _numsTokens []*big.Int, _prevPollIDs []*big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.CommitVotes(&_CivilPLCRVotingContract.TransactOpts, _pollIDs, _secretHashes, _numsTokens, _prevPollIDs)
}

// CommitVotes is a paid mutator transaction binding the contract method 0x3ec36b99.
//
// Solidity: function commitVotes(_pollIDs uint256[], _secretHashes bytes32[], _numsTokens uint256[], _prevPollIDs uint256[]) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactorSession) CommitVotes(_pollIDs []*big.Int, _secretHashes [][32]byte, _numsTokens []*big.Int, _prevPollIDs []*big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.CommitVotes(&_CivilPLCRVotingContract.TransactOpts, _pollIDs, _secretHashes, _numsTokens, _prevPollIDs)
}

// RequestVotingRights is a paid mutator transaction binding the contract method 0xa25236fe.
//
// Solidity: function requestVotingRights(_numTokens uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactor) RequestVotingRights(opts *bind.TransactOpts, _numTokens *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.contract.Transact(opts, "requestVotingRights", _numTokens)
}

// RequestVotingRights is a paid mutator transaction binding the contract method 0xa25236fe.
//
// Solidity: function requestVotingRights(_numTokens uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) RequestVotingRights(_numTokens *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.RequestVotingRights(&_CivilPLCRVotingContract.TransactOpts, _numTokens)
}

// RequestVotingRights is a paid mutator transaction binding the contract method 0xa25236fe.
//
// Solidity: function requestVotingRights(_numTokens uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactorSession) RequestVotingRights(_numTokens *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.RequestVotingRights(&_CivilPLCRVotingContract.TransactOpts, _numTokens)
}

// RescueTokens is a paid mutator transaction binding the contract method 0x97603560.
//
// Solidity: function rescueTokens(_pollID uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactor) RescueTokens(opts *bind.TransactOpts, _pollID *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.contract.Transact(opts, "rescueTokens", _pollID)
}

// RescueTokens is a paid mutator transaction binding the contract method 0x97603560.
//
// Solidity: function rescueTokens(_pollID uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) RescueTokens(_pollID *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.RescueTokens(&_CivilPLCRVotingContract.TransactOpts, _pollID)
}

// RescueTokens is a paid mutator transaction binding the contract method 0x97603560.
//
// Solidity: function rescueTokens(_pollID uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactorSession) RescueTokens(_pollID *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.RescueTokens(&_CivilPLCRVotingContract.TransactOpts, _pollID)
}

// RescueTokensInMultiplePolls is a paid mutator transaction binding the contract method 0xbb11ed7e.
//
// Solidity: function rescueTokensInMultiplePolls(_pollIDs uint256[]) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactor) RescueTokensInMultiplePolls(opts *bind.TransactOpts, _pollIDs []*big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.contract.Transact(opts, "rescueTokensInMultiplePolls", _pollIDs)
}

// RescueTokensInMultiplePolls is a paid mutator transaction binding the contract method 0xbb11ed7e.
//
// Solidity: function rescueTokensInMultiplePolls(_pollIDs uint256[]) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) RescueTokensInMultiplePolls(_pollIDs []*big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.RescueTokensInMultiplePolls(&_CivilPLCRVotingContract.TransactOpts, _pollIDs)
}

// RescueTokensInMultiplePolls is a paid mutator transaction binding the contract method 0xbb11ed7e.
//
// Solidity: function rescueTokensInMultiplePolls(_pollIDs uint256[]) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactorSession) RescueTokensInMultiplePolls(_pollIDs []*big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.RescueTokensInMultiplePolls(&_CivilPLCRVotingContract.TransactOpts, _pollIDs)
}

// RevealVote is a paid mutator transaction binding the contract method 0xb11d8bb8.
//
// Solidity: function revealVote(_pollID uint256, _voteOption uint256, _salt uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactor) RevealVote(opts *bind.TransactOpts, _pollID *big.Int, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.contract.Transact(opts, "revealVote", _pollID, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0xb11d8bb8.
//
// Solidity: function revealVote(_pollID uint256, _voteOption uint256, _salt uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) RevealVote(_pollID *big.Int, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.RevealVote(&_CivilPLCRVotingContract.TransactOpts, _pollID, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0xb11d8bb8.
//
// Solidity: function revealVote(_pollID uint256, _voteOption uint256, _salt uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactorSession) RevealVote(_pollID *big.Int, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.RevealVote(&_CivilPLCRVotingContract.TransactOpts, _pollID, _voteOption, _salt)
}

// RevealVotes is a paid mutator transaction binding the contract method 0x8090f92e.
//
// Solidity: function revealVotes(_pollIDs uint256[], _voteOptions uint256[], _salts uint256[]) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactor) RevealVotes(opts *bind.TransactOpts, _pollIDs []*big.Int, _voteOptions []*big.Int, _salts []*big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.contract.Transact(opts, "revealVotes", _pollIDs, _voteOptions, _salts)
}

// RevealVotes is a paid mutator transaction binding the contract method 0x8090f92e.
//
// Solidity: function revealVotes(_pollIDs uint256[], _voteOptions uint256[], _salts uint256[]) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) RevealVotes(_pollIDs []*big.Int, _voteOptions []*big.Int, _salts []*big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.RevealVotes(&_CivilPLCRVotingContract.TransactOpts, _pollIDs, _voteOptions, _salts)
}

// RevealVotes is a paid mutator transaction binding the contract method 0x8090f92e.
//
// Solidity: function revealVotes(_pollIDs uint256[], _voteOptions uint256[], _salts uint256[]) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactorSession) RevealVotes(_pollIDs []*big.Int, _voteOptions []*big.Int, _salts []*big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.RevealVotes(&_CivilPLCRVotingContract.TransactOpts, _pollIDs, _voteOptions, _salts)
}

// StartPoll is a paid mutator transaction binding the contract method 0x32ed3d60.
//
// Solidity: function startPoll(_voteQuorum uint256, _commitDuration uint256, _revealDuration uint256) returns(pollID uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactor) StartPoll(opts *bind.TransactOpts, _voteQuorum *big.Int, _commitDuration *big.Int, _revealDuration *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.contract.Transact(opts, "startPoll", _voteQuorum, _commitDuration, _revealDuration)
}

// StartPoll is a paid mutator transaction binding the contract method 0x32ed3d60.
//
// Solidity: function startPoll(_voteQuorum uint256, _commitDuration uint256, _revealDuration uint256) returns(pollID uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) StartPoll(_voteQuorum *big.Int, _commitDuration *big.Int, _revealDuration *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.StartPoll(&_CivilPLCRVotingContract.TransactOpts, _voteQuorum, _commitDuration, _revealDuration)
}

// StartPoll is a paid mutator transaction binding the contract method 0x32ed3d60.
//
// Solidity: function startPoll(_voteQuorum uint256, _commitDuration uint256, _revealDuration uint256) returns(pollID uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactorSession) StartPoll(_voteQuorum *big.Int, _commitDuration *big.Int, _revealDuration *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.StartPoll(&_CivilPLCRVotingContract.TransactOpts, _voteQuorum, _commitDuration, _revealDuration)
}

// WithdrawVotingRights is a paid mutator transaction binding the contract method 0xe7b1d43c.
//
// Solidity: function withdrawVotingRights(_numTokens uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactor) WithdrawVotingRights(opts *bind.TransactOpts, _numTokens *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.contract.Transact(opts, "withdrawVotingRights", _numTokens)
}

// WithdrawVotingRights is a paid mutator transaction binding the contract method 0xe7b1d43c.
//
// Solidity: function withdrawVotingRights(_numTokens uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractSession) WithdrawVotingRights(_numTokens *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.WithdrawVotingRights(&_CivilPLCRVotingContract.TransactOpts, _numTokens)
}

// WithdrawVotingRights is a paid mutator transaction binding the contract method 0xe7b1d43c.
//
// Solidity: function withdrawVotingRights(_numTokens uint256) returns()
func (_CivilPLCRVotingContract *CivilPLCRVotingContractTransactorSession) WithdrawVotingRights(_numTokens *big.Int) (*types.Transaction, error) {
	return _CivilPLCRVotingContract.Contract.WithdrawVotingRights(&_CivilPLCRVotingContract.TransactOpts, _numTokens)
}

// CivilPLCRVotingContractPollCreatedIterator is returned from FilterPollCreated and is used to iterate over the raw logs and unpacked data for PollCreated events raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractPollCreatedIterator struct {
	Event *CivilPLCRVotingContractPollCreated // Event containing the contract specifics and raw log

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
func (it *CivilPLCRVotingContractPollCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilPLCRVotingContractPollCreated)
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
		it.Event = new(CivilPLCRVotingContractPollCreated)
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
func (it *CivilPLCRVotingContractPollCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilPLCRVotingContractPollCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilPLCRVotingContractPollCreated represents a PollCreated event raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractPollCreated struct {
	VoteQuorum    *big.Int
	CommitEndDate *big.Int
	RevealEndDate *big.Int
	PollID        *big.Int
	Creator       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPollCreated is a free log retrieval operation binding the contract event 0x404f1f1c229d9eb2a949e7584da6ffde9d059ef2169f487ca815434cce0640d0.
//
// Solidity: e _PollCreated(voteQuorum uint256, commitEndDate uint256, revealEndDate uint256, pollID indexed uint256, creator indexed address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) FilterPollCreated(opts *bind.FilterOpts, pollID []*big.Int, creator []common.Address) (*CivilPLCRVotingContractPollCreatedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _CivilPLCRVotingContract.contract.FilterLogs(opts, "_PollCreated", pollIDRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &CivilPLCRVotingContractPollCreatedIterator{contract: _CivilPLCRVotingContract.contract, event: "_PollCreated", logs: logs, sub: sub}, nil
}

// WatchPollCreated is a free log subscription operation binding the contract event 0x404f1f1c229d9eb2a949e7584da6ffde9d059ef2169f487ca815434cce0640d0.
//
// Solidity: e _PollCreated(voteQuorum uint256, commitEndDate uint256, revealEndDate uint256, pollID indexed uint256, creator indexed address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) WatchPollCreated(opts *bind.WatchOpts, sink chan<- *CivilPLCRVotingContractPollCreated, pollID []*big.Int, creator []common.Address) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _CivilPLCRVotingContract.contract.WatchLogs(opts, "_PollCreated", pollIDRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilPLCRVotingContractPollCreated)
				if err := _CivilPLCRVotingContract.contract.UnpackLog(event, "_PollCreated", log); err != nil {
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

// CivilPLCRVotingContractTokensRescuedIterator is returned from FilterTokensRescued and is used to iterate over the raw logs and unpacked data for TokensRescued events raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractTokensRescuedIterator struct {
	Event *CivilPLCRVotingContractTokensRescued // Event containing the contract specifics and raw log

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
func (it *CivilPLCRVotingContractTokensRescuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilPLCRVotingContractTokensRescued)
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
		it.Event = new(CivilPLCRVotingContractTokensRescued)
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
func (it *CivilPLCRVotingContractTokensRescuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilPLCRVotingContractTokensRescuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilPLCRVotingContractTokensRescued represents a TokensRescued event raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractTokensRescued struct {
	PollID *big.Int
	Voter  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensRescued is a free log retrieval operation binding the contract event 0x402507661c8c8cb90e0a796450b8bdd28b6c516f05279c0cd29e84c344e1699a.
//
// Solidity: e _TokensRescued(pollID indexed uint256, voter indexed address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) FilterTokensRescued(opts *bind.FilterOpts, pollID []*big.Int, voter []common.Address) (*CivilPLCRVotingContractTokensRescuedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _CivilPLCRVotingContract.contract.FilterLogs(opts, "_TokensRescued", pollIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &CivilPLCRVotingContractTokensRescuedIterator{contract: _CivilPLCRVotingContract.contract, event: "_TokensRescued", logs: logs, sub: sub}, nil
}

// WatchTokensRescued is a free log subscription operation binding the contract event 0x402507661c8c8cb90e0a796450b8bdd28b6c516f05279c0cd29e84c344e1699a.
//
// Solidity: e _TokensRescued(pollID indexed uint256, voter indexed address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) WatchTokensRescued(opts *bind.WatchOpts, sink chan<- *CivilPLCRVotingContractTokensRescued, pollID []*big.Int, voter []common.Address) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _CivilPLCRVotingContract.contract.WatchLogs(opts, "_TokensRescued", pollIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilPLCRVotingContractTokensRescued)
				if err := _CivilPLCRVotingContract.contract.UnpackLog(event, "_TokensRescued", log); err != nil {
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

// CivilPLCRVotingContractVoteCommittedIterator is returned from FilterVoteCommitted and is used to iterate over the raw logs and unpacked data for VoteCommitted events raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractVoteCommittedIterator struct {
	Event *CivilPLCRVotingContractVoteCommitted // Event containing the contract specifics and raw log

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
func (it *CivilPLCRVotingContractVoteCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilPLCRVotingContractVoteCommitted)
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
		it.Event = new(CivilPLCRVotingContractVoteCommitted)
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
func (it *CivilPLCRVotingContractVoteCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilPLCRVotingContractVoteCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilPLCRVotingContractVoteCommitted represents a VoteCommitted event raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractVoteCommitted struct {
	PollID    *big.Int
	NumTokens *big.Int
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoteCommitted is a free log retrieval operation binding the contract event 0xea7979e4280d7e6bffc1c7d83a1ac99f16d02ecc14465ce41016226783b663d7.
//
// Solidity: e _VoteCommitted(pollID indexed uint256, numTokens uint256, voter indexed address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) FilterVoteCommitted(opts *bind.FilterOpts, pollID []*big.Int, voter []common.Address) (*CivilPLCRVotingContractVoteCommittedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _CivilPLCRVotingContract.contract.FilterLogs(opts, "_VoteCommitted", pollIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &CivilPLCRVotingContractVoteCommittedIterator{contract: _CivilPLCRVotingContract.contract, event: "_VoteCommitted", logs: logs, sub: sub}, nil
}

// WatchVoteCommitted is a free log subscription operation binding the contract event 0xea7979e4280d7e6bffc1c7d83a1ac99f16d02ecc14465ce41016226783b663d7.
//
// Solidity: e _VoteCommitted(pollID indexed uint256, numTokens uint256, voter indexed address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) WatchVoteCommitted(opts *bind.WatchOpts, sink chan<- *CivilPLCRVotingContractVoteCommitted, pollID []*big.Int, voter []common.Address) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _CivilPLCRVotingContract.contract.WatchLogs(opts, "_VoteCommitted", pollIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilPLCRVotingContractVoteCommitted)
				if err := _CivilPLCRVotingContract.contract.UnpackLog(event, "_VoteCommitted", log); err != nil {
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

// CivilPLCRVotingContractVoteRevealedIterator is returned from FilterVoteRevealed and is used to iterate over the raw logs and unpacked data for VoteRevealed events raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractVoteRevealedIterator struct {
	Event *CivilPLCRVotingContractVoteRevealed // Event containing the contract specifics and raw log

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
func (it *CivilPLCRVotingContractVoteRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilPLCRVotingContractVoteRevealed)
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
		it.Event = new(CivilPLCRVotingContractVoteRevealed)
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
func (it *CivilPLCRVotingContractVoteRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilPLCRVotingContractVoteRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilPLCRVotingContractVoteRevealed represents a VoteRevealed event raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractVoteRevealed struct {
	PollID       *big.Int
	NumTokens    *big.Int
	VotesFor     *big.Int
	VotesAgainst *big.Int
	Choice       *big.Int
	Voter        common.Address
	Salt         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVoteRevealed is a free log retrieval operation binding the contract event 0x9b19aaec524fad29c0ced9b9973a15e3045d7c3be156d71394ab40f0d5f119ff.
//
// Solidity: e _VoteRevealed(pollID indexed uint256, numTokens uint256, votesFor uint256, votesAgainst uint256, choice indexed uint256, voter indexed address, salt uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) FilterVoteRevealed(opts *bind.FilterOpts, pollID []*big.Int, choice []*big.Int, voter []common.Address) (*CivilPLCRVotingContractVoteRevealedIterator, error) {

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

	logs, sub, err := _CivilPLCRVotingContract.contract.FilterLogs(opts, "_VoteRevealed", pollIDRule, choiceRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &CivilPLCRVotingContractVoteRevealedIterator{contract: _CivilPLCRVotingContract.contract, event: "_VoteRevealed", logs: logs, sub: sub}, nil
}

// WatchVoteRevealed is a free log subscription operation binding the contract event 0x9b19aaec524fad29c0ced9b9973a15e3045d7c3be156d71394ab40f0d5f119ff.
//
// Solidity: e _VoteRevealed(pollID indexed uint256, numTokens uint256, votesFor uint256, votesAgainst uint256, choice indexed uint256, voter indexed address, salt uint256)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) WatchVoteRevealed(opts *bind.WatchOpts, sink chan<- *CivilPLCRVotingContractVoteRevealed, pollID []*big.Int, choice []*big.Int, voter []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CivilPLCRVotingContract.contract.WatchLogs(opts, "_VoteRevealed", pollIDRule, choiceRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilPLCRVotingContractVoteRevealed)
				if err := _CivilPLCRVotingContract.contract.UnpackLog(event, "_VoteRevealed", log); err != nil {
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

// CivilPLCRVotingContractVotingRightsGrantedIterator is returned from FilterVotingRightsGranted and is used to iterate over the raw logs and unpacked data for VotingRightsGranted events raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractVotingRightsGrantedIterator struct {
	Event *CivilPLCRVotingContractVotingRightsGranted // Event containing the contract specifics and raw log

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
func (it *CivilPLCRVotingContractVotingRightsGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilPLCRVotingContractVotingRightsGranted)
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
		it.Event = new(CivilPLCRVotingContractVotingRightsGranted)
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
func (it *CivilPLCRVotingContractVotingRightsGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilPLCRVotingContractVotingRightsGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilPLCRVotingContractVotingRightsGranted represents a VotingRightsGranted event raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractVotingRightsGranted struct {
	NumTokens *big.Int
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVotingRightsGranted is a free log retrieval operation binding the contract event 0xf7aaf024511d9982df8cd0d437c71c30106e6848cd1ba3d288d7a9c0e276aeda.
//
// Solidity: e _VotingRightsGranted(numTokens uint256, voter indexed address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) FilterVotingRightsGranted(opts *bind.FilterOpts, voter []common.Address) (*CivilPLCRVotingContractVotingRightsGrantedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _CivilPLCRVotingContract.contract.FilterLogs(opts, "_VotingRightsGranted", voterRule)
	if err != nil {
		return nil, err
	}
	return &CivilPLCRVotingContractVotingRightsGrantedIterator{contract: _CivilPLCRVotingContract.contract, event: "_VotingRightsGranted", logs: logs, sub: sub}, nil
}

// WatchVotingRightsGranted is a free log subscription operation binding the contract event 0xf7aaf024511d9982df8cd0d437c71c30106e6848cd1ba3d288d7a9c0e276aeda.
//
// Solidity: e _VotingRightsGranted(numTokens uint256, voter indexed address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) WatchVotingRightsGranted(opts *bind.WatchOpts, sink chan<- *CivilPLCRVotingContractVotingRightsGranted, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _CivilPLCRVotingContract.contract.WatchLogs(opts, "_VotingRightsGranted", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilPLCRVotingContractVotingRightsGranted)
				if err := _CivilPLCRVotingContract.contract.UnpackLog(event, "_VotingRightsGranted", log); err != nil {
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

// CivilPLCRVotingContractVotingRightsWithdrawnIterator is returned from FilterVotingRightsWithdrawn and is used to iterate over the raw logs and unpacked data for VotingRightsWithdrawn events raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractVotingRightsWithdrawnIterator struct {
	Event *CivilPLCRVotingContractVotingRightsWithdrawn // Event containing the contract specifics and raw log

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
func (it *CivilPLCRVotingContractVotingRightsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilPLCRVotingContractVotingRightsWithdrawn)
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
		it.Event = new(CivilPLCRVotingContractVotingRightsWithdrawn)
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
func (it *CivilPLCRVotingContractVotingRightsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilPLCRVotingContractVotingRightsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilPLCRVotingContractVotingRightsWithdrawn represents a VotingRightsWithdrawn event raised by the CivilPLCRVotingContract contract.
type CivilPLCRVotingContractVotingRightsWithdrawn struct {
	NumTokens *big.Int
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVotingRightsWithdrawn is a free log retrieval operation binding the contract event 0xfaeb7dbb9992397d26ea1944efd40c80b40f702faf69b46c67ad10aba68ccb79.
//
// Solidity: e _VotingRightsWithdrawn(numTokens uint256, voter indexed address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) FilterVotingRightsWithdrawn(opts *bind.FilterOpts, voter []common.Address) (*CivilPLCRVotingContractVotingRightsWithdrawnIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _CivilPLCRVotingContract.contract.FilterLogs(opts, "_VotingRightsWithdrawn", voterRule)
	if err != nil {
		return nil, err
	}
	return &CivilPLCRVotingContractVotingRightsWithdrawnIterator{contract: _CivilPLCRVotingContract.contract, event: "_VotingRightsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchVotingRightsWithdrawn is a free log subscription operation binding the contract event 0xfaeb7dbb9992397d26ea1944efd40c80b40f702faf69b46c67ad10aba68ccb79.
//
// Solidity: e _VotingRightsWithdrawn(numTokens uint256, voter indexed address)
func (_CivilPLCRVotingContract *CivilPLCRVotingContractFilterer) WatchVotingRightsWithdrawn(opts *bind.WatchOpts, sink chan<- *CivilPLCRVotingContractVotingRightsWithdrawn, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _CivilPLCRVotingContract.contract.WatchLogs(opts, "_VotingRightsWithdrawn", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilPLCRVotingContractVotingRightsWithdrawn)
				if err := _CivilPLCRVotingContract.contract.UnpackLog(event, "_VotingRightsWithdrawn", log); err != nil {
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
