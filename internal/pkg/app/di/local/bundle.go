package local

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/wjojf/go-uber-fx/internal/api/http"
	fiberFX "github.com/wjojf/go-uber-fx/internal/pkg/app/di/local/fiber"
	"github.com/wjojf/go-uber-fx/internal/pkg/app/di/local/repository"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"github.com/wjojf/go-uber-fx/internal/pkg/logging"
	"github.com/wjojf/go-uber-fx/internal/storage/postgres"
	"go.uber.org/fx"
)

func Bundle(cfg config.Config) fx.Option {

	return fx.Options(

		// Repository
		repository.Module,

		// Usecases / Services

		// Infrastructure
		fx.Supply(cfg),

		logging.Module,

		postgres.Module,

		fx.Supply(fiberFX.ConfigLocal),
		http.Module,

		fx.Invoke(registerHooks),
	)
}

func registerHooks(
	lc fx.Lifecycle,
	cfg config.Config,
	log *slog.Logger,
	http *fiber.App,
) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return onStartHook(ctx, cfg, log, http)
			},
			OnStop: func(ctx context.Context) error {
				return onStopHook(ctx, log, http)
			},
		},
	)
}

func onStartHook(ctx context.Context, cfg config.Config, log *slog.Logger, http *fiber.App) error {
	log.Info("Starting application")
	go http.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
	return nil
}

func onStopHook(ctx context.Context, log *slog.Logger, http *fiber.App) error {
	log.Info("Stopping application")
	return http.Shutdown()
}
