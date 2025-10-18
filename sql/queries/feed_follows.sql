-- name: CreateFeedFollow :one
WITH inserted_feed_follow as (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT inserted_feed_follow.*, users.name as user_name, feeds.name as feed_name
FROM inserted_feed_follow 
JOIN users
ON inserted_feed_follow.user_id = users.id 
JOIN feeds
ON inserted_feed_follow.id = feed_id;

-- name: GetFeedFollowsForUser :many
SELECT * FROM feed_follows
JOIN feeds 
ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;