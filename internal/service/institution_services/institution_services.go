package service

import (
	dto "optitech/internal/dto/institution_services"
	dto_services "optitech/internal/dto/services"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type service_institution_services struct {
	institutionServiceRepository interfaces.IInstitutionServiceRepository
}

func NewServiceInstitutionServices(r interfaces.IInstitutionServiceRepository) interfaces.IServicesInstitutionService {
	return &service_institution_services{
		institutionServiceRepository: r,
	}
}

func (s *service_institution_services) List(InstitutionID int32) (*[]dto_services.GetServiceRes, error) {

	return s.institutionServiceRepository.ListInstitutionServices(InstitutionID)

}

func (s *service_institution_services) Create(req *[]sq.CreateInstitutionServicesParams) error {
	return s.institutionServiceRepository.CreateInstitutionService(req)
}

func (s *service_institution_services) DeleteById(req dto.GetInstitutionServicesReq) (bool, error) {
	arg := &sq.DeleteInstitutionServiceByIdParams{
		ServiceID:     req.ServiceID,
		InstitutionID: req.InstitutionID,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	err := s.institutionServiceRepository.DeleteInstitutionServiceById(arg)

	if err != nil {
		return false, err
	}

	return true, err
}
func (s *service_institution_services) Exists(req *sq.ExistsInstitutionServiceParams) bool {
	return s.Exists(req)
}

func (s *service_institution_services) DeleteByInstitution(institutionID int32) (bool, error) {
	arg := &sq.DeleteInstitutionServicesByInstitutionParams{
		InstitutionID: institutionID,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	err := s.institutionServiceRepository.DeleteInstitutionServiceByInstitution(arg)

	if err != nil {
		return false, err
	}

	return true, err
}
