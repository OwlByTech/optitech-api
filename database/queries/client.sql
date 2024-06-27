-- name: GetClient :one
SELECT * FROM client
WHERE client_id = $1 LIMIT 1;

-- name: LoginClient :one
SELECT * FROM client
WHERE password = $1 AND email= $2 LIMIT 1;

-- name: ListClients :many
SELECT * FROM client
ORDER BY given_name;

-- name: GetClientByEmail :one
SELECT * FROM client
WHERE email = $1;

-- name: CreateClient :one
INSERT INTO client (given_name, surname, email, password, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateClientById :exec
UPDATE client
SET  given_name = $2, password = $3, surname = $4, email = $5, updated_at = $6
WHERE client_id = $1;

-- name: UpdateClientPhoto :exec
UPDATE client
SET  photo= $2, updated_at = $3
WHERE client_id = $1;

-- name: UpdateClientStatusById :exec
UPDATE client
SET status= $2, updated_at = $3
WHERE client_id = $1;

-- name: DeleteClientById :exec
UPDATE client
SET deleted_at = $2
WHERE client_id = $1;

-- name: DeleteAllClients :execresult
UPDATE client
SET deleted_at = $1
WHERE deleted_at IS NULL;
