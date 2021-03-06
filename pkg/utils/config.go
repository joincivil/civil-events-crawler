// Package utils contains various common utils separate by utility types
package utils

import (
	"context"
	"errors"
	"fmt"
	"strings"

	log "github.com/golang/glog"

	cconfig "github.com/joincivil/go-common/pkg/config"
	cstrings "github.com/joincivil/go-common/pkg/strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kelseyhightower/envconfig"
	"github.com/shurcooL/graphql"
)

const (
	envVarPrefix = "crawl"
)

// NOTE(PN): After envconfig populates CrawlerConfig with the environment vars,
// there is nothing preventing the CrawlerConfig fields from being mutated.

// CrawlerConfig is the master config for the crawler derived from environment
// variables.
type CrawlerConfig struct {
	EthAPIURL     string `envconfig:"eth_api_url" required:"true" desc:"Ethereum HTTP API address"`
	EthWsAPIURL   string `envconfig:"eth_ws_api_url" desc:"Ethereum Websocket API address (optional, disables watchers if empty)"`
	EthStartBlock uint64 `envconfig:"eth_start_block" desc:"Sets the start Eth block (default 0)" default:"0"`

	// Enables polling mode, disables websockets
	PollingEnabled      bool `envconfig:"polling_enabled" desc:"Enable polling only mode (true disables listeners)"`
	PollingIntervalSecs int  `envconfig:"polling_int_secs" desc:"Sets the polling interval"`

	// CivilListingsGraphqlURL enables call to retrieve newsroom listings from Civil.
	// Should pass in the URL to the GraphQL endpoint to enable.
	CivilListingsGraphqlURL string `envconfig:"civil_listing_graphql_url" desc:"URL of the Civil Listings GraphQL endpoint"`

	// ContractAddresses map a contract type to a string of contract addresses.  If there are more than 1
	// contract to be tracked for a particular type, delimit the addresses with '|'.
	ContractAddresses   map[string]string           `split_words:"true" required:"true" desc:"<contract name>:<contract addr>. Delimit contract address with '|' for multiple addresses"`
	ContractAddressObjs map[string][]common.Address `ignored:"true"`

	PersisterType             cconfig.PersisterType `ignored:"true"`
	PersisterTypeName         string                `split_words:"true" required:"true" desc:"Sets the persister type to use"`
	PersisterPostgresAddress  string                `split_words:"true" desc:"If persister type is Postgresql, sets the address"`
	PersisterPostgresPort     int                   `split_words:"true" desc:"If persister type is Postgresql, sets the port"`
	PersisterPostgresDbname   string                `split_words:"true" desc:"If persister type is Postgresql, sets the database name"`
	PersisterPostgresUser     string                `split_words:"true" desc:"If persister type is Postgresql, sets the database user"`
	PersisterPostgresPw       string                `split_words:"true" desc:"If persister type is Postgresql, sets the database password"`
	PersisterPostgresMaxConns *int                  `split_words:"true" desc:"If persister type is Postgresql, sets the max conns in pool"`
	PersisterPostgresMaxIdle  *int                  `split_words:"true" desc:"If persister type is Postgresql, sets the max idle conns in pool"`
	PersisterPostgresConnLife *int                  `split_words:"true" desc:"If persister type is Postgresql, sets the max conn lifetime in secs"`

	PubSubProjectID string `split_words:"true" desc:"Sets the Google Cloud PubSub Project ID"`
	PubSubTopicName string `split_words:"true" desc:"Sets the Google Cloud PubSub Topic name"`

	// VersionNumber is the version of DB for postgres persistence
	VersionNumber string `split_words:"true" desc:"Sets the version for table"`

	// Configuration for supported error reporting (Stackdriver, Sentry)
	StackdriverProjectID      string `split_words:"true" desc:"Sets the Stackdriver Google Cloud project ID. If empty, will disable logging"`
	StackdriverServiceName    string `split_words:"true" desc:"Sets the Stackdriver service name"`
	StackdriverServiceVersion string `split_words:"true" desc:"Sets the Stackdriver service version"`

	SentryDsn        string `split_words:"true" desc:"Sets the Sentry DSN"`
	SentryEnv        string `split_words:"true" desc:"Sets the Sentry environment"`
	SentryLoggerName string `split_words:"true" desc:"Sets the Sentry logger name"`
	SentryRelease    string `split_words:"true" desc:"Sets the Sentry release value"`

	// Configuration for pprof profiling
	PprofEnable                bool   `split_words:"true" desc:"Enables the local pprof endpoints for debugging and profiling"`
	CloudProfileProjectID      string `split_words:"true" desc:"Sets the cloud profiler Google Cloud project ID. If empty, will disable cloud profiler agent"`
	CloudProfileServiceName    string `split_words:"true" desc:"Sets the service name to with in the Google Cloud profiler"`
	CloudProfileServiceVersion string `split_words:"true" desc:"Sets the Google Cloud profiler service version"`

	PreemptSecs *int `split_words:"true" desc:"Sets the secs delay before listener prempt restart. 0 if no preempt"`
}

// Edge represents an edge field in the query
type Edge struct {
	Node struct {
		Name            graphql.String
		ContractAddress graphql.String
		Whitelisted     graphql.Boolean
		ApprovalDate    graphql.Int
		LastGovState    graphql.String
	}
}

// PageInfo represents a pageinfo field in the query
type PageInfo struct {
	EndCursor   graphql.String
	HasNextPage graphql.Boolean
}

var allListingAddressesQuery struct {
	AllListingAddresses []graphql.String `graphql:"allListingAddresses"`
}

var allMultiSigAddressesQuery struct {
	AllMultiSigAddresses []graphql.String `graphql:"allMultiSigAddresses"`
}

// FetchListingAddresses retrieves the list of Civil newsroom addresses if given
// the endpoint URL
func (c *CrawlerConfig) FetchListingAddresses() error {
	if c.CivilListingsGraphqlURL == "" {
		return nil
	}

	var addressStrings []string
	var addressObjs []common.Address

	newsroomContractName := "newsroom"
	if c.ContractAddresses[newsroomContractName] != "" {
		addressStrings = strings.Split(c.ContractAddresses[newsroomContractName], "|")
	}

	client := graphql.NewClient(c.CivilListingsGraphqlURL, nil)

	// Fetch all the known listing addreses
	q := allListingAddressesQuery
	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		return err
	}
	// Look at all the addresses and add to slices
	addrs := q.AllListingAddresses
	for _, addr := range addrs {
		addressStrings = append(addressStrings, string(addr))
		addressObjs = append(addressObjs, common.HexToAddress(string(addr)))
	}

	c.ContractAddresses[newsroomContractName] = strings.Join(addressStrings, "|")
	c.ContractAddressObjs[newsroomContractName] = append(
		c.ContractAddressObjs[newsroomContractName], addressObjs...)

	return nil
}

// FetchMultiSigAddresses retrieves the list of Civil multi sig addresses if given
// the endpoint URL
func (c *CrawlerConfig) FetchMultiSigAddresses() error {
	if c.CivilListingsGraphqlURL == "" {
		return nil
	}

	var addressStrings []string
	var addressObjs []common.Address

	multiSigContractName := "multisigwallet"
	if c.ContractAddresses[multiSigContractName] != "" {
		addressStrings = strings.Split(c.ContractAddresses[multiSigContractName], "|")
	}

	client := graphql.NewClient(c.CivilListingsGraphqlURL, nil)

	// Fetch all the known multi sig addreses
	q := allMultiSigAddressesQuery
	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		return err
	}
	// Look at all the addresses and add to slices
	addrs := q.AllMultiSigAddresses
	for _, addr := range addrs {
		addressStrings = append(addressStrings, string(addr))
		addressObjs = append(addressObjs, common.HexToAddress(string(addr)))
	}

	c.ContractAddresses[multiSigContractName] = strings.Join(addressStrings, "|")
	c.ContractAddressObjs[multiSigContractName] = append(
		c.ContractAddressObjs[multiSigContractName], addressObjs...)

	return nil
}

// PersistType returns the persister type, implements PersisterConfig
func (c *CrawlerConfig) PersistType() cconfig.PersisterType {
	return c.PersisterType
}

// PostgresAddress returns the postgres persister address, implements PersisterConfig
func (c *CrawlerConfig) Address() string {
	return c.PersisterPostgresAddress
}

// PostgresPort returns the postgres persister port, implements PersisterConfig
func (c *CrawlerConfig) Port() int {
	return c.PersisterPostgresPort
}

// PostgresDbname returns the postgres persister db name, implements PersisterConfig
func (c *CrawlerConfig) Dbname() string {
	return c.PersisterPostgresDbname
}

// PostgresUser returns the postgres persister user, implements PersisterConfig
func (c *CrawlerConfig) User() string {
	return c.PersisterPostgresUser
}

// PostgresPw returns the postgres persister password, implements PersisterConfig
func (c *CrawlerConfig) Password() string {
	return c.PersisterPostgresPw
}

// PoolMaxConns returns the max conns for a pool, if configured, implements PersisterConfig
func (c *CrawlerConfig) PoolMaxConns() *int {
	return c.PersisterPostgresMaxConns
}

// PoolMaxIdleConns returns the max idleconns for a pool, if configured, implements PersisterConfig
func (c *CrawlerConfig) PoolMaxIdleConns() *int {
	return c.PersisterPostgresMaxIdle
}

// PoolConnLifetimeSecs returns the conn lifetime for a pool, if configured, implements PersisterConfig
func (c *CrawlerConfig) PoolConnLifetimeSecs() *int {
	return c.PersisterPostgresConnLife
}

// DataVersion returns the version number from the config
func (c *CrawlerConfig) DataVersion() string {
	return c.VersionNumber
}

// OutputUsage prints the usage string to os.Stdout
func (c *CrawlerConfig) OutputUsage() {
	cconfig.OutputUsage(c, envVarPrefix, envVarPrefix)
}

// PopulateFromEnv processes the environment vars, populates CrawlerConfig
// with the respective values, and validates the values.
func (c *CrawlerConfig) PopulateFromEnv() error {
	envEnvVar := fmt.Sprintf("%v_ENV", strings.ToUpper(envVarPrefix))
	err := cconfig.PopulateFromDotEnv(envEnvVar)
	if err != nil {
		return err
	}

	err = envconfig.Process(envVarPrefix, c)
	if err != nil {
		return err
	}

	err = c.validateAPIURL()
	if err != nil {
		return err
	}

	err = c.validateContractAddresses()
	if err != nil {
		return err
	}

	c.populateContractAddressObjs()

	err = c.FetchListingAddresses()
	if err != nil {
		log.Errorf("Unable to fetch the Civil listing addresses: err: %v", err)
	}

	err = c.FetchMultiSigAddresses()
	if err != nil {
		log.Errorf("Unable to fetch the Civil multi sig addresses: err: %v", err)
	}

	err = c.populatePersisterType()
	if err != nil {
		return err
	}

	return c.validatePersister()
}

func (c *CrawlerConfig) populatePersisterType() error {
	var err error
	c.PersisterType, err = cconfig.PersisterTypeFromName(c.PersisterTypeName)
	return err
}

func (c *CrawlerConfig) populateContractAddressObjs() {
	c.ContractAddressObjs = map[string][]common.Address{}
	for contractName, addrStr := range c.ContractAddresses {
		addrs := splitStrByPipe(addrStr)
		addrList := make([]common.Address, len(addrs))
		for i, addr := range addrs {
			addrList[i] = common.HexToAddress(addr)
		}
		c.ContractAddressObjs[contractName] = addrList
	}
}

func (c *CrawlerConfig) validateContractAddresses() error {
	for _, addrStr := range c.ContractAddresses {
		if addrStr == "" {
			return fmt.Errorf("Invalid contract address: '%v'", addrStr)
		}
		addrs := splitStrByPipe(addrStr)
		for _, addr := range addrs {
			if !cstrings.IsValidContractAddress(addr) {
				return fmt.Errorf("Invalid contract address: '%v'", addr)
			}
		}
	}
	return nil
}

func (c *CrawlerConfig) validateAPIURL() error {
	if c.EthAPIURL == "" || !cstrings.IsValidEthAPIURL(c.EthAPIURL) {
		return fmt.Errorf("Invalid eth API URL: '%v'", c.EthAPIURL)
	}
	return nil
}

func (c *CrawlerConfig) validatePersister() error {
	var err error
	if c.PersisterType == cconfig.PersisterTypePostgresql {
		err = c.validatePostgresqlPersister()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CrawlerConfig) validatePostgresqlPersister() error {
	if c.PersisterPostgresAddress == "" {
		return errors.New("Postgresql address required")
	}
	if c.PersisterPostgresPort == 0 {
		return errors.New("Postgresql port required")
	}
	if c.PersisterPostgresDbname == "" {
		return errors.New("Postgresql db name required")
	}
	return nil
}

func splitStrByPipe(str string) []string {
	return strings.Split(str, "|")
}
