-- name: GetServices :one
SELECT * FROM services
WHERE services_id = $1 LIMIT 1;

-- name: ListServicess :many
SELECT * FROM services
ORDER BY services_id;

-- name: GetServicesByName :one
SELECT service_name
FROM services
WHERE services_id = $1;

-- name: CreateServices :one
INSERT INTO services(service_name, created_at)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateServicesById :exec
UPDATE services
SET service_name = $2, updated_at = $3
WHERE services_id = $1;

-- name: DeleteServicesById :exec
UPDATE services
SET deleted_at = $2
WHERE services_id = $1;

-- name: DeleteAllServicess :execresult
UPDATE services
SET deleted_at = $1
WHERE deleted_at IS NULL;


