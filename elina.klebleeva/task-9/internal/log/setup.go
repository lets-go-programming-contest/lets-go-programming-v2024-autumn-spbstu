package log

import (
	"log/slog"
	"os"

	"github.com/EmptyInsid/task-9/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func Setup(env config.LoggerConfig) *slog.Logger {
	var log *slog.Logger

	switch env.Mod {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
