-- name: GetInstitution :one
SELECT * FROM institution
WHERE institution_id = $1 LIMIT 1;

-- name: ListInstitutions :many
SELECT * FROM institution
ORDER BY institution_name;

-- name: GetInstitutionByName :one
SELECT  institution_name, logo, description, services
FROM institution
WHERE institution_name = $1;

-- name: CreateInstitution :one
INSERT INTO institution (institution_name, logo, description, services, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateInstitutionById :exec
UPDATE institution
SET institution_name = $2, logo = $3, description = $4, services = $5, updated_at=$6
WHERE institution_id = $1;

-- name: DeleteInstitutionById :exec
UPDATE institution
SET deleted_at = $2
WHERE institution_id = $1 AND deleted_at IS NULL;

-- name: DeleteAllInstitutions :exec
UPDATE institution
set deleted_at = $1 
where deleted_at is NULL;