package repository

import (
	"context"
	dto "optitech/internal/dto/client_role"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryClientRole struct {
	ClientRoleRepository *sq.Queries
}

func (r *repositoryClientRole) GetByClientId(clientId int32) (*sq.GetClientRoleByClientIdRow, error) {
	ctx := context.Background()
	res, err := r.ClientRoleRepository.GetClientRoleByClientId(ctx, clientId)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func NewRepositoryClientRole(q *sq.Queries) interfaces.IClientRoleRepository {
	return &repositoryClientRole{
		ClientRoleRepository: q,
	}
}

func (r *repositoryClientRole) CreateClientRole(arg *sq.CreateClientRoleParams) (*dto.CreateClientRoleRes, error) {
	ctx := context.Background()
	res, err := r.ClientRoleRepository.CreateClientRole(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateClientRoleRes{
		Id:       res.ClientRoleID,
		ClientId: res.ClientID,
		RoleId:   res.RoleID,
	}, err
}

func (r *repositoryClientRole) DeleteClientRole(arg *sq.DeleteClientRoleByIdParams) error {
	ctx := context.Background()
	return r.ClientRoleRepository.DeleteClientRoleById(ctx, *arg)
}

func (r *repositoryClientRole) GetClientRole(id int64) (*dto.GetClientRoleRes, error) {
	ctx := context.Background()
	res, err := r.ClientRoleRepository.GetClientRole(ctx, id)

	if err != nil {
		return nil, err
	}

	return &dto.GetClientRoleRes{
		Id:       res.ClientRoleID,
		ClientId: res.ClientID,
		RoleId:   res.RoleID,
	}, nil
}

func (r *repositoryClientRole) ListClientRoles() (*[]dto.GetClientRoleRes, error) {
	ctx := context.Background()
	repoRes, err := r.ClientRoleRepository.ListClientRoles(ctx)

	if err != nil {
		return nil, err
	}

	ClientRoles := make([]dto.GetClientRoleRes, len(repoRes))
	for i, inst := range repoRes {
		ClientRoles[i] = dto.GetClientRoleRes{
			Id:       inst.ClientRoleID,
			ClientId: inst.ClientID,
			RoleId:   inst.RoleID,
		}
	}

	return &ClientRoles, nil
}

func (r *repositoryClientRole) UpdateClientRole(arg *sq.UpdateClientRoleByIdParams) error {
	ctx := context.Background()
	return r.ClientRoleRepository.UpdateClientRoleById(ctx, *arg)
}
