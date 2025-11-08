package logging

import (
	"log/slog"
	"os"
	"withoutforget/cider/internal/config"

	"github.com/lmittmann/tint"
)

func InitLogger(cfg *config.Logging) {
	var logger *slog.Logger

	var level slog.Level

	switch cfg.Level {
	case "info":
		level = slog.LevelInfo
	case "debug":
		level = slog.LevelDebug
	}

	if cfg.HumanReadable {
		logger = slog.New(tint.NewHandler(os.Stderr, &tint.Options{
			Level:     level,
			AddSource: true,
		}))
	} else {
		logger = slog.New(
			slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
				AddSource: true,
				Level:     level,
			}))
	}

	slog.SetDefault(logger)
}
