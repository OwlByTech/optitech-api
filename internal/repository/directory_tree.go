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

func NewRepositoryDirectoryTree(q *sq.Queries) interfaces.IDirectoryRepository {
	return &repositoryDirectoryTree{
		directoryRepository: q,
	}
}

func (r *repositoryDirectoryTree) GetDirectory(directoryID int64) (*dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()

	repoRes, err := r.directoryRepository.GetDirectoryTree(ctx, directoryID)

	if err != nil {
		return nil, err
	}

	return &dto.GetDirectoryTreeRes{
		Id:       repoRes.DirectoryID,
		ParentID: int64(repoRes.ParentID.Int32),
		Name:     repoRes.Name.String,
	}, nil
}

func (r *repositoryDirectoryTree) CreateDirectory(arg *sq.CreateDirectoryTreeParams) (*dto.CreateDirectoryTreeRes, error) {
	ctx := context.Background()

	res, err := r.directoryRepository.CreateDirectoryTree(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateDirectoryTreeRes{
		Id:          res.DirectoryID,
		DirectoryId: res.DirectoryID,
		ParentID:    int64(res.ParentID.Int32),
		Name:        res.Name.String,
	}, nil
}
