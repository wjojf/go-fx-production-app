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

	// map of event handlers: extend here
	var eventHandlers = map[string][]pubsub.Handler{
		events.TopicUserCreated: {
			userVerifyHandler,
		},
	}

	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				// subscribe to the events
				for topic, handlers := range eventHandlers {
					for _, handler := range handlers {
						err := subscriber.Subscribe(topic, pubsub.NewAdaptedHandler(handler))
						if err != nil {
							log.Error("Failed to subscribe",
								slog.String("topic", topic), slog.Any("err", err),
							)
						}
					}
				}

				return nil
			},
		},
	)
}
