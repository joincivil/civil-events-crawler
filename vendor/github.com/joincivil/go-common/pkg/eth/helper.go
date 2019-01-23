package eth

import (
	"crypto/ecdsa"
	"math/big"

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

// NewSimulatedBackendHelper creates a new Helper using an ethereum SimulatedBackend
func NewSimulatedBackendHelper() (*Helper, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 1000000000)

	accounts := make(map[string]Account)
	accounts["genesis"] = Account{Key: key, Auth: auth, Address: GetEthAddressFromPrivateKey(key)}

	return &Helper{
		Blockchain: blockchain,
		Key:        key,
		Auth:       auth,
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

// AddAccount generates a new account and adds it to the account mapping
func (h *Helper) AddAccount(name string) (Account, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return Account{}, err
	}
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)

	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}

	h.Accounts[name] = Account{
		Key:     key,
		Auth:    auth,
		Address: crypto.PubkeyToAddress(key.PublicKey),
	}

	return h.Accounts[name], nil
}
