// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: role_permission.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createRolePermission = `-- name: CreateRolePermission :one
INSERT INTO role_permission(role_id, permission_id, created_at)
VALUES ($1, $2, $3)
RETURNING role_permission_id, role_id, permission_id, created_at, updated_at, deleted_at
`

type CreateRolePermissionParams struct {
	RoleID       int32            `json:"role_id"`
	PermissionID int32            `json:"permission_id"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateRolePermission(ctx context.Context, arg CreateRolePermissionParams) (RolePermission, error) {
	row := q.db.QueryRow(ctx, createRolePermission, arg.RoleID, arg.PermissionID, arg.CreatedAt)
	var i RolePermission
	err := row.Scan(
		&i.RolePermissionID,
		&i.RoleID,
		&i.PermissionID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllRolePermissions = `-- name: DeleteAllRolePermissions :execresult
UPDATE role_permission
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllRolePermissions(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllRolePermissions, deletedAt)
}

const deleteRolePermissionById = `-- name: DeleteRolePermissionById :exec
UPDATE role_permission
SET deleted_at = $2
WHERE role_permission_id = $1
`

type DeleteRolePermissionByIdParams struct {
	RolePermissionID int64            `json:"role_permission_id"`
	DeletedAt        pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteRolePermissionById(ctx context.Context, arg DeleteRolePermissionByIdParams) error {
	_, err := q.db.Exec(ctx, deleteRolePermissionById, arg.RolePermissionID, arg.DeletedAt)
	return err
}

const getRolePermission = `-- name: GetRolePermission :one
SELECT role_id, permission_id
FROM role_permission
WHERE role_permission_id = $1
`

type GetRolePermissionRow struct {
	RoleID       int32 `json:"role_id"`
	PermissionID int32 `json:"permission_id"`
}

func (q *Queries) GetRolePermission(ctx context.Context, rolePermissionID int64) (GetRolePermissionRow, error) {
	row := q.db.QueryRow(ctx, getRolePermission, rolePermissionID)
	var i GetRolePermissionRow
	err := row.Scan(&i.RoleID, &i.PermissionID)
	return i, err
}

const listPermissionByRoleId = `-- name: ListPermissionByRoleId :many
SELECT p.permission_id, p.name, p.code, p.description, p.created_at, p.updated_at, p.deleted_at, r.role_id, r.role_name, r.description, r.created_at, r.updated_at, r.deleted_at, rp.role_permission_id, rp.role_id, rp.permission_id, rp.created_at, rp.updated_at, rp.deleted_at
FROM role_permission rp
JOIN permission p ON rp.permission_id = p.permission_id
JOIN roles r ON rp.role_id = r.role_id
WHERE rp.role_id = $1
`

type ListPermissionByRoleIdRow struct {
	Permission     Permission     `json:"permission"`
	Role           Role           `json:"role"`
	RolePermission RolePermission `json:"role_permission"`
}

func (q *Queries) ListPermissionByRoleId(ctx context.Context, roleID int32) ([]ListPermissionByRoleIdRow, error) {
	rows, err := q.db.Query(ctx, listPermissionByRoleId, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPermissionByRoleIdRow
	for rows.Next() {
		var i ListPermissionByRoleIdRow
		if err := rows.Scan(
			&i.Permission.PermissionID,
			&i.Permission.Name,
			&i.Permission.Code,
			&i.Permission.Description,
			&i.Permission.CreatedAt,
			&i.Permission.UpdatedAt,
			&i.Permission.DeletedAt,
			&i.Role.RoleID,
			&i.Role.RoleName,
			&i.Role.Description,
			&i.Role.CreatedAt,
			&i.Role.UpdatedAt,
			&i.Role.DeletedAt,
			&i.RolePermission.RolePermissionID,
			&i.RolePermission.RoleID,
			&i.RolePermission.PermissionID,
			&i.RolePermission.CreatedAt,
			&i.RolePermission.UpdatedAt,
			&i.RolePermission.DeletedAt,
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

const listRolePermissions = `-- name: ListRolePermissions :many
SELECT role_permission_id, role_id, permission_id, created_at, updated_at, deleted_at FROM role_permission
ORDER BY role_permission_id
`

func (q *Queries) ListRolePermissions(ctx context.Context) ([]RolePermission, error) {
	rows, err := q.db.Query(ctx, listRolePermissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RolePermission
	for rows.Next() {
		var i RolePermission
		if err := rows.Scan(
			&i.RolePermissionID,
			&i.RoleID,
			&i.PermissionID,
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

const updateRolePermissionById = `-- name: UpdateRolePermissionById :exec
UPDATE role_permission
SET role_id = $2, permission_id = $3, updated_at = $4
WHERE role_permission_id = $1
`

type UpdateRolePermissionByIdParams struct {
	RolePermissionID int64            `json:"role_permission_id"`
	RoleID           int32            `json:"role_id"`
	PermissionID     int32            `json:"permission_id"`
	UpdatedAt        pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateRolePermissionById(ctx context.Context, arg UpdateRolePermissionByIdParams) error {
	_, err := q.db.Exec(ctx, updateRolePermissionById,
		arg.RolePermissionID,
		arg.RoleID,
		arg.PermissionID,
		arg.UpdatedAt,
	)
	return err
}
