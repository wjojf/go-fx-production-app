package postgres

import (
	"github.com/wjojf/go-uber-fx/internal/storage"
	"github.com/wjojf/go-uber-fx/internal/storage/postgres/outbox"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"postgres",

	fx.Provide(NewPool),

	fx.Provide(
		fx.Annotate(
			outbox.NewProducer,
			fx.As(new(storage.OutboxProducer)),
		),
	),
)
