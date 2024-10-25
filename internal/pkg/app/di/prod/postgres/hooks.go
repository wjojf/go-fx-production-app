package postgres

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wjojf/go-uber-fx/internal/storage"
	"go.uber.org/fx"
)

func PostgresJobs(lc fx.Lifecycle, log *slog.Logger, p storage.OutboxProducer, pool *pgxpool.Pool) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go p.StartProducing()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info("Closing postgres connection pool")
				pool.Close()
				return nil
			},
		},
	)

	lc.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				log.Info("Stopping outbox producer")
				p.Stop()
				return nil
			},
		},
	)

}
