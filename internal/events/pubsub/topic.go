package pubsub

import (
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
)

var (
	ErrCreatingTopic = errors.New("error creating topic")
)

func GetTopic(cl *pubsub.Client, topic string) (*pubsub.Topic, error) {

	var ctx context.Context = context.TODO()

	t := cl.Topic(topic)

	ok, err := t.Exists(ctx)

	if err != nil {
		return nil, err
	}

	if !ok {
		t, err = cl.CreateTopic(ctx, topic)
		if err != nil {
			return nil, ErrCreatingTopic
		}
	}

	return t, nil
}
