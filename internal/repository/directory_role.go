package repository

import (
	"context"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryDirectoryRole struct {
	directoryRoleRepository *sq.Queries
}

func NewRepositoryDIrectoryRole(q *sq.Queries) interfaces.IDirectoryRoleRepository {
	return &repositoryDirectoryRole{
		directoryRoleRepository: q,
	}
}

func (r *repositoryDirectoryRole) CreateDirectoryRole(arg *[]sq.CreateDirectoryRoleParams) error {
	ctx := context.Background()
	_, err := r.directoryRoleRepository.CreateDirectoryRole(ctx, *arg)
	return err
}
