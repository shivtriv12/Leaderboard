-- +goose Up
CREATE INDEX idx_username_trgm ON users USING GIN(username gin_trgm_ops);

-- +goose Down
DROP INDEX idx_username_trgm;