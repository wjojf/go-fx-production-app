package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/wjojf/go-uber-fx/internal/events"
)

// Implements events.Handler interface
type Handler interface {
	ID() string
	Handle(ctx context.Context, message *pubsub.Message) error
}

type AdaptedHandler struct {
	h Handler
}

func NewAdaptedHandler(h Handler) events.Handler {
	return AdaptedHandler{
		h: h,
	}
}

func (ah AdaptedHandler) Handle(ctx context.Context, message interface{}) error {
	msg, ok := message.(*pubsub.Message)
	if !ok {
		return nil
	}

	return ah.h.Handle(ctx, msg)
}
