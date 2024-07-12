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

func (r *repositoryDocument) ListDocumentByDirectory(directoryID int32) (*[]dto.GetDocumentRes, error) {
	ctx := context.Background()

	repoRes, err := r.documentRepository.ListDocumentsByDirectory(ctx, directoryID)

	if err != nil {
		return nil, err
	}
	documents := make([]dto.GetDocumentRes, len(repoRes))
	for i, inst := range repoRes {
		documents[i] = dto.GetDocumentRes{
			Id:          inst.DocumentID,
			Name:        inst.Name,
			DirectoryId: inst.DirectoryID,
			FormatId:    inst.FormatID.Int32,
			FileRute:    inst.FileRute,
			Status:      string(inst.Status),
		}
	}
	return &documents, nil
}

func (r *repositoryDocument) CreateDocument(arg *sq.CreateDocumentParams) (*dto.CreateDocumentRes, error) {
	ctx := context.Background()

	res, err := r.documentRepository.CreateDocument(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateDocumentRes{
		DirectoryId: res.DirectoryID,
		FormatId:    res.FormatID.Int32,
		FileRute:    res.FileRute,
		Status:      string(res.Status),
	}, nil

}
