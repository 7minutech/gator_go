-- +goose Up
CREATE TABLE users (
    id  uuid PRIMARY KEY NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;