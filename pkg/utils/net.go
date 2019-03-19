package utils

import (
	"net/url"

	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/golang/glog"
	"github.com/joincivil/go-common/pkg/eth"
	"github.com/pkg/errors"
)

// TODO(PN): Move this to go-common

// IsWebsocketURL return true if the given URL is a websocket URL
func IsWebsocketURL(rawurl string) bool {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Infof("Unable to parse URL: err: %v", err)
		return false
	}
	if u.Scheme == "ws" || u.Scheme == "wss" {
		return true
	}
	return false
}

// SetupHTTPEthClient returns an HTTP ethclient if URL is valid
func SetupHTTPEthClient(ethURL string) (*ethclient.Client, error) {
	if IsWebsocketURL(ethURL) {
		return nil, errors.Errorf(
			"fatal: Valid HTTP eth client URL required: configured url: %v",
			ethURL,
		)
	}

	client, err := ethclient.Dial(ethURL)
	if err != nil {
		return nil, errors.Wrap(err, "error dialing http client")
	}

	return client, nil
}

// SetupWebsocketEthClient returns an websocket ethclient if URL is valid.  Sets up a
// "command" ping if killChan is not nil and pingDelay > 0.  If message is sent
// to killChan, stops the ping.
func SetupWebsocketEthClient(ethURL string, killChan <-chan bool, pingDelay int) (*ethclient.Client, error) {
	if ethURL == "" {
		return nil, nil
	}

	if !IsWebsocketURL(ethURL) {
		return nil, errors.Errorf(
			"fatal: Valid websocket eth client URL required: configured url: %v",
			ethURL,
		)
	}

	client, err := ethclient.Dial(ethURL)
	if err != nil {
		return nil, errors.Wrap(err, "error dialing ws client")
	}

	if killChan != nil && pingDelay != 0 {
		go eth.WebsocketPing(client, killChan, pingDelay)
	}

	return client, nil
}
