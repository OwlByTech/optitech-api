-- name: GetInstitution :one
SELECT * FROM institution
WHERE institution_id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: ListInstitutions :many
SELECT * FROM institution
WHERE deleted_at IS NULL
ORDER BY institution_name ;

-- name: GetInstitutionByName :one
SELECT  institution_name, logo, description
FROM institution
WHERE institution_name = $1;

-- name: CreateInstitution :one
INSERT INTO institution (institution_name, logo, description, created_at,asesor_id)
VALUES ($1, $2, $3, $4,$5)
RETURNING *;

-- name: UpdateInstitution :exec
UPDATE institution
SET institution_name = $2, logo = $3, description = $4,  updated_at=$5,asesor_id= $6
WHERE institution_id = $1;


-- name: UpdateAsesorInstitution :exec
UPDATE institution
SET asesor_id= $2 ,updated_at=$3
WHERE institution_id = $1;

-- name: DeleteInstitution :exec
UPDATE institution
SET deleted_at = $2
WHERE institution_id = $1 AND deleted_at IS NULL;

-- name: DeleteAllInstitutions :exec
UPDATE institution
set deleted_at = $1
where deleted_at is NULL;
