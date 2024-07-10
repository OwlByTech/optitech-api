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

func (r *repositoryDocumentClient) GetDocumentClient(documentClientID int32) (*dto.GetDocumentClientRes, error) {
	ctx := context.Background()

	repoRes, err := r.documentClientRepository.GetDocumentClient(ctx, (documentClientID))

	if err != nil {
		return nil, err
	}

	return &dto.GetDocumentClientRes{
		//TODO: ADD ID
		ClientId:   repoRes.ClientID,
		DocumentId: repoRes.DocumentID,
		Action:     string(repoRes.Action),
	}, nil
}
