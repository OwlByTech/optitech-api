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
		ServiceID:   repoRes.ServiceID,
		ServiceName: repoRes.ServiceName,
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
			ServiceID:   inst.ServiceID,
			ServiceName: inst.ServiceName,
		}
	}

	return &services, nil
}

func (r *repositoryService) CreateService(arg *sq.CreateServiceParams) (*dto.CreateServiceRes, error) {
	ctx := context.Background()
	res, err := r.serviceRepository.CreateService(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateServiceRes{
		ServiceName: res.ServiceName,
	}, nil
}

func (r *repositoryService) UpdateService(arg *sq.UpdateServiceParams) error {
	ctx := context.Background()
	return r.serviceRepository.UpdateService(ctx, *arg)

}

func (r *repositoryService) DeleteService(arg *sq.DeleteServiceParams) error {
	ctx := context.Background()
	return r.serviceRepository.DeleteService(ctx, *arg)

}
