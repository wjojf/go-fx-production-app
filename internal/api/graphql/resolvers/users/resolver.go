package users

import (
	"log/slog"

	"github.com/wjojf/go-uber-fx/internal/domain/users/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	r   repository.UsersRepository
	log *slog.Logger
}

func NewResolver(r repository.UsersRepository, log *slog.Logger) *Resolver {
	return &Resolver{
		r:   r,
		log: log,
	}
}
