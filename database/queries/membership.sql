 -- name: GetMembership :one
SELECT * FROM membership
WHERE membership_id = $1 LIMIT 1;

-- name: ListMemberships :many
SELECT * FROM membership
ORDER BY membership_id;

-- name: GetMembershipByName :one
SELECT membership_type_id, finish_at
FROM membership
WHERE membership_id = $1;

-- name: CreateMembership :one
INSERT INTO membership(membership_type_id, created_at, finish_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateMembershipById :exec
UPDATE membership
SET membership_type_id = $2, finish_at = $3, updated_at = $4
WHERE membership_id = $1;

-- name: DeleteMembershipById :exec
UPDATE membership
SET deleted_at = $2
WHERE membership_id = $1;

-- name: DeleteRoleAllPermissions :execresult
UPDATE membership
SET deleted_at = $1
WHERE deleted_at IS NULL;


