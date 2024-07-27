package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type Manager struct {
	cl *pubsub.Client

	opFunc     OperationFunc
	configFunc SubscriptionConfigurator
	ctxFunc    ContextFunc
}

func NewManager(cl *pubsub.Client, opFunc OperationFunc, configFunc SubscriptionConfigurator) *Manager {
	return &Manager{
		cl:         cl,
		opFunc:     opFunc,
		configFunc: configFunc,
	}
}

func (m *Manager) Subscribe(topicName string, handler Handler) error {
	var ctx context.Context = m.ctxFunc()

	topic, err := GetTopic(m.cl, topicName)
	if err != nil {
		return err
	}

	opId := m.opFunc(topicName)
	cfg := m.configFunc(topic)

	sub := m.cl.Subscription(opId)

	ok, err := sub.Exists(ctx)
	if err != nil {
		return err
	}

	if !ok {
		sub, err = m.cl.CreateSubscription(ctx, opId, cfg)
		if err != nil {
			return err
		}
	}

	return sub.Receive(ctx, handler.Handle)
}
