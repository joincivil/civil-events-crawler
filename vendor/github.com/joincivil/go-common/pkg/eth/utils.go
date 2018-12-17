package eth

import (
	"github.com/ethereum/go-ethereum/common"
)

// NormalizeEthAddress takes a string address to normalize the
// case of the ethereum address when it is a string.
// Runs through common.Address.Hex().
func NormalizeEthAddress(addr string) string {
	address := common.HexToAddress(addr)
	return address.Hex()
}
