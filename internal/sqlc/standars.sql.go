// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: standars.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createStandard = `-- name: CreateStandard :one
INSERT INTO standards(service_id, name, complexity, modality, article, section, paragraph, criteria, comply, applys, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING standard_id, service_id, name, complexity, modality, article, section, paragraph, criteria, comply, applys, created_at, updated_at, deleted_at
`

type CreateStandardParams struct {
	ServiceID  int32            `json:"service_id"`
	Name       string           `json:"name"`
	Complexity pgtype.Text      `json:"complexity"`
	Modality   string           `json:"modality"`
	Article    string           `json:"article"`
	Section    string           `json:"section"`
	Paragraph  pgtype.Text      `json:"paragraph"`
	Criteria   string           `json:"criteria"`
	Comply     pgtype.Bool      `json:"comply"`
	Applys     pgtype.Bool      `json:"applys"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateStandard(ctx context.Context, arg CreateStandardParams) (Standard, error) {
	row := q.db.QueryRow(ctx, createStandard,
		arg.ServiceID,
		arg.Name,
		arg.Complexity,
		arg.Modality,
		arg.Article,
		arg.Section,
		arg.Paragraph,
		arg.Criteria,
		arg.Comply,
		arg.Applys,
		arg.CreatedAt,
	)
	var i Standard
	err := row.Scan(
		&i.StandardID,
		&i.ServiceID,
		&i.Name,
		&i.Complexity,
		&i.Modality,
		&i.Article,
		&i.Section,
		&i.Paragraph,
		&i.Criteria,
		&i.Comply,
		&i.Applys,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllStandards = `-- name: DeleteAllStandards :execresult
UPDATE standards
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllStandards(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllStandards, deletedAt)
}

const deleteStandardById = `-- name: DeleteStandardById :exec
UPDATE standards
SET name = $2, complexity = $3, modality =$4, article = $5, section = $6, paragraph = $7, criteria = $8, comply = $9, applys = $10, updated_at = $11
WHERE standard_id = $1
`

type DeleteStandardByIdParams struct {
	StandardID int32            `json:"standard_id"`
	Name       string           `json:"name"`
	Complexity pgtype.Text      `json:"complexity"`
	Modality   string           `json:"modality"`
	Article    string           `json:"article"`
	Section    string           `json:"section"`
	Paragraph  pgtype.Text      `json:"paragraph"`
	Criteria   string           `json:"criteria"`
	Comply     pgtype.Bool      `json:"comply"`
	Applys     pgtype.Bool      `json:"applys"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) DeleteStandardById(ctx context.Context, arg DeleteStandardByIdParams) error {
	_, err := q.db.Exec(ctx, deleteStandardById,
		arg.StandardID,
		arg.Name,
		arg.Complexity,
		arg.Modality,
		arg.Article,
		arg.Section,
		arg.Paragraph,
		arg.Criteria,
		arg.Comply,
		arg.Applys,
		arg.UpdatedAt,
	)
	return err
}

const getStandard = `-- name: GetStandard :one
SELECT standard_id, service_id, name, complexity, modality, article, section, paragraph, criteria, comply, applys, created_at, updated_at, deleted_at FROM standards
WHERE standard_id = $1 LIMIT 1
`

func (q *Queries) GetStandard(ctx context.Context, standardID int32) (Standard, error) {
	row := q.db.QueryRow(ctx, getStandard, standardID)
	var i Standard
	err := row.Scan(
		&i.StandardID,
		&i.ServiceID,
		&i.Name,
		&i.Complexity,
		&i.Modality,
		&i.Article,
		&i.Section,
		&i.Paragraph,
		&i.Criteria,
		&i.Comply,
		&i.Applys,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getStandardByName = `-- name: GetStandardByName :one
SELECT name, complexity, modality, article, section, paragraph, criteria, comply, applys
FROM standards
WHERE standard_id = $1
`

type GetStandardByNameRow struct {
	Name       string      `json:"name"`
	Complexity pgtype.Text `json:"complexity"`
	Modality   string      `json:"modality"`
	Article    string      `json:"article"`
	Section    string      `json:"section"`
	Paragraph  pgtype.Text `json:"paragraph"`
	Criteria   string      `json:"criteria"`
	Comply     pgtype.Bool `json:"comply"`
	Applys     pgtype.Bool `json:"applys"`
}

func (q *Queries) GetStandardByName(ctx context.Context, standardID int32) (GetStandardByNameRow, error) {
	row := q.db.QueryRow(ctx, getStandardByName, standardID)
	var i GetStandardByNameRow
	err := row.Scan(
		&i.Name,
		&i.Complexity,
		&i.Modality,
		&i.Article,
		&i.Section,
		&i.Paragraph,
		&i.Criteria,
		&i.Comply,
		&i.Applys,
	)
	return i, err
}

const listStandards = `-- name: ListStandards :many
SELECT standard_id, service_id, name, complexity, modality, article, section, paragraph, criteria, comply, applys, created_at, updated_at, deleted_at FROM standards
ORDER BY standard_id
`

func (q *Queries) ListStandards(ctx context.Context) ([]Standard, error) {
	rows, err := q.db.Query(ctx, listStandards)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Standard
	for rows.Next() {
		var i Standard
		if err := rows.Scan(
			&i.StandardID,
			&i.ServiceID,
			&i.Name,
			&i.Complexity,
			&i.Modality,
			&i.Article,
			&i.Section,
			&i.Paragraph,
			&i.Criteria,
			&i.Comply,
			&i.Applys,
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

const updateStandardById = `-- name: UpdateStandardById :exec
UPDATE standards
SET deleted_at = $2
WHERE standard_id = $1
`

type UpdateStandardByIdParams struct {
	StandardID int32            `json:"standard_id"`
	DeletedAt  pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) UpdateStandardById(ctx context.Context, arg UpdateStandardByIdParams) error {
	_, err := q.db.Exec(ctx, updateStandardById, arg.StandardID, arg.DeletedAt)
	return err
}
