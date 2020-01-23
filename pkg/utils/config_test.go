// Package time_test contains tests for the config utils
package utils_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/joincivil/civil-events-crawler/pkg/utils"
	cconfig "github.com/joincivil/go-common/pkg/config"
)

// CRAWL_ETH_API_URL=http://ethaddress.com CRAWL_CONTRACT_ADDRESSES=civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55 CRAWL_PERSISTER_TYPE_NAME=postgresql CRAWL_PERSISTER_POSTGRES_ADDRESS=localhost CRAWL_PERSISTER_POSTGRES_PORT=5432 CRAWL_PERSISTER_POSTGRES_DBNAME=civil_crawler go run cmd/crawler/main.go
func setEnvironmentVariables() {
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
}

func TestCrawlerConfig(t *testing.T) {
	setEnvironmentVariables()
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err != nil {
		t.Errorf("Failed to populate from environment: err: %v", err)
	}
}

func TestCrawlerConfigMultiAddresses(t *testing.T) {
	setEnvironmentVariables()
	os.Setenv(
		"CRAWL_CONTRACT_ADDRESSES",
		"civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55|0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d|0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d",
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
	setEnvironmentVariables()
	os.Setenv(
		"CRAWL_ETH_API_URL",
		"ethaddress.com",
	)

	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to allow bad URL from environment: err: %v", err)
	}
}

func TestBadAddressCrawlerConfig(t *testing.T) {
	setEnvironmentVariables()
	// Bad address
	os.Setenv(
		"CRAWL_CONTRACT_ADDRESSES",
		"civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311def,newsroom:0xdfe273082089bb7f70ee36eebcde64832fe97e55",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to allow bad address from environment: err: %v", err)
	}
}

func TestBadPersisterNameCrawlerConfig(t *testing.T) {
	setEnvironmentVariables()
	//Bad persister name
	os.Setenv(
		"CRAWL_PERSISTER_TYPE_NAME",
		"mysql",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to allow bad persister type from environment: err: %v", err)
	}
}

func TestBadPersisterPostgresqlAddressCrawlerConfig(t *testing.T) {
	setEnvironmentVariables()
	//Bad persister postgresql address
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_ADDRESS",
		"",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to allow bad postgres address from environment: err: %v", err)
	}
}

func TestBadPersisterPostgresqlPortCrawlerConfig(t *testing.T) {
	setEnvironmentVariables()
	//Bad persister postgresql port
	os.Setenv(
		"CRAWL_PERSISTER_POSTGRES_PORT",
		"0",
	)
	config := &utils.CrawlerConfig{}
	err := config.PopulateFromEnv()
	if err == nil {
		t.Errorf("Should have failed to allow bad postgres port from environment: err: %v", err)
	}
}

func TestBadPersisterPostgresqlDBNameCrawlerConfig(t *testing.T) {
	setEnvironmentVariables()
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
	_, err := cconfig.PersisterTypeFromName("")
	if err == nil {
		t.Error("Should have failed to retrieve a persister type from empty name")
	}

	persisterType, err := cconfig.PersisterTypeFromName("none")
	if err != nil {
		t.Error("Should have retrieved a persister type for 'none' name")
	}
	if persisterType != cconfig.PersisterTypeNone {
		t.Error("Should have retrieved a persister type for 'none' value")
	}

	persisterType, err = cconfig.PersisterTypeFromName("postgresql")
	if err != nil {
		t.Error("Should have retrieved a persister type for 'postgresql' name")
	}
	if persisterType != cconfig.PersisterTypePostgresql {
		t.Error("Should have retrieved a persister type for 'postgresql' value")
	}
}

func testGraphqlResponse(w http.ResponseWriter, r *http.Request) {
	message := `
	{
		"data": {
			"allListingAddresses": [
				"0xADbB46098E06dBE18aFF2416920FF03EA6814e7b",
				"0x5572DBfa985b1127219ff38f4A10AdB10311725b",
				"0xF71B43B1d4a0462fA9a37F7A3E5f947804A73bfA",
				"0x76a1f346aAA3a1Dc27A5b967b76B096b787055D9",
				"0xA3a7056f4727d9E8094957D937b993adB35f21fF",
				"0xc2A0456154456f0d4e73F6f5acbCd08Ea6A6B2E8",
				"0xcFfd0E01AD3712B776740aa9766034850dbA2725"
				]
		}
	}`
	w.Write([]byte(message)) // nolint: errcheck
}

func testServer(t *testing.T, handler func(http.ResponseWriter, *http.Request)) *http.Server {
	srv := &http.Server{Addr: ":8889"}
	http.HandleFunc("/", handler)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			t.Logf("Error w test service: %s", err)
		}
	}()

	// returning reference so caller can call Shutdown()
	return srv
}

func TestFetchListingAddresses(t *testing.T) {
	server := testServer(t, testGraphqlResponse)
	os.Setenv(
		"CRAWL_ETH_API_URL",
		"http://ethaddress.com",
	)
	os.Setenv(
		"CRAWL_CIVIL_LISTING_GRAPHQL_URL",
		"http://localhost:8889/query",
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
		t.Errorf("Should have populated env vars: err: %v", err)
	}
	if len(config.ContractAddresses["newsroom"]) <= 0 {
		t.Errorf("Should have fetched listing addresses: len: %v, err: %v",
			len(config.ContractAddresses["newsroom"]), err)
	}
	// 1 newsroom from config, 7 from graphql
	if len(config.ContractAddressObjs["newsroom"]) != 8 {
		t.Errorf("Should have fetched listing 8 address objs: len: %v, err: %v",
			len(config.ContractAddressObjs["newsroom"]), err)
	}
	server.Shutdown(context.TODO()) // nolint: errcheck
}

func testGraphqlMultiSigResponse(w http.ResponseWriter, r *http.Request) {
	message := `
	{
		"data": {
			"allMultiSigAddresses": [
				"0xADbB46098E06dBE18aFF2416920FF03EA6814e7b",
				"0x5572DBfa985b1127219ff38f4A10AdB10311725b",
				"0xF71B43B1d4a0462fA9a37F7A3E5f947804A73bfA",
				"0x76a1f346aAA3a1Dc27A5b967b76B096b787055D9",
				"0xA3a7056f4727d9E8094957D937b993adB35f21fF",
				"0xc2A0456154456f0d4e73F6f5acbCd08Ea6A6B2E8",
				"0xcFfd0E01AD3712B776740aa9766034850dbA2725"
				]
		}
	}`
	w.Write([]byte(message)) // nolint: errcheck
}

func TestFetchMultiSigAddresses(t *testing.T) {
	server := testServer(t, testGraphqlMultiSigResponse)
	os.Setenv(
		"CRAWL_ETH_API_URL",
		"http://ethaddress.com",
	)
	os.Setenv(
		"CRAWL_CIVIL_LISTING_GRAPHQL_URL",
		"http://localhost:8889/query",
	)
	os.Setenv(
		"CRAWL_CONTRACT_ADDRESSES",
		"civiltcr:0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d",
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
		t.Errorf("Should have populated env vars: err: %v", err)
	}
	if len(config.ContractAddresses["multisigwallet"]) <= 0 {
		t.Errorf("Should have fetched listing addresses: len: %v, err: %v",
			len(config.ContractAddresses["multisigwallet"]), err)
	}
	// 0 multi sigs from config, 7 from graphql
	if len(config.ContractAddressObjs["multisigwallet"]) != 7 {
		t.Errorf("Should have fetched listing 7 address objs: len: %v, err: %v",
			len(config.ContractAddressObjs["multisigwallet"]), err)
	}
	server.Shutdown(context.TODO()) // nolint: errcheck
}
