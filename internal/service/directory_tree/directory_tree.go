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
	documentService         interfaces.IDocumentService
}

func NewServiceDirectory(r interfaces.IDirectoryRepository, documentService interfaces.IDocumentService) interfaces.IDirectoryService {
	return &serviceDirectoryTree{
		directoryTreeRepository: r,
		documentService:         documentService,
	}
}

func (s *serviceDirectoryTree) Get(req dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error) {
	return s.directoryTreeRepository.GetDirectory(req.Id)
}

func (s *serviceDirectoryTree) Create(req *dto.CreateDirectoryTreeReq) (*dto.CreateDirectoryTreeRes, error) {
	var parentID pgtype.Int4
	if req.ParentID == 0 {
		parentID = pgtype.Int4{Valid: false}
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

func (s *serviceDirectoryTree) List() (*[]dto.GetDirectoryTreeRes, error) {
	repoRes, err := s.directoryTreeRepository.ListDirectory()
	if err != nil {
		return nil, err
	}
	return repoRes, nil
}
func (s *serviceDirectoryTree) ListByParent(req dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeByParentRes, error) {
	repoRes, err := s.directoryTreeRepository.ListDirectoryByParent(int32(req.Id))
	if err != nil {
		return nil, err
	}
	documents, err := s.documentService.ListByDirectory(dto.GetDirectoryTreeReq{Id: int64(req.Id)})
	if err != nil {
		return nil, err
	}
	directory, err := s.Get(dto.GetDirectoryTreeReq{Id: int64(req.Id)})
	if err != nil {
		return nil, err
	}

	return &dto.GetDirectoryTreeByParentRes{
		Id:        directory.Id,
		Name:      directory.Name,
		ParentID:  directory.ParentID,
		Directory: *repoRes, Document: *documents}, nil
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
