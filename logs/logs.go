//go:build !debug
// +build !debug

package logs

import (
	"flag"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
	With().
	Timestamp().
	Caller().
	Int("pid", os.Getpid()).
	Logger()

func loadLogs() zerolog.Logger {
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if e := log.Debug(); e.Enabled() {
		// Compute log output only if enabled.
		value := "bar"
		e.Str("debugStatus", value).Msg("enabled")
	}
	return logger
}

func Logger() zerolog.Logger {
	return logger
}

func init() {
	loadLogs()
}
