-- name: GetNotification :one
SELECT * FROM notification
WHERE notification_id = $1 AND to_id = $2 LIMIT 1;

-- name: ListNotifications :many
SELECT * FROM notification
ORDER BY created_at;

-- name: ListNotitifacionByType :many
SELECT * FROM notification
ORDER BY type;

-- name: CreateNofication :one
INSERT INTO notification ("from", "to", from_id, to_id, message, title, payload, type, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateNotificationVisualized :exec
UPDATE notification
SET visualized = $2
WHERE notification_id = $1;