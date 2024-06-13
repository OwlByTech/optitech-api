package interfaces

import (
	dto "optitech/internal/dto/institution_services"
	dto_service "optitech/internal/dto/services"
	models "optitech/internal/sqlc"
)

type IServiceInstitutionService interface {
	List(InstitutionID int32) (*[]dto_service.GetServiceRes, error)
	Create(req *[]models.CreateInstitutionServicesParams) error
	Exists(req *models.ExistsInstitutionServiceParams) bool
	Recover(arg *models.RecoverInstitutionServiceParams) error
	Update(req *dto.UpdateInstitutionServicesReq) bool
	DeleteById(req *dto.GetInstitutionServicesReq) error
	DeleteByInstitution(InstitutionID int32) (bool, error)
}
type IInstitutionServiceRepository interface {
	ListInstitutionServices(institutionID int32) (*[]dto_service.GetServiceRes, error)
	CreateInstitutionService(arg *[]models.CreateInstitutionServicesParams) error
	ExistsInstitutionService(arg *models.ExistsInstitutionServiceParams) bool
	RecoverInstitutionService(arg *models.RecoverInstitutionServiceParams) error
	DeleteInstitutionServiceById(arg *models.DeleteInstitutionServiceByIdParams) error
	DeleteInstitutionServiceByInstitution(arg *models.DeleteInstitutionServicesByInstitutionParams) error
}
