package dto

import (
	"encoding/json"
)

type GetNotificationReq struct {
	ID int64 `json:"id" validate:"required"`
}

type GetNotificationRes struct {
	ID         int64            `json:"id"`
	From       FromNotification `json:"from"`
	To         ToNotification   `json:"to"`
	FromID     int32            `json:"fromId"`
	ToID       int32            `json:"toId"`
	Message    string           `json:"message"`
	Title      string           `json:"title"`
	Payload    json.RawMessage  `json:"payload"`
	Type       TypeNotification `json:"type"`
	Visualized bool             `json:"visualized"`
}
