package service

import (
	dto "optitech/internal/dto/asesor"
	ddto "optitech/internal/dto/directory_tree"
	sdto "optitech/internal/dto/services"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceAsesor struct {
	asesorRepository     interfaces.IAsesorRepository
	directoryTreeService interfaces.IDirectoryService
	servicesRepository   interfaces.IService
}

func NewServiceAsesor(r interfaces.IAsesorRepository, serviceDirectory interfaces.IDirectoryService, serviceServices interfaces.IService) interfaces.IAsesorService {
	return &serviceAsesor{
		asesorRepository:     r,
		directoryTreeService: serviceDirectory,
		servicesRepository:   serviceServices,
	}
}

func (s *serviceAsesor) Get(req dto.GetAsesorReq) (*dto.GetAsesorRes, error) {
	return s.asesorRepository.GetAsesor(req.Id)
}

func (s *serviceAsesor) Create(req *dto.CreateAsesorReq) (*dto.CreateAsesorRes, error) {
	repoReq := &sq.CreateAsesorParams{
		AsesorID:  req.ClientId,
		About:     pgtype.Text{String: req.About, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := s.asesorRepository.CreateAsesor(repoReq)
	if err != nil {
		return nil, err
	}

	directoryTreeRoot := &ddto.CreateDirectoryTreeReq{
		AsesorID: r.Id,
		Name:     "/",
	}

	rootDirectoryID, err := s.directoryTreeService.Create(directoryTreeRoot)
	if err != nil {
		return nil, err
	}

	var services *[]sdto.GetServiceRes
	services, err = s.servicesRepository.List()
	if err != nil {
		return nil, err
	}

	for _, service := range *services {

		serviceDirectoryReq := &ddto.CreateDirectoryTreeReq{
			ParentID: rootDirectoryID.DirectoryId,
			Name:     service.Name,
			AsesorID: r.Id,
		}

		_, err = s.directoryTreeService.Create(serviceDirectoryReq)
		if err != nil {
			return nil, err
		}
	}

	return r, nil
}

func (s *serviceAsesor) Update(req *dto.UpdateAsesorReq) (bool, error) {
	repoReq := &sq.UpdateAsesorByIdParams{
		AsesorID:  req.AsesorID,
		About:     pgtype.Text{String: req.About, Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	err := s.asesorRepository.UpdateAsesor(repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *serviceAsesor) List() (*[]dto.GetAsesorRes, error) {
	repoRes, err := s.asesorRepository.ListAsesor()
	if err != nil {
		return nil, err
	}
	return repoRes, nil
}

func (s *serviceAsesor) Delete(req dto.GetAsesorReq) (bool, error) {
	repoReq := &sq.DeleteAsesorByIdParams{
		AsesorID:  req.Id,
		DeletedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if err := s.asesorRepository.DeleteAsesor(repoReq); err != nil {
		return false, pgtype.ErrScanTargetTypeChanged
	}

	return true, nil
}
