package eth

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// WebsocketPing periodically makes a call over the given websocket conn
// to the Eth node to ensure the connection stays alive.
// Since there is no built in facility to do pings with the go-eth lib,
// need to do this ourselves by making a eth_getHeaderByNumber RPC call.
// NOTE(PN): Need to ensure the client passed in is a websocket client.
// XXX(PN): I'm not sure of it's effectiveness since all nodes may be configured
// differently in regards to keeping connections open.
func WebsocketPing(client *ethclient.Client, killChan <-chan bool, pingIntervalSecs int) {
	log.Infof("Starting ws ping at %v sec intervals...", pingIntervalSecs)
Loop:
	for {
		select {
		case <-time.After(time.Duration(pingIntervalSecs) * time.Second):
			header, err := client.HeaderByNumber(context.TODO(), nil)
			if err != nil {
				log.Errorf("Ping header by number failed: err: %v", err)
				continue
			}
			log.Infof("Ping success: block number: %v", header.Number)

		case <-killChan:
			log.Infof("Closing websocket ping")
			break Loop
		}
	}
}

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
	params ...interface{},
) (common.Address, *types.Transaction, *bind.BoundContract, error) {

	for libraryName, libraryAddress := range libraries {
		bin = ABILinkLibrary(bin, libraryName, libraryAddress)
	}

	parsed, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	return bind.DeployContract(opts, parsed, common.FromHex(bin), backend, params...)
}
