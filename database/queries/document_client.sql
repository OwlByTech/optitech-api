-- name: GetDocumentClient :one
SELECT * FROM document_client
WHERE document_client_id = $1 LIMIT 1;

-- name: ListDocumentClients :many
SELECT * FROM document_client
ORDER BY document_client_id;

-- name: GetDocumentClientByName :one
SELECT client_id, document_id, action
FROM document_client
WHERE document_client_id = $1;

-- name: CreateDocumentClient :one
INSERT INTO document_client(client_id, document_id, action, create_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteDocumentClient :exec
DELETE FROM document_client
WHERE document_client_id = $1;

-- name: DeleteAllDocumentClients :execresult
DELETE FROM document_client;
