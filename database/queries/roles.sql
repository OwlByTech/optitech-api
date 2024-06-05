-- name: GetRole :one
SELECT * FROM roles
WHERE role_id = $1 LIMIT 1;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY role_id;

-- name: GetRoleByName :one
SELECT role_name
FROM roles
WHERE role_id = $1;

-- name: CreateRole :one
INSERT INTO roles(role_name, created_at)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateRoleById :exec
UPDATE roles
SET role_name = $2, updated_at = $3
WHERE role_id = $1;

-- name: DeleteRoleById :exec
UPDATE roles
SET deleted_at = $2
WHERE role_id = $1;

-- name: DeleteAllRoles :execresult
UPDATE roles
SET deleted_at = $1
WHERE role_id IS NULL;


