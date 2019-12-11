package contractspecs

import "github.com/joincivil/go-common/pkg/generated/contract"

var (
	// ContractTypeToSpecs contains a map from ContractType to the contract specs,
	// which is the location of the contract and contract ABI, along with contract
	// metadata used to generate the watchers/filterers.
	// To be kept up to date with supported contracts
	// TODO(PN): Convert this to a YAML or JSON file that loads into this spec struct
	ContractTypeToSpecs = CSpecs{
		specs: map[ContractType]*ContractSpecs{
			CivilTcrContractType: {
				name:        "CivilTCRContract",
				simpleName:  "civiltcr",
				abiStr:      contract.CivilTCRContractABI,
				importPath:  "github.com/joincivil/go-common/pkg/generated/contract",
				typePackage: "contract",
			},
			NewsroomContractType: {
				name:        "NewsroomContract",
				simpleName:  "newsroom",
				abiStr:      contract.NewsroomContractABI,
				importPath:  "github.com/joincivil/go-common/pkg/generated/contract",
				typePackage: "contract",
			},
			CivilPLCRVotingContractType: {
				name:        "CivilPLCRVotingContract",
				simpleName:  "civilplcrvoting",
				abiStr:      contract.CivilPLCRVotingContractABI,
				importPath:  "github.com/joincivil/go-common/pkg/generated/contract",
				typePackage: "contract",
			},
			CVLTokenContractType: {
				name:        "CVLTokenContract",
				simpleName:  "cvltoken",
				abiStr:      contract.CVLTokenContractABI,
				importPath:  "github.com/joincivil/go-common/pkg/generated/contract",
				typePackage: "contract",
			},
			CivilParameterizerContractType: {
				name:        "ParameterizerContract",
				simpleName:  "civilparameterizer",
				abiStr:      contract.ParameterizerContractABI,
				importPath:  "github.com/joincivil/go-common/pkg/generated/contract",
				typePackage: "contract",
			},
			CivilGovernmentContractType: {
				name:        "GovernmentContract",
				simpleName:  "civilgovernment",
				abiStr:      contract.GovernmentContractABI,
				importPath:  "github.com/joincivil/go-common/pkg/generated/contract",
				typePackage: "contract",
			},
		},
	}

	// DisableCrawl is a map of contract:event keys to a bool to disable
	// crawling on a particular event.  The key is to be of the format
	// <contract type name in lower>:<event name in lower>
	// ex. newsroomcontract:watchcontentpublished.
	// If true, disables tracking of that event
	// Use FlagKey() to consistently create keys for this map
	DisableCrawl = map[string]bool{}

	// EnableListener is a map of contract:event keys to a bool to enable
	// listeners on a particular event. If DisableCrawl is set for the same event,
	// the configuration here will not be used.
	// The key is to be of the format <contract type name in lower>:<event name in lower>
	// ex. newsroomcontract:watchcontentpublished.
	// If true, enables websocket listener for that event
	//
	// Use FlagKey() to consistently create keys for this map
	//
	// To enable listeners for all the events, add an "all": true item to this
	// config map. Use the EnableAllListenersKey to set this value.
	EnableListener = map[string]bool{
		EnableAllListenersKey: true,
	}

	// NameToContractTypes is the map from a simple name to ContractType
	// To be kept up to date with supported contracts
	NameToContractTypes = CTypes{}
)
