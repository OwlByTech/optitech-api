-- name: GetInstitution :one
SELECT * FROM institution
WHERE institution_id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: GetInstitutionLogo :one
SELECT logo,institution_name FROM institution
WHERE institution_id = $1;

-- name: GetInstitutionByClient :one
SELECT  i.institution_id FROM institution i
INNER JOIN institution_client ON i.institution_id=institution_client.institution_id
WHERE  institution_client.client_id = $1
AND i.deleted_at IS NULL
LIMIT 1;

-- name: ListInstitutions :many
SELECT * FROM institution
WHERE deleted_at IS NULL
ORDER BY institution_name;

-- name: GetInstitutionByName :one
SELECT  institution_name, logo, description
FROM institution
WHERE institution_name = $1;

-- name: CreateInstitution :one
INSERT INTO institution (institution_name , description, created_at,asesor_id)
VALUES ($1, $2, $3, $4)
RETURNING institution.institution_id;

-- name: UpdateInstitution :exec
UPDATE institution
SET institution_name = $2, description = $3, updated_at = $4, asesor_id = $5
WHERE institution_id = $1;

-- name: UpdateLogoInstitution :exec
UPDATE institution
SET logo = $2 ,updated_at = $3 WHERE institution_id = $1;

-- name: UpdateAsesorInstitution :exec
UPDATE institution
SET asesor_id = $2 ,updated_at = $3
WHERE institution_id = $1;

-- name: DeleteInstitution :exec
UPDATE institution
SET deleted_at = $2
WHERE institution_id = $1 AND deleted_at IS NULL;

-- name: DeleteAllInstitutions :exec
UPDATE institution
set deleted_at = $1
where deleted_at is NULL;

-- name: GetInstitutionByAsesor :many
SELECT  i.institution_id FROM institution i
INNER JOIN institution_client ON i.institution_id=institution_client.institution_id
WHERE  i.asesor_id = $1
AND i.deleted_at IS NULL
LIMIT 1;