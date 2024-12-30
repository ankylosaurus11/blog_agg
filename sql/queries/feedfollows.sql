-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
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

SELECT 
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds
ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users
ON inserted_feed_follow.user_id = users.id;

-- name: GetFeed :one
SELECT feeds.name, feeds.id
FROM feeds
WHERE $1 = feeds.url;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
INNER JOIN users ON users.id = feed_follows.user_id
WHERE $1 = feed_follows.user_id;