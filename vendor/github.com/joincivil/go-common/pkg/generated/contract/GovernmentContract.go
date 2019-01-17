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
const GovernmentContractBin = `0x608060405262093a80600b553480156200001857600080fd5b50604051620023473803806200234783398101604090815281516020830151918301516060840151608085015160a086015160c087015160e08801516101008901516101208a01516101408b01516101608c01516101808d01519a9c999a98999798969795969495939492939192909101600160a060020a038d1615156200010157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f617070656c6c6174654164647220616464726573732069732030000000000000604482015290519081900360640190fd5b600160a060020a038c1615156200017957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f676f7665726e6d656e74436f6e74726f6c6c6572416464722069732030000000604482015290519081900360640190fd5b8c6000806101000a815481600160a060020a030219169083600160a060020a031602179055508b600160006101000a815481600160a060020a030219169083600160a060020a031602179055508a600a60006101000a815481600160a060020a030219169083600160a060020a031602179055506200023d6040805190810160405280601081526020017f7265717565737441707065616c4c656e000000000000000000000000000000008152508a6200044f640100000000026401000000009004565b60408051808201909152600e81527f6a7564676541707065616c4c656e00000000000000000000000000000000000060208201526200028690896401000000006200044f810204565b60408051808201909152600981527f61707065616c46656500000000000000000000000000000000000000000000006020820152620002cf908b6401000000006200044f810204565b60408051808201909152601481527f61707065616c566f746550657263656e7461676500000000000000000000000060208201526200031890886401000000006200044f810204565b6200038f606060405190810160405280602281526020017f61707065616c4368616c6c656e6765566f746544697370656e736174696f6e5081526020017f6374000000000000000000000000000000000000000000000000000000000000815250876200044f640100000000026401000000009004565b60408051808201909152601381527f676f767450436f6d6d697453746167654c656e000000000000000000000000006020820152620003d890856401000000006200044f810204565b60408051808201909152601381527f676f76745052657665616c53746167654c656e0000000000000000000000000060208201526200042190846401000000006200044f810204565b600282905580516200043b90600390602084019062000575565b50505050505050505050505050506200061a565b8060046000846040518082805190602001908083835b60208310620004865780518252601f19909201916020918201910162000465565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916600019168152602001908152602001600020819055507f0e92bd4b74871caaf73a4a51ca5ad4f01e5c5215e940a2f2a1f4c755b955066c82826040518080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015620005355781810151838201526020016200051b565b50505050905090810190601f168015620005635780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620005b857805160ff1916838001178555620005e8565b82800160010185558215620005e8579182015b82811115620005e8578251825591602001919060010190620005cb565b50620005f6929150620005fa565b5090565b6200061791905b80821115620005f6576000815560010162000601565b90565b611d1d806200062a6000396000f3006080604052600436106100ec5763ffffffff60e060020a60003504166229514f81146100f157806330490e911461011857806332ed5b121461013257806335300990146101d7578063551224251461020357806356e1fb88146102245780635793b9cf14610255578063693ec85e1461026a5780638ca7f51c146102c35780639a99fd111461034d578063a33c91d814610362578063bade1c5414610377578063c7d93fd4146103d2578063d5fd9e66146103e7578063d704626c146103fc578063dc6ab5271461045a578063f2a2129b14610472578063fce1ccca14610487578063ffa1bdf01461049c575b600080fd5b3480156100fd57600080fd5b506101066104b4565b60408051918252519081900360200190f35b34801561012457600080fd5b506101306004356104ba565b005b34801561013e57600080fd5b5061014a6004356104d4565b6040518085815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b83811015610199578181015183820152602001610181565b50505050905090810190601f1680156101c65780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b3480156101e357600080fd5b506101ef600435610587565b604080519115158252519081900360200190f35b34801561020f57600080fd5b50610130600160a060020a036004351661059d565b34801561023057600080fd5b5061023961070c565b60408051600160a060020a039092168252519081900360200190f35b34801561026157600080fd5b5061023961071c565b34801561027657600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261010694369492936024939284019190819084018382808284375094975061072b9650505050505050565b3480156102cf57600080fd5b506102d86107a2565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103125781810151838201526020016102fa565b50505050905090810190601f16801561033f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561035957600080fd5b50610130610830565b34801561036e57600080fd5b506101ef610847565b34801561038357600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261010694369492936024939284019190819084018382808284375094975050933594506108f49350505050565b3480156103de57600080fd5b50610106610f89565b3480156103f357600080fd5b50610239610f8f565b34801561040857600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610130958335953695604494919390910191908190840183828082843750949750610f9e9650505050505050565b34801561046657600080fd5b506101066004356112f4565b34801561047e57600080fd5b50610239611306565b34801561049357600080fd5b50610239611315565b3480156104a857600080fd5b506101ef600435611324565b600b5481565b6104c381611324565b156100ec576104d1816114a8565b50565b6005602090815260009182526040918290208054600180830180548651600293821615610100026000190190911692909204601f8101869004860283018601909652858252919492939092908301828280156105715780601f1061054657610100808354040283529160200191610571565b820191906000526020600020905b81548152906001019060200180831161055457829003601f168201915b5050505050908060020154908060030154905084565b6000908152600560205260408120600201541190565b600154600160a060020a03163314610625576040805160e560020a62461bcd02815260206004820152602360248201527f53656e646572206973206e6f7420476f7665726e6d656e7420436f6e74726f6c60448201527f6c65720000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a03811615156106ab576040805160e560020a62461bcd02815260206004820152602260248201527f6e6577417070656c6c6174652061646472657373206d757374206e6f7420626560448201527f2030000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f759a9d1715f38685bd08c7fb25060b7b6795cddf54214336e02a0533c5c7b89e9181900360200190a150565b600054600160a060020a03165b90565b600154600160a060020a031690565b600060046000836040518082805190602001908083835b602083106107615780518252601f199092019160209182019101610742565b51815160209384036101000a600019018019909216911617905260408051929094018290039091208652850195909552929092016000205495945050505050565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108285780601f106107fd57610100808354040283529160200191610828565b820191906000526020600020905b81548152906001019060200180831161080b57829003601f168201915b505050505081565b610838610847565b156100ec576108456116fd565b565b6000806006600001541180156108ef5750600a54600654604080517fee684830000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a039092169163ee684830916024808201926020929091908290030181600087803b1580156108c257600080fd5b505af11580156108d6573d6000803e3d6000fd5b505050506040513d60208110156108ec57600080fd5b50515b905090565b6000805481908190600160a060020a0316331461095b576040805160e560020a62461bcd02815260206004820152601760248201527f53656e646572206973206e6f7420417070656c6c617465000000000000000000604482015290519081900360640190fd5b84846040518083805190602001908083835b6020831061098c5780518252601f19909201916020918201910161096d565b51815160001960209485036101000a019081169019919091161790529201938452506040805193849003820184207f61707065616c566f746550657263656e74616765000000000000000000000000855290519384900360140184208a5191975094508993925082918401908083835b60208310610a1b5780518252601f1990920191602091820191016109fc565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020600019161480610b115750604080517f61707065616c4368616c6c656e6765566f746544697370656e736174696f6e5081527f637400000000000000000000000000000000000000000000000000000000000060208083019190915291519081900360220181208751909288929182918401908083835b60208310610ade5780518252601f199092019160209182019101610abf565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916145b15610b95576064841115610b95576040805160e560020a62461bcd02815260206004820152603760248201527f50657263656e7461676520706172616d6574657273206d757374206265206c6560448201527f7373207468616e206f7220657175616c20746f20313030000000000000000000606482015290519081900360840190fd5b610b9e82610587565b15610c19576040805160e560020a62461bcd02815260206004820152602a60248201527f50726f706f736564207265706172616d65746572697a6174696f6e20616c726560448201527f6164792065786973747300000000000000000000000000000000000000000000606482015290519081900360840190fd5b83610c238661072b565b1415610c9f576040805160e560020a62461bcd02815260206004820152603c60248201527f50726f706f736564207265706172616d65746572697a6174696f6e20776f756c60448201527f64206e6f74206368616e676520706172616d657465722076616c756500000000606482015290519081900360840190fd5b600a5460408051808201909152601481527f61707065616c566f746550657263656e746167650000000000000000000000006020820152600160a060020a03909116906332ed3d6090610cf19061072b565b610d2f6040805190810160405280601381526020017f676f767450436f6d6d697453746167654c656e0000000000000000000000000081525061072b565b610d5b604080519081016040528060138152602001600080516020611cd283398151915281525061072b565b6040518463ffffffff1660e060020a028152600401808481526020018381526020018281526020019350505050602060405180830381600087803b158015610da257600080fd5b505af1158015610db6573d6000803e3d6000fd5b505050506040513d6020811015610dcc57600080fd5b5051604080516080810182528281526020818101899052600b548351808501855260138152600080516020611cd283398151915292810192909252939450909291830191610e7b91610e6f90610e219061072b565b610e6f610e626040805190810160405280601381526020017f676f767450436f6d6d697453746167654c656e0000000000000000000000000081525061072b565b429063ffffffff611a2316565b9063ffffffff611a2316565b8152602090810186905260008481526005825260409020825181558282015180519192610eb092600185019290910190611b57565b5060408201518160020155606082015181600301559050507f74adf299a4c734e1ae114977ab264221c2f4a914c02243561aaa9158735d32248585848460405180806020018581526020018460001916600019168152602001838152602001828103825286818151815260200191508051906020019080838360005b83811015610f44578181015183820152602001610f2c565b50505050905090810190601f168015610f715780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1509392505050565b60025481565b600054600160a060020a031681565b60008054600160a060020a03163314611001576040805160e560020a62461bcd02815260206004820152601760248201527f53656e646572206973206e6f7420417070656c6c617465000000000000000000604482015290519081900360640190fd5b6009541561107f576040805160e560020a62461bcd02815260206004820152602860248201527f4e657720436f6e737469747574696f6e2070726f706f73616c20616c7265616460448201527f7920616374697665000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600a5460408051808201909152601481527f61707065616c566f746550657263656e746167650000000000000000000000006020820152600160a060020a03909116906332ed3d60906110d19061072b565b61110f6040805190810160405280601381526020017f676f767450436f6d6d697453746167654c656e0000000000000000000000000081525061072b565b61113b604080519081016040528060138152602001600080516020611cd283398151915281525061072b565b6040518463ffffffff1660e060020a028152600401808481526020018381526020018281526020019350505050602060405180830381600087803b15801561118257600080fd5b505af1158015611196573d6000803e3d6000fd5b505050506040513d60208110156111ac57600080fd5b5051604080516080810182528281526020818101879052818301869052600b54835180850190945260138452600080516020611cd2833981519152918401919091529293509160608301916112099190610e6f90610e219061072b565b90528051600690815560208083015160075560408301518051611230926008920190611b57565b50606082015181600301559050507f2f9bf5073932602ad4ef9cc3fc0537d243d84e66ef4beebc84c2e406d72a5e8f83838360405180846000191660001916815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156112b357818101518382015260200161129b565b50505050905090810190601f1680156112e05780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1505050565b60046020526000908152604090205481565b600154600160a060020a031681565b600a54600160a060020a031681565b600061132e611bd5565b600083815260056020908152604091829020825160808101845281548152600180830180548651600260001994831615610100029490940190911692909204601f81018690048602830186019096528582529194929385810193919291908301828280156113dd5780601f106113b2576101008083540402835291602001916113dd565b820191906000526020600020905b8154815290600101906020018083116113c057829003601f168201915b50505050508152602001600282015481526020016003820154815250509050600081600001511180156114a15750600a548151604080517fee684830000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a039092169163ee684830916024808201926020929091908290030181600087803b15801561147457600080fd5b505af1158015611488573d6000803e3d6000fd5b505050506040513d602081101561149e57600080fd5b50515b9392505050565b6000818152600560209081526040808320600a54815483517f49403183000000000000000000000000000000000000000000000000000000008152600481019190915292519194600160a060020a03909116936349403183936024808201949293918390030190829087803b15801561152057600080fd5b505af1158015611534573d6000803e3d6000fd5b505050506040513d602081101561154a57600080fd5b50511561168c57428160020154111561164957611607816001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156115f85780601f106115cd576101008083540402835291602001916115f8565b820191906000526020600020905b8154815290600101906020018083116115db57829003601f168201915b50505050508260030154611a36565b805460408051848152602081019290925280517fe040346a7ca6935dfd5ccdb81e13933d6af35a399b27c7c61d2888b4960336479281900390910190a1611687565b805460408051848152602081019290925280517f0571dcf79f562f7040389aac4b84570b60ed77c3a1f6d9f10f2e3dc86d647e8f9281900390910190a15b6116ca565b805460408051848152602081019290925280517fcbc6fb3892c732a14043baca80213f571ebc1a385c676b25d9907fe8e7a2e37b9281900390910190a15b6000828152600560205260408120818155906116e96001830182611bfe565b506000600282018190556003909101555050565b600a54600654604080517f49403183000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a03909216916349403183916024808201926020929091908290030181600087803b15801561176757600080fd5b505af115801561177b573d6000803e3d6000fd5b505050506040513d602081101561179157600080fd5b50511561194657600954421015611888576007546002908155600880546117ca9260039291610100600182161502600019011604611c42565b50600754604080518281526020810182815260088054600261010060018316150260001901909116049383018490527f948095b5aab587891b27052c824e5672ed50ec1f9a79e83170a31ec6a3426ab89493909291906060830190849080156118745780601f1061184957610100808354040283529160200191611874565b820191906000526020600020905b81548152906001019060200180831161185757829003601f168201915b5050935050505060405180910390a1611941565b600754604080518281526020810182815260088054600261010060018316150260001901909116049383018490527fc0904308100a434603288b335e40b0e90899de3ab2b8a46c29c6c10a81e15cc99493909291906060830190849080156119315780601f1061190657610100808354040283529160200191611931565b820191906000526020600020905b81548152906001019060200180831161191457829003601f168201915b5050935050505060405180910390a15b6119ff565b600754604080518281526020810182815260088054600261010060018316150260001901909116049383018490527fd5a88495cdd0eccf33577193df1aee9985321cc18fb46ff43942cfc7c280d2a79493909291906060830190849080156119ef5780601f106119c4576101008083540402835291602001916119ef565b820191906000526020600020905b8154815290600101906020018083116119d257829003601f168201915b5050935050505060405180910390a15b60006006818155600782905590611a17600882611bfe565b60038201600090555050565b81810182811015611a3057fe5b92915050565b8060046000846040518082805190602001908083835b60208310611a6b5780518252601f199092019160209182019101611a4c565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916600019168152602001908152602001600020819055507f0e92bd4b74871caaf73a4a51ca5ad4f01e5c5215e940a2f2a1f4c755b955066c82826040518080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015611b18578181015183820152602001611b00565b50505050905090810190601f168015611b455780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611b9857805160ff1916838001178555611bc5565b82800160010185558215611bc5579182015b82811115611bc5578251825591602001919060010190611baa565b50611bd1929150611cb7565b5090565b608060405190810160405280600081526020016060815260200160008152602001600081525090565b50805460018160011615610100020316600290046000825580601f10611c2457506104d1565b601f0160209004906000526020600020908101906104d19190611cb7565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611c7b5780548555611bc5565b82800160010185558215611bc557600052602060002091601f016020900482015b82811115611bc5578254825591600101919060010190611c9c565b61071991905b80821115611bd15760008155600101611cbd5600676f76745052657665616c53746167654c656e00000000000000000000000000a165627a7a72305820287bcc8c71e3883565ce7be0e14218da0938e9b4ffb61093c3599f35a4dd06300029`

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
