package users

import "github.com/wjojf/go-uber-fx/internal/domain/users/models"

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func FromUser(user models.User) UserResponse {
	return UserResponse{
		ID:       user.ID(),
		Username: user.Username(),
	}
}
