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
	var parentID pgtype.Int8
	if req.ParentID == 0 {
		parentID.Valid = false
	} else {
		parentID.Int64 = int64(req.ParentID)
		parentID.Valid = true
	}

	repoReq := &sq.CreateDirectoryTreeParams{
		ParentID:      parentID,
		Name:          pgtype.Text{String: req.Name, Valid: true},
		InstitutionID: pgtype.Int4{Int32: int32(req.InstitutionID), Valid: true},
		CreatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := s.directoryTreeRepository.CreateDirectory(repoReq)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *serviceDirectoryTree) List() (*[]dto.GetDirectoryTreeRes, error) {
	repoRes, err := s.directoryTreeRepository.ListDirectory()
	if err != nil {
		return nil, err
	}
	return repoRes, nil
}

func (s *serviceDirectoryTree) Delete(req dto.GetDirectoryTreeReq) (bool, error) {
	repoReq := &sq.DeleteDirectoryTreeByIdParams{
		DirectoryID: req.Id,
		DeletedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if err := s.directoryTreeRepository.DeleteDirectory(repoReq); err != nil {
		return false, pgtype.ErrScanTargetTypeChanged
	}

	return true, nil
}
