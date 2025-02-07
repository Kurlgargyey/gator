-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: GetPostsForUser :many
WITH user_follows as (
    SELECT feed_id from feed_follows
    WHERE feed_follows.user_id = $1
)

SELECT posts.* FROM posts
INNER JOIN user_follows ON posts.feed_id = user_follows.feed_id
ORDER BY published_at DESC
LIMIT sqlc.arg(post_limit);