package service

import (
	dto "optitech/internal/dto/permission"
	"optitech/internal/interfaces"
)

type servicePermission struct {
	permissionRepository interfaces.IPermissionRepository
}

func NewServicePermission(r interfaces.IPermissionRepository) interfaces.IPermissionService {
	return &servicePermission{
		permissionRepository: r,
	}
}

func (s *servicePermission) List() (*[]dto.GetPermissionRes, error) {
	repoRes, err := s.permissionRepository.ListPermissions()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}
