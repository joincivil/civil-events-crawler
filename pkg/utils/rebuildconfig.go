// Package utils contains various common utils separate by utility types
package utils

import (
	"errors"

	cconfig "github.com/joincivil/go-common/pkg/config"

	"github.com/kelseyhightower/envconfig"
)

const (
	rebuildEnvVarPrefix = "rebuild"
)

// RebuildConfig is the master config for the rebuild script derived from environment
// variables.
type RebuildConfig struct {
	PersisterType            cconfig.PersisterType `ignored:"true"`
	PersisterPostgresAddress string                `split_words:"true" required:"true" desc:"If persister type is Postgresql, sets the address"`
	PersisterPostgresPort    int                   `split_words:"true" required:"true" desc:"If persister type is Postgresql, sets the port"`
	PersisterPostgresDbname  string                `split_words:"true" required:"true" desc:"If persister type is Postgresql, sets the database name"`
	PersisterPostgresUser    string                `split_words:"true" required:"true" desc:"If persister type is Postgresql, sets the database user"`
	PersisterPostgresPw      string                `split_words:"true" required:"true" desc:"If persister type is Postgresql, sets the database password"`
}

// PopulateFromEnv processes the environment vars, populates RebuildConfig
// with the respective values, and validates the values.
func (r *RebuildConfig) PopulateFromEnv() error {
	err := envconfig.Process(rebuildEnvVarPrefix, r)
	if err != nil {
		return err
	}
	r.populatePersisterType()

	return r.validatePersister()
}

func (r *RebuildConfig) populatePersisterType() {
	r.PersisterType = cconfig.PersisterTypePostgresql
}

func (r *RebuildConfig) validatePersister() error {
	return r.validatePostgresqlPersister()
}

// OutputUsage prints the usage string to os.Stdout
func (r *RebuildConfig) OutputUsage() {
	cconfig.OutputUsage(r, rebuildEnvVarPrefix, rebuildEnvVarPrefix)
}

func (r *RebuildConfig) validatePostgresqlPersister() error {
	if r.PersisterPostgresAddress == "" {
		return errors.New("Postgresql address required")
	}
	if r.PersisterPostgresPort == 0 {
		return errors.New("Postgresql port required")
	}
	if r.PersisterPostgresDbname == "" {
		return errors.New("Postgresql db name required")
	}
	return nil
}
