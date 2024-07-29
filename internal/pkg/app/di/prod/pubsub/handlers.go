package pubsub

import (
	"context"
	"log/slog"

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

	// logger
	log *slog.Logger,

) {
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				// Verify user handler
				vh := pubsub.NewAdaptedHandler(userVerifyHandler)
				err := subscriber.Subscribe(events.TopicUserCreated, vh)
				if err != nil {
					log.Error("Failed to subscribe to user verify handler", slog.Any("err", err))
				}

				return nil
			},
		},
	)
}
