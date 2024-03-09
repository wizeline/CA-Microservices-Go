package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewZeroLog(t *testing.T) {
	zl := NewZeroLog(DefaultTimeFormat)
	assert.NotNil(t, zl)
}
