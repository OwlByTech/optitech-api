package interfaces

import (
	i "optitech/internal/dto/institution"
	models "optitech/internal/sqlc"
)

type IDirectoryInstitutionService interface {
	ListInstitutionByDirectoryId(directoryId int32) (*[]i.GetInstitutionReq, error)
}

type IDirectoryInstitutionRepository interface {
	ListInstitutionByDirectoryId(directoryId int32) (*[]models.GetDirectoryInstitutionByDirectoryIdRow, error)
}

type IDirectoryInstitutionHandler interface{}
