package repository

import (
	"context"
	dto_client "optitech/internal/dto/client"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repository_institution_client struct {
	insititutionClientRepository *sq.Queries
}

func NewRepositoryInstitutionClient(q *sq.Queries) interfaces.IInstitutionClientRepository {
	return &repository_institution_client{
		insititutionClientRepository: q,
	}
}

func (r *repository_institution_client) ListInstitutionClient(InstitutionID int32) (*[]dto_client.GetClientRes, error) {
	ctx := context.Background()
	repoRes, err := r.insititutionClientRepository.ListInstitutionClients(ctx, InstitutionID)
	if err != nil {
		return nil, err
	}

	institution_client := make([]dto_client.GetClientRes, len(repoRes))
	for i, inst := range repoRes {
		institution_client[i] = dto_client.GetClientRes{
			ClientID:  inst.ClientID,
			GivenName: inst.GivenName,
			Surname:   inst.Surname,
		}
	}

	return &institution_client, nil
}

func (r *repository_institution_client) ExistsInstitutionClient(arg *sq.ExistsInstitutionClientParams) bool {
	ctx := context.Background()
	_, err := r.insititutionClientRepository.ExistsInstitutionClient(ctx, *arg)
	if err != nil {
		return false
	}
	return true

}

func (r *repository_institution_client) CreateInstitutionClient(arg *[]sq.CreateInstitutionClientParams) error {
	ctx := context.Background()
	_, err := r.insititutionClientRepository.CreateInstitutionClient(ctx, *arg)
	return err
}

func (r *repository_institution_client) DeleteInstitutionClientById(arg *sq.DeleteinstInstitutionClientByClientAndInstitutionParams) error {
	ctx := context.Background()
	return r.insititutionClientRepository.DeleteinstInstitutionClientByClientAndInstitution(ctx, *arg)

}
func (r *repository_institution_client) DeleteInstitutionClientByInstitution(arg *sq.DeleteInstitutionClientByInstitutionParams) error {
	ctx := context.Background()
	return r.insititutionClientRepository.DeleteInstitutionClientByInstitution(ctx, *arg)

}
