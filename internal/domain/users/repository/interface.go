package repository

import (
	"context"

	"github.com/wjojf/go-uber-fx/internal/domain/users/models"
)

type UsersRepository interface {
	GetUserByID(ctx context.Context, id string) (models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	SaveUser(ctx context.Context, user models.UserValueObject) (models.User, error)
	UpdateUserByID(ctx context.Context, userID string, user models.UserValueObjectPartial) (models.User, error)
}
