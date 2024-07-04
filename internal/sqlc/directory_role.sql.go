// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: directory_role.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createDirectoryRole = `-- name: CreateDirectoryRole :one
INSERT INTO directory_role(directory_id, role_id, created_at)
VALUES ($1, $2, $3)
RETURNING directory_role_id, directory_id, role_id, created_at, updated_at, deleted_at
`

type CreateDirectoryRoleParams struct {
	DirectoryID pgtype.Int4      `json:"directory_id"`
	RoleID      pgtype.Int4      `json:"role_id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateDirectoryRole(ctx context.Context, arg CreateDirectoryRoleParams) (DirectoryRole, error) {
	row := q.db.QueryRow(ctx, createDirectoryRole, arg.DirectoryID, arg.RoleID, arg.CreatedAt)
	var i DirectoryRole
	err := row.Scan(
		&i.DirectoryRoleID,
		&i.DirectoryID,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllDirectoryRoles = `-- name: DeleteAllDirectoryRoles :execresult
UPDATE directory_role
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllDirectoryRoles(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllDirectoryRoles, deletedAt)
}

const deleteDirectoryRoleById = `-- name: DeleteDirectoryRoleById :exec
UPDATE directory_role
SET deleted_at = $2
WHERE directory_role_id = $1
`

type DeleteDirectoryRoleByIdParams struct {
	DirectoryRoleID int64            `json:"directory_role_id"`
	DeletedAt       pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteDirectoryRoleById(ctx context.Context, arg DeleteDirectoryRoleByIdParams) error {
	_, err := q.db.Exec(ctx, deleteDirectoryRoleById, arg.DirectoryRoleID, arg.DeletedAt)
	return err
}

const getDirectoryRole = `-- name: GetDirectoryRole :one
SELECT directory_role_id, directory_id, role_id, created_at, updated_at, deleted_at FROM directory_role
WHERE directory_role_id = $1 LIMIT 1
`

func (q *Queries) GetDirectoryRole(ctx context.Context, directoryRoleID int64) (DirectoryRole, error) {
	row := q.db.QueryRow(ctx, getDirectoryRole, directoryRoleID)
	var i DirectoryRole
	err := row.Scan(
		&i.DirectoryRoleID,
		&i.DirectoryID,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getDirectoryRoleByName = `-- name: GetDirectoryRoleByName :one
SELECT directory_id, role_id
FROM directory_role
WHERE directory_role_id = $1
`

type GetDirectoryRoleByNameRow struct {
	DirectoryID pgtype.Int4 `json:"directory_id"`
	RoleID      pgtype.Int4 `json:"role_id"`
}

func (q *Queries) GetDirectoryRoleByName(ctx context.Context, directoryRoleID int64) (GetDirectoryRoleByNameRow, error) {
	row := q.db.QueryRow(ctx, getDirectoryRoleByName, directoryRoleID)
	var i GetDirectoryRoleByNameRow
	err := row.Scan(&i.DirectoryID, &i.RoleID)
	return i, err
}

const listDirectoryRoles = `-- name: ListDirectoryRoles :many
SELECT directory_role_id, directory_id, role_id, created_at, updated_at, deleted_at FROM directory_role
ORDER BY directory_role_id
`

func (q *Queries) ListDirectoryRoles(ctx context.Context) ([]DirectoryRole, error) {
	rows, err := q.db.Query(ctx, listDirectoryRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DirectoryRole
	for rows.Next() {
		var i DirectoryRole
		if err := rows.Scan(
			&i.DirectoryRoleID,
			&i.DirectoryID,
			&i.RoleID,
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

const updateDirectoryRole = `-- name: UpdateDirectoryRole :exec
UPDATE directory_role
SET role_id = $2, updated_at = $3
WHERE directory_id = $1
`

type UpdateDirectoryRoleParams struct {
	DirectoryID pgtype.Int4      `json:"directory_id"`
	RoleID      pgtype.Int4      `json:"role_id"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateDirectoryRole(ctx context.Context, arg UpdateDirectoryRoleParams) error {
	_, err := q.db.Exec(ctx, updateDirectoryRole, arg.DirectoryID, arg.RoleID, arg.UpdatedAt)
	return err
}
