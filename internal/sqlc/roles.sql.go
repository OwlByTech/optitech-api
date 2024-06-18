// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: roles.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createRole = `-- name: CreateRole :one
INSERT INTO roles(role_name, description, created_at)
VALUES ($1, $2, $3)
RETURNING role_id, role_name, description, created_at, updated_at, deleted_at
`

type CreateRoleParams struct {
	RoleName    string           `json:"role_name"`
	Description string           `json:"description"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRow(ctx, createRole, arg.RoleName, arg.Description, arg.CreatedAt)
	var i Role
	err := row.Scan(
		&i.RoleID,
		&i.RoleName,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllRoles = `-- name: DeleteAllRoles :execresult
UPDATE roles
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllRoles(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllRoles, deletedAt)
}

const deleteRoleById = `-- name: DeleteRoleById :exec
UPDATE roles
SET deleted_at = $2
WHERE role_id = $1
`

type DeleteRoleByIdParams struct {
	RoleID    int64            `json:"role_id"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteRoleById(ctx context.Context, arg DeleteRoleByIdParams) error {
	_, err := q.db.Exec(ctx, deleteRoleById, arg.RoleID, arg.DeletedAt)
	return err
}

const getRole = `-- name: GetRole :one
SELECT role_id, role_name, description, created_at, updated_at, deleted_at FROM roles
WHERE role_id = $1 LIMIT 1
`

func (q *Queries) GetRole(ctx context.Context, roleID int64) (Role, error) {
	row := q.db.QueryRow(ctx, getRole, roleID)
	var i Role
	err := row.Scan(
		&i.RoleID,
		&i.RoleName,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getRoleByName = `-- name: GetRoleByName :one
SELECT role_name, description
FROM roles
WHERE role_id = $1
`

type GetRoleByNameRow struct {
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
}

func (q *Queries) GetRoleByName(ctx context.Context, roleID int64) (GetRoleByNameRow, error) {
	row := q.db.QueryRow(ctx, getRoleByName, roleID)
	var i GetRoleByNameRow
	err := row.Scan(&i.RoleName, &i.Description)
	return i, err
}

const listRoles = `-- name: ListRoles :many
SELECT role_id, role_name, description, created_at, updated_at, deleted_at FROM roles
ORDER BY role_id
`

func (q *Queries) ListRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.Query(ctx, listRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.RoleID,
			&i.RoleName,
			&i.Description,
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

const updateRoleById = `-- name: UpdateRoleById :exec
UPDATE roles
SET role_name = $2, description = $3, updated_at = $4
WHERE role_id = $1
`

type UpdateRoleByIdParams struct {
	RoleID      int64            `json:"role_id"`
	RoleName    string           `json:"role_name"`
	Description string           `json:"description"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateRoleById(ctx context.Context, arg UpdateRoleByIdParams) error {
	_, err := q.db.Exec(ctx, updateRoleById,
		arg.RoleID,
		arg.RoleName,
		arg.Description,
		arg.UpdatedAt,
	)
	return err
}
