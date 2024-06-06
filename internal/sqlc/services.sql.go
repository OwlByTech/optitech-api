// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: services.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createServices = `-- name: CreateServices :one
INSERT INTO services(service_name, created_at)
VALUES ($1, $2)
RETURNING services_id, service_name, created_at, updated_at, deleted_at
`

type CreateServicesParams struct {
	ServiceName string    `json:"service_name"`
	CreatedAt   time.Time `json:"created_at"`
}

func (q *Queries) CreateServices(ctx context.Context, arg CreateServicesParams) (Service, error) {
	row := q.db.QueryRowContext(ctx, createServices, arg.ServiceName, arg.CreatedAt)
	var i Service
	err := row.Scan(
		&i.ServicesID,
		&i.ServiceName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllServicess = `-- name: DeleteAllServicess :execresult
UPDATE services
SET deleted_at = $1
WHERE services_id IS NULL
`

func (q *Queries) DeleteAllServicess(ctx context.Context, deletedAt sql.NullTime) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteAllServicess, deletedAt)
}

const deleteServicesById = `-- name: DeleteServicesById :exec
UPDATE services
SET deleted_at = $2
WHERE services_id = $1
`

type DeleteServicesByIdParams struct {
	ServicesID int64        `json:"services_id"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
}

func (q *Queries) DeleteServicesById(ctx context.Context, arg DeleteServicesByIdParams) error {
	_, err := q.db.ExecContext(ctx, deleteServicesById, arg.ServicesID, arg.DeletedAt)
	return err
}

const getServices = `-- name: GetServices :one
SELECT services_id, service_name, created_at, updated_at, deleted_at FROM services
WHERE services_id = $1 LIMIT 1
`

func (q *Queries) GetServices(ctx context.Context, servicesID int64) (Service, error) {
	row := q.db.QueryRowContext(ctx, getServices, servicesID)
	var i Service
	err := row.Scan(
		&i.ServicesID,
		&i.ServiceName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getServicesByName = `-- name: GetServicesByName :one
SELECT service_name
FROM services
WHERE services_id = $1
`

func (q *Queries) GetServicesByName(ctx context.Context, servicesID int64) (string, error) {
	row := q.db.QueryRowContext(ctx, getServicesByName, servicesID)
	var service_name string
	err := row.Scan(&service_name)
	return service_name, err
}

const listServicess = `-- name: ListServicess :many
SELECT services_id, service_name, created_at, updated_at, deleted_at FROM services
ORDER BY services_id
`

func (q *Queries) ListServicess(ctx context.Context) ([]Service, error) {
	rows, err := q.db.QueryContext(ctx, listServicess)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Service
	for rows.Next() {
		var i Service
		if err := rows.Scan(
			&i.ServicesID,
			&i.ServiceName,
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

const updateServicesById = `-- name: UpdateServicesById :exec
UPDATE services
SET service_name = $2, updated_at = $3
WHERE services_id = $1
`

type UpdateServicesByIdParams struct {
	ServicesID  int64        `json:"services_id"`
	ServiceName string       `json:"service_name"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

func (q *Queries) UpdateServicesById(ctx context.Context, arg UpdateServicesByIdParams) error {
	_, err := q.db.ExecContext(ctx, updateServicesById, arg.ServicesID, arg.ServiceName, arg.UpdatedAt)
	return err
}