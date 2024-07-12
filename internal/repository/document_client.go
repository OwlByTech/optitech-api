package repository

import (
	"context"
	dto "optitech/internal/dto/document_client"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryDocumentClient struct {
	documentClientRepository *sq.Queries
}

func NewRepositoryDocumentClient(q *sq.Queries) interfaces.IDocumentClientRepository {
	return &repositoryDocumentClient{
		documentClientRepository: q,
	}
}

func (r *repositoryDocumentClient) GetDocumentClient(documentClientID int64) (*dto.GetDocumentClientRes, error) {
	ctx := context.Background()

	repoRes, err := r.documentClientRepository.GetDocumentClient(ctx, (documentClientID))

	if err != nil {
		return nil, err
	}

	return &dto.GetDocumentClientRes{
		Id:         documentClientID,
		ClientId:   repoRes.ClientID,
		DocumentId: repoRes.DocumentID,
		Action:     string(repoRes.Action),
	}, nil
}

func (r *repositoryDocumentClient) CreateDocumentClient(arg *sq.CreateDocumentClientParams) (*dto.CreateDocumentClientRes, error) {
	ctx := context.Background()

	res, err := r.documentClientRepository.CreateDocumentClient(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateDocumentClientRes{
		ClientId:   res.ClientID,
		DocumentId: res.DocumentID,
		Action:     string(res.Action),
	}, nil

}
