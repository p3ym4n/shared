package log

import (
	stdlog "log"
)

//go:generate mockery --case underscore --dir ./ --name Logger --output ./logtest --outpkg logtest
type Logger interface {
	// GetStd return the inner standard logger
	GetStd() *stdlog.Logger
	GetLevel() Level
	GetFormat() Format

	// Log LogMsg for logging with dynamic level with msg and details
	Log(level Level, msg string, args map[string]any)
	// Msg for logging with dynamic level and only msg
	Msg(level Level, msg string)

	Debug(msg string, args map[string]any)
	Info(msg string, args map[string]any)
	Error(msg string, args map[string]any)
	Fatal(msg string, args map[string]any)
}
