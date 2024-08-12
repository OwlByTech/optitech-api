// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: format.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createFormat = `-- name: CreateFormat :one
INSERT INTO format(updated_format_id, asesor_id, service_id, format_name, description, extension, version, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING format_id, updated_format_id, asesor_id, service_id, format_name, description, extension, version, created_at, updated_at, deleted_at
`

type CreateFormatParams struct {
	UpdatedFormatID pgtype.Int4      `json:"updated_format_id"`
	AsesorID        int32            `json:"asesor_id"`
	ServiceID       pgtype.Int4      `json:"service_id"`
	FormatName      string           `json:"format_name"`
	Description     string           `json:"description"`
	Extension       Extensions       `json:"extension"`
	Version         string           `json:"version"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateFormat(ctx context.Context, arg CreateFormatParams) (Format, error) {
	row := q.db.QueryRow(ctx, createFormat,
		arg.UpdatedFormatID,
		arg.AsesorID,
		arg.ServiceID,
		arg.FormatName,
		arg.Description,
		arg.Extension,
		arg.Version,
		arg.CreatedAt,
	)
	var i Format
	err := row.Scan(
		&i.FormatID,
		&i.UpdatedFormatID,
		&i.AsesorID,
		&i.ServiceID,
		&i.FormatName,
		&i.Description,
		&i.Extension,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllFormats = `-- name: DeleteAllFormats :execresult
UPDATE format
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllFormats(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllFormats, deletedAt)
}

const deleteFormatById = `-- name: DeleteFormatById :exec
UPDATE format
SET deleted_at = $2
WHERE format_id = $1
`

type DeleteFormatByIdParams struct {
	FormatID  int32            `json:"format_id"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteFormatById(ctx context.Context, arg DeleteFormatByIdParams) error {
	_, err := q.db.Exec(ctx, deleteFormatById, arg.FormatID, arg.DeletedAt)
	return err
}

const getFormat = `-- name: GetFormat :one
SELECT format_id, updated_format_id, asesor_id, service_id, format_name, description, extension, version, created_at, updated_at, deleted_at FROM format
WHERE format_id = $1 LIMIT 1
`

func (q *Queries) GetFormat(ctx context.Context, formatID int32) (Format, error) {
	row := q.db.QueryRow(ctx, getFormat, formatID)
	var i Format
	err := row.Scan(
		&i.FormatID,
		&i.UpdatedFormatID,
		&i.AsesorID,
		&i.ServiceID,
		&i.FormatName,
		&i.Description,
		&i.Extension,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getFormatByName = `-- name: GetFormatByName :one
SELECT description, extension, version
FROM format
WHERE format_name = $1
`

type GetFormatByNameRow struct {
	Description string     `json:"description"`
	Extension   Extensions `json:"extension"`
	Version     string     `json:"version"`
}

func (q *Queries) GetFormatByName(ctx context.Context, formatName string) (GetFormatByNameRow, error) {
	row := q.db.QueryRow(ctx, getFormatByName, formatName)
	var i GetFormatByNameRow
	err := row.Scan(&i.Description, &i.Extension, &i.Version)
	return i, err
}

const listFormats = `-- name: ListFormats :many
SELECT format_id, updated_format_id, asesor_id, service_id, format_name, description, extension, version, created_at, updated_at, deleted_at FROM format
ORDER BY format_name
`

func (q *Queries) ListFormats(ctx context.Context) ([]Format, error) {
	rows, err := q.db.Query(ctx, listFormats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Format
	for rows.Next() {
		var i Format
		if err := rows.Scan(
			&i.FormatID,
			&i.UpdatedFormatID,
			&i.AsesorID,
			&i.ServiceID,
			&i.FormatName,
			&i.Description,
			&i.Extension,
			&i.Version,
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

const updateFormatById = `-- name: UpdateFormatById :exec
UPDATE format
SET format_name = $2, description = $3, extension=$4, version=$5, service_id=$6, updated_at=$7
WHERE format_id = $1
`

type UpdateFormatByIdParams struct {
	FormatID    int32            `json:"format_id"`
	FormatName  string           `json:"format_name"`
	Description string           `json:"description"`
	Extension   Extensions       `json:"extension"`
	Version     string           `json:"version"`
	ServiceID   pgtype.Int4      `json:"service_id"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateFormatById(ctx context.Context, arg UpdateFormatByIdParams) error {
	_, err := q.db.Exec(ctx, updateFormatById,
		arg.FormatID,
		arg.FormatName,
		arg.Description,
		arg.Extension,
		arg.Version,
		arg.ServiceID,
		arg.UpdatedAt,
	)
	return err
}
