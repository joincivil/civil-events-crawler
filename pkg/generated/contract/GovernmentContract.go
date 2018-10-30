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
const GovernmentContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"PROCESSBY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposals\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"processBy\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"constitutionURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"constitutionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appellate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"params\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governmentController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voting\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"appellateAddr\",\"type\":\"address\"},{\"name\":\"governmentControllerAddr\",\"type\":\"address\"},{\"name\":\"plcrAddr\",\"type\":\"address\"},{\"name\":\"appealFeeAmount\",\"type\":\"uint256\"},{\"name\":\"requestAppealLength\",\"type\":\"uint256\"},{\"name\":\"judgeAppealLength\",\"type\":\"uint256\"},{\"name\":\"appealSupermajorityPercentage\",\"type\":\"uint256\"},{\"name\":\"appealChallengeVoteDispensationPct\",\"type\":\"uint256\"},{\"name\":\"pDeposit\",\"type\":\"uint256\"},{\"name\":\"pCommitStageLength\",\"type\":\"uint256\"},{\"name\":\"pRevealStageLength\",\"type\":\"uint256\"},{\"name\":\"constHash\",\"type\":\"bytes32\"},{\"name\":\"constURI\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAppellate\",\"type\":\"address\"}],\"name\":\"_AppellateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"_ParameterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"propID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"pollID\",\"type\":\"uint256\"}],\"name\":\"_GovtReparameterizationProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"propId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"pollID\",\"type\":\"uint256\"}],\"name\":\"_ProposalPassed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"propId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"pollID\",\"type\":\"uint256\"}],\"name\":\"_ProposalExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"propId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"pollID\",\"type\":\"uint256\"}],\"name\":\"_ProposalFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proposedHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"proposedURI\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"pollID\",\"type\":\"uint256\"}],\"name\":\"_NewConstProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"constHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"constURI\",\"type\":\"string\"}],\"name\":\"_NewConstProposalPassed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"constHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"constURI\",\"type\":\"string\"}],\"name\":\"_NewConstProposalExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"constHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"constURI\",\"type\":\"string\"}],\"name\":\"_NewConstProposalFailed\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAppellate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGovernmentController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"proposeReparameterization\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newConstHash\",\"type\":\"bytes32\"},{\"name\":\"_newConstURI\",\"type\":\"string\"}],\"name\":\"proposeNewConstitution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_propID\",\"type\":\"bytes32\"}],\"name\":\"processProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"processConstChangeProp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_propID\",\"type\":\"bytes32\"}],\"name\":\"propExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_propID\",\"type\":\"bytes32\"}],\"name\":\"propCanBeResolved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"constChangePropCanBeResolved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAppellate\",\"type\":\"address\"}],\"name\":\"setAppellate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GovernmentContractBin is the compiled bytecode used for deploying new contracts.
const GovernmentContractBin = `0x608060405262093a80600b553480156200001857600080fd5b5060405162001f4738038062001f4783398101604090815281516020830151918301516060840151608085015160a086015160c087015160e08801516101008901516101208a01516101408b01516101608c01516101808d01519a9c999a98999798969795969495939492939192909101600160a060020a038d1615156200009f57600080fd5b600160a060020a038c161515620000b557600080fd5b8c6000806101000a815481600160a060020a030219169083600160a060020a031602179055508b600160006101000a815481600160a060020a030219169083600160a060020a031602179055508a600a60006101000a815481600160a060020a030219169083600160a060020a03160217905550620001796040805190810160405280601081526020017f7265717565737441707065616c4c656e000000000000000000000000000000008152508a6200038b640100000000026401000000009004565b60408051808201909152600e81527f6a7564676541707065616c4c656e0000000000000000000000000000000000006020820152620001c290896401000000006200038b810204565b60408051808201909152600981527f61707065616c466565000000000000000000000000000000000000000000000060208201526200020b908b6401000000006200038b810204565b60408051808201909152601481527f61707065616c566f746550657263656e7461676500000000000000000000000060208201526200025490886401000000006200038b810204565b620002cb606060405190810160405280602281526020017f61707065616c4368616c6c656e6765566f746544697370656e736174696f6e5081526020017f6374000000000000000000000000000000000000000000000000000000000000815250876200038b640100000000026401000000009004565b60408051808201909152601381527f676f767450436f6d6d697453746167654c656e0000000000000000000000000060208201526200031490856401000000006200038b810204565b60408051808201909152601381527f676f76745052657665616c53746167654c656e0000000000000000000000000060208201526200035d90846401000000006200038b810204565b6002829055805162000377906003906020840190620004b1565b505050505050505050505050505062000556565b8060046000846040518082805190602001908083835b60208310620003c25780518252601f199092019160209182019101620003a1565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916600019168152602001908152602001600020819055507f0e92bd4b74871caaf73a4a51ca5ad4f01e5c5215e940a2f2a1f4c755b955066c82826040518080602001838152602001828103825284818151815260200191508051906020019080838360005b838110156200047157818101518382015260200162000457565b50505050905090810190601f1680156200049f5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004f457805160ff191683800117855562000524565b8280016001018555821562000524579182015b828111156200052457825182559160200191906001019062000507565b506200053292915062000536565b5090565b6200055391905b808211156200053257600081556001016200053d565b90565b6119e180620005666000396000f3006080604052600436106100ec5763ffffffff60e060020a60003504166229514f81146100f157806330490e911461011857806332ed5b121461013257806335300990146101d7578063551224251461020357806356e1fb88146102245780635793b9cf14610255578063693ec85e1461026a5780638ca7f51c146102c35780639a99fd111461034d578063a33c91d814610362578063bade1c5414610377578063c7d93fd4146103d2578063d5fd9e66146103e7578063d704626c146103fc578063dc6ab5271461045a578063f2a2129b14610472578063fce1ccca14610487578063ffa1bdf01461049c575b600080fd5b3480156100fd57600080fd5b506101066104b4565b60408051918252519081900360200190f35b34801561012457600080fd5b506101306004356104ba565b005b34801561013e57600080fd5b5061014a6004356104d4565b6040518085815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b83811015610199578181015183820152602001610181565b50505050905090810190601f1680156101c65780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b3480156101e357600080fd5b506101ef600435610587565b604080519115158252519081900360200190f35b34801561020f57600080fd5b50610130600160a060020a036004351661059d565b34801561023057600080fd5b5061023961062a565b60408051600160a060020a039092168252519081900360200190f35b34801561026157600080fd5b5061023961063a565b34801561027657600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101069436949293602493928401919081908401838280828437509497506106499650505050505050565b3480156102cf57600080fd5b506102d86106c0565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103125781810151838201526020016102fa565b50505050905090810190601f16801561033f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561035957600080fd5b5061013061074e565b34801561036e57600080fd5b506101ef610765565b34801561038357600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261010694369492936024939284019190819084018382808284375094975050933594506108129350505050565b3480156103de57600080fd5b50610106610d09565b3480156103f357600080fd5b50610239610d0f565b34801561040857600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610130958335953695604494919390910191908190840183828082843750949750610d1e9650505050505050565b34801561046657600080fd5b50610106600435610fb8565b34801561047e57600080fd5b50610239610fca565b34801561049357600080fd5b50610239610fd9565b3480156104a857600080fd5b506101ef600435610fe8565b600b5481565b6104c381610fe8565b156100ec576104d18161116c565b50565b6005602090815260009182526040918290208054600180830180548651600293821615610100026000190190911692909204601f8101869004860283018601909652858252919492939092908301828280156105715780601f1061054657610100808354040283529160200191610571565b820191906000526020600020905b81548152906001019060200180831161055457829003601f168201915b5050505050908060020154908060030154905084565b6000908152600560205260408120600201541190565b600154600160a060020a031633146105b457600080fd5b600160a060020a03811615156105c957600080fd5b60008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f759a9d1715f38685bd08c7fb25060b7b6795cddf54214336e02a0533c5c7b89e9181900360200190a150565b600054600160a060020a03165b90565b600154600160a060020a031690565b600060046000836040518082805190602001908083835b6020831061067f5780518252601f199092019160209182019101610660565b51815160209384036101000a600019018019909216911617905260408051929094018290039091208652850195909552929092016000205495945050505050565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156107465780601f1061071b57610100808354040283529160200191610746565b820191906000526020600020905b81548152906001019060200180831161072957829003601f168201915b505050505081565b610756610765565b156100ec576107636113c1565b565b60008060066000015411801561080d5750600a54600654604080517fee684830000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a039092169163ee684830916024808201926020929091908290030181600087803b1580156107e057600080fd5b505af11580156107f4573d6000803e3d6000fd5b505050506040513d602081101561080a57600080fd5b50515b905090565b6000805481908190600160a060020a0316331461082e57600080fd5b84846040518083805190602001908083835b6020831061085f5780518252601f199092019160209182019101610840565b51815160001960209485036101000a019081169019919091161790529201938452506040805193849003820184207f61707065616c566f746550657263656e74616765000000000000000000000000855290519384900360140184208a5191975094508993925082918401908083835b602083106108ee5780518252601f1990920191602091820191016108cf565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614806109e45750604080517f61707065616c4368616c6c656e6765566f746544697370656e736174696f6e5081527f637400000000000000000000000000000000000000000000000000000000000060208083019190915291519081900360220181208751909288929182918401908083835b602083106109b15780518252601f199092019160209182019101610992565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916145b156109f75760648411156109f757600080fd5b610a0082610587565b15610a0a57600080fd5b83610a1486610649565b1415610a1f57600080fd5b600a5460408051808201909152601481527f61707065616c566f746550657263656e746167650000000000000000000000006020820152600160a060020a03909116906332ed3d6090610a7190610649565b610aaf6040805190810160405280601381526020017f676f767450436f6d6d697453746167654c656e00000000000000000000000000815250610649565b610adb604080519081016040528060138152602001600080516020611996833981519152815250610649565b6040518463ffffffff1660e060020a028152600401808481526020018381526020018281526020019350505050602060405180830381600087803b158015610b2257600080fd5b505af1158015610b36573d6000803e3d6000fd5b505050506040513d6020811015610b4c57600080fd5b5051604080516080810182528281526020818101899052600b54835180850185526013815260008051602061199683398151915292810192909252939450909291830191610bfb91610bef90610ba190610649565b610bef610be26040805190810160405280601381526020017f676f767450436f6d6d697453746167654c656e00000000000000000000000000815250610649565b429063ffffffff6116e716565b9063ffffffff6116e716565b8152602090810186905260008481526005825260409020825181558282015180519192610c309260018501929091019061181b565b5060408201518160020155606082015181600301559050507f74adf299a4c734e1ae114977ab264221c2f4a914c02243561aaa9158735d32248585848460405180806020018581526020018460001916600019168152602001838152602001828103825286818151815260200191508051906020019080838360005b83811015610cc4578181015183820152602001610cac565b50505050905090810190601f168015610cf15780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1509392505050565b60025481565b600054600160a060020a031681565b60008054600160a060020a03163314610d3657600080fd5b60095415610d4357600080fd5b600a5460408051808201909152601481527f61707065616c566f746550657263656e746167650000000000000000000000006020820152600160a060020a03909116906332ed3d6090610d9590610649565b610dd36040805190810160405280601381526020017f676f767450436f6d6d697453746167654c656e00000000000000000000000000815250610649565b610dff604080519081016040528060138152602001600080516020611996833981519152815250610649565b6040518463ffffffff1660e060020a028152600401808481526020018381526020018281526020019350505050602060405180830381600087803b158015610e4657600080fd5b505af1158015610e5a573d6000803e3d6000fd5b505050506040513d6020811015610e7057600080fd5b5051604080516080810182528281526020818101879052818301869052600b5483518085019094526013845260008051602061199683398151915291840191909152929350916060830191610ecd9190610bef90610ba190610649565b90528051600690815560208083015160075560408301518051610ef492600892019061181b565b50606082015181600301559050507f2f9bf5073932602ad4ef9cc3fc0537d243d84e66ef4beebc84c2e406d72a5e8f83838360405180846000191660001916815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610f77578181015183820152602001610f5f565b50505050905090810190601f168015610fa45780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1505050565b60046020526000908152604090205481565b600154600160a060020a031681565b600a54600160a060020a031681565b6000610ff2611899565b600083815260056020908152604091829020825160808101845281548152600180830180548651600260001994831615610100029490940190911692909204601f81018690048602830186019096528582529194929385810193919291908301828280156110a15780601f10611076576101008083540402835291602001916110a1565b820191906000526020600020905b81548152906001019060200180831161108457829003601f168201915b50505050508152602001600282015481526020016003820154815250509050600081600001511180156111655750600a548151604080517fee684830000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a039092169163ee684830916024808201926020929091908290030181600087803b15801561113857600080fd5b505af115801561114c573d6000803e3d6000fd5b505050506040513d602081101561116257600080fd5b50515b9392505050565b6000818152600560209081526040808320600a54815483517f49403183000000000000000000000000000000000000000000000000000000008152600481019190915292519194600160a060020a03909116936349403183936024808201949293918390030190829087803b1580156111e457600080fd5b505af11580156111f8573d6000803e3d6000fd5b505050506040513d602081101561120e57600080fd5b50511561135057428160020154111561130d576112cb816001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112bc5780601f10611291576101008083540402835291602001916112bc565b820191906000526020600020905b81548152906001019060200180831161129f57829003601f168201915b505050505082600301546116fa565b805460408051848152602081019290925280517fe040346a7ca6935dfd5ccdb81e13933d6af35a399b27c7c61d2888b4960336479281900390910190a161134b565b805460408051848152602081019290925280517f0571dcf79f562f7040389aac4b84570b60ed77c3a1f6d9f10f2e3dc86d647e8f9281900390910190a15b61138e565b805460408051848152602081019290925280517fcbc6fb3892c732a14043baca80213f571ebc1a385c676b25d9907fe8e7a2e37b9281900390910190a15b6000828152600560205260408120818155906113ad60018301826118c2565b506000600282018190556003909101555050565b600a54600654604080517f49403183000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a03909216916349403183916024808201926020929091908290030181600087803b15801561142b57600080fd5b505af115801561143f573d6000803e3d6000fd5b505050506040513d602081101561145557600080fd5b50511561160a5760095442101561154c5760075460029081556008805461148e9260039291610100600182161502600019011604611906565b50600754604080518281526020810182815260088054600261010060018316150260001901909116049383018490527f948095b5aab587891b27052c824e5672ed50ec1f9a79e83170a31ec6a3426ab89493909291906060830190849080156115385780601f1061150d57610100808354040283529160200191611538565b820191906000526020600020905b81548152906001019060200180831161151b57829003601f168201915b5050935050505060405180910390a1611605565b600754604080518281526020810182815260088054600261010060018316150260001901909116049383018490527fc0904308100a434603288b335e40b0e90899de3ab2b8a46c29c6c10a81e15cc99493909291906060830190849080156115f55780601f106115ca576101008083540402835291602001916115f5565b820191906000526020600020905b8154815290600101906020018083116115d857829003601f168201915b5050935050505060405180910390a15b6116c3565b600754604080518281526020810182815260088054600261010060018316150260001901909116049383018490527fd5a88495cdd0eccf33577193df1aee9985321cc18fb46ff43942cfc7c280d2a79493909291906060830190849080156116b35780601f10611688576101008083540402835291602001916116b3565b820191906000526020600020905b81548152906001019060200180831161169657829003601f168201915b5050935050505060405180910390a15b600060068181556007829055906116db6008826118c2565b60038201600090555050565b818101828110156116f457fe5b92915050565b8060046000846040518082805190602001908083835b6020831061172f5780518252601f199092019160209182019101611710565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916600019168152602001908152602001600020819055507f0e92bd4b74871caaf73a4a51ca5ad4f01e5c5215e940a2f2a1f4c755b955066c82826040518080602001838152602001828103825284818151815260200191508051906020019080838360005b838110156117dc5781810151838201526020016117c4565b50505050905090810190601f1680156118095780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061185c57805160ff1916838001178555611889565b82800160010185558215611889579182015b8281111561188957825182559160200191906001019061186e565b5061189592915061197b565b5090565b608060405190810160405280600081526020016060815260200160008152602001600081525090565b50805460018160011615610100020316600290046000825580601f106118e857506104d1565b601f0160209004906000526020600020908101906104d1919061197b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061193f5780548555611889565b8280016001018555821561188957600052602060002091601f016020900482015b82811115611889578254825591600101919060010190611960565b61063791905b8082111561189557600081556001016119815600676f76745052657665616c53746167654c656e00000000000000000000000000a165627a7a72305820e481a4f219ad4512b7f7056c9a8906b2a827b6e18cd902726135554c3a68254e0029`

// DeployGovernmentContract deploys a new Ethereum contract, binding an instance of GovernmentContract to it.
func DeployGovernmentContract(auth *bind.TransactOpts, backend bind.ContractBackend, appellateAddr common.Address, governmentControllerAddr common.Address, plcrAddr common.Address, appealFeeAmount *big.Int, requestAppealLength *big.Int, judgeAppealLength *big.Int, appealSupermajorityPercentage *big.Int, appealChallengeVoteDispensationPct *big.Int, pDeposit *big.Int, pCommitStageLength *big.Int, pRevealStageLength *big.Int, constHash [32]byte, constURI string) (common.Address, *types.Transaction, *GovernmentContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernmentContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernmentContractBin), backend, appellateAddr, governmentControllerAddr, plcrAddr, appealFeeAmount, requestAppealLength, judgeAppealLength, appealSupermajorityPercentage, appealChallengeVoteDispensationPct, pDeposit, pCommitStageLength, pRevealStageLength, constHash, constURI)
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

// PROCESSBY is a free data retrieval call binding the contract method 0x0029514f.
//
// Solidity: function PROCESSBY() constant returns(uint256)
func (_GovernmentContract *GovernmentContractCaller) PROCESSBY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "PROCESSBY")
	return *ret0, err
}

// PROCESSBY is a free data retrieval call binding the contract method 0x0029514f.
//
// Solidity: function PROCESSBY() constant returns(uint256)
func (_GovernmentContract *GovernmentContractSession) PROCESSBY() (*big.Int, error) {
	return _GovernmentContract.Contract.PROCESSBY(&_GovernmentContract.CallOpts)
}

// PROCESSBY is a free data retrieval call binding the contract method 0x0029514f.
//
// Solidity: function PROCESSBY() constant returns(uint256)
func (_GovernmentContract *GovernmentContractCallerSession) PROCESSBY() (*big.Int, error) {
	return _GovernmentContract.Contract.PROCESSBY(&_GovernmentContract.CallOpts)
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

// ConstChangePropCanBeResolved is a free data retrieval call binding the contract method 0xa33c91d8.
//
// Solidity: function constChangePropCanBeResolved() constant returns(bool)
func (_GovernmentContract *GovernmentContractCaller) ConstChangePropCanBeResolved(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "constChangePropCanBeResolved")
	return *ret0, err
}

// ConstChangePropCanBeResolved is a free data retrieval call binding the contract method 0xa33c91d8.
//
// Solidity: function constChangePropCanBeResolved() constant returns(bool)
func (_GovernmentContract *GovernmentContractSession) ConstChangePropCanBeResolved() (bool, error) {
	return _GovernmentContract.Contract.ConstChangePropCanBeResolved(&_GovernmentContract.CallOpts)
}

// ConstChangePropCanBeResolved is a free data retrieval call binding the contract method 0xa33c91d8.
//
// Solidity: function constChangePropCanBeResolved() constant returns(bool)
func (_GovernmentContract *GovernmentContractCallerSession) ConstChangePropCanBeResolved() (bool, error) {
	return _GovernmentContract.Contract.ConstChangePropCanBeResolved(&_GovernmentContract.CallOpts)
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

// PropCanBeResolved is a free data retrieval call binding the contract method 0xffa1bdf0.
//
// Solidity: function propCanBeResolved(_propID bytes32) constant returns(bool)
func (_GovernmentContract *GovernmentContractCaller) PropCanBeResolved(opts *bind.CallOpts, _propID [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "propCanBeResolved", _propID)
	return *ret0, err
}

// PropCanBeResolved is a free data retrieval call binding the contract method 0xffa1bdf0.
//
// Solidity: function propCanBeResolved(_propID bytes32) constant returns(bool)
func (_GovernmentContract *GovernmentContractSession) PropCanBeResolved(_propID [32]byte) (bool, error) {
	return _GovernmentContract.Contract.PropCanBeResolved(&_GovernmentContract.CallOpts, _propID)
}

// PropCanBeResolved is a free data retrieval call binding the contract method 0xffa1bdf0.
//
// Solidity: function propCanBeResolved(_propID bytes32) constant returns(bool)
func (_GovernmentContract *GovernmentContractCallerSession) PropCanBeResolved(_propID [32]byte) (bool, error) {
	return _GovernmentContract.Contract.PropCanBeResolved(&_GovernmentContract.CallOpts, _propID)
}

// PropExists is a free data retrieval call binding the contract method 0x35300990.
//
// Solidity: function propExists(_propID bytes32) constant returns(bool)
func (_GovernmentContract *GovernmentContractCaller) PropExists(opts *bind.CallOpts, _propID [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "propExists", _propID)
	return *ret0, err
}

// PropExists is a free data retrieval call binding the contract method 0x35300990.
//
// Solidity: function propExists(_propID bytes32) constant returns(bool)
func (_GovernmentContract *GovernmentContractSession) PropExists(_propID [32]byte) (bool, error) {
	return _GovernmentContract.Contract.PropExists(&_GovernmentContract.CallOpts, _propID)
}

// PropExists is a free data retrieval call binding the contract method 0x35300990.
//
// Solidity: function propExists(_propID bytes32) constant returns(bool)
func (_GovernmentContract *GovernmentContractCallerSession) PropExists(_propID [32]byte) (bool, error) {
	return _GovernmentContract.Contract.PropExists(&_GovernmentContract.CallOpts, _propID)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals( bytes32) constant returns(pollID uint256, name string, processBy uint256, value uint256)
func (_GovernmentContract *GovernmentContractCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	PollID    *big.Int
	Name      string
	ProcessBy *big.Int
	Value     *big.Int
}, error) {
	ret := new(struct {
		PollID    *big.Int
		Name      string
		ProcessBy *big.Int
		Value     *big.Int
	})
	out := ret
	err := _GovernmentContract.contract.Call(opts, out, "proposals", arg0)
	return *ret, err
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals( bytes32) constant returns(pollID uint256, name string, processBy uint256, value uint256)
func (_GovernmentContract *GovernmentContractSession) Proposals(arg0 [32]byte) (struct {
	PollID    *big.Int
	Name      string
	ProcessBy *big.Int
	Value     *big.Int
}, error) {
	return _GovernmentContract.Contract.Proposals(&_GovernmentContract.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals( bytes32) constant returns(pollID uint256, name string, processBy uint256, value uint256)
func (_GovernmentContract *GovernmentContractCallerSession) Proposals(arg0 [32]byte) (struct {
	PollID    *big.Int
	Name      string
	ProcessBy *big.Int
	Value     *big.Int
}, error) {
	return _GovernmentContract.Contract.Proposals(&_GovernmentContract.CallOpts, arg0)
}

// Voting is a free data retrieval call binding the contract method 0xfce1ccca.
//
// Solidity: function voting() constant returns(address)
func (_GovernmentContract *GovernmentContractCaller) Voting(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernmentContract.contract.Call(opts, out, "voting")
	return *ret0, err
}

// Voting is a free data retrieval call binding the contract method 0xfce1ccca.
//
// Solidity: function voting() constant returns(address)
func (_GovernmentContract *GovernmentContractSession) Voting() (common.Address, error) {
	return _GovernmentContract.Contract.Voting(&_GovernmentContract.CallOpts)
}

// Voting is a free data retrieval call binding the contract method 0xfce1ccca.
//
// Solidity: function voting() constant returns(address)
func (_GovernmentContract *GovernmentContractCallerSession) Voting() (common.Address, error) {
	return _GovernmentContract.Contract.Voting(&_GovernmentContract.CallOpts)
}

// ProcessConstChangeProp is a paid mutator transaction binding the contract method 0x9a99fd11.
//
// Solidity: function processConstChangeProp() returns()
func (_GovernmentContract *GovernmentContractTransactor) ProcessConstChangeProp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernmentContract.contract.Transact(opts, "processConstChangeProp")
}

// ProcessConstChangeProp is a paid mutator transaction binding the contract method 0x9a99fd11.
//
// Solidity: function processConstChangeProp() returns()
func (_GovernmentContract *GovernmentContractSession) ProcessConstChangeProp() (*types.Transaction, error) {
	return _GovernmentContract.Contract.ProcessConstChangeProp(&_GovernmentContract.TransactOpts)
}

// ProcessConstChangeProp is a paid mutator transaction binding the contract method 0x9a99fd11.
//
// Solidity: function processConstChangeProp() returns()
func (_GovernmentContract *GovernmentContractTransactorSession) ProcessConstChangeProp() (*types.Transaction, error) {
	return _GovernmentContract.Contract.ProcessConstChangeProp(&_GovernmentContract.TransactOpts)
}

// ProcessProposal is a paid mutator transaction binding the contract method 0x30490e91.
//
// Solidity: function processProposal(_propID bytes32) returns()
func (_GovernmentContract *GovernmentContractTransactor) ProcessProposal(opts *bind.TransactOpts, _propID [32]byte) (*types.Transaction, error) {
	return _GovernmentContract.contract.Transact(opts, "processProposal", _propID)
}

// ProcessProposal is a paid mutator transaction binding the contract method 0x30490e91.
//
// Solidity: function processProposal(_propID bytes32) returns()
func (_GovernmentContract *GovernmentContractSession) ProcessProposal(_propID [32]byte) (*types.Transaction, error) {
	return _GovernmentContract.Contract.ProcessProposal(&_GovernmentContract.TransactOpts, _propID)
}

// ProcessProposal is a paid mutator transaction binding the contract method 0x30490e91.
//
// Solidity: function processProposal(_propID bytes32) returns()
func (_GovernmentContract *GovernmentContractTransactorSession) ProcessProposal(_propID [32]byte) (*types.Transaction, error) {
	return _GovernmentContract.Contract.ProcessProposal(&_GovernmentContract.TransactOpts, _propID)
}

// ProposeNewConstitution is a paid mutator transaction binding the contract method 0xd704626c.
//
// Solidity: function proposeNewConstitution(_newConstHash bytes32, _newConstURI string) returns()
func (_GovernmentContract *GovernmentContractTransactor) ProposeNewConstitution(opts *bind.TransactOpts, _newConstHash [32]byte, _newConstURI string) (*types.Transaction, error) {
	return _GovernmentContract.contract.Transact(opts, "proposeNewConstitution", _newConstHash, _newConstURI)
}

// ProposeNewConstitution is a paid mutator transaction binding the contract method 0xd704626c.
//
// Solidity: function proposeNewConstitution(_newConstHash bytes32, _newConstURI string) returns()
func (_GovernmentContract *GovernmentContractSession) ProposeNewConstitution(_newConstHash [32]byte, _newConstURI string) (*types.Transaction, error) {
	return _GovernmentContract.Contract.ProposeNewConstitution(&_GovernmentContract.TransactOpts, _newConstHash, _newConstURI)
}

// ProposeNewConstitution is a paid mutator transaction binding the contract method 0xd704626c.
//
// Solidity: function proposeNewConstitution(_newConstHash bytes32, _newConstURI string) returns()
func (_GovernmentContract *GovernmentContractTransactorSession) ProposeNewConstitution(_newConstHash [32]byte, _newConstURI string) (*types.Transaction, error) {
	return _GovernmentContract.Contract.ProposeNewConstitution(&_GovernmentContract.TransactOpts, _newConstHash, _newConstURI)
}

// ProposeReparameterization is a paid mutator transaction binding the contract method 0xbade1c54.
//
// Solidity: function proposeReparameterization(_name string, _value uint256) returns(bytes32)
func (_GovernmentContract *GovernmentContractTransactor) ProposeReparameterization(opts *bind.TransactOpts, _name string, _value *big.Int) (*types.Transaction, error) {
	return _GovernmentContract.contract.Transact(opts, "proposeReparameterization", _name, _value)
}

// ProposeReparameterization is a paid mutator transaction binding the contract method 0xbade1c54.
//
// Solidity: function proposeReparameterization(_name string, _value uint256) returns(bytes32)
func (_GovernmentContract *GovernmentContractSession) ProposeReparameterization(_name string, _value *big.Int) (*types.Transaction, error) {
	return _GovernmentContract.Contract.ProposeReparameterization(&_GovernmentContract.TransactOpts, _name, _value)
}

// ProposeReparameterization is a paid mutator transaction binding the contract method 0xbade1c54.
//
// Solidity: function proposeReparameterization(_name string, _value uint256) returns(bytes32)
func (_GovernmentContract *GovernmentContractTransactorSession) ProposeReparameterization(_name string, _value *big.Int) (*types.Transaction, error) {
	return _GovernmentContract.Contract.ProposeReparameterization(&_GovernmentContract.TransactOpts, _name, _value)
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

// FilterAppellateSet is a free log retrieval operation binding the contract event 0x759a9d1715f38685bd08c7fb25060b7b6795cddf54214336e02a0533c5c7b89e.
//
// Solidity: e _AppellateSet(newAppellate address)
func (_GovernmentContract *GovernmentContractFilterer) FilterAppellateSet(opts *bind.FilterOpts) (*GovernmentContractAppellateSetIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "_AppellateSet")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractAppellateSetIterator{contract: _GovernmentContract.contract, event: "_AppellateSet", logs: logs, sub: sub}, nil
}

// WatchAppellateSet is a free log subscription operation binding the contract event 0x759a9d1715f38685bd08c7fb25060b7b6795cddf54214336e02a0533c5c7b89e.
//
// Solidity: e _AppellateSet(newAppellate address)
func (_GovernmentContract *GovernmentContractFilterer) WatchAppellateSet(opts *bind.WatchOpts, sink chan<- *GovernmentContractAppellateSet) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "_AppellateSet")
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
				if err := _GovernmentContract.contract.UnpackLog(event, "_AppellateSet", log); err != nil {
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

// GovernmentContractGovtReparameterizationProposalIterator is returned from FilterGovtReparameterizationProposal and is used to iterate over the raw logs and unpacked data for GovtReparameterizationProposal events raised by the GovernmentContract contract.
type GovernmentContractGovtReparameterizationProposalIterator struct {
	Event *GovernmentContractGovtReparameterizationProposal // Event containing the contract specifics and raw log

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
func (it *GovernmentContractGovtReparameterizationProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernmentContractGovtReparameterizationProposal)
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
		it.Event = new(GovernmentContractGovtReparameterizationProposal)
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
func (it *GovernmentContractGovtReparameterizationProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernmentContractGovtReparameterizationProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernmentContractGovtReparameterizationProposal represents a GovtReparameterizationProposal event raised by the GovernmentContract contract.
type GovernmentContractGovtReparameterizationProposal struct {
	Name   string
	Value  *big.Int
	PropID [32]byte
	PollID *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGovtReparameterizationProposal is a free log retrieval operation binding the contract event 0x74adf299a4c734e1ae114977ab264221c2f4a914c02243561aaa9158735d3224.
//
// Solidity: e _GovtReparameterizationProposal(name string, value uint256, propID bytes32, pollID uint256)
func (_GovernmentContract *GovernmentContractFilterer) FilterGovtReparameterizationProposal(opts *bind.FilterOpts) (*GovernmentContractGovtReparameterizationProposalIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "_GovtReparameterizationProposal")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractGovtReparameterizationProposalIterator{contract: _GovernmentContract.contract, event: "_GovtReparameterizationProposal", logs: logs, sub: sub}, nil
}

// WatchGovtReparameterizationProposal is a free log subscription operation binding the contract event 0x74adf299a4c734e1ae114977ab264221c2f4a914c02243561aaa9158735d3224.
//
// Solidity: e _GovtReparameterizationProposal(name string, value uint256, propID bytes32, pollID uint256)
func (_GovernmentContract *GovernmentContractFilterer) WatchGovtReparameterizationProposal(opts *bind.WatchOpts, sink chan<- *GovernmentContractGovtReparameterizationProposal) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "_GovtReparameterizationProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernmentContractGovtReparameterizationProposal)
				if err := _GovernmentContract.contract.UnpackLog(event, "_GovtReparameterizationProposal", log); err != nil {
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

// GovernmentContractNewConstProposalIterator is returned from FilterNewConstProposal and is used to iterate over the raw logs and unpacked data for NewConstProposal events raised by the GovernmentContract contract.
type GovernmentContractNewConstProposalIterator struct {
	Event *GovernmentContractNewConstProposal // Event containing the contract specifics and raw log

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
func (it *GovernmentContractNewConstProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernmentContractNewConstProposal)
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
		it.Event = new(GovernmentContractNewConstProposal)
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
func (it *GovernmentContractNewConstProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernmentContractNewConstProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernmentContractNewConstProposal represents a NewConstProposal event raised by the GovernmentContract contract.
type GovernmentContractNewConstProposal struct {
	ProposedHash [32]byte
	ProposedURI  string
	PollID       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewConstProposal is a free log retrieval operation binding the contract event 0x2f9bf5073932602ad4ef9cc3fc0537d243d84e66ef4beebc84c2e406d72a5e8f.
//
// Solidity: e _NewConstProposal(proposedHash bytes32, proposedURI string, pollID uint256)
func (_GovernmentContract *GovernmentContractFilterer) FilterNewConstProposal(opts *bind.FilterOpts) (*GovernmentContractNewConstProposalIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "_NewConstProposal")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractNewConstProposalIterator{contract: _GovernmentContract.contract, event: "_NewConstProposal", logs: logs, sub: sub}, nil
}

// WatchNewConstProposal is a free log subscription operation binding the contract event 0x2f9bf5073932602ad4ef9cc3fc0537d243d84e66ef4beebc84c2e406d72a5e8f.
//
// Solidity: e _NewConstProposal(proposedHash bytes32, proposedURI string, pollID uint256)
func (_GovernmentContract *GovernmentContractFilterer) WatchNewConstProposal(opts *bind.WatchOpts, sink chan<- *GovernmentContractNewConstProposal) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "_NewConstProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernmentContractNewConstProposal)
				if err := _GovernmentContract.contract.UnpackLog(event, "_NewConstProposal", log); err != nil {
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

// GovernmentContractNewConstProposalExpiredIterator is returned from FilterNewConstProposalExpired and is used to iterate over the raw logs and unpacked data for NewConstProposalExpired events raised by the GovernmentContract contract.
type GovernmentContractNewConstProposalExpiredIterator struct {
	Event *GovernmentContractNewConstProposalExpired // Event containing the contract specifics and raw log

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
func (it *GovernmentContractNewConstProposalExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernmentContractNewConstProposalExpired)
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
		it.Event = new(GovernmentContractNewConstProposalExpired)
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
func (it *GovernmentContractNewConstProposalExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernmentContractNewConstProposalExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernmentContractNewConstProposalExpired represents a NewConstProposalExpired event raised by the GovernmentContract contract.
type GovernmentContractNewConstProposalExpired struct {
	ConstHash [32]byte
	ConstURI  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewConstProposalExpired is a free log retrieval operation binding the contract event 0xc0904308100a434603288b335e40b0e90899de3ab2b8a46c29c6c10a81e15cc9.
//
// Solidity: e _NewConstProposalExpired(constHash bytes32, constURI string)
func (_GovernmentContract *GovernmentContractFilterer) FilterNewConstProposalExpired(opts *bind.FilterOpts) (*GovernmentContractNewConstProposalExpiredIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "_NewConstProposalExpired")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractNewConstProposalExpiredIterator{contract: _GovernmentContract.contract, event: "_NewConstProposalExpired", logs: logs, sub: sub}, nil
}

// WatchNewConstProposalExpired is a free log subscription operation binding the contract event 0xc0904308100a434603288b335e40b0e90899de3ab2b8a46c29c6c10a81e15cc9.
//
// Solidity: e _NewConstProposalExpired(constHash bytes32, constURI string)
func (_GovernmentContract *GovernmentContractFilterer) WatchNewConstProposalExpired(opts *bind.WatchOpts, sink chan<- *GovernmentContractNewConstProposalExpired) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "_NewConstProposalExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernmentContractNewConstProposalExpired)
				if err := _GovernmentContract.contract.UnpackLog(event, "_NewConstProposalExpired", log); err != nil {
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

// GovernmentContractNewConstProposalFailedIterator is returned from FilterNewConstProposalFailed and is used to iterate over the raw logs and unpacked data for NewConstProposalFailed events raised by the GovernmentContract contract.
type GovernmentContractNewConstProposalFailedIterator struct {
	Event *GovernmentContractNewConstProposalFailed // Event containing the contract specifics and raw log

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
func (it *GovernmentContractNewConstProposalFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernmentContractNewConstProposalFailed)
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
		it.Event = new(GovernmentContractNewConstProposalFailed)
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
func (it *GovernmentContractNewConstProposalFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernmentContractNewConstProposalFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernmentContractNewConstProposalFailed represents a NewConstProposalFailed event raised by the GovernmentContract contract.
type GovernmentContractNewConstProposalFailed struct {
	ConstHash [32]byte
	ConstURI  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewConstProposalFailed is a free log retrieval operation binding the contract event 0xd5a88495cdd0eccf33577193df1aee9985321cc18fb46ff43942cfc7c280d2a7.
//
// Solidity: e _NewConstProposalFailed(constHash bytes32, constURI string)
func (_GovernmentContract *GovernmentContractFilterer) FilterNewConstProposalFailed(opts *bind.FilterOpts) (*GovernmentContractNewConstProposalFailedIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "_NewConstProposalFailed")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractNewConstProposalFailedIterator{contract: _GovernmentContract.contract, event: "_NewConstProposalFailed", logs: logs, sub: sub}, nil
}

// WatchNewConstProposalFailed is a free log subscription operation binding the contract event 0xd5a88495cdd0eccf33577193df1aee9985321cc18fb46ff43942cfc7c280d2a7.
//
// Solidity: e _NewConstProposalFailed(constHash bytes32, constURI string)
func (_GovernmentContract *GovernmentContractFilterer) WatchNewConstProposalFailed(opts *bind.WatchOpts, sink chan<- *GovernmentContractNewConstProposalFailed) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "_NewConstProposalFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernmentContractNewConstProposalFailed)
				if err := _GovernmentContract.contract.UnpackLog(event, "_NewConstProposalFailed", log); err != nil {
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

// GovernmentContractNewConstProposalPassedIterator is returned from FilterNewConstProposalPassed and is used to iterate over the raw logs and unpacked data for NewConstProposalPassed events raised by the GovernmentContract contract.
type GovernmentContractNewConstProposalPassedIterator struct {
	Event *GovernmentContractNewConstProposalPassed // Event containing the contract specifics and raw log

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
func (it *GovernmentContractNewConstProposalPassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernmentContractNewConstProposalPassed)
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
		it.Event = new(GovernmentContractNewConstProposalPassed)
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
func (it *GovernmentContractNewConstProposalPassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernmentContractNewConstProposalPassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernmentContractNewConstProposalPassed represents a NewConstProposalPassed event raised by the GovernmentContract contract.
type GovernmentContractNewConstProposalPassed struct {
	ConstHash [32]byte
	ConstURI  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewConstProposalPassed is a free log retrieval operation binding the contract event 0x948095b5aab587891b27052c824e5672ed50ec1f9a79e83170a31ec6a3426ab8.
//
// Solidity: e _NewConstProposalPassed(constHash bytes32, constURI string)
func (_GovernmentContract *GovernmentContractFilterer) FilterNewConstProposalPassed(opts *bind.FilterOpts) (*GovernmentContractNewConstProposalPassedIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "_NewConstProposalPassed")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractNewConstProposalPassedIterator{contract: _GovernmentContract.contract, event: "_NewConstProposalPassed", logs: logs, sub: sub}, nil
}

// WatchNewConstProposalPassed is a free log subscription operation binding the contract event 0x948095b5aab587891b27052c824e5672ed50ec1f9a79e83170a31ec6a3426ab8.
//
// Solidity: e _NewConstProposalPassed(constHash bytes32, constURI string)
func (_GovernmentContract *GovernmentContractFilterer) WatchNewConstProposalPassed(opts *bind.WatchOpts, sink chan<- *GovernmentContractNewConstProposalPassed) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "_NewConstProposalPassed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernmentContractNewConstProposalPassed)
				if err := _GovernmentContract.contract.UnpackLog(event, "_NewConstProposalPassed", log); err != nil {
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

// FilterParameterSet is a free log retrieval operation binding the contract event 0x0e92bd4b74871caaf73a4a51ca5ad4f01e5c5215e940a2f2a1f4c755b955066c.
//
// Solidity: e _ParameterSet(name string, value uint256)
func (_GovernmentContract *GovernmentContractFilterer) FilterParameterSet(opts *bind.FilterOpts) (*GovernmentContractParameterSetIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "_ParameterSet")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractParameterSetIterator{contract: _GovernmentContract.contract, event: "_ParameterSet", logs: logs, sub: sub}, nil
}

// WatchParameterSet is a free log subscription operation binding the contract event 0x0e92bd4b74871caaf73a4a51ca5ad4f01e5c5215e940a2f2a1f4c755b955066c.
//
// Solidity: e _ParameterSet(name string, value uint256)
func (_GovernmentContract *GovernmentContractFilterer) WatchParameterSet(opts *bind.WatchOpts, sink chan<- *GovernmentContractParameterSet) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "_ParameterSet")
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
				if err := _GovernmentContract.contract.UnpackLog(event, "_ParameterSet", log); err != nil {
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

// GovernmentContractProposalExpiredIterator is returned from FilterProposalExpired and is used to iterate over the raw logs and unpacked data for ProposalExpired events raised by the GovernmentContract contract.
type GovernmentContractProposalExpiredIterator struct {
	Event *GovernmentContractProposalExpired // Event containing the contract specifics and raw log

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
func (it *GovernmentContractProposalExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernmentContractProposalExpired)
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
		it.Event = new(GovernmentContractProposalExpired)
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
func (it *GovernmentContractProposalExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernmentContractProposalExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernmentContractProposalExpired represents a ProposalExpired event raised by the GovernmentContract contract.
type GovernmentContractProposalExpired struct {
	PropId [32]byte
	PollID *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProposalExpired is a free log retrieval operation binding the contract event 0x0571dcf79f562f7040389aac4b84570b60ed77c3a1f6d9f10f2e3dc86d647e8f.
//
// Solidity: e _ProposalExpired(propId bytes32, pollID uint256)
func (_GovernmentContract *GovernmentContractFilterer) FilterProposalExpired(opts *bind.FilterOpts) (*GovernmentContractProposalExpiredIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "_ProposalExpired")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractProposalExpiredIterator{contract: _GovernmentContract.contract, event: "_ProposalExpired", logs: logs, sub: sub}, nil
}

// WatchProposalExpired is a free log subscription operation binding the contract event 0x0571dcf79f562f7040389aac4b84570b60ed77c3a1f6d9f10f2e3dc86d647e8f.
//
// Solidity: e _ProposalExpired(propId bytes32, pollID uint256)
func (_GovernmentContract *GovernmentContractFilterer) WatchProposalExpired(opts *bind.WatchOpts, sink chan<- *GovernmentContractProposalExpired) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "_ProposalExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernmentContractProposalExpired)
				if err := _GovernmentContract.contract.UnpackLog(event, "_ProposalExpired", log); err != nil {
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

// GovernmentContractProposalFailedIterator is returned from FilterProposalFailed and is used to iterate over the raw logs and unpacked data for ProposalFailed events raised by the GovernmentContract contract.
type GovernmentContractProposalFailedIterator struct {
	Event *GovernmentContractProposalFailed // Event containing the contract specifics and raw log

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
func (it *GovernmentContractProposalFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernmentContractProposalFailed)
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
		it.Event = new(GovernmentContractProposalFailed)
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
func (it *GovernmentContractProposalFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernmentContractProposalFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernmentContractProposalFailed represents a ProposalFailed event raised by the GovernmentContract contract.
type GovernmentContractProposalFailed struct {
	PropId [32]byte
	PollID *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProposalFailed is a free log retrieval operation binding the contract event 0xcbc6fb3892c732a14043baca80213f571ebc1a385c676b25d9907fe8e7a2e37b.
//
// Solidity: e _ProposalFailed(propId bytes32, pollID uint256)
func (_GovernmentContract *GovernmentContractFilterer) FilterProposalFailed(opts *bind.FilterOpts) (*GovernmentContractProposalFailedIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "_ProposalFailed")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractProposalFailedIterator{contract: _GovernmentContract.contract, event: "_ProposalFailed", logs: logs, sub: sub}, nil
}

// WatchProposalFailed is a free log subscription operation binding the contract event 0xcbc6fb3892c732a14043baca80213f571ebc1a385c676b25d9907fe8e7a2e37b.
//
// Solidity: e _ProposalFailed(propId bytes32, pollID uint256)
func (_GovernmentContract *GovernmentContractFilterer) WatchProposalFailed(opts *bind.WatchOpts, sink chan<- *GovernmentContractProposalFailed) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "_ProposalFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernmentContractProposalFailed)
				if err := _GovernmentContract.contract.UnpackLog(event, "_ProposalFailed", log); err != nil {
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

// GovernmentContractProposalPassedIterator is returned from FilterProposalPassed and is used to iterate over the raw logs and unpacked data for ProposalPassed events raised by the GovernmentContract contract.
type GovernmentContractProposalPassedIterator struct {
	Event *GovernmentContractProposalPassed // Event containing the contract specifics and raw log

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
func (it *GovernmentContractProposalPassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernmentContractProposalPassed)
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
		it.Event = new(GovernmentContractProposalPassed)
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
func (it *GovernmentContractProposalPassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernmentContractProposalPassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernmentContractProposalPassed represents a ProposalPassed event raised by the GovernmentContract contract.
type GovernmentContractProposalPassed struct {
	PropId [32]byte
	PollID *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProposalPassed is a free log retrieval operation binding the contract event 0xe040346a7ca6935dfd5ccdb81e13933d6af35a399b27c7c61d2888b496033647.
//
// Solidity: e _ProposalPassed(propId bytes32, pollID uint256)
func (_GovernmentContract *GovernmentContractFilterer) FilterProposalPassed(opts *bind.FilterOpts) (*GovernmentContractProposalPassedIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "_ProposalPassed")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractProposalPassedIterator{contract: _GovernmentContract.contract, event: "_ProposalPassed", logs: logs, sub: sub}, nil
}

// WatchProposalPassed is a free log subscription operation binding the contract event 0xe040346a7ca6935dfd5ccdb81e13933d6af35a399b27c7c61d2888b496033647.
//
// Solidity: e _ProposalPassed(propId bytes32, pollID uint256)
func (_GovernmentContract *GovernmentContractFilterer) WatchProposalPassed(opts *bind.WatchOpts, sink chan<- *GovernmentContractProposalPassed) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "_ProposalPassed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernmentContractProposalPassed)
				if err := _GovernmentContract.contract.UnpackLog(event, "_ProposalPassed", log); err != nil {
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
