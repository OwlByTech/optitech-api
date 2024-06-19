// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: institution.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createInstitution = `-- name: CreateInstitution :one
INSERT INTO institution (institution_name, logo, description, created_at)
VALUES ($1, $2, $3, $4)
RETURNING institution_id, asesor_id, institution_name, logo, description, created_at, updated_at, deleted_at
`

type CreateInstitutionParams struct {
	InstitutionName string           `json:"institution_name"`
	Logo            pgtype.Text      `json:"logo"`
	Description     string           `json:"description"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateInstitution(ctx context.Context, arg CreateInstitutionParams) (Institution, error) {
	row := q.db.QueryRow(ctx, createInstitution,
		arg.InstitutionName,
		arg.Logo,
		arg.Description,
		arg.CreatedAt,
	)
	var i Institution
	err := row.Scan(
		&i.InstitutionID,
		&i.AsesorID,
		&i.InstitutionName,
		&i.Logo,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllInstitutions = `-- name: DeleteAllInstitutions :exec
UPDATE institution
set deleted_at = $1 
where deleted_at is NULL
`

func (q *Queries) DeleteAllInstitutions(ctx context.Context, deletedAt pgtype.Timestamp) error {
	_, err := q.db.Exec(ctx, deleteAllInstitutions, deletedAt)
	return err
}

const deleteInstitutionById = `-- name: DeleteInstitutionById :exec
UPDATE institution
SET deleted_at = $2
WHERE institution_id = $1 AND deleted_at IS NULL
`

type DeleteInstitutionByIdParams struct {
	InstitutionID int64            `json:"institution_id"`
	DeletedAt     pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteInstitutionById(ctx context.Context, arg DeleteInstitutionByIdParams) error {
	_, err := q.db.Exec(ctx, deleteInstitutionById, arg.InstitutionID, arg.DeletedAt)
	return err
}

const getInstitution = `-- name: GetInstitution :one
SELECT institution_id, asesor_id, institution_name, logo, description, created_at, updated_at, deleted_at FROM institution
WHERE institution_id = $1 LIMIT 1
`

func (q *Queries) GetInstitution(ctx context.Context, institutionID int64) (Institution, error) {
	row := q.db.QueryRow(ctx, getInstitution, institutionID)
	var i Institution
	err := row.Scan(
		&i.InstitutionID,
		&i.AsesorID,
		&i.InstitutionName,
		&i.Logo,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getInstitutionByName = `-- name: GetInstitutionByName :one
SELECT  institution_name, logo, description
FROM institution
WHERE institution_name = $1
`

type GetInstitutionByNameRow struct {
	InstitutionName string      `json:"institution_name"`
	Logo            pgtype.Text `json:"logo"`
	Description     string      `json:"description"`
}

func (q *Queries) GetInstitutionByName(ctx context.Context, institutionName string) (GetInstitutionByNameRow, error) {
	row := q.db.QueryRow(ctx, getInstitutionByName, institutionName)
	var i GetInstitutionByNameRow
	err := row.Scan(&i.InstitutionName, &i.Logo, &i.Description)
	return i, err
}

const listInstitutions = `-- name: ListInstitutions :many
SELECT institution_id, asesor_id, institution_name, logo, description, created_at, updated_at, deleted_at FROM institution
ORDER BY institution_name
`

func (q *Queries) ListInstitutions(ctx context.Context) ([]Institution, error) {
	rows, err := q.db.Query(ctx, listInstitutions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Institution
	for rows.Next() {
		var i Institution
		if err := rows.Scan(
			&i.InstitutionID,
			&i.AsesorID,
			&i.InstitutionName,
			&i.Logo,
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

const updateInstitutionById = `-- name: UpdateInstitutionById :exec
UPDATE institution
SET institution_name = $2, logo = $3, description = $4,  updated_at=$5
WHERE institution_id = $1
`

type UpdateInstitutionByIdParams struct {
	InstitutionID   int64            `json:"institution_id"`
	InstitutionName string           `json:"institution_name"`
	Logo            pgtype.Text      `json:"logo"`
	Description     string           `json:"description"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateInstitutionById(ctx context.Context, arg UpdateInstitutionByIdParams) error {
	_, err := q.db.Exec(ctx, updateInstitutionById,
		arg.InstitutionID,
		arg.InstitutionName,
		arg.Logo,
		arg.Description,
		arg.UpdatedAt,
	)
	return err
}
