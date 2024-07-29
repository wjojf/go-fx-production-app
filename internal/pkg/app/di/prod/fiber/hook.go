package fiber

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"go.uber.org/fx"
)

func ServerHooks(
	lc fx.Lifecycle,
	cfg config.Config,
	log *slog.Logger,
	http *fiber.App,
) {
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				log.Info("Starting HTTP Server...")
				go http.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
				return nil
			},
			OnStop: func(context.Context) error {
				log.Info("Stopping HTTP Server...")
				return http.Shutdown()
			},
		},
	)
}
