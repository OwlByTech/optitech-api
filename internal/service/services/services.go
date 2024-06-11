package service

import (
	"github.com/jackc/pgx/v5/pgtype"
	dto "optitech/internal/dto/services"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"
)

type service_services struct {
	servicesRepository interfaces.IServiceRepository
}

func NewServiceServices(r interfaces.IServiceRepository) interfaces.IService {
	return &service_services{
		servicesRepository: r,
	}
}
func (s *service_services) Get(req dto.GetServiceReq) (*dto.GetServiceRes, error) {
	repoRes, err := s.servicesRepository.GetService(req.ServiceID)

	if err != nil {
		return nil, err
	}

	return repoRes, err
}

func (s *service_services) List() (*[]dto.GetServiceRes, error) {
	repoRes, err := s.servicesRepository.ListServices()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}

func (s *service_services) Create(req *dto.CreateServiceReq) (*dto.CreateServiceRes, error) {
	repoReq := &sq.CreateServiceParams{
		ServiceName: req.ServiceName,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := s.servicesRepository.CreateService(repoReq)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *service_services) Update(req *dto.UpdateServiceReq) (bool, error) {
	repoReq := &sq.UpdateServiceParams{
		ServiceID:   req.ServiceID,
		ServiceName: req.ServiceName,
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	err := s.servicesRepository.UpdateService(repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service_services) Delete(req dto.GetServiceReq) (bool, error) {
	repoReq := &sq.DeleteServiceParams{
		ServiceID: req.ServiceID,
		DeletedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	err := s.servicesRepository.DeleteService(repoReq)

	if err != nil {
		return false, err
	}

	return true, err
}
