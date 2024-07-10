package service

import (
	dto "optitech/internal/dto/document_client"
	"optitech/internal/interfaces"
)

type serviceDocumentClient struct {
	documentClientRepository interfaces.IDocumentClientRepository
}

func NewServiceDocumentClient(f interfaces.IDocumentClientRepository) interfaces.IDocumentClientService {
	return &serviceDocumentClient{
		documentClientRepository: f,
	}
}

func (s *serviceDocumentClient) Get(req dto.GetDocumentClientReq) (*dto.GetDocumentClientRes, error) {
	return s.documentClientRepository.GetDocumentClient(req.Id)
}
