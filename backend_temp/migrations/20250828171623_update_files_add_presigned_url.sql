-- +goose Up
-- +goose StatementBegin
ALTER TABLE files 
    ADD COLUMN file_url VARCHAR(255),
    ADD COLUMN file_url_expiry DATE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE files
    DROP COLUMN IF EXISTS file_url,                  -- Drop file_url column
    DROP COLUMN IF EXISTS file_url_expiry; 
-- +goose StatementEnd
