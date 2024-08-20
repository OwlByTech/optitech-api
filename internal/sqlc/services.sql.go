// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: services.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createServices = `-- name: CreateServices :one
INSERT INTO services(name, created_at)
VALUES ($1, $2)
RETURNING service_id, name, created_at, updated_at, deleted_at
`

type CreateServicesParams struct {
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateServices(ctx context.Context, arg CreateServicesParams) (Service, error) {
	row := q.db.QueryRow(ctx, createServices, arg.Name, arg.CreatedAt)
	var i Service
	err := row.Scan(
		&i.ServiceID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllServicess = `-- name: DeleteAllServicess :execresult
UPDATE services
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllServicess(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllServicess, deletedAt)
}

const deleteService = `-- name: DeleteService :exec
UPDATE services
SET deleted_at = $2
WHERE service_id = $1
`

type DeleteServiceParams struct {
	ServiceID int32            `json:"service_id"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteService(ctx context.Context, arg DeleteServiceParams) error {
	_, err := q.db.Exec(ctx, deleteService, arg.ServiceID, arg.DeletedAt)
	return err
}

const getService = `-- name: GetService :one
SELECT service_id, name, created_at, updated_at, deleted_at FROM services
WHERE service_id = $1 AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetService(ctx context.Context, serviceID int32) (Service, error) {
	row := q.db.QueryRow(ctx, getService, serviceID)
	var i Service
	err := row.Scan(
		&i.ServiceID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getServicesByName = `-- name: GetServicesByName :one
SELECT name
FROM services
WHERE service_id = $1
`

func (q *Queries) GetServicesByName(ctx context.Context, serviceID int32) (string, error) {
	row := q.db.QueryRow(ctx, getServicesByName, serviceID)
	var name string
	err := row.Scan(&name)
	return name, err
}

const listServices = `-- name: ListServices :many
SELECT service_id, name, created_at, updated_at, deleted_at FROM services
ORDER BY service_id
`

func (q *Queries) ListServices(ctx context.Context) ([]Service, error) {
	rows, err := q.db.Query(ctx, listServices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Service
	for rows.Next() {
		var i Service
		if err := rows.Scan(
			&i.ServiceID,
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

const updateService = `-- name: UpdateService :exec
UPDATE services
SET name = $2, updated_at = $3
WHERE service_id = $1
`

type UpdateServiceParams struct {
	ServiceID int32            `json:"service_id"`
	Name      string           `json:"name"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateService(ctx context.Context, arg UpdateServiceParams) error {
	_, err := q.db.Exec(ctx, updateService, arg.ServiceID, arg.Name, arg.UpdatedAt)
	return err
}
