-- name: GetClientRole :one
SELECT * FROM client_role
WHERE client_role_id = $1 LIMIT 1;

-- name: ListClientRoles :many
SELECT * FROM client_role
ORDER BY role_id;

-- name: GetClientRoleByName :one
SELECT client_id, role_id
FROM client_role
WHERE client_role_id = $1;

-- name: CreateClientRole :one
INSERT INTO client_role(client_id, role_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateClientRoleById :exec
UPDATE client_role
SET client_id = $2, role_id = $3, updated_at = $4
WHERE client_role_id = $1;

-- name: DeleteClientRoleById :exec
UPDATE client_role
SET deleted_at = $2
WHERE client_role_id = $1;

-- name: DeleteAllClientRoles :execresult
UPDATE client_role
SET deleted_at = $1
WHERE deleted_at IS NULL;

-- name: GetClientRoleByClientId :many
SELECT sqlc.embed(c), sqlc.embed(r), sqlc.embed(cr)
FROM client_role cr
JOIN client c ON cr.client_id = c.client_id
JOIN roles r ON cr.role_id = r.role_id
WHERE cr.client_id = $1;
