// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: clyde_user_totp_credential.sql

package database

import (
	"context"
)

const deleteTotpCredential = `-- name: DeleteTotpCredential :exec
DELETE FROM clyde_user_totp_credential
WHERE user_id = $1
`

func (q *Queries) DeleteTotpCredential(ctx context.Context, userID string) error {
	_, err := q.db.Exec(ctx, deleteTotpCredential, userID)
	return err
}

const findTotpCredentialByUserId = `-- name: FindTotpCredentialByUserId :one
SELECT user_id, created_at, updated_at, key FROM clyde_user_totp_credential
WHERE user_id = $1
`

func (q *Queries) FindTotpCredentialByUserId(ctx context.Context, userID string) (ClydeUserTotpCredential, error) {
	row := q.db.QueryRow(ctx, findTotpCredentialByUserId, userID)
	var i ClydeUserTotpCredential
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Key,
	)
	return i, err
}

const insertTotpCredential = `-- name: InsertTotpCredential :exec
INSERT INTO clyde_user_totp_credential (user_id, key)
VALUES ($1, $2)
`

type InsertTotpCredentialParams struct {
	UserID string
	Key    []byte
}

func (q *Queries) InsertTotpCredential(ctx context.Context, arg InsertTotpCredentialParams) error {
	_, err := q.db.Exec(ctx, insertTotpCredential, arg.UserID, arg.Key)
	return err
}

const listTotpCredentials = `-- name: ListTotpCredentials :many
SELECT user_id, created_at, updated_at, key FROM clyde_user_totp_credential
ORDER BY created_at ASC
`

func (q *Queries) ListTotpCredentials(ctx context.Context) ([]ClydeUserTotpCredential, error) {
	rows, err := q.db.Query(ctx, listTotpCredentials)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClydeUserTotpCredential
	for rows.Next() {
		var i ClydeUserTotpCredential
		if err := rows.Scan(
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Key,
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

const updateTotpCredential = `-- name: UpdateTotpCredential :exec
UPDATE clyde_user_totp_credential
SET updated_at = CURRENT_TIMESTAMP,
    key = COALESCE($2, key)
WHERE user_id = $1
`

type UpdateTotpCredentialParams struct {
	UserID string
	Key    []byte
}

func (q *Queries) UpdateTotpCredential(ctx context.Context, arg UpdateTotpCredentialParams) error {
	_, err := q.db.Exec(ctx, updateTotpCredential, arg.UserID, arg.Key)
	return err
}
