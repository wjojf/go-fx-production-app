package users

import "github.com/wjojf/go-uber-fx/internal/domain/users/models"

type UserResponse struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	IsVerified bool   `json:"is_verified"`
}

func FromUser(user models.User) UserResponse {
	return UserResponse{
		ID:         user.ID(),
		Username:   user.Username(),
		IsVerified: user.IsVerified(),
	}
}
