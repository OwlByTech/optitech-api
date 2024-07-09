package service

import (
	"optitech/internal/config"
	drdto "optitech/internal/dto/directory_tree"
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

func (s *serviceDocument) ListByDirectory(req drdto.GetDirectoryTreeReq) (*[]dto.GetDocumentRes, error) {
	return s.documentRepository.ListDocumentByDirectory(int32(req.Id))
}

func (s *serviceDocument) Create(req *dto.CreateDocumentReq) (*dto.CreateDocumentRes, error) {

	repoReq := &sq.CreateDocumentParams{
		DirectoryID: req.DirectoryId,
		FormatID:    pgtype.Int4{Int32: req.FormatId, Valid: true},
		FileRute:    UploadDocument(),
		Name:        req.Name,
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

	do_conf := config.DigitalOcean

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(do_conf.DigitalOceanKey, do_conf.DigitalOceanSecret, ""), // Specifies your credentials.
		Endpoint:         aws.String(do_conf.DigitalOceanEndpoint),                                                  // Find your endpoint in the control panel, under Settings. Prepend "https://".
		S3ForcePathStyle: aws.Bool(false),                                                                           // // Configures to use subdomain/virtual calling format. Depending on your version, alternatively use o.UsePathStyle = false
		Region:           aws.String(do_conf.DigitalOceanRegion),                                                    // Must be "us-east-1" when creating new Spaces. Otherwise, use the region in your endpoint, such as "nyc3".
	}

	sess, err := session.NewSession(s3Config)

	if err != nil {
		return "Error"
	}

	//TODO: Change this for multipart

	filename := "/uploads/archivo.txt"
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		return "Error"
	}

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(do_conf.DigitalOceanBucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		return "Error"
	}
	return aws.StringValue(&result.Location)
}
