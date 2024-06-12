package service

import (
	dto "optitech/internal/dto/institution_services"
	dtoServices "optitech/internal/dto/services"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceInstitutionService struct {
	institutionServiceRepository interfaces.IInstitutionServiceRepository
}

func NewServiceInstitutionServices(r interfaces.IInstitutionServiceRepository) interfaces.IServicesInstitutionService {
	return &serviceInstitutionService{
		institutionServiceRepository: r,
	}
}

func (s *serviceInstitutionService) List(InstitutionID int32) (*[]dtoServices.GetServiceRes, error) {

	return s.institutionServiceRepository.ListInstitutionServices(InstitutionID)

}

func (s *serviceInstitutionService) Create(req *[]sq.CreateInstitutionServicesParams) error {
	return s.institutionServiceRepository.CreateInstitutionService(req)
}

func (s *serviceInstitutionService) DeleteById(req dto.GetInstitutionServicesReq) (bool, error) {
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
func (s *serviceInstitutionService) Exists(req *sq.ExistsInstitutionServiceParams) bool {
	return s.Exists(req)
}

func (s *serviceInstitutionService) DeleteByInstitution(institutionID int32) (bool, error) {
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
