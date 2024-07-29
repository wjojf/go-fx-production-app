package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

var (
	cl *pubsub.Client = nil
)

func NewClient(opts ClientOpts) (*pubsub.Client, error) {
	if cl != nil {
		return cl, nil
	}

	keyOption := option.WithCredentialsFile(opts.KeyFilepath)

	client, err := pubsub.NewClient(opts.Ctx, opts.ProjectID, keyOption)
	if err != nil {
		return nil, err
	}

	cl = client

	return cl, nil
}

type ClientOpts struct {
	Ctx         context.Context
	ProjectID   string
	KeyFilepath string
}
