-- name: GetDirectoryInstitution :one
SELECT * FROM directory_institution
WHERE directory_institution_id = $1 LIMIT 1;

-- name: ListDirectoryInstitutions :many
SELECT * FROM directory_institution
ORDER BY directory_institution_id;

-- name: GetDirectoryInstitutionByName :one
SELECT institution_id, directory_id
FROM directory_institution
WHERE directory_institution_id = $1;

-- name: CreateDirectoryInstitution :one
INSERT INTO directory_institution(institution_id, directory_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteDirectoryInstitutionById :exec
UPDATE directory_institution
SET deleted_at = $2
WHERE directory_institution_id = $1;

-- name: DeleteAllDirectoryInstitutions :execresult
UPDATE directory_institution
SET deleted_at = $1
WHERE deleted_at IS NULL;

