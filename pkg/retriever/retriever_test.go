// Package retriever_test contains tests for retriever package
package retriever_test

import (
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/golang/glog"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
	"testing"
)

// Using rinkeby for now but will change to simulated backend
const (
	rinkebyAddress = "https://rinkeby.infura.io"
	testTCRAddress = "0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d"
	startBlock     = 2335623
)

func setupRinkebyClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(rinkebyAddress)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}
	return client, nil
}

// TestEventCollection tests that events are being collected,
func TestEventCollection(t *testing.T) {
	client, err := setupRinkebyClient()
	if err != nil {
		t.Errorf("Error connecting to rinkeby: %v", err)
	}
	retrieve := retriever.NewCivilEventRetriever(client, testTCRAddress, startBlock)
	err = retrieve.Retrieve()
	if err != nil {
		t.Errorf("Error retrieving events: %v", err)
	}
	pastEvents := retrieve.PastEvents
	if len(pastEvents) == 0 {
		t.Error("No events collected")
	}
}
