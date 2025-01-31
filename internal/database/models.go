// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ClydeEmailUpdateRequest struct {
	ID        string
	UserID    string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	ExpiresAt pgtype.Timestamptz
	Email     string
	Code      string
}

type ClydeEmailVerificationRequest struct {
	UserID    string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	ExpiresAt pgtype.Timestamptz
	Code      string
}

type ClydePasskeyCredential struct {
	ID              string
	UserID          string
	Name            string
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
	CoseAlgorithmID int32
	PublicKey       []byte
}

type ClydePasswordResetRequest struct {
	ID        string
	UserID    string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	ExpiresAt pgtype.Timestamptz
	CodeHash  string
}

type ClydeSecurityKey struct {
	ID              string
	UserID          string
	Name            string
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
	CoseAlgorithmID int32
	PublicKey       []byte
}

type ClydeUser struct {
	ID              string
	Email           string
	PasswordHash    string
	RecoveryCode    string
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
	Name            string
	Role            string
	IsEmailVerified pgtype.Bool
	Is2faEnabled    pgtype.Bool
}

type ClydeUserTotpCredential struct {
	UserID    string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	Key       []byte
}
