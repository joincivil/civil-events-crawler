// Package strings contains various common utils for strings
package strings

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

var (
	validEthAddressExp = regexp.MustCompile(`^(http|https|ws|wss):\/\/((.+?)\.(.{2,5})|localhost|ethereum|127\.0\.0\.1)(\:[0-9]{2,})*.*$`)
)

// IsValidEthAPIURL returns true if the given string matches a valid
// eth endpoint URL
func IsValidEthAPIURL(url string) bool {
	return validEthAddressExp.MatchString(url)
}

// IsValidContractAddress returns true is the given string matches a valid
// smart contract address
func IsValidContractAddress(address string) bool {
	return common.IsHexAddress(address)
}

// RandomHexStr generates a hex string from a byte slice of n random numbers.
func RandomHexStr(n int) (string, error) {
	bys := make([]byte, n)
	if _, err := rand.Read(bys); err != nil {
		return "", err
	}
	return hex.EncodeToString(bys), nil
}

// ListCommonAddressToListString converts a list of common.address to list of string
func ListCommonAddressToListString(addresses []common.Address) []string {
	addressesString := make([]string, len(addresses))
	for i, address := range addresses {
		addressesString[i] = address.Hex()
	}
	return addressesString
}

// ListStringToListCommonAddress converts a list of strings to list of common.address
func ListStringToListCommonAddress(addresses []string) []common.Address {
	addressesCommon := make([]common.Address, len(addresses))
	for i, address := range addresses {
		addressesCommon[i] = common.HexToAddress(address)
	}
	return addressesCommon
}

// ListCommonAddressesToString converts a list of common.address to a comma delimited string
func ListCommonAddressesToString(addresses []common.Address) string {
	addressesString := ListCommonAddressToListString(addresses)
	return strings.Join(addressesString, ",")
}

// ListIntToListString converts a list of big.int to a list of string
func ListIntToListString(listInt []int) []string {
	listString := make([]string, len(listInt))
	for idx, i := range listInt {
		listString[idx] = strconv.Itoa(i)
	}
	return listString
}

// StringToCommonAddressesList converts a comma delimited string to a list of common.address
func StringToCommonAddressesList(addresses string) []common.Address {
	addressesString := strings.Split(addresses, ",")
	return ListStringToListCommonAddress(addressesString)
}
