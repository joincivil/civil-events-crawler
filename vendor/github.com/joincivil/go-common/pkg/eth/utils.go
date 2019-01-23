package eth

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// NormalizeEthAddress takes a string address to normalize the
// case of the ethereum address when it is a string.
// Runs through common.Address.Hex().
func NormalizeEthAddress(addr string) string {
	address := common.HexToAddress(addr)
	return address.Hex()
}

// ABILinkLibrary replaces references to a library
// with the actual addresses to those library contracts
func ABILinkLibrary(bin string, libraryName string, libraryAddress common.Address) string {
	libstr := fmt.Sprintf("_+%v_+", libraryName)
	libraryRexp := regexp.MustCompile(libstr)

	// Remove the 0x prefix from those addresses, just need the actual hex string
	cleanLibraryAddr := strings.Replace(libraryAddress.Hex(), "0x", "", -1)

	modifiedBin := libraryRexp.ReplaceAllString(bin, cleanLibraryAddr)

	return modifiedBin
}

// DeployContractWithLinks patches a contract bin with provided library addresses
func DeployContractWithLinks(
	opts *bind.TransactOpts,
	backend bind.ContractBackend,
	abiString string,
	bin string,
	libraries map[string]common.Address,
) (common.Address, *types.Transaction, *bind.BoundContract, error) {

	linkedBin := bin
	for libraryName, libraryAddress := range libraries {
		linkedBin = ABILinkLibrary(bin, libraryName, libraryAddress)
	}

	parsed, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return bind.DeployContract(opts, parsed, common.FromHex(linkedBin), backend)

}
