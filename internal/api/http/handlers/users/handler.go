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
	return Handler{
		r:   r,
		log: log,
	}
}

func (h Handler) Repository() repository.UsersRepository {
	return h.r
}

func (h Handler) Logger() *slog.Logger {
	return h.log
}
