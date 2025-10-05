-- +goose Up
-- +goose StatementBegin
ALTER TABLE files
ALTER COLUMN file_url TYPE TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE files
ALTER COLUMN file_url VARCHAR(255);
-- +goose StatementEnd
