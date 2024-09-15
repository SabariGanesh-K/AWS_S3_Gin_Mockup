// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddFile(ctx context.Context, arg AddFileParams) (Users, error)
	CreateFile(ctx context.Context, arg CreateFileParams) (Files, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Sessions, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmails, error)
	GetSession(ctx context.Context, id uuid.UUID) (Sessions, error)
	GetUser(ctx context.Context, username string) (Users, error)
	// type: SearchFileParams:MinCreatedAt=time.Time,SearchFileParams:MaxCreatedAt=time.Time
	SearchFile(ctx context.Context, arg SearchFileParams) (Files, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmails, error)
}

var _ Querier = (*Queries)(nil)
