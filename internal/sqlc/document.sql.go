// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: document.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createDocument = `-- name: CreateDocument :one
INSERT INTO document(directory_id,name, format_id, file_rute, status, created_at)
VALUES ($1, $2, $3, $4, $5,$6)
RETURNING document_id, directory_id, format_id, name, file_rute, status, created_at, updated_at, deleted_at
`

type CreateDocumentParams struct {
	DirectoryID int32            `json:"directory_id"`
	Name        string           `json:"name"`
	FormatID    pgtype.Int4      `json:"format_id"`
	FileRute    string           `json:"file_rute"`
	Status      Status           `json:"status"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateDocument(ctx context.Context, arg CreateDocumentParams) (Document, error) {
	row := q.db.QueryRow(ctx, createDocument,
		arg.DirectoryID,
		arg.Name,
		arg.FormatID,
		arg.FileRute,
		arg.Status,
		arg.CreatedAt,
	)
	var i Document
	err := row.Scan(
		&i.DocumentID,
		&i.DirectoryID,
		&i.FormatID,
		&i.Name,
		&i.FileRute,
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

func (q *Queries) DeleteAllDocuments(ctx context.Context, deletedAt pgtype.Timestamp) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAllDocuments, deletedAt)
}

const deleteDocumentById = `-- name: DeleteDocumentById :exec
UPDATE document
SET deleted_at = $2
WHERE document_id = $1
`

type DeleteDocumentByIdParams struct {
	DocumentID int64            `json:"document_id"`
	DeletedAt  pgtype.Timestamp `json:"deleted_at"`
}

func (q *Queries) DeleteDocumentById(ctx context.Context, arg DeleteDocumentByIdParams) error {
	_, err := q.db.Exec(ctx, deleteDocumentById, arg.DocumentID, arg.DeletedAt)
	return err
}

const getDocument = `-- name: GetDocument :one
SELECT document_id, directory_id, format_id, name, file_rute, status, created_at, updated_at, deleted_at FROM document
WHERE document_id = $1 AND deleted_at is null
LIMIT 1
`

func (q *Queries) GetDocument(ctx context.Context, documentID int64) (Document, error) {
	row := q.db.QueryRow(ctx, getDocument, documentID)
	var i Document
	err := row.Scan(
		&i.DocumentID,
		&i.DirectoryID,
		&i.FormatID,
		&i.Name,
		&i.FileRute,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getDocumentByName = `-- name: GetDocumentByName :one
SELECT directory_id, format_id, file_rute, status
FROM document
WHERE document_id = $1
`

type GetDocumentByNameRow struct {
	DirectoryID int32       `json:"directory_id"`
	FormatID    pgtype.Int4 `json:"format_id"`
	FileRute    string      `json:"file_rute"`
	Status      Status      `json:"status"`
}

func (q *Queries) GetDocumentByName(ctx context.Context, documentID int64) (GetDocumentByNameRow, error) {
	row := q.db.QueryRow(ctx, getDocumentByName, documentID)
	var i GetDocumentByNameRow
	err := row.Scan(
		&i.DirectoryID,
		&i.FormatID,
		&i.FileRute,
		&i.Status,
	)
	return i, err
}

const listDocuments = `-- name: ListDocuments :many
SELECT document_id, directory_id, format_id, name, file_rute, status, created_at, updated_at, deleted_at FROM document
ORDER BY document_id
`

func (q *Queries) ListDocuments(ctx context.Context) ([]Document, error) {
	rows, err := q.db.Query(ctx, listDocuments)
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
			&i.Name,
			&i.FileRute,
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

const listDocumentsByDirectory = `-- name: ListDocumentsByDirectory :many
SELECT document_id, directory_id, format_id, name, file_rute, status, created_at, updated_at, deleted_at FROM document
WHERE directory_id= $1
ORDER BY document_id
`

func (q *Queries) ListDocumentsByDirectory(ctx context.Context, directoryID int32) ([]Document, error) {
	rows, err := q.db.Query(ctx, listDocumentsByDirectory, directoryID)
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
			&i.Name,
			&i.FileRute,
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

const updateDocumentById = `-- name: UpdateDocumentById :exec
UPDATE document
SET  directory_id = $2, format_id = $3, file_rute = $4, status = $5, updated_at = $6
WHERE document_id = $1
`

type UpdateDocumentByIdParams struct {
	DocumentID  int64            `json:"document_id"`
	DirectoryID int32            `json:"directory_id"`
	FormatID    pgtype.Int4      `json:"format_id"`
	FileRute    string           `json:"file_rute"`
	Status      Status           `json:"status"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateDocumentById(ctx context.Context, arg UpdateDocumentByIdParams) error {
	_, err := q.db.Exec(ctx, updateDocumentById,
		arg.DocumentID,
		arg.DirectoryID,
		arg.FormatID,
		arg.FileRute,
		arg.Status,
		arg.UpdatedAt,
	)
	return err
}
