package dto

import (
	"mime/multipart"
)

type UpdateInstitutionReq struct {
	InstitutionName string `json:"institutionName" `
	Description     string `json:"description"`
	LogoFile        *multipart.FileHeader
	AsesorID        int32    `json:"asesor_id"`
	Services        []string `json:"services" `
}

type UpdateInstitutionRes struct {
	Id              int64
	InstitutionName string
	Description     string
	Services        []string
}
