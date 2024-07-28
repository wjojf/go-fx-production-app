package pubsub

import (
	"github.com/wjojf/go-uber-fx/internal/events"
	"github.com/wjojf/go-uber-fx/internal/events/pubsub/handlers"
	"go.uber.org/fx"
)

// This module needs to be provided with following types:
//   - pubsub.OperationFunc
//   - pubsub.ContextFunc
//   - pubsub.SubscriptionConfigurator
//   - pubsub.ClientOpts
var Module = fx.Module(
	"Google Pub/Sub",

	fx.Provide(NewClient),

	fx.Provide(NewSubscriber),
	fx.Provide(
		fx.Annotate(
			NewAdaptedSubscriber,
			fx.As(new(events.Subscriber)),
		),
	),

	fx.Provide(
		fx.Annotate(
			NewPublisher,
			fx.As(new(events.Publisher)),
		),
	),

	handlers.Module,
)
