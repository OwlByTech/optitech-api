-- name: GetInstitutionClient :one
SELECT * FROM institution_client
WHERE institution_client_id = $1 LIMIT 1;

-- name: ListInstitutionClients :many
SELECT * FROM institution_client
ORDER BY institution_client_id;

-- name: GetInstitutionClientByName :one
SELECT client_id, institution_id
FROM institution_client
WHERE institution_client_id = $1;

-- name: CreateInstitutionClient :one
INSERT INTO institution_client(client_id, institution_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateInstitutionClientById :exec
UPDATE institution_client
SET client_id = $2, institution_id = $3, updated_at = $4
WHERE institution_client_id = $1;

-- name: DeleteinstInstitutionClientById :exec
UPDATE institution_client
SET deleted_at = $2
WHERE institution_client_id = $1;

-- name: DeleteAllInstitutionClient :execresult
UPDATE institution_client
SET deleted_at = $1
WHERE institution_client_id IS NULL;
