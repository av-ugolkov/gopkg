package logger

import (
	"fmt"
	"log"
	"log/slog"
)

func Fatalf(msg string, args ...any) {
	log.Fatalf(msg, args...)
}

func Infof(msg string, args ...any) {
	slog.Info(fmt.Sprintf(msg, args...))
}

func Debugf(msg string, args ...any) {
	slog.Debug(fmt.Sprintf(msg, args...))
}

func Warnf(msg string, args ...any) {
	slog.Warn(fmt.Sprintf(msg, args...))
}

func Errorf(msg string, args ...any) {
	slog.Error(fmt.Sprintf(msg, args...))
}
