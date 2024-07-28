package pubsub

import (
	"log/slog"

	"cloud.google.com/go/pubsub"
)

type Publisher struct {
	cl      *pubsub.Client
	ctxFunc ContextFunc
	log     *slog.Logger
}

func NewPublisher(cl *pubsub.Client, ctxFunc ContextFunc, log *slog.Logger) *Publisher {
	return &Publisher{
		cl:      cl,
		ctxFunc: ctxFunc,
		log:     log,
	}
}

func (p Publisher) Publish(topicName string, payload []byte) error {

	ctx := p.ctxFunc()
	topic, err := GetTopic(p.cl, topicName)
	if err != nil {
		p.log.Error("Failed to get Pub/Sub topic. Error: %s", err.Error())
		return err
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: payload,
	})

	msgID, err := result.Get(ctx)
	if err != nil {
		p.log.Error("Failed to publish Pub/Sub message. Error: %s", err.Error())
		return err
	}

	p.log.Info("Published message with ID: %s", msgID)
	return nil
}
