 -- name: GetMembershipType :one
SELECT * FROM membership_type
WHERE membership_type_id = $1 LIMIT 1;

-- name: ListMembershipType :many
SELECT * FROM membership_type
ORDER BY membership_type_id;

-- name: GetMembershipTypeByName :one
SELECT membership_name, users
FROM membership_type
WHERE membership_type_id = $1;

-- name: CreateMembershipType :one
INSERT INTO membership_type(membership_name, users, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateMembershipTypeById :exec
UPDATE membership_type
SET membership_name = $2, users = $3, updated_at = $4
WHERE membership_type_id = $1;

-- name: DeleteMembershipTypeById :exec
UPDATE membership_type
SET deleted_at = $2
WHERE membership_type_id = $1;

-- name: DeleteMembershipType :execresult
UPDATE membership_type
SET deleted_at = $1
WHERE membership_type_id IS NULL;


