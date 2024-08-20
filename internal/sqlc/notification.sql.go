// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: notification.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createNofication = `-- name: CreateNofication :one
INSERT INTO notification ("from", "to", from_id, to_id, message, title, payload, type, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING notification_id, "from", "to", from_id, to_id, message, title, visualized, payload, type, created_at
`

type CreateNoficationParams struct {
	From      FromNotification     `json:"from"`
	To        ToNotification       `json:"to"`
	FromID    int32                `json:"from_id"`
	ToID      int32                `json:"to_id"`
	Message   string               `json:"message"`
	Title     string               `json:"title"`
	Payload   []byte               `json:"payload"`
	Type      NullTypeNotification `json:"type"`
	CreatedAt pgtype.Timestamp     `json:"created_at"`
}

func (q *Queries) CreateNofication(ctx context.Context, arg CreateNoficationParams) (Notification, error) {
	row := q.db.QueryRow(ctx, createNofication,
		arg.From,
		arg.To,
		arg.FromID,
		arg.ToID,
		arg.Message,
		arg.Title,
		arg.Payload,
		arg.Type,
		arg.CreatedAt,
	)
	var i Notification
	err := row.Scan(
		&i.NotificationID,
		&i.From,
		&i.To,
		&i.FromID,
		&i.ToID,
		&i.Message,
		&i.Title,
		&i.Visualized,
		&i.Payload,
		&i.Type,
		&i.CreatedAt,
	)
	return i, err
}

const getNotification = `-- name: GetNotification :one
SELECT notification_id, "from", "to", from_id, to_id, message, title, visualized, payload, type, created_at FROM notification
WHERE notification_id = $1 LIMIT 1
`

func (q *Queries) GetNotification(ctx context.Context, notificationID int64) (Notification, error) {
	row := q.db.QueryRow(ctx, getNotification, notificationID)
	var i Notification
	err := row.Scan(
		&i.NotificationID,
		&i.From,
		&i.To,
		&i.FromID,
		&i.ToID,
		&i.Message,
		&i.Title,
		&i.Visualized,
		&i.Payload,
		&i.Type,
		&i.CreatedAt,
	)
	return i, err
}

const listNotifications = `-- name: ListNotifications :many
SELECT notification_id, "from", "to", from_id, to_id, message, title, visualized, payload, type, created_at FROM notification
ORDER BY created_at
`

func (q *Queries) ListNotifications(ctx context.Context) ([]Notification, error) {
	rows, err := q.db.Query(ctx, listNotifications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Notification
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.NotificationID,
			&i.From,
			&i.To,
			&i.FromID,
			&i.ToID,
			&i.Message,
			&i.Title,
			&i.Visualized,
			&i.Payload,
			&i.Type,
			&i.CreatedAt,
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

const listNotitifacionByType = `-- name: ListNotitifacionByType :many
SELECT notification_id, "from", "to", from_id, to_id, message, title, visualized, payload, type, created_at FROM notification
ORDER BY type
`

func (q *Queries) ListNotitifacionByType(ctx context.Context) ([]Notification, error) {
	rows, err := q.db.Query(ctx, listNotitifacionByType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Notification
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.NotificationID,
			&i.From,
			&i.To,
			&i.FromID,
			&i.ToID,
			&i.Message,
			&i.Title,
			&i.Visualized,
			&i.Payload,
			&i.Type,
			&i.CreatedAt,
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

const updateNotificationVisualized = `-- name: UpdateNotificationVisualized :exec
UPDATE notification
SET visualized = $2
WHERE notification_id = $1
`

type UpdateNotificationVisualizedParams struct {
	NotificationID int64       `json:"notification_id"`
	Visualized     pgtype.Bool `json:"visualized"`
}

func (q *Queries) UpdateNotificationVisualized(ctx context.Context, arg UpdateNotificationVisualizedParams) error {
	_, err := q.db.Exec(ctx, updateNotificationVisualized, arg.NotificationID, arg.Visualized)
	return err
}