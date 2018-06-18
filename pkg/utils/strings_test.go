// Package time_test contains tests for the string utils
package utils_test

import (
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"testing"
)

func TestIsValidEthAPIURL(t *testing.T) {
	if utils.IsValidEthAPIURL("thisisnotavalidurl") {
		t.Error("Should have failed on an invalid eth API url")
	}
	if utils.IsValidEthAPIURL("http//thisisnotavalidurl.com") {
		t.Error("Should have failed on an invalid eth API url")
	}
	if utils.IsValidEthAPIURL("http/thisisnotavalidurl.com") {
		t.Error("Should have failed on an invalid eth API url")
	}
	if !utils.IsValidEthAPIURL("http://thisisvalid.co") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !utils.IsValidEthAPIURL("http://thisisvalid.com") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !utils.IsValidEthAPIURL("https://thisisvalid.com") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !utils.IsValidEthAPIURL("https://thisisvalid.longtld") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !utils.IsValidEthAPIURL("ws://thisisvalid.ether/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !utils.IsValidEthAPIURL("wss://thisisvalid.com/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !utils.IsValidEthAPIURL("wss://localhost/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !utils.IsValidEthAPIURL("wss://localhost:8545/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !utils.IsValidEthAPIURL("wss://127.0.0.1/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !utils.IsValidEthAPIURL("wss://127.0.0.1:8545/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
}

func TestIsValidContractAddress(t *testing.T) {
	if utils.IsValidContractAddress("") {
		t.Error("Should have failed on an empty contract address")
	}
	if utils.IsValidContractAddress("thisisnotavalidaddress") {
		t.Error("Should have failed on an invalid contract address")
	}
	if utils.IsValidContractAddress("0xdfe273082089bb7f70ee36eebcde64832fe97e55f") {
		t.Error("Should have failed on an invalid contract address")
	}
	if !utils.IsValidContractAddress("0xdfe273082089bb7f70ee36eebcde64832fe97e55") {
		t.Error("Should have not have failed on an valid contract address")
	}

}
