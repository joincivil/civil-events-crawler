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
			MultiSigWalletContractType: {
				name:        "MultiSigWalletContract",
				simpleName:  "multisigwallet",
				abiStr:      contract.MultiSigWalletContractABI,
				importPath:  "github.com/joincivil/go-common/pkg/generated/contract",
				typePackage: "contract",
			},
			MultiSigWalletFactoryContractType: {
				name:        "MultiSigWalletFactoryContract",
				simpleName:  "multisigwalletfactory",
				abiStr:      contract.MultiSigWalletFactoryContractABI,
				importPath:  "github.com/joincivil/go-common/pkg/generated/contract",
				typePackage: "contract",
			},
			NewsroomFactoryType: {
				name:        "NewsroomFactory",
				simpleName:  "newsroomfactory",
				abiStr:      contract.NewsroomFactoryABI,
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
	//
	// Use FlagKey() to consistently create keys for this map
	//
	// NOTE: The currently disabled events match up items Civil is not processing.
	// Feel free to alter and update as needed.
	// Should be a config file or env vars in the future
	DisableCrawl = map[string]bool{
		FlagKey("CivilTCRContract", "GovernmentTransfered"): true,

		FlagKey("NewsroomContract", "ContentPublished"):   true,
		FlagKey("NewsroomContract", "OwnershipRenounced"): true,
		FlagKey("NewsroomContract", "RevisionSigned"):     true,
		FlagKey("NewsroomContract", "RoleAdded"):          true,
		FlagKey("NewsroomContract", "RoleRemoved"):        true,

		FlagKey("CVLTokenContract", "OwnershipRenounced"):   true,
		FlagKey("CVLTokenContract", "OwnershipTransferred"): true,
		FlagKey("CVLTokenContract", "Approval"):             true,

		FlagKey("ParameterizerContract", "RewardClaimed"): true,

		FlagKey("CivilPLCRVotingContract", "VotingRightsGranted"):   true,
		FlagKey("CivilPLCRVotingContract", "VotingRightsWithdrawn"): true,

		FlagKey("GovernmentContract", "AppellateSet"):                   true,
		FlagKey("GovernmentContract", "GovtReparameterizationProposal"): true,
		FlagKey("GovernmentContract", "NewConstSet"):                    true,
		FlagKey("GovernmentContract", "ParameterSet"):                   true,
		FlagKey("GovernmentContract", "ProposalExpired"):                true,
		FlagKey("GovernmentContract", "ProposalFailed"):                 true,
		FlagKey("GovernmentContract", "ProposalPassed"):                 true,
	}

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
	// NOTE: The currently enabled events match up with Civil needs.
	// Feel free to alter as needed for your deployments.
	// Should be a config file or env vars in the future
	EnableListener = map[string]bool{
		FlagKey("CivilTCRContract", "AppealGranted"):                 true,
		FlagKey("CivilTCRContract", "AppealRequested"):               true,
		FlagKey("CivilTCRContract", "Application"):                   true,
		FlagKey("CivilTCRContract", "ApplicationRemoved"):            true,
		FlagKey("CivilTCRContract", "ApplicationWhitelisted"):        true,
		FlagKey("CivilTCRContract", "Challenge"):                     true,
		FlagKey("CivilTCRContract", "ChallengeFailed"):               true,
		FlagKey("CivilTCRContract", "ChallengeSucceeded"):            true,
		FlagKey("CivilTCRContract", "Deposit"):                       true,
		FlagKey("CivilTCRContract", "FailedChallengeOverturned"):     true,
		FlagKey("CivilTCRContract", "GrantedAppealChallenged"):       true,
		FlagKey("CivilTCRContract", "GrantedAppealConfirmed"):        true,
		FlagKey("CivilTCRContract", "GrantedAppealOverturned"):       true,
		FlagKey("CivilTCRContract", "ListingRemoved"):                true,
		FlagKey("CivilTCRContract", "ListingWithdrawn"):              true,
		FlagKey("CivilTCRContract", "RewardClaimed"):                 true,
		FlagKey("CivilTCRContract", "SuccessfulChallengeOverturned"): true,
		FlagKey("CivilTCRContract", "TouchAndRemoved"):               true,
		FlagKey("CivilTCRContract", "Withdrawal"):                    true,
		// EnableAllListenersKey: true,
	}

	// NameToContractTypes is the map from a simple name to ContractType
	// To be kept up to date with supported contracts
	NameToContractTypes = CTypes{}
)
