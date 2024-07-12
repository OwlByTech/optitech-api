-- name: GetDirectoryTree :one
SELECT * FROM directory_tree
WHERE directory_id = $1 LIMIT 1;

-- name: ListDirectoryTrees :many
SELECT * FROM directory_tree
ORDER BY directory_id;

-- name: ListDirectoryChildByParent :many
SELECT *
FROM directory_tree
WHERE parent_id= $1;


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


-- name: ListDirectoryHierarchyById :many
WITH RECURSIVE directory  AS (
  SELECT directory_id,name,parent_id
  FROM directory_tree dt
  WHERE parent_id IS null
  UNION ALL
  SELECT e.directory_id, e.name, e.parent_id
  FROM directory_tree  e
  INNER JOIN directory_tree eh ON e.parent_id = eh.directory_id where  e.directory_id<=$1
)
SELECT * FROM directory;
