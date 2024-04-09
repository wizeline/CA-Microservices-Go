package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewZeroLog(t *testing.T) {
	zl := NewZeroLog()
	assert.NotEqual(t, zl, ZeroLog{})
	zl.Log().Info().Msg("NewZeroLog tested successfully")
}
