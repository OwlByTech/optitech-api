-- name: GetStandard :one
SELECT * FROM standards
WHERE standard_id = $1 LIMIT 1;

-- name: ListStandards :many
SELECT * FROM standards
ORDER BY standard_id;

-- name: GetStandardByName :one
SELECT standard, complexity, modality, article, section, paragraph, criteria, comply, applys
FROM standards
WHERE standard_id = $1;

-- name: CreateStandard :one
INSERT INTO standards(service_id, standard, complexity, modality, article, section, paragraph, criteria, comply, applys, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: UpdateStandardById :exec
UPDATE standards
SET deleted_at = $2
WHERE standard_id = $1;

-- name: DeleteStandardById :exec
UPDATE standards
SET standard = $2, complexity = $3, modality =$4, article = $5, section = $6, paragraph = $7, criteria = $8, comply = $9, applys = $10, updated_at = $11
WHERE standard_id = $1;

-- name: DeleteAllStandards :execresult
UPDATE standards
SET deleted_at = $1
WHERE deleted_at IS NULL;


