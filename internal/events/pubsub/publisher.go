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
		p.log.Error("Failed to get Pub/Sub topic. Error: %s", slog.Any("err", err))
		return err
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: payload,
	})

	msgID, err := result.Get(ctx)
	if err != nil {
		p.log.Error("Failed to publish Pub/Sub message", slog.Any("err", err))
		return err
	}

	p.log.Info(
		"Published Pub/Sub message",
		slog.Any("msgID", msgID),
		slog.String("topic", topicName),
		slog.String("payload", string(payload)),
	)
	return nil
}
