package service

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	dto "optitech/internal/dto/services"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
	"time"
)

func GetService(req dto.GetServiceReq) (*dto.GetServiceRes, error) {
	ctx := context.Background()

	repoRes, err := repository.Queries.GetServices(ctx, req.ServiceID)

	if err != nil {
		return nil, err
	}

	return &dto.GetServiceRes{
		ServiceID:   repoRes.ServiceID,
		ServiceName: repoRes.ServiceName,
	}, nil
}

func ListServices() ([]*dto.GetServiceRes, error) {
	ctx := context.Background()
	repoRes, err := repository.Queries.ListServicess(ctx)
	if err != nil {
		return nil, err
	}

	services := make([]*dto.GetServiceRes, len(repoRes))
	for i, inst := range repoRes {
		services[i] = &dto.GetServiceRes{
			ServiceID:   inst.ServiceID,
			ServiceName: inst.ServiceName,
		}
	}

	return services, nil
}

func CreateService(req dto.CreateServiceReq) (*sq.Service, error) {
	ctx := context.Background()
	repoReq := sq.CreateServicesParams{
		ServiceName: req.ServiceName,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := repository.Queries.CreateServices(ctx, repoReq)

	if err != nil {
		return nil, err
	}

	return &r, nil
}

func UpdateService(req dto.UpdateServiceReq) (bool, error) {
	ctx := context.Background()
	repoReq := sq.UpdateServicesByIdParams{
		ServiceName: req.ServiceName,
	}

	err := repository.Queries.UpdateServicesById(ctx, repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteService(req dto.GetServiceReq) (bool, error) {
	ctx := context.Background()
	repoReq := sq.DeleteServicesByIdParams{
		ServiceID: req.ServiceID,
		DeletedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	err := repository.Queries.DeleteServicesById(ctx, repoReq)

	if err != nil {
		return false, err
	}

	return true, err
}
