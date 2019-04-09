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
const GovernmentContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"PROCESSBY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposals\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"processBy\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"constitutionURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"constitutionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appellate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"params\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governmentController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voting\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"appellateAddr\",\"type\":\"address\"},{\"name\":\"governmentControllerAddr\",\"type\":\"address\"},{\"name\":\"plcrAddr\",\"type\":\"address\"},{\"name\":\"appealFeeAmount\",\"type\":\"uint256\"},{\"name\":\"requestAppealLength\",\"type\":\"uint256\"},{\"name\":\"judgeAppealLength\",\"type\":\"uint256\"},{\"name\":\"appealSupermajorityPercentage\",\"type\":\"uint256\"},{\"name\":\"appealChallengeVoteDispensationPct\",\"type\":\"uint256\"},{\"name\":\"pDeposit\",\"type\":\"uint256\"},{\"name\":\"pCommitStageLength\",\"type\":\"uint256\"},{\"name\":\"pRevealStageLength\",\"type\":\"uint256\"},{\"name\":\"constHash\",\"type\":\"bytes32\"},{\"name\":\"constURI\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAppellate\",\"type\":\"address\"}],\"name\":\"_AppellateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"_ParameterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"propID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"pollID\",\"type\":\"uint256\"}],\"name\":\"_GovtReparameterizationProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"propId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"pollID\",\"type\":\"uint256\"}],\"name\":\"_ProposalPassed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"propId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"pollID\",\"type\":\"uint256\"}],\"name\":\"_ProposalExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"propId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"pollID\",\"type\":\"uint256\"}],\"name\":\"_ProposalFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proposedHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"proposedURI\",\"type\":\"string\"}],\"name\":\"_NewConstSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAppellate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGovernmentController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"proposeReparameterization\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newConstHash\",\"type\":\"bytes32\"},{\"name\":\"_newConstURI\",\"type\":\"string\"}],\"name\":\"setNewConstitution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_propID\",\"type\":\"bytes32\"}],\"name\":\"processProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_propID\",\"type\":\"bytes32\"}],\"name\":\"propExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_propID\",\"type\":\"bytes32\"}],\"name\":\"propCanBeResolved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAppellate\",\"type\":\"address\"}],\"name\":\"setAppellate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GovernmentContractBin is the compiled bytecode used for deploying new contracts.
const GovernmentContractBin = `0x608060405262093a806007553480156200001857600080fd5b5060405162001c6638038062001c6683398101604090815281516020830151918301516060840151608085015160a086015160c087015160e08801516101008901516101208a01516101408b01516101608c01516101808d01519a9c999a98999798969795969495939492939192909101600160a060020a038d1615156200010157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f617070656c6c6174654164647220616464726573732069732030000000000000604482015290519081900360640190fd5b600160a060020a038c1615156200017957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f676f7665726e6d656e74436f6e74726f6c6c6572416464722069732030000000604482015290519081900360640190fd5b8c6000806101000a815481600160a060020a030219169083600160a060020a031602179055508b600160006101000a815481600160a060020a030219169083600160a060020a031602179055508a600660006101000a815481600160a060020a030219169083600160a060020a031602179055506200023d6040805190810160405280601081526020017f7265717565737441707065616c4c656e000000000000000000000000000000008152508a6200044f640100000000026401000000009004565b60408051808201909152600e81527f6a7564676541707065616c4c656e00000000000000000000000000000000000060208201526200028690896401000000006200044f810204565b60408051808201909152600981527f61707065616c46656500000000000000000000000000000000000000000000006020820152620002cf908b6401000000006200044f810204565b60408051808201909152601481527f61707065616c566f746550657263656e7461676500000000000000000000000060208201526200031890886401000000006200044f810204565b6200038f606060405190810160405280602281526020017f61707065616c4368616c6c656e6765566f746544697370656e736174696f6e5081526020017f6374000000000000000000000000000000000000000000000000000000000000815250876200044f640100000000026401000000009004565b60408051808201909152601381527f676f767450436f6d6d697453746167654c656e000000000000000000000000006020820152620003d890856401000000006200044f810204565b60408051808201909152601381527f676f76745052657665616c53746167654c656e0000000000000000000000000060208201526200042190846401000000006200044f810204565b600282905580516200043b90600390602084019062000575565b50505050505050505050505050506200061a565b8060046000846040518082805190602001908083835b60208310620004865780518252601f19909201916020918201910162000465565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916600019168152602001908152602001600020819055507f0e92bd4b74871caaf73a4a51ca5ad4f01e5c5215e940a2f2a1f4c755b955066c82826040518080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015620005355781810151838201526020016200051b565b50505050905090810190601f168015620005635780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620005b857805160ff1916838001178555620005e8565b82800160010185558215620005e8579182015b82811115620005e8578251825591602001919060010190620005cb565b50620005f6929150620005fa565b5090565b6200061791905b80821115620005f6576000815560010162000601565b90565b61163c806200062a6000396000f3006080604052600436106100ef5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166229514f81146100f457806330490e911461011b57806332ed5b121461013557806335300990146101da578063551224251461020657806356e1fb88146102275780635793b9cf14610258578063693ec85e1461026d5780638ca7f51c146102c6578063b0924d6e14610350578063bade1c54146103ae578063c7d93fd414610409578063d5fd9e661461041e578063dc6ab52714610433578063f2a2129b1461044b578063fce1ccca14610460578063ffa1bdf014610475575b600080fd5b34801561010057600080fd5b5061010961048d565b60408051918252519081900360200190f35b34801561012757600080fd5b50610133600435610493565b005b34801561014157600080fd5b5061014d6004356104ad565b6040518085815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561019c578181015183820152602001610184565b50505050905090810190601f1680156101c95780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b3480156101e657600080fd5b506101f2600435610560565b604080519115158252519081900360200190f35b34801561021257600080fd5b50610133600160a060020a0360043516610576565b34801561023357600080fd5b5061023c6106e5565b60408051600160a060020a039092168252519081900360200190f35b34801561026457600080fd5b5061023c6106f5565b34801561027957600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101099436949293602493928401919081908401838280828437509497506107049650505050505050565b3480156102d257600080fd5b506102db61077b565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103155781810151838201526020016102fd565b50505050905090810190601f1680156103425780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561035c57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101339583359536956044949193909101919081908401838280828437509497506108099650505050505050565b3480156103ba57600080fd5b506040805160206004803580820135601f8101849004840285018401909552848452610109943694929360249392840191908190840183828082843750949750509335945061092b9350505050565b34801561041557600080fd5b50610109610ffd565b34801561042a57600080fd5b5061023c611003565b34801561043f57600080fd5b50610109600435611012565b34801561045757600080fd5b5061023c611024565b34801561046c57600080fd5b5061023c611033565b34801561048157600080fd5b506101f2600435611042565b60075481565b61049c81611042565b156100ef576104aa816111c6565b50565b6005602090815260009182526040918290208054600180830180548651600293821615610100026000190190911692909204601f81018690048602830186019096528582529194929390929083018282801561054a5780601f1061051f5761010080835404028352916020019161054a565b820191906000526020600020905b81548152906001019060200180831161052d57829003601f168201915b5050505050908060020154908060030154905084565b6000908152600560205260408120600201541190565b600154600160a060020a031633146105fe576040805160e560020a62461bcd02815260206004820152602360248201527f53656e646572206973206e6f7420476f7665726e6d656e7420436f6e74726f6c60448201527f6c65720000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0381161515610684576040805160e560020a62461bcd02815260206004820152602260248201527f6e6577417070656c6c6174652061646472657373206d757374206e6f7420626560448201527f2030000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f759a9d1715f38685bd08c7fb25060b7b6795cddf54214336e02a0533c5c7b89e9181900360200190a150565b600054600160a060020a03165b90565b600154600160a060020a031690565b600060046000836040518082805190602001908083835b6020831061073a5780518252601f19909201916020918201910161071b565b51815160209384036101000a600019018019909216911617905260408051929094018290039091208652850195909552929092016000205495945050505050565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108015780601f106107d657610100808354040283529160200191610801565b820191906000526020600020905b8154815290600101906020018083116107e457829003601f168201915b505050505081565b600054600160a060020a0316331461086b576040805160e560020a62461bcd02815260206004820152601760248201527f53656e646572206973206e6f7420417070656c6c617465000000000000000000604482015290519081900360640190fd5b6002829055805161088390600390602084019061150f565b5060408051838152602080820183815284519383019390935283517f2f6679e95449d4806445cd50a14e77e4b83ea193ae84e30f8a3247436442c25593869386939092606084019185019080838360005b838110156108ec5781810151838201526020016108d4565b50505050905090810190601f1680156109195780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b6000805481908190600160a060020a03163314610992576040805160e560020a62461bcd02815260206004820152601760248201527f53656e646572206973206e6f7420417070656c6c617465000000000000000000604482015290519081900360640190fd5b84846040518083805190602001908083835b602083106109c35780518252601f1990920191602091820191016109a4565b51815160001960209485036101000a019081169019919091161790529201938452506040805193849003820184207f61707065616c566f746550657263656e74616765000000000000000000000000855290519384900360140184208a5191975094508993925082918401908083835b60208310610a525780518252601f199092019160209182019101610a33565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020600019161480610b485750604080517f61707065616c4368616c6c656e6765566f746544697370656e736174696f6e5081527f637400000000000000000000000000000000000000000000000000000000000060208083019190915291519081900360220181208751909288929182918401908083835b60208310610b155780518252601f199092019160209182019101610af6565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916145b15610bcc576064841115610bcc576040805160e560020a62461bcd02815260206004820152603760248201527f50657263656e7461676520706172616d6574657273206d757374206265206c6560448201527f7373207468616e206f7220657175616c20746f20313030000000000000000000606482015290519081900360840190fd5b610bd582610560565b15610c50576040805160e560020a62461bcd02815260206004820152602a60248201527f50726f706f736564207265706172616d65746572697a6174696f6e20616c726560448201527f6164792065786973747300000000000000000000000000000000000000000000606482015290519081900360840190fd5b83610c5a86610704565b1415610cd6576040805160e560020a62461bcd02815260206004820152603c60248201527f50726f706f736564207265706172616d65746572697a6174696f6e20776f756c60448201527f64206e6f74206368616e676520706172616d657465722076616c756500000000606482015290519081900360840190fd5b60065460408051808201909152601481527f61707065616c566f746550657263656e746167650000000000000000000000006020820152600160a060020a03909116906332ed3d6090610d2890610704565b610d666040805190810160405280601381526020017f676f767450436f6d6d697453746167654c656e00000000000000000000000000815250610704565b610da46040805190810160405280601381526020017f676f76745052657665616c53746167654c656e00000000000000000000000000815250610704565b6040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808481526020018381526020018281526020019350505050602060405180830381600087803b158015610e0457600080fd5b505af1158015610e18573d6000803e3d6000fd5b505050506040513d6020811015610e2e57600080fd5b505160408051608081018252828152602081810189905260075483518085018552601381527f676f76745052657665616c53746167654c656e0000000000000000000000000092810192909252939450909291830191610eef91610ee390610e9590610704565b610ee3610ed66040805190810160405280601381526020017f676f767450436f6d6d697453746167654c656e00000000000000000000000000815250610704565b429063ffffffff61141b16565b9063ffffffff61141b16565b8152602090810186905260008481526005825260409020825181558282015180519192610f249260018501929091019061150f565b5060408201518160020155606082015181600301559050507f74adf299a4c734e1ae114977ab264221c2f4a914c02243561aaa9158735d32248585848460405180806020018581526020018460001916600019168152602001838152602001828103825286818151815260200191508051906020019080838360005b83811015610fb8578181015183820152602001610fa0565b50505050905090810190601f168015610fe55780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1509392505050565b60025481565b600054600160a060020a031681565b60046020526000908152604090205481565b600154600160a060020a031681565b600654600160a060020a031681565b600061104c61158d565b600083815260056020908152604091829020825160808101845281548152600180830180548651600260001994831615610100029490940190911692909204601f81018690048602830186019096528582529194929385810193919291908301828280156110fb5780601f106110d0576101008083540402835291602001916110fb565b820191906000526020600020905b8154815290600101906020018083116110de57829003601f168201915b50505050508152602001600282015481526020016003820154815250509050600081600001511180156111bf57506006548151604080517fee684830000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a039092169163ee684830916024808201926020929091908290030181600087803b15801561119257600080fd5b505af11580156111a6573d6000803e3d6000fd5b505050506040513d60208110156111bc57600080fd5b50515b9392505050565b6000818152600560209081526040808320600654815483517f49403183000000000000000000000000000000000000000000000000000000008152600481019190915292519194600160a060020a03909116936349403183936024808201949293918390030190829087803b15801561123e57600080fd5b505af1158015611252573d6000803e3d6000fd5b505050506040513d602081101561126857600080fd5b5051156113aa57428160020154111561136757611325816001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113165780601f106112eb57610100808354040283529160200191611316565b820191906000526020600020905b8154815290600101906020018083116112f957829003601f168201915b5050505050826003015461142e565b805460408051848152602081019290925280517fe040346a7ca6935dfd5ccdb81e13933d6af35a399b27c7c61d2888b4960336479281900390910190a16113a5565b805460408051848152602081019290925280517f0571dcf79f562f7040389aac4b84570b60ed77c3a1f6d9f10f2e3dc86d647e8f9281900390910190a15b6113e8565b805460408051848152602081019290925280517fcbc6fb3892c732a14043baca80213f571ebc1a385c676b25d9907fe8e7a2e37b9281900390910190a15b60008281526005602052604081208181559061140760018301826115b6565b506000600282018190556003909101555050565b8181018281101561142857fe5b92915050565b8060046000846040518082805190602001908083835b602083106114635780518252601f199092019160209182019101611444565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916600019168152602001908152602001600020819055507f0e92bd4b74871caaf73a4a51ca5ad4f01e5c5215e940a2f2a1f4c755b955066c8282604051808060200183815260200182810382528481815181526020019150805190602001908083836000838110156108ec5781810151838201526020016108d4565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061155057805160ff191683800117855561157d565b8280016001018555821561157d579182015b8281111561157d578251825591602001919060010190611562565b506115899291506115f6565b5090565b608060405190810160405280600081526020016060815260200160008152602001600081525090565b50805460018160011615610100020316600290046000825580601f106115dc57506104aa565b601f0160209004906000526020600020908101906104aa91905b6106f291905b8082111561158957600081556001016115fc5600a165627a7a7230582051df1615dd2c537cbf5db0585d8bf81750a103d029b61d03466054fc19bce43e0029`

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

// SetNewConstitution is a paid mutator transaction binding the contract method 0xb0924d6e.
//
// Solidity: function setNewConstitution(_newConstHash bytes32, _newConstURI string) returns()
func (_GovernmentContract *GovernmentContractTransactor) SetNewConstitution(opts *bind.TransactOpts, _newConstHash [32]byte, _newConstURI string) (*types.Transaction, error) {
	return _GovernmentContract.contract.Transact(opts, "setNewConstitution", _newConstHash, _newConstURI)
}

// SetNewConstitution is a paid mutator transaction binding the contract method 0xb0924d6e.
//
// Solidity: function setNewConstitution(_newConstHash bytes32, _newConstURI string) returns()
func (_GovernmentContract *GovernmentContractSession) SetNewConstitution(_newConstHash [32]byte, _newConstURI string) (*types.Transaction, error) {
	return _GovernmentContract.Contract.SetNewConstitution(&_GovernmentContract.TransactOpts, _newConstHash, _newConstURI)
}

// SetNewConstitution is a paid mutator transaction binding the contract method 0xb0924d6e.
//
// Solidity: function setNewConstitution(_newConstHash bytes32, _newConstURI string) returns()
func (_GovernmentContract *GovernmentContractTransactorSession) SetNewConstitution(_newConstHash [32]byte, _newConstURI string) (*types.Transaction, error) {
	return _GovernmentContract.Contract.SetNewConstitution(&_GovernmentContract.TransactOpts, _newConstHash, _newConstURI)
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

// GovernmentContractNewConstSetIterator is returned from FilterNewConstSet and is used to iterate over the raw logs and unpacked data for NewConstSet events raised by the GovernmentContract contract.
type GovernmentContractNewConstSetIterator struct {
	Event *GovernmentContractNewConstSet // Event containing the contract specifics and raw log

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
func (it *GovernmentContractNewConstSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernmentContractNewConstSet)
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
		it.Event = new(GovernmentContractNewConstSet)
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
func (it *GovernmentContractNewConstSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernmentContractNewConstSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernmentContractNewConstSet represents a NewConstSet event raised by the GovernmentContract contract.
type GovernmentContractNewConstSet struct {
	ProposedHash [32]byte
	ProposedURI  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewConstSet is a free log retrieval operation binding the contract event 0x2f6679e95449d4806445cd50a14e77e4b83ea193ae84e30f8a3247436442c255.
//
// Solidity: e _NewConstSet(proposedHash bytes32, proposedURI string)
func (_GovernmentContract *GovernmentContractFilterer) FilterNewConstSet(opts *bind.FilterOpts) (*GovernmentContractNewConstSetIterator, error) {

	logs, sub, err := _GovernmentContract.contract.FilterLogs(opts, "_NewConstSet")
	if err != nil {
		return nil, err
	}
	return &GovernmentContractNewConstSetIterator{contract: _GovernmentContract.contract, event: "_NewConstSet", logs: logs, sub: sub}, nil
}

// WatchNewConstSet is a free log subscription operation binding the contract event 0x2f6679e95449d4806445cd50a14e77e4b83ea193ae84e30f8a3247436442c255.
//
// Solidity: e _NewConstSet(proposedHash bytes32, proposedURI string)
func (_GovernmentContract *GovernmentContractFilterer) WatchNewConstSet(opts *bind.WatchOpts, sink chan<- *GovernmentContractNewConstSet) (event.Subscription, error) {

	logs, sub, err := _GovernmentContract.contract.WatchLogs(opts, "_NewConstSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernmentContractNewConstSet)
				if err := _GovernmentContract.contract.UnpackLog(event, "_NewConstSet", log); err != nil {
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
