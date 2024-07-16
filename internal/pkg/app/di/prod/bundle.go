package prod

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	fiberFX "github.com/wjojf/go-uber-fx/internal/pkg/app/di/prod/fiber"
	"github.com/wjojf/go-uber-fx/internal/pkg/app/di/prod/repository"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"github.com/wjojf/go-uber-fx/internal/pkg/logging"
	"github.com/wjojf/go-uber-fx/internal/storage/postgres"
	"go.uber.org/fx"
)

func Bundle(cfg config.Config) fx.Option {

	return fx.Options(
		// Config
		fx.Supply(cfg),

		// Repository
		repository.Module,

		// Infrastructure
		logging.Module,

		postgres.Module,

		fiberFX.Module(cfg),

		// Main Activity
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
