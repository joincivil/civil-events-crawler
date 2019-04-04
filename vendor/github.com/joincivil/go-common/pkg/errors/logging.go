package errors

import (
	"context"
	"net/http"
	"sync"

	log "github.com/golang/glog"

	"cloud.google.com/go/errorreporting"
	"github.com/getsentry/raven-go"
)

const (
	defaultServiceName    = "civil_default"
	defaultServiceVersion = "0.1"
)

var onceSentry sync.Once

// ErrorMeta is any additional metadata to attach to the error
type ErrorMeta struct {
	UserID         *string
	Tags           map[string]string
	RelatedRequest *http.Request
}

// ErrorReporter is an interface to the error reporting system
// Could by Stackdriver, could be Sentry
type ErrorReporter interface {
	// Msg sends a message asynchronously
	Msg(msg string, meta *ErrorMeta)
	// Error sends an error type asynchronously
	Error(err error, meta *ErrorMeta)
}

// NullErrorReporter is a reporter that does nothing.  Used for testing or
// if an "empty" reporter that logs to nowhere.
type NullErrorReporter struct {
}

// Msg does nothing here
func (n *NullErrorReporter) Msg(msg string, meta *ErrorMeta) {
}

// Error does nothing here
func (n *NullErrorReporter) Error(err error, meta *ErrorMeta) {
}

// MetaErrorReporterConfig configures the meta error reporting
type MetaErrorReporterConfig struct {
	StackDriverProjectID      string
	StackDriverServiceName    string
	StackDriverServiceVersion string

	SentryDSN        string
	SentryRelease    string
	SentryLoggerName string
	SentryDebug      bool
	SentryEnv        string
	SentrySampleRate float32
}

func onStackdriverError(err error) {
	log.Infof("Error stackdriver reporting: err: %v", err)
}

// NewMetaErrorReporter returns a initialized MetaErrorReporter
func NewMetaErrorReporter(config *MetaErrorReporterConfig) (*MetaErrorReporter, error) {
	// Create the Stackdriver client.
	if config.StackDriverProjectID == "" {
		log.Infof("Stackdriver not enabled, projectID is nil")
		return nil, nil
	}
	if config.StackDriverServiceName == "" {
		config.StackDriverServiceName = defaultServiceName
	}
	if config.StackDriverServiceVersion == "" {
		config.StackDriverServiceName = defaultServiceVersion
	}

	ctx := context.Background()
	ec, err := errorreporting.NewClient(ctx, config.StackDriverProjectID, errorreporting.Config{
		ServiceName:    config.StackDriverServiceName,
		ServiceVersion: config.StackDriverServiceVersion,
		OnError:        onStackdriverError,
	})
	if err != nil {
		return nil, err
	}

	if config.SentryDSN != "" {
		// Set up Sentry config only once.
		onceSentry.Do(func() {
			// Startup the Sentry client
			err := raven.SetDSN(config.SentryDSN)
			if err != nil {
				log.Errorf("Error setting Sentry DSN: %v", err)
				return
			}
			raven.SetEnvironment(config.SentryEnv)
			raven.SetRelease(config.SentryRelease)
			err = raven.SetSampleRate(config.SentrySampleRate)
			if err != nil {
				log.Errorf("Error setting Sentry sample rate: %v", err)
				return
			}
			raven.SetDefaultLoggerName(config.SentryLoggerName)
		})
	}

	// Create the Sentry client
	return &MetaErrorReporter{
		stackDriver: ec,
		sentryDSN:   config.SentryDSN,
	}, nil
}

// MetaErrorReporter is an ErrorReporter that reports to multiple services.
// Reports to Stackdriver and Sentry
type MetaErrorReporter struct {
	stackDriver *errorreporting.Client
	sentryDSN   string
}

// Msg sends a message asynchronously
func (m *MetaErrorReporter) Msg(msg string, meta *ErrorMeta) {
	// Stackdriver only logs errors
	// Sentry logs error, tags, req
	if m.sentryDSN != "" {
		if meta != nil {
			raven.CaptureMessage(msg, meta.Tags, raven.NewHttp(meta.RelatedRequest))
		} else {
			raven.CaptureMessage(msg, nil, nil)
		}

	}
}

// Error sends an error type asynchronously
func (m *MetaErrorReporter) Error(err error, meta *ErrorMeta) {
	// Stackdriver logs error, userID, req
	if m.stackDriver != nil {
		logEntry := &errorreporting.Entry{}
		logEntry.Error = err
		if meta != nil {
			if meta.UserID != nil {
				logEntry.User = *meta.UserID
			}
			if meta.RelatedRequest != nil {
				logEntry.Req = meta.RelatedRequest
			}
		}
		m.stackDriver.Report(*logEntry)
	}

	// Sentry logs error, tags, req
	if m.sentryDSN != "" {
		if meta != nil {
			raven.CaptureError(err, meta.Tags, raven.NewHttp(meta.RelatedRequest),
				&raven.User{ID: *meta.UserID})

		} else {
			raven.CaptureError(err, nil, nil)
		}
	}
}
