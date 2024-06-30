package dto

import (
	"mime/multipart"
	dto_client "optitech/internal/dto/client"
	dto "optitech/internal/dto/services"
)

type CreateInstitutionReq struct {
	InstitutionName string `json:"name" validate:"required"`
	Description     string `json:"description" validate:"required"`
	LogoFile        *multipart.FileHeader
	AsesorID        int32   `json:"asesorId"`
	Services        []int32 `json:"services" validate:"required"`
	Clients         []int32 `json:"clients" `
}

type CreateInstitutionRes struct {
	InstitutionId   int32                     `json:"id"`
	InstitutionName string                    `json:"name"`
	Description     string                    `json:"description"`
	Services        []dto.GetServiceRes       `json:"services"`
	Clients         []dto_client.GetClientRes `json:"clients"`
}
