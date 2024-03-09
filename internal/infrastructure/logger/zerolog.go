package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// Note: add time format types supported by zerolog as needed
const (
	DefaultTimeFormat TimeFormat = "2006/01/02 15:04:05"
)

type ZeroLog struct {
	logger *zerolog.Logger
}

type TimeFormat string

func (tf TimeFormat) String() string {
	return string(tf)
}

func NewZeroLog(timeFormat TimeFormat) *ZeroLog {
	writer := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stderr
		w.TimeFormat = timeFormat.String()
	})

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zl := zerolog.New(writer).With().Timestamp().Logger()

	return &ZeroLog{
		logger: &zl,
	}
}
