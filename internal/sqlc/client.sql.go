// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: client.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createClient = `-- name: CreateClient :one
INSERT INTO client (given_name, surname, email, password, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING client_id, given_name, surname, photo, email, password, status, created_at, updated_at, deleted_at
`

type CreateClientParams struct {
	GivenName string           `json:"given_name"`
	Surname   string           `json:"surname"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (Client, error) {
	row := q.db.QueryRow(ctx, createClient,
		arg.GivenName,
		arg.Surname,
		arg.Email,
		arg.Password,
		arg.CreatedAt,
	)
	var i Client
	err := row.Scan(
		&i.ClientID,
		&i.GivenName,
		&i.Surname,
		&i.Photo,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllClients = `-- name: DeleteAllClients :execresult
UPDATE client
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllClients(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllClients, deletedAt)
}

const deleteClientById = `-- name: DeleteClientById :exec
UPDATE client
SET deleted_at = $2
WHERE client_id = $1
`

type DeleteClientByIdParams struct {
	ClientID  int32            `json:"client_id"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteClientById(ctx context.Context, arg DeleteClientByIdParams) error {
	_, err := q.db.Exec(ctx, deleteClientById, arg.ClientID, arg.DeletedAt)
	return err
}

const getClient = `-- name: GetClient :one
SELECT client_id, given_name, surname, photo, email, password, status, created_at, updated_at, deleted_at FROM client
WHERE client_id = $1 LIMIT 1
`

func (q *Queries) GetClient(ctx context.Context, clientID int32) (Client, error) {
	row := q.db.QueryRow(ctx, getClient, clientID)
	var i Client
	err := row.Scan(
		&i.ClientID,
		&i.GivenName,
		&i.Surname,
		&i.Photo,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getClientByEmail = `-- name: GetClientByEmail :one
SELECT client_id, given_name, surname, photo, email, password, status, created_at, updated_at, deleted_at FROM client
WHERE email = $1
`

func (q *Queries) GetClientByEmail(ctx context.Context, email string) (Client, error) {
	row := q.db.QueryRow(ctx, getClientByEmail, email)
	var i Client
	err := row.Scan(
		&i.ClientID,
		&i.GivenName,
		&i.Surname,
		&i.Photo,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getClientPhoto = `-- name: GetClientPhoto :one
SELECT photo FROM client
WHERE client_id = $1
`

func (q *Queries) GetClientPhoto(ctx context.Context, clientID int32) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, getClientPhoto, clientID)
	var photo pgtype.Text
	err := row.Scan(&photo)
	return photo, err
}

const listClients = `-- name: ListClients :many
SELECT client_id, given_name, surname, photo, email, password, status, created_at, updated_at, deleted_at FROM client
ORDER BY given_name
`

func (q *Queries) ListClients(ctx context.Context) ([]Client, error) {
	rows, err := q.db.Query(ctx, listClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Client
	for rows.Next() {
		var i Client
		if err := rows.Scan(
			&i.ClientID,
			&i.GivenName,
			&i.Surname,
			&i.Photo,
			&i.Email,
			&i.Password,
			&i.Status,
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

const loginClient = `-- name: LoginClient :one
SELECT client_id, given_name, surname, photo, email, password, status, created_at, updated_at, deleted_at FROM client
WHERE password = $1 AND email= $2 LIMIT 1
`

type LoginClientParams struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (q *Queries) LoginClient(ctx context.Context, arg LoginClientParams) (Client, error) {
	row := q.db.QueryRow(ctx, loginClient, arg.Password, arg.Email)
	var i Client
	err := row.Scan(
		&i.ClientID,
		&i.GivenName,
		&i.Surname,
		&i.Photo,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateClientById = `-- name: UpdateClientById :exec
UPDATE client
SET  given_name = $2, password = $3, surname = $4, email = $5, updated_at = $6
WHERE client_id = $1
`

type UpdateClientByIdParams struct {
	ClientID  int32            `json:"client_id"`
	GivenName string           `json:"given_name"`
	Password  string           `json:"password"`
	Surname   string           `json:"surname"`
	Email     string           `json:"email"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateClientById(ctx context.Context, arg UpdateClientByIdParams) error {
	_, err := q.db.Exec(ctx, updateClientById,
		arg.ClientID,
		arg.GivenName,
		arg.Password,
		arg.Surname,
		arg.Email,
		arg.UpdatedAt,
	)
	return err
}

const updateClientPhoto = `-- name: UpdateClientPhoto :exec
UPDATE client
SET  photo = $2, updated_at = $3
WHERE client_id = $1
`

type UpdateClientPhotoParams struct {
	ClientID  int32            `json:"client_id"`
	Photo     pgtype.Text      `json:"photo"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateClientPhoto(ctx context.Context, arg UpdateClientPhotoParams) error {
	_, err := q.db.Exec(ctx, updateClientPhoto, arg.ClientID, arg.Photo, arg.UpdatedAt)
	return err
}

const updateClientStatusById = `-- name: UpdateClientStatusById :exec
UPDATE client
SET status = $2, updated_at = $3
WHERE client_id = $1
`

type UpdateClientStatusByIdParams struct {
	ClientID  int32            `json:"client_id"`
	Status    StatusClient     `json:"status"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateClientStatusById(ctx context.Context, arg UpdateClientStatusByIdParams) error {
	_, err := q.db.Exec(ctx, updateClientStatusById, arg.ClientID, arg.Status, arg.UpdatedAt)
	return err
}
