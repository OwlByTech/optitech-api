package repository

import (
	"context"
	dto "optitech/internal/dto/directory_tree"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type repositoryDirectoryTree struct {
	directoryRepository *sq.Queries
}

func NewRepositoryDirectoryTree(q *sq.Queries) interfaces.IDirectoryRepository {
	return &repositoryDirectoryTree{
		directoryRepository: q,
	}
}

func (r *repositoryDirectoryTree) GetDirectory(req *sq.GetDirectoryTreeParams) (*dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()

	repoRes, err := r.directoryRepository.GetDirectoryTree(ctx, *req)

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

func (r *repositoryDirectoryTree) GetDirectoryParentInstitution(institutionId int32) (*dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()

	repoRes, err := r.directoryRepository.GetDirectoryTreeParent(ctx, pgtype.Int4{Int32: institutionId, Valid: true})

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
			Id:            inst.DirectoryID,
			ParentID:      inst.ParentID.Int64,
			Name:          inst.Name.String,
			InstitutionID: inst.InstitutionID.Int32,
		}
	}
	return &directorys, nil
}
func (r *repositoryDirectoryTree) ListDirectoryByParent(parentId int64, institutionId int32) ([]*dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()
	repoRes, err := r.directoryRepository.ListDirectoryChildByParent(ctx, sq.ListDirectoryChildByParentParams{ParentID: pgtype.Int8{Int64: parentId, Valid: true}, InstitutionID: pgtype.Int4{Int32: institutionId, Valid: true}})

	if err != nil {
		return nil, err
	}
	directorys := make([]*dto.GetDirectoryTreeRes, len(repoRes))
	for i, inst := range repoRes {
		directorys[i] = &dto.GetDirectoryTreeRes{
			Id:            inst.DirectoryID,
			ParentID:      inst.ParentID.Int64,
			Name:          inst.Name.String,
			InstitutionID: inst.InstitutionID.Int32,
		}
	}
	return directorys, nil
}
func (r *repositoryDirectoryTree) ListDirectoryHierarchy(childId int64, institutionId int32) (*[]dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()
	repoRes, err := r.directoryRepository.ListDirectoryHierarchyById(ctx, sq.ListDirectoryHierarchyByIdParams{InstitutionID: pgtype.Int4{Int32: institutionId, Valid: true}, DirectoryID: childId})

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

func (r *repositoryDirectoryTree) UpdateDirectoryTree(arg *sq.UpdateDirectoryTreeByIdParams) error {
	ctx := context.Background()
	return r.directoryRepository.UpdateDirectoryTreeById(ctx, *arg)
}
