-- name: GetDirectoryTree :one
SELECT * FROM directory_tree
WHERE directory_id = $1 LIMIT 1;

-- name: ListDirectoryTrees :many
SELECT * FROM directory_tree
ORDER BY directory_id;

-- name: GetDirectoryChildByParent :many
SELECT sqlc.embed(dr), sqlc.embed(d)
FROM directory_tree dr
JOIN document d ON dr.directory_id= d.directory_id
WHERE dr.parent_id= $1;

-- name: GetDirectoryTreeByName :one
SELECT name
FROM directory_tree
WHERE directory_id = $1;

-- name: CreateDirectoryTree :one
INSERT INTO directory_tree(parent_id, name, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteDirectoryTreeById :exec
UPDATE directory_tree
SET deleted_at = $2
WHERE directory_id = $1;

-- name: DeleteAllDirectoryTrees :execresult
UPDATE directory_tree
SET deleted_at = $1
WHERE deleted_at IS NULL;


