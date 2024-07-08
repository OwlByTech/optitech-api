package service

import (
	cnf "optitech/internal/config"
	dto "optitech/internal/dto/document"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
		FileRute:    UploadDocument(),
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

func UploadDocument() string {

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(cnf.DigitalOcean.DigitalOceanKey, cnf.DigitalOcean.DigitalOceanSecret, ""),
		Endpoint:         aws.String(cnf.DigitalOcean.DigitalOceanEndpoint),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String(cnf.DigitalOcean.DigitalOceanRegion),
	}

	sess, err := session.NewSession(s3Config)

	if err != nil {
		return "Error"
	}

	//TODO: Change this for multipart

	filename := "internal/hola1.txt"
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		return "Error"
	}

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cnf.DigitalOcean.DigitalOceanBucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		return "Error"
	}
	return aws.StringValue(&result.Location)
}
