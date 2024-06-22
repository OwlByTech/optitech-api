package service

import (
	c "optitech/internal/dto/client"
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

func (s *serviceClientRole) GetByClientId(clientId int32) (*dto.GetClientRole, error) {
	repoRes, err := s.ClientRoleRepository.GetByClientId(clientId)

	if err != nil {
		return nil, err
	}

	return &dto.GetClientRole{
		Client: c.GetClientRes{
			Id:        repoRes.Client.ClientID,
			GivenName: repoRes.Client.GivenName,
			Surname:   repoRes.Client.Surname,
			Email:     repoRes.Client.Email,
		},
		Role: r.GetRoleRes{
			Id:          repoRes.Role.RoleID,
			RoleName:    repoRes.Role.RoleName,
			Description: repoRes.Role.Description,
		},
	}, nil
}

func (s *serviceClientRole) List() (*[]dto.GetClientRoleRes, error) {
	repoRes, err := s.ClientRoleRepository.ListClientRoles()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}
