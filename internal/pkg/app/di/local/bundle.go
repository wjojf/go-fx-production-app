package local

import (
	"github.com/wjojf/go-uber-fx/internal/api/graphql"
	"github.com/wjojf/go-uber-fx/internal/pkg/app/di/local/domain"
	graphqlFX "github.com/wjojf/go-uber-fx/internal/pkg/app/di/local/graphql"
	postgresFX "github.com/wjojf/go-uber-fx/internal/pkg/app/di/local/postgres"
	pubsubFX "github.com/wjojf/go-uber-fx/internal/pkg/app/di/local/pubsub"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"github.com/wjojf/go-uber-fx/internal/pkg/logging"
	"github.com/wjojf/go-uber-fx/internal/pkg/tracing"
	"github.com/wjojf/go-uber-fx/internal/storage/postgres"
	"go.uber.org/fx"
)

func Bundle(cfg config.Config) fx.Option {
	return fx.Options(

		// Config
		fx.Supply(cfg),

		// Domain
		domain.Module,

		// Repository
		postgresFX.Repositories,

		// Infrastructure
		logging.Module,

		// Tracing
		tracing.Module,

		// Postgres Connection
		postgres.Module,

		// GraphQL API
		graphql.Module,

		// PubSub
		pubsubFX.Module,

		fx.Invoke(
			// Start event handlers and listeners
			pubsubFX.PubSubHooks,

			// Start the outbox producer
			postgresFX.PostgresJobs,

			// Start the graphql server
			graphqlFX.Hooks,
		),
	)
}
