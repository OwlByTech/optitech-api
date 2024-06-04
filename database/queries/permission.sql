-- name: GetPermission :one
SELECT * FROM permission
WHERE permission_id = $1 LIMIT 1;

-- name: ListPermissions :many
SELECT * FROM permission
ORDER BY permission_id;

-- name: GetPermissionByName :one
SELECT permission_type
FROM permission
WHERE permission_id = $1;

-- name: CreatePermission :one
INSERT INTO permission(permission_type, created_at)
VALUES ($1, $2)
RETURNING *;

-- name: UpdatePermissionById :exec
UPDATE permission
SET permission_type = $2, updated_at = $3
WHERE permission_id = $1;

-- name: DeletePermissionById :exec
UPDATE permission
SET deleted_at = $2
WHERE permission_id = $1;

-- name: DeleteAllPermissions :execresult
UPDATE permission
SET deleted_at = $1
WHERE document_client_id IS NULL;


