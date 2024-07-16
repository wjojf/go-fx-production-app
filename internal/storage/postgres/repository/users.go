package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wjojf/go-uber-fx/internal/domain/users/models"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return UserRepository{pool: pool}
}

func (r UserRepository) GetUserByID(id int) (models.User, error) {
	var err error

	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return models.User{}, err
	}

	defer conn.Release()

	if err = conn.Ping(context.Background()); err != nil {
		return models.User{}, err
	}

	return models.NewUser(id, "John Doe")
}
