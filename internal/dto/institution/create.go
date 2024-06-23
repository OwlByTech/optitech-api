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
	InstitutionId   int32                     `json:"institutionId"`
	InstitutionName string                    `json:"institutionName"`
	Description     string                    `json:"description"`
	Services        []dto.GetServiceRes       `json:"services"`
	Clients         []dto_client.GetClientRes `json:"clients"`
}
