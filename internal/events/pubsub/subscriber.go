package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type Subscriber struct {
	cl *pubsub.Client

	opFunc     OperationFunc
	configFunc SubscriptionConfigurator
	ctxFunc    ContextFunc
}

func NewSubscriber(cl *pubsub.Client, opFunc OperationFunc, configFunc SubscriptionConfigurator) *Subscriber {
	return &Subscriber{
		cl:         cl,
		opFunc:     opFunc,
		configFunc: configFunc,
	}
}

func (s Subscriber) Subscribe(topicName string, handler Handler) error {
	var ctx context.Context = s.ctxFunc()

	sub, err := s.getSub(ctx, topicName)
	if err != nil {
		return err
	}

	return sub.Receive(ctx, handler.Handle)
}

func (s Subscriber) getSub(ctx context.Context, topicName string) (*pubsub.Subscription, error) {
	topic, err := GetTopic(s.cl, topicName)
	if err != nil {
		return nil, err
	}

	opId := s.opFunc(topicName)
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
