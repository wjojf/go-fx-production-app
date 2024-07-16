package logging

import (
	"context"
	"log/slog"
	"os"

	"github.com/wjojf/go-uber-fx/internal/pkg/config"
)

func SetupLogger(cfg config.Config) *slog.Logger {

	var log *slog.Logger

	switch cfg.Environment {

	case config.EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case config.EnvDev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func MustGetLogger(ctx context.Context) *slog.Logger {
	log, _ := ctx.Value("log").(*slog.Logger)
	return log
}
