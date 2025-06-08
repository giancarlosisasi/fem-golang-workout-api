-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tokens (
  hash BYTEA PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id) on DELETE CASCADE,
  expiry TIMESTAMP(0) with TIME ZONE NOT NULL,
  scope TEXT NOT NULL
  -- created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  -- updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tokens;
-- +goose StatementEnd