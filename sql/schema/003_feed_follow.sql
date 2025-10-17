-- +goose Up
CREATE TABLE feed_follows (
    id uuid PRIMARY KEY,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    user_id uuid references users(id) ON DELETE CASCADE NOT NULL,
    feed_id uuid references feeds(id) ON DELETE CASCADE NOT NULL,
    UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;