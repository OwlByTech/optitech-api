-- name: GetDirectoryTree :one
SELECT * FROM directory_tree
WHERE directory_id = $1 AND deleted_at IS NULL AND institution_id=$2 LIMIT 1 ;

-- name: GetDirectoryTreeParent :one
SELECT * FROM directory_tree
WHERE parent_id IS NULL AND deleted_at IS NULL AND institution_id=$1 LIMIT 1 ;

-- name: ListDirectoryTrees :many
SELECT * FROM directory_tree
ORDER BY directory_id AND deleted_at IS NULL;

-- name: ListDirectoryChildByParent :many
SELECT *
FROM directory_tree
WHERE parent_id= $1 AND deleted_at IS NULL AND institution_id=$2;

-- name: GetDirectoryTreeByName :one
SELECT name
FROM directory_tree
WHERE directory_id = $1;

-- name: CreateDirectoryTree :one
INSERT INTO directory_tree(parent_id, name, created_at, institution_id, asesor_id)
VALUES ($1, $2, $3, $4, $5)
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
  WHERE parent_id IS null AND dt.deleted_at IS NULL AND dt.institution_id=$1
  UNION ALL
  SELECT e.directory_id, e.name, e.parent_id
  FROM directory_tree  e
  INNER JOIN directory_tree eh ON e.parent_id = eh.directory_id
    where  e.directory_id<=$2 AND e.deleted_at IS NULL  AND e.institution_id=$1
)
SELECT * FROM directory;

-- name: UpdateDirectoryTreeById :exec
UPDATE directory_tree
SET name = $2, updated_at = $3, parent_id = $4, asesor_id = $5
WHERE directory_id = $1;

-- name: GetInstitutionNameByDirectoryId :one
SELECT sqlc.embed(i), sqlc.embed(dt)
FROM directory_tree dt
JOIN institution i ON dt.institution_id = i.institution_id
WHERE dt.directory_id = $1;