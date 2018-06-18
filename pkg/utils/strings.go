// Package utils contains various common utils separate by utility types
package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"regexp"
)

var (
	validEthAddressExp = regexp.MustCompile(`^(http|https|ws|wss):\/\/((.+?)\.(.{2,5})|localhost|127\.0\.0\.1)(\:[0-9]{2,})*.*$`)
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
