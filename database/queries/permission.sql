-- name: GetPermission :one
SELECT * FROM permission
WHERE permission_id = $1 LIMIT 1;

-- name: ListPermissions :many
SELECT * FROM permission
ORDER BY permission_id;

-- name: GetPermissionByName :one
SELECT permission_name, permission_code, permission_description
FROM permission
WHERE permission_id = $1;

-- name: CreatePermission :one
INSERT INTO permission(permission_name, permission_code, permission_description, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdatePermissionById :exec
UPDATE permission
SET permission_name = $2, permission_code = $3, permission_description = $4, updated_at = $5
WHERE permission_id = $1;

-- name: DeletePermissionById :exec
UPDATE permission
SET deleted_at = $2
WHERE permission_id = $1;

-- name: DeleteAllPermissions :execresult
UPDATE permission
SET deleted_at = $1
WHERE deleted_at IS NULL;


