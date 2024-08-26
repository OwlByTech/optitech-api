-- name: GetService :one
SELECT * FROM services
WHERE service_id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: ListServices :many
SELECT * FROM services
ORDER BY service_id;

-- name: GetServicesByName :one
SELECT name
FROM services
WHERE service_id = $1;

-- name: CreateServices :one
INSERT INTO services(name, created_at)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateService :exec
UPDATE services
SET name = $2, updated_at = $3
WHERE service_id = $1;

-- name: DeleteService :exec
UPDATE services
SET deleted_at = $2
WHERE service_id = $1;

-- name: DeleteAllServicess :execresult
UPDATE services
SET deleted_at = $1
WHERE deleted_at IS NULL;


