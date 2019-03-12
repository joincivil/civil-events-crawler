package postgres // import "github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

import (
	"fmt"
)

const (
	// VersionTableName is the table name for this model
	VersionTableName = "data_version"
	// VersionFieldName is the name for the version field in Version struct
	VersionFieldName = "version"
	// ServiceFieldName is the name for the service_name field in Version struct
	ServiceFieldName = "service_name"
	// LastUpdatedTsFieldName is the name for last_updated_timestamp field in Version struct
	LastUpdatedTsFieldName = "last_updated_timestamp"
	// ExistsFieldName is the name for exists field in Version struct
	ExistsFieldName = "exists"
)

// CreateVersionTableQuery returns the query to create this table
// version and service_name are unique.
func CreateVersionTableQuery(tableName string) string {
	queryString := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s(
            version TEXT,
            service_name TEXT,
            last_updated_timestamp INT,
            exists BOOL,
            PRIMARY KEY(version, service_name)
        );
    `, tableName)
	return queryString
}

// Version is the model for the version table in DB
type Version struct {
	Version           *string `db:"version"`
	ServiceName       string  `db:"service_name"`
	LastUpdatedDateTs int64   `db:"last_updated_timestamp"`
	Exists            bool    `db:"exists"`
}
