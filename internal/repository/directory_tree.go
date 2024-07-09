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
		Id:       int32(repoRes.DirectoryID),
		ParentID: int32(repoRes.ParentID.Int32),
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
		Id:          int32(res.DirectoryID),
		DirectoryId: int32(res.DirectoryID),
		ParentID:    int32(res.ParentID.Int32),
		Name:        res.Name.String,
	}, nil
}

func (r *repositoryDirectoryTree) ListDirectory() (*[]dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()
	repoRes, err := r.directoryRepository.ListDirectoryTrees(ctx)

	if err != nil {
		return nil, err
	}

	directorys := make([]dto.GetDirectoryTreeRes, len(repoRes))
	for i, inst := range repoRes {
		directorys[i] = dto.GetDirectoryTreeRes{
			Id:       int32(inst.DirectoryID),
			ParentID: int32(inst.ParentID.Int32),
			Name:     inst.Name.String,
		}
	}
	return &directorys, nil
}

func (r *repositoryDirectoryTree) DeleteDirectory(arg *sq.DeleteDirectoryTreeByIdParams) error {
	ctx := context.Background()
	return r.directoryRepository.DeleteDirectoryTreeById(ctx, *arg)
}
