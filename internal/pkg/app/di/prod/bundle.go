package prod

import (
	"log/slog"

	"github.com/wjojf/go-uber-fx/internal/api/graphql"

	"github.com/wjojf/go-uber-fx/internal/pkg/app/di/prod/domain"
	fiberFX "github.com/wjojf/go-uber-fx/internal/pkg/app/di/prod/fiber"
	postgresFX "github.com/wjojf/go-uber-fx/internal/pkg/app/di/prod/postgres"
	pubsubFX "github.com/wjojf/go-uber-fx/internal/pkg/app/di/prod/pubsub"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"github.com/wjojf/go-uber-fx/internal/pkg/logging"
	"github.com/wjojf/go-uber-fx/internal/storage/postgres"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func AppBundle(cfg config.Config) fx.Option {

	return fx.Options(
		// Config
		fx.Supply(cfg),

		// FX Event Loggers
		fx.WithLogger(
			func(log *slog.Logger) fxevent.Logger {
				return &fxevent.SlogLogger{
					Logger: log,
				}
			},
		),

		// Domain
		domain.Module,

		// Repository
		postgresFX.Repositories,

		// Infrastructure
		logging.Module,

		// Postgres Connection
		postgres.Module,

		// Fiber API
		fiberFX.Module(),

		// PubSub
		pubsubFX.Module,

		// GraphQL
		graphql.Module,

		// Main Activity
		fx.Invoke(
			// Start the fiber server
			fiberFX.ServerHooks,

			// Start the outbox producer
			postgresFX.PostgresJobs,
		),
	)
}

func ConsumerBundle(cfg config.Config) fx.Option {
	return fx.Options(
		// Config
		fx.Supply(cfg),

		// FX Event Loggers
		fx.WithLogger(
			func(log *slog.Logger) fxevent.Logger {
				return &fxevent.SlogLogger{
					Logger: log,
				}
			},
		),

		// Domain
		domain.Module,

		// Repository
		postgresFX.Repositories,

		// Infrastructure
		logging.Module,

		// Postgres Connection
		postgres.Module,

		// PubSub
		pubsubFX.Module,

		// Main Activity
		fx.Invoke(
			// Start event handlers and listeners
			pubsubFX.PubSubHooks,
		),
	)
}
