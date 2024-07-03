package repository

import (
	"context"
	dto "optitech/internal/dto/document"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryDocument struct {
	documentRepository *sq.Queries
}

func NewRepositoryDocument(q *sq.Queries) interfaces.IDocumentRepository {
	return &repositoryDocument{
		documentRepository: q,
	}
}

func (r *repositoryDocument) GetDocument(documentID int64) (*dto.GetDocumentRes, error) {
	ctx := context.Background()

	repoRes, err := r.documentRepository.GetDocument(ctx, (documentID))

	if err != nil {
		return nil, err
	}

	return &dto.GetDocumentRes{
		Id:          repoRes.DocumentID,
		DirectoryId: repoRes.DirectoryID,
		FormatId:    repoRes.FormatID.Int32,
		FileRute:    repoRes.FileRute,
		Status:      string(repoRes.Status),
	}, nil
}
