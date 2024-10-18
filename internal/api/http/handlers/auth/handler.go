package auth

import (
	"github.com/wjojf/go-uber-fx/internal/api/http/service"
	users "github.com/wjojf/go-uber-fx/internal/domain/users/repository"
)

type Handler struct {
	r          users.UsersRepository
	jwtService service.JwtService
}

func New(r users.UsersRepository, jwtService service.JwtService) Handler {
	return Handler{
		r:          r,
		jwtService: jwtService,
	}
}
