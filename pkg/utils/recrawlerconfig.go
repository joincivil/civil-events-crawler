package utils

import (
	"github.com/ethereum/go-ethereum/common"
	log "github.com/golang/glog"
	cconfig "github.com/joincivil/go-common/pkg/config"
	"github.com/kelseyhightower/envconfig"
)

const (
	envVarPrefixRecrawler = "recrawl"
)

// NOTE(PN): After envconfig populates RecrawlerConfig with the environment vars,
// there is nothing preventing the RecrawlerConfig fields from being mutated.

// RecrawlerConfig is the master config for the recrawler derived from environment
// variables.
type RecrawlerConfig struct {
	EthAPIURL     string `envconfig:"eth_api_url" required:"true" desc:"Ethereum HTTP API address"`
	EthStartBlock uint64 `envconfig:"eth_start_block" desc:"Sets the start Eth block (default 0)" default:"0"`

	WetRun bool `envconfig:"wet_run" desc:"If set to true, will update the events table with missing events"`

	// These are the contracts you want to try and recrawl.
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
}

// OutputUsage prints the usage string to os.Stdout
func (c *RecrawlerConfig) OutputUsage() {
	cconfig.OutputUsage(c, envVarPrefixRecrawler, envVarPrefixRecrawler)
}

// PopulateFromEnv processes the environment vars, populates CrawlerConfig
// with the respective values, and validates the values.
func (c *RecrawlerConfig) PopulateFromEnv() error {
	err := envconfig.Process(envVarPrefixRecrawler, c)
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

	err = c.populatePersisterType()
	if err != nil {
		return err
	}

	return c.validatePersister()
}

func (c *RecrawlerConfig) populatePersisterType() error {
	var err error
	c.PersisterType, err = cconfig.PersisterTypeFromName(c.PersisterTypeName)
	return err
}

func (c *RecrawlerConfig) populateContractAddressObjs() {
	var err error
	c.ContractAddressObjs, err = populateContractAddressObjs(c.ContractAddresses)
	if err != nil {
		log.Errorf("Error populating contract address objs: %v", err)
	}
}

func (c *RecrawlerConfig) validateContractAddresses() error {
	return validateContractAddresses(c.ContractAddresses)
}

func (c *RecrawlerConfig) validateAPIURL() error {
	return validateAPIURL(c.EthAPIURL)
}

func (c *RecrawlerConfig) validatePersister() error {
	var err error
	if c.PersisterType == cconfig.PersisterTypePostgresql {
		err = validatePostgresqlPersister(
			c.PersisterPostgresAddress,
			c.PersisterPostgresPort,
			c.PersisterPostgresDbname,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// Address returns the address of the persister
// Implements PersisterConfig
func (c *RecrawlerConfig) Address() string {
	return c.PersisterPostgresAddress
}

// Port returns the port of the persister
// Implements PersisterConfig
func (c *RecrawlerConfig) Port() int {
	return c.PersisterPostgresPort
}

// Type returns the persister type
// Implements PersisterConfig
func (c *RecrawlerConfig) Type() cconfig.PersisterType {
	return c.PersisterType
}

// Username returns the username to access the persister
// Implements PersisterConfig
func (c *RecrawlerConfig) Username() string {
	return c.PersisterPostgresUser
}

// Password returns the password to access the persister
// Implements PersisterConfig
func (c *RecrawlerConfig) Password() string {
	return c.PersisterPostgresPw
}

// Dbname returns the "dbname" to access the persister
// Implements PersisterConfig
func (c *RecrawlerConfig) Dbname() string {
	return c.PersisterPostgresDbname
}
