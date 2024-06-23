 -- name: GetRolePermission :one
SELECT * FROM role_permission
WHERE role_permission_id = $1 LIMIT 1;

-- name: ListRolePermissions :many
SELECT * FROM role_permission
ORDER BY role_permission_id;

-- name: GetRolePermission :one
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

-- name: DeleteAllRolePermissions :execresult
UPDATE role_permission
SET deleted_at = $1
WHERE deleted_at IS NULL;

-- name: ListPermissionByRoleId :many
SELECT sqlc.embed(p), sqlc.embed(r), sqlc.embed(rp)
FROM role_permission rp
JOIN permission p ON rp.permission_id = p.permission_id
JOIN roles r ON rp.role_id = r.role_id
WHERE rp.role_id = $1;