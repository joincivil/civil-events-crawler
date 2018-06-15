package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joincivil/civil-events-crawler/pkg/contractutils"
	"github.com/joincivil/civil-events-crawler/pkg/generated/filterer"
	"github.com/joincivil/civil-events-crawler/pkg/generated/watcher"
	"github.com/joincivil/civil-events-crawler/pkg/listener"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
)

func main() {
	client, err := contractutils.SetupRinkebyClient()
	tcrAddr := "0x77e5aabddb760fba989a1c4b2cdd4aa8fa3d311d"
	startBlock := 0
	if err != nil {
		fmt.Printf("Error connecting to rinkeby: %v", err)
		os.Exit(1)
	}

	filterers := []model.ContractFilterers{
		filterer.NewCivilTCRContractFilterers(common.HexToAddress(tcrAddr)),
	}
	retrieve := retriever.NewCivilEventRetriever(client, startBlock, filterers)
	err = retrieve.Retrieve()
	if err != nil {
		fmt.Printf("Error Retrieving events: %v\n", err)
		os.Exit(1)
	}
	pastEvents := retrieve.PastEvents
	if len(pastEvents) == 0 {
		fmt.Printf("No events collected")
	}

	for _, event := range pastEvents {
		fmt.Println("Event: ", event)
	}

	watchers := []model.ContractWatchers{
		watcher.NewCivilTCRContractWatchers(common.HexToAddress(tcrAddr)),
	}
	listener := listener.NewCivilEventListener(client, watchers)
	if listener == nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	err = listener.Start()
	if err != nil {
		fmt.Printf("Listener should have started with no errors: %v", err)
	}
	defer cleanup(listener)

	quitChan := make(chan interface{})
	eventRecv := make(chan bool)

	go func(quit <-chan interface{}, recv chan<- bool) {
		for {
			select {
			case event := <-listener.EventRecvChan:
				fmt.Println("New event", event)
				recv <- true
			case <-quit:
				return
			}
		}
	}(quitChan, eventRecv)

	numEvents := 0
	for s := range eventRecv {
		fmt.Println("incrementing events", s)
		numEvents++
	}
}

func cleanup(listener *listener.CivilEventListener) {
	err := listener.Stop()
	if err != nil {
		fmt.Println("error stopping listener")
	}
}
