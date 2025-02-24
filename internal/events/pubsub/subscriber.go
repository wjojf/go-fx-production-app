package pubsub

import (
	"context"
	"errors"
	"log/slog"

	"cloud.google.com/go/pubsub"
	"github.com/wjojf/go-uber-fx/internal/events"
)

type Subscriber struct {
	cl *pubsub.Client

	opFunc     OperationFunc
	configFunc SubscriptionConfigurator
	ctxFunc    ContextFunc

	log *slog.Logger
}

type AdaptedSubscriber struct {
	s Subscriber
}

func NewAdaptedSubscriber(s Subscriber) AdaptedSubscriber {
	return AdaptedSubscriber{
		s: s,
	}
}

func NewSubscriber(
	cl *pubsub.Client,
	opFunc OperationFunc,
	configFunc SubscriptionConfigurator,
	ctxFunc ContextFunc,
	log *slog.Logger,
) Subscriber {
	return Subscriber{
		cl:         cl,
		opFunc:     opFunc,
		configFunc: configFunc,
		ctxFunc:    ctxFunc,
		log:        log,
	}
}

func (s Subscriber) Subscribe(topicName string, handler Handler) error {

	go func() {
		s.log.Info("Subscribing to topic", slog.String("topic", topicName))

		var ctx context.Context = s.ctxFunc()

		sub, err := s.getSub(ctx, topicName, handler)
		if err != nil {
			s.log.Error("Failed to get subscription", slog.Any("err", err))
			return
		}

		for {
			var ctx context.Context = s.ctxFunc()

			s.log.Debug("Listening for messages...", slog.String("topic", topicName))
			err := sub.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
				handler.Handle(ctx, message)
				message.Ack()
			})

			if err != nil {
				s.log.Error("Failed to receive message", slog.Any("err", err))
			}
		}
	}()

	return nil
}

func (s Subscriber) getSub(ctx context.Context, topicName string, handler Handler) (*pubsub.Subscription, error) {
	topic, err := GetTopic(s.cl, topicName)
	if err != nil {
		return nil, err
	}

	opId := s.opFunc(topicName, handler)
	cfg := s.configFunc(topic)

	sub := s.cl.Subscription(opId)

	ok, err := sub.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ok {
		return s.cl.CreateSubscription(ctx, opId, cfg)
	}

	return sub, nil
}

func (as AdaptedSubscriber) Subscribe(topicName string, handler events.Handler) error {

	h, ok := handler.(AdaptedHandler)
	if !ok {
		return errors.New("handler is not a pubsub.Handler")
	}

	return as.s.Subscribe(topicName, h.h)
}
