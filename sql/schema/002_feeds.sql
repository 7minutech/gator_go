-- +goose Up
CREATE TABLE feeds (
    id  uuid PRIMARY KEY NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    url VARCHAR(255) UNIQUE NOT NULL,
    user_id uuid references users(id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE feeds;
