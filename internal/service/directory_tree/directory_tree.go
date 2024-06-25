package service

import (
	dto "optitech/internal/dto/directory_tree"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceDirectoryTree struct {
	directoryTreeRepository interfaces.IDirectoryRepositoy
}

func NewServicDirectory(r interfaces.IDirectoryRepositoy) interfaces.IDirectoryService {
	return &serviceDirectoryTree{
		directoryTreeRepository: r,
	}
}

func (s *serviceDirectoryTree) Get(req dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error) {
	return s.directoryTreeRepository.GetDirectroy(req.Id)
}

func (s *serviceDirectoryTree) Create(req *dto.CreateDirectoryTreeReq) (*dto.CreateDirectoryTreeRes, error) {
	repoReq := &sq.CreateDirectoryTreeParams{
		ParentID:  pgtype.Int4{Int32: int32(req.ParentID), Valid: true},
		Name:      pgtype.Text{String: req.Name, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := s.directoryTreeRepository.CreateDirectoy(repoReq)
	if err != nil {
		return nil, err
	}

	// TODO: Ubicar el token si es requerido

	return r, nil
}
