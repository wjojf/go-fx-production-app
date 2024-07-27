package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

var (
	cl *pubsub.Client = nil
)

func NewClient(opts ClientOpts) (*pubsub.Client, error) {
	if cl != nil {
		return cl, nil
	}

	client, err := pubsub.NewClient(opts.Ctx, opts.ProjectID)
	if err != nil {
		return nil, err
	}

	cl = client

	return cl, nil
}

type ClientOpts struct {
	Ctx       context.Context
	ProjectID string
}
