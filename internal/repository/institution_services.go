package repository

import (
	"context"
	dtoService "optitech/internal/dto/services"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryInstitutionService struct {
	institutionRepository *sq.Queries
}

func NewRepositoryInstitutionServices(q *sq.Queries) interfaces.IInstitutionServiceRepository {
	return &repositoryInstitutionService{
		institutionRepository: q,
	}
}

func (r *repositoryInstitutionService) ListInstitutionServices(InstitutionID int32) (*[]dtoService.GetServiceRes, error) {
	ctx := context.Background()
	repoRes, err := r.institutionRepository.ListInstitutionServices(ctx, InstitutionID)
	if err != nil {
		return nil, err
	}

	institution_services := make([]dtoService.GetServiceRes, len(repoRes))
	for i, inst := range repoRes {
		institution_services[i] = dtoService.GetServiceRes{
			ServiceID:   inst.ServiceID,
			ServiceName: inst.ServiceName,
		}
	}

	return &institution_services, nil
}
func (r *repositoryInstitutionService) ExistsInstitutionService(arg *sq.ExistsInstitutionServiceParams) bool {
	ctx := context.Background()
	_, err := r.institutionRepository.ExistsInstitutionService(ctx, *arg)
	if err != nil {
		return false
	}
	return true

}

func (r *repositoryInstitutionService) CreateInstitutionService(arg *[]sq.CreateInstitutionServicesParams) error {
	ctx := context.Background()
	_, err := r.institutionRepository.CreateInstitutionServices(ctx, *arg)
	return err
}

func (r *repositoryInstitutionService) DeleteInstitutionServiceById(arg *sq.DeleteInstitutionServiceByIdParams) error {
	ctx := context.Background()
	return r.institutionRepository.DeleteInstitutionServiceById(ctx, *arg)

}
func (r *repositoryInstitutionService) DeleteInstitutionServiceByInstitution(arg *sq.DeleteInstitutionServicesByInstitutionParams) error {
	ctx := context.Background()
	return r.institutionRepository.DeleteInstitutionServicesByInstitution(ctx, *arg)

}
