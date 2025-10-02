-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
  DROP COLUMN IF EXISTS first_name,
  DROP COLUMN IF EXISTS last_name,
  DROP COLUMN IF EXISTS email,
  DROP COLUMN IF EXISTS password;

ALTER TABLE users
  ADD COLUMN IF NOT EXISTS username VARCHAR(100) NOT NULL,
  ADD COLUMN IF NOT EXISTS password VARCHAR(255) NOT NULL;

ALTER TABLE users
  ADD COLUMN IF NOT EXISTS id BIGSERIAL PRIMARY KEY,
  ADD COLUMN IF NOT EXISTS created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ;

-- Create the function for the trigger
CREATE OR REPLACE FUNCTION update_users_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_users_updated_at
  BEFORE UPDATE ON users
  FOR EACH ROW
  EXECUTE FUNCTION update_users_updated_at();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
  ADD COLUMN first_name VARCHAR(255) NOT NULL,
  ADD COLUMN last_name VARCHAR(255) NOT NULL,
  ADD COLUMN email VARCHAR(255) UNIQUE NOT NULL,
  ADD COLUMN password VARCHAR(255) NOT NULL;
    

-- Drop the newly added column
ALTER TABLE users
  DROP COLUMN IF EXISTS username;

-- Optionally, revert GORM model fields
ALTER TABLE users
  DROP COLUMN IF EXISTS id,
  DROP COLUMN IF EXISTS created_at,
  DROP COLUMN IF EXISTS updated_at,
  DROP COLUMN IF EXISTS deleted_at;
-- +goose StatementEnd
