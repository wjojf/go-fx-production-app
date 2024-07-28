package pubsub

import (
	"context"

	"github.com/wjojf/go-uber-fx/internal/events"
	"github.com/wjojf/go-uber-fx/internal/events/pubsub"
	"github.com/wjojf/go-uber-fx/internal/events/pubsub/handlers/user"
	"go.uber.org/fx"
)

func PubSubHooks(
	lc fx.Lifecycle,

	// Subscriber for the events
	subscriber events.Subscriber,

	// Handlers for the subscriber
	userVerifyHandler user.VerifyHandler,

) {
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				subscriber.Subscribe(events.TopicUserCreated, pubsub.NewAdaptedHandler(userVerifyHandler))

				return nil
			},
		},
	)
}
