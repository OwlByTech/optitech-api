package service

import (
	"fmt"
	"log"
	"mime/multipart"
	cnf "optitech/internal/config"
	drdto "optitech/internal/dto/directory_tree"
	dto "optitech/internal/dto/document"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"path/filepath"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jackc/pgx/v5/pgtype"
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

	repoReq := &sq.CreateDocumentParams{
		DirectoryID: req.DirectoryId,
		FormatID:    pgtype.Int4{Int32: req.FormatId, Valid: false},
		Name:        req.File.Filename,
		Status:      sq.Status(req.Status),
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	rute := fmt.Sprintf("%s%s", strconv.FormatInt(time.Now().UTC().UnixMicro(), 10), filepath.Ext(req.File.Filename))

	log.Print(rute)

	fileRute, err := UploadDocument(req.File, rute, institutionName.Institution.InstitutionName)

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

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cnf.Env.DigitalOceanBucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", institutionName, name)),
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	defer file.Close()
	return name, nil
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

	repoReq := &sq.UpdateDocumentNameByIdParams{
		DocumentID: req.Id,
		Name:       req.Name,
		UpdatedAt:  pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	/*
		repoRes, err := s.documentRepository.GetDocument(req.Id)

		if err != nil {
			return false, err
		}

		RenameDocument(repoRes.Name, fileName)

	*/

	if err := s.documentRepository.UpdateDocument(repoReq); err != nil {
		return false, nil
	}
	return true, nil
}

/*
func RenameDocument(oldName string, newName string) error {

	log.Print(oldName, newName)

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
*/
