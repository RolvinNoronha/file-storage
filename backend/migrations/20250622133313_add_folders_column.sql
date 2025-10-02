-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS folders (
    id BIGSERIAL PRIMARY KEY,                            -- Auto-incrementing ID using BIGSERIAL
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,     -- Created timestamp with timezone
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,     -- Updated timestamp with timezone
    deleted_at TIMESTAMPTZ,                               -- Optional deleted timestamp with timezone
    name VARCHAR(100) NOT NULL,                           -- Folder name (cannot be NULL)
    user_id BIGINT NOT NULL,                              -- Foreign key referencing users table
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE  -- Foreign key to users table with cascading delete
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS folders;
-- +goose StatementEnd
