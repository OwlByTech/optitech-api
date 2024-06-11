package dto

import (
	"mime/multipart"
)

type UpdateInstitutionReq struct {
	InstitutionID   int32
	InstitutionName string `json:"institutionName" `
	Description     string `json:"description"`
	LogoFile        *multipart.FileHeader
	AsesorID        int32    `json:"asesor_id"`
	Services        []string `json:"services" `
}
