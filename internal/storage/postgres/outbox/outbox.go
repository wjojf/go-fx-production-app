package outbox

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func StoreOutboxEvent(tx pgx.Tx, dto OutboxEvent) error {
	var query string = `
		INSERT INTO outbox (event_name, payload)
		VALUES ($1, $2)
	`
	_, err := tx.Exec(context.Background(), query, dto.EventName, dto.Payload)
	return err
}
