-- +goose Up
CREATE TABLE posts (
    id uuid PRIMARY KEY,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL UNIQUE,
    description Text NOT NULL,
    published_at VARCHAR(255) NOT NULL,
    feed_id uuid references feeds(id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE posts;