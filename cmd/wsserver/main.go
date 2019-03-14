package main

// This is for test purposes only.
// This starts up a simple websocket service to test abruptly closing connections on
// the crawler.

import (
	"fmt"
	"net/http"
	"time"

	cstrings "github.com/joincivil/go-common/pkg/strings"
	ctime "github.com/joincivil/go-common/pkg/time"
	"golang.org/x/net/websocket"
)

var (
	startTime = 0

	timeToCloseConnSecs = 90
)

type ethSubscribe struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type ethSubscribeResp struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

func (e *ethSubscribeResp) genResult() {
	randHex, err := cstrings.RandomHexStr(32)
	if err != nil {
		fmt.Printf("Error getting random hex str: err: %v", err)
		e.Result = ""
		return
	}
	e.Result = randHex
}

func waitCloseConn(ws *websocket.Conn) {
	time.Sleep(time.Duration(timeToCloseConnSecs) * time.Second)
	err := ws.Close()
	if err != nil {
		fmt.Printf("Error closing conn: err: %v\n", err)
	}
}

func testWsCloseServer(ws *websocket.Conn) {
	go waitCloseConn(ws)
	for {

		var msg ethSubscribe
		fmt.Printf("waiting to read\n")
		if err := websocket.JSON.Receive(ws, &msg); err != nil {
			fmt.Printf("Error reading : %v\n", err)
			return
		}

		respMsg := &ethSubscribeResp{
			JSONRPC: "2.0",
			ID:      msg.ID,
		}

		respMsg.genResult()
		if err := websocket.JSON.Send(ws, respMsg); err != nil {
			fmt.Printf("Error writing resp : %v\n", err)
			return
		}

		fmt.Printf("done responding\n")
	}
}

func main() {
	startTime = ctime.CurrentEpochSecsInInt()
	http.Handle("/ws", websocket.Handler(testWsCloseServer))
	fmt.Printf("Starting up on port :9090\n")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic("Error starting testWsCloseServer: err: %v" + err.Error())
	}
}
