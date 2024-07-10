package repository

import (
	"context"
	dto "optitech/internal/dto/directory_institution"
	sq "optitech/internal/sqlc"
)

type repositoryDirectoryInstitution struct {
	directoryInstitutionRepository *sq.Queries
}

func (r *repositoryDirectoryInstitution) ListDirectoryInstitution() (*[]dto.GetDirectoryInstitutionRes, error) {
	ctx := context.Background()
	repoRes, err := r.directoryInstitutionRepository.ListDirectoryInstitutions(ctx)
	if err != nil {
		return nil, err
	}

	institutions := make([]dto.GetDirectoryInstitutionRes, len(repoRes))
	for i, inst := range repoRes {
		institutions[i] = dto.GetDirectoryInstitutionRes{
			InstitutionId: inst.InstitutionID,
			DirectoryId:   inst.DirectoryID,
		}
	}

	return &institutions, nil
}

func (r *repositoryDirectoryInstitution) ListByDirectoryId(directoryInstitutionId int32) (*[]sq.GetDirectoryInstitutionByDirectoryIdRow, error) {
	ctx := context.Background()
	res, err := r.directoryInstitutionRepository.GetDirectoryInstitutionByDirectoryId(ctx, directoryInstitutionId)

	if err != nil {
		return nil, err
	}
	return &res, nil
}
