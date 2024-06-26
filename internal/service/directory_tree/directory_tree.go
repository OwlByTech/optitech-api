package service

import (
	dto "optitech/internal/dto/directory_tree"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceDirectoryTree struct {
	directoryTreeRepository interfaces.IDirectoryRepository
}

func NewServiceDirectory(r interfaces.IDirectoryRepository) interfaces.IDirectoryService {
	return &serviceDirectoryTree{
		directoryTreeRepository: r,
	}
}

func (s *serviceDirectoryTree) Get(req dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error) {
	return s.directoryTreeRepository.GetDirectory(req.Id)
}

func (s *serviceDirectoryTree) Create(req *dto.CreateDirectoryTreeReq) (*dto.CreateDirectoryTreeRes, error) {
	var parentID pgtype.Int4
	if req.ParentID == 0 {
		parentID = pgtype.Int4{Valid: false} // Asignar null si ParentID es 0
	} else {
		parentID = pgtype.Int4{Int32: int32(req.ParentID), Valid: true}
	}

	repoReq := &sq.CreateDirectoryTreeParams{
		ParentID:  parentID,
		Name:      pgtype.Text{String: req.Name, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := s.directoryTreeRepository.CreateDirectory(repoReq)
	if err != nil {
		return nil, err
	}

	return r, nil
}
