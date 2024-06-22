package service

import (
	dto "optitech/internal/dto/role_permission"
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

func (s *serviceRolePermission) List() (*[]dto.GetRolePermissionRes, error) {
	repoRes, err := s.RolePermissionRepository.ListRolePermissions()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}
