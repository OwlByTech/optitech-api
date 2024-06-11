package repository

import (
	"context"
	dto_service "optitech/internal/dto/services"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repository_institution_services struct {
	institutionRepository *sq.Queries
}

func NewRepositoryInstitutionServices(q *sq.Queries) interfaces.IInstitutionServiceRepository {
	return &repository_institution_services{
		institutionRepository: q,
	}
}

func (r *repository_institution_services) ListInstitutionServices(InstitutionID int32) (*[]dto_service.GetServiceRes, error) {
	ctx := context.Background()
	repoRes, err := r.institutionRepository.ListInstitutionServices(ctx, InstitutionID)
	if err != nil {
		return nil, err
	}

	institution_services := make([]dto_service.GetServiceRes, len(repoRes))
	for i, inst := range repoRes {
		institution_services[i] = dto_service.GetServiceRes{
			ServiceID:   inst.ServiceID,
			ServiceName: inst.ServiceName,
		}
	}

	return &institution_services, nil
}
func (r *repository_institution_services) ExistsInstitutionService(arg *sq.ExistsInstitutionServiceParams) bool {
	ctx := context.Background()
	_, err := r.institutionRepository.ExistsInstitutionService(ctx, *arg)
	if err != nil {
		return false
	}
	return true

}

func (r *repository_institution_services) CreateInstitutionService(arg *[]sq.CreateInstitutionServicesParams) error {
	ctx := context.Background()
	_, err := r.institutionRepository.CreateInstitutionServices(ctx, *arg)
	return err
}

func (r *repository_institution_services) DeleteInstitutionServiceById(arg *sq.DeleteInstitutionServiceByIdParams) error {
	ctx := context.Background()
	return r.institutionRepository.DeleteInstitutionServiceById(ctx, *arg)

}
func (r *repository_institution_services) DeleteInstitutionServiceByInstitution(arg *sq.DeleteInstitutionServicesByInstitutionParams) error {
	ctx := context.Background()
	return r.institutionRepository.DeleteInstitutionServicesByInstitution(ctx, *arg)

}
