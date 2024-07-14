package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
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
		Id:            repoRes.DirectoryID,
		ParentID:      repoRes.ParentID.Int64,
		Name:          repoRes.Name.String,
		InstitutionID: repoRes.InstitutionID.Int32,
	}, nil
}

func (r *repositoryDirectoryTree) CreateDirectory(arg *sq.CreateDirectoryTreeParams) (*dto.CreateDirectoryTreeRes, error) {
	ctx := context.Background()

	res, err := r.directoryRepository.CreateDirectoryTree(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateDirectoryTreeRes{
		DirectoryId:   res.DirectoryID,
		ParentID:      res.ParentID.Int64,
		Name:          res.Name.String,
		InstitutionID: res.InstitutionID.Int32,
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
			Id:       inst.DirectoryID,
			ParentID: inst.ParentID.Int64,
			Name:     inst.Name.String,
			InstitutionID: inst.InstitutionID.Int32,
		}
	}
	return &directorys, nil
}
func (r *repositoryDirectoryTree) ListDirectoryByParent(parentId int64) ([]*dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()
	repoRes, err := r.directoryRepository.ListDirectoryChildByParent(ctx, pgtype.Int8{Int64: parentId, Valid: true})

	if err != nil {
		return nil, err
	}
	directorys := make([]*dto.GetDirectoryTreeRes, len(repoRes))
	for i, inst := range repoRes {
		directorys[i] = &dto.GetDirectoryTreeRes{
			Id:       inst.DirectoryID,
			ParentID: inst.ParentID.Int64,
			Name:     inst.Name.String,
			InstitutionID: inst.InstitutionID.Int32,
		}
	}
	return directorys, nil
}
func (r *repositoryDirectoryTree) ListDirectoryHierarchy(childId int64) (*[]dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()
	repoRes, err := r.directoryRepository.ListDirectoryHierarchyById(ctx, childId)

	if err != nil {
		return nil, err
	}
	directorys := make([]dto.GetDirectoryTreeRes, len(repoRes))
	for i, inst := range repoRes {
		directorys[i] = dto.GetDirectoryTreeRes{
			Id:       inst.DirectoryID,
			ParentID: inst.ParentID.Int64,
			Name:     inst.Name.String,
		}
	}
	return &directorys, nil
}

func (r *repositoryDirectoryTree) DeleteDirectory(arg *sq.DeleteDirectoryTreeByIdParams) error {
	ctx := context.Background()
	return r.directoryRepository.DeleteDirectoryTreeById(ctx, *arg)
}
