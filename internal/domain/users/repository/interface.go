package repository

import "github.com/wjojf/go-uber-fx/internal/domain/users/models"

type UsersRepository interface {
	GetUserByID(id string) (models.User, error)
	SaveUser(user models.UserValueObject) (models.User, error)
}
