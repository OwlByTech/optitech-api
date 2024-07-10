package interfaces

import (
	dto "optitech/internal/dto/directory_institution"
	i "optitech/internal/dto/institution"
	models "optitech/internal/sqlc"
)

type IDirectoryInstitutionService interface {
	ListInstitutionByDirectoryId(directoryId int32) (*[]i.GetInstitutionReq, error)
	ListDirectoryInstitution() (*[]dto.GetDirectoryInstitutionRes, error)
}

type IDirectoryInstitutionRepository interface {
	ListInstitutionByDirectoryId(directoryId int32) (*[]models.GetDirectoryInstitutionByDirectoryIdRow, error)
	ListDirectoryInstitution() (*[]dto.GetDirectoryInstitutionRes, error)
}

type IDirectoryInstitutionHandler interface{}
