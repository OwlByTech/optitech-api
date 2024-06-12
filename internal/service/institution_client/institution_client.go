package service

import (
	dtoClient "optitech/internal/dto/client"
	dto "optitech/internal/dto/institution_client"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
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
	return s.Exists(req)
}

func (s *serviceInstitutionClient) Create(req *[]sq.CreateInstitutionClientParams) error {
	return s.institutionClientRepository.CreateInstitutionClient(req)
}

func (s *serviceInstitutionClient) DeleteById(req *dto.GetInstitutionClientReq) (bool, error) {
	arg := &sq.DeleteinstInstitutionClientByClientAndInstitutionParams{
		InstitutionID: req.InstitutionID,
		ClientID:      req.ClientID,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	err := s.institutionClientRepository.DeleteInstitutionClientById(arg)

	if err != nil {
		return false, err
	}

	return true, err
}

func (s *serviceInstitutionClient) DeleteByInstitution(InstitutionID int32) (bool, error) {
	arg := &sq.DeleteInstitutionClientByInstitutionParams{
		InstitutionID: InstitutionID,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	err := s.institutionClientRepository.DeleteInstitutionClientByInstitution(arg)

	if err != nil {
		return false, err
	}

	return true, err
}
