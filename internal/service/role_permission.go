package service

import (
	p "optitech/internal/dto/permission"
	dto "optitech/internal/dto/role_permission"
	r "optitech/internal/dto/roles"
	"optitech/internal/interfaces"
)

type serviceRolePermission struct {
	RolePermissionRepository interfaces.IRolePermissionRepository
}

func NewServiceRolePermission(r interfaces.IRolePermissionRepository) interfaces.IRolePermissionService {
	return &serviceRolePermission{
		RolePermissionRepository: r,
	}
}

func (s *serviceRolePermission) GetByRoleId(roleId int32) (*dto.GetRolePermission, error) {
	repoRes, err := s.RolePermissionRepository.GetRolePermissionByRoleId(roleId)

	if err != nil {
		return nil, err
	}

	return &dto.GetRolePermission{
		Permission: p.GetPermissionRes{
			Id:          repoRes.Permission.PermissionID,
			Name:        repoRes.Permission.Name,
			Code:        repoRes.Permission.Code,
			Description: repoRes.Permission.Description,
		},
		Role: r.GetRoleRes{
			Id:          repoRes.Role.RoleID,
			RoleName:    repoRes.Role.RoleName,
			Description: repoRes.Permission.Description,
		},
	}, nil
}

func (s *serviceRolePermission) List() (*[]dto.GetRolePermissionRes, error) {
	repoRes, err := s.RolePermissionRepository.ListRolePermissions()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}
