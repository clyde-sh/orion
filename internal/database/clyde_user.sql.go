// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: clyde_user.sql

package database

import (
	"context"
)

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM clyde_user
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, email, password_hash, recovery_code, created_at, updated_at, name, role FROM clyde_user
WHERE email = $1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (ClydeUser, error) {
	row := q.db.QueryRow(ctx, findUserByEmail, email)
	var i ClydeUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.RecoveryCode,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Role,
	)
	return i, err
}

const findUserById = `-- name: FindUserById :one
SELECT id, email, password_hash, recovery_code, created_at, updated_at, name, role FROM clyde_user
WHERE id = $1
`

func (q *Queries) FindUserById(ctx context.Context, id string) (ClydeUser, error) {
	row := q.db.QueryRow(ctx, findUserById, id)
	var i ClydeUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.RecoveryCode,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Role,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO clyde_user (id, email, password_hash, recovery_code, name, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, COALESCE($6, 'user'), COALESCE($7, CURRENT_TIMESTAMP), COALESCE($8, CURRENT_TIMESTAMP))
`

type InsertUserParams struct {
	ID           string
	Email        string
	PasswordHash string
	RecoveryCode string
	Name         string
	Column6      interface{}
	Column7      interface{}
	Column8      interface{}
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.Exec(ctx, insertUser,
		arg.ID,
		arg.Email,
		arg.PasswordHash,
		arg.RecoveryCode,
		arg.Name,
		arg.Column6,
		arg.Column7,
		arg.Column8,
	)
	return err
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, password_hash, recovery_code, created_at, updated_at, name, role FROM clyde_user
ORDER BY created_at ASC
`

func (q *Queries) ListUsers(ctx context.Context) ([]ClydeUser, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClydeUser
	for rows.Next() {
		var i ClydeUser
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.PasswordHash,
			&i.RecoveryCode,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE clyde_user
SET email = COALESCE($2, email),
    password_hash = COALESCE($3, password_hash),
    recovery_code = COALESCE($4, recovery_code),
    name = COALESCE($5, name),
    role = COALESCE($6, role),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateUserParams struct {
	ID           string
	Email        string
	PasswordHash string
	RecoveryCode string
	Name         string
	Role         string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.Email,
		arg.PasswordHash,
		arg.RecoveryCode,
		arg.Name,
		arg.Role,
	)
	return err
}
