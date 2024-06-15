package repository

import (
	"context"
	dto "optitech/internal/dto/client"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryClient struct {
	clientRepository *sq.Queries
}

func NewRepositoryClient(q *sq.Queries) interfaces.IClientRepository {
	return &repositoryClient{
		clientRepository: q,
	}
}

func (r *repositoryClient) GetClient(institutionID int64) (*dto.GetClientRes, error) {
	ctx := context.Background()

	repoRes, err := r.clientRepository.GetClient(ctx, (institutionID))

	if err != nil {
		return nil, err
	}

	return &dto.GetClientRes{
		Id:        repoRes.ClientID,
		GivenName: repoRes.GivenName,
		Surname:   repoRes.Surname,
		Email:     repoRes.Email,
	}, nil
}

func (r *repositoryClient) CreateClient(arg *sq.CreateClientParams) (*dto.CreateClientRes, error) {
	ctx := context.Background()

	res, err := r.clientRepository.CreateClient(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateClientRes{
		Id:        res.ClientID,
		GivenName: res.GivenName,
		Surname:   res.Surname,
		Email:     res.Email,
	}, nil
}

func (r *repositoryClient) UpdateClient(arg *sq.UpdateClientByIdParams) error {
	ctx := context.Background()
	return r.clientRepository.UpdateClientById(ctx, *arg)
}

func (r *repositoryClient) ListClient() (*[]dto.GetClientRes, error) {
	return nil, nil
}

func (r *repositoryClient) DeleteClient(arg *sq.DeleteClientByIdParams) error {
	return nil
}
