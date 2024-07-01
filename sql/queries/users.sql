-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, password, api_key)
VALUES ($1, $2, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE api_key = $1;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1;

-- name: DeleteUser :one
DELETE FROM users WHERE api_key = $1
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name = $1, password = $2, updated_at = $3
WHERE api_key = $4
RETURNING *;