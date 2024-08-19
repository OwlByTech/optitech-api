package service

import (
	"github.com/jackc/pgx/v5/pgtype"
	dto "optitech/internal/dto/asesor"
	ddto "optitech/internal/dto/directory_tree"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"
)

type serviceAsesor struct {
	asesorRepository     interfaces.IAsesorRepository
	directoryTreeService interfaces.IDirectoryService
}

func NewServiceAsesor(r interfaces.IAsesorRepository, serviceDirectory interfaces.IDirectoryService) interfaces.IAsesorService {
	return &serviceAsesor{
		asesorRepository:     r,
		directoryTreeService: serviceDirectory,
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
	_, err = s.directoryTreeService.Create(&ddto.CreateDirectoryTreeReq{
		AsesorID: r.Id,
		Name:     "/",
	})

	if err != nil {
		return nil, err
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
