-- name: CreatePostLike :one
INSERT INTO post_likes(id, created_at, updated_at, user_id, post_id)
VALUES($1, $2, $2, $3, $4)
RETURNING *;

-- name: DeletePostLike :one
DELETE FROM post_likes WHERE id = $1
RETURNING *; 

-- name: GetPostLikesByUser :many
SELECT * FROM post_likes WHERE user_id = $1;

-- name: GetPostLikeById :one
SELECT * FROM post_likes WHERE id = $1;