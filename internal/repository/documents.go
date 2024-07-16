package repository

import (
	"context"
	cnf "optitech/internal/config"
	dto "optitech/internal/dto/document"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
		Name:        repoRes.Name,
		Id:          repoRes.DocumentID,
		DirectoryId: repoRes.DirectoryID,
		FormatId:    repoRes.FormatID.Int32,
		FileRute:    repoRes.FileRute,
		Status:      string(repoRes.Status),
	}, nil
}

func DownloadDocument(name string) (string, error) {

	s3Config := cnf.GetS3Config()

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(cnf.Env.DigitalOceanBucket),
		Key:    aws.String(name),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}

func (r *repositoryDocument) DownloadDocumentById(documentID int64) (string, error) {
	ctx := context.Background()

	repoRes, err := r.documentRepository.GetDocument(ctx, (documentID))

	if err != nil {
		return "", err
	}

	return DownloadDocument(repoRes.Name)
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
func (r *repositoryDocument) DeleteDocument(arg *sq.DeleteDocumentByIdParams) error {
	ctx := context.Background()
	return r.documentRepository.DeleteDocumentById(ctx, *arg)
}

func (r *repositoryDocument) UpdateDocument(arg *sq.UpdateDocumentNameByIdParams) error {
	ctx := context.Background()
	return r.documentRepository.UpdateDocumentNameById(ctx, *arg)
}

func (r *repositoryDocument) ExistsDocuments(documentID int64) bool {
	ctx := context.Background()
	_, err := r.documentRepository.ExistsDocument(ctx, (documentID))
	return err == nil
}
