package service

import (
	dto_client "optitech/internal/dto/client"
	dto "optitech/internal/dto/institution_client"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type service_institution_client struct {
	institutionClientRepository interfaces.IInstitutionClientRepository
}

func NewServiceInstitutionClient(r interfaces.IInstitutionClientRepository) interfaces.IInstitutionClientService {
	return &service_institution_client{
		institutionClientRepository: r,
	}
}

func (s *service_institution_client) List(InstitutionID int32) (*[]dto_client.GetClientRes, error) {

	return s.institutionClientRepository.ListInstitutionClient(InstitutionID)

}
func (s *service_institution_client) Exists(req *sq.ExistsInstitutionClientParams) bool {
	return s.Exists(req)
}

func (s *service_institution_client) Create(req *[]sq.CreateInstitutionClientParams) error {
	return s.institutionClientRepository.CreateInstitutionClient(req)
}

func (s *service_institution_client) DeleteById(req *dto.GetInstitutionClientReq) (bool, error) {
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

func (s *service_institution_client) DeleteByInstitution(InstitutionID int32) (bool, error) {
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
