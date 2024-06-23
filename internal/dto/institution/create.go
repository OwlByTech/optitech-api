package dto

import (
	"mime/multipart"
	dto_client "optitech/internal/dto/client"
	dto "optitech/internal/dto/services"
)

type CreateInstitutionReq struct {
	InstitutionName string `json:"institutionName" validate:"required"`
	Description     string `json:"description" validate:"required"`
	LogoFile        *multipart.FileHeader
	AsesorID        int32   `json:"asesor_id"`
	Services        []int32 `json:"services" validate:"required"`
	Clients         []int32 `json:"clients" validate:"required"`
}

type CreateInstitutionRes struct {
	InstitutionID   int32
	InstitutionName string
	Description     string
	Services        []dto.GetServiceRes
	Clients         []dto_client.GetClientRes
}
