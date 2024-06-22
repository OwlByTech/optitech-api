package service

import (
	dto "optitech/internal/dto/client_role"
	"optitech/internal/interfaces"
)

type serviceClientRole struct {
	ClientRoleRepository interfaces.IClientRoleRepository
}

func NewServiceClientRole(r interfaces.IClientRoleRepository) interfaces.IClientRoleService {
	return &serviceClientRole{
		ClientRoleRepository: r,
	}
}

func (s *serviceClientRole) List() (*[]dto.GetClientRoleRes, error) {
	repoRes, err := s.ClientRoleRepository.ListClientRoles()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}
