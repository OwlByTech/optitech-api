-- name: GetDocument :one
SELECT * FROM document
WHERE document_id = $1 LIMIT 1;

-- name: ListDocuments :many
SELECT * FROM document
ORDER BY document_id;

-- name: GetDocumentByName :one
SELECT format_id, institution_id, client_id, file_rute, status
FROM document
WHERE document_id = $1;

-- name: CreateDocument :one
INSERT INTO document(format_id, institution_id, client_id, file_rute, status, create_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateDocumentById :exec
UPDATE document
SET file_rute = $2, status = $3, update_at = $4
WHERE document_id = $1;

-- name: DeleteDocument :exec
DELETE FROM document
WHERE document_id = $1;

-- name: DeleteAllDocuments :execresult
DELETE FROM document;
