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
	arg := &sq.DeleteinstInstitutionClientByClientAndInstitutionParams{
		InstitutionID: req.InstitutionID,
		ClientID:      req.ClientID,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	return s.institutionClientRepository.DeleteInstitutionClientById(arg)

}
func (s *serviceInstitutionClient) Update(req dto.UpdateInstitutionClientReq) bool {
	res, err := s.List(req.InstitutionID)
	if err != nil {
		return false
	}
	var listValid []int32
	var listCreate []sq.CreateInstitutionClientParams
	for _, client := range *res {
		if slices.Index(req.Clients, client.ClientID) == -1 {
			if s.DeleteById(&dto.GetInstitutionClientReq{InstitutionID: req.InstitutionID, ClientID: client.ClientID}) != nil {
				return false
			}
		} else {
			listValid = append(listValid, client.ClientID)
		}
	}
	for _, client := range listValid {
		if s.Exists(&sq.ExistsInstitutionClientParams{InstitutionID: req.InstitutionID, ClientID: client}) {
			if err := s.Recover(&sq.RecoverInstitutionClientParams{InstitutionID: req.InstitutionID, ClientID: client}); err != nil {
				return false
			}
		} else {
			listCreate = append(listCreate, sq.CreateInstitutionClientParams{InstitutionID: req.InstitutionID, ClientID: client, CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true}})
		}
	}
	if err := s.Create(&listCreate); err != nil {
		return false
	}

	return true

}

func (s *serviceInstitutionClient) DeleteByInstitution(InstitutionID int32) error {
	arg := &sq.DeleteInstitutionClientByInstitutionParams{
		InstitutionID: InstitutionID,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	return s.institutionClientRepository.DeleteInstitutionClientByInstitution(arg)

}

func (s *serviceInstitutionClient) Recover(req *sq.RecoverInstitutionClientParams) error {
	return s.institutionClientRepository.RecoverInstitutionClient(req)
}
