package postgres // import "github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

import (
    "fmt"
    // "github.com/ethereum/go-ethereum/common"
)

const (
    versionTableName = "version"
)

// CreateVersionTableQuery returns the query to create the version table
func CreateVersionTableQuery() string {
    return CreateVersionTableQueryString(versionTableName)
}

// CreateVersionTableQueryString returns the query to create this table
func CreateVersionTableQueryString(tableName string) string {
    queryString := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s(
            version TEXT,
            service_name TEXT
        );
    `, versionTableName)
    return queryString
}

// VersionData is the model for the version table in DB
type VersionData struct {
    Version     string `db:"version"`
    ServiceName string `db:"service_name"`
}
