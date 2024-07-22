-- name: GetRole :one
SELECT * FROM roles
WHERE role_id = $1 LIMIT 1;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY role_id;

-- name: GetRoleByName :one
SELECT role_name, description
FROM roles
WHERE role_id = $1;

-- name: CreateRole :one
INSERT INTO roles(role_name, description, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateRoleById :exec
UPDATE roles
SET role_name = $2, description = $3, updated_at = $4
WHERE role_id = $1;

-- name: DeleteRoleById :exec
UPDATE roles
SET deleted_at = $2
WHERE role_id = $1;

-- name: DeleteAllRoles :execresult
UPDATE roles
SET deleted_at = $1
WHERE deleted_at IS NULL;


