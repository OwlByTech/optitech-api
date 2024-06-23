package repository

import (
	"context"
	dto "optitech/internal/dto/services"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryService struct {
	serviceRepository *sq.Queries
}

func NewRepositoryService(q *sq.Queries) interfaces.IServiceRepository {
	return &repositoryService{
		serviceRepository: q,
	}
}

func (r *repositoryService) GetService(ServiceID int32) (*dto.GetServiceRes, error) {
	ctx := context.Background()
	repoRes, err := r.serviceRepository.GetService(ctx, ServiceID)

	if err != nil {
		return nil, err
	}
	return &dto.GetServiceRes{
		Id:   repoRes.ServiceID,
		Name: repoRes.Name,
	}, nil
}

func (r *repositoryService) ListServices() (*[]dto.GetServiceRes, error) {
	ctx := context.Background()
	repoRes, err := r.serviceRepository.ListServices(ctx)
	if err != nil {
		return nil, err
	}

	services := make([]dto.GetServiceRes, len(repoRes))
	for i, inst := range repoRes {
		services[i] = dto.GetServiceRes{
			Id:   inst.ServiceID,
			Name: inst.Name,
		}
	}

	return &services, nil
}
