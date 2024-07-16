package repository

import "github.com/wjojf/go-uber-fx/internal/domain/users/models"

type UsersRepository interface {
	GetUserByID(id int) (models.User, error)
}
