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

// NewsroomContractABI is the input ABI used to generate the binding from.
const NewsroomContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"newsroomName\",\"type\":\"string\"},{\"name\":\"charterUri\",\"type\":\"string\"},{\"name\":\"charterHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"editor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"contentId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ContentPublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contentId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"revisionId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"author\",\"type\":\"address\"}],\"name\":\"RevisionSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"editor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"contentId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"revisionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"RevisionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"granter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"grantee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"granter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"grantee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"contentId\",\"type\":\"uint256\"}],\"name\":\"getContent\",\"outputs\":[{\"name\":\"contentHash\",\"type\":\"bytes32\"},{\"name\":\"uri\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"author\",\"type\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contentId\",\"type\":\"uint256\"},{\"name\":\"revisionId\",\"type\":\"uint256\"}],\"name\":\"getRevision\",\"outputs\":[{\"name\":\"contentHash\",\"type\":\"bytes32\"},{\"name\":\"uri\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"author\",\"type\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contentId\",\"type\":\"uint256\"}],\"name\":\"revisionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contentId\",\"type\":\"uint256\"}],\"name\":\"isContentSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contentId\",\"type\":\"uint256\"},{\"name\":\"revisionId\",\"type\":\"uint256\"}],\"name\":\"isRevisionSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"string\"}],\"name\":\"addRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"addEditor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"string\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"contentUri\",\"type\":\"string\"},{\"name\":\"contentHash\",\"type\":\"bytes32\"},{\"name\":\"author\",\"type\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"publishContent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"contentId\",\"type\":\"uint256\"},{\"name\":\"contentUri\",\"type\":\"string\"},{\"name\":\"contentHash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateRevision\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"contentId\",\"type\":\"uint256\"},{\"name\":\"revisionId\",\"type\":\"uint256\"},{\"name\":\"author\",\"type\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"signRevision\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NewsroomContractBin is the compiled bytecode used for deploying new contracts.
const NewsroomContractBin = `0x60806040523480156200001157600080fd5b506040516200232738038062002327833981016040908152815160208301519183015160008054600160a060020a0319163317905590830192919091019062000063836401000000006200009b810204565b620000918282600060206040519081016040528060008152506200018c640100000000026401000000009004565b50505050620009f2565b600054600160a060020a03163314620000b357600080fd5b8051600010620000c257600080fd5b8051620000d79060059060208401906200094d565b5060408051602080825260058054600260001961010060018416150201909116049183018290527f4737457377f528cc8afd815f73ecb8b05df80d047dbffc41c17750a4033592bc939092918291820190849080156200017b5780601f106200014f576101008083540402835291602001916200017b565b820191906000526020600020905b8154815290600101906020018083116200015d57829003601f168201915b50509250505060405180910390a150565b6000806040805190810160405280600681526020017f656469746f720000000000000000000000000000000000000000000000000000815250620001df336200034a640100000000026401000000009004565b80620001fb5750620001fb33826401000000006200035e810204565b15156200020757600080fd5b60048054600181019091559150600160a060020a0385161580156200022b57508351155b806200024b5750600160a060020a038516158015906200024b5750835115155b15156200025757600080fd5b60008281526002602052604090206001018054600160a060020a031916600160a060020a0387161790556200029882888887640100000000620003e4810204565b508133600160a060020a03167f1ede735f9b446d8014022fed176848ac3894c54942bef9ff452f7ae42b50d5ae896040518080602001828103825283818151815260200191508051906020019080838360005b8381101562000305578181015183820152602001620002eb565b50505050905090810190601f168015620003335780820380516001836020036101000a031916815260200191505b509250505060405180910390a35095945050505050565b600054600160a060020a0390811691161490565b60006001826040518082805190602001908083835b60208310620003945780518252601f19909201916020918201910162000373565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201909420600160a060020a0397909716600090815296905250509092205460ff1692915050565b600080600060045487101515620003fa57600080fd5b861515620004225762000416336401000000006200034a810204565b15156200042257600080fd5b50506000858152600260209081526040808320805482516080810184528881528085018a815242948201949094526060810188905260018083018085558488529686902082516004850290910190815594518051949793969395929492936200049293928501929101906200094d565b506040820151600282015560608201518051620004ba9160038401916020909101906200094d565b50505050600182015482546200050691600160a060020a0316908990859085908110620004e357fe5b9060005260206000209060040201620005fd640100000000026401000000009004565b156200054b576001820154604051600160a060020a0390911690829089907f605611fc50d3effbe4af88e82f5daebfcffe0fb8f3b34ed32f1a746290ccbc6190600090a45b808733600160a060020a03167f18b6b5c485f8822a270464dd544d0715597dc8f1a007ee2b0252b62b8b9fb390896040518080602001828103825283818151815260200191508051906020019080838360005b83811015620005b85781810151838201526020016200059e565b50505050905090810190601f168015620005e65780820380516001836020036101000a031916815260200191505b509250505060405180910390a45050949350505050565b60008080600160a060020a03861615806200062d5750600384015460026000196101006001841615020190911604155b156200066057600384015460026000196101006001841615020190911604156200065657600080fd5b60009250620007bc565b8354604080516c0100000000000000000000000030028152601481019290925251908190036034019020620006a39064010000000062001671620007c582021704565b600385018054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152939550600160a060020a038a1693620007629390929091830182828015620007455780601f10620007195761010080835404028352916020019162000745565b820191906000526020600020905b8154815290600101906020018083116200072757829003601f168201915b5088949350506401000000006200171b6200087182021704915050565b600160a060020a0316146200077657600080fd5b506000818152600360205260409020805460ff1615806200079a5750848160010154145b1515620007a657600080fd5b805460ff19166001908117825581810186905592505b50509392505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c80830185905283518084039091018152605c909201928390528151600093918291908401908083835b602083106200083f5780518252601f1990920191602091820191016200081e565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050565b600080600080845160411415156200088d576000935062000944565b50505060208201516040830151606084015160001a601b60ff82161015620008b357601b015b8060ff16601b14158015620008cc57508060ff16601c14155b15620008dc576000935062000944565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af115801562000937573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200099057805160ff1916838001178555620009c0565b82800160010185558215620009c0579182015b82811115620009c0578251825591602001919060010190620009a3565b50620009ce929150620009d2565b5090565b620009ef91905b80821115620009ce5760008155600101620009d9565b90565b6119258062000a026000396000f3006080604052600436106101065763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461010b5780630b7ad54c146101955780631bfe0308146102aa578063217fe6c6146102d95780632f54bf6e146103545780635614bdc8146103755780636192e3e81461039f57806365462d96146103ba578063715018a6146103ef5780637d72aa651461040457806384a1176c146104315780638da5cb5b146104de578063a54d19881461050f578063c47f00271461052a578063cc45969614610583578063e45e1c7d14610598578063e5975bdc146105cc578063efc97390146105ed578063f2fde38b14610605575b600080fd5b34801561011757600080fd5b50610120610626565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561015a578181015183820152602001610142565b50505050905090810190601f1680156101875780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101a157600080fd5b506101ad6004356106b4565b60408051868152908101849052600160a060020a038316606082015260a060208083018281528751928401929092528651608084019160c08501919089019080838360005b8381101561020a5781810151838201526020016101f2565b50505050905090810190601f1680156102375780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561026a578181015183820152602001610252565b50505050905090810190601f1680156102975780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3480156102b657600080fd5b506102d760048035600160a060020a031690602480359081019101356106ed565b005b3480156102e557600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610340958335600160a060020a03169536956044949193909101919081908401838280828437509497506107719650505050505050565b604080519115158252519081900360200190f35b34801561036057600080fd5b50610340600160a060020a03600435166107f5565b34801561038157600080fd5b5061038d600435610809565b60408051918252519081900360200190f35b3480156103ab57600080fd5b506101ad60043560243561081b565b3480156103c657600080fd5b506102d760048035906024803591600160a060020a0360443516916064359081019101356109b6565b3480156103fb57600080fd5b506102d7610b2b565b34801561041057600080fd5b506102d760048035600160a060020a03169060248035908101910135610b97565b34801561043d57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261038d94369492936024939284019190819084018382808284375050604080516020888301358a018035601f8101839004830284018301909452838352979a89359a8a830135600160a060020a03169a91999098506060909101965091945090810192508190840183828082843750949750610c159650505050505050565b3480156104ea57600080fd5b506104f3610d9c565b60408051600160a060020a039092168252519081900360200190f35b34801561051b57600080fd5b50610340600435602435610dab565b34801561053657600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526102d7943694929360249392840191908190840183828082843750949750610e059650505050505050565b34801561058f57600080fd5b5061038d610eee565b3480156105a457600080fd5b506102d760048035906024803580820192908101359160443591606435918201910135610ef4565b3480156105d857600080fd5b506102d7600160a060020a0360043516610fac565b3480156105f957600080fd5b5061034060043561101f565b34801561061157600080fd5b506102d7600160a060020a0360043516611043565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106ac5780601f10610681576101008083540402835291602001916106ac565b820191906000526020600020905b81548152906001019060200180831161068f57829003601f168201915b505050505081565b6000818152600260205260408120546060908290819083906106db9087906000190161081b565b939a9299509097509550909350915050565b604080518082019091526006815260d160020a6532b234ba37b9026020820152610716336107f5565b8061072657506107263382610771565b151561073157600080fd5b61076b8484848080601f01602080910402602001604051908101604052809392919081815260200183838082843750611066945050505050565b50505050565b60006001826040518082805190602001908083835b602083106107a55780518252601f199092019160209182019101610786565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201909420600160a060020a0397909716600090815296905250509092205460ff1692915050565b600054600160a060020a0390811691161490565b60009081526002602052604090205490565b600082815260026020526040812080546060918391829184918390881061084157600080fd5b815482908990811061084f57fe5b600091825260209182902060049190910201805460028083015460018781015481860180546040805161010095831615959095026000190190911695909504601f810189900489028401890190955284835295975093959193600160a060020a031692600388019286919083018282801561090b5780601f106108e05761010080835404028352916020019161090b565b820191906000526020600020905b8154815290600101906020018083116108ee57829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959950869450925084019050828280156109995780601f1061096e57610100808354040283529160200191610999565b820191906000526020600020905b81548152906001019060200180831161097c57829003601f168201915b505050505090509650965096509650965050509295509295909350565b60008060408051908101604052806006815260200160d160020a6532b234ba37b9028152506109e4336107f5565b806109f457506109f43382610771565b15156109ff57600080fd5b6004548810610a0d57600080fd5b60008881526002602052604090206001810154909350600160a060020a03161580610a4757506001830154600160a060020a038781169116145b1515610a5257600080fd5b82548710610a5f57600080fd5b871515610a7a57610a6f336107f5565b1515610a7a57600080fd5b60018301805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388161790558254839088908110610ab257fe5b600091825260209091206004909102019150610ad26003830186866117f0565b50610ade868984611180565b1515610ae957600080fd5b85600160a060020a031687897f605611fc50d3effbe4af88e82f5daebfcffe0fb8f3b34ed32f1a746290ccbc6160405160405180910390a45050505050505050565b600054600160a060020a03163314610b4257600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b604080518082019091526006815260d160020a6532b234ba37b9026020820152610bc0336107f5565b80610bd05750610bd03382610771565b1515610bdb57600080fd5b61076b8484848080601f01602080910402602001604051908101604052809392919081815260200183838082843750611323945050505050565b60008060408051908101604052806006815260200160d160020a6532b234ba37b902815250610c43336107f5565b80610c535750610c533382610771565b1515610c5e57600080fd5b60048054600181019091559150600160a060020a038516158015610c8157508351155b80610c9f5750600160a060020a03851615801590610c9f5750835115155b1515610caa57600080fd5b6000828152600260205260409020600101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038716179055610ced82888887611403565b508133600160a060020a03167f1ede735f9b446d8014022fed176848ac3894c54942bef9ff452f7ae42b50d5ae896040518080602001828103825283818151815260200191508051906020019080838360005b83811015610d58578181015183820152602001610d40565b50505050905090810190601f168015610d855780820380516001836020036101000a031916815260200191505b509250505060405180910390a35095945050505050565b600054600160a060020a031681565b600082815260026020526040812080548310610dc657600080fd5b8083815481101515610dd457fe5b6000918252602090912060036004909202010154600260001961010060018416150201909116041515949350505050565b600054600160a060020a03163314610e1c57600080fd5b8051600010610e2a57600080fd5b8051610e3d90600590602084019061186e565b5060408051602080825260058054600260001961010060018416150201909116049183018290527f4737457377f528cc8afd815f73ecb8b05df80d047dbffc41c17750a4033592bc93909291829182019084908015610edd5780601f10610eb257610100808354040283529160200191610edd565b820191906000526020600020905b815481529060010190602001808311610ec057829003601f168201915b50509250505060405180910390a150565b60045481565b604080518082019091526006815260d160020a6532b234ba37b9026020820152610f1d336107f5565b80610f2d5750610f2d3382610771565b1515610f3857600080fd5b610fa28787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375050604080516020601f8c018190048102820181019092528a81528c955093508a92508991508190840183828082843750611403945050505050565b5050505050505050565b604080518082019091526006815260d160020a6532b234ba37b9026020820152610fd5336107f5565b80610fe55750610fe53382610771565b1515610ff057600080fd5b61101b8260408051908101604052806006815260200160d160020a6532b234ba37b902815250611323565b5050565b60008181526002602052604081205461103d90839060001901610dab565b92915050565b600054600160a060020a0316331461105a57600080fd5b611063816115f4565b50565b6001816040518082805190602001908083835b602083106110985780518252601f199092019160209182019101611079565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520600160a060020a0388166000818152918452828220805460ff19169055838752875187850152875190963396507f6a52fb0cb0e75e6a6721483d2e539b38273ec6fe95b648a41e1a901594aeccb895508894909384939084019291860191908190849084905b8381101561114257818101518382015260200161112a565b50505050905090810190601f16801561116f5780820380516001836020036101000a031916815260200191505b509250505060405180910390a35050565b60008080600160a060020a03861615806111af5750600384015460026000196101006001841615020190911604155b156111df57600384015460026000196101006001841615020190911604156111d657600080fd5b6000925061131a565b8354604080516c010000000000000000000000003002815260148101929092525190819003603401902061121290611671565b600385018054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152939550600160a060020a038a16936112c393909290918301828280156112af5780601f10611284576101008083540402835291602001916112af565b820191906000526020600020905b81548152906001019060200180831161129257829003601f168201915b50889493505063ffffffff61171b16915050565b600160a060020a0316146112d657600080fd5b506000818152600360205260409020805460ff1615806112f95750848160010154145b151561130457600080fd5b805460ff19166001908117825581810186905592505b50509392505050565b600180826040518082805190602001908083835b602083106113565780518252601f199092019160209182019101611337565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520600160a060020a0389166000818152918452828220805460ff191698151598909817909755828652875186840152875133967fa40c1dc2b34f6b51c3ea614b688f243e50047ed9fa3ea19010303d70dac781ed965089955093849384019290860191908190849084908381101561114257818101518382015260200161112a565b60008060006004548710151561141857600080fd5b86151561143357611428336107f5565b151561143357600080fd5b50506000858152600260209081526040808320805482516080810184528881528085018a815242948201949094526060810188905260018083018085558488529686902082516004850290910190815594518051949793969395929492936114a1939285019291019061186e565b5060408201516002820155606082015180516114c791600384019160209091019061186e565b505050506001820154825461150191600160a060020a03169089908590859081106114ee57fe5b9060005260206000209060040201611180565b15611545576001820154604051600160a060020a0390911690829089907f605611fc50d3effbe4af88e82f5daebfcffe0fb8f3b34ed32f1a746290ccbc6190600090a45b808733600160a060020a03167f18b6b5c485f8822a270464dd544d0715597dc8f1a007ee2b0252b62b8b9fb390896040518080602001828103825283818151815260200191508051906020019080838360005b838110156115b0578181015183820152602001611598565b50505050905090810190601f1680156115dd5780820380516001836020036101000a031916815260200191505b509250505060405180910390a45050949350505050565b600160a060020a038116151561160957600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c80830185905283518084039091018152605c909201928390528151600093918291908401908083835b602083106116e95780518252601f1990920191602091820191016116ca565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050565b6000806000808451604114151561173557600093506117e7565b50505060208201516040830151606084015160001a601b60ff8216101561175a57601b015b8060ff16601b1415801561177257508060ff16601c14155b1561178057600093506117e7565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af11580156117da573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106118315782800160ff1982351617855561185e565b8280016001018555821561185e579182015b8281111561185e578235825591602001919060010190611843565b5061186a9291506118dc565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106118af57805160ff191683800117855561185e565b8280016001018555821561185e579182015b8281111561185e5782518255916020019190600101906118c1565b6118f691905b8082111561186a57600081556001016118e2565b905600a165627a7a72305820bcc9c0cfe282f811125487adf087f3850f2c4e12410c37a565e3c3c12aee91140029`

// DeployNewsroomContract deploys a new Ethereum contract, binding an instance of NewsroomContract to it.
func DeployNewsroomContract(auth *bind.TransactOpts, backend bind.ContractBackend, newsroomName string, charterUri string, charterHash [32]byte) (common.Address, *types.Transaction, *NewsroomContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NewsroomContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NewsroomContractBin), backend, newsroomName, charterUri, charterHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NewsroomContract{NewsroomContractCaller: NewsroomContractCaller{contract: contract}, NewsroomContractTransactor: NewsroomContractTransactor{contract: contract}, NewsroomContractFilterer: NewsroomContractFilterer{contract: contract}}, nil
}

// NewsroomContract is an auto generated Go binding around an Ethereum contract.
type NewsroomContract struct {
	NewsroomContractCaller     // Read-only binding to the contract
	NewsroomContractTransactor // Write-only binding to the contract
	NewsroomContractFilterer   // Log filterer for contract events
}

// NewsroomContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type NewsroomContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewsroomContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NewsroomContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewsroomContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NewsroomContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewsroomContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NewsroomContractSession struct {
	Contract     *NewsroomContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NewsroomContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NewsroomContractCallerSession struct {
	Contract *NewsroomContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// NewsroomContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NewsroomContractTransactorSession struct {
	Contract     *NewsroomContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// NewsroomContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type NewsroomContractRaw struct {
	Contract *NewsroomContract // Generic contract binding to access the raw methods on
}

// NewsroomContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NewsroomContractCallerRaw struct {
	Contract *NewsroomContractCaller // Generic read-only contract binding to access the raw methods on
}

// NewsroomContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NewsroomContractTransactorRaw struct {
	Contract *NewsroomContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNewsroomContract creates a new instance of NewsroomContract, bound to a specific deployed contract.
func NewNewsroomContract(address common.Address, backend bind.ContractBackend) (*NewsroomContract, error) {
	contract, err := bindNewsroomContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NewsroomContract{NewsroomContractCaller: NewsroomContractCaller{contract: contract}, NewsroomContractTransactor: NewsroomContractTransactor{contract: contract}, NewsroomContractFilterer: NewsroomContractFilterer{contract: contract}}, nil
}

// NewNewsroomContractCaller creates a new read-only instance of NewsroomContract, bound to a specific deployed contract.
func NewNewsroomContractCaller(address common.Address, caller bind.ContractCaller) (*NewsroomContractCaller, error) {
	contract, err := bindNewsroomContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NewsroomContractCaller{contract: contract}, nil
}

// NewNewsroomContractTransactor creates a new write-only instance of NewsroomContract, bound to a specific deployed contract.
func NewNewsroomContractTransactor(address common.Address, transactor bind.ContractTransactor) (*NewsroomContractTransactor, error) {
	contract, err := bindNewsroomContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NewsroomContractTransactor{contract: contract}, nil
}

// NewNewsroomContractFilterer creates a new log filterer instance of NewsroomContract, bound to a specific deployed contract.
func NewNewsroomContractFilterer(address common.Address, filterer bind.ContractFilterer) (*NewsroomContractFilterer, error) {
	contract, err := bindNewsroomContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NewsroomContractFilterer{contract: contract}, nil
}

// bindNewsroomContract binds a generic wrapper to an already deployed contract.
func bindNewsroomContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NewsroomContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NewsroomContract *NewsroomContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NewsroomContract.Contract.NewsroomContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NewsroomContract *NewsroomContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewsroomContract.Contract.NewsroomContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NewsroomContract *NewsroomContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NewsroomContract.Contract.NewsroomContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NewsroomContract *NewsroomContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NewsroomContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NewsroomContract *NewsroomContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewsroomContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NewsroomContract *NewsroomContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NewsroomContract.Contract.contract.Transact(opts, method, params...)
}

// ContentCount is a free data retrieval call binding the contract method 0xcc459696.
//
// Solidity: function contentCount() constant returns(uint256)
func (_NewsroomContract *NewsroomContractCaller) ContentCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NewsroomContract.contract.Call(opts, out, "contentCount")
	return *ret0, err
}

// ContentCount is a free data retrieval call binding the contract method 0xcc459696.
//
// Solidity: function contentCount() constant returns(uint256)
func (_NewsroomContract *NewsroomContractSession) ContentCount() (*big.Int, error) {
	return _NewsroomContract.Contract.ContentCount(&_NewsroomContract.CallOpts)
}

// ContentCount is a free data retrieval call binding the contract method 0xcc459696.
//
// Solidity: function contentCount() constant returns(uint256)
func (_NewsroomContract *NewsroomContractCallerSession) ContentCount() (*big.Int, error) {
	return _NewsroomContract.Contract.ContentCount(&_NewsroomContract.CallOpts)
}

// GetContent is a free data retrieval call binding the contract method 0x0b7ad54c.
//
// Solidity: function getContent(contentId uint256) constant returns(contentHash bytes32, uri string, timestamp uint256, author address, signature bytes)
func (_NewsroomContract *NewsroomContractCaller) GetContent(opts *bind.CallOpts, contentId *big.Int) (struct {
	ContentHash [32]byte
	Uri         string
	Timestamp   *big.Int
	Author      common.Address
	Signature   []byte
}, error) {
	ret := new(struct {
		ContentHash [32]byte
		Uri         string
		Timestamp   *big.Int
		Author      common.Address
		Signature   []byte
	})
	out := ret
	err := _NewsroomContract.contract.Call(opts, out, "getContent", contentId)
	return *ret, err
}

// GetContent is a free data retrieval call binding the contract method 0x0b7ad54c.
//
// Solidity: function getContent(contentId uint256) constant returns(contentHash bytes32, uri string, timestamp uint256, author address, signature bytes)
func (_NewsroomContract *NewsroomContractSession) GetContent(contentId *big.Int) (struct {
	ContentHash [32]byte
	Uri         string
	Timestamp   *big.Int
	Author      common.Address
	Signature   []byte
}, error) {
	return _NewsroomContract.Contract.GetContent(&_NewsroomContract.CallOpts, contentId)
}

// GetContent is a free data retrieval call binding the contract method 0x0b7ad54c.
//
// Solidity: function getContent(contentId uint256) constant returns(contentHash bytes32, uri string, timestamp uint256, author address, signature bytes)
func (_NewsroomContract *NewsroomContractCallerSession) GetContent(contentId *big.Int) (struct {
	ContentHash [32]byte
	Uri         string
	Timestamp   *big.Int
	Author      common.Address
	Signature   []byte
}, error) {
	return _NewsroomContract.Contract.GetContent(&_NewsroomContract.CallOpts, contentId)
}

// GetRevision is a free data retrieval call binding the contract method 0x6192e3e8.
//
// Solidity: function getRevision(contentId uint256, revisionId uint256) constant returns(contentHash bytes32, uri string, timestamp uint256, author address, signature bytes)
func (_NewsroomContract *NewsroomContractCaller) GetRevision(opts *bind.CallOpts, contentId *big.Int, revisionId *big.Int) (struct {
	ContentHash [32]byte
	Uri         string
	Timestamp   *big.Int
	Author      common.Address
	Signature   []byte
}, error) {
	ret := new(struct {
		ContentHash [32]byte
		Uri         string
		Timestamp   *big.Int
		Author      common.Address
		Signature   []byte
	})
	out := ret
	err := _NewsroomContract.contract.Call(opts, out, "getRevision", contentId, revisionId)
	return *ret, err
}

// GetRevision is a free data retrieval call binding the contract method 0x6192e3e8.
//
// Solidity: function getRevision(contentId uint256, revisionId uint256) constant returns(contentHash bytes32, uri string, timestamp uint256, author address, signature bytes)
func (_NewsroomContract *NewsroomContractSession) GetRevision(contentId *big.Int, revisionId *big.Int) (struct {
	ContentHash [32]byte
	Uri         string
	Timestamp   *big.Int
	Author      common.Address
	Signature   []byte
}, error) {
	return _NewsroomContract.Contract.GetRevision(&_NewsroomContract.CallOpts, contentId, revisionId)
}

// GetRevision is a free data retrieval call binding the contract method 0x6192e3e8.
//
// Solidity: function getRevision(contentId uint256, revisionId uint256) constant returns(contentHash bytes32, uri string, timestamp uint256, author address, signature bytes)
func (_NewsroomContract *NewsroomContractCallerSession) GetRevision(contentId *big.Int, revisionId *big.Int) (struct {
	ContentHash [32]byte
	Uri         string
	Timestamp   *big.Int
	Author      common.Address
	Signature   []byte
}, error) {
	return _NewsroomContract.Contract.GetRevision(&_NewsroomContract.CallOpts, contentId, revisionId)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(user address, role string) constant returns(bool)
func (_NewsroomContract *NewsroomContractCaller) HasRole(opts *bind.CallOpts, user common.Address, role string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NewsroomContract.contract.Call(opts, out, "hasRole", user, role)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(user address, role string) constant returns(bool)
func (_NewsroomContract *NewsroomContractSession) HasRole(user common.Address, role string) (bool, error) {
	return _NewsroomContract.Contract.HasRole(&_NewsroomContract.CallOpts, user, role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(user address, role string) constant returns(bool)
func (_NewsroomContract *NewsroomContractCallerSession) HasRole(user common.Address, role string) (bool, error) {
	return _NewsroomContract.Contract.HasRole(&_NewsroomContract.CallOpts, user, role)
}

// IsContentSigned is a free data retrieval call binding the contract method 0xefc97390.
//
// Solidity: function isContentSigned(contentId uint256) constant returns(bool)
func (_NewsroomContract *NewsroomContractCaller) IsContentSigned(opts *bind.CallOpts, contentId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NewsroomContract.contract.Call(opts, out, "isContentSigned", contentId)
	return *ret0, err
}

// IsContentSigned is a free data retrieval call binding the contract method 0xefc97390.
//
// Solidity: function isContentSigned(contentId uint256) constant returns(bool)
func (_NewsroomContract *NewsroomContractSession) IsContentSigned(contentId *big.Int) (bool, error) {
	return _NewsroomContract.Contract.IsContentSigned(&_NewsroomContract.CallOpts, contentId)
}

// IsContentSigned is a free data retrieval call binding the contract method 0xefc97390.
//
// Solidity: function isContentSigned(contentId uint256) constant returns(bool)
func (_NewsroomContract *NewsroomContractCallerSession) IsContentSigned(contentId *big.Int) (bool, error) {
	return _NewsroomContract.Contract.IsContentSigned(&_NewsroomContract.CallOpts, contentId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(user address) constant returns(bool)
func (_NewsroomContract *NewsroomContractCaller) IsOwner(opts *bind.CallOpts, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NewsroomContract.contract.Call(opts, out, "isOwner", user)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(user address) constant returns(bool)
func (_NewsroomContract *NewsroomContractSession) IsOwner(user common.Address) (bool, error) {
	return _NewsroomContract.Contract.IsOwner(&_NewsroomContract.CallOpts, user)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(user address) constant returns(bool)
func (_NewsroomContract *NewsroomContractCallerSession) IsOwner(user common.Address) (bool, error) {
	return _NewsroomContract.Contract.IsOwner(&_NewsroomContract.CallOpts, user)
}

// IsRevisionSigned is a free data retrieval call binding the contract method 0xa54d1988.
//
// Solidity: function isRevisionSigned(contentId uint256, revisionId uint256) constant returns(bool)
func (_NewsroomContract *NewsroomContractCaller) IsRevisionSigned(opts *bind.CallOpts, contentId *big.Int, revisionId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NewsroomContract.contract.Call(opts, out, "isRevisionSigned", contentId, revisionId)
	return *ret0, err
}

// IsRevisionSigned is a free data retrieval call binding the contract method 0xa54d1988.
//
// Solidity: function isRevisionSigned(contentId uint256, revisionId uint256) constant returns(bool)
func (_NewsroomContract *NewsroomContractSession) IsRevisionSigned(contentId *big.Int, revisionId *big.Int) (bool, error) {
	return _NewsroomContract.Contract.IsRevisionSigned(&_NewsroomContract.CallOpts, contentId, revisionId)
}

// IsRevisionSigned is a free data retrieval call binding the contract method 0xa54d1988.
//
// Solidity: function isRevisionSigned(contentId uint256, revisionId uint256) constant returns(bool)
func (_NewsroomContract *NewsroomContractCallerSession) IsRevisionSigned(contentId *big.Int, revisionId *big.Int) (bool, error) {
	return _NewsroomContract.Contract.IsRevisionSigned(&_NewsroomContract.CallOpts, contentId, revisionId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NewsroomContract *NewsroomContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NewsroomContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NewsroomContract *NewsroomContractSession) Name() (string, error) {
	return _NewsroomContract.Contract.Name(&_NewsroomContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NewsroomContract *NewsroomContractCallerSession) Name() (string, error) {
	return _NewsroomContract.Contract.Name(&_NewsroomContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NewsroomContract *NewsroomContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NewsroomContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NewsroomContract *NewsroomContractSession) Owner() (common.Address, error) {
	return _NewsroomContract.Contract.Owner(&_NewsroomContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NewsroomContract *NewsroomContractCallerSession) Owner() (common.Address, error) {
	return _NewsroomContract.Contract.Owner(&_NewsroomContract.CallOpts)
}

// RevisionCount is a free data retrieval call binding the contract method 0x5614bdc8.
//
// Solidity: function revisionCount(contentId uint256) constant returns(uint256)
func (_NewsroomContract *NewsroomContractCaller) RevisionCount(opts *bind.CallOpts, contentId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NewsroomContract.contract.Call(opts, out, "revisionCount", contentId)
	return *ret0, err
}

// RevisionCount is a free data retrieval call binding the contract method 0x5614bdc8.
//
// Solidity: function revisionCount(contentId uint256) constant returns(uint256)
func (_NewsroomContract *NewsroomContractSession) RevisionCount(contentId *big.Int) (*big.Int, error) {
	return _NewsroomContract.Contract.RevisionCount(&_NewsroomContract.CallOpts, contentId)
}

// RevisionCount is a free data retrieval call binding the contract method 0x5614bdc8.
//
// Solidity: function revisionCount(contentId uint256) constant returns(uint256)
func (_NewsroomContract *NewsroomContractCallerSession) RevisionCount(contentId *big.Int) (*big.Int, error) {
	return _NewsroomContract.Contract.RevisionCount(&_NewsroomContract.CallOpts, contentId)
}

// AddEditor is a paid mutator transaction binding the contract method 0xe5975bdc.
//
// Solidity: function addEditor(who address) returns()
func (_NewsroomContract *NewsroomContractTransactor) AddEditor(opts *bind.TransactOpts, who common.Address) (*types.Transaction, error) {
	return _NewsroomContract.contract.Transact(opts, "addEditor", who)
}

// AddEditor is a paid mutator transaction binding the contract method 0xe5975bdc.
//
// Solidity: function addEditor(who address) returns()
func (_NewsroomContract *NewsroomContractSession) AddEditor(who common.Address) (*types.Transaction, error) {
	return _NewsroomContract.Contract.AddEditor(&_NewsroomContract.TransactOpts, who)
}

// AddEditor is a paid mutator transaction binding the contract method 0xe5975bdc.
//
// Solidity: function addEditor(who address) returns()
func (_NewsroomContract *NewsroomContractTransactorSession) AddEditor(who common.Address) (*types.Transaction, error) {
	return _NewsroomContract.Contract.AddEditor(&_NewsroomContract.TransactOpts, who)
}

// AddRole is a paid mutator transaction binding the contract method 0x7d72aa65.
//
// Solidity: function addRole(who address, role string) returns()
func (_NewsroomContract *NewsroomContractTransactor) AddRole(opts *bind.TransactOpts, who common.Address, role string) (*types.Transaction, error) {
	return _NewsroomContract.contract.Transact(opts, "addRole", who, role)
}

// AddRole is a paid mutator transaction binding the contract method 0x7d72aa65.
//
// Solidity: function addRole(who address, role string) returns()
func (_NewsroomContract *NewsroomContractSession) AddRole(who common.Address, role string) (*types.Transaction, error) {
	return _NewsroomContract.Contract.AddRole(&_NewsroomContract.TransactOpts, who, role)
}

// AddRole is a paid mutator transaction binding the contract method 0x7d72aa65.
//
// Solidity: function addRole(who address, role string) returns()
func (_NewsroomContract *NewsroomContractTransactorSession) AddRole(who common.Address, role string) (*types.Transaction, error) {
	return _NewsroomContract.Contract.AddRole(&_NewsroomContract.TransactOpts, who, role)
}

// PublishContent is a paid mutator transaction binding the contract method 0x84a1176c.
//
// Solidity: function publishContent(contentUri string, contentHash bytes32, author address, signature bytes) returns(uint256)
func (_NewsroomContract *NewsroomContractTransactor) PublishContent(opts *bind.TransactOpts, contentUri string, contentHash [32]byte, author common.Address, signature []byte) (*types.Transaction, error) {
	return _NewsroomContract.contract.Transact(opts, "publishContent", contentUri, contentHash, author, signature)
}

// PublishContent is a paid mutator transaction binding the contract method 0x84a1176c.
//
// Solidity: function publishContent(contentUri string, contentHash bytes32, author address, signature bytes) returns(uint256)
func (_NewsroomContract *NewsroomContractSession) PublishContent(contentUri string, contentHash [32]byte, author common.Address, signature []byte) (*types.Transaction, error) {
	return _NewsroomContract.Contract.PublishContent(&_NewsroomContract.TransactOpts, contentUri, contentHash, author, signature)
}

// PublishContent is a paid mutator transaction binding the contract method 0x84a1176c.
//
// Solidity: function publishContent(contentUri string, contentHash bytes32, author address, signature bytes) returns(uint256)
func (_NewsroomContract *NewsroomContractTransactorSession) PublishContent(contentUri string, contentHash [32]byte, author common.Address, signature []byte) (*types.Transaction, error) {
	return _NewsroomContract.Contract.PublishContent(&_NewsroomContract.TransactOpts, contentUri, contentHash, author, signature)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x1bfe0308.
//
// Solidity: function removeRole(who address, role string) returns()
func (_NewsroomContract *NewsroomContractTransactor) RemoveRole(opts *bind.TransactOpts, who common.Address, role string) (*types.Transaction, error) {
	return _NewsroomContract.contract.Transact(opts, "removeRole", who, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x1bfe0308.
//
// Solidity: function removeRole(who address, role string) returns()
func (_NewsroomContract *NewsroomContractSession) RemoveRole(who common.Address, role string) (*types.Transaction, error) {
	return _NewsroomContract.Contract.RemoveRole(&_NewsroomContract.TransactOpts, who, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x1bfe0308.
//
// Solidity: function removeRole(who address, role string) returns()
func (_NewsroomContract *NewsroomContractTransactorSession) RemoveRole(who common.Address, role string) (*types.Transaction, error) {
	return _NewsroomContract.Contract.RemoveRole(&_NewsroomContract.TransactOpts, who, role)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NewsroomContract *NewsroomContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewsroomContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NewsroomContract *NewsroomContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _NewsroomContract.Contract.RenounceOwnership(&_NewsroomContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NewsroomContract *NewsroomContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NewsroomContract.Contract.RenounceOwnership(&_NewsroomContract.TransactOpts)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(newName string) returns()
func (_NewsroomContract *NewsroomContractTransactor) SetName(opts *bind.TransactOpts, newName string) (*types.Transaction, error) {
	return _NewsroomContract.contract.Transact(opts, "setName", newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(newName string) returns()
func (_NewsroomContract *NewsroomContractSession) SetName(newName string) (*types.Transaction, error) {
	return _NewsroomContract.Contract.SetName(&_NewsroomContract.TransactOpts, newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(newName string) returns()
func (_NewsroomContract *NewsroomContractTransactorSession) SetName(newName string) (*types.Transaction, error) {
	return _NewsroomContract.Contract.SetName(&_NewsroomContract.TransactOpts, newName)
}

// SignRevision is a paid mutator transaction binding the contract method 0x65462d96.
//
// Solidity: function signRevision(contentId uint256, revisionId uint256, author address, signature bytes) returns()
func (_NewsroomContract *NewsroomContractTransactor) SignRevision(opts *bind.TransactOpts, contentId *big.Int, revisionId *big.Int, author common.Address, signature []byte) (*types.Transaction, error) {
	return _NewsroomContract.contract.Transact(opts, "signRevision", contentId, revisionId, author, signature)
}

// SignRevision is a paid mutator transaction binding the contract method 0x65462d96.
//
// Solidity: function signRevision(contentId uint256, revisionId uint256, author address, signature bytes) returns()
func (_NewsroomContract *NewsroomContractSession) SignRevision(contentId *big.Int, revisionId *big.Int, author common.Address, signature []byte) (*types.Transaction, error) {
	return _NewsroomContract.Contract.SignRevision(&_NewsroomContract.TransactOpts, contentId, revisionId, author, signature)
}

// SignRevision is a paid mutator transaction binding the contract method 0x65462d96.
//
// Solidity: function signRevision(contentId uint256, revisionId uint256, author address, signature bytes) returns()
func (_NewsroomContract *NewsroomContractTransactorSession) SignRevision(contentId *big.Int, revisionId *big.Int, author common.Address, signature []byte) (*types.Transaction, error) {
	return _NewsroomContract.Contract.SignRevision(&_NewsroomContract.TransactOpts, contentId, revisionId, author, signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_NewsroomContract *NewsroomContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _NewsroomContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_NewsroomContract *NewsroomContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _NewsroomContract.Contract.TransferOwnership(&_NewsroomContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_NewsroomContract *NewsroomContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _NewsroomContract.Contract.TransferOwnership(&_NewsroomContract.TransactOpts, _newOwner)
}

// UpdateRevision is a paid mutator transaction binding the contract method 0xe45e1c7d.
//
// Solidity: function updateRevision(contentId uint256, contentUri string, contentHash bytes32, signature bytes) returns()
func (_NewsroomContract *NewsroomContractTransactor) UpdateRevision(opts *bind.TransactOpts, contentId *big.Int, contentUri string, contentHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _NewsroomContract.contract.Transact(opts, "updateRevision", contentId, contentUri, contentHash, signature)
}

// UpdateRevision is a paid mutator transaction binding the contract method 0xe45e1c7d.
//
// Solidity: function updateRevision(contentId uint256, contentUri string, contentHash bytes32, signature bytes) returns()
func (_NewsroomContract *NewsroomContractSession) UpdateRevision(contentId *big.Int, contentUri string, contentHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _NewsroomContract.Contract.UpdateRevision(&_NewsroomContract.TransactOpts, contentId, contentUri, contentHash, signature)
}

// UpdateRevision is a paid mutator transaction binding the contract method 0xe45e1c7d.
//
// Solidity: function updateRevision(contentId uint256, contentUri string, contentHash bytes32, signature bytes) returns()
func (_NewsroomContract *NewsroomContractTransactorSession) UpdateRevision(contentId *big.Int, contentUri string, contentHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _NewsroomContract.Contract.UpdateRevision(&_NewsroomContract.TransactOpts, contentId, contentUri, contentHash, signature)
}

// NewsroomContractContentPublishedIterator is returned from FilterContentPublished and is used to iterate over the raw logs and unpacked data for ContentPublished events raised by the NewsroomContract contract.
type NewsroomContractContentPublishedIterator struct {
	Event *NewsroomContractContentPublished // Event containing the contract specifics and raw log

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
func (it *NewsroomContractContentPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewsroomContractContentPublished)
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
		it.Event = new(NewsroomContractContentPublished)
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
func (it *NewsroomContractContentPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewsroomContractContentPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewsroomContractContentPublished represents a ContentPublished event raised by the NewsroomContract contract.
type NewsroomContractContentPublished struct {
	Editor    common.Address
	ContentId *big.Int
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContentPublished is a free log retrieval operation binding the contract event 0x1ede735f9b446d8014022fed176848ac3894c54942bef9ff452f7ae42b50d5ae.
//
// Solidity: e ContentPublished(editor indexed address, contentId indexed uint256, uri string)
func (_NewsroomContract *NewsroomContractFilterer) FilterContentPublished(opts *bind.FilterOpts, editor []common.Address, contentId []*big.Int) (*NewsroomContractContentPublishedIterator, error) {

	var editorRule []interface{}
	for _, editorItem := range editor {
		editorRule = append(editorRule, editorItem)
	}
	var contentIdRule []interface{}
	for _, contentIdItem := range contentId {
		contentIdRule = append(contentIdRule, contentIdItem)
	}

	logs, sub, err := _NewsroomContract.contract.FilterLogs(opts, "ContentPublished", editorRule, contentIdRule)
	if err != nil {
		return nil, err
	}
	return &NewsroomContractContentPublishedIterator{contract: _NewsroomContract.contract, event: "ContentPublished", logs: logs, sub: sub}, nil
}

// WatchContentPublished is a free log subscription operation binding the contract event 0x1ede735f9b446d8014022fed176848ac3894c54942bef9ff452f7ae42b50d5ae.
//
// Solidity: e ContentPublished(editor indexed address, contentId indexed uint256, uri string)
func (_NewsroomContract *NewsroomContractFilterer) WatchContentPublished(opts *bind.WatchOpts, sink chan<- *NewsroomContractContentPublished, editor []common.Address, contentId []*big.Int) (event.Subscription, error) {

	var editorRule []interface{}
	for _, editorItem := range editor {
		editorRule = append(editorRule, editorItem)
	}
	var contentIdRule []interface{}
	for _, contentIdItem := range contentId {
		contentIdRule = append(contentIdRule, contentIdItem)
	}

	logs, sub, err := _NewsroomContract.contract.WatchLogs(opts, "ContentPublished", editorRule, contentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewsroomContractContentPublished)
				if err := _NewsroomContract.contract.UnpackLog(event, "ContentPublished", log); err != nil {
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

// NewsroomContractNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the NewsroomContract contract.
type NewsroomContractNameChangedIterator struct {
	Event *NewsroomContractNameChanged // Event containing the contract specifics and raw log

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
func (it *NewsroomContractNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewsroomContractNameChanged)
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
		it.Event = new(NewsroomContractNameChanged)
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
func (it *NewsroomContractNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewsroomContractNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewsroomContractNameChanged represents a NameChanged event raised by the NewsroomContract contract.
type NewsroomContractNameChanged struct {
	NewName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0x4737457377f528cc8afd815f73ecb8b05df80d047dbffc41c17750a4033592bc.
//
// Solidity: e NameChanged(newName string)
func (_NewsroomContract *NewsroomContractFilterer) FilterNameChanged(opts *bind.FilterOpts) (*NewsroomContractNameChangedIterator, error) {

	logs, sub, err := _NewsroomContract.contract.FilterLogs(opts, "NameChanged")
	if err != nil {
		return nil, err
	}
	return &NewsroomContractNameChangedIterator{contract: _NewsroomContract.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0x4737457377f528cc8afd815f73ecb8b05df80d047dbffc41c17750a4033592bc.
//
// Solidity: e NameChanged(newName string)
func (_NewsroomContract *NewsroomContractFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *NewsroomContractNameChanged) (event.Subscription, error) {

	logs, sub, err := _NewsroomContract.contract.WatchLogs(opts, "NameChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewsroomContractNameChanged)
				if err := _NewsroomContract.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// NewsroomContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the NewsroomContract contract.
type NewsroomContractOwnershipRenouncedIterator struct {
	Event *NewsroomContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *NewsroomContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewsroomContractOwnershipRenounced)
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
		it.Event = new(NewsroomContractOwnershipRenounced)
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
func (it *NewsroomContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewsroomContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewsroomContractOwnershipRenounced represents a OwnershipRenounced event raised by the NewsroomContract contract.
type NewsroomContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_NewsroomContract *NewsroomContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*NewsroomContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _NewsroomContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NewsroomContractOwnershipRenouncedIterator{contract: _NewsroomContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_NewsroomContract *NewsroomContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *NewsroomContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _NewsroomContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewsroomContractOwnershipRenounced)
				if err := _NewsroomContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// NewsroomContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NewsroomContract contract.
type NewsroomContractOwnershipTransferredIterator struct {
	Event *NewsroomContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NewsroomContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewsroomContractOwnershipTransferred)
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
		it.Event = new(NewsroomContractOwnershipTransferred)
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
func (it *NewsroomContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewsroomContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewsroomContractOwnershipTransferred represents a OwnershipTransferred event raised by the NewsroomContract contract.
type NewsroomContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_NewsroomContract *NewsroomContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NewsroomContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NewsroomContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NewsroomContractOwnershipTransferredIterator{contract: _NewsroomContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_NewsroomContract *NewsroomContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NewsroomContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NewsroomContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewsroomContractOwnershipTransferred)
				if err := _NewsroomContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// NewsroomContractRevisionSignedIterator is returned from FilterRevisionSigned and is used to iterate over the raw logs and unpacked data for RevisionSigned events raised by the NewsroomContract contract.
type NewsroomContractRevisionSignedIterator struct {
	Event *NewsroomContractRevisionSigned // Event containing the contract specifics and raw log

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
func (it *NewsroomContractRevisionSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewsroomContractRevisionSigned)
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
		it.Event = new(NewsroomContractRevisionSigned)
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
func (it *NewsroomContractRevisionSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewsroomContractRevisionSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewsroomContractRevisionSigned represents a RevisionSigned event raised by the NewsroomContract contract.
type NewsroomContractRevisionSigned struct {
	ContentId  *big.Int
	RevisionId *big.Int
	Author     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevisionSigned is a free log retrieval operation binding the contract event 0x605611fc50d3effbe4af88e82f5daebfcffe0fb8f3b34ed32f1a746290ccbc61.
//
// Solidity: e RevisionSigned(contentId indexed uint256, revisionId indexed uint256, author indexed address)
func (_NewsroomContract *NewsroomContractFilterer) FilterRevisionSigned(opts *bind.FilterOpts, contentId []*big.Int, revisionId []*big.Int, author []common.Address) (*NewsroomContractRevisionSignedIterator, error) {

	var contentIdRule []interface{}
	for _, contentIdItem := range contentId {
		contentIdRule = append(contentIdRule, contentIdItem)
	}
	var revisionIdRule []interface{}
	for _, revisionIdItem := range revisionId {
		revisionIdRule = append(revisionIdRule, revisionIdItem)
	}
	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}

	logs, sub, err := _NewsroomContract.contract.FilterLogs(opts, "RevisionSigned", contentIdRule, revisionIdRule, authorRule)
	if err != nil {
		return nil, err
	}
	return &NewsroomContractRevisionSignedIterator{contract: _NewsroomContract.contract, event: "RevisionSigned", logs: logs, sub: sub}, nil
}

// WatchRevisionSigned is a free log subscription operation binding the contract event 0x605611fc50d3effbe4af88e82f5daebfcffe0fb8f3b34ed32f1a746290ccbc61.
//
// Solidity: e RevisionSigned(contentId indexed uint256, revisionId indexed uint256, author indexed address)
func (_NewsroomContract *NewsroomContractFilterer) WatchRevisionSigned(opts *bind.WatchOpts, sink chan<- *NewsroomContractRevisionSigned, contentId []*big.Int, revisionId []*big.Int, author []common.Address) (event.Subscription, error) {

	var contentIdRule []interface{}
	for _, contentIdItem := range contentId {
		contentIdRule = append(contentIdRule, contentIdItem)
	}
	var revisionIdRule []interface{}
	for _, revisionIdItem := range revisionId {
		revisionIdRule = append(revisionIdRule, revisionIdItem)
	}
	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}

	logs, sub, err := _NewsroomContract.contract.WatchLogs(opts, "RevisionSigned", contentIdRule, revisionIdRule, authorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewsroomContractRevisionSigned)
				if err := _NewsroomContract.contract.UnpackLog(event, "RevisionSigned", log); err != nil {
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

// NewsroomContractRevisionUpdatedIterator is returned from FilterRevisionUpdated and is used to iterate over the raw logs and unpacked data for RevisionUpdated events raised by the NewsroomContract contract.
type NewsroomContractRevisionUpdatedIterator struct {
	Event *NewsroomContractRevisionUpdated // Event containing the contract specifics and raw log

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
func (it *NewsroomContractRevisionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewsroomContractRevisionUpdated)
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
		it.Event = new(NewsroomContractRevisionUpdated)
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
func (it *NewsroomContractRevisionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewsroomContractRevisionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewsroomContractRevisionUpdated represents a RevisionUpdated event raised by the NewsroomContract contract.
type NewsroomContractRevisionUpdated struct {
	Editor     common.Address
	ContentId  *big.Int
	RevisionId *big.Int
	Uri        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevisionUpdated is a free log retrieval operation binding the contract event 0x18b6b5c485f8822a270464dd544d0715597dc8f1a007ee2b0252b62b8b9fb390.
//
// Solidity: e RevisionUpdated(editor indexed address, contentId indexed uint256, revisionId indexed uint256, uri string)
func (_NewsroomContract *NewsroomContractFilterer) FilterRevisionUpdated(opts *bind.FilterOpts, editor []common.Address, contentId []*big.Int, revisionId []*big.Int) (*NewsroomContractRevisionUpdatedIterator, error) {

	var editorRule []interface{}
	for _, editorItem := range editor {
		editorRule = append(editorRule, editorItem)
	}
	var contentIdRule []interface{}
	for _, contentIdItem := range contentId {
		contentIdRule = append(contentIdRule, contentIdItem)
	}
	var revisionIdRule []interface{}
	for _, revisionIdItem := range revisionId {
		revisionIdRule = append(revisionIdRule, revisionIdItem)
	}

	logs, sub, err := _NewsroomContract.contract.FilterLogs(opts, "RevisionUpdated", editorRule, contentIdRule, revisionIdRule)
	if err != nil {
		return nil, err
	}
	return &NewsroomContractRevisionUpdatedIterator{contract: _NewsroomContract.contract, event: "RevisionUpdated", logs: logs, sub: sub}, nil
}

// WatchRevisionUpdated is a free log subscription operation binding the contract event 0x18b6b5c485f8822a270464dd544d0715597dc8f1a007ee2b0252b62b8b9fb390.
//
// Solidity: e RevisionUpdated(editor indexed address, contentId indexed uint256, revisionId indexed uint256, uri string)
func (_NewsroomContract *NewsroomContractFilterer) WatchRevisionUpdated(opts *bind.WatchOpts, sink chan<- *NewsroomContractRevisionUpdated, editor []common.Address, contentId []*big.Int, revisionId []*big.Int) (event.Subscription, error) {

	var editorRule []interface{}
	for _, editorItem := range editor {
		editorRule = append(editorRule, editorItem)
	}
	var contentIdRule []interface{}
	for _, contentIdItem := range contentId {
		contentIdRule = append(contentIdRule, contentIdItem)
	}
	var revisionIdRule []interface{}
	for _, revisionIdItem := range revisionId {
		revisionIdRule = append(revisionIdRule, revisionIdItem)
	}

	logs, sub, err := _NewsroomContract.contract.WatchLogs(opts, "RevisionUpdated", editorRule, contentIdRule, revisionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewsroomContractRevisionUpdated)
				if err := _NewsroomContract.contract.UnpackLog(event, "RevisionUpdated", log); err != nil {
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

// NewsroomContractRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the NewsroomContract contract.
type NewsroomContractRoleAddedIterator struct {
	Event *NewsroomContractRoleAdded // Event containing the contract specifics and raw log

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
func (it *NewsroomContractRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewsroomContractRoleAdded)
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
		it.Event = new(NewsroomContractRoleAdded)
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
func (it *NewsroomContractRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewsroomContractRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewsroomContractRoleAdded represents a RoleAdded event raised by the NewsroomContract contract.
type NewsroomContractRoleAdded struct {
	Granter common.Address
	Grantee common.Address
	Role    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xa40c1dc2b34f6b51c3ea614b688f243e50047ed9fa3ea19010303d70dac781ed.
//
// Solidity: e RoleAdded(granter indexed address, grantee indexed address, role string)
func (_NewsroomContract *NewsroomContractFilterer) FilterRoleAdded(opts *bind.FilterOpts, granter []common.Address, grantee []common.Address) (*NewsroomContractRoleAddedIterator, error) {

	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _NewsroomContract.contract.FilterLogs(opts, "RoleAdded", granterRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return &NewsroomContractRoleAddedIterator{contract: _NewsroomContract.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xa40c1dc2b34f6b51c3ea614b688f243e50047ed9fa3ea19010303d70dac781ed.
//
// Solidity: e RoleAdded(granter indexed address, grantee indexed address, role string)
func (_NewsroomContract *NewsroomContractFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *NewsroomContractRoleAdded, granter []common.Address, grantee []common.Address) (event.Subscription, error) {

	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _NewsroomContract.contract.WatchLogs(opts, "RoleAdded", granterRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewsroomContractRoleAdded)
				if err := _NewsroomContract.contract.UnpackLog(event, "RoleAdded", log); err != nil {
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

// NewsroomContractRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the NewsroomContract contract.
type NewsroomContractRoleRemovedIterator struct {
	Event *NewsroomContractRoleRemoved // Event containing the contract specifics and raw log

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
func (it *NewsroomContractRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewsroomContractRoleRemoved)
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
		it.Event = new(NewsroomContractRoleRemoved)
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
func (it *NewsroomContractRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewsroomContractRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewsroomContractRoleRemoved represents a RoleRemoved event raised by the NewsroomContract contract.
type NewsroomContractRoleRemoved struct {
	Granter common.Address
	Grantee common.Address
	Role    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0x6a52fb0cb0e75e6a6721483d2e539b38273ec6fe95b648a41e1a901594aeccb8.
//
// Solidity: e RoleRemoved(granter indexed address, grantee indexed address, role string)
func (_NewsroomContract *NewsroomContractFilterer) FilterRoleRemoved(opts *bind.FilterOpts, granter []common.Address, grantee []common.Address) (*NewsroomContractRoleRemovedIterator, error) {

	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _NewsroomContract.contract.FilterLogs(opts, "RoleRemoved", granterRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return &NewsroomContractRoleRemovedIterator{contract: _NewsroomContract.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0x6a52fb0cb0e75e6a6721483d2e539b38273ec6fe95b648a41e1a901594aeccb8.
//
// Solidity: e RoleRemoved(granter indexed address, grantee indexed address, role string)
func (_NewsroomContract *NewsroomContractFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *NewsroomContractRoleRemoved, granter []common.Address, grantee []common.Address) (event.Subscription, error) {

	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _NewsroomContract.contract.WatchLogs(opts, "RoleRemoved", granterRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewsroomContractRoleRemoved)
				if err := _NewsroomContract.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
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
