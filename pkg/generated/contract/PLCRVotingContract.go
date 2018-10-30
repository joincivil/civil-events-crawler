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
const PLCRVotingContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_POLL_NONCE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"voteTokenBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pollMap\",\"outputs\":[{\"name\":\"commitEndDate\",\"type\":\"uint256\"},{\"name\":\"revealEndDate\",\"type\":\"uint256\"},{\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"name\":\"votesFor\",\"type\":\"uint256\"},{\"name\":\"votesAgainst\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pollNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_VoteCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesFor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesAgainst\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"choice\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"_VoteRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commitEndDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealEndDate\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"_PollCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_VotingRightsGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_VotingRightsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_TokensRescued\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_numTokens\",\"type\":\"uint256\"}],\"name\":\"requestVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_numTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollIDs\",\"type\":\"uint256[]\"}],\"name\":\"rescueTokensInMultiplePolls\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"},{\"name\":\"_numTokens\",\"type\":\"uint256\"},{\"name\":\"_prevPollID\",\"type\":\"uint256\"}],\"name\":\"commitVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollIDs\",\"type\":\"uint256[]\"},{\"name\":\"_secretHashes\",\"type\":\"bytes32[]\"},{\"name\":\"_numsTokens\",\"type\":\"uint256[]\"},{\"name\":\"_prevPollIDs\",\"type\":\"uint256[]\"}],\"name\":\"commitVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prevID\",\"type\":\"uint256\"},{\"name\":\"_nextID\",\"type\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_numTokens\",\"type\":\"uint256\"}],\"name\":\"validPosition\",\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollIDs\",\"type\":\"uint256[]\"},{\"name\":\"_voteOptions\",\"type\":\"uint256[]\"},{\"name\":\"_salts\",\"type\":\"uint256[]\"}],\"name\":\"revealVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"getNumPassingTokens\",\"outputs\":[{\"name\":\"correctVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteQuorum\",\"type\":\"uint256\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"}],\"name\":\"startPoll\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"isPassed\",\"outputs\":[{\"name\":\"passed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getTotalNumberOfTokensForWinningOption\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"pollEnded\",\"outputs\":[{\"name\":\"ended\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"commitPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"revealPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"didCommit\",\"outputs\":[{\"name\":\"committed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"didReveal\",\"outputs\":[{\"name\":\"revealed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"pollExists\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getCommitHash\",\"outputs\":[{\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getNumTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getLastNode\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getLockedTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_numTokens\",\"type\":\"uint256\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getInsertPointForNumTokens\",\"outputs\":[{\"name\":\"prevNode\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_terminationDate\",\"type\":\"uint256\"}],\"name\":\"isExpired\",\"outputs\":[{\"name\":\"expired\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"attrUUID\",\"outputs\":[{\"name\":\"UUID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PLCRVotingContractBin is the compiled bytecode used for deploying new contracts.
const PLCRVotingContractBin = `0x608060405234801561001057600080fd5b50604051602080611c7e8339810160405251600160a060020a038116151561003757600080fd5b60058054600160a060020a03909216600160a060020a031990921691909117905560008055611c138061006b6000396000f3006080604052600436106101715763ffffffff60e060020a600035041663053e71a681146101765780632173a10f146101a05780632c052031146101b557806332ed3d60146101dc5780633b930294146101fa5780633ec36b991461021b578063427fa1d214610261578063441c77c01461028257806349403183146102ae5780636148fed5146102c65780636b2d95d4146103095780636cbf9c5e1461032a5780637f97e8361461034b5780638090f92e1461036f578063819b0293146103a757806388d21ff3146103d157806397508f36146103e957806397603560146103fe578063a1103f3714610416578063a25236fe1461043a578063a4439dc514610452578063aa7ca4641461046a578063b11d8bb81461048e578063b43bd069146104ac578063bb11ed7e146104d3578063d138209214610528578063d901402b1461054c578063d9548e5314610570578063e7b1d43c14610588578063ee684830146105a0578063fc0c546a146105b8575b600080fd5b34801561018257600080fd5b5061018e6004356105e9565b60408051918252519081900360200190f35b3480156101ac57600080fd5b5061018e61063e565b3480156101c157600080fd5b5061018e600160a060020a0360043516602435604435610643565b3480156101e857600080fd5b5061018e60043560243560443561080c565b34801561020657600080fd5b5061018e600160a060020a03600435166108e8565b34801561022757600080fd5b5061025f60246004803582810192908201359181358083019290820135916044358083019290820135916064359182019101356108fa565b005b34801561026d57600080fd5b5061018e600160a060020a036004351661099a565b34801561028e57600080fd5b5061029a600435610a54565b604080519115158252519081900360200190f35b3480156102ba57600080fd5b5061029a600435610a9f565b3480156102d257600080fd5b506102de600435610b1c565b6040805195865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561031557600080fd5b5061018e600160a060020a0360043516610b4a565b34801561033657600080fd5b5061025f600435602435604435606435610b5e565b34801561035757600080fd5b5061029a600160a060020a03600435166024356110d0565b34801561037b57600080fd5b5061025f6024600480358281019290820135918135808301929082013591604435918201910135611115565b3480156103b357600080fd5b5061029a600435602435600160a060020a0360443516606435611191565b3480156103dd57600080fd5b5061029a6004356111d2565b3480156103f557600080fd5b5061018e6111e7565b34801561040a57600080fd5b5061025f6004356111ed565b34801561042257600080fd5b5061018e600160a060020a0360043516602435611391565b34801561044657600080fd5b5061025f6004356113c8565b34801561045e57600080fd5b5061029a60043561155b565b34801561047657600080fd5b5061029a600160a060020a0360043516602435611589565b34801561049a57600080fd5b5061025f6004356024356044356115ce565b3480156104b857600080fd5b5061018e600160a060020a03600435166024356044356117d1565b3480156104df57600080fd5b506040805160206004803580820135838102808601850190965280855261025f9536959394602494938501929182918501908490808284375094975061187d9650505050505050565b34801561053457600080fd5b5061018e600160a060020a03600435166024356118b5565b34801561055857600080fd5b5061018e600160a060020a0360043516602435611987565b34801561057c57600080fd5b5061029a600435611a26565b34801561059457600080fd5b5061025f600435611a2b565b3480156105ac57600080fd5b5061029a600435611b57565b3480156105c457600080fd5b506105cd611b89565b60408051600160a060020a039092168252519081900360200190f35b60006105f482611b57565b15156105ff57600080fd5b61060882610a9f565b156106255750600081815260016020526040902060030154610639565b506000818152600160205260409020600401545b919050565b600081565b60008060006106518661099a565b915061065d86836118b5565b90505b81156107ff5761067086836118b5565b9050848111610742578382141561073a57600160a060020a03861660009081526003602090815260409182902082517f30fe0a0a000000000000000000000000000000000000000000000000000000008152600481019190915260248101859052915173__DLL___________________________________926330fe0a0a926044808301939192829003018186803b15801561070b57600080fd5b505af415801561071f573d6000803e3d6000fd5b505050506040513d602081101561073557600080fd5b505191505b819250610803565b600160a060020a03861660009081526003602090815260409182902082517f30fe0a0a000000000000000000000000000000000000000000000000000000008152600481019190915260248101859052915173__DLL___________________________________926330fe0a0a926044808301939192829003018186803b1580156107cc57600080fd5b505af41580156107e0573d6000803e3d6000fd5b505050506040513d60208110156107f657600080fd5b50519150610660565b8192505b50509392505050565b6000805460010181558080610827428663ffffffff611b9816565b9150610839828563ffffffff611b9816565b6040805160a08101825284815260208082018481528284018b8152600060608086018281526080870183815283548452600180885289852098518955955195880195909555925160028701559151600386015591516004909401939093555483518b8152918201879052818401859052925193945033937f404f1f1c229d9eb2a949e7584da6ffde9d059ef2169f487ca815434cce0640d0929181900390910190a35050600054949350505050565b60026020526000908152604090205481565b600087861461090857600080fd5b87841461091457600080fd5b87821461092057600080fd5b5060005b8781101561098f5761098789898381811061093b57fe5b90506020020135888884818110151561095057fe5b6020029190910135905087878581811061096657fe5b90506020020135868686818110151561097b57fe5b90506020020135610b5e565b600101610924565b505050505050505050565b600160a060020a038116600090815260036020908152604080832081517f30fe0a0a000000000000000000000000000000000000000000000000000000008152600481019190915260248101849052905173__DLL___________________________________926330fe0a0a9260448082019391829003018186803b158015610a2257600080fd5b505af4158015610a36573d6000803e3d6000fd5b505050506040513d6020811015610a4c57600080fd5b505192915050565b6000610a5f826111d2565b1515610a6a57600080fd5b60008281526001602081905260409091200154610a8690611a26565b158015610a995750610a978261155b565b155b92915050565b6000610aa9611bb7565b610ab283611b57565b1515610abd57600080fd5b5050600090815260016020818152604092839020835160a0810185528154815292810154918301919091526002810154928201839052600381015460608301819052600490910154608090920182905290810190910260649091021190565b6001602081905260009182526040909120805491810154600282015460038301546004909301549192909185565b6000610a9982610b598461099a565b6118b5565b6000806000610b6c8761155b565b1515610b7757600080fd5b33600090815260026020526040902054851115610bba5733600090815260026020526040902054610baf90869063ffffffff611ba516565b9250610bba836113c8565b33600090815260026020526040902054851115610bd657600080fd5b861515610be257600080fd5b851515610bee57600080fd5b831580610ca457503360009081526003602090815260409182902082517f366a5ba2000000000000000000000000000000000000000000000000000000008152600481019190915260248101879052915173__DLL___________________________________9263366a5ba2926044808301939192829003018186803b158015610c7757600080fd5b505af4158015610c8b573d6000803e3d6000fd5b505050506040513d6020811015610ca157600080fd5b50515b1515610caf57600080fd5b3360009081526003602090815260409182902082517f07d29ac9000000000000000000000000000000000000000000000000000000008152600481019190915260248101879052915173__DLL___________________________________926307d29ac9926044808301939192829003018186803b158015610d3057600080fd5b505af4158015610d44573d6000803e3d6000fd5b505050506040513d6020811015610d5a57600080fd5b5051915086821415610e16573360009081526003602090815260409182902082517f07d29ac90000000000000000000000000000000000000000000000000000000081526004810191909152602481018a9052915173__DLL___________________________________926307d29ac9926044808301939192829003018186803b158015610de757600080fd5b505af4158015610dfb573d6000803e3d6000fd5b505050506040513d6020811015610e1157600080fd5b505191505b610e2284833388611191565b1515610e2d57600080fd5b3360009081526003602052604080822081517f9735c51b000000000000000000000000000000000000000000000000000000008152600481019190915260248101879052604481018a905260648101859052905173__DLL___________________________________92639735c51b9260848082019391829003018186803b158015610eb857600080fd5b505af4158015610ecc573d6000803e3d6000fd5b50505050610eda3388611391565b604080517f977aa031000000000000000000000000000000000000000000000000000000008152600481810152602481018390526064810188905260806044820152600960848201527f6e756d546f6b656e73000000000000000000000000000000000000000000000060a4820152905191925073__AttributeStore________________________9163977aa0319160c480820192600092909190829003018186803b158015610f8a57600080fd5b505af4158015610f9e573d6000803e3d6000fd5b5050604080517f977aa03100000000000000000000000000000000000000000000000000000000815260048181015260248101859052606481018a905260806044820152600a60848201527f636f6d6d6974486173680000000000000000000000000000000000000000000060a4820152905173__AttributeStore________________________935063977aa031925060c4808301926000929190829003018186803b15801561104e57600080fd5b505af4158015611062573d6000803e3d6000fd5b5050506000888152600160208181526040808420338086526005909101835293819020805460ff1916909317909255815189815291519293508a927fea7979e4280d7e6bffc1c7d83a1ac99f16d02ecc14465ce41016226783b663d79281900390910190a350505050505050565b60006110db826111d2565b15156110e657600080fd5b506000908152600160209081526040808320600160a060020a0394909416835260059093019052205460ff1690565b600085841461112357600080fd5b85821461112f57600080fd5b5060005b858110156111885761118087878381811061114a57fe5b90506020020135868684818110151561115f57fe5b90506020020135858585818110151561117457fe5b905060200201356115ce565b600101611133565b50505050505050565b60008060006111a085886118b5565b84101591506111af85876118b5565b841115806111bb575085155b90508180156111c75750805b979650505050505050565b60008115801590610a99575050600054101590565b60005481565b6000818152600160208190526040909120015461120990611a26565b151561121457600080fd5b3360009081526003602090815260409182902082517f366a5ba2000000000000000000000000000000000000000000000000000000008152600481019190915260248101849052915173__DLL___________________________________9263366a5ba2926044808301939192829003018186803b15801561129557600080fd5b505af41580156112a9573d6000803e3d6000fd5b505050506040513d60208110156112bf57600080fd5b505115156112cc57600080fd5b3360009081526003602052604080822081517f6d900ed0000000000000000000000000000000000000000000000000000000008152600481019190915260248101849052905173__DLL___________________________________92636d900ed09260448082019391829003018186803b15801561134957600080fd5b505af415801561135d573d6000803e3d6000fd5b50506040513392508391507f402507661c8c8cb90e0a796450b8bdd28b6c516f05279c0cd29e84c344e1699a90600090a350565b604080516c01000000000000000000000000600160a060020a03851602815260148101839052905190819003603401902092915050565b600554604080517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015290518392600160a060020a0316916370a082319160248083019260209291908290030181600087803b15801561142d57600080fd5b505af1158015611441573d6000803e3d6000fd5b505050506040513d602081101561145757600080fd5b5051101561146457600080fd5b33600081815260026020908152604080832080548601905560055481517f23b872dd0000000000000000000000000000000000000000000000000000000081526004810195909552306024860152604485018690529051600160a060020a03909116936323b872dd9360648083019493928390030190829087803b1580156114eb57600080fd5b505af11580156114ff573d6000803e3d6000fd5b505050506040513d602081101561151557600080fd5b5051151561152257600080fd5b60408051828152905133917ff7aaf024511d9982df8cd0d437c71c30106e6848cd1ba3d288d7a9c0e276aeda919081900360200190a250565b6000611566826111d2565b151561157157600080fd5b600082815260016020526040902054610a9790611a26565b6000611594826111d2565b151561159f57600080fd5b506000908152600160209081526040808320600160a060020a0394909416835260069093019052205460ff1690565b60006115d984610a54565b15156115e457600080fd5b600084815260016020908152604080832033845260050190915290205460ff16151561160f57600080fd5b600084815260016020908152604080832033845260060190915290205460ff161561163957600080fd5b6116433385611987565b604080518581526020810185905281519081900390910190201461166657600080fd5b61167033856118b5565b905082600114156116975760008481526001602052604090206003018054820190556116af565b60008481526001602052604090206004018054820190555b3360009081526003602052604080822081517f6d900ed0000000000000000000000000000000000000000000000000000000008152600481019190915260248101879052905173__DLL___________________________________92636d900ed09260448082019391829003018186803b15801561172c57600080fd5b505af4158015611740573d6000803e3d6000fd5b505050600085815260016020818152604080842033808652600682018452828620805460ff191686179055948a9052928252600383015460049093015481518781529283019390935281810192909252606081018690529051919250859187917f9b19aaec524fad29c0ced9b9973a15e3045d7c3be156d71394ab40f0d5f119ff919081900360800190a450505050565b6000806000806117e086611b57565b15156117eb57600080fd5b6000868152600160209081526040808320600160a060020a038b16845260060190915290205460ff16151561181f57600080fd5b61182886610a9f565b611833576000611836565b60015b60ff16925082856040518083815260200182815260200192505050604051809103902091506118658787611987565b905081811461187357600080fd5b6111c787876118b5565b60005b81518110156118b1576118a9828281518110151561189a57fe5b906020019060200201516111ed565b600101611880565b5050565b600073__AttributeStore________________________6350389f5c60046118dd8686611391565b6040805160e060020a63ffffffff86160281526004810193909352602483019190915260606044830152600960648301527f6e756d546f6b656e73000000000000000000000000000000000000000000000060848301525160a4808301926020929190829003018186803b15801561195457600080fd5b505af4158015611968573d6000803e3d6000fd5b505050506040513d602081101561197e57600080fd5b50519392505050565b600073__AttributeStore________________________6350389f5c60046119af8686611391565b6040805160e060020a63ffffffff86160281526004810193909352602483019190915260606044830152600a60648301527f636f6d6d6974486173680000000000000000000000000000000000000000000060848301525160a4808301926020929190829003018186803b15801561195457600080fd5b421190565b6000611a55611a3933610b4a565b336000908152600260205260409020549063ffffffff611ba516565b905081811015611a6457600080fd5b3360008181526002602090815260408083208054879003905560055481517fa9059cbb0000000000000000000000000000000000000000000000000000000081526004810195909552602485018790529051600160a060020a039091169363a9059cbb9360448083019493928390030190829087803b158015611ae657600080fd5b505af1158015611afa573d6000803e3d6000fd5b505050506040513d6020811015611b1057600080fd5b50511515611b1d57600080fd5b60408051838152905133917ffaeb7dbb9992397d26ea1944efd40c80b40f702faf69b46c67ad10aba68ccb79919081900360200190a25050565b6000611b62826111d2565b1515611b6d57600080fd5b60008281526001602081905260409091200154610a9990611a26565b600554600160a060020a031681565b81810182811015610a9957fe5b600082821115611bb157fe5b50900390565b60a060405190810160405280600081526020016000815260200160008152602001600081526020016000815250905600a165627a7a7230582089ca3f2d5c8ac4b21b13500b9f4eee13e3c88783c6cd02c44310390988f37bb20029`

// DeployPLCRVotingContract deploys a new Ethereum contract, binding an instance of PLCRVotingContract to it.
func DeployPLCRVotingContract(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *PLCRVotingContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PLCRVotingContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PLCRVotingContractBin), backend, _token)
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

// CommitVotes is a paid mutator transaction binding the contract method 0x3ec36b99.
//
// Solidity: function commitVotes(_pollIDs uint256[], _secretHashes bytes32[], _numsTokens uint256[], _prevPollIDs uint256[]) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactor) CommitVotes(opts *bind.TransactOpts, _pollIDs []*big.Int, _secretHashes [][32]byte, _numsTokens []*big.Int, _prevPollIDs []*big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.contract.Transact(opts, "commitVotes", _pollIDs, _secretHashes, _numsTokens, _prevPollIDs)
}

// CommitVotes is a paid mutator transaction binding the contract method 0x3ec36b99.
//
// Solidity: function commitVotes(_pollIDs uint256[], _secretHashes bytes32[], _numsTokens uint256[], _prevPollIDs uint256[]) returns()
func (_PLCRVotingContract *PLCRVotingContractSession) CommitVotes(_pollIDs []*big.Int, _secretHashes [][32]byte, _numsTokens []*big.Int, _prevPollIDs []*big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.CommitVotes(&_PLCRVotingContract.TransactOpts, _pollIDs, _secretHashes, _numsTokens, _prevPollIDs)
}

// CommitVotes is a paid mutator transaction binding the contract method 0x3ec36b99.
//
// Solidity: function commitVotes(_pollIDs uint256[], _secretHashes bytes32[], _numsTokens uint256[], _prevPollIDs uint256[]) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactorSession) CommitVotes(_pollIDs []*big.Int, _secretHashes [][32]byte, _numsTokens []*big.Int, _prevPollIDs []*big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.CommitVotes(&_PLCRVotingContract.TransactOpts, _pollIDs, _secretHashes, _numsTokens, _prevPollIDs)
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

// RescueTokensInMultiplePolls is a paid mutator transaction binding the contract method 0xbb11ed7e.
//
// Solidity: function rescueTokensInMultiplePolls(_pollIDs uint256[]) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactor) RescueTokensInMultiplePolls(opts *bind.TransactOpts, _pollIDs []*big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.contract.Transact(opts, "rescueTokensInMultiplePolls", _pollIDs)
}

// RescueTokensInMultiplePolls is a paid mutator transaction binding the contract method 0xbb11ed7e.
//
// Solidity: function rescueTokensInMultiplePolls(_pollIDs uint256[]) returns()
func (_PLCRVotingContract *PLCRVotingContractSession) RescueTokensInMultiplePolls(_pollIDs []*big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.RescueTokensInMultiplePolls(&_PLCRVotingContract.TransactOpts, _pollIDs)
}

// RescueTokensInMultiplePolls is a paid mutator transaction binding the contract method 0xbb11ed7e.
//
// Solidity: function rescueTokensInMultiplePolls(_pollIDs uint256[]) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactorSession) RescueTokensInMultiplePolls(_pollIDs []*big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.RescueTokensInMultiplePolls(&_PLCRVotingContract.TransactOpts, _pollIDs)
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

// RevealVotes is a paid mutator transaction binding the contract method 0x8090f92e.
//
// Solidity: function revealVotes(_pollIDs uint256[], _voteOptions uint256[], _salts uint256[]) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactor) RevealVotes(opts *bind.TransactOpts, _pollIDs []*big.Int, _voteOptions []*big.Int, _salts []*big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.contract.Transact(opts, "revealVotes", _pollIDs, _voteOptions, _salts)
}

// RevealVotes is a paid mutator transaction binding the contract method 0x8090f92e.
//
// Solidity: function revealVotes(_pollIDs uint256[], _voteOptions uint256[], _salts uint256[]) returns()
func (_PLCRVotingContract *PLCRVotingContractSession) RevealVotes(_pollIDs []*big.Int, _voteOptions []*big.Int, _salts []*big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.RevealVotes(&_PLCRVotingContract.TransactOpts, _pollIDs, _voteOptions, _salts)
}

// RevealVotes is a paid mutator transaction binding the contract method 0x8090f92e.
//
// Solidity: function revealVotes(_pollIDs uint256[], _voteOptions uint256[], _salts uint256[]) returns()
func (_PLCRVotingContract *PLCRVotingContractTransactorSession) RevealVotes(_pollIDs []*big.Int, _voteOptions []*big.Int, _salts []*big.Int) (*types.Transaction, error) {
	return _PLCRVotingContract.Contract.RevealVotes(&_PLCRVotingContract.TransactOpts, _pollIDs, _voteOptions, _salts)
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
// Solidity: e _PollCreated(voteQuorum uint256, commitEndDate uint256, revealEndDate uint256, pollID indexed uint256, creator indexed address)
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
// Solidity: e _PollCreated(voteQuorum uint256, commitEndDate uint256, revealEndDate uint256, pollID indexed uint256, creator indexed address)
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
// Solidity: e _TokensRescued(pollID indexed uint256, voter indexed address)
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
// Solidity: e _TokensRescued(pollID indexed uint256, voter indexed address)
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
// Solidity: e _VoteCommitted(pollID indexed uint256, numTokens uint256, voter indexed address)
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
// Solidity: e _VoteCommitted(pollID indexed uint256, numTokens uint256, voter indexed address)
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
	Salt         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVoteRevealed is a free log retrieval operation binding the contract event 0x9b19aaec524fad29c0ced9b9973a15e3045d7c3be156d71394ab40f0d5f119ff.
//
// Solidity: e _VoteRevealed(pollID indexed uint256, numTokens uint256, votesFor uint256, votesAgainst uint256, choice indexed uint256, voter indexed address, salt uint256)
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

// WatchVoteRevealed is a free log subscription operation binding the contract event 0x9b19aaec524fad29c0ced9b9973a15e3045d7c3be156d71394ab40f0d5f119ff.
//
// Solidity: e _VoteRevealed(pollID indexed uint256, numTokens uint256, votesFor uint256, votesAgainst uint256, choice indexed uint256, voter indexed address, salt uint256)
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
// Solidity: e _VotingRightsGranted(numTokens uint256, voter indexed address)
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
// Solidity: e _VotingRightsGranted(numTokens uint256, voter indexed address)
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
// Solidity: e _VotingRightsWithdrawn(numTokens uint256, voter indexed address)
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
// Solidity: e _VotingRightsWithdrawn(numTokens uint256, voter indexed address)
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
