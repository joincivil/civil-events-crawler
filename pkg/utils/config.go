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
	PollingEnabled      bool `envconfig:"polling_enabled" desc:"Enable polling mode (true disables listeners)"`
	PollingIntervalSecs int  `envconfig:"polling_int_secs" desc:"Sets the polling interval"`

	// CivilListingsGraphqlURL enables call to retrieve newsroom listings from Civil.
	// Should pass in the URL to the GraphQL endpoint to enable.
	CivilListingsGraphqlURL string `envconfig:"civil_listing_graphql_url" desc:"URL of the Civil Listings GraphQL endpoint"`

	// ContractAddresses map a contract type to a string of contract addresses.  If there are more than 1
	// contract to be tracked for a particular type, delimit the addresses with '|'.
	ContractAddresses   map[string]string           `split_words:"true" required:"true" desc:"<contract name>:<contract addr>. Delimit contract address with '|' for multiple addresses"`
	ContractAddressObjs map[string][]common.Address `ignored:"true"`

	PersisterType            cconfig.PersisterType `ignored:"true"`
	PersisterTypeName        string                `split_words:"true" required:"true" desc:"Sets the persister type to use"`
	PersisterPostgresAddress string                `split_words:"true" desc:"If persister type is Postgresql, sets the address"`
	PersisterPostgresPort    int                   `split_words:"true" desc:"If persister type is Postgresql, sets the port"`
	PersisterPostgresDbname  string                `split_words:"true" desc:"If persister type is Postgresql, sets the database name"`
	PersisterPostgresUser    string                `split_words:"true" desc:"If persister type is Postgresql, sets the database user"`
	PersisterPostgresPw      string                `split_words:"true" desc:"If persister type is Postgresql, sets the database password"`

	PubSubProjectID string `split_words:"true" desc:"Sets the Google Cloud Project ID name"`
	PubSubTopicName string `split_words:"true" desc:"Sets the Google Cloud PubSub Topic name"`

	SentryDsn string `split_words:"true" desc:"Sets the Sentry DSN"`
	SentryEnv string `split_words:"true" desc:"Sets the Sentry environment"`
}

// FetchListingAddresses retrieves the list of Civil newsroom listings if given
// the endpoint URL
func (c *CrawlerConfig) FetchListingAddresses() error {
	if c.CivilListingsGraphqlURL == "" {
		return nil
	}

	var listingQuery struct {
		Listings []struct {
			Name            graphql.String
			ContractAddress graphql.String
		} `graphql:"listings(whitelistedOnly: true)"`
	}

	client := graphql.NewClient(c.CivilListingsGraphqlURL, nil)
	err := client.Query(context.Background(), &listingQuery, nil)
	if err != nil {
		return err
	}

	newsroomContractName := "newsroom"
	var addressStrings []string
	var addressObjs []common.Address

	if c.ContractAddresses[newsroomContractName] != "" {
		addressStrings = strings.Split(c.ContractAddresses[newsroomContractName], "|")
	}

	for _, listing := range listingQuery.Listings {
		log.Infof("adding newsroom contract: %v, %v", listing.Name, string(listing.ContractAddress))
		addressStrings = append(addressStrings, string(listing.ContractAddress))
		addressObjs = append(addressObjs, common.HexToAddress(string(listing.ContractAddress)))
	}

	c.ContractAddresses["newsroom"] = strings.Join(addressStrings, "|")
	c.ContractAddressObjs["newsroom"] = append(c.ContractAddressObjs["newsroom"], addressObjs...)
	return nil
}

// OutputUsage prints the usage string to os.Stdout
func (c *CrawlerConfig) OutputUsage() {
	cconfig.OutputUsage(c, envVarPrefix, envVarPrefix)
}

// PopulateFromEnv processes the environment vars, populates CrawlerConfig
// with the respective values, and validates the values.
func (c *CrawlerConfig) PopulateFromEnv() error {
	err := envconfig.Process(envVarPrefix, c)
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
