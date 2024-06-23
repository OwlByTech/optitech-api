package service

import (
	"github.com/jackc/pgx/v5/pgtype"
	dto "optitech/internal/dto/asesor"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"
)

type serviceAsesor struct {
	clientRepository interfaces.IAsesorRepository
}

func NewServiceAsesor(r interfaces.IAsesorRepository) interfaces.IAsesorService {
	return &serviceAsesor{
		clientRepository: r,
	}
}

func (s *serviceAsesor) Get(req dto.GetAsesorReq) (*dto.GetAsesorRes, error) {
	return s.clientRepository.GetAsesor(req.Id)
}

func (s *serviceAsesor) Create(req *dto.CreateAsesorReq) (*dto.CreateAsesorRes, error) {
	repoReq := &sq.CreateAsesorParams{
		AsesorID:  req.ClientId,
		Photo:     req.Photo,
		About:     req.About,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := s.clientRepository.CreateAsesor(repoReq)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *serviceAsesor) Update(req *dto.UpdateAsesorReq) (bool, error) {
	repoReq := &sq.UpdateAsesorByIdParams{
		AsesorID:  req.AsesorID,
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if req.Photo != "" {
		repoReq.Photo = req.Photo
	}

	if req.About != "" {
		repoReq.About = req.About
	}

	err := s.clientRepository.UpdateAsesor(repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *serviceAsesor) List() (*[]dto.GetAsesorRes, error) {
	repoRes, err := s.clientRepository.ListAsesor()
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

	if err := s.clientRepository.DeleteAsesor(repoReq); err != nil {
		return false, pgtype.ErrScanTargetTypeChanged
	}

	return true, nil
}
