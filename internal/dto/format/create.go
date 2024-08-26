package dto

import "mime/multipart"

type CreateFormatReq struct {
	UpdateFormatID int32  `json:"updateFormatId"`
	AsesorId       int32  `json:"asesorId"`
	ServiceID      int32  `json:"serviceId"`
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description" validate:"required"`
	Extension      string `json:"extension" validate:"required"`
	Version        string `json:"version" validate:"required"`
	DirectoryId    int64  `json:"directoryId" validate:"required"`
	FormatFile     *multipart.FileHeader
}

type CreateFormatRes struct {
	Id          int32  `json:"id"`
	AsesorId    int32  `json:"asesorId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Extension   string `json:"extension"`
	Version     string `json:"version"`
}
