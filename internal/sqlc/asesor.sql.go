// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: asesor.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createAsesor = `-- name: CreateAsesor :one
INSERT INTO asesor (client_id, username, photo, about, create_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING asesor_id, client_id, username, photo, about, create_at, update_at
`

type CreateAsesorParams struct {
	ClientID int32     `json:"client_id"`
	Username string    `json:"username"`
	Photo    string    `json:"photo"`
	About    string    `json:"about"`
	CreateAt time.Time `json:"create_at"`
}

func (q *Queries) CreateAsesor(ctx context.Context, arg CreateAsesorParams) (Asesor, error) {
	row := q.db.QueryRowContext(ctx, createAsesor,
		arg.ClientID,
		arg.Username,
		arg.Photo,
		arg.About,
		arg.CreateAt,
	)
	var i Asesor
	err := row.Scan(
		&i.AsesorID,
		&i.ClientID,
		&i.Username,
		&i.Photo,
		&i.About,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteAllAsesors = `-- name: DeleteAllAsesors :execresult
DELETE FROM asesor
`

func (q *Queries) DeleteAllAsesors(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteAllAsesors)
}

const deleteAsesor = `-- name: DeleteAsesor :exec
DELETE FROM asesor
WHERE asesor_id = $1
`

func (q *Queries) DeleteAsesor(ctx context.Context, asesorID int64) error {
	_, err := q.db.ExecContext(ctx, deleteAsesor, asesorID)
	return err
}

const getAsesor = `-- name: GetAsesor :one
SELECT asesor_id, client_id, username, photo, about, create_at, update_at FROM asesor
WHERE asesor_id = $1 LIMIT 1
`

func (q *Queries) GetAsesor(ctx context.Context, asesorID int64) (Asesor, error) {
	row := q.db.QueryRowContext(ctx, getAsesor, asesorID)
	var i Asesor
	err := row.Scan(
		&i.AsesorID,
		&i.ClientID,
		&i.Username,
		&i.Photo,
		&i.About,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const getAsesorByUsername = `-- name: GetAsesorByUsername :one
SELECT username photo, about
FROM asesor
WHERE username = $1
`

type GetAsesorByUsernameRow struct {
	Photo string `json:"photo"`
	About string `json:"about"`
}

func (q *Queries) GetAsesorByUsername(ctx context.Context, username string) (GetAsesorByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getAsesorByUsername, username)
	var i GetAsesorByUsernameRow
	err := row.Scan(&i.Photo, &i.About)
	return i, err
}

const listAsesors = `-- name: ListAsesors :many
SELECT asesor_id, client_id, username, photo, about, create_at, update_at FROM asesor
ORDER BY username
`

func (q *Queries) ListAsesors(ctx context.Context) ([]Asesor, error) {
	rows, err := q.db.QueryContext(ctx, listAsesors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Asesor
	for rows.Next() {
		var i Asesor
		if err := rows.Scan(
			&i.AsesorID,
			&i.ClientID,
			&i.Username,
			&i.Photo,
			&i.About,
			&i.CreateAt,
			&i.UpdateAt,
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

const updateAsesorById = `-- name: UpdateAsesorById :exec
UPDATE asesor
SET username = $2, photo = $3, about = $4, update_at = $5
WHERE asesor_id = $1
`

type UpdateAsesorByIdParams struct {
	AsesorID int64        `json:"asesor_id"`
	Username string       `json:"username"`
	Photo    string       `json:"photo"`
	About    string       `json:"about"`
	UpdateAt sql.NullTime `json:"update_at"`
}

func (q *Queries) UpdateAsesorById(ctx context.Context, arg UpdateAsesorByIdParams) error {
	_, err := q.db.ExecContext(ctx, updateAsesorById,
		arg.AsesorID,
		arg.Username,
		arg.Photo,
		arg.About,
		arg.UpdateAt,
	)
	return err
}
