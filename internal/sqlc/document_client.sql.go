// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: document_client.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createDocumentClient = `-- name: CreateDocumentClient :one
INSERT INTO document_client(client_id, document_id, action, created_at)
VALUES ($1, $2, $3, $4)
RETURNING document_client_id, client_id, document_id, action, created_at, updated_at, deleted_at
`

type CreateDocumentClientParams struct {
	ClientID   int32            `json:"client_id"`
	DocumentID int32            `json:"document_id"`
	Action     Action           `json:"action"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateDocumentClient(ctx context.Context, arg CreateDocumentClientParams) (DocumentClient, error) {
	row := q.db.QueryRow(ctx, createDocumentClient,
		arg.ClientID,
		arg.DocumentID,
		arg.Action,
		arg.CreatedAt,
	)
	var i DocumentClient
	err := row.Scan(
		&i.DocumentClientID,
		&i.ClientID,
		&i.DocumentID,
		&i.Action,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllDocumentClients = `-- name: DeleteAllDocumentClients :execresult
UPDATE document_client
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllDocumentClients(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllDocumentClients, deletedAt)
}

const deleteDocumentClientById = `-- name: DeleteDocumentClientById :exec
UPDATE document_client
SET deleted_at = $2
WHERE document_client_id = $1
`

type DeleteDocumentClientByIdParams struct {
	DocumentClientID int64            `json:"document_client_id"`
	DeletedAt        pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteDocumentClientById(ctx context.Context, arg DeleteDocumentClientByIdParams) error {
	_, err := q.db.Exec(ctx, deleteDocumentClientById, arg.DocumentClientID, arg.DeletedAt)
	return err
}

const getDocumentClient = `-- name: GetDocumentClient :one
SELECT document_client_id, client_id, document_id, action, created_at, updated_at, deleted_at FROM document_client
WHERE document_client_id = $1 LIMIT 1
`

func (q *Queries) GetDocumentClient(ctx context.Context, documentClientID int64) (DocumentClient, error) {
	row := q.db.QueryRow(ctx, getDocumentClient, documentClientID)
	var i DocumentClient
	err := row.Scan(
		&i.DocumentClientID,
		&i.ClientID,
		&i.DocumentID,
		&i.Action,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getDocumentClientByName = `-- name: GetDocumentClientByName :one
SELECT client_id, document_id, action
FROM document_client
WHERE document_client_id = $1
`

type GetDocumentClientByNameRow struct {
	ClientID   int32  `json:"client_id"`
	DocumentID int32  `json:"document_id"`
	Action     Action `json:"action"`
}

func (q *Queries) GetDocumentClientByName(ctx context.Context, documentClientID int64) (GetDocumentClientByNameRow, error) {
	row := q.db.QueryRow(ctx, getDocumentClientByName, documentClientID)
	var i GetDocumentClientByNameRow
	err := row.Scan(&i.ClientID, &i.DocumentID, &i.Action)
	return i, err
}

const listDocumentClients = `-- name: ListDocumentClients :many
SELECT document_client_id, client_id, document_id, action, created_at, updated_at, deleted_at FROM document_client
ORDER BY document_client_id
`

func (q *Queries) ListDocumentClients(ctx context.Context) ([]DocumentClient, error) {
	rows, err := q.db.Query(ctx, listDocumentClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DocumentClient
	for rows.Next() {
		var i DocumentClient
		if err := rows.Scan(
			&i.DocumentClientID,
			&i.ClientID,
			&i.DocumentID,
			&i.Action,
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
