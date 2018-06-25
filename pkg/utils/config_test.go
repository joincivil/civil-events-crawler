// Package time_test contains tests for the config utils
package utils_test

import (
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"os"
	"testing"
)

// CRAWL_ETH_API_URL=http://ethaddress.com CRAWL_CONTRACT_ADDRESSES=civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55 CRAWL_PERSISTER_TYPE_NAME=postgresql CRAWL_PERSISTER_POSTGRES_ADDRESS=localhost CRAWL_PERSISTER_POSTGRES_PORT=5432 CRAWL_PERSISTER_POSTGRES_DBNAME=civil_crawler go run cmd/crawler/main.go

func TestCrawlerConfig(t *testing.T) {
	os.Setenv(
		"CRAWL_ETH_API_URL",
		"http://ethaddress.com",
	)
	os.Setenv(
		"CRAWL_CONTRACT_ADDRESSES",
		"civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55",
	)
	os.Setenv(
		"CRAWL_PERSISTER_TYPE_NAME",
		"postgresql",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_ADDRESS",
		"localhost",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_PORT",
		"5432",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_DBNAME",
		"civil_crawler",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err != nil {
		t.Errorf("Failed to populate from environment: err: %v", err)
	}
}

func TestCrawlerConfigMultiAddresses(t *testing.T) {
	os.Setenv(
		"CRAWL_ETH_API_URL",
		"http://ethaddress.com",
	)
	os.Setenv(
		"CRAWL_CONTRACT_ADDRESSES",
		"civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55|0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d|0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d",
	)
	os.Setenv(
		"CRAWL_PERSISTER_TYPE_NAME",
		"postgresql",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_ADDRESS",
		"localhost",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_PORT",
		"5432",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_DBNAME",
		"civil_crawler",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err != nil {
		t.Errorf("Failed to populate from environment: err: %v", err)
	}
	objs, ok := config.ContractAddressObjs["newsroom"]
	if !ok {
		t.Error("Should have seen the newsroom addresses")
	}
	if len(objs) != 3 {
		t.Error("Should have seen the 3 newsroom addresses")
	}
}

func TestBadEthURLCrawlerConfig(t *testing.T) {
	// Bad URL
	os.Setenv(
		"CRAWL_ETH_API_URL",
		"ethaddress.com",
	)
	os.Setenv(
		"CRAWL_CONTRACT_ADDRESSES",
		"civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55",
	)
	os.Setenv(
		"CRAWL_PERSISTER_TYPE_NAME",
		"postgresql",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_ADDRESS",
		"localhost",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_PORT",
		"5432",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_DBNAME",
		"civil_crawler",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to allow bad URL from environment: err: %v", err)
	}
}

func TestBadAddressCrawlerConfig(t *testing.T) {
	os.Setenv(
		"CRAWL_ETH_API_URL",
		"http://ethaddress.com",
	)
	// Bad address
	os.Setenv(
		"CRAWL_CONTRACT_ADDRESSES",
		"civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311def,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55",
	)
	os.Setenv(
		"CRAWL_PERSISTER_TYPE_NAME",
		"postgresql",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_ADDRESS",
		"localhost",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_PORT",
		"5432",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_DBNAME",
		"civil_crawler",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to allow bad address from environment: err: %v", err)
	}
}

func TestBadPersisterNameCrawlerConfig(t *testing.T) {
	os.Setenv(
		"CRAWL_ETH_API_URL",
		"http://ethaddress.com",
	)
	os.Setenv(
		"CRAWL_CONTRACT_ADDRESSES",
		"civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55",
	)
	//Bad persister name
	os.Setenv(
		"CRAWL_PERSISTER_TYPE_NAME",
		"mysql",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_ADDRESS",
		"localhost",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_PORT",
		"5432",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_DBNAME",
		"civil_crawler",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to allow bad persister type from environment: err: %v", err)
	}
}

func TestBadPersisterPostgresqlAddressCrawlerConfig(t *testing.T) {
	os.Setenv(
		"CRAWL_ETH_API_URL",
		"http://ethaddress.com",
	)
	os.Setenv(
		"CRAWL_CONTRACT_ADDRESSES",
		"civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55",
	)
	os.Setenv(
		"CRAWL_PERSISTER_TYPE_NAME",
		"postgresql",
	)
	//Bad persister postgresql address
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_ADDRESS",
		"",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_PORT",
		"5432",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_DBNAME",
		"civil_crawler",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to allow bad postgres address from environment: err: %v", err)
	}
}

func TestBadPersisterPostgresqlPortCrawlerConfig(t *testing.T) {
	os.Setenv(
		"CRAWL_ETH_API_URL",
		"http://ethaddress.com",
	)
	os.Setenv(
		"CRAWL_CONTRACT_ADDRESSES",
		"civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55",
	)
	os.Setenv(
		"CRAWL_PERSISTER_TYPE_NAME",
		"postgresql",
	)
	//Bad persister postgresql address
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_ADDRESS",
		"localhost",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_PORT",
		"0",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_DBNAME",
		"civil_crawler",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to allow bad postgres port from environment: err: %v", err)
	}
}

func TestBadPersisterPostgresqlDBNameCrawlerConfig(t *testing.T) {
	os.Setenv(
		"CRAWL_ETH_API_URL",
		"http://ethaddress.com",
	)
	os.Setenv(
		"CRAWL_CONTRACT_ADDRESSES",
		"civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55",
	)
	os.Setenv(
		"CRAWL_PERSISTER_TYPE_NAME",
		"postgresql",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_ADDRESS",
		"localhost",
	)
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_PORT",
		"5432",
	)
	//Bad persister dbname
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_DBNAME",
		"",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to allow bad postgres dbname from environment: err: %v", err)
	}
}

func TestEmptyCrawlerConfig(t *testing.T) {
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to populate from empty environment: err: %v", err)
	}
}

func TestCrawlerConfigUsage(t *testing.T) {
	config := &utils.CrawlerConfig{}
	config.OutputUsage()
}

func TestPersisterTypeFromName(t *testing.T) {
	_, err := utils.PersisterTypeFromName("")
	if err == nil {
		t.Error("Should have failed to retrieve a persister type from empty name")
	}

	persisterType, err := utils.PersisterTypeFromName("none")
	if err != nil {
		t.Error("Should have retrieved a persister type for 'none' name")
	}
	if persisterType != utils.PersisterTypeNone {
		t.Error("Should have retrieved a persister type for 'none' value")
	}

	persisterType, err = utils.PersisterTypeFromName("postgresql")
	if err != nil {
		t.Error("Should have retrieved a persister type for 'postgresql' name")
	}
	if persisterType != utils.PersisterTypePostgresql {
		t.Error("Should have retrieved a persister type for 'postgresql' value")
	}
}
