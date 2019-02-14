package eth

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

// Helper provides methods to create Ethereum transactions
type Helper struct {
	Blockchain bind.ContractBackend
	Key        *ecdsa.PrivateKey
	Auth       *bind.TransactOpts
	Accounts   map[string]Account
}

// Account groups a private key, authentication, and ETH address
type Account struct {
	Key     *ecdsa.PrivateKey
	Auth    *bind.TransactOpts
	Address common.Address
}

// NewETHClientHelper creates a new Helper using an ethclient with the provided URL
// accountKeys is a mapping of name->privateKey hex string
func NewETHClientHelper(ethAPIURL string, accountKeys map[string]string) (*Helper, error) {
	blockchain, err := ethclient.Dial(ethAPIURL)
	if err != nil {
		return nil, err
	}

	accounts := make(map[string]Account)
	for keyName, privateKey := range accountKeys {
		account, err := AccountFromPK(privateKey)
		if err != nil {
			return nil, err
		}
		accounts[keyName] = account
	}

	helper := &Helper{
		Blockchain: blockchain,
		Accounts:   accounts,
	}

	if (accounts["default"] != Account{}) {
		helper.Key = accounts["default"].Key
		helper.Auth = accounts["default"].Auth
	}

	return helper, nil
}

// NewSimulatedBackendHelper creates a new Helper using an ethereum SimulatedBackend
// generates accounts for "genesis", "alice", "bob", "carol", "dan", "erin"
func NewSimulatedBackendHelper() (*Helper, error) {
	alloc := make(core.GenesisAlloc)

	accounts := make(map[string]Account)

	accountNames := []string{"genesis", "alice", "bob", "carol", "dan", "eric"}

	for _, name := range accountNames {
		account, err := MakeAccount()
		if err != nil {
			return nil, err
		}
		accounts[name] = account
		alloc[account.Address] = core.GenesisAccount{Balance: big.NewInt(9223372036854775807)}
	}

	blockchain := backends.NewSimulatedBackend(alloc, 1000000000)

	return &Helper{
		Blockchain: blockchain,
		Key:        accounts["genesis"].Key,
		Auth:       accounts["genesis"].Auth,
		Accounts:   accounts,
	}, nil
}

// Transact creates a barebones bind.TransactOpts using the helper's account
func (h *Helper) Transact() *bind.TransactOpts {
	return &bind.TransactOpts{
		From:   h.Auth.From,
		Signer: h.Auth.Signer,
		Value:  nil,
	}
}

// TransactWithGasLimit creates a bind.TransactOpts object with a default gas limit using the helper's account
func (h *Helper) TransactWithGasLimit() *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     h.Auth.From,
		Signer:   h.Auth.Signer,
		Value:    nil,
		GasLimit: 100000,
	}
}

// Call creates an empty bind.CallOpts instance
func (h *Helper) Call() *bind.CallOpts {
	return &bind.CallOpts{}
}

// AccountFromPK constructs an Account from the provided ECDSA private key hex string
func AccountFromPK(privateKey string) (Account, error) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return Account{}, err
	}
	auth := bind.NewKeyedTransactor(pk)
	return Account{Key: pk, Auth: auth, Address: GetEthAddressFromPrivateKey(pk)}, nil
}

// AddAccount generates a new account and adds it to the account mapping
func (h *Helper) AddAccount(name string) (Account, error) {
	account, err := MakeAccount()
	if err != nil {
		return Account{}, err
	}

	h.Accounts[name] = account

	return h.Accounts[name], nil
}

// MakeAccount generates a new random Account
func MakeAccount() (Account, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return Account{}, err
	}
	auth := bind.NewKeyedTransactor(key)
	return Account{
		Key:     key,
		Auth:    auth,
		Address: crypto.PubkeyToAddress(key.PublicKey),
	}, nil

}
