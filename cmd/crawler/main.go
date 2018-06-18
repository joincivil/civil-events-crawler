// Package main contains commands to run
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

func main() {
	config := &utils.CrawlerConfig{}
	flag.Usage = func() {
		config.OutputUsage()
		os.Exit(0)
	}
	flag.Parse()

	err := config.PopulateFromEnv()
	if err != nil {
		config.OutputUsage()
		fmt.Printf("Invalid crawler config: err: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("config = %v", config)
}
