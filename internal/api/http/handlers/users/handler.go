package users

import (
	"github.com/opentracing/opentracing-go"
	"log/slog"

	"github.com/wjojf/go-uber-fx/internal/domain/users/repository"
)

type Handler struct {
	r      repository.UsersRepository
	log    *slog.Logger
	tracer opentracing.Tracer
}

func New(r repository.UsersRepository, log *slog.Logger, tracer opentracing.Tracer) Handler {
	return Handler{
		r:      r,
		log:    log,
		tracer: tracer,
	}
}

func (h Handler) Repository() repository.UsersRepository {
	return h.r
}

func (h Handler) Logger() *slog.Logger {
	return h.log
}
