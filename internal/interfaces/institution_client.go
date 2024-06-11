package interfaces

import (
	dto_client "optitech/internal/dto/client"
	dto "optitech/internal/dto/institution_client"
	models "optitech/internal/sqlc"
)

type IInstitutionClientService interface {
	List(InstitutionID int32) (*[]dto_client.GetClientRes, error)
	Create(req *[]models.CreateInstitutionClientParams) error
	Exists(req *models.ExistsInstitutionClientParams) bool
	DeleteById(req *dto.GetInstitutionClientReq) (bool, error)
	DeleteByInstitution(InstitutionID int32) (bool, error)
}
type IInstitutionClientRepository interface {
	ListInstitutionClient(InstitutionID int32) (*[]dto_client.GetClientRes, error)
	CreateInstitutionClient(arg *[]models.CreateInstitutionClientParams) error
	ExistsInstitutionClient(arg *models.ExistsInstitutionClientParams) bool
	DeleteInstitutionClientById(arg *models.DeleteinstInstitutionClientByClientAndInstitutionParams) error
	DeleteInstitutionClientByInstitution(arg *models.DeleteInstitutionClientByInstitutionParams) error
}
