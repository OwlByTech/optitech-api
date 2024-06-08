-- name: ListInstitutionServices :many
SELECT services.service_name ,services.service_id FROM institution_services
INNER JOIN services ON  institution_services.services_id=services.service_id
WHERE institution_services.institution_id= $1
ORDER BY services.service_id;

-- name: CreateInstitutionService :copyfrom
INSERT INTO institution_services(institution_id, service_id, created_at)
VALUES ($1, $2, $3);

-- name: DeleteAllInstitutionServices :execresult
UPDATE institution_services
SET deleted_at = $1
WHERE deleted_at IS NULL;


