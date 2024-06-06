-- name: GetInstitutionService :one
SELECT * FROM institution_services
WHERE institution_services_id = $1 LIMIT 1;

-- name: ListInstitutionServices :many
SELECT * FROM institution_services
ORDER BY institution_services_id;

-- name: GetInstitutionServiceByName :one
SELECT institution_id, services_id
FROM institution_services
WHERE institution_services_id = $1;

-- name: CreateInstitutionService :one
INSERT INTO institution_services(institution_id, services_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteInstitutionServiceById :exec
UPDATE institution_services
SET deleted_at = $2
WHERE institution_services_id = $1;

-- name: DeleteAllInstitutionServices :execresult
UPDATE institution_services
SET deleted_at = $1
WHERE deleted_at IS NULL;


