package repository

import (
	"context"
	dto "optitech/internal/dto/institution"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repository_institution struct {
	institutionRepository *sq.Queries
}

func NewRepositoryInstitution(q *sq.Queries) interfaces.IInstitutionRepository {
	return &repository_institution{
		institutionRepository: q,
	}
}

func (r *repository_institution) GetInstitution(InstitutionID int64) (*dto.GetInstitutionRes, error) {
	ctx := context.Background()
	repoRes, err := r.institutionRepository.GetInstitution(ctx, InstitutionID)

	if err != nil {
		return nil, err
	}
	return &dto.GetInstitutionRes{
		InstitutionID:   repoRes.InstitutionID,
		InstitutionName: repoRes.InstitutionName,
		Description:     repoRes.Description,
	}, nil
}

func (r *repository_institution) ListInstitutions() ([]*dto.GetInstitutionRes, error) {
	ctx := context.Background()
	repoRes, err := r.institutionRepository.ListInstitutions(ctx)
	if err != nil {
		return nil, err
	}

	institutions := make([]*dto.GetInstitutionRes, len(repoRes))
	for i, inst := range repoRes {
		services := []string{}
		institutions[i] = &dto.GetInstitutionRes{
			InstitutionID:   inst.InstitutionID,
			Description:     inst.Description,
			InstitutionName: inst.InstitutionName,
			Logo:            inst.Logo.String,
			Services:        services,
		}
	}

	return institutions, nil
}

func (r *repository_institution) CreateInstitution(arg *sq.CreateInstitutionParams) (*dto.CreateInstitutionRes, error) {
	ctx := context.Background()
	res, err := r.institutionRepository.CreateInstitution(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateInstitutionRes{
		InstitutionID:   res.InstitutionID,
		InstitutionName: res.InstitutionName,
		Description:     res.Description,
	}, nil
}

func (r *repository_institution) UpdateInstitution(arg *sq.UpdateInstitutionParams) error {
	ctx := context.Background()
	return r.institutionRepository.UpdateInstitution(ctx, *arg)

}

func (r *repository_institution) DeleteInstitution(arg *sq.DeleteInstitutionParams) error {
	ctx := context.Background()
	return r.institutionRepository.DeleteInstitution(ctx, *arg)

}
