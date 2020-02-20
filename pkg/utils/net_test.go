package utils_test

import (
	"testing"

	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

func TestIsWebsocketURL(t *testing.T) {
	if utils.IsWebsocketURL("http://mydomain.com") {
		t.Errorf("Should have returned false for http url")
	}
	if utils.IsWebsocketURL("https://mydomain.com") {
		t.Errorf("Should have returned false for https url")
	}
	if !utils.IsWebsocketURL("ws://mydomain.com/ws") {
		t.Errorf("Should have returned true for ws url")
	}
	if !utils.IsWebsocketURL("wss://mydomain.com/ws") {
		t.Errorf("Should have returned true for wss url")
	}
}

func TestSetupHTTPEthClient(t *testing.T) {
	client, err := utils.SetupHTTPEthClient("https://rinkeby.infura.io/")
	if err != nil {
		t.Errorf("Should not have gotten error: err: %v", err)
	}
	if client == nil {
		t.Errorf("Should not have nil client")
	}

	_, err = utils.SetupHTTPEthClient("wss://rinkeby.infura.io/ws")
	if err == nil {
		t.Errorf("Should have gotten error")
	}
}
