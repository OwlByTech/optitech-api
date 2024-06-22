package repository

import (
	"context"
	dto "optitech/internal/dto/permission"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryPermission struct {
	permissionRepository *sq.Queries
}

func NewRepositoryPermission(q *sq.Queries) interfaces.IPermissionRepository {
	return &repositoryPermission{
		permissionRepository: q,
	}
}

func (r *repositoryPermission) CreatePermission(arg *sq.CreatePermissionParams) (*dto.CreatePermissionRes, error) {
	ctx := context.Background()
	res, err := r.permissionRepository.CreatePermission(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreatePermissionRes{
		Id:          res.PermissionID,
		Name:        res.Name,
		Code:        res.Code,
		Description: res.Description,
	}, err
}

func (r *repositoryPermission) DeletePermission(arg *sq.DeletePermissionByIdParams) error {
	ctx := context.Background()
	return r.permissionRepository.DeletePermissionById(ctx, *arg)
}

func (r *repositoryPermission) GetPermission(id int64) (*dto.GetPermissionRes, error) {
	ctx := context.Background()
	res, err := r.permissionRepository.GetPermission(ctx, id)

	if err != nil {
		return nil, err
	}

	return &dto.GetPermissionRes{
		Id:          res.PermissionID,
		Name:        res.Name,
		Code:        res.Code,
		Description: res.Description,
	}, nil
}

func (r *repositoryPermission) ListPermissions() (*[]dto.GetPermissionRes, error) {
	ctx := context.Background()
	repoRes, err := r.permissionRepository.ListPermissions(ctx)

	if err != nil {
		return nil, err
	}

	permissions := make([]dto.GetPermissionRes, len(repoRes))
	for i, inst := range repoRes {
		permissions[i] = dto.GetPermissionRes{
			Id:          inst.PermissionID,
			Name:        inst.Name,
			Code:        inst.Code,
			Description: inst.Description,
		}
	}

	return &permissions, nil
}

func (r *repositoryPermission) UpdatePermission(arg *sq.UpdatePermissionByIdParams) error {
	ctx := context.Background()
	return r.permissionRepository.UpdatePermissionById(ctx, *arg)
}
