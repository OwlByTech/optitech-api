 -- name: GetRolePermissionn :one
SELECT * FROM role_permission
WHERE role_permission_id = $1 LIMIT 1;

-- name: ListPeRolermissions :many
SELECT * FROM role_permission
ORDER BY role_permission_id;

-- name: GetRolePermissionnByName :one
SELECT role_id, permission_id
FROM role_permission
WHERE role_permission_id = $1;

-- name: CreateRolePermission :one
INSERT INTO role_permission(role_id, permission_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateRolePermissionById :exec
UPDATE role_permission
SET role_id = $2, permission_id = $3, updated_at = $4
WHERE role_permission_id = $1;

-- name: DeleteRolePermissionById :exec
UPDATE role_permission
SET deleted_at = $2
WHERE role_permission_id = $1;

-- name: DeleteRoleAllPermissions :execresult
UPDATE role_permission
SET deleted_at = $1
WHERE document_client_id IS NULL;


