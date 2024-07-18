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

func (r UserRepository) GetUserByID(id string) (models.User, error) {
	var err error

	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return models.User{}, err
	}

	defer conn.Release()

	var query string = "SELECT username FROM users WHERE id = $1"
	var username string

	err = conn.QueryRow(context.Background(), query, id).Scan(&username)

	return models.NewUser(id, username)
}

func (r UserRepository) SaveUser(user models.UserValueObject) (models.User, error) {
	var err error
	var ctx context.Context = context.Background()

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return models.User{}, err
	}

	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return models.User{}, err
	}

	var query string = "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username"

	var id, name string

	row := tx.QueryRow(ctx, query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id, &name); err != nil {
		tx.Rollback(ctx)
		return models.User{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return models.User{}, err
	}

	return models.NewUser(id, name)
}
