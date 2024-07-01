-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $2, $3, $4)
RETURNING *;

-- name: DeleteFeedFollow :one
DELETE FROM feed_follows WHERE id = $1
RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE user_id = $1;