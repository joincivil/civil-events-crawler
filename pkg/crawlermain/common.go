package crawlermain

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	elog "github.com/ethereum/go-ethereum/log"
	log "github.com/golang/glog"

	"github.com/joincivil/civil-events-crawler/pkg/eventcollector"
	"github.com/joincivil/civil-events-crawler/pkg/generated/handlerlist"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

	cerrors "github.com/joincivil/go-common/pkg/errors"
)

// InitErrorReporter initializes the error reporter
func InitErrorReporter(config *utils.CrawlerConfig) (cerrors.ErrorReporter, error) {
	errRepConfig := &cerrors.MetaErrorReporterConfig{
		StackDriverProjectID:      "civil-media",
		StackDriverServiceName:    "crawler",
		StackDriverServiceVersion: "1.0",
		SentryDSN:                 config.SentryDsn,
		SentryDebug:               false,
		SentryEnv:                 config.SentryEnv,
		SentryLoggerName:          "crawler_logger",
		SentryRelease:             "1.0",
		SentrySampleRate:          1.0,
	}
	reporter, err := cerrors.NewMetaErrorReporter(errRepConfig)
	if err != nil {
		log.Errorf("Error creating meta reporter: %v", err)
		return nil, err
	}
	if reporter == nil {
		log.Infof("Enabling null error reporter")
		return &cerrors.NullErrorReporter{}, nil
	}
	return reporter, nil
}

func contractFilterers(contractAddressObjs map[string][]common.Address) []model.ContractFilterers {
	return handlerlist.ContractFilterers(contractAddressObjs)
}

func contractWatchers(contractAddressObjs map[string][]common.Address) []model.ContractWatchers {
	return handlerlist.ContractWatchers(contractAddressObjs)
}

func eventTriggers(config *utils.CrawlerConfig) []eventcollector.Trigger {
	return []eventcollector.Trigger{
		&eventcollector.AddNewsroomWatchersTrigger{},
		&eventcollector.RemoveNewsroomWatchersTrigger{},
	}
}

func cleanup(eventCol *eventcollector.EventCollector, killChan chan<- bool) {
	log.Info("Stopping crawler...")
	err := eventCol.StopCollection(false)
	if err != nil {
		log.Errorf("Error stopping collection: err: %v", err)
	}
	if killChan != nil {
		close(killChan)
	}
	log.Info("Crawler stopped")
}

func setupKillNotify(eventCol *eventcollector.EventCollector, killChan chan<- bool) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup(eventCol, killChan)
		os.Exit(1)
	}()
}

func enableGoEtherumLogging() {
	glog := elog.NewGlogHandler(elog.StreamHandler(os.Stderr, elog.TerminalFormat(false)))
	glog.Verbosity(elog.Lvl(elog.LvlDebug)) // nolint: unconvert
	glog.Vmodule("")                        // nolint: errcheck, gosec
	elog.Root().SetHandler(glog)
}
