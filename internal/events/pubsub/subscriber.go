package pubsub

import (
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
	"github.com/wjojf/go-uber-fx/internal/events"
)

type Subscriber struct {
	cl *pubsub.Client

	opFunc     OperationFunc
	configFunc SubscriptionConfigurator
	ctxFunc    ContextFunc
}

type AdaptedSubscriber struct {
	s Subscriber
}

func NewAdaptedSubscriber(s Subscriber) AdaptedSubscriber {
	return AdaptedSubscriber{
		s: s,
	}
}

func NewSubscriber(cl *pubsub.Client, opFunc OperationFunc, configFunc SubscriptionConfigurator, ctxFunc ContextFunc) Subscriber {
	return Subscriber{
		cl:         cl,
		opFunc:     opFunc,
		configFunc: configFunc,
		ctxFunc:    ctxFunc,
	}
}

func (s Subscriber) Subscribe(topicName string, handler Handler) error {
	var ctx context.Context = s.ctxFunc()

	sub, err := s.getSub(ctx, topicName, handler)
	if err != nil {
		return err
	}

	go sub.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
		handler.Handle(ctx, message)
	})

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
