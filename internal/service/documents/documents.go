package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jackc/pgx/v5/pgtype"
	"mime/multipart"
	cnf "optitech/internal/config"
	drdto "optitech/internal/dto/directory_tree"
	dto "optitech/internal/dto/document"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"
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

	institutionName, err := s.documentRepository.GetInstitutionByDocumentId(int64(req.DirectoryId))

	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s%s/%s", cnf.Env.DigitalOceanFilesEndpoint, institutionName.Institution.InstitutionName, req.File.Filename)
	exists, err := s.documentRepository.GetEndpointExists(endpoint)

	if exists {
		return nil, err
	}

	repoReq := &sq.CreateDocumentParams{
		DirectoryID: req.DirectoryId,
		FormatID:    pgtype.Int4{Int32: req.FormatId, Valid: false},
		Name:        req.File.Filename,
		Status:      sq.Status(req.Status),
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	fileRute, err := UploadDocument(req.File, req.File.Filename, institutionName.Institution.InstitutionName)

	if err != nil {
		return nil, err
	}
	if req.FormatId > 0 {
		repoReq.FormatID = pgtype.Int4{Int32: req.FormatId, Valid: true}
	}
	repoReq.FileRute = fileRute
	repoRes, err := s.documentRepository.CreateDocument(repoReq)

	if err != nil {
		return nil, err
	}

	return repoRes, err
}

func UploadDocument(fileHeader *multipart.FileHeader, name string, institutionName string) (string, error) {

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(cnf.Env.DigitalOceanKey, cnf.Env.DigitalOceanSecret, ""),
		Endpoint:         aws.String(cnf.Env.DigitalOceanEndpoint),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String(cnf.Env.DigitalOceanRegion),
	}

	sess, err := session.NewSession(s3Config)

	if err != nil {
		return "", err
	}
	uploader := s3manager.NewUploader(sess)

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cnf.Env.DigitalOceanBucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", institutionName, name)),
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	defer file.Close()
	return aws.StringValue(&result.Location), nil
}

func (s *serviceDocument) DownloadDocumentById(req dto.GetDocumentReq) (string, error) {

	exist, err := s.documentRepository.ExistsDocuments(req.Id)
	if err != nil {
		return "", err
	}

	if exist {
		return "", fmt.Errorf("the document does not exist")
	}

	document, err := s.documentRepository.DownloadDocumentById(req.Id)
	if err != nil {
		return "", err
	}

	return document, err
}

func (s *serviceDocument) DeleteDocument(req dto.GetDocumentReq) (bool, error) {
	repoReq := &sq.DeleteDocumentByIdParams{
		DocumentID: req.Id,
		DeletedAt:  pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if err := s.documentRepository.DeleteDocument(repoReq); err != nil {
		return false, pgtype.ErrScanTargetTypeChanged
	}

	return true, nil
}

func (s *serviceDocument) UpdateDocument(req *dto.UpdateDocumentReq) (bool, error) {

	institutionName, err := s.documentRepository.GetInstitutionByDocumentId(int64(req.DirectoryID))

	if err != nil {
		return false, err
	}

	repoReq := &sq.UpdateDocumentNameByIdParams{
		DocumentID: req.Id,
		Name:       req.Name,
		FileRute:   fmt.Sprintf("%s%s/%s", cnf.Env.DigitalOceanFilesEndpoint, institutionName.Institution.InstitutionName, req.Name),
		UpdatedAt:  pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	repoRes, err := s.documentRepository.GetDocument(req.Id)

	if err != nil {
		return false, err
	}

	RenameDocument(repoRes.Name, req.Name)

	if err := s.documentRepository.UpdateDocument(repoReq); err != nil {
		return false, nil
	}
	return true, nil
}

func RenameDocument(oldName, newName string) error {
	s3Config := cnf.GetS3Config()

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return err
	}

	svc := s3.New(sess)

	_, err = svc.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(cnf.Env.DigitalOceanBucket),
		CopySource: aws.String(cnf.Env.DigitalOceanBucket + "/" + oldName),
		Key:        aws.String(newName),
	})
	if err != nil {
		return err
	}

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(cnf.Env.DigitalOceanBucket),
		Key:    aws.String(oldName),
	})
	if err != nil {
		return err
	}

	return nil
}
