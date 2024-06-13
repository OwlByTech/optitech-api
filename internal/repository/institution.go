package repository

import (
	"context"
	dto "optitech/internal/dto/institution"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryInstitution struct {
	institutionRepository *sq.Queries
}

func NewRepositoryInstitution(q *sq.Queries) interfaces.IInstitutionRepository {
	return &repositoryInstitution{
		institutionRepository: q,
	}
}

func (r *repositoryInstitution) GetInstitution(InstitutionID int32) (*dto.GetInstitutionRes, error) {
	ctx := context.Background()
	repoRes, err := r.institutionRepository.GetInstitution(ctx, InstitutionID)

	if err != nil {
		return nil, err
	}
	return &dto.GetInstitutionRes{
		Id:              repoRes.InstitutionID,
		InstitutionName: repoRes.InstitutionName,
		Description:     repoRes.Description,
		Logo:            repoRes.Logo.String,
	}, nil
}

func (r *repositoryInstitution) ListInstitutions() (*[]dto.GetInstitutionRes, error) {
	ctx := context.Background()
	repoRes, err := r.institutionRepository.ListInstitutions(ctx)
	if err != nil {
		return nil, err
	}

	institutions := make([]dto.GetInstitutionRes, len(repoRes))
	for i, inst := range repoRes {
		institutions[i] = dto.GetInstitutionRes{
			Id:              inst.InstitutionID,
			Description:     inst.Description,
			InstitutionName: inst.InstitutionName,
			Logo:            inst.Logo.String,
		}
	}

	return &institutions, nil
}

func (r *repositoryInstitution) CreateInstitution(arg *sq.CreateInstitutionParams) (*dto.CreateInstitutionRes, error) {
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

func (r *repositoryInstitution) UpdateAsesorInstitution(arg *sq.UpdateAsesorInstitutionParams) error {
	ctx := context.Background()
	return r.institutionRepository.UpdateAsesorInstitution(ctx, *arg)
}

func (r *repositoryInstitution) UpdateInstitution(arg *sq.UpdateInstitutionParams) error {
	ctx := context.Background()
	return r.institutionRepository.UpdateInstitution(ctx, *arg)

}

func (r *repositoryInstitution) DeleteInstitution(arg *sq.DeleteInstitutionParams) error {
	ctx := context.Background()
	return r.institutionRepository.DeleteInstitution(ctx, *arg)

}
