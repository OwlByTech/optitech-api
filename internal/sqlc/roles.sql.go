// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: roles.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createRole = `-- name: CreateRole :one
INSERT INTO roles(role_name, created_at)
VALUES ($1, $2)
RETURNING role_id, role_name, created_at, updated_at, deleted_at
`

type CreateRoleParams struct {
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, createRole, arg.RoleName, arg.CreatedAt)
	var i Role
	err := row.Scan(
		&i.RoleID,
		&i.RoleName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllRoles = `-- name: DeleteAllRoles :execresult
UPDATE roles
SET deleted_at = $1
WHERE role_id IS NULL
`

func (q *Queries) DeleteAllRoles(ctx context.Context, deletedAt sql.NullTime) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteAllRoles, deletedAt)
}

const deleteRoleById = `-- name: DeleteRoleById :exec
UPDATE roles
SET deleted_at = $2
WHERE role_id = $1
`

type DeleteRoleByIdParams struct {
	RoleID    int64        `json:"role_id"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func (q *Queries) DeleteRoleById(ctx context.Context, arg DeleteRoleByIdParams) error {
	_, err := q.db.ExecContext(ctx, deleteRoleById, arg.RoleID, arg.DeletedAt)
	return err
}

const getRole = `-- name: GetRole :one
SELECT role_id, role_name, created_at, updated_at, deleted_at FROM roles
WHERE role_id = $1 LIMIT 1
`

func (q *Queries) GetRole(ctx context.Context, roleID int64) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRole, roleID)
	var i Role
	err := row.Scan(
		&i.RoleID,
		&i.RoleName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getRoleByName = `-- name: GetRoleByName :one
SELECT role_name
FROM roles
WHERE role_id = $1
`

func (q *Queries) GetRoleByName(ctx context.Context, roleID int64) (string, error) {
	row := q.db.QueryRowContext(ctx, getRoleByName, roleID)
	var role_name string
	err := row.Scan(&role_name)
	return role_name, err
}

const listRoles = `-- name: ListRoles :many
SELECT role_id, role_name, created_at, updated_at, deleted_at FROM roles
ORDER BY role_id
`

func (q *Queries) ListRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.QueryContext(ctx, listRoles)
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
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRoleById = `-- name: UpdateRoleById :exec
UPDATE roles
SET role_name = $2, updated_at = $3
WHERE role_id = $1
`

type UpdateRoleByIdParams struct {
	RoleID    int64        `json:"role_id"`
	RoleName  string       `json:"role_name"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

func (q *Queries) UpdateRoleById(ctx context.Context, arg UpdateRoleByIdParams) error {
	_, err := q.db.ExecContext(ctx, updateRoleById, arg.RoleID, arg.RoleName, arg.UpdatedAt)
	return err
}
