package users

import (
	"log/slog"

	"github.com/wjojf/go-uber-fx/internal/domain/users/repository"
)

type Handler struct {
	r   repository.UsersRepository
	log *slog.Logger
}

func New(r repository.UsersRepository, log *slog.Logger) Handler {
	return Handler{r: r, log: log}
}
