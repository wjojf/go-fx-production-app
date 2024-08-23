package pubsub

import (
	"github.com/wjojf/go-uber-fx/internal/events/pubsub"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"pub-sub-prod",

	pubsub.Module,

	fx.Provide(GetClientOpts),
	fx.Provide(GetContextFunc),
	fx.Provide(GetOperationFunc),
	fx.Provide(GetSubscriptionConfigurator),
)
