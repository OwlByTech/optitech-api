// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: client_role.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createClientRole = `-- name: CreateClientRole :one
INSERT INTO client_role(client_id, role_id, created_at)
VALUES ($1, $2, $3)
RETURNING client_role_id, client_id, role_id, created_at, updated_at, deleted_at
`

type CreateClientRoleParams struct {
	ClientID  int32            `json:"client_id"`
	RoleID    int32            `json:"role_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateClientRole(ctx context.Context, arg CreateClientRoleParams) (ClientRole, error) {
	row := q.db.QueryRow(ctx, createClientRole, arg.ClientID, arg.RoleID, arg.CreatedAt)
	var i ClientRole
	err := row.Scan(
		&i.ClientRoleID,
		&i.ClientID,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllClientRoles = `-- name: DeleteAllClientRoles :execresult
UPDATE client_role
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllClientRoles(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllClientRoles, deletedAt)
}

const deleteClientRoleById = `-- name: DeleteClientRoleById :exec
UPDATE client_role
SET deleted_at = $2
WHERE client_role_id = $1
`

type DeleteClientRoleByIdParams struct {
	ClientRoleID int64            `json:"client_role_id"`
	DeletedAt    pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteClientRoleById(ctx context.Context, arg DeleteClientRoleByIdParams) error {
	_, err := q.db.Exec(ctx, deleteClientRoleById, arg.ClientRoleID, arg.DeletedAt)
	return err
}

const getClientRole = `-- name: GetClientRole :one
SELECT client_role_id, client_id, role_id, created_at, updated_at, deleted_at FROM client_role
WHERE client_role_id = $1 LIMIT 1
`

func (q *Queries) GetClientRole(ctx context.Context, clientRoleID int64) (ClientRole, error) {
	row := q.db.QueryRow(ctx, getClientRole, clientRoleID)
	var i ClientRole
	err := row.Scan(
		&i.ClientRoleID,
		&i.ClientID,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getClientRoleByName = `-- name: GetClientRoleByName :one
SELECT client_id, role_id
FROM client_role
WHERE client_role_id = $1
`

type GetClientRoleByNameRow struct {
	ClientID int32 `json:"client_id"`
	RoleID   int32 `json:"role_id"`
}

func (q *Queries) GetClientRoleByName(ctx context.Context, clientRoleID int64) (GetClientRoleByNameRow, error) {
	row := q.db.QueryRow(ctx, getClientRoleByName, clientRoleID)
	var i GetClientRoleByNameRow
	err := row.Scan(&i.ClientID, &i.RoleID)
	return i, err
}

const listClientRoles = `-- name: ListClientRoles :many
SELECT client_role_id, client_id, role_id, created_at, updated_at, deleted_at FROM client_role
ORDER BY client_role_id
`

func (q *Queries) ListClientRoles(ctx context.Context) ([]ClientRole, error) {
	rows, err := q.db.Query(ctx, listClientRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClientRole
	for rows.Next() {
		var i ClientRole
		if err := rows.Scan(
			&i.ClientRoleID,
			&i.ClientID,
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

const updateClientRoleById = `-- name: UpdateClientRoleById :exec
UPDATE client_role
SET client_id = $2, role_id = $3, updated_at = $4
WHERE client_role_id = $1
`

type UpdateClientRoleByIdParams struct {
	ClientRoleID int64            `json:"client_role_id"`
	ClientID     int32            `json:"client_id"`
	RoleID       int32            `json:"role_id"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateClientRoleById(ctx context.Context, arg UpdateClientRoleByIdParams) error {
	_, err := q.db.Exec(ctx, updateClientRoleById,
		arg.ClientRoleID,
		arg.ClientID,
		arg.RoleID,
		arg.UpdatedAt,
	)
	return err
}
