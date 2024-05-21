-- name: GetAsesor :one
SELECT * FROM asesor
WHERE asesor_id = $1 LIMIT 1;

-- name: ListAsesors :many
SELECT * FROM asesor
ORDER BY username;

-- name: GetAsesorByUsername :one
SELECT username photo, about
FROM asesor
WHERE username = $1;

-- name: CreateAsesor :one
INSERT INTO asesor (username, photo, about, create_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateAsesorById :exec
UPDATE asesor
SET username = $2, photo = $3, about = $4, update_at = $5
WHERE asesor_id = $1;

-- name: DeleteAsesor :exec
DELETE FROM asesor
WHERE asesor_id = $1;

-- name: DeleteAllAsesors :execresult
DELETE FROM asesor;
