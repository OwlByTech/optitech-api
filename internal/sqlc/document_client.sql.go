// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: document_client.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createDocumentClient = `-- name: CreateDocumentClient :one
INSERT INTO document_client(client_id, document_id, action, create_at)
VALUES ($1, $2, $3, $4)
RETURNING document_client_id, client_id, document_id, action, create_at
`

type CreateDocumentClientParams struct {
	ClientID   int32     `json:"client_id"`
	DocumentID int32     `json:"document_id"`
	Action     Action    `json:"action"`
	CreateAt   time.Time `json:"create_at"`
}

func (q *Queries) CreateDocumentClient(ctx context.Context, arg CreateDocumentClientParams) (DocumentClient, error) {
	row := q.db.QueryRowContext(ctx, createDocumentClient,
		arg.ClientID,
		arg.DocumentID,
		arg.Action,
		arg.CreateAt,
	)
	var i DocumentClient
	err := row.Scan(
		&i.DocumentClientID,
		&i.ClientID,
		&i.DocumentID,
		&i.Action,
		&i.CreateAt,
	)
	return i, err
}

const deleteAllDocumentClients = `-- name: DeleteAllDocumentClients :execresult
DELETE FROM document_client
`

func (q *Queries) DeleteAllDocumentClients(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteAllDocumentClients)
}

const deleteDocumentClient = `-- name: DeleteDocumentClient :exec
DELETE FROM document_client
WHERE document_client_id = $1
`

func (q *Queries) DeleteDocumentClient(ctx context.Context, documentClientID int64) error {
	_, err := q.db.ExecContext(ctx, deleteDocumentClient, documentClientID)
	return err
}

const getDocumentClient = `-- name: GetDocumentClient :one
SELECT document_client_id, client_id, document_id, action, create_at FROM document_client
WHERE document_client_id = $1 LIMIT 1
`

func (q *Queries) GetDocumentClient(ctx context.Context, documentClientID int64) (DocumentClient, error) {
	row := q.db.QueryRowContext(ctx, getDocumentClient, documentClientID)
	var i DocumentClient
	err := row.Scan(
		&i.DocumentClientID,
		&i.ClientID,
		&i.DocumentID,
		&i.Action,
		&i.CreateAt,
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
	row := q.db.QueryRowContext(ctx, getDocumentClientByName, documentClientID)
	var i GetDocumentClientByNameRow
	err := row.Scan(&i.ClientID, &i.DocumentID, &i.Action)
	return i, err
}

const listDocumentClients = `-- name: ListDocumentClients :many
SELECT document_client_id, client_id, document_id, action, create_at FROM document_client
ORDER BY document_client_id
`

func (q *Queries) ListDocumentClients(ctx context.Context) ([]DocumentClient, error) {
	rows, err := q.db.QueryContext(ctx, listDocumentClients)
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
			&i.CreateAt,
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
