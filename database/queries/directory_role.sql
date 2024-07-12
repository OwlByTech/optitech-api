-- name: GetDirectoryRole :one
SELECT * FROM directory_role
WHERE user_id = $1 LIMIT 1;

-- name: ListDirectoryRoles :many
SELECT * FROM directory_role
ORDER BY user_id;

-- name: CreateDirectoryRole :one
INSERT INTO directory_role(directory_id, user_id, status, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteDirectoryRoleById :exec
UPDATE directory_role
SET deleted_at = $3
WHERE directory_id = $1 AND user_id = $2;

-- name: DeleteAllDirectoryRoles :execresult
UPDATE directory_role
SET deleted_at = $1
WHERE deleted_at IS NULL;

-- name: UpdateDirectoryRole :exec
UPDATE directory_role
SET user_id = $2, updated_at = $3, status = $4
WHERE directory_id = $1;
