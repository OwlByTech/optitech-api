package repository

import (
	"context"
	dto "optitech/internal/dto/directory_role"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryDirectoryRole struct {
	directoryRoleRepository *sq.Queries
}

func NewRepositoryDirectoryRole(q *sq.Queries) interfaces.IDirectoryRoleRepository {
	return &repositoryDirectoryRole{
		directoryRoleRepository: q,
	}
}

func (r *repositoryDirectoryRole) CreateDirectoryRole(arg *sq.CreateDirectoryRoleParams) (*dto.CreateDirectoryRoleRes, error) {
	ctx := context.Background()
	res, err := r.directoryRoleRepository.CreateDirectoryRole(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateDirectoryRoleRes{
		DirectoryId: int64(res.DirectoryID.Int32),
		UserId:      int64(res.UserID.Int32),
		Status:      string(res.Status),
	}, nil
}
