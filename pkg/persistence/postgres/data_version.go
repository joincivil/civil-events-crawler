package postgres // import "github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

import (
	"fmt"
)

const (
	// VersionTableName is the table name for this model
	VersionTableName = "data_version"
)

// CreateVersionTableQuery returns the query to create this table
func CreateVersionTableQuery(tableName string) string {
	queryString := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s(
            version TEXT,
            service_name TEXT,
            last_updated_timestamp INT
        );
    `, tableName)
	return queryString
}

// Version is the model for the version table in DB
type Version struct {
	Version           *string `db:"version"`
	ServiceName       string  `db:"service_name"`
	LastUpdatedDateTs int64   `db:"last_updated_timestamp"`
}
