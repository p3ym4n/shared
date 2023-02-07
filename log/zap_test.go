package v1

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestNewZapFromEnv(t *testing.T) {
	logger, closer := NewZapFromEnv()
	assert.IsType(t, func() {}, closer)
	assert.NotNil(t, logger)
}

func TestNewZap(t *testing.T) {
	format := FormatConsole
	level := LevelError
	logger, closer := NewZap(format, level, zapcore.Lock(os.Stdout))
	assert.IsType(t, func() {}, closer)
	assert.NotNil(t, logger)
}
