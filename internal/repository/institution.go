package repository

import (
	"context"
	dto "optitech/internal/dto/institution"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
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
		AsesorId:        repoRes.AsesorID.Int32,
	}, nil
}

func (r *repositoryInstitution) GetInstitutionLogo(InstitutionID int32) (*dto.GetInstitutionRes, error) {

	ctx := context.Background()
	repoRes, err := r.institutionRepository.GetInstitutionLogo(ctx, InstitutionID)

	if err != nil {
		return nil, err
	}
	return &dto.GetInstitutionRes{
		InstitutionName: repoRes.InstitutionName,
		Logo:            repoRes.Logo.String,
	}, nil

}

func (r *repositoryInstitution) GetInstitutionByClient(ClientID int32) (int32, error) {
	ctx := context.Background()
	return r.institutionRepository.GetInstitutionByClient(ctx, ClientID)
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

func (r *repositoryInstitution) CreateInstitution(arg *sq.CreateInstitutionParams) (int32, error) {
	ctx := context.Background()
	return r.institutionRepository.CreateInstitution(ctx, *arg)

}

func (r *repositoryInstitution) UpdateAsesorInstitution(arg *sq.UpdateAsesorInstitutionParams) error {
	ctx := context.Background()
	return r.institutionRepository.UpdateAsesorInstitution(ctx, *arg)
}

func (r *repositoryInstitution) UpdateInstitution(arg *sq.UpdateInstitutionParams) error {
	ctx := context.Background()
	return r.institutionRepository.UpdateInstitution(ctx, *arg)

}
func (r *repositoryInstitution) UpdateLogoInstitution(arg *sq.UpdateLogoInstitutionParams) error {
	ctx := context.Background()
	return r.institutionRepository.UpdateLogoInstitution(ctx, *arg)

}

func (r *repositoryInstitution) DeleteInstitution(arg *sq.DeleteInstitutionParams) error {
	ctx := context.Background()
	return r.institutionRepository.DeleteInstitution(ctx, *arg)

}

func (r *repositoryInstitution) GetInstitutionByAsesor(ClientID int32) (*[]dto.GetInstitutionRes, error) {
	ctx := context.Background()
	pgClientID := pgtype.Int4{
		Int32: ClientID,
		Valid: true,
	}

	institutionIDs, err := r.institutionRepository.GetInstitutionByAsesor(ctx, pgClientID)
	if err != nil {
		return nil, err
	}

	var institutions []dto.GetInstitutionRes
	for _, id := range institutionIDs {
		institution, err := r.institutionRepository.GetInstitution(ctx, id)
		if err != nil {
			return nil, err
		}

		institutions = append(institutions, dto.GetInstitutionRes{
			Id:              institution.InstitutionID,
			InstitutionName: institution.InstitutionName,
			Logo:            institution.Logo.String,
			Description:     institution.Description,
			AsesorId:        institution.AsesorID.Int32,
		})
	}

	return &institutions, nil
}
