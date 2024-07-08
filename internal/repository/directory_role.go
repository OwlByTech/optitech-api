package repository

import (
	"context"
	dto "optitech/internal/dto/directory_role"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
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

func (r *repositoryDirectoryRole) GetDirectoryRole(userID int64) (*dto.GetDirectoryRoleRes, error) {
	ctx := context.Background()
	repoRes, err := r.directoryRoleRepository.GetDirectoryRole(ctx, pgtype.Int4{Int32: int32(userID), Valid: true})

	if err != nil {
		return nil, err
	}

	return &dto.GetDirectoryRoleRes{
		DirectoryId: int64(repoRes.DirectoryID.Int32),
		UserId:      int64(repoRes.UserID.Int32),
		Status:      string(repoRes.Status),
	}, nil
}

func (r *repositoryDirectoryRole) ListDirectoryRole() (*[]dto.GetDirectoryRoleRes, error) {
	ctx := context.Background()
	repoRes, err := r.directoryRoleRepository.ListDirectoryRoles(ctx)

	if err != nil {
		return nil, err
	}

	directorys := make([]dto.GetDirectoryRoleRes, len(repoRes))
	for i, inst := range repoRes {
		directorys[i] = dto.GetDirectoryRoleRes{
			DirectoryId: int64(inst.DirectoryID.Int32),
			UserId:      int64(inst.UserID.Int32),
			Status:      string(inst.Status),
		}
	}

	return &directorys, nil
}
