package resolvers

import (
	"log/slog"

	"github.com/wjojf/go-uber-fx/internal/domain/users/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log *slog.Logger
	r   repository.UsersRepository
}

func New(log *slog.Logger, r repository.UsersRepository) *Resolver {
	return &Resolver{
		log: log,
		r:   r,
	}
}
