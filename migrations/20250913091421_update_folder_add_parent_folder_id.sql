-- +goose Up
-- +goose StatementBegin
ALTER TABLE folders 
    ADD COLUMN parent_folder_id BIGINT;

ALTER TABLE folders
    ADD CONSTRAINT fk_parent_folder FOREIGN KEY (parent_folder_id) REFERENCES folders(id) ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE folders
    DROP CONSTRAINT fk_parent_folder;

ALTER TABLE folders
    DROP COLUMN parent_folder_id;

-- +goose StatementEnd
