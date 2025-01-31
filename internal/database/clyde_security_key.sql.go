// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: clyde_security_key.sql

package database

import (
	"context"
)

const deleteSecurityKey = `-- name: DeleteSecurityKey :exec
DELETE FROM clyde_security_key
WHERE id = $1
`

func (q *Queries) DeleteSecurityKey(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteSecurityKey, id)
	return err
}

const findSecurityKeyById = `-- name: FindSecurityKeyById :one
SELECT id, user_id, name, created_at, updated_at, cose_algorithm_id, public_key FROM clyde_security_key
WHERE id = $1
`

func (q *Queries) FindSecurityKeyById(ctx context.Context, id string) (ClydeSecurityKey, error) {
	row := q.db.QueryRow(ctx, findSecurityKeyById, id)
	var i ClydeSecurityKey
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

const insertSecurityKey = `-- name: InsertSecurityKey :exec
INSERT INTO clyde_security_key (id, user_id, name, cose_algorithm_id, public_key)
VALUES ($1, $2, $3, $4, $5)
`

type InsertSecurityKeyParams struct {
	ID              string
	UserID          string
	Name            string
	CoseAlgorithmID int32
	PublicKey       []byte
}

func (q *Queries) InsertSecurityKey(ctx context.Context, arg InsertSecurityKeyParams) error {
	_, err := q.db.Exec(ctx, insertSecurityKey,
		arg.ID,
		arg.UserID,
		arg.Name,
		arg.CoseAlgorithmID,
		arg.PublicKey,
	)
	return err
}

const listSecurityKeys = `-- name: ListSecurityKeys :many
SELECT id, user_id, name, created_at, updated_at, cose_algorithm_id, public_key FROM clyde_security_key
ORDER BY created_at ASC
`

func (q *Queries) ListSecurityKeys(ctx context.Context) ([]ClydeSecurityKey, error) {
	rows, err := q.db.Query(ctx, listSecurityKeys)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClydeSecurityKey
	for rows.Next() {
		var i ClydeSecurityKey
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

const updateSecurityKey = `-- name: UpdateSecurityKey :exec
UPDATE clyde_security_key
SET updated_at = CURRENT_TIMESTAMP,
    name = COALESCE($2, name),
    cose_algorithm_id = COALESCE($3, cose_algorithm_id),
    public_key = COALESCE($4, public_key)
WHERE id = $1
`

type UpdateSecurityKeyParams struct {
	ID              string
	Name            string
	CoseAlgorithmID int32
	PublicKey       []byte
}

func (q *Queries) UpdateSecurityKey(ctx context.Context, arg UpdateSecurityKeyParams) error {
	_, err := q.db.Exec(ctx, updateSecurityKey,
		arg.ID,
		arg.Name,
		arg.CoseAlgorithmID,
		arg.PublicKey,
	)
	return err
}
