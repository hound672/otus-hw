package logger

import (
	"log"
	"log/slog"
	"os"
)

var globalLogger *slog.Logger //nolint: gochecknoglobals // global by design

const defaultLevel = slog.LevelError

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

func InitLogger(config *Config) error {
	levelMapping := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}

	level, ok := levelMapping[config.Level]
	if !ok {
		log.Printf("Can't init logger with log level from config: %s, will be used default level: error", config.Level)
		level = defaultLevel
	}
	logLevel := &slog.LevelVar{}
	logLevel.Set(level)

	globalLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	return nil
}

func Debug(msg string, args ...any) {
	if globalLogger == nil {
		return
	}
	globalLogger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	if globalLogger == nil {
		return
	}
	globalLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	if globalLogger == nil {
		return
	}
	globalLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	if globalLogger == nil {
		return
	}
	globalLogger.Error(msg, args...)
}
