// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Files struct {
	ID        string    `json:"id"`
	Owner     string    `json:"owner"`
	S3Url     string    `json:"s3_url"`
	FileName  string    `json:"file_name"`
	CreatedAt time.Time `json:"created_at"`
	Filesize  string    `json:"filesize"`
	Filetype  string    `json:"filetype"`
}

type Sessions struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Users struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	FilesOwned        []string  `json:"files_owned"`
	IsEmailVerified   bool      `json:"is_email_verified"`
	Role              string    `json:"role"`
}

type VerifyEmails struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	SecretCode string    `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}
