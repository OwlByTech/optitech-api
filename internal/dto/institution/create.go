package dto

import "mime/multipart"

type CreateInstitutionReq struct {
	InstitutionName string `json:"institutionName" validate:"required"`
	Description     string `json:"description" validate:"required"`
	LogoFile        *multipart.FileHeader
	AsesorID        int32    `json:"asesor_id"`
	Services        []string `json:"services" validate:"required"`
}

type CreateInstitutionRes struct {
	Id              int64
	InstitutionName string
	Description     string
	Services        []string
}
