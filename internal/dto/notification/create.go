package dto

import "encoding/json"

type FromNotification string

const (
	FromNotificationInstitution FromNotification = "institution"
	FromNotificationAsesor      FromNotification = "asesor"
	FromNotificationSuperUser   FromNotification = "super_user"
)

type ToNotification string

const (
	ToNotificationInstitution ToNotification = "institution"
	ToNotificationAsesor      ToNotification = "asesor"
	ToNotificationSuperUser   ToNotification = "super_user"
	ToNotificationAll         ToNotification = "add"
)

type TypeNotification string

const (
	TypeNotificationInformation TypeNotification = "information"
	TypeNotificationCorrection  TypeNotification = "correction"
	TypeNotificationError       TypeNotification = "error"
	TypeNotificationAproved     TypeNotification = "aproved"
)

type CreateNorificationReq struct {
	From    FromNotification `json:"from" validate:"required"`
	To      ToNotification   `json:"to" validate:"required"`
	FromID  int32            `json:"fromId" validate:"required"`
	ToID    int32            `json:"toId" validate:"required"`
	Message string           `json:"message" validate:"required"`
	Title   string           `json:"title" validate:"required"`
	Payload json.RawMessage  `json:"payload"`
	Type    TypeNotification `json:"type" validate:"required"`
}

type CreateNotificationRes struct {
	From    FromNotification `json:"from"`
	To      ToNotification   `json:"to"`
	FromID  int32            `json:"fromId"`
	ToID    int32            `json:"toId"`
	Message string           `json:"message"`
	Title   string           `json:"title"`
	Payload json.RawMessage  `json:"payload"`
	Type    TypeNotification `json:"type"`
}
