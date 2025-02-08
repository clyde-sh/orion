-- name: FindUserById :one
SELECT * FROM thorfinn_users WHERE id = $1;

-- name: FindUserByEmail :one
SELECT * FROM thorfinn_users WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM thorfinn_users ORDER BY created_at DESC;

-- name: CreateUser :one
INSERT INTO thorfinn_users (id, email, password_hash) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUser :one
UPDATE thorfinn_users SET email = $2, password_hash = $3 WHERE id = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM thorfinn_users WHERE id = $1;
