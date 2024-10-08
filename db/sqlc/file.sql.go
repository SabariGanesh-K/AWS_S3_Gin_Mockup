// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: file.sql

package db

import (
	"context"
	"time"
)

const createFile = `-- name: CreateFile :one
INSERT INTO files (
  id,
  owner,
  s3_url,
  file_name,
   created_at,
   filesize,
   filetype
) VALUES (
  $1, $2, $3, $4,$5,$6,$7
) RETURNING id, owner, s3_url, file_name, created_at, filesize, filetype
`

type CreateFileParams struct {
	ID        string    `json:"id"`
	Owner     string    `json:"owner"`
	S3Url     string    `json:"s3_url"`
	FileName  string    `json:"file_name"`
	CreatedAt time.Time `json:"created_at"`
	Filesize  string    `json:"filesize"`
	Filetype  string    `json:"filetype"`
}

func (q *Queries) CreateFile(ctx context.Context, arg CreateFileParams) (Files, error) {
	row := q.db.QueryRowContext(ctx, createFile,
		arg.ID,
		arg.Owner,
		arg.S3Url,
		arg.FileName,
		arg.CreatedAt,
		arg.Filesize,
		arg.Filetype,
	)
	var i Files
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.S3Url,
		&i.FileName,
		&i.CreatedAt,
		&i.Filesize,
		&i.Filetype,
	)
	return i, err
}

const searchFile = `-- name: SearchFile :one
SELECT id, owner, s3_url, file_name, created_at, filesize, filetype
FROM files
WHERE (
    id = COALESCE($1, id) OR $1 IS NULL

AND file_name ILIKE COALESCE($2, '%')
AND (filetype = COALESCE($3, filetype) OR $3 IS NULL)
AND ( $4 IS NULL OR  $5 IS NULL) OR (created_at BETWEEN COALESCE($4, '0001-01-01'::timestamp) AND COALESCE($5, '9999-12-31'::timestamp))
) LIMIT 1
`

type SearchFileParams struct {
	ID           string      `json:"id"`
	FileName     string      `json:"file_name"`
	Filetype     string      `json:"filetype"`
	MinCreatedAt interface{} `json:"min_created_at"`
	MaxCreatedAt interface{} `json:"max_created_at"`
}

// type: SearchFileParams:MinCreatedAt=time.Time,SearchFileParams:MaxCreatedAt=time.Time
func (q *Queries) SearchFile(ctx context.Context, arg SearchFileParams) (Files, error) {
	row := q.db.QueryRowContext(ctx, searchFile,
		arg.ID,
		arg.FileName,
		arg.Filetype,
		arg.MinCreatedAt,
		arg.MaxCreatedAt,
	)
	var i Files
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.S3Url,
		&i.FileName,
		&i.CreatedAt,
		&i.Filesize,
		&i.Filetype,
	)
	return i, err
}
