-- name: GetDocument :one
SELECT * FROM document
WHERE document_id = $1 LIMIT 1;

-- name: ListDocuments :many
SELECT * FROM document
ORDER BY document_id;

-- name: GetDocumentByName :one
SELECT directory_id, format_id, url, status
FROM document
WHERE document_id = $1;

-- name: CreateDocument :one
INSERT INTO document(directory_id, format_id, url, status, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateDocumentById :exec
UPDATE document
SET  directory_id = $2 format_id = $3 url = $4, status = $5, updated_at = $6
WHERE document_id = $1;

-- name: DeleteDocumentById :exec
UPDATE document
SET deleted_at = $2
WHERE document_id = $1

-- name: DeleteAllDocuments :execresult
UPDATE document;
SET deleted_at = $1
WHERE document_id IS NULL;
