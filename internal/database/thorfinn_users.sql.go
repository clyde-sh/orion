// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: thorfinn_users.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO thorfinn_users (id, email, password_hash) VALUES ($1, $2, $3) RETURNING id, email, password_hash, verified, two_factor_enabled, created_at, updated_at
`

type CreateUserParams struct {
	ID           string
	Email        string
	PasswordHash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (ThorfinnUser, error) {
	row := q.db.QueryRow(ctx, createUser, arg.ID, arg.Email, arg.PasswordHash)
	var i ThorfinnUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Verified,
		&i.TwoFactorEnabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM thorfinn_users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, email, password_hash, verified, two_factor_enabled, created_at, updated_at FROM thorfinn_users WHERE email = $1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (ThorfinnUser, error) {
	row := q.db.QueryRow(ctx, findUserByEmail, email)
	var i ThorfinnUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Verified,
		&i.TwoFactorEnabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserById = `-- name: FindUserById :one
SELECT id, email, password_hash, verified, two_factor_enabled, created_at, updated_at FROM thorfinn_users WHERE id = $1
`

func (q *Queries) FindUserById(ctx context.Context, id string) (ThorfinnUser, error) {
	row := q.db.QueryRow(ctx, findUserById, id)
	var i ThorfinnUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Verified,
		&i.TwoFactorEnabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, verified, two_factor_enabled, created_at, updated_at
FROM thorfinn_users
ORDER BY created_at DESC
`

type ListUsersRow struct {
	ID               string
	Email            string
	Verified         bool
	TwoFactorEnabled bool
	CreatedAt        pgtype.Timestamptz
	UpdatedAt        pgtype.Timestamptz
}

func (q *Queries) ListUsers(ctx context.Context) ([]ListUsersRow, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersRow
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Verified,
			&i.TwoFactorEnabled,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateUser = `-- name: UpdateUser :one
UPDATE thorfinn_users
SET email = $2, password_hash = $3, verified = $4, two_factor_enabled = $5
WHERE id = $1
RETURNING id, email, password_hash, verified, two_factor_enabled, created_at, updated_at
`

type UpdateUserParams struct {
	ID               string
	Email            string
	PasswordHash     string
	Verified         bool
	TwoFactorEnabled bool
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (ThorfinnUser, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Email,
		arg.PasswordHash,
		arg.Verified,
		arg.TwoFactorEnabled,
	)
	var i ThorfinnUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Verified,
		&i.TwoFactorEnabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :one
UPDATE thorfinn_users
SET password_hash = $2
WHERE id = $1
RETURNING id, email, password_hash, verified, two_factor_enabled, created_at, updated_at
`

type UpdateUserPasswordParams struct {
	ID           string
	PasswordHash string
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (ThorfinnUser, error) {
	row := q.db.QueryRow(ctx, updateUserPassword, arg.ID, arg.PasswordHash)
	var i ThorfinnUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Verified,
		&i.TwoFactorEnabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserVerified = `-- name: UpdateUserVerified :one
UPDATE thorfinn_users
SET verified = $2
WHERE id = $1
RETURNING id, email, password_hash, verified, two_factor_enabled, created_at, updated_at
`

type UpdateUserVerifiedParams struct {
	ID       string
	Verified bool
}

func (q *Queries) UpdateUserVerified(ctx context.Context, arg UpdateUserVerifiedParams) (ThorfinnUser, error) {
	row := q.db.QueryRow(ctx, updateUserVerified, arg.ID, arg.Verified)
	var i ThorfinnUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Verified,
		&i.TwoFactorEnabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
