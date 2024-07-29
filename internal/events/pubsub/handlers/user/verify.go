package user

import (
	"context"
	"encoding/json"
	"log/slog"

	"cloud.google.com/go/pubsub"
	"github.com/wjojf/go-uber-fx/internal/domain/users/events"
	"github.com/wjojf/go-uber-fx/internal/domain/users/repository"
)

type VerifyHandler struct {
	log *slog.Logger
	r   repository.UsersRepository
}

func NewVerifyHandler(log *slog.Logger, r repository.UsersRepository) VerifyHandler {
	return VerifyHandler{
		log: log,
		r:   r,
	}
}

func (h VerifyHandler) ID() string {
	return "verify-user-handler"
}

func (h VerifyHandler) Handle(ctx context.Context, message *pubsub.Message) error {
	var payload events.UserCreatedPayload
	if err := json.Unmarshal(message.Data, &payload); err != nil {
		h.log.Error("failed to unmarshal payload: %v", slog.Any("err", err))
		return err
	}

	h.log.Info("handling user created event", slog.Any("payload", payload))

	return nil
}
