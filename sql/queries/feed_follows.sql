--name CreateFeedFollow :many
WITH inserted_feed_follow as (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    )
)
SELECT inserted_feed_follow.*, users.name, feeds.name
FROM inserted_feed_follow 
JOIN users
ON inserted_feed_follow.user_id = users.id
FROM inserted_feed_follow 
JOIN feeds
ON inserted_feed_follow.id = feed_id;