package sentry

import (
	"time"

	"github.com/getsentry/sentry-go"
)

type Sentry struct{}

func (s Sentry) Init() error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://57c983391500f6aeeb84ffd1f77fc84b@o4506716010971136.ingest.sentry.io/4506716011167744",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s Sentry) Flush(seconds int) {
	sentry.Flush(time.Duration(seconds) * time.Second)
}

func (s Sentry) CaptureMessage(message string) {
	sentry.CaptureMessage(message)
}
