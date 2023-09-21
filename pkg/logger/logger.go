package logger

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

func NewLogger() *slog.Logger {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: nil,
		TimeFormat:  time.RFC3339,
		NoColor:     false,
	}))
	return logger
}
