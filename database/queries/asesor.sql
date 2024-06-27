-- name: GetAsesor :one
SELECT * FROM asesor
WHERE asesor_id = $1 LIMIT 1;

-- name: ListAsesors :many
SELECT * FROM asesor
ORDER BY asesor_id;

-- name: CreateAsesor :one
INSERT INTO asesor (asesor_id, about, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateAsesorById :exec
UPDATE asesor
SET   about = $2, updated_at = $3
WHERE asesor_id = $1;

-- name: DeleteAsesorById :exec
UPDATE asesor
SET deleted_at = $2
WHERE asesor_id = $1;

-- name: DeleteAllAsesors :execresult
UPDATE asesor
SET deleted_at = $1
WHERE deleted_at is NULL;
