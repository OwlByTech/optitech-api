package dto

type UpdateNotificationVisualizedReq struct {
	NotificationID int64 `json:"id" validate:"required"`
	Visualized     bool  `json:"visualized" validate:"required"`
}
