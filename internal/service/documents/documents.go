package service

import (
	"fmt"
	drdto "optitech/internal/dto/directory_tree"
	dto "optitech/internal/dto/document"
	"optitech/internal/interfaces"
	digitalOcean "optitech/internal/service/digital_ocean"
	sq "optitech/internal/sqlc"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceDocument struct {
	documentRepository interfaces.IDocumentRepository
}

func NewServiceDocument(d interfaces.IDocumentRepository) interfaces.IDocumentService {
	return &serviceDocument{
		documentRepository: d,
	}
}

func (s *serviceDocument) Get(req dto.GetDocumentReq) (*dto.GetDocumentRes, error) {
	return s.documentRepository.GetDocument(req.Id)
}

func (s *serviceDocument) ListByDirectory(req drdto.GetDirectoryTreeReq) (*[]dto.GetDocumentRes, error) {
	return s.documentRepository.ListDocumentByDirectory(int32(req.Id))
}

func (s *serviceDocument) Create(req *dto.CreateDocumentReq) (*dto.CreateDocumentRes, error) {

	institutionName, err := s.documentRepository.GetInstitutionByDocumentId(int64(req.DirectoryId))

	if err != nil {
		return nil, err
	}

	repoReq := &sq.CreateDocumentParams{
		DirectoryID: req.DirectoryId,
		FormatID:    pgtype.Int4{Int32: req.FormatId, Valid: false},
		Name:        req.File.Filename,
		Status:      sq.Status(req.Status),
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	rute := fmt.Sprintf("%s%s", strconv.FormatInt(time.Now().UTC().UnixMicro(), 10), filepath.Ext(req.File.Filename))

	fileRute, err := digitalOcean.UploadDocument(req.File, rute, institutionName.Institution.InstitutionName)

	if err != nil {
		return nil, err
	}

	if req.FormatId > 0 {
		repoReq.FormatID = pgtype.Int4{Int32: req.FormatId, Valid: true}
	}
	repoReq.FileRute = fileRute
	repoRes, err := s.documentRepository.CreateDocument(repoReq)

	if err != nil {
		return nil, err
	}

	return repoRes, err
}

func (s *serviceDocument) DownloadDocumentById(req dto.GetDocumentReq) (string, error) {

	exist, err := s.documentRepository.ExistsDocuments(req.Id)
	if err != nil {
		return "", err
	}

	if exist {
		return "", fmt.Errorf("the document does not exist")
	}

	document, err := s.documentRepository.DownloadDocumentById(req.Id)
	if err != nil {
		return "", err
	}

	route, err := digitalOcean.DownloadDocument(document.FileRute, document.InstitutionName)
	if err != nil {
		return "", err
	}
	return route, err
}

func (s *serviceDocument) DeleteDocument(req dto.GetDocumentReq) (bool, error) {
	repoReq := &sq.DeleteDocumentByIdParams{
		DocumentID: req.Id,
		DeletedAt:  pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if err := s.documentRepository.DeleteDocument(repoReq); err != nil {
		return false, pgtype.ErrScanTargetTypeChanged
	}

	return true, nil
}

func (s *serviceDocument) UpdateDocument(req *dto.UpdateDocumentReq) (bool, error) {

	repoReq := &sq.UpdateDocumentNameByIdParams{
		DocumentID: req.Id,
		Name:       req.Name,
		UpdatedAt:  pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if err := s.documentRepository.UpdateDocument(repoReq); err != nil {
		return false, nil
	}
	return true, nil
}
