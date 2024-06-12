package service

import (
	dto "optitech/internal/dto/services"
	"optitech/internal/interfaces"
)

type serviceService struct {
	servicesRepository interfaces.IServiceRepository
}

func NewServiceServices(r interfaces.IServiceRepository) interfaces.IService {
	return &serviceService{
		servicesRepository: r,
	}
}
func (s *serviceService) Get(req dto.GetServiceReq) (*dto.GetServiceRes, error) {
	repoRes, err := s.servicesRepository.GetService(req.ServiceID)

	if err != nil {
		return nil, err
	}

	return repoRes, err
}

func (s *serviceService) List() (*[]dto.GetServiceRes, error) {
	repoRes, err := s.servicesRepository.ListServices()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}
