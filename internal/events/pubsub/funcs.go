package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

type SubscriptionConfigurator func(topic *pubsub.Topic) pubsub.SubscriptionConfig

type OperationFunc func(topicName string, handler Handler) string

type ContextFunc func() context.Context

func DefaultSubscriptionConfigurator(topic *pubsub.Topic) pubsub.SubscriptionConfig {
	return pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 10,
	}
}

func DefaultOperationGenerator(topicName string, handler Handler) string {
	return fmt.Sprintf("%s-%s", topicName, handler.ID())
}

func DefaultContextFunc() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}
