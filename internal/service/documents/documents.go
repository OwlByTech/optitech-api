package service

import (
	"fmt"
	drdto "optitech/internal/dto/directory_tree"
	dto "optitech/internal/dto/document"
	"optitech/internal/interfaces"
	digitalOcean "optitech/internal/service/digital_ocean"
	sq "optitech/internal/sqlc"
	"optitech/internal/tools"
	"path/filepath"
	"time"

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

func folderPath(institutionId int32, asesorId int32) string {
	if institutionId > 0 {
		return tools.FolderTypePath(tools.InstitutionFolderType, institutionId)
	}

	return tools.FolderTypePath(tools.AsesorFolderType, asesorId)
}

func (s *serviceDocument) Get(req dto.GetDocumentReq) (*dto.GetDocumentRes, error) {
	return s.documentRepository.GetDocument(req.Id)
}

func (s *serviceDocument) ListByDirectory(req drdto.GetDirectoryTreeReq) (*[]dto.GetDocumentRes, error) {
	return s.documentRepository.ListDocumentByDirectory(req.Id)
}

func (s *serviceDocument) Create(req *dto.CreateDocumentByteReq) (*dto.CreateDocumentRes, error) {
	repoReq := &sq.CreateDocumentParams{
		DirectoryID: req.DirectoryId,
		Name:        req.Filename,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if req.Status != "" {
		repoReq.Status = sq.NullStatus{Status: sq.Status(req.Status), Valid: true}
	}

	if req.FormatId > 0 {
		repoReq.FormatID = pgtype.Int4{Int32: req.FormatId, Valid: true}
	}

	folder := folderPath(req.InstitutionId, req.AsesorId)
	filename := tools.NormalizeFilename(req.Filename)
	filePath := filepath.Join(folder, filename)

	if err := digitalOcean.UploadDocument(*req.File, filePath); err != nil {
		return nil, err
	}

	repoReq.FileRute = filePath
	repoRes, err := s.documentRepository.CreateDocument(repoReq)
	if err != nil {
		return nil, err
	}

	return repoRes, err
}

func (s *serviceDocument) DownloadDocumentById(req dto.GetDocumentReq) (*string, error) {
	doc, err := s.documentRepository.DownloadDocumentById(req.Id)
	if err != nil {
		return nil, err
	}

	if doc.AsesorId != req.AsesorId && doc.InstitutionId != req.InstitutionId {
		return nil, fmt.Errorf("the document does not exist")
	}

	url, err := digitalOcean.DownloadDocumentWithFilename(doc.FileRute, doc.Filename)
	if err != nil {
		return nil, err
	}

	return url, nil
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
