package testutils

import "os"

type DBCreds struct {
	Port     int
	Dbname   string
	User     string
	Password string
	Host     string
}

// GetTestDBCreds returns the credentials for the local docker instance
// dependent on env vars.
func GetTestDBCreds() DBCreds {
	var creds DBCreds
	if os.Getenv("CI") == "true" {
		creds = DBCreds{
			Port:     5432,
			Dbname:   "circle_test",
			User:     "root",
			Password: "root",
			Host:     "localhost",
		}
	} else {
		creds = DBCreds{
			Port:     5432,
			Dbname:   "civil_crawler",
			User:     "docker",
			Password: "docker",
			Host:     "localhost",
		}
	}
	return creds
}
