package v1

import (
	stdlog "log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	zap    *zap.Logger
	level  Level
	format Format
}

func NewZapFromEnv() (Logger, func()) {
	format, level := configFromEnv()
	return NewZap(format, level, zapcore.Lock(os.Stdout))
}

func NewZap(format Format, level Level, writer zapcore.WriteSyncer) (Logger, func()) {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	if format == FormatConsole {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	var coreLevel zap.AtomicLevel
	switch level {
	case LevelDebug:
		coreLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case LevelInfo:
		coreLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case LevelError:
		coreLevel = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		coreLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	logger := zap.New(zapcore.NewCore(encoder, writer, coreLevel))
	zap.ReplaceGlobals(logger)
	_ = zap.RedirectStdLog(logger)

	return &ZapLogger{
			zap:    logger,
			level:  level,
			format: format,
		}, func() {
			_ = logger.Sync()
		}
}

func (l *ZapLogger) GetStd() *stdlog.Logger {
	return zap.NewStdLog(l.zap)
}

func (l *ZapLogger) GetLevel() Level {
	return l.level
}

func (l *ZapLogger) GetFormat() Format {
	return l.format
}

func (l *ZapLogger) Log(level Level, msg string, args map[string]any) {
	fields := make([]zap.Field, 0, len(args))
	for k, v := range args {
		fields = append(fields, zap.Any(k, v))
	}
	switch level {
	case LevelDebug:
		l.zap.Debug(msg, fields...)
	case LevelInfo:
		l.zap.Info(msg, fields...)
	case LevelError:
		l.zap.Error(msg, fields...)
	case LevelFatal:
		l.zap.Fatal(msg, fields...)
	default:
		l.zap.Error(msg, fields...)
	}
}

func (l *ZapLogger) LogMsg(level Level, msg string) {
	switch level {
	case LevelDebug:
		l.zap.Debug(msg)
	case LevelInfo:
		l.zap.Info(msg)
	case LevelError:
		l.zap.Error(msg)
	case LevelFatal:
		l.zap.Fatal(msg)
	default:
		l.zap.Error(msg)
	}
}

func (l *ZapLogger) Debug(msg string, args map[string]any) {
	fields := make([]zap.Field, 0, len(args))
	for k, v := range args {
		fields = append(fields, zap.Any(k, v))
	}
	l.zap.Debug(msg, fields...)
}

func (l *ZapLogger) Info(msg string, args map[string]any) {
	fields := make([]zap.Field, 0, len(args))
	for k, v := range args {
		fields = append(fields, zap.Any(k, v))
	}
	l.zap.Info(msg, fields...)
}

func (l *ZapLogger) Error(msg string, args map[string]any) {
	fields := make([]zap.Field, 0, len(args))
	for k, v := range args {
		fields = append(fields, zap.Any(k, v))
	}
	l.zap.Error(msg, fields...)
}

func (l *ZapLogger) Fatal(msg string, args map[string]any) {
	fields := make([]zap.Field, 0, len(args))
	for k, v := range args {
		fields = append(fields, zap.Any(k, v))
	}
	l.zap.Fatal(msg, fields...)
}
