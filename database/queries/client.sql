-- name: GetClient :one
SELECT * FROM client
WHERE client_id = $1 LIMIT 1;

-- name: ListClients :many
SELECT * FROM client
ORDER BY given_name;

-- name: GetClientByEmail :one
SELECT given_name, surname, email, pass, asesor_id, institution_id
FROM client
WHERE email = $1;

-- name: CreateClient :one
INSERT INTO client (given_name, surname, email, pass, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateClientById :exec
UPDATE client
SET given_name = $2, pass = $3, surname = $4, email = $5, asesor_id=$6, institution_id=$7, updated_at = $8
WHERE client_id = $1;

-- name: DeleteClient :exec
DELETE FROM client
WHERE client_id = $1;

-- name: DeleteAllClients :execresult
DELETE FROM client;
