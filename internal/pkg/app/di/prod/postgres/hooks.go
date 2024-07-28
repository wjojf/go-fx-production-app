package postgres

import (
	"context"

	"github.com/wjojf/go-uber-fx/internal/storage"
	"go.uber.org/fx"
)

func PostgresJobs(lc fx.Lifecycle, p storage.OutboxProducer) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return p.StartProducing()
			},
		},
	)
}
