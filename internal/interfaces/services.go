package interfaces

import (
	dto "optitech/internal/dto/services"
	models "optitech/internal/sqlc"
)

type IService interface {
	Get(req dto.GetServiceReq) (*dto.GetServiceRes, error)
	Create(req *dto.CreateServiceReq) (*dto.CreateServiceRes, error)
	Update(req *dto.UpdateServiceReq) (bool, error)
	List() (*[]dto.GetServiceRes, error)
	Delete(req dto.GetServiceReq) (bool, error)
}
type IServiceRepository interface {
	GetService(institutionID int32) (*dto.GetServiceRes, error)
	CreateService(arg *models.CreateServiceParams) (*dto.CreateServiceRes, error)
	UpdateService(arg *models.UpdateServiceParams) error
	ListServices() (*[]dto.GetServiceRes, error)
	DeleteService(arg *models.DeleteServiceParams) error
}
