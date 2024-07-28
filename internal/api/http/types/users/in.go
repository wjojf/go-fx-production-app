package users

import "github.com/wjojf/go-uber-fx/internal/domain/users/models"

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r CreateUserRequest) ToValueObject() (models.UserValueObject, error) {
	return models.NewUserValueObject(r.Username, r.Email, r.Password, false)
}
