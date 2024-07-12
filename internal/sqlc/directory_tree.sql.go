// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: directory_tree.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createDirectoryTree = `-- name: CreateDirectoryTree :one
INSERT INTO directory_tree(parent_id, name, created_at)
VALUES ($1, $2, $3)
RETURNING directory_id, parent_id, name, created_at, updated_at, deleted_at
`

type CreateDirectoryTreeParams struct {
	ParentID  pgtype.Int4      `json:"parent_id"`
	Name      pgtype.Text      `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateDirectoryTree(ctx context.Context, arg CreateDirectoryTreeParams) (DirectoryTree, error) {
	row := q.db.QueryRow(ctx, createDirectoryTree, arg.ParentID, arg.Name, arg.CreatedAt)
	var i DirectoryTree
	err := row.Scan(
		&i.DirectoryID,
		&i.ParentID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllDirectoryTrees = `-- name: DeleteAllDirectoryTrees :execresult
UPDATE directory_tree
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllDirectoryTrees(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllDirectoryTrees, deletedAt)
}

const deleteDirectoryTreeById = `-- name: DeleteDirectoryTreeById :exec
UPDATE directory_tree
SET deleted_at = $2
WHERE directory_id = $1
`

type DeleteDirectoryTreeByIdParams struct {
	DirectoryID int64            `json:"directory_id"`
	DeletedAt   pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteDirectoryTreeById(ctx context.Context, arg DeleteDirectoryTreeByIdParams) error {
	_, err := q.db.Exec(ctx, deleteDirectoryTreeById, arg.DirectoryID, arg.DeletedAt)
	return err
}

const getDirectoryTree = `-- name: GetDirectoryTree :one
SELECT directory_id, parent_id, name, created_at, updated_at, deleted_at FROM directory_tree
WHERE directory_id = $1 LIMIT 1
`

func (q *Queries) GetDirectoryTree(ctx context.Context, directoryID int64) (DirectoryTree, error) {
	row := q.db.QueryRow(ctx, getDirectoryTree, directoryID)
	var i DirectoryTree
	err := row.Scan(
		&i.DirectoryID,
		&i.ParentID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getDirectoryTreeByName = `-- name: GetDirectoryTreeByName :one
SELECT name
FROM directory_tree
WHERE directory_id = $1
`

func (q *Queries) GetDirectoryTreeByName(ctx context.Context, directoryID int64) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, getDirectoryTreeByName, directoryID)
	var name pgtype.Text
	err := row.Scan(&name)
	return name, err
}

const listDirectoryChildByParent = `-- name: ListDirectoryChildByParent :many
SELECT directory_id, parent_id, name, created_at, updated_at, deleted_at
FROM directory_tree
WHERE parent_id= $1
`

func (q *Queries) ListDirectoryChildByParent(ctx context.Context, parentID pgtype.Int4) ([]DirectoryTree, error) {
	rows, err := q.db.Query(ctx, listDirectoryChildByParent, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DirectoryTree
	for rows.Next() {
		var i DirectoryTree
		if err := rows.Scan(
			&i.DirectoryID,
			&i.ParentID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listDirectoryHierarchyById = `-- name: ListDirectoryHierarchyById :many
WITH RECURSIVE directory  AS (
  SELECT directory_id,name,parent_id
  FROM directory_tree dt
  WHERE parent_id IS null
  UNION ALL
  SELECT e.directory_id, e.name, e.parent_id
  FROM directory_tree  e
  INNER JOIN directory_tree eh ON e.parent_id = eh.directory_id where  e.directory_id<=$1
)
SELECT directory_id, name, parent_id FROM directory
`

type ListDirectoryHierarchyByIdRow struct {
	DirectoryID int64       `json:"directory_id"`
	Name        pgtype.Text `json:"name"`
	ParentID    pgtype.Int4 `json:"parent_id"`
}

func (q *Queries) ListDirectoryHierarchyById(ctx context.Context, directoryID int64) ([]ListDirectoryHierarchyByIdRow, error) {
	rows, err := q.db.Query(ctx, listDirectoryHierarchyById, directoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListDirectoryHierarchyByIdRow
	for rows.Next() {
		var i ListDirectoryHierarchyByIdRow
		if err := rows.Scan(&i.DirectoryID, &i.Name, &i.ParentID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listDirectoryTrees = `-- name: ListDirectoryTrees :many
SELECT directory_id, parent_id, name, created_at, updated_at, deleted_at FROM directory_tree
ORDER BY directory_id
`

func (q *Queries) ListDirectoryTrees(ctx context.Context) ([]DirectoryTree, error) {
	rows, err := q.db.Query(ctx, listDirectoryTrees)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DirectoryTree
	for rows.Next() {
		var i DirectoryTree
		if err := rows.Scan(
			&i.DirectoryID,
			&i.ParentID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
