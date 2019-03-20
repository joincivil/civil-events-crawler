// Package main contains commands to run
package main

import (
	"flag"
	"os"
	"time"

	log "github.com/golang/glog"

	"github.com/joincivil/civil-events-crawler/pkg/crawlermain"
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
		log.Errorf("Invalid crawler config: err: %v\n", err)
		os.Exit(2)
	}

	errRep, err := crawlermain.InitErrorReporter(config)
	if err != nil {
		log.Errorf("Error init error reporting: err: %+v\n", err)
		os.Exit(2)
	}

	err = crawlermain.StartUpCrawler(config, errRep)
	if err != nil {
		log.Errorf("Crawler error: err: %+v\n", err)
		errRep.Error(err, nil)
		// XXX(PN): Ensure the error gets sent before we die
		time.Sleep(3 * time.Second)
		os.Exit(2)
	}
	log.Info("Crawler stopped")
}
