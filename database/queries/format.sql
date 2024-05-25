-- name: GetFormat :one
SELECT * FROM format
WHERE format_id = $1 LIMIT 1;

-- name: ListFormats :many
SELECT * FROM format
ORDER BY format_name;

-- name: GetFormatByName :one
SELECT description, items, extension, version
FROM format
WHERE format_name = $1;

-- name: CreateFormat :one
INSERT INTO format(asesor_id, format_name, description, items, extension, version, create_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateFormatById :exec
UPDATE format
SET format_name = $2, description = $3, items = $4, extension=$5, version=$6, update_at=$7
WHERE format_id = $1;

-- name: DeleteFormat :exec
DELETE FROM format
WHERE format_id = $1;

-- name: DeleteAllFormats :execresult
DELETE FROM format;
