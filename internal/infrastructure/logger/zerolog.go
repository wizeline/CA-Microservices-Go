package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// defaultTimeFormat is the default time format used by the logger.
// Note: here you might change the time format as needed.
const defaultTimeFormat = "2006/01/02 15:04:05"

// ZeroLog is a logger type with zerolog.Logger support
type ZeroLog struct {
	logger *zerolog.Logger
}

// NewZeroLog returns an implemented ZeroLog instance.
func NewZeroLog() ZeroLog {
	writer := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stderr
		w.TimeFormat = defaultTimeFormat
	})

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zl := zerolog.New(writer).With().Timestamp().Logger()

	return ZeroLog{
		logger: &zl,
	}
}

// Log returns the underlying zerolog.Logger instance
func (l ZeroLog) Log() *zerolog.Logger {
	return l.logger
}
