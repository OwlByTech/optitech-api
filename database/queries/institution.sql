-- name: GetInstitution :one
SELECT * FROM institution
WHERE institution_id = $1 LIMIT 1;

-- name: ListInstitutions :many
SELECT * FROM institution
ORDER BY institution_name;

-- name: GetInstitutionByName :one
SELECT  institution_name, logo, descrip, services
FROM institution
WHERE institution_name = $1;

-- name: CreateInstitution :one
INSERT INTO institution (institution_name, logo, descrip, services, create_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateInstitutionById :exec
UPDATE institution
SET institution_name = $2, logo = $3, descrip = $4, services = $5, update_at=$6
WHERE institution_id = $1;

-- name: DeleteInstitution :exec
DELETE FROM institution
WHERE institution_id = $1;

-- name: DeleteAllInstitutions :execresult
DELETE FROM institution;
