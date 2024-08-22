// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: asesor.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createAsesor = `-- name: CreateAsesor :one
INSERT INTO asesor (asesor_id, about, created_at)
VALUES ($1, $2, $3)
RETURNING asesor_id, about, created_at, updated_at, deleted_at
`

type CreateAsesorParams struct {
	AsesorID  int32            `json:"asesor_id"`
	About     pgtype.Text      `json:"about"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateAsesor(ctx context.Context, arg CreateAsesorParams) (Asesor, error) {
	row := q.db.QueryRow(ctx, createAsesor, arg.AsesorID, arg.About, arg.CreatedAt)
	var i Asesor
	err := row.Scan(
		&i.AsesorID,
		&i.About,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllAsesors = `-- name: DeleteAllAsesors :execresult
UPDATE asesor
SET deleted_at = $1
WHERE deleted_at is NULL
`

func (q *Queries) DeleteAllAsesors(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllAsesors, deletedAt)
}

const deleteAsesorById = `-- name: DeleteAsesorById :exec
UPDATE asesor
SET deleted_at = $2
WHERE asesor_id = $1
`

type DeleteAsesorByIdParams struct {
	AsesorID  int32            `json:"asesor_id"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteAsesorById(ctx context.Context, arg DeleteAsesorByIdParams) error {
	_, err := q.db.Exec(ctx, deleteAsesorById, arg.AsesorID, arg.DeletedAt)
	return err
}

const getAsesor = `-- name: GetAsesor :one
SELECT asesor_id, about, created_at, updated_at, deleted_at FROM asesor
WHERE asesor_id = $1 LIMIT 1
`

func (q *Queries) GetAsesor(ctx context.Context, asesorID int32) (Asesor, error) {
	row := q.db.QueryRow(ctx, getAsesor, asesorID)
	var i Asesor
	err := row.Scan(
		&i.AsesorID,
		&i.About,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listAsesors = `-- name: ListAsesors :many
SELECT asesor_id, about, created_at, updated_at, deleted_at FROM asesor
ORDER BY asesor_id
`

func (q *Queries) ListAsesors(ctx context.Context) ([]Asesor, error) {
	rows, err := q.db.Query(ctx, listAsesors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Asesor
	for rows.Next() {
		var i Asesor
		if err := rows.Scan(
			&i.AsesorID,
			&i.About,
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

const updateAsesorById = `-- name: UpdateAsesorById :exec
UPDATE asesor
SET   about = $2, updated_at = $3
WHERE asesor_id = $1
`

type UpdateAsesorByIdParams struct {
	AsesorID  int32            `json:"asesor_id"`
	About     pgtype.Text      `json:"about"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateAsesorById(ctx context.Context, arg UpdateAsesorByIdParams) error {
	_, err := q.db.Exec(ctx, updateAsesorById, arg.AsesorID, arg.About, arg.UpdatedAt)
	return err
}
