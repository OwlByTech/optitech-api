package service

import (
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type serviceDirectoryRole struct {
	directoryRoleRepository interfaces.IDirectoryRoleRepository
}

func NewServiceDirectory(r interfaces.IDirectoryRoleRepository) interfaces.IDirectoryRoleService {
	return &serviceDirectoryRole{
		directoryRoleRepository: r,
	}
}

func (s *serviceDirectoryRole) Create(req *[]sq.CreateDirectoryRoleParams) error {
	return s.directoryRoleRepository.CreateDirectoryRole(req)
}
