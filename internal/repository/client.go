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

func (r *repositoryClient) GetClient(clientID int32) (*dto.GetClientRes, error) {
	ctx := context.Background()

	repoRes, err := r.clientRepository.GetClient(ctx, (clientID))

	if err != nil {
		return nil, err
	}

	return &dto.GetClientRes{
		Id:        repoRes.ClientID,
		GivenName: repoRes.GivenName,
		Photo:     repoRes.Photo.String,
		Status:    dto.StatusClient(repoRes.Status),
		Surname:   repoRes.Surname,
		Email:     repoRes.Email,
	}, nil
}

func (r *repositoryClient) CreateClient(arg *sq.CreateClientParams) (*dto.CreateClient, error) {
	ctx := context.Background()

	res, err := r.clientRepository.CreateClient(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateClient{
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
func (r *repositoryClient) UpdateStatusClient(arg *sq.UpdateClientStatusByIdParams) error {
	ctx := context.Background()
	return r.clientRepository.UpdateClientStatusById(ctx, *arg)
}

func (r *repositoryClient) UpdatePhotoClient(arg *sq.UpdateClientPhotoParams) error {
	ctx := context.Background()
	return r.clientRepository.UpdateClientPhoto(ctx, *arg)
}

func (r *repositoryClient) ListClient() (*[]dto.GetClientRes, error) {
	ctx := context.Background()
	repoRes, err := r.clientRepository.ListClients(ctx)

	if err != nil {
		return nil, err
	}

	clients := make([]dto.GetClientRes, len(repoRes))
	for i, inst := range repoRes {
		clients[i] = dto.GetClientRes{
			Id:        inst.ClientID,
			GivenName: inst.GivenName,
			Surname:   inst.Surname,
			Email:     inst.Email,
		}
	}
	return &clients, nil
}

func (r *repositoryClient) DeleteClient(arg *sq.DeleteClientByIdParams) error {
	ctx := context.Background()
	return r.clientRepository.DeleteClientById(ctx, *arg)
}

func (r *repositoryClient) LoginClient(email string) (*sq.Client, error) {
	ctx := context.Background()
	res, err := r.clientRepository.GetClientByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &res, nil

}
