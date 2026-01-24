-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    ratings INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE users;