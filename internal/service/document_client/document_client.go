package service

import (
	dto "optitech/internal/dto/document_client"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceDocumentClient struct {
	documentClientRepository interfaces.IDocumentClientRepository
}

func NewServiceDocumentClient(f interfaces.IDocumentClientRepository) interfaces.IDocumentClientService {
	return &serviceDocumentClient{
		documentClientRepository: f,
	}
}

func (s *serviceDocumentClient) GetDocumentClient(req dto.GetDocumentClientReq) (*dto.GetDocumentClientRes, error) {
	return s.documentClientRepository.GetDocumentClient(req.Id)
}

func (s *serviceDocumentClient) CreateDocumentClient(req *dto.CreateDocumentClientReq) (*dto.CreateDocumentClientRes, error) {

	repoReq := &sq.CreateDocumentClientParams{
		ClientID:   req.ClientId,
		DocumentID: req.DocumentId,
		Action:     sq.Action(req.Action),
		CreatedAt:  pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	repoRes, err := s.documentClientRepository.CreateDocumentClient(repoReq)

	if err != nil {
		return nil, err
	}

	document_client := &dto.CreateDocumentClientRes{
		Id:         repoRes.Id,
		ClientId:   repoReq.ClientID,
		DocumentId: repoReq.DocumentID,
		Action:     repoRes.Action,
	}

	return document_client, err

}
