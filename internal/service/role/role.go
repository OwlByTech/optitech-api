package service

import (
	dto "optitech/internal/dto/roles"
	"optitech/internal/interfaces"
)

type serviceRole struct {
	RoleRepository interfaces.IRoleRepository
}

func NewServiceRole(r interfaces.IRoleRepository) interfaces.IRoleService {
	return &serviceRole{
		RoleRepository: r,
	}
}

func (s *serviceRole) List() (*[]dto.GetRoleRes, error) {
	repoRes, err := s.RoleRepository.ListRoles()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}
