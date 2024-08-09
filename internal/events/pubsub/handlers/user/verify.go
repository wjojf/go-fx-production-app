package user

import (
	"context"
	"encoding/json"
	"log/slog"

	"cloud.google.com/go/pubsub"
	"github.com/wjojf/go-uber-fx/internal/domain/users/events"
	"github.com/wjojf/go-uber-fx/internal/domain/users/models"
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

	vo := models.UserValueObjectPartial{
		UserValueObject: models.UserValueObject{
			IsVerified: true,
		},
		Fields: []string{models.FieldIsVerified},
	}

	user, err := h.r.UpdateUserByID(ctx, payload.UserID, vo)
	if err != nil {
		h.log.Error("failed to update user", slog.Any("err", err))
		return err
	}

	h.log.Info("user verified", slog.Any("user_id", user.ID()))

	return nil
}
