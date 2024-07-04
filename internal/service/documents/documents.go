package service

import (
	dto "optitech/internal/dto/document"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
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

func (s *serviceDocument) Create(req *dto.CreateDocumentReq) (*dto.CreateDocumentRes, error) {

	repoReq := &sq.CreateDocumentParams{
		DirectoryID: req.DirectoryId,
		FormatID:    pgtype.Int4{Int32: req.FormatId, Valid: true},
		FileRute:    req.FileRute,
		Status:      sq.Status(req.Status),
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	repoRes, err := s.documentRepository.CreateDocument(repoReq)

	if err != nil {
		return nil, err
	}

	// TODO: RETURNS EMPTY JSON

	document := &dto.CreateDocumentRes{
		Id: repoRes.Id,
	}
	return document, err

}
