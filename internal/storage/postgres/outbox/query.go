package outbox

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func StoreOutboxEvent(tx pgx.Tx, dto OutboxEvent) error {
	var query string = `
		INSERT INTO outbox (event_name, payload)
		VALUES ($1, $2)
	`
	_, err := tx.Exec(context.Background(), query, dto.EventName, dto.Payload)
	return err
}

func GetOutboxEventsToProduce(pool *pgxpool.Pool) ([]OutboxEvent, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	var query string = `SELECT id, event_name, payload FROM outbox WHERE is_published=false;`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var outboxEvents []OutboxEvent
	for rows.Next() {
		var dto OutboxEvent
		err = rows.Scan(&dto.ID, &dto.EventName, &dto.Payload)
		if err != nil {
			return nil, err
		}
		outboxEvents = append(outboxEvents, dto)
	}

	return outboxEvents, nil
}

func MarkOutboxEventAsProduced(pool *pgxpool.Pool, id int) error {

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}

	defer conn.Release()

	var query string = `UPDATE outbox SET is_published=true WHERE id=$1;`
	_, err = conn.Exec(context.Background(), query, id)

	return err
}

func DeleteProducesOutboxEvents(pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}

	defer conn.Release()

	var query string = `DELETE FROM outbox WHERE is_published=true;`
	_, err = conn.Exec(context.Background(), query)

	return err
}
