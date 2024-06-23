package repository

import (
	"context"
	dto "optitech/internal/dto/roles"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryRole struct {
	RoleRepository *sq.Queries
}

func NewRepositoryRole(q *sq.Queries) interfaces.IRoleRepository {
	return &repositoryRole{
		RoleRepository: q,
	}
}

func (r *repositoryRole) CreateRole(arg *sq.CreateRoleParams) (*dto.CreateRoleRes, error) {
	ctx := context.Background()
	res, err := r.RoleRepository.CreateRole(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateRoleRes{
		Id:          res.RoleID,
		RoleName:    res.RoleName,
		Description: res.Description,
	}, err
}

func (r *repositoryRole) DeleteRole(arg *sq.DeleteRoleByIdParams) error {
	ctx := context.Background()
	return r.RoleRepository.DeleteRoleById(ctx, *arg)
}

func (r *repositoryRole) GetRole(id int64) (*dto.GetRoleRes, error) {
	ctx := context.Background()
	res, err := r.RoleRepository.GetRole(ctx, id)

	if err != nil {
		return nil, err
	}

	return &dto.GetRoleRes{
		Id:          res.RoleID,
		RoleName:    res.RoleName,
		Description: res.Description,
	}, nil
}

func (r *repositoryRole) ListRoles() (*[]dto.GetRoleRes, error) {
	ctx := context.Background()
	repoRes, err := r.RoleRepository.ListRoles(ctx)

	if err != nil {
		return nil, err
	}

	Roles := make([]dto.GetRoleRes, len(repoRes))
	for i, inst := range repoRes {
		Roles[i] = dto.GetRoleRes{
			Id:          inst.RoleID,
			RoleName:    inst.RoleName,
			Description: inst.Description,
		}
	}

	return &Roles, nil
}

func (r *repositoryRole) UpdateRole(arg *sq.UpdateRoleByIdParams) error {
	ctx := context.Background()
	return r.RoleRepository.UpdateRoleById(ctx, *arg)
}
