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

// CivilTokenControllerContractABI is the input ABI used to generate the binding from.
const CivilTokenControllerContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"SUCCESS_CODE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addToCivilians\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeFromUnlocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"newsroomMultisigList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"managerAddress\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeFromCore\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addToUnlocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MUST_BE_A_CIVILIAN_CODE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeFromVerified\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addToNewsroomMultisigs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addToCore\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MUST_BE_A_CIVILIAN_ERROR\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"verifiedList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockedList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"coreList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeFromCivilians\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"managerAddress\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addToStorefront\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeFromStorefront\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MUST_BE_VERIFIED_ERROR\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MUST_BE_VERIFIED_CODE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isManager\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MUST_BE_UNLOCKED_CODE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SUCCESS_MESSAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeFromNewsroomMultisigs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"checkProofOfUse\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addToVerified\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"civilianList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"storefrontList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MUST_BE_UNLOCKED_ERROR\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"managerAddress\",\"type\":\"address\"}],\"name\":\"checkManagerStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"managers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"detectTransferRestriction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"restrictionCode\",\"type\":\"uint8\"}],\"name\":\"messageForTransferRestriction\",\"outputs\":[{\"name\":\"message\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"onRequestVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CivilTokenControllerContractBin is the compiled bytecode used for deploying new contracts.
const CivilTokenControllerContractBin = `0x60806040523480156200001157600080fd5b5060008054600160a060020a03191633178155604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020820190815291517f6ef3054c0000000000000000000000000000000000000000000000000000000081526008600482018181526024830186905260606044840190815284516064850152845173__MessagesAndCodes______________________97636ef3054c97949690959094608490910191808383895b83811015620000e7578181015183820152602001620000cd565b50505050905090810190601f168015620001155780820380516001836020036101000a031916815260200191505b5094505050505060206040518083038186803b1580156200013557600080fd5b505af41580156200014a573d6000803e3d6000fd5b505050506040513d60208110156200016157600080fd5b5050604080518082018252601281527f4d5553545f42455f415f434956494c49414e00000000000000000000000000006020820190815291517f6ef3054c00000000000000000000000000000000000000000000000000000000815260086004820181815260016024840181905260606044850190815285516064860152855173__MessagesAndCodes______________________97636ef3054c979596939594936084019180838360005b83811015620002275781810151838201526020016200020d565b50505050905090810190601f168015620002555780820380516001836020036101000a031916815260200191505b5094505050505060206040518083038186803b1580156200027557600080fd5b505af41580156200028a573d6000803e3d6000fd5b505050506040513d6020811015620002a157600080fd5b5050604080518082018252601081527f4d5553545f42455f554e4c4f434b4544000000000000000000000000000000006020820190815291517f6ef3054c00000000000000000000000000000000000000000000000000000000815260086004820181815260026024840181905260606044850190815285516064860152855173__MessagesAndCodes______________________97636ef3054c979596939594936084019180838360005b83811015620003675781810151838201526020016200034d565b50505050905090810190601f168015620003955780820380516001836020036101000a031916815260200191505b5094505050505060206040518083038186803b158015620003b557600080fd5b505af4158015620003ca573d6000803e3d6000fd5b505050506040513d6020811015620003e157600080fd5b5050604080518082018252601081527f4d5553545f42455f5645524946494544000000000000000000000000000000006020820190815291517f6ef3054c00000000000000000000000000000000000000000000000000000000815260086004820181815260036024840181905260606044850190815285516064860152855173__MessagesAndCodes______________________97636ef3054c979596939594936084019180838360005b83811015620004a75781810151838201526020016200048d565b50505050905090810190601f168015620004d55780820380516001836020036101000a031916815260200191505b5094505050505060206040518083038186803b158015620004f557600080fd5b505af41580156200050a573d6000803e3d6000fd5b505050506040513d60208110156200052157600080fd5b505061143480620005336000396000f3006080604052600436106101d75763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630e969a0581146101dc5780631299090a1461020757806314862ea51461022a578063166542b31461024b5780632d06177a146102805780633c421424146102a15780633f16b282146102c2578063439b650b146102e35780634b3d1485146102f8578063607c60bb14610319578063715018a61461033a578063725248731461034f57806373c601821461037357806378a21ddb146103945780637f4ab1dd1461041e57806381601496146104395780638bdeb4521461045a5780638da5cb5b1461047b5780639685cc57146104ac578063ab3d0c7a146104cd578063ac18de43146104ee578063ac1f38531461050f578063adbb916014610530578063bdcadb3b14610551578063c0e9794a14610566578063c56a3e881461057b578063ca3aaa9a14610590578063d4ce1415146105a5578063e7984d17146105cf578063e79a4fd4146105e4578063e99f29a414610605578063ee37c29f14610626578063ee56f4fa14610647578063f0fbca0614610668578063f198391814610689578063f281e7d11461069e578063f2fde38b146106bf578063fdff9b4d146106e0575b600080fd5b3480156101e857600080fd5b506101f1610701565b6040805160ff9092168252519081900360200190f35b34801561021357600080fd5b50610228600160a060020a0360043516610706565b005b34801561023657600080fd5b50610228600160a060020a03600435166107a0565b34801561025757600080fd5b5061026c600160a060020a0360043516610837565b604080519115158252519081900360200190f35b34801561028c57600080fd5b50610228600160a060020a036004351661084c565b3480156102ad57600080fd5b50610228600160a060020a036004351661088a565b3480156102ce57600080fd5b50610228600160a060020a0360043516610921565b3480156102ef57600080fd5b506101f16109bb565b34801561030457600080fd5b50610228600160a060020a03600435166109c0565b34801561032557600080fd5b50610228600160a060020a0360043516610a57565b34801561034657600080fd5b50610228610af1565b34801561035b57600080fd5b50610228600160a060020a0360043516602435610b5d565b34801561037f57600080fd5b50610228600160a060020a0360043516610b6a565b3480156103a057600080fd5b506103a9610c04565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103e35781810151838201526020016103cb565b50505050905090810190601f1680156104105780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561042a57600080fd5b506103a960ff60043516610c3b565b34801561044557600080fd5b5061026c600160a060020a0360043516610ce0565b34801561046657600080fd5b5061026c600160a060020a0360043516610cf5565b34801561048757600080fd5b50610490610d0a565b60408051600160a060020a039092168252519081900360200190f35b3480156104b857600080fd5b5061026c600160a060020a0360043516610d19565b3480156104d957600080fd5b50610228600160a060020a0360043516610d2e565b3480156104fa57600080fd5b50610228600160a060020a0360043516610dc5565b34801561051b57600080fd5b50610228600160a060020a0360043516610dfd565b34801561053c57600080fd5b50610228600160a060020a0360043516610e97565b34801561055d57600080fd5b506103a9610f2e565b34801561057257600080fd5b506101f1610f65565b34801561058757600080fd5b5061026c610f6a565b34801561059c57600080fd5b506101f1610f7a565b3480156105b157600080fd5b506101f1600160a060020a0360043581169060243516604435610f7f565b3480156105db57600080fd5b506103a961112c565b3480156105f057600080fd5b50610228600160a060020a0360043516611163565b34801561061157600080fd5b50610228600160a060020a03600435166111fa565b34801561063257600080fd5b50610228600160a060020a03600435166111fd565b34801561065357600080fd5b5061026c600160a060020a0360043516611297565b34801561067457600080fd5b5061026c600160a060020a03600435166112ac565b34801561069557600080fd5b506103a96112c1565b3480156106aa57600080fd5b5061026c600160a060020a03600435166112f8565b3480156106cb57600080fd5b50610228600160a060020a0360043516611316565b3480156106ec57600080fd5b5061026c600160a060020a0360043516611336565b600081565b61070f336112f8565b806107245750600054600160a060020a031633145b151561077c576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600360205260409020805460ff19166001179055565b6107a9336112f8565b806107be5750600054600160a060020a031633145b1515610816576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600460205260409020805460ff19169055565b60076020526000908152604090205460ff1681565b600054600160a060020a0316331461086357600080fd5b600160a060020a03166000908152600160208190526040909120805460ff19169091179055565b610893336112f8565b806108a85750600054600160a060020a031633145b1515610900576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600260205260409020805460ff19169055565b61092a336112f8565b8061093f5750600054600160a060020a031633145b1515610997576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600460205260409020805460ff19166001179055565b600181565b6109c9336112f8565b806109de5750600054600160a060020a031633145b1515610a36576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600560205260409020805460ff19169055565b610a60336112f8565b80610a755750600054600160a060020a031633145b1515610acd576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600760205260409020805460ff19166001179055565b600054600160a060020a03163314610b0857600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b610b6682610921565b5050565b610b73336112f8565b80610b885750600054600160a060020a031633145b1515610be0576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600260205260409020805460ff19166001179055565b60408051808201909152601281527f4d5553545f42455f415f434956494c49414e0000000000000000000000000000602082015281565b60ff811660009081526008602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610cd45780601f10610ca957610100808354040283529160200191610cd4565b820191906000526020600020905b815481529060010190602001808311610cb757829003601f168201915b50505050509050919050565b60056020526000908152604090205460ff1681565b60046020526000908152604090205460ff1681565b600054600160a060020a031681565b60026020526000908152604090205460ff1681565b610d37336112f8565b80610d4c5750600054600160a060020a031633145b1515610da4576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600360205260409020805460ff19169055565b600054600160a060020a03163314610ddc57600080fd5b600160a060020a03166000908152600160205260409020805460ff19169055565b610e06336112f8565b80610e1b5750600054600160a060020a031633145b1515610e73576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600660205260409020805460ff19166001179055565b610ea0336112f8565b80610eb55750600054600160a060020a031633145b1515610f0d576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600660205260409020805460ff19169055565b60408051808201909152601081527f4d5553545f42455f564552494649454400000000000000000000000000000000602082015281565b600381565b6000610f75336112f8565b905090565b600281565b600160a060020a03831660009081526002602052604081205460ff1680610fbe5750600160a060020a03841660009081526004602052604090205460ff165b15610fcb57506000611125565b600160a060020a03841660009081526006602052604090205460ff161561104057600160a060020a03831660009081526005602052604090205460ff168061102b5750600160a060020a03831660009081526002602052604090205460ff165b1561103857506000611125565b506003611125565b600160a060020a03841660009081526007602052604090205460ff16156110b557600160a060020a03831660009081526002602052604090205460ff16806110a05750600160a060020a03831660009081526003602052604090205460ff165b156110ad57506000611125565b506002611125565b600160a060020a03841660009081526003602052604090205460ff161561112157600160a060020a03831660009081526002602052604090205460ff16806110a05750600160a060020a03831660009081526007602052604090205460ff16156110ad57506000611125565b5060015b9392505050565b60408051808201909152600781527f5355434345535300000000000000000000000000000000000000000000000000602082015281565b61116c336112f8565b806111815750600054600160a060020a031633145b15156111d9576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600760205260409020805460ff19169055565b50565b611206336112f8565b8061121b5750600054600160a060020a031633145b1515611273576040805160e560020a62461bcd02815260206004820152602f60248201526000805160206113e983398151915260448201526000805160206113c9833981519152606482015290519081900360840190fd5b600160a060020a03166000908152600560205260409020805460ff19166001179055565b60036020526000908152604090205460ff1681565b60066020526000908152604090205460ff1681565b60408051808201909152601081527f4d5553545f42455f554e4c4f434b454400000000000000000000000000000000602082015281565b600160a060020a031660009081526001602052604090205460ff1690565b600054600160a060020a0316331461132d57600080fd5b6111fa8161134b565b60016020526000908152604090205460ff1681565b600160a060020a038116151561136057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905556006f726d207468697320616374696f6e00000000000000000000000000000000004f6e6c79206d616e6167657273206f72206f776e657273206d61792070657266a165627a7a723058206eee73c4dd333fb677620e1d09da7387908d904dccbbb76c2670ed302be59ff60029`

// DeployCivilTokenControllerContract deploys a new Ethereum contract, binding an instance of CivilTokenControllerContract to it.
func DeployCivilTokenControllerContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CivilTokenControllerContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CivilTokenControllerContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CivilTokenControllerContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CivilTokenControllerContract{CivilTokenControllerContractCaller: CivilTokenControllerContractCaller{contract: contract}, CivilTokenControllerContractTransactor: CivilTokenControllerContractTransactor{contract: contract}, CivilTokenControllerContractFilterer: CivilTokenControllerContractFilterer{contract: contract}}, nil
}

// CivilTokenControllerContract is an auto generated Go binding around an Ethereum contract.
type CivilTokenControllerContract struct {
	CivilTokenControllerContractCaller     // Read-only binding to the contract
	CivilTokenControllerContractTransactor // Write-only binding to the contract
	CivilTokenControllerContractFilterer   // Log filterer for contract events
}

// CivilTokenControllerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CivilTokenControllerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilTokenControllerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CivilTokenControllerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilTokenControllerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CivilTokenControllerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilTokenControllerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CivilTokenControllerContractSession struct {
	Contract     *CivilTokenControllerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CivilTokenControllerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CivilTokenControllerContractCallerSession struct {
	Contract *CivilTokenControllerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// CivilTokenControllerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CivilTokenControllerContractTransactorSession struct {
	Contract     *CivilTokenControllerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// CivilTokenControllerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CivilTokenControllerContractRaw struct {
	Contract *CivilTokenControllerContract // Generic contract binding to access the raw methods on
}

// CivilTokenControllerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CivilTokenControllerContractCallerRaw struct {
	Contract *CivilTokenControllerContractCaller // Generic read-only contract binding to access the raw methods on
}

// CivilTokenControllerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CivilTokenControllerContractTransactorRaw struct {
	Contract *CivilTokenControllerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCivilTokenControllerContract creates a new instance of CivilTokenControllerContract, bound to a specific deployed contract.
func NewCivilTokenControllerContract(address common.Address, backend bind.ContractBackend) (*CivilTokenControllerContract, error) {
	contract, err := bindCivilTokenControllerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CivilTokenControllerContract{CivilTokenControllerContractCaller: CivilTokenControllerContractCaller{contract: contract}, CivilTokenControllerContractTransactor: CivilTokenControllerContractTransactor{contract: contract}, CivilTokenControllerContractFilterer: CivilTokenControllerContractFilterer{contract: contract}}, nil
}

// NewCivilTokenControllerContractCaller creates a new read-only instance of CivilTokenControllerContract, bound to a specific deployed contract.
func NewCivilTokenControllerContractCaller(address common.Address, caller bind.ContractCaller) (*CivilTokenControllerContractCaller, error) {
	contract, err := bindCivilTokenControllerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CivilTokenControllerContractCaller{contract: contract}, nil
}

// NewCivilTokenControllerContractTransactor creates a new write-only instance of CivilTokenControllerContract, bound to a specific deployed contract.
func NewCivilTokenControllerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CivilTokenControllerContractTransactor, error) {
	contract, err := bindCivilTokenControllerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CivilTokenControllerContractTransactor{contract: contract}, nil
}

// NewCivilTokenControllerContractFilterer creates a new log filterer instance of CivilTokenControllerContract, bound to a specific deployed contract.
func NewCivilTokenControllerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CivilTokenControllerContractFilterer, error) {
	contract, err := bindCivilTokenControllerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CivilTokenControllerContractFilterer{contract: contract}, nil
}

// bindCivilTokenControllerContract binds a generic wrapper to an already deployed contract.
func bindCivilTokenControllerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CivilTokenControllerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CivilTokenControllerContract *CivilTokenControllerContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CivilTokenControllerContract.Contract.CivilTokenControllerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CivilTokenControllerContract *CivilTokenControllerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.CivilTokenControllerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CivilTokenControllerContract *CivilTokenControllerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.CivilTokenControllerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CivilTokenControllerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.contract.Transact(opts, method, params...)
}

// MUSTBEACIVILIANCODE is a free data retrieval call binding the contract method 0x439b650b.
//
// Solidity: function MUST_BE_A_CIVILIAN_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) MUSTBEACIVILIANCODE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "MUST_BE_A_CIVILIAN_CODE")
	return *ret0, err
}

// MUSTBEACIVILIANCODE is a free data retrieval call binding the contract method 0x439b650b.
//
// Solidity: function MUST_BE_A_CIVILIAN_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) MUSTBEACIVILIANCODE() (uint8, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEACIVILIANCODE(&_CivilTokenControllerContract.CallOpts)
}

// MUSTBEACIVILIANCODE is a free data retrieval call binding the contract method 0x439b650b.
//
// Solidity: function MUST_BE_A_CIVILIAN_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) MUSTBEACIVILIANCODE() (uint8, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEACIVILIANCODE(&_CivilTokenControllerContract.CallOpts)
}

// MUSTBEACIVILIANERROR is a free data retrieval call binding the contract method 0x78a21ddb.
//
// Solidity: function MUST_BE_A_CIVILIAN_ERROR() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) MUSTBEACIVILIANERROR(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "MUST_BE_A_CIVILIAN_ERROR")
	return *ret0, err
}

// MUSTBEACIVILIANERROR is a free data retrieval call binding the contract method 0x78a21ddb.
//
// Solidity: function MUST_BE_A_CIVILIAN_ERROR() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) MUSTBEACIVILIANERROR() (string, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEACIVILIANERROR(&_CivilTokenControllerContract.CallOpts)
}

// MUSTBEACIVILIANERROR is a free data retrieval call binding the contract method 0x78a21ddb.
//
// Solidity: function MUST_BE_A_CIVILIAN_ERROR() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) MUSTBEACIVILIANERROR() (string, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEACIVILIANERROR(&_CivilTokenControllerContract.CallOpts)
}

// MUSTBEUNLOCKEDCODE is a free data retrieval call binding the contract method 0xca3aaa9a.
//
// Solidity: function MUST_BE_UNLOCKED_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) MUSTBEUNLOCKEDCODE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "MUST_BE_UNLOCKED_CODE")
	return *ret0, err
}

// MUSTBEUNLOCKEDCODE is a free data retrieval call binding the contract method 0xca3aaa9a.
//
// Solidity: function MUST_BE_UNLOCKED_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) MUSTBEUNLOCKEDCODE() (uint8, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEUNLOCKEDCODE(&_CivilTokenControllerContract.CallOpts)
}

// MUSTBEUNLOCKEDCODE is a free data retrieval call binding the contract method 0xca3aaa9a.
//
// Solidity: function MUST_BE_UNLOCKED_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) MUSTBEUNLOCKEDCODE() (uint8, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEUNLOCKEDCODE(&_CivilTokenControllerContract.CallOpts)
}

// MUSTBEUNLOCKEDERROR is a free data retrieval call binding the contract method 0xf1983918.
//
// Solidity: function MUST_BE_UNLOCKED_ERROR() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) MUSTBEUNLOCKEDERROR(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "MUST_BE_UNLOCKED_ERROR")
	return *ret0, err
}

// MUSTBEUNLOCKEDERROR is a free data retrieval call binding the contract method 0xf1983918.
//
// Solidity: function MUST_BE_UNLOCKED_ERROR() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) MUSTBEUNLOCKEDERROR() (string, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEUNLOCKEDERROR(&_CivilTokenControllerContract.CallOpts)
}

// MUSTBEUNLOCKEDERROR is a free data retrieval call binding the contract method 0xf1983918.
//
// Solidity: function MUST_BE_UNLOCKED_ERROR() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) MUSTBEUNLOCKEDERROR() (string, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEUNLOCKEDERROR(&_CivilTokenControllerContract.CallOpts)
}

// MUSTBEVERIFIEDCODE is a free data retrieval call binding the contract method 0xc0e9794a.
//
// Solidity: function MUST_BE_VERIFIED_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) MUSTBEVERIFIEDCODE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "MUST_BE_VERIFIED_CODE")
	return *ret0, err
}

// MUSTBEVERIFIEDCODE is a free data retrieval call binding the contract method 0xc0e9794a.
//
// Solidity: function MUST_BE_VERIFIED_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) MUSTBEVERIFIEDCODE() (uint8, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEVERIFIEDCODE(&_CivilTokenControllerContract.CallOpts)
}

// MUSTBEVERIFIEDCODE is a free data retrieval call binding the contract method 0xc0e9794a.
//
// Solidity: function MUST_BE_VERIFIED_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) MUSTBEVERIFIEDCODE() (uint8, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEVERIFIEDCODE(&_CivilTokenControllerContract.CallOpts)
}

// MUSTBEVERIFIEDERROR is a free data retrieval call binding the contract method 0xbdcadb3b.
//
// Solidity: function MUST_BE_VERIFIED_ERROR() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) MUSTBEVERIFIEDERROR(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "MUST_BE_VERIFIED_ERROR")
	return *ret0, err
}

// MUSTBEVERIFIEDERROR is a free data retrieval call binding the contract method 0xbdcadb3b.
//
// Solidity: function MUST_BE_VERIFIED_ERROR() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) MUSTBEVERIFIEDERROR() (string, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEVERIFIEDERROR(&_CivilTokenControllerContract.CallOpts)
}

// MUSTBEVERIFIEDERROR is a free data retrieval call binding the contract method 0xbdcadb3b.
//
// Solidity: function MUST_BE_VERIFIED_ERROR() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) MUSTBEVERIFIEDERROR() (string, error) {
	return _CivilTokenControllerContract.Contract.MUSTBEVERIFIEDERROR(&_CivilTokenControllerContract.CallOpts)
}

// SUCCESSCODE is a free data retrieval call binding the contract method 0x0e969a05.
//
// Solidity: function SUCCESS_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) SUCCESSCODE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "SUCCESS_CODE")
	return *ret0, err
}

// SUCCESSCODE is a free data retrieval call binding the contract method 0x0e969a05.
//
// Solidity: function SUCCESS_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) SUCCESSCODE() (uint8, error) {
	return _CivilTokenControllerContract.Contract.SUCCESSCODE(&_CivilTokenControllerContract.CallOpts)
}

// SUCCESSCODE is a free data retrieval call binding the contract method 0x0e969a05.
//
// Solidity: function SUCCESS_CODE() constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) SUCCESSCODE() (uint8, error) {
	return _CivilTokenControllerContract.Contract.SUCCESSCODE(&_CivilTokenControllerContract.CallOpts)
}

// SUCCESSMESSAGE is a free data retrieval call binding the contract method 0xe7984d17.
//
// Solidity: function SUCCESS_MESSAGE() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) SUCCESSMESSAGE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "SUCCESS_MESSAGE")
	return *ret0, err
}

// SUCCESSMESSAGE is a free data retrieval call binding the contract method 0xe7984d17.
//
// Solidity: function SUCCESS_MESSAGE() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) SUCCESSMESSAGE() (string, error) {
	return _CivilTokenControllerContract.Contract.SUCCESSMESSAGE(&_CivilTokenControllerContract.CallOpts)
}

// SUCCESSMESSAGE is a free data retrieval call binding the contract method 0xe7984d17.
//
// Solidity: function SUCCESS_MESSAGE() constant returns(string)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) SUCCESSMESSAGE() (string, error) {
	return _CivilTokenControllerContract.Contract.SUCCESSMESSAGE(&_CivilTokenControllerContract.CallOpts)
}

// CheckManagerStatus is a free data retrieval call binding the contract method 0xf281e7d1.
//
// Solidity: function checkManagerStatus(managerAddress address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) CheckManagerStatus(opts *bind.CallOpts, managerAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "checkManagerStatus", managerAddress)
	return *ret0, err
}

// CheckManagerStatus is a free data retrieval call binding the contract method 0xf281e7d1.
//
// Solidity: function checkManagerStatus(managerAddress address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) CheckManagerStatus(managerAddress common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.CheckManagerStatus(&_CivilTokenControllerContract.CallOpts, managerAddress)
}

// CheckManagerStatus is a free data retrieval call binding the contract method 0xf281e7d1.
//
// Solidity: function checkManagerStatus(managerAddress address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) CheckManagerStatus(managerAddress common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.CheckManagerStatus(&_CivilTokenControllerContract.CallOpts, managerAddress)
}

// CivilianList is a free data retrieval call binding the contract method 0xee56f4fa.
//
// Solidity: function civilianList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) CivilianList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "civilianList", arg0)
	return *ret0, err
}

// CivilianList is a free data retrieval call binding the contract method 0xee56f4fa.
//
// Solidity: function civilianList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) CivilianList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.CivilianList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// CivilianList is a free data retrieval call binding the contract method 0xee56f4fa.
//
// Solidity: function civilianList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) CivilianList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.CivilianList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// CoreList is a free data retrieval call binding the contract method 0x9685cc57.
//
// Solidity: function coreList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) CoreList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "coreList", arg0)
	return *ret0, err
}

// CoreList is a free data retrieval call binding the contract method 0x9685cc57.
//
// Solidity: function coreList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) CoreList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.CoreList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// CoreList is a free data retrieval call binding the contract method 0x9685cc57.
//
// Solidity: function coreList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) CoreList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.CoreList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// DetectTransferRestriction is a free data retrieval call binding the contract method 0xd4ce1415.
//
// Solidity: function detectTransferRestriction(from address, to address, value uint256) constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) DetectTransferRestriction(opts *bind.CallOpts, from common.Address, to common.Address, value *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "detectTransferRestriction", from, to, value)
	return *ret0, err
}

// DetectTransferRestriction is a free data retrieval call binding the contract method 0xd4ce1415.
//
// Solidity: function detectTransferRestriction(from address, to address, value uint256) constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) DetectTransferRestriction(from common.Address, to common.Address, value *big.Int) (uint8, error) {
	return _CivilTokenControllerContract.Contract.DetectTransferRestriction(&_CivilTokenControllerContract.CallOpts, from, to, value)
}

// DetectTransferRestriction is a free data retrieval call binding the contract method 0xd4ce1415.
//
// Solidity: function detectTransferRestriction(from address, to address, value uint256) constant returns(uint8)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) DetectTransferRestriction(from common.Address, to common.Address, value *big.Int) (uint8, error) {
	return _CivilTokenControllerContract.Contract.DetectTransferRestriction(&_CivilTokenControllerContract.CallOpts, from, to, value)
}

// IsManager is a free data retrieval call binding the contract method 0xc56a3e88.
//
// Solidity: function isManager() constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) IsManager(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "isManager")
	return *ret0, err
}

// IsManager is a free data retrieval call binding the contract method 0xc56a3e88.
//
// Solidity: function isManager() constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) IsManager() (bool, error) {
	return _CivilTokenControllerContract.Contract.IsManager(&_CivilTokenControllerContract.CallOpts)
}

// IsManager is a free data retrieval call binding the contract method 0xc56a3e88.
//
// Solidity: function isManager() constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) IsManager() (bool, error) {
	return _CivilTokenControllerContract.Contract.IsManager(&_CivilTokenControllerContract.CallOpts)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) Managers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "managers", arg0)
	return *ret0, err
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) Managers(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.Managers(&_CivilTokenControllerContract.CallOpts, arg0)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) Managers(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.Managers(&_CivilTokenControllerContract.CallOpts, arg0)
}

// MessageForTransferRestriction is a free data retrieval call binding the contract method 0x7f4ab1dd.
//
// Solidity: function messageForTransferRestriction(restrictionCode uint8) constant returns(message string)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) MessageForTransferRestriction(opts *bind.CallOpts, restrictionCode uint8) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "messageForTransferRestriction", restrictionCode)
	return *ret0, err
}

// MessageForTransferRestriction is a free data retrieval call binding the contract method 0x7f4ab1dd.
//
// Solidity: function messageForTransferRestriction(restrictionCode uint8) constant returns(message string)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) MessageForTransferRestriction(restrictionCode uint8) (string, error) {
	return _CivilTokenControllerContract.Contract.MessageForTransferRestriction(&_CivilTokenControllerContract.CallOpts, restrictionCode)
}

// MessageForTransferRestriction is a free data retrieval call binding the contract method 0x7f4ab1dd.
//
// Solidity: function messageForTransferRestriction(restrictionCode uint8) constant returns(message string)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) MessageForTransferRestriction(restrictionCode uint8) (string, error) {
	return _CivilTokenControllerContract.Contract.MessageForTransferRestriction(&_CivilTokenControllerContract.CallOpts, restrictionCode)
}

// NewsroomMultisigList is a free data retrieval call binding the contract method 0x166542b3.
//
// Solidity: function newsroomMultisigList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) NewsroomMultisigList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "newsroomMultisigList", arg0)
	return *ret0, err
}

// NewsroomMultisigList is a free data retrieval call binding the contract method 0x166542b3.
//
// Solidity: function newsroomMultisigList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) NewsroomMultisigList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.NewsroomMultisigList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// NewsroomMultisigList is a free data retrieval call binding the contract method 0x166542b3.
//
// Solidity: function newsroomMultisigList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) NewsroomMultisigList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.NewsroomMultisigList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) Owner() (common.Address, error) {
	return _CivilTokenControllerContract.Contract.Owner(&_CivilTokenControllerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) Owner() (common.Address, error) {
	return _CivilTokenControllerContract.Contract.Owner(&_CivilTokenControllerContract.CallOpts)
}

// StorefrontList is a free data retrieval call binding the contract method 0xf0fbca06.
//
// Solidity: function storefrontList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) StorefrontList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "storefrontList", arg0)
	return *ret0, err
}

// StorefrontList is a free data retrieval call binding the contract method 0xf0fbca06.
//
// Solidity: function storefrontList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) StorefrontList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.StorefrontList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// StorefrontList is a free data retrieval call binding the contract method 0xf0fbca06.
//
// Solidity: function storefrontList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) StorefrontList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.StorefrontList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// UnlockedList is a free data retrieval call binding the contract method 0x8bdeb452.
//
// Solidity: function unlockedList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) UnlockedList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "unlockedList", arg0)
	return *ret0, err
}

// UnlockedList is a free data retrieval call binding the contract method 0x8bdeb452.
//
// Solidity: function unlockedList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) UnlockedList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.UnlockedList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// UnlockedList is a free data retrieval call binding the contract method 0x8bdeb452.
//
// Solidity: function unlockedList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) UnlockedList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.UnlockedList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// VerifiedList is a free data retrieval call binding the contract method 0x81601496.
//
// Solidity: function verifiedList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCaller) VerifiedList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTokenControllerContract.contract.Call(opts, out, "verifiedList", arg0)
	return *ret0, err
}

// VerifiedList is a free data retrieval call binding the contract method 0x81601496.
//
// Solidity: function verifiedList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) VerifiedList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.VerifiedList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// VerifiedList is a free data retrieval call binding the contract method 0x81601496.
//
// Solidity: function verifiedList( address) constant returns(bool)
func (_CivilTokenControllerContract *CivilTokenControllerContractCallerSession) VerifiedList(arg0 common.Address) (bool, error) {
	return _CivilTokenControllerContract.Contract.VerifiedList(&_CivilTokenControllerContract.CallOpts, arg0)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(managerAddress address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) AddManager(opts *bind.TransactOpts, managerAddress common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "addManager", managerAddress)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(managerAddress address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) AddManager(managerAddress common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddManager(&_CivilTokenControllerContract.TransactOpts, managerAddress)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(managerAddress address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) AddManager(managerAddress common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddManager(&_CivilTokenControllerContract.TransactOpts, managerAddress)
}

// AddToCivilians is a paid mutator transaction binding the contract method 0x1299090a.
//
// Solidity: function addToCivilians(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) AddToCivilians(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "addToCivilians", operator)
}

// AddToCivilians is a paid mutator transaction binding the contract method 0x1299090a.
//
// Solidity: function addToCivilians(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) AddToCivilians(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToCivilians(&_CivilTokenControllerContract.TransactOpts, operator)
}

// AddToCivilians is a paid mutator transaction binding the contract method 0x1299090a.
//
// Solidity: function addToCivilians(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) AddToCivilians(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToCivilians(&_CivilTokenControllerContract.TransactOpts, operator)
}

// AddToCore is a paid mutator transaction binding the contract method 0x73c60182.
//
// Solidity: function addToCore(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) AddToCore(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "addToCore", operator)
}

// AddToCore is a paid mutator transaction binding the contract method 0x73c60182.
//
// Solidity: function addToCore(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) AddToCore(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToCore(&_CivilTokenControllerContract.TransactOpts, operator)
}

// AddToCore is a paid mutator transaction binding the contract method 0x73c60182.
//
// Solidity: function addToCore(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) AddToCore(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToCore(&_CivilTokenControllerContract.TransactOpts, operator)
}

// AddToNewsroomMultisigs is a paid mutator transaction binding the contract method 0x607c60bb.
//
// Solidity: function addToNewsroomMultisigs(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) AddToNewsroomMultisigs(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "addToNewsroomMultisigs", operator)
}

// AddToNewsroomMultisigs is a paid mutator transaction binding the contract method 0x607c60bb.
//
// Solidity: function addToNewsroomMultisigs(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) AddToNewsroomMultisigs(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToNewsroomMultisigs(&_CivilTokenControllerContract.TransactOpts, operator)
}

// AddToNewsroomMultisigs is a paid mutator transaction binding the contract method 0x607c60bb.
//
// Solidity: function addToNewsroomMultisigs(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) AddToNewsroomMultisigs(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToNewsroomMultisigs(&_CivilTokenControllerContract.TransactOpts, operator)
}

// AddToStorefront is a paid mutator transaction binding the contract method 0xac1f3853.
//
// Solidity: function addToStorefront(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) AddToStorefront(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "addToStorefront", operator)
}

// AddToStorefront is a paid mutator transaction binding the contract method 0xac1f3853.
//
// Solidity: function addToStorefront(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) AddToStorefront(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToStorefront(&_CivilTokenControllerContract.TransactOpts, operator)
}

// AddToStorefront is a paid mutator transaction binding the contract method 0xac1f3853.
//
// Solidity: function addToStorefront(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) AddToStorefront(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToStorefront(&_CivilTokenControllerContract.TransactOpts, operator)
}

// AddToUnlocked is a paid mutator transaction binding the contract method 0x3f16b282.
//
// Solidity: function addToUnlocked(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) AddToUnlocked(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "addToUnlocked", operator)
}

// AddToUnlocked is a paid mutator transaction binding the contract method 0x3f16b282.
//
// Solidity: function addToUnlocked(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) AddToUnlocked(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToUnlocked(&_CivilTokenControllerContract.TransactOpts, operator)
}

// AddToUnlocked is a paid mutator transaction binding the contract method 0x3f16b282.
//
// Solidity: function addToUnlocked(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) AddToUnlocked(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToUnlocked(&_CivilTokenControllerContract.TransactOpts, operator)
}

// AddToVerified is a paid mutator transaction binding the contract method 0xee37c29f.
//
// Solidity: function addToVerified(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) AddToVerified(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "addToVerified", operator)
}

// AddToVerified is a paid mutator transaction binding the contract method 0xee37c29f.
//
// Solidity: function addToVerified(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) AddToVerified(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToVerified(&_CivilTokenControllerContract.TransactOpts, operator)
}

// AddToVerified is a paid mutator transaction binding the contract method 0xee37c29f.
//
// Solidity: function addToVerified(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) AddToVerified(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.AddToVerified(&_CivilTokenControllerContract.TransactOpts, operator)
}

// CheckProofOfUse is a paid mutator transaction binding the contract method 0xe99f29a4.
//
// Solidity: function checkProofOfUse(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) CheckProofOfUse(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "checkProofOfUse", operator)
}

// CheckProofOfUse is a paid mutator transaction binding the contract method 0xe99f29a4.
//
// Solidity: function checkProofOfUse(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) CheckProofOfUse(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.CheckProofOfUse(&_CivilTokenControllerContract.TransactOpts, operator)
}

// CheckProofOfUse is a paid mutator transaction binding the contract method 0xe99f29a4.
//
// Solidity: function checkProofOfUse(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) CheckProofOfUse(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.CheckProofOfUse(&_CivilTokenControllerContract.TransactOpts, operator)
}

// OnRequestVotingRights is a paid mutator transaction binding the contract method 0x72524873.
//
// Solidity: function onRequestVotingRights(user address, tokenAmount uint256) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) OnRequestVotingRights(opts *bind.TransactOpts, user common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "onRequestVotingRights", user, tokenAmount)
}

// OnRequestVotingRights is a paid mutator transaction binding the contract method 0x72524873.
//
// Solidity: function onRequestVotingRights(user address, tokenAmount uint256) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) OnRequestVotingRights(user common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.OnRequestVotingRights(&_CivilTokenControllerContract.TransactOpts, user, tokenAmount)
}

// OnRequestVotingRights is a paid mutator transaction binding the contract method 0x72524873.
//
// Solidity: function onRequestVotingRights(user address, tokenAmount uint256) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) OnRequestVotingRights(user common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.OnRequestVotingRights(&_CivilTokenControllerContract.TransactOpts, user, tokenAmount)
}

// RemoveFromCivilians is a paid mutator transaction binding the contract method 0xab3d0c7a.
//
// Solidity: function removeFromCivilians(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) RemoveFromCivilians(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "removeFromCivilians", operator)
}

// RemoveFromCivilians is a paid mutator transaction binding the contract method 0xab3d0c7a.
//
// Solidity: function removeFromCivilians(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) RemoveFromCivilians(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromCivilians(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveFromCivilians is a paid mutator transaction binding the contract method 0xab3d0c7a.
//
// Solidity: function removeFromCivilians(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) RemoveFromCivilians(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromCivilians(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveFromCore is a paid mutator transaction binding the contract method 0x3c421424.
//
// Solidity: function removeFromCore(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) RemoveFromCore(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "removeFromCore", operator)
}

// RemoveFromCore is a paid mutator transaction binding the contract method 0x3c421424.
//
// Solidity: function removeFromCore(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) RemoveFromCore(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromCore(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveFromCore is a paid mutator transaction binding the contract method 0x3c421424.
//
// Solidity: function removeFromCore(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) RemoveFromCore(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromCore(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveFromNewsroomMultisigs is a paid mutator transaction binding the contract method 0xe79a4fd4.
//
// Solidity: function removeFromNewsroomMultisigs(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) RemoveFromNewsroomMultisigs(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "removeFromNewsroomMultisigs", operator)
}

// RemoveFromNewsroomMultisigs is a paid mutator transaction binding the contract method 0xe79a4fd4.
//
// Solidity: function removeFromNewsroomMultisigs(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) RemoveFromNewsroomMultisigs(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromNewsroomMultisigs(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveFromNewsroomMultisigs is a paid mutator transaction binding the contract method 0xe79a4fd4.
//
// Solidity: function removeFromNewsroomMultisigs(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) RemoveFromNewsroomMultisigs(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromNewsroomMultisigs(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveFromStorefront is a paid mutator transaction binding the contract method 0xadbb9160.
//
// Solidity: function removeFromStorefront(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) RemoveFromStorefront(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "removeFromStorefront", operator)
}

// RemoveFromStorefront is a paid mutator transaction binding the contract method 0xadbb9160.
//
// Solidity: function removeFromStorefront(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) RemoveFromStorefront(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromStorefront(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveFromStorefront is a paid mutator transaction binding the contract method 0xadbb9160.
//
// Solidity: function removeFromStorefront(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) RemoveFromStorefront(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromStorefront(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveFromUnlocked is a paid mutator transaction binding the contract method 0x14862ea5.
//
// Solidity: function removeFromUnlocked(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) RemoveFromUnlocked(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "removeFromUnlocked", operator)
}

// RemoveFromUnlocked is a paid mutator transaction binding the contract method 0x14862ea5.
//
// Solidity: function removeFromUnlocked(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) RemoveFromUnlocked(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromUnlocked(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveFromUnlocked is a paid mutator transaction binding the contract method 0x14862ea5.
//
// Solidity: function removeFromUnlocked(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) RemoveFromUnlocked(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromUnlocked(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveFromVerified is a paid mutator transaction binding the contract method 0x4b3d1485.
//
// Solidity: function removeFromVerified(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) RemoveFromVerified(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "removeFromVerified", operator)
}

// RemoveFromVerified is a paid mutator transaction binding the contract method 0x4b3d1485.
//
// Solidity: function removeFromVerified(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) RemoveFromVerified(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromVerified(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveFromVerified is a paid mutator transaction binding the contract method 0x4b3d1485.
//
// Solidity: function removeFromVerified(operator address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) RemoveFromVerified(operator common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveFromVerified(&_CivilTokenControllerContract.TransactOpts, operator)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(managerAddress address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) RemoveManager(opts *bind.TransactOpts, managerAddress common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "removeManager", managerAddress)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(managerAddress address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) RemoveManager(managerAddress common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveManager(&_CivilTokenControllerContract.TransactOpts, managerAddress)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(managerAddress address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) RemoveManager(managerAddress common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RemoveManager(&_CivilTokenControllerContract.TransactOpts, managerAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RenounceOwnership(&_CivilTokenControllerContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.RenounceOwnership(&_CivilTokenControllerContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.TransferOwnership(&_CivilTokenControllerContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_CivilTokenControllerContract *CivilTokenControllerContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _CivilTokenControllerContract.Contract.TransferOwnership(&_CivilTokenControllerContract.TransactOpts, _newOwner)
}

// CivilTokenControllerContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the CivilTokenControllerContract contract.
type CivilTokenControllerContractOwnershipRenouncedIterator struct {
	Event *CivilTokenControllerContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *CivilTokenControllerContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTokenControllerContractOwnershipRenounced)
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
		it.Event = new(CivilTokenControllerContractOwnershipRenounced)
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
func (it *CivilTokenControllerContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTokenControllerContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTokenControllerContractOwnershipRenounced represents a OwnershipRenounced event raised by the CivilTokenControllerContract contract.
type CivilTokenControllerContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_CivilTokenControllerContract *CivilTokenControllerContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*CivilTokenControllerContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _CivilTokenControllerContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CivilTokenControllerContractOwnershipRenouncedIterator{contract: _CivilTokenControllerContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_CivilTokenControllerContract *CivilTokenControllerContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *CivilTokenControllerContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _CivilTokenControllerContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTokenControllerContractOwnershipRenounced)
				if err := _CivilTokenControllerContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// CivilTokenControllerContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CivilTokenControllerContract contract.
type CivilTokenControllerContractOwnershipTransferredIterator struct {
	Event *CivilTokenControllerContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CivilTokenControllerContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTokenControllerContractOwnershipTransferred)
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
		it.Event = new(CivilTokenControllerContractOwnershipTransferred)
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
func (it *CivilTokenControllerContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTokenControllerContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTokenControllerContractOwnershipTransferred represents a OwnershipTransferred event raised by the CivilTokenControllerContract contract.
type CivilTokenControllerContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CivilTokenControllerContract *CivilTokenControllerContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CivilTokenControllerContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CivilTokenControllerContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CivilTokenControllerContractOwnershipTransferredIterator{contract: _CivilTokenControllerContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CivilTokenControllerContract *CivilTokenControllerContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CivilTokenControllerContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CivilTokenControllerContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTokenControllerContractOwnershipTransferred)
				if err := _CivilTokenControllerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
