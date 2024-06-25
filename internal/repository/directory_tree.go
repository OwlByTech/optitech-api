package repository

import (
	"context"
	dto "optitech/internal/dto/directory_tree"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryDirectoryTree struct {
	directoryRepository *sq.Queries
}

func NewRepositoryDirectoryTree(q *sq.Queries) interfaces.IDirectoryRepositoy {
	return &repositoryDirectoryTree{
		directoryRepository: q,
	}
}

func (r *repositoryDirectoryTree) GetDirectroy(directoryID int64) (*dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()

	repoRes, err := r.directoryRepository.GetDirectoryTree(ctx, (directoryID))

	if err != nil {
		return nil, err
	}

	return &dto.GetDirectoryTreeRes{
		Id:          repoRes.DirectoryID,
		DirectoryId: repoRes.DirectoryID,
	}, nil
}

func (r *repositoryDirectoryTree) CreateDirectoy(arg *sq.CreateDirectoryTreeParams) (*dto.CreateDirectoryTreeRes, error) {
	ctx := context.Background()

	res, err := r.directoryRepository.CreateDirectoryTree(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateDirectoryTreeRes{
		Id:          res.DirectoryID,
		DirectoryId: res.DirectoryID,
	}, nil
}
