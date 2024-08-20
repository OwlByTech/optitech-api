package dto

import (
	dto_client "optitech/internal/dto/client"
	dto_services "optitech/internal/dto/services"
)

type GetInstitutionReq struct {
	Id int32 `json:"id" validate:"required"`
}

type GetInstitutionRes struct {
	Id              int32                        `json:"id"`
	InstitutionName string                       `json:"institutionName"`
	Logo            string                       `json:"logo"`
	Description     string                       `json:"description"`
	Services        []dto_services.GetServiceRes `json:"services"`
	AsesorId        int32                        `json:"asesorId"`
	Clients         []dto_client.GetClientRes    `json:"clients"`
}
