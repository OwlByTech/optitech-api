package repository

import (
	"context"
	sq "optitech/internal/sqlc"
)

type repositoryDirectoryInstitution struct {
	DirectoryInstitutionRepository *sq.Queries
}

func (r *repositoryDirectoryInstitution) ListByDirectoryId(directoryInstitutionId int32) (*[]sq.GetDirectoryInstitutionByDirectoryIdRow, error) {
	ctx := context.Background()
	res, err := r.DirectoryInstitutionRepository.GetDirectoryInstitutionByDirectoryId(ctx, directoryInstitutionId)

	if err != nil {
		return nil, err
	}
	return &res, nil
}
