package service

import (
	p "optitech/internal/dto/permission"
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

func (s *serviceRolePermission) ListPermissionsByRoleId(roleId int32) (*[]p.GetPermissionRes, error) {
	repoRes, err := s.RolePermissionRepository.ListPermissionByRoleId(roleId)

	if err != nil {
		return nil, err
	}

	rolePermissions := make([]p.GetPermissionRes, len(*repoRes))

	for i, inst := range *repoRes {
		rolePermissions[i] = p.GetPermissionRes{
			Id:          inst.Permission.PermissionID,
			Name:        inst.Permission.Name,
			Code:        inst.Permission.Code,
			Description: inst.Permission.Description,
		}
	}

	return &rolePermissions, nil
}

func (s *serviceRolePermission) List() (*[]dto.GetRolePermissionRes, error) {
	repoRes, err := s.RolePermissionRepository.ListRolePermissions()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}
