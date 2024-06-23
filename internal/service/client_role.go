package service

import (
	dto "optitech/internal/dto/client_role"
	r "optitech/internal/dto/roles"
	"optitech/internal/interfaces"
	"optitech/internal/sqlc"
)

type serviceClientRole struct {
	ClientRoleRepository interfaces.IClientRoleRepository
}

func NewServiceClientRole(r interfaces.IClientRoleRepository) interfaces.IClientRoleService {
	return &serviceClientRole{
		ClientRoleRepository: r,
	}
}

func (s *serviceClientRole) Create(arg *sqlc.CreateClientRoleParams) (*dto.CreateClientRoleRes, error) {
	repoRes, err := s.ClientRoleRepository.CreateClientRole(arg)

	if err != nil {
		return nil, err
	}

	return repoRes, nil
}

func (s *serviceClientRole) ListByClientId(clientId int32) (*[]r.GetRoleRes, error) {
	repoRes, err := s.ClientRoleRepository.ListByClientId(clientId)

	if err != nil {
		return nil, err
	}

	roles := make([]r.GetRoleRes, len(*repoRes))

	for i, inst := range *repoRes {
		roles[i] =
			r.GetRoleRes{
				Id:          inst.Role.RoleID,
				RoleName:    inst.Role.RoleName,
				Description: inst.Role.Description,
			}
	}

	return &roles, nil
}

func (s *serviceClientRole) List() (*[]dto.GetClientRoleRes, error) {
	repoRes, err := s.ClientRoleRepository.ListClientRoles()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}

func (s *serviceClientRole) Delete(arg *sqlc.DeleteClientRoleByIdParams) error {
	return s.ClientRoleRepository.DeleteClientRole(arg)
}
