// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: institution_services.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateInstitutionServicesParams struct {
	InstitutionID int32            `json:"institution_id"`
	ServiceID     int32            `json:"service_id"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
}

const deleteAllInstitutionServices = `-- name: DeleteAllInstitutionServices :execresult
UPDATE institution_services
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllInstitutionServices(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllInstitutionServices, deletedAt)
}

const deleteInstitutionServiceById = `-- name: DeleteInstitutionServiceById :exec
UPDATE institution_services
SET deleted_at = $2
WHERE institution_id= $1 AND service_id= $3
`

type DeleteInstitutionServiceByIdParams struct {
	InstitutionID int32            `json:"institution_id"`
	DeletedAt     pgtype.Timestamp `json:"deleted_at"`
	ServiceID     int32            `json:"service_id"`
}

func (q *Queries) DeleteInstitutionServiceById(ctx context.Context, arg DeleteInstitutionServiceByIdParams) error {
	_, err := q.db.Exec(ctx, deleteInstitutionServiceById, arg.InstitutionID, arg.DeletedAt, arg.ServiceID)
	return err
}

const deleteInstitutionServicesByInstitution = `-- name: DeleteInstitutionServicesByInstitution :exec
UPDATE institution_services
SET deleted_at = $2
WHERE institution_id= $1
`

type DeleteInstitutionServicesByInstitutionParams struct {
	InstitutionID int32            `json:"institution_id"`
	DeletedAt     pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteInstitutionServicesByInstitution(ctx context.Context, arg DeleteInstitutionServicesByInstitutionParams) error {
	_, err := q.db.Exec(ctx, deleteInstitutionServicesByInstitution, arg.InstitutionID, arg.DeletedAt)
	return err
}

const existsInstitutionService = `-- name: ExistsInstitutionService :one
SELECT institution_id, service_id, created_at, updated_at, deleted_at FROM institution_services
WHERE service_id = $1 AND institution_id=$2 and deleted_at IS NOT NULL
`

type ExistsInstitutionServiceParams struct {
	ServiceID     int32 `json:"service_id"`
	InstitutionID int32 `json:"institution_id"`
}

func (q *Queries) ExistsInstitutionService(ctx context.Context, arg ExistsInstitutionServiceParams) (InstitutionService, error) {
	row := q.db.QueryRow(ctx, existsInstitutionService, arg.ServiceID, arg.InstitutionID)
	var i InstitutionService
	err := row.Scan(
		&i.InstitutionID,
		&i.ServiceID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listInstitutionServices = `-- name: ListInstitutionServices :many
SELECT services.name ,services.service_id FROM institution_services
INNER JOIN services ON  institution_services.service_id=services.service_id
WHERE institution_services.institution_id= $1
AND institution_services.deleted_at IS NULL
ORDER BY services.service_id
`

type ListInstitutionServicesRow struct {
	Name      string `json:"name"`
	ServiceID int32  `json:"service_id"`
}

func (q *Queries) ListInstitutionServices(ctx context.Context, institutionID int32) ([]ListInstitutionServicesRow, error) {
	rows, err := q.db.Query(ctx, listInstitutionServices, institutionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListInstitutionServicesRow
	for rows.Next() {
		var i ListInstitutionServicesRow
		if err := rows.Scan(&i.Name, &i.ServiceID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const recoverInstitutionService = `-- name: RecoverInstitutionService :exec
UPDATE institution_services
SET deleted_at = NULL,updated_at= $2
WHERE institution_id= $1 AND service_id= $3
`

type RecoverInstitutionServiceParams struct {
	InstitutionID int32            `json:"institution_id"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	ServiceID     int32            `json:"service_id"`
}

func (q *Queries) RecoverInstitutionService(ctx context.Context, arg RecoverInstitutionServiceParams) error {
	_, err := q.db.Exec(ctx, recoverInstitutionService, arg.InstitutionID, arg.UpdatedAt, arg.ServiceID)
	return err
}
