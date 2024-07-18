-- name: GetDocument :one
SELECT * FROM document
WHERE document_id = $1 AND deleted_at is null
LIMIT 1;

-- name: ListDocuments :many
SELECT * FROM document
ORDER BY document_id;


-- name: ListDocumentsByDirectory :many
SELECT * FROM document
WHERE directory_id= $1
ORDER BY document_id;


-- name: GetDocumentByName :one
SELECT directory_id, format_id, file_rute, status
FROM document
WHERE document_id = $1;

-- name: CreateDocument :one
INSERT INTO document(directory_id,name, format_id, file_rute, status, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateDocumentById :exec
UPDATE document
SET  directory_id = $2, format_id = $3, file_rute = $4, status = $5, updated_at = $6
WHERE document_id = $1;

-- name: UpdateDocumentNameById :exec
UPDATE document
SET name = $2, file_rute = $3, updated_at = $4
WHERE document_id = $1;

-- name: DeleteDocumentById :exec
UPDATE document
SET deleted_at = $2
WHERE document_id = $1;

-- name: DeleteAllDocuments :execresult
UPDATE document
SET deleted_at = $1
WHERE deleted_at IS NULL;

-- name: ExistsDocument :one
SELECT COUNT(*) > 0
FROM document
WHERE document_id = $1 AND deleted_at IS NOT NULL
LIMIT 1;

-- name: ExistEndpoint :one
SELECT COUNT(file_rute) > 0
FROM document
WHERE file_rute = $1
LIMIT 1;
