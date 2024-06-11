// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: document.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createDocument = `-- name: CreateDocument :one
INSERT INTO document(directory_id, format_id, url, status, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING document_id, directory_id, format_id, url, status, created_at, updated_at, deleted_at
`

type CreateDocumentParams struct {
	DirectoryID int32         `json:"directory_id"`
	FormatID    sql.NullInt32 `json:"format_id"`
	Url         string        `json:"url"`
	Status      Status        `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
}

func (q *Queries) CreateDocument(ctx context.Context, arg CreateDocumentParams) (Document, error) {
	row := q.db.QueryRowContext(ctx, createDocument,
		arg.DirectoryID,
		arg.FormatID,
		arg.Url,
		arg.Status,
		arg.CreatedAt,
	)
	var i Document
	err := row.Scan(
		&i.DocumentID,
		&i.DirectoryID,
		&i.FormatID,
		&i.Url,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllDocuments = `-- name: DeleteAllDocuments :execresult
UPDATE document
SET deleted_at = $1
WHERE deleted_at IS NULL
`

func (q *Queries) DeleteAllDocuments(ctx context.Context, deletedAt sql.NullTime) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteAllDocuments, deletedAt)
}

const deleteDocumentById = `-- name: DeleteDocumentById :exec
UPDATE document
SET deleted_at = $2
WHERE document_id = $1
`

type DeleteDocumentByIdParams struct {
	DocumentID int64        `json:"document_id"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
}

func (q *Queries) DeleteDocumentById(ctx context.Context, arg DeleteDocumentByIdParams) error {
	_, err := q.db.ExecContext(ctx, deleteDocumentById, arg.DocumentID, arg.DeletedAt)
	return err
}

const getDocument = `-- name: GetDocument :one
SELECT document_id, directory_id, format_id, url, status, created_at, updated_at, deleted_at FROM document
WHERE document_id = $1 LIMIT 1
`

func (q *Queries) GetDocument(ctx context.Context, documentID int64) (Document, error) {
	row := q.db.QueryRowContext(ctx, getDocument, documentID)
	var i Document
	err := row.Scan(
		&i.DocumentID,
		&i.DirectoryID,
		&i.FormatID,
		&i.Url,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getDocumentByName = `-- name: GetDocumentByName :one
SELECT directory_id, format_id, url, status
FROM document
WHERE document_id = $1
`

type GetDocumentByNameRow struct {
	DirectoryID int32         `json:"directory_id"`
	FormatID    sql.NullInt32 `json:"format_id"`
	Url         string        `json:"url"`
	Status      Status        `json:"status"`
}

func (q *Queries) GetDocumentByName(ctx context.Context, documentID int64) (GetDocumentByNameRow, error) {
	row := q.db.QueryRowContext(ctx, getDocumentByName, documentID)
	var i GetDocumentByNameRow
	err := row.Scan(
		&i.DirectoryID,
		&i.FormatID,
		&i.Url,
		&i.Status,
	)
	return i, err
}

const listDocuments = `-- name: ListDocuments :many
SELECT document_id, directory_id, format_id, url, status, created_at, updated_at, deleted_at FROM document
ORDER BY document_id
`

func (q *Queries) ListDocuments(ctx context.Context) ([]Document, error) {
	rows, err := q.db.QueryContext(ctx, listDocuments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Document
	for rows.Next() {
		var i Document
		if err := rows.Scan(
			&i.DocumentID,
			&i.DirectoryID,
			&i.FormatID,
			&i.Url,
			&i.Status,
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

const updateDocumentById = `-- name: UpdateDocumentById :exec
UPDATE document
SET  directory_id = $2, format_id = $3, url = $4, status = $5, updated_at = $6
WHERE document_id = $1
`

type UpdateDocumentByIdParams struct {
	DocumentID  int64         `json:"document_id"`
	DirectoryID int32         `json:"directory_id"`
	FormatID    sql.NullInt32 `json:"format_id"`
	Url         string        `json:"url"`
	Status      Status        `json:"status"`
	UpdatedAt   sql.NullTime  `json:"updated_at"`
}

func (q *Queries) UpdateDocumentById(ctx context.Context, arg UpdateDocumentByIdParams) error {
	_, err := q.db.ExecContext(ctx, updateDocumentById,
		arg.DocumentID,
		arg.DirectoryID,
		arg.FormatID,
		arg.Url,
		arg.Status,
		arg.UpdatedAt,
	)
	return err
}