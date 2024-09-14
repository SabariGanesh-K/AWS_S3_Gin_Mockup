-- name: CreateFile :one
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
) RETURNING *;

-- name: SearchFile :one
-- type: SearchFileParams:MinCreatedAt=time.Time,SearchFileParams:MaxCreatedAt=time.Time
SELECT *
FROM files
WHERE (
    id = COALESCE(sqlc.arg(id), id) OR sqlc.arg(id) IS NULL

AND file_name ILIKE COALESCE(sqlc.arg(file_name), '%')
AND (filetype = COALESCE(sqlc.arg(filetype), filetype) OR sqlc.arg(filetype) IS NULL)
AND ( sqlc.arg(min_created_at) IS NULL OR  sqlc.arg(max_created_at) IS NULL) OR (created_at BETWEEN COALESCE(sqlc.arg(min_created_at), '0001-01-01'::timestamp) AND COALESCE(sqlc.arg(max_created_at), '9999-12-31'::timestamp))
) LIMIT 1;




