package dto

import (
	"mime/multipart"
)

type UpdateInstitutionReq struct {
	InstitutionID   int32
	InstitutionName string `json:"institutionName" `
	Description     string `json:"description"`
	LogoFile        *multipart.FileHeader
	Services        []int32 `json:"services" `
}
