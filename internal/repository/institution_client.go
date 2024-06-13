package repository

import (
	"context"
	dtoClient "optitech/internal/dto/client"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryInstitutionClient struct {
	insititutionClientRepository *sq.Queries
}

func NewRepositoryInstitutionClient(q *sq.Queries) interfaces.IInstitutionClientRepository {
	return &repositoryInstitutionClient{
		insititutionClientRepository: q,
	}
}

func (r *repositoryInstitutionClient) ListInstitutionClient(InstitutionID int32) (*[]dtoClient.GetClientRes, error) {
	ctx := context.Background()
	repoRes, err := r.insititutionClientRepository.ListInstitutionClients(ctx, InstitutionID)
	if err != nil {
		return nil, err
	}

	institutionClient := make([]dtoClient.GetClientRes, len(repoRes))
	for i, inst := range repoRes {
		institutionClient[i] = dtoClient.GetClientRes{
			ClientID:  inst.ClientID,
			GivenName: inst.GivenName,
			Surname:   inst.Surname,
		}
	}

	return &institutionClient, nil
}

func (r *repositoryInstitutionClient) ExistsInstitutionClient(arg *sq.ExistsInstitutionClientParams) bool {
	ctx := context.Background()
	_, err := r.insititutionClientRepository.ExistsInstitutionClient(ctx, *arg)
	if err != nil {
		return false
	}
	return true

}
func (r *repositoryInstitutionClient) RecoverInstitutionClient(arg *sq.RecoverInstitutionClientParams) error {
	ctx := context.Background()
	return r.insititutionClientRepository.RecoverInstitutionClient(ctx, *arg)

}

func (r *repositoryInstitutionClient) CreateInstitutionClient(arg *[]sq.CreateInstitutionClientParams) error {
	ctx := context.Background()
	_, err := r.insititutionClientRepository.CreateInstitutionClient(ctx, *arg)
	return err
}

func (r *repositoryInstitutionClient) DeleteInstitutionClientById(arg *sq.DeleteinstInstitutionClientByClientAndInstitutionParams) error {
	ctx := context.Background()
	return r.insititutionClientRepository.DeleteinstInstitutionClientByClientAndInstitution(ctx, *arg)

}
func (r *repositoryInstitutionClient) DeleteInstitutionClientByInstitution(arg *sq.DeleteInstitutionClientByInstitutionParams) error {
	ctx := context.Background()
	return r.insititutionClientRepository.DeleteInstitutionClientByInstitution(ctx, *arg)

}
