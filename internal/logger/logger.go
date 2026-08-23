package logger

import "log/slog"

func Info(msg string) {
	slog.Info(msg)
}

func Error(msg string) {
	slog.Error(msg)
}

func Fatal(msg string) {
	slog.Error("fatal: " + msg)
}
