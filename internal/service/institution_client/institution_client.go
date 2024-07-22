package service

import (
	dtoClient "optitech/internal/dto/client"
	dto "optitech/internal/dto/institution_client"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"slices"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceInstitutionClient struct {
	institutionClientRepository interfaces.IInstitutionClientRepository
}

func NewServiceInstitutionClient(r interfaces.IInstitutionClientRepository) interfaces.IInstitutionClientService {
	return &serviceInstitutionClient{
		institutionClientRepository: r,
	}
}

func (s *serviceInstitutionClient) List(InstitutionID int32) (*[]dtoClient.GetClientRes, error) {
	return s.institutionClientRepository.ListInstitutionClient(InstitutionID)
}
func (s *serviceInstitutionClient) Exists(req *sq.ExistsInstitutionClientParams) bool {
	return s.institutionClientRepository.ExistsInstitutionClient(req)
}

func (s *serviceInstitutionClient) Create(req *[]sq.CreateInstitutionClientParams) error {
	return s.institutionClientRepository.CreateInstitutionClient(req)
}

func (s *serviceInstitutionClient) DeleteById(req *dto.GetInstitutionClientReq) error {
	arg := &sq.DeleteInstitutionByClientParams{
		InstitutionID: req.InstitutionId,
		ClientID:      req.ClientId,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	return s.institutionClientRepository.DeleteInstitutionClientById(arg)

}
func (s *serviceInstitutionClient) Update(req dto.UpdateInstitutionClientReq) (bool, error) {
	res, err := s.List(req.InstitutionID)

	if err != nil {
		return false, err
	}

	listClients := []int32{}
	for _, client := range *res {
		listClients = append(listClients, client.Id)
	}
	var listCreate []sq.CreateInstitutionClientParams
	deleteClient, createClient := FindMissing(listClients, req.Clients)

	for _, client := range deleteClient {
		if err := s.DeleteById(&dto.GetInstitutionClientReq{InstitutionId: req.InstitutionID, ClientId: client}); err != nil {
			return false, err
		}
	}
	for _, client := range createClient {
		if s.Exists(&sq.ExistsInstitutionClientParams{InstitutionID: req.InstitutionID, ClientID: client}) {
			if err := s.Recover(&sq.RecoverInstitutionClientParams{InstitutionID: req.InstitutionID, ClientID: client}); err != nil {
				return false, err
			}
		} else {
			listCreate = append(listCreate, sq.CreateInstitutionClientParams{InstitutionID: req.InstitutionID, ClientID: client, CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true}})

		}
	}
	if err := s.Create(&listCreate); err != nil {
		return false, err
	}

	return true, nil

}

func (s *serviceInstitutionClient) DeleteByInstitution(InstitutionID int32) error {
	arg := &sq.DeleteInstitutionClientParams{
		InstitutionID: InstitutionID,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	return s.institutionClientRepository.DeleteInstitutionClientByInstitution(arg)

}

func (s *serviceInstitutionClient) Recover(req *sq.RecoverInstitutionClientParams) error {
	return s.institutionClientRepository.RecoverInstitutionClient(req)
}

func FindMissing(list []int32, listTwo []int32) ([]int32, []int32) {
	missingInList := []int32{}
	missingInListTwo := []int32{}
	for _, element := range list {
		if slices.Index(list, element) == -1 {
			missingInList = append(missingInList, element)
		}
	}
	for _, element := range listTwo {
		if slices.Index(list, element) == -1 {
			missingInListTwo = append(missingInListTwo, element)
		}
	}

	return missingInList, missingInListTwo
}
