package pubsub

import (
	"context"

	"cloud.google.com/go/internal/pubsub"
)

// Implements events.Handler interface
type Handler interface {
	Handle(ctx context.Context, payload *pubsub.Message)
}
