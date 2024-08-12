package verify

import (
	"context"
	"log/slog"

	"github.com/wjojf/go-uber-fx/internal/domain/users/events"
	"github.com/wjojf/go-uber-fx/internal/domain/users/models"
	"github.com/wjojf/go-uber-fx/internal/domain/users/repository"
)

type Handler struct {
	r   repository.UsersRepository
	log *slog.Logger
}

func NewHandler(r repository.UsersRepository, log *slog.Logger) Handler {
	return Handler{
		r:   r,
		log: log,
	}
}

func (h Handler) VerifyUser(ctx context.Context, payload events.UserCreatedPayload) error {
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
