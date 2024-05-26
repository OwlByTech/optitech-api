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
INSERT INTO institution_client(client_id, institution_id, vinculated_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateInstitutionClientById :exec
UPDATE institution_client
SET client_id = $2, institution_id = $3, update_at = $4
WHERE institution_client_id = $1;

-- name: DeleteInstitutionClient :exec
DELETE FROM institution_client
WHERE institution_client_id = $1;

-- name: DeleteAllInstitutionClients :execresult
DELETE FROM institution_client;
