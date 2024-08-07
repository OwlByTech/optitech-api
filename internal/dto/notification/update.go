package dto

type UpdateNotificationVisualizedReq struct {
	NotificationID int64 `json:"notificationId" validate:"required"`
	Visualized     bool  `json:"visualized" validate:"required"`
}
