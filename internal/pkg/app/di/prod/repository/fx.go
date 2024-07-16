package repository

import (
	users "github.com/wjojf/go-uber-fx/internal/domain/users/repository"
	postgres "github.com/wjojf/go-uber-fx/internal/storage/postgres/repository"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"repository-prod",

	fx.Provide(
		fx.Annotate(
			postgres.NewUserRepository,
			fx.As(new(users.UsersRepository)),
		),
	),
)
