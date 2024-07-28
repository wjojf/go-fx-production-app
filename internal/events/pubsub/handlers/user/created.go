package user

import (
	"log/slog"

	"github.com/wjojf/go-uber-fx/internal/domain/users/repository"
)

type UserCreatedHandler struct {
	log *slog.Logger
	r   repository.UsersRepository
}
