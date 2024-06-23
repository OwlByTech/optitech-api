package dto

import (
	dto_client "optitech/internal/dto/client"
	dto_services "optitech/internal/dto/services"
)

type GetInstitutionReq struct {
	Id int32 `validate:"required"`
}

type GetInstitutionRes struct {
	Id              int32
	InstitutionName string
	Logo            string
	Description     string
	AsesorId        int32
	Services        []dto_services.GetServiceRes
	Clients         []dto_client.GetClientRes
}
