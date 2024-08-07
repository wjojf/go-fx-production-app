package adapters

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/tracelog"
)

type SlogTracer struct {
	log *slog.Logger
}

func NewSlogTracer(log *slog.Logger) *SlogTracer {
	return &SlogTracer{log: log}
}

func (t *SlogTracer) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {

	t.log.Debug("in slog trace", slog.String("level_received", level.String()))

	switch level {
	case tracelog.LogLevelDebug:
		t.log.Debug(msg, slog.Any("meta", data))
	case tracelog.LogLevelInfo:
		t.log.Info(msg, slog.Any("meta", data))
	case tracelog.LogLevelWarn:
		t.log.Warn(msg, slog.Any("meta", data))
	case tracelog.LogLevelError:
		t.log.Error(msg, slog.Any("meta", data))
	}
}
