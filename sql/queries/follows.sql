-- name: CreateFeedFollow :one
WITH inserted_follow as (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
) SELECT inserted_follow.*, users.name as user_name, feeds.name as feed_name from inserted_follow
    INNER JOIN users ON users.id = inserted_follow.user_id
    INNER JOIN feeds ON feeds.id = inserted_follow.feed_id;

-- name: GetFeedFollowsForUser :many
WITH user_follows as (
    SELECT * from feed_follows
    WHERE feed_follows.user_id = $1
)   SELECT user_follows.*, users.name as user_name, feeds.name as feed_name from user_follows
    INNER JOIN users ON users.id = user_follows.user_id
    INNER JOIN feeds ON feeds.id = user_follows.feed_id;