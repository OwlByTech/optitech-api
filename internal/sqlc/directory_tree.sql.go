// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: directory_tree.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createDirectoryTree = `-- name: CreateDirectoryTree :one
INSERT INTO directory_tree(parent_id, name, created_at, institution_id, asesor_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING directory_id, parent_id, institution_id, name, asesor_id, created_at, updated_at, deleted_at
`

type CreateDirectoryTreeParams struct {
	ParentID      pgtype.Int8      `json:"parent_id"`
	Name          pgtype.Text      `json:"name"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	InstitutionID pgtype.Int4      `json:"institution_id"`
	AsesorID      pgtype.Int4      `json:"asesor_id"`
}

func (q *Queries) CreateDirectoryTree(ctx context.Context, arg CreateDirectoryTreeParams) (DirectoryTree, error) {
	row := q.db.QueryRow(ctx, createDirectoryTree,
		arg.ParentID,
		arg.Name,
		arg.CreatedAt,
		arg.InstitutionID,
		arg.AsesorID,
	)
	var i DirectoryTree
	err := row.Scan(
		&i.DirectoryID,
		&i.ParentID,
		&i.InstitutionID,
		&i.Name,
		&i.AsesorID,
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

const getDirectoryAsesorTreeParent = `-- name: GetDirectoryAsesorTreeParent :one
SELECT directory_id, parent_id, institution_id, name, asesor_id, created_at, updated_at, deleted_at FROM directory_tree
WHERE parent_id IS NULL AND deleted_at IS NULL AND asesor_id=$1 LIMIT 1
`

func (q *Queries) GetDirectoryAsesorTreeParent(ctx context.Context, asesorID pgtype.Int4) (DirectoryTree, error) {
	row := q.db.QueryRow(ctx, getDirectoryAsesorTreeParent, asesorID)
	var i DirectoryTree
	err := row.Scan(
		&i.DirectoryID,
		&i.ParentID,
		&i.InstitutionID,
		&i.Name,
		&i.AsesorID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getDirectoryIdByAsesorId = `-- name: GetDirectoryIdByAsesorId :one
SELECT directory_id,name,parent_id
  FROM directory_tree dt
  WHERE parent_id IS null AND dt.deleted_at IS NULL AND dt.asesor_id = $1
`

type GetDirectoryIdByAsesorIdRow struct {
	DirectoryID int64       `json:"directory_id"`
	Name        pgtype.Text `json:"name"`
	ParentID    pgtype.Int8 `json:"parent_id"`
}

func (q *Queries) GetDirectoryIdByAsesorId(ctx context.Context, asesorID pgtype.Int4) (GetDirectoryIdByAsesorIdRow, error) {
	row := q.db.QueryRow(ctx, getDirectoryIdByAsesorId, asesorID)
	var i GetDirectoryIdByAsesorIdRow
	err := row.Scan(&i.DirectoryID, &i.Name, &i.ParentID)
	return i, err
}

const getDirectoryIdByInstitutionId = `-- name: GetDirectoryIdByInstitutionId :one
SELECT directory_id,name,parent_id
  FROM directory_tree dt
  WHERE parent_id IS null AND dt.deleted_at IS NULL AND dt.institution_id = $1
`

type GetDirectoryIdByInstitutionIdRow struct {
	DirectoryID int64       `json:"directory_id"`
	Name        pgtype.Text `json:"name"`
	ParentID    pgtype.Int8 `json:"parent_id"`
}

func (q *Queries) GetDirectoryIdByInstitutionId(ctx context.Context, institutionID pgtype.Int4) (GetDirectoryIdByInstitutionIdRow, error) {
	row := q.db.QueryRow(ctx, getDirectoryIdByInstitutionId, institutionID)
	var i GetDirectoryIdByInstitutionIdRow
	err := row.Scan(&i.DirectoryID, &i.Name, &i.ParentID)
	return i, err
}

const getDirectoryInstitutionTreeParent = `-- name: GetDirectoryInstitutionTreeParent :one
SELECT directory_id, parent_id, institution_id, name, asesor_id, created_at, updated_at, deleted_at FROM directory_tree
WHERE parent_id IS NULL AND deleted_at IS NULL AND institution_id=$1 LIMIT 1
`

func (q *Queries) GetDirectoryInstitutionTreeParent(ctx context.Context, institutionID pgtype.Int4) (DirectoryTree, error) {
	row := q.db.QueryRow(ctx, getDirectoryInstitutionTreeParent, institutionID)
	var i DirectoryTree
	err := row.Scan(
		&i.DirectoryID,
		&i.ParentID,
		&i.InstitutionID,
		&i.Name,
		&i.AsesorID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getDirectoryTreeByAsesor = `-- name: GetDirectoryTreeByAsesor :one
SELECT directory_id, parent_id, institution_id, name, asesor_id, created_at, updated_at, deleted_at FROM directory_tree
WHERE directory_id = $1 AND deleted_at IS NULL AND asesor_id=$2 LIMIT 1
`

type GetDirectoryTreeByAsesorParams struct {
	DirectoryID int64       `json:"directory_id"`
	AsesorID    pgtype.Int4 `json:"asesor_id"`
}

func (q *Queries) GetDirectoryTreeByAsesor(ctx context.Context, arg GetDirectoryTreeByAsesorParams) (DirectoryTree, error) {
	row := q.db.QueryRow(ctx, getDirectoryTreeByAsesor, arg.DirectoryID, arg.AsesorID)
	var i DirectoryTree
	err := row.Scan(
		&i.DirectoryID,
		&i.ParentID,
		&i.InstitutionID,
		&i.Name,
		&i.AsesorID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getDirectoryTreeById = `-- name: GetDirectoryTreeById :one
SELECT directory_id, parent_id, institution_id, name, asesor_id, created_at, updated_at, deleted_at FROM directory_tree
WHERE directory_id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetDirectoryTreeById(ctx context.Context, directoryID int64) (DirectoryTree, error) {
	row := q.db.QueryRow(ctx, getDirectoryTreeById, directoryID)
	var i DirectoryTree
	err := row.Scan(
		&i.DirectoryID,
		&i.ParentID,
		&i.InstitutionID,
		&i.Name,
		&i.AsesorID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getDirectoryTreeByInstitution = `-- name: GetDirectoryTreeByInstitution :one
SELECT directory_id, parent_id, institution_id, name, asesor_id, created_at, updated_at, deleted_at FROM directory_tree
WHERE directory_id = $1 AND deleted_at IS NULL AND institution_id=$2 LIMIT 1
`

type GetDirectoryTreeByInstitutionParams struct {
	DirectoryID   int64       `json:"directory_id"`
	InstitutionID pgtype.Int4 `json:"institution_id"`
}

func (q *Queries) GetDirectoryTreeByInstitution(ctx context.Context, arg GetDirectoryTreeByInstitutionParams) (DirectoryTree, error) {
	row := q.db.QueryRow(ctx, getDirectoryTreeByInstitution, arg.DirectoryID, arg.InstitutionID)
	var i DirectoryTree
	err := row.Scan(
		&i.DirectoryID,
		&i.ParentID,
		&i.InstitutionID,
		&i.Name,
		&i.AsesorID,
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

const getInstitutionNameByDirectoryId = `-- name: GetInstitutionNameByDirectoryId :one
SELECT i.institution_id, i.asesor_id, i.institution_name, i.logo, i.description, i.created_at, i.updated_at, i.deleted_at, dt.directory_id, dt.parent_id, dt.institution_id, dt.name, dt.asesor_id, dt.created_at, dt.updated_at, dt.deleted_at
FROM directory_tree dt
LEFT JOIN institution i ON dt.institution_id = i.institution_id
WHERE dt.directory_id = $1
`

type GetInstitutionNameByDirectoryIdRow struct {
	Institution   Institution   `json:"institution"`
	DirectoryTree DirectoryTree `json:"directory_tree"`
}

func (q *Queries) GetInstitutionNameByDirectoryId(ctx context.Context, directoryID int64) (GetInstitutionNameByDirectoryIdRow, error) {
	row := q.db.QueryRow(ctx, getInstitutionNameByDirectoryId, directoryID)
	var i GetInstitutionNameByDirectoryIdRow
	err := row.Scan(
		&i.Institution.InstitutionID,
		&i.Institution.AsesorID,
		&i.Institution.InstitutionName,
		&i.Institution.Logo,
		&i.Institution.Description,
		&i.Institution.CreatedAt,
		&i.Institution.UpdatedAt,
		&i.Institution.DeletedAt,
		&i.DirectoryTree.DirectoryID,
		&i.DirectoryTree.ParentID,
		&i.DirectoryTree.InstitutionID,
		&i.DirectoryTree.Name,
		&i.DirectoryTree.AsesorID,
		&i.DirectoryTree.CreatedAt,
		&i.DirectoryTree.UpdatedAt,
		&i.DirectoryTree.DeletedAt,
	)
	return i, err
}

const listDirectorAsesoryChildByParent = `-- name: ListDirectorAsesoryChildByParent :many
SELECT directory_id, parent_id, institution_id, name, asesor_id, created_at, updated_at, deleted_at
FROM directory_tree
WHERE parent_id= $1 AND deleted_at IS NULL AND asesor_id=$2
`

type ListDirectorAsesoryChildByParentParams struct {
	ParentID pgtype.Int8 `json:"parent_id"`
	AsesorID pgtype.Int4 `json:"asesor_id"`
}

func (q *Queries) ListDirectorAsesoryChildByParent(ctx context.Context, arg ListDirectorAsesoryChildByParentParams) ([]DirectoryTree, error) {
	rows, err := q.db.Query(ctx, listDirectorAsesoryChildByParent, arg.ParentID, arg.AsesorID)
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
			&i.InstitutionID,
			&i.Name,
			&i.AsesorID,
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

const listDirectoryHierarchyAsesorById = `-- name: ListDirectoryHierarchyAsesorById :many
WITH RECURSIVE directory  AS (
  SELECT directory_id,name,parent_id
  FROM directory_tree dt
  WHERE parent_id IS null AND dt.deleted_at IS NULL AND dt.asesor_id=$1
  UNION ALL
  SELECT e.directory_id, e.name, e.parent_id
  FROM directory_tree  e
  INNER JOIN directory_tree eh ON e.parent_id = eh.directory_id
    where  e.directory_id<=$2 AND e.deleted_at IS NULL  AND e.asesor_id=$1
)
SELECT directory_id, name, parent_id FROM directory
`

type ListDirectoryHierarchyAsesorByIdParams struct {
	AsesorID    pgtype.Int4 `json:"asesor_id"`
	DirectoryID int64       `json:"directory_id"`
}

type ListDirectoryHierarchyAsesorByIdRow struct {
	DirectoryID int64       `json:"directory_id"`
	Name        pgtype.Text `json:"name"`
	ParentID    pgtype.Int8 `json:"parent_id"`
}

func (q *Queries) ListDirectoryHierarchyAsesorById(ctx context.Context, arg ListDirectoryHierarchyAsesorByIdParams) ([]ListDirectoryHierarchyAsesorByIdRow, error) {
	rows, err := q.db.Query(ctx, listDirectoryHierarchyAsesorById, arg.AsesorID, arg.DirectoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListDirectoryHierarchyAsesorByIdRow
	for rows.Next() {
		var i ListDirectoryHierarchyAsesorByIdRow
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

const listDirectoryHierarchyInstitutionById = `-- name: ListDirectoryHierarchyInstitutionById :many
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
SELECT directory_id, name, parent_id FROM directory
`

type ListDirectoryHierarchyInstitutionByIdParams struct {
	InstitutionID pgtype.Int4 `json:"institution_id"`
	DirectoryID   int64       `json:"directory_id"`
}

type ListDirectoryHierarchyInstitutionByIdRow struct {
	DirectoryID int64       `json:"directory_id"`
	Name        pgtype.Text `json:"name"`
	ParentID    pgtype.Int8 `json:"parent_id"`
}

func (q *Queries) ListDirectoryHierarchyInstitutionById(ctx context.Context, arg ListDirectoryHierarchyInstitutionByIdParams) ([]ListDirectoryHierarchyInstitutionByIdRow, error) {
	rows, err := q.db.Query(ctx, listDirectoryHierarchyInstitutionById, arg.InstitutionID, arg.DirectoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListDirectoryHierarchyInstitutionByIdRow
	for rows.Next() {
		var i ListDirectoryHierarchyInstitutionByIdRow
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

const listDirectoryInstitutionChildByParent = `-- name: ListDirectoryInstitutionChildByParent :many
SELECT directory_id, parent_id, institution_id, name, asesor_id, created_at, updated_at, deleted_at
FROM directory_tree
WHERE parent_id= $1 AND deleted_at IS NULL AND institution_id=$2
`

type ListDirectoryInstitutionChildByParentParams struct {
	ParentID      pgtype.Int8 `json:"parent_id"`
	InstitutionID pgtype.Int4 `json:"institution_id"`
}

func (q *Queries) ListDirectoryInstitutionChildByParent(ctx context.Context, arg ListDirectoryInstitutionChildByParentParams) ([]DirectoryTree, error) {
	rows, err := q.db.Query(ctx, listDirectoryInstitutionChildByParent, arg.ParentID, arg.InstitutionID)
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
			&i.InstitutionID,
			&i.Name,
			&i.AsesorID,
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

const listDirectoryTrees = `-- name: ListDirectoryTrees :many
SELECT directory_id, parent_id, institution_id, name, asesor_id, created_at, updated_at, deleted_at FROM directory_tree
ORDER BY directory_id AND deleted_at IS NULL
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
			&i.InstitutionID,
			&i.Name,
			&i.AsesorID,
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

const updateDirectoryTreeById = `-- name: UpdateDirectoryTreeById :exec
UPDATE directory_tree
SET name = $2, updated_at = $3, parent_id = $4, asesor_id = $5
WHERE directory_id = $1
`

type UpdateDirectoryTreeByIdParams struct {
	DirectoryID int64            `json:"directory_id"`
	Name        pgtype.Text      `json:"name"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	ParentID    pgtype.Int8      `json:"parent_id"`
	AsesorID    pgtype.Int4      `json:"asesor_id"`
}

func (q *Queries) UpdateDirectoryTreeById(ctx context.Context, arg UpdateDirectoryTreeByIdParams) error {
	_, err := q.db.Exec(ctx, updateDirectoryTreeById,
		arg.DirectoryID,
		arg.Name,
		arg.UpdatedAt,
		arg.ParentID,
		arg.AsesorID,
	)
	return err
}
