-- name: ListInstitutionServices :many
SELECT services.service_name ,services.service_id FROM institution_services
INNER JOIN services ON  institution_services.service_id=services.service_id
WHERE institution_services.institution_id= $1
AND institution_services.deleted_at IS NULL
ORDER BY services.service_id;

-- name: CreateInstitutionServices :copyfrom
INSERT INTO institution_services(institution_id, service_id, created_at)
VALUES ($1, $2, $3);

-- name: ExistsInstitutionService :one
SELECT * FROM institution_services
WHERE service_id = $1 AND institution_id=$2 and deleted_at IS NOT NULL ;


-- name: RecoverInstitutionService :exec
UPDATE institution_services
SET deleted_at = NULL,updated_at= $2
WHERE institution_id= $1 AND service_id= $3 ;

-- name: DeleteAllInstitutionServices :execresult
UPDATE institution_services
SET deleted_at = $1
WHERE deleted_at IS NULL;

-- name: DeleteInstitutionServiceById :exec
UPDATE institution_services
SET deleted_at = $2
WHERE institution_id= $1 AND service_id= $3;

-- name: DeleteInstitutionServicesByInstitution :exec
UPDATE institution_services
SET deleted_at = $2
WHERE institution_id= $1;
