-- name: GetDirectoryRole :one
SELECT * FROM directory_role
WHERE directory_role_id = $1 LIMIT 1;

-- name: ListDirectoryRoles :many
SELECT * FROM directory_role
ORDER BY directory_role_id;

-- name: GetDirectoryRoleByName :one
SELECT directory_id, role_id
FROM directory_role
WHERE directory_role_id = $1;

-- name: CreateDirectoryRole :one
INSERT INTO directory_role(directory_id, role_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteDirectoryRoleById :exec
UPDATE directory_role
SET deleted_at = $2
WHERE directory_role_id = $1;

-- name: DeleteAllDirectoryRoles :execresult
UPDATE directory_role
SET deleted_at = $1
WHERE directory_id IS NULL;


