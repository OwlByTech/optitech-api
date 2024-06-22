package repository

import (
	"context"
	dto "optitech/internal/dto/role_permission"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryRolePermission struct {
	RolePermissionRepository *sq.Queries
}

func NewRepositoryRolePermission(q *sq.Queries) interfaces.IRolePermissionRepository {
	return &repositoryRolePermission{
		RolePermissionRepository: q,
	}
}

func (r *repositoryRolePermission) GetRolePermissionByRoleId(roleId int32) (*sq.GetRolePermissionByRoleIdRow, error) {
	ctx := context.Background()
	res, err := r.RolePermissionRepository.GetRolePermissionByRoleId(ctx, roleId)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *repositoryRolePermission) CreateRolePermission(arg *sq.CreateRolePermissionParams) (*dto.CreateRolePermissionRes, error) {
	ctx := context.Background()
	res, err := r.RolePermissionRepository.CreateRolePermission(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateRolePermissionRes{
		Id:           res.RolePermissionID,
		RoleId:       res.RoleID,
		PermissionId: res.PermissionID,
	}, err
}

func (r *repositoryRolePermission) DeleteRolePermission(arg *sq.DeleteRolePermissionByIdParams) error {
	ctx := context.Background()
	return r.RolePermissionRepository.DeleteRolePermissionById(ctx, *arg)
}

func (r *repositoryRolePermission) GetRolePermission(id int64) (*dto.GetRolePermissionRes, error) {
	ctx := context.Background()
	res, err := r.RolePermissionRepository.GetRolePermission(ctx, id)

	if err != nil {
		return nil, err
	}

	return &dto.GetRolePermissionRes{
		Id:           res.RoleID,
		RoleId:       res.RoleID,
		PermissionId: res.PermissionID,
	}, nil
}

func (r *repositoryRolePermission) ListRolePermissions() (*[]dto.GetRolePermissionRes, error) {
	ctx := context.Background()
	repoRes, err := r.RolePermissionRepository.ListRolePermissions(ctx)

	if err != nil {
		return nil, err
	}

	RolePermissions := make([]dto.GetRolePermissionRes, len(repoRes))
	for i, inst := range repoRes {
		RolePermissions[i] = dto.GetRolePermissionRes{
			Id:           inst.RoleID,
			RoleId:       inst.RoleID,
			PermissionId: inst.PermissionID,
		}
	}

	return &RolePermissions, nil
}

func (r *repositoryRolePermission) UpdateRolePermission(arg *sq.UpdateRolePermissionByIdParams) error {
	ctx := context.Background()
	return r.RolePermissionRepository.UpdateRolePermissionById(ctx, *arg)
}
