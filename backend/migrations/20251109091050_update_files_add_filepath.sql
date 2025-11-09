-- +goose Up
-- +goose StatementBegin
ALTER TABLE files
ADD COLUMN file_path VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE files
DROP COLUMN file_path;
-- +goose StatementEnd
