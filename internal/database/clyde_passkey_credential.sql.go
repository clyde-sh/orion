// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: clyde_passkey_credential.sql

package database

import (
	"context"
)

const deletePasskeyCredential = `-- name: DeletePasskeyCredential :exec
DELETE FROM clyde_passkey_credential
WHERE id = $1
`

func (q *Queries) DeletePasskeyCredential(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deletePasskeyCredential, id)
	return err
}

const findPasskeyCredentialById = `-- name: FindPasskeyCredentialById :one
SELECT id, user_id, name, created_at, updated_at, cose_algorithm_id, public_key FROM clyde_passkey_credential
WHERE id = $1
`

func (q *Queries) FindPasskeyCredentialById(ctx context.Context, id string) (ClydePasskeyCredential, error) {
	row := q.db.QueryRow(ctx, findPasskeyCredentialById, id)
	var i ClydePasskeyCredential
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CoseAlgorithmID,
		&i.PublicKey,
	)
	return i, err
}

const insertPasskeyCredential = `-- name: InsertPasskeyCredential :exec
INSERT INTO clyde_passkey_credential (id, user_id, name, cose_algorithm_id, public_key)
VALUES ($1, $2, $3, $4, $5)
`

type InsertPasskeyCredentialParams struct {
	ID              string
	UserID          string
	Name            string
	CoseAlgorithmID int32
	PublicKey       []byte
}

func (q *Queries) InsertPasskeyCredential(ctx context.Context, arg InsertPasskeyCredentialParams) error {
	_, err := q.db.Exec(ctx, insertPasskeyCredential,
		arg.ID,
		arg.UserID,
		arg.Name,
		arg.CoseAlgorithmID,
		arg.PublicKey,
	)
	return err
}

const listPasskeyCredentials = `-- name: ListPasskeyCredentials :many
SELECT id, user_id, name, created_at, updated_at, cose_algorithm_id, public_key FROM clyde_passkey_credential
ORDER BY created_at ASC
`

func (q *Queries) ListPasskeyCredentials(ctx context.Context) ([]ClydePasskeyCredential, error) {
	rows, err := q.db.Query(ctx, listPasskeyCredentials)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClydePasskeyCredential
	for rows.Next() {
		var i ClydePasskeyCredential
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CoseAlgorithmID,
			&i.PublicKey,
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

const updatePasskeyCredential = `-- name: UpdatePasskeyCredential :exec
UPDATE clyde_passkey_credential
SET updated_at = CURRENT_TIMESTAMP,
    name = COALESCE($2, name),
    cose_algorithm_id = COALESCE($3, cose_algorithm_id),
    public_key = COALESCE($4, public_key)
WHERE id = $1
`

type UpdatePasskeyCredentialParams struct {
	ID              string
	Name            string
	CoseAlgorithmID int32
	PublicKey       []byte
}

func (q *Queries) UpdatePasskeyCredential(ctx context.Context, arg UpdatePasskeyCredentialParams) error {
	_, err := q.db.Exec(ctx, updatePasskeyCredential,
		arg.ID,
		arg.Name,
		arg.CoseAlgorithmID,
		arg.PublicKey,
	)
	return err
}
