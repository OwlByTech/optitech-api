-- name: GetFormat :one
SELECT * FROM format
WHERE format_id = $1 LIMIT 1;

-- name: ListFormats :many
SELECT * FROM format
ORDER BY format_name;

-- name: GetFormatByName :one
SELECT description, extension, version
FROM format
WHERE format_name = $1;

-- name: CreateFormat :one
INSERT INTO format(asesor_id, format_name, description, extension, version, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateFormatById :exec
UPDATE format
SET format_name = $2, description = $3, extension=$4, version=$5, updated_at=$6
WHERE format_id = $1;

-- name: DeleteFormatById :exec
UPDATE format
SET deleted_at = $2
WHERE format_id = $1;

-- name: DeleteAllFormats :execresult
UPDATE format
SET deleted_at = $1
WHERE deleted_at IS NULL;

