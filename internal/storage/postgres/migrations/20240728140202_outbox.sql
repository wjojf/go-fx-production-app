-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS outbox (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    event_name TEXT NOT NULL,
    payload JSONB NOT NULL,
    is_published BOOLEAN DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS outbox;
-- +goose StatementEnd
