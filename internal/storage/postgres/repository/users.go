package repository

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	types "github.com/wjojf/go-uber-fx/internal/domain/users/events"
	"github.com/wjojf/go-uber-fx/internal/domain/users/models"
	"github.com/wjojf/go-uber-fx/internal/events"
	"github.com/wjojf/go-uber-fx/internal/storage/postgres/mapper"
	"github.com/wjojf/go-uber-fx/internal/storage/postgres/outbox"
)

type UserRepository struct {
	pool *pgxpool.Pool
	log  *slog.Logger
}

func NewUserRepository(pool *pgxpool.Pool, log *slog.Logger) UserRepository {
	return UserRepository{
		pool: pool,
		log:  log,
	}
}

func (r UserRepository) GetUserByID(ctx context.Context, id string) (models.User, error) {
	var err error

	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return models.User{}, err
	}

	defer conn.Release()

	var query string = "SELECT username, is_verified FROM users WHERE id = $1"
	var username string
	var isVerified bool

	err = conn.QueryRow(context.Background(), query, id).Scan(&username, &isVerified)
	if err != nil {
		return models.User{}, err
	}

	return models.NewUser(id, username, isVerified)
}

func (r UserRepository) SaveUser(ctx context.Context, user models.UserValueObject) (models.User, error) {
	var err error

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return models.User{}, err
	}

	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return models.User{}, err
	}

	var query string = `
		INSERT INTO users (username, email, password, is_verified)
		VALUES ($1, $2, $3, $4)
		RETURNING id, username, is_verified
	`

	var id, name string
	var isVerified bool

	row := tx.QueryRow(ctx, query, user.Username, user.Email, user.Password, user.IsVerified)
	if err := row.Scan(&id, &name, &isVerified); err != nil {
		tx.Rollback(ctx)
		return models.User{}, err
	}

	// Marshal the event payload
	payload, err := json.Marshal(
		types.UserCreatedPayload{
			EventID: uuid.NewString(),
			UserID:  id,
		},
	)
	if err != nil {
		tx.Rollback(ctx)
		return models.User{}, errors.Wrap(err, "failed to marshal payload")
	}

	// Store the event in the outbox
	if err := outbox.StoreOutboxEvent(
		tx, outbox.OutboxEvent{
			EventName: events.TopicUserCreated,
			Payload:   payload,
		},
	); err != nil {
		tx.Rollback(ctx)
		return models.User{}, errors.Wrap(err, "failed to store outbox event")
	}

	if err = tx.Commit(ctx); err != nil {
		return models.User{}, errors.Wrap(err, "failed to commit transaction")
	}

	return models.NewUser(id, name, isVerified)
}

func (r UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var err error
	var users []models.User

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return users, err
	}

	defer conn.Release()

	var query string = "SELECT id, username, is_verified FROM users"
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var id, username string
		var isVerified bool

		if err = rows.Scan(&id, &username, &isVerified); err != nil {
			return users, err
		}

		user, err := models.NewUser(id, username, isVerified)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r UserRepository) UpdateUserByID(ctx context.Context, userID string, user models.UserValueObjectPartial) (models.User, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return models.User{}, err
	}

	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return models.User{}, err
	}

	var query string = mapper.GetUpdateUserQuery(userID, user)

	_, err = tx.Exec(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		return models.User{}, err
	}

	// Marshal the event payload
	payload, err := json.Marshal(
		types.UserUpdatedPayload{
			EventID: uuid.NewString(),
			UserID:  userID,
		},
	)
	if err != nil {
		tx.Rollback(ctx)
		return models.User{}, errors.Wrap(err, "failed to marshal payload")
	}

	if err := outbox.StoreOutboxEvent(
		tx, outbox.OutboxEvent{
			EventName: events.TopicUserUpdated,
			Payload:   payload,
		},
	); err != nil {
		tx.Rollback(ctx)
		return models.User{}, errors.Wrap(err, "failed to store outbox event")
	}

	if err = tx.Commit(ctx); err != nil {
		return models.User{}, err
	}

	return r.GetUserByID(ctx, userID)
}

func (r UserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var err error

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return models.User{}, err
	}

	defer conn.Release()

	var query string = "SELECT id, username, is_verified FROM users WHERE email = $1"
	var id, username string
	var isVerified bool

	err = conn.QueryRow(ctx, query, email).Scan(&id, &username, &isVerified)
	if err != nil {
		return models.User{}, err
	}

	return models.NewUser(id, username, isVerified)
}
