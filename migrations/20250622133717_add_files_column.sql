-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS files (
    id BIGSERIAL PRIMARY KEY,                            -- Auto-incrementing ID using BIGSERIAL
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,     -- Created timestamp (with timezone)
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,     -- Updated timestamp (with timezone)
    deleted_at TIMESTAMPTZ,                               -- Optional: Deleted timestamp (with timezone)
    name VARCHAR(100) NOT NULL,                           -- Name of the file
    path VARCHAR(255) NOT NULL,                           -- Path of the file
    file_type VARCHAR(100),                               -- File type (optional)
    file_size INTEGER,                                    -- File size (optional)
    user_id BIGINT NOT NULL,                              -- Foreign key referencing users table
    folder_id BIGINT,                                    -- Optional foreign key referencing folders table
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,    -- Foreign key to users table
    CONSTRAINT fk_folder FOREIGN KEY (folder_id) REFERENCES folders (id) ON DELETE SET NULL   -- Foreign key to folders table
);

CREATE INDEX IF NOT EXISTS idx_folder_id ON files (folder_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files;
-- +goose StatementEnd
