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

func (r *repositoryDirectoryTree) GetDirectory(req *dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()

	if req.AsesorID > 0 {
		repoRes, err := r.directoryRepository.GetDirectoryTreeByAsesor(ctx, sq.GetDirectoryTreeByAsesorParams{
			DirectoryID: req.Id,
			AsesorID:    pgtype.Int4{Int32: req.AsesorID, Valid: true},
		})

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
	repoRes, err := r.directoryRepository.GetDirectoryTreeByInstitution(ctx, sq.GetDirectoryTreeByInstitutionParams{
		DirectoryID:   req.Id,
		InstitutionID: pgtype.Int4{Int32: req.InstitutionID, Valid: true},
	})

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

func (r *repositoryDirectoryTree) GetDirectoryParent(req *dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()
	if req.InstitutionID > 0 {
		repoRes, err := r.directoryRepository.GetDirectoryInstitutionTreeParent(ctx, pgtype.Int4{Int32: req.InstitutionID, Valid: true})

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
	repoRes, err := r.directoryRepository.GetDirectoryAsesorTreeParent(ctx, pgtype.Int4{Int32: req.AsesorID, Valid: true})

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
func (r *repositoryDirectoryTree) ListDirectoryByParent(req *dto.GetDirectoryTreeReq) ([]*dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()
	if req.InstitutionID > 0 {
		repoRes, err := r.directoryRepository.ListDirectoryInstitutionChildByParent(ctx, sq.ListDirectoryInstitutionChildByParentParams{ParentID: pgtype.Int8{Int64: req.Id, Valid: true}, InstitutionID: pgtype.Int4{Int32: req.InstitutionID, Valid: true}})

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
	repoRes, err := r.directoryRepository.ListDirectorAsesoryChildByParent(ctx, sq.ListDirectorAsesoryChildByParentParams{ParentID: pgtype.Int8{Int64: req.Id, Valid: true}, AsesorID: pgtype.Int4{Int32: req.AsesorID, Valid: true}})

	if err != nil {
		return nil, err
	}
	directorys := make([]*dto.GetDirectoryTreeRes, len(repoRes))
	for i, inst := range repoRes {
		directorys[i] = &dto.GetDirectoryTreeRes{
			Id:       inst.DirectoryID,
			ParentID: inst.ParentID.Int64,
			Name:     inst.Name.String,
			AsesorID: inst.AsesorID.Int32,
		}
	}
	return directorys, nil
}
func (r *repositoryDirectoryTree) ListDirectoryHierarchy(req *dto.GetDirectoryTreeReq) (*[]dto.GetDirectoryTreeRes, error) {
	ctx := context.Background()
	if req.InstitutionID > 0 {
		repoRes, err := r.directoryRepository.ListDirectoryHierarchyInstitutionById(ctx, sq.ListDirectoryHierarchyInstitutionByIdParams{InstitutionID: pgtype.Int4{Int32: req.InstitutionID, Valid: true}, DirectoryID: req.Id})

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
	repoRes, err := r.directoryRepository.ListDirectoryHierarchyAsesorById(ctx, sq.ListDirectoryHierarchyAsesorByIdParams{AsesorID: pgtype.Int4{Int32: req.AsesorID, Valid: true}, DirectoryID: req.Id})

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
