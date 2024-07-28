-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN is_verified BOOLEAN DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN is_verified;
-- +goose StatementEnd
