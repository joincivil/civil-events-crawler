// Package contractutils contains utilities related to smart contracts
// and testing smart contracts
package contractutils

// NOTE: very specific to Civil's smart contract implementation

import (
	"math/big"
	"regexp"
	"strings"

	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "fmt"
	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
)

const (
	// rinkebyAddress = "wss://rinkeby.infura.io/ws"
	gasLimit = uint64(8000000000000)

	minDeposit                         = 10
	pMinDeposit                        = 100
	pDeposit                           = 150
	applyStageLength                   = 0 // 18000
	pApplyStageLength                  = 120
	commitStageLength                  = 18000
	pCommitStageLength                 = 120
	revealStageLength                  = 18000
	pRevealStageLength                 = 120
	dispensationPct                    = 50
	pDispensationPct                   = 50
	voteQuorum                         = 50
	pVoteQuorum                        = 50
	pProcessBy                         = 18000
	appealFeeAmount                    = 1000
	challengeAppealLength              = 18000
	requestAppealPhaseLength           = 36000
	judgeAppealPhaseLength             = 36000
	appealChallengeCommitStageLength   = 16000
	appealChallengeRevealStageLength   = 16000
	appealSupermajorityPercentage      = 66
	appealChallengeVoteDispensationPct = 66
)

// // SetupRinkebyClient returns an instance of the ethclient setup on Rinkeby
// func SetupRinkebyClient() (*ethclient.Client, error) {
// 	client, err := ethclient.Dial(rinkebyAddress)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return client, nil
// }

// SetupSimulatedClient returns an an instance of the simulated backend.
func SetupSimulatedClient(gasLimit uint64) (*backends.SimulatedBackend, *bind.TransactOpts) {
	key, _ := crypto.GenerateKey() // nolint: gosec
	auth := bind.NewKeyedTransactor(key)
	genAlloc := make(core.GenesisAlloc)
	genAlloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(9223372036854775807)}

	sim := backends.NewSimulatedBackend(genAlloc, gasLimit)
	return sim, auth
}

// AllTestContracts contains the values returned from SetupAllTestContracts
type AllTestContracts struct {
	TokenAddr         common.Address
	TokenContract     *contract.EIP20Contract
	PlcrAddr          common.Address
	PlcrContract      *contract.CivilPLCRVotingContract
	ParamAddr         common.Address
	ParamContract     *contract.ParameterizerContract
	GovtAddr          common.Address
	GovtContract      *contract.GovernmentContract
	CivilTcrAddr      common.Address
	CivilTcrContract  *contract.CivilTCRContract
	NewsroomAddr      common.Address
	NewsroomContract  *contract.NewsroomContract
	TokenTeleAddr     common.Address
	TokenTeleContract *contract.DummyTokenTelemetryContract
	Client            *backends.SimulatedBackend
	Auth              *bind.TransactOpts
}

// SetupAllTestContracts returns a struct with all the test contracts deployed to the
// simulated backend.
func SetupAllTestContracts() (*AllTestContracts, error) {
	client, auth := SetupSimulatedClient(gasLimit)

	// Deploy the required libraries
	asAddr, err := setupAttributeStoreContract(client, auth)
	if err != nil {
		log.Fatalf("Unable to deploy the attribute store lib: %v", err)
	}
	client.Commit()

	dllAddr, err := setupDLLContract(client, auth)
	if err != nil {
		log.Fatalf("Unable to deploy the DLL library: %v", err)
	}
	client.Commit()

	// Setup the TESTCVL token
	tokenAddr, eip20, err := setupTestEIP20Contract(client, auth)
	if err != nil {
		log.Fatalf("Unable to deploy a test token: %v", err)
	}

	client.Commit()

	balance, _ := eip20.BalanceOf(&bind.CallOpts{}, auth.From) // nolint: gosec
	approveOpts := &bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}

	tokenTeleAddr, tokenTele, err := setupTestDummyTokenTelemetryContract(client, auth)
	if err != nil {
		log.Fatalf("Unable to deploy a test dummy token telemetry contract: err: %v", err)
		return nil, err
	}

	_, err = eip20.Approve(approveOpts, tokenTeleAddr, balance)
	if err != nil {
		log.Errorf("Unable to approve user for token telemetry: %v", err)
		return nil, err
	}

	client.Commit()

	plcrAddr, plcr, err := setupTestCivilPLCRVotingContract(client, auth, tokenAddr, tokenTeleAddr, dllAddr, asAddr)
	if err != nil {
		log.Fatalf("Unable to deploy a test Civil PLCR voting contract: err: %v", err)
		return nil, err
	}

	_, err = eip20.Approve(approveOpts, plcrAddr, balance)
	if err != nil {
		log.Errorf("Unable to approve user for PLCR: %v", err)
		return nil, err
	}

	client.Commit()

	paramAddr, param, err := setupTestParameterizerContract(client, auth, tokenAddr, plcrAddr)
	if err != nil {
		log.Fatalf("Unable to deploy a test parameterizer contract: err: %v", err)
		return nil, err
	}

	_, err = eip20.Approve(approveOpts, paramAddr, balance)
	if err != nil {
		log.Errorf("Unable to approve user for parameterizer: %v", err)
		return nil, err
	}

	client.Commit()

	govtAddr, govt, err := setupTestGovernmentContract(client, auth, auth.From,
		auth.From, plcrAddr)
	if err != nil {
		log.Fatalf("Unable to deploy a test government contract: err: %v", err)
		return nil, err
	}

	_, err = eip20.Approve(approveOpts, govtAddr, balance)
	if err != nil {
		log.Errorf("Unable to approve user for govt: %v", err)
		return nil, err
	}

	client.Commit()

	civilTcrAddr, civilTcr, err := setupTestCivilTCRContract(client, auth, tokenAddr,
		plcrAddr, paramAddr, govtAddr, tokenTeleAddr)
	if err != nil {
		log.Fatalf("Unable to deploy a test civil tcr contract: err: %v", err)
		return nil, err
	}

	_, err = eip20.Approve(approveOpts, civilTcrAddr, balance)
	if err != nil {
		log.Errorf("Unable to approve user for tcr: %v", err)
		return nil, err
	}

	client.Commit()

	newsroomAddr, newsroom, err := setupTestNewsroomContract(client, auth)
	if err != nil {
		log.Fatalf("Unable to deploy a test newsroom contract: err: %v", err)
		return nil, err
	}

	_, err = eip20.Approve(approveOpts, newsroomAddr, balance)
	if err != nil {
		log.Errorf("Unable to approve user for newsroom: err: %v", err)
		return nil, err
	}

	client.Commit()

	return &AllTestContracts{
		TokenAddr:         tokenAddr,
		TokenContract:     eip20,
		PlcrAddr:          plcrAddr,
		PlcrContract:      plcr,
		ParamAddr:         paramAddr,
		ParamContract:     param,
		GovtAddr:          govtAddr,
		GovtContract:      govt,
		CivilTcrAddr:      civilTcrAddr,
		CivilTcrContract:  civilTcr,
		NewsroomAddr:      newsroomAddr,
		NewsroomContract:  newsroom,
		TokenTeleAddr:     tokenTeleAddr,
		TokenTeleContract: tokenTele,
		Client:            client,
		Auth:              auth,
	}, nil

}

// setupTestNewsroomContract deploys a test newsroom contract to the given ContractBackend.
func setupTestNewsroomContract(client bind.ContractBackend, auth *bind.TransactOpts) (common.Address,
	*contract.NewsroomContract, error) {
	address, _, contract, err := contract.DeployNewsroomContract(
		auth,
		client,
		"Your Newsroom Here",
		"newsroom.com/charter",
		[32]byte{},
	)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, contract, nil
}

// setupTestEIP20Contract deploys a test token contract to the given ContractBackend.
func setupTestEIP20Contract(client bind.ContractBackend, auth *bind.TransactOpts) (common.Address,
	*contract.EIP20Contract, error) {
	address, _, contract, err := contract.DeployEIP20Contract(
		auth,
		client,
		big.NewInt(9223372036854775807),
		"CivilTokenTest",
		18,
		"CVLT",
	)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, contract, nil
}

// modifyByteCodeWithDllAttrStore replaces references to the dll/addr store library
// with the actual addresses to those library contracts
func modifyByteCodeWithDllAttrStore(bin string, dllAddr common.Address,
	attrStoreAddr common.Address) string {
	// To add the DLL and AttributeStore addresses, replace occurrences of
	// _DLL__________ and _AttributeStore__________ with the respective contract addresses
	ddlRexp := regexp.MustCompile("_+DLL_+")
	asRexp := regexp.MustCompile("_+AttributeStore_+")

	// Remove the 0x prefix from those addresses, just need the actual hex string
	cleanDllAddr := strings.Replace(dllAddr.Hex(), "0x", "", -1)
	cleanAttrStoreAddr := strings.Replace(attrStoreAddr.Hex(), "0x", "", -1)

	modifiedBin := ddlRexp.ReplaceAllString(contract.CivilPLCRVotingContractBin, cleanDllAddr)
	modifiedBin = asRexp.ReplaceAllString(modifiedBin, cleanAttrStoreAddr)
	return modifiedBin
}

// deployCivilPLCRVotingContractModified deploys a new Ethereum contract, binding an instance of CivilPLCRVotingContract to it.
// Hacking this to modify the BIN to add the DLL and AttributeStore addresses.
func deployCivilPLCRVotingContractModified(auth *bind.TransactOpts, backend bind.ContractBackend,
	tokenAddr common.Address, telemetryAddr common.Address, dllAddr common.Address,
	attrStoreAddr common.Address) (common.Address, *types.Transaction, *contract.CivilPLCRVotingContract, error) {
	parsed, err := abi.JSON(strings.NewReader(contract.CivilPLCRVotingContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, _, err := bind.DeployContract(
		auth,
		parsed,
		common.FromHex(modifyByteCodeWithDllAttrStore(
			contract.CivilPLCRVotingContractBin,
			dllAddr,
			attrStoreAddr,
		)),
		backend,
		tokenAddr,
		telemetryAddr,
	)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	newContract, err := contract.NewCivilPLCRVotingContract(address, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, newContract, nil
}

// setupTestCivilPLCRVotingContract deploys a test Civil PLCR voting contract to the given ContractBackend.
func setupTestCivilPLCRVotingContract(client bind.ContractBackend, auth *bind.TransactOpts,
	tokenAddr common.Address, teleAddr common.Address, dllAddr common.Address,
	attributeStoreAddr common.Address) (common.Address, *contract.CivilPLCRVotingContract, error) {
	address, _, contract, err := deployCivilPLCRVotingContractModified(
		auth,
		client,
		tokenAddr,
		teleAddr,
		dllAddr,
		attributeStoreAddr,
	)
	if err != nil {
		return common.Address{}, nil, err
	}

	return address, contract, nil
}

// setupTestDummyTokenTelemetryContract deploys a test dummy token telemetry contract to the given ContractBackend.
func setupTestDummyTokenTelemetryContract(client bind.ContractBackend, auth *bind.TransactOpts) (
	common.Address, *contract.DummyTokenTelemetryContract, error) {
	address, _, contract, err := contract.DeployDummyTokenTelemetryContract(
		auth,
		client,
	)
	if err != nil {
		return common.Address{}, nil, err
	}

	return address, contract, nil
}

// setupAttributeStoreContract deploys the required AttributeStore contract
func setupAttributeStoreContract(client bind.ContractBackend, auth *bind.TransactOpts) (common.Address, error) {
	address, _, _, err := contract.DeployAttributeStoreContract(
		auth,
		client,
	)
	if err != nil {
		return common.Address{}, err
	}

	return address, nil
}

// setupDLLContract deploys the required DLL contract
func setupDLLContract(client bind.ContractBackend, auth *bind.TransactOpts) (common.Address, error) {
	address, _, _, err := contract.DeployDLLContract(
		auth,
		client,
	)
	if err != nil {
		return common.Address{}, err
	}

	return address, nil
}

// setupTestParameterizerContract deploys a test parameterizer voting contract to the given ContractBackend.
func setupTestParameterizerContract(client bind.ContractBackend, auth *bind.TransactOpts,
	tokenAddr common.Address, plcrAddr common.Address) (common.Address, *contract.ParameterizerContract, error) {
	params := []*big.Int{
		big.NewInt(minDeposit),
		big.NewInt(pMinDeposit),
		big.NewInt(applyStageLength),
		big.NewInt(pApplyStageLength),
		big.NewInt(commitStageLength),
		big.NewInt(pCommitStageLength),
		big.NewInt(revealStageLength),
		big.NewInt(pRevealStageLength),
		big.NewInt(dispensationPct),
		big.NewInt(pDispensationPct),
		big.NewInt(voteQuorum),
		big.NewInt(pVoteQuorum),
		big.NewInt(pProcessBy),
		big.NewInt(challengeAppealLength),
		big.NewInt(appealChallengeCommitStageLength),
		big.NewInt(appealChallengeRevealStageLength),
	}
	address, _, contract, err := contract.DeployParameterizerContract(
		auth,
		client,
		tokenAddr,
		plcrAddr,
		params,
	)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, contract, nil
}

// setupTestGovernmentContract deploys a test government contract to the given ContractBackend.
func setupTestGovernmentContract(client bind.ContractBackend, auth *bind.TransactOpts,
	appellateAddr common.Address, governmentControllerAddr common.Address,
	plcrAddr common.Address) (common.Address, *contract.GovernmentContract, error) {
	address, _, contract, err := contract.DeployGovernmentContract(
		auth,
		client,
		appellateAddr,
		governmentControllerAddr,
		plcrAddr,
		big.NewInt(appealFeeAmount),
		big.NewInt(requestAppealPhaseLength),
		big.NewInt(judgeAppealPhaseLength),
		big.NewInt(appealSupermajorityPercentage),
		big.NewInt(appealChallengeVoteDispensationPct),
		big.NewInt(pDeposit),
		big.NewInt(pCommitStageLength),
		big.NewInt(pRevealStageLength),
		[32]byte{},
		"http://madeupURL.com",
	)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, contract, nil
}

// setupTestCivilTCRContract deploys a test Civil TCR contract to the given ContractBackend.
func setupTestCivilTCRContract(client bind.ContractBackend, auth *bind.TransactOpts,
	tokenAddr common.Address, plcrAddr common.Address, paramAddr common.Address,
	govtAddr common.Address, tokenTeleAddr common.Address) (common.Address,
	*contract.CivilTCRContract, error) {
	address, _, contract, err := contract.DeployCivilTCRContract(
		auth,
		client,
		tokenAddr,
		plcrAddr,
		paramAddr,
		govtAddr,
		tokenTeleAddr,
	)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, contract, nil
}
