package service

import (
	dto "optitech/internal/dto/institution_services"
	dtoServices "optitech/internal/dto/services"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"slices"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceInstitutionService struct {
	institutionServiceRepository interfaces.IInstitutionServiceRepository
}

func NewServiceInstitutionServices(r interfaces.IInstitutionServiceRepository) interfaces.IServiceInstitutionService {
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

func (s *serviceInstitutionService) Update(req *dto.UpdateInstitutionServicesReq) bool {
	res, err := s.List(req.InstitutionID)
	if err != nil {
		return false
	}
	var listValid []int32
	var listCreate []sq.CreateInstitutionServicesParams
	for _, service := range *res {
		if slices.Index(req.Services, service.ServiceID) == -1 {
			if s.DeleteById(&dto.GetInstitutionServicesReq{InstitutionID: req.InstitutionID, ServiceID: service.ServiceID}) != nil {
				return false
			}
		} else {
			listValid = append(listValid, service.ServiceID)
		}
	}
	for _, service := range listValid {
		if s.Exists(&sq.ExistsInstitutionServiceParams{ServiceID: service, InstitutionID: req.InstitutionID}) {
			if err := s.Recover(&sq.RecoverInstitutionServiceParams{ServiceID: service, InstitutionID: req.InstitutionID}); err != nil {
				return false
			}
		} else {
			listCreate = append(listCreate, sq.CreateInstitutionServicesParams{InstitutionID: req.InstitutionID, ServiceID: service})
		}
	}
	if err := s.Create(&listCreate); err != nil {
		return false
	}

	return true

}

func (s *serviceInstitutionService) DeleteById(req *dto.GetInstitutionServicesReq) error {
	arg := &sq.DeleteInstitutionServiceByIdParams{
		ServiceID:     req.ServiceID,
		InstitutionID: req.InstitutionID,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	return s.institutionServiceRepository.DeleteInstitutionServiceById(arg)

}
func (s *serviceInstitutionService) Exists(req *sq.ExistsInstitutionServiceParams) bool {
	return s.institutionServiceRepository.ExistsInstitutionService(req)
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
func (s *serviceInstitutionService) Recover(req *sq.RecoverInstitutionServiceParams) error {
	return s.institutionServiceRepository.RecoverInstitutionService(req)
}
