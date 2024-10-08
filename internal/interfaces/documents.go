package interfaces

import (
	dto "optitech/internal/dto/directory_tree"
	d "optitech/internal/dto/document"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IDocumentService interface {
	Get(req d.GetDocumentReq) (*d.GetDocumentRes, error)
	Create(arg *d.CreateDocumentByteReq) (*d.CreateDocumentRes, error)
	CreateVersion(arg *d.CreateDocumentVersionByteReq) (bool, error)
	ListByDirectory(req dto.GetDirectoryTreeReq) (*[]d.GetDocumentRes, error)
	DeleteDocument(req d.GetDocumentReq) (bool, error)
	DownloadDocumentById(req d.GetDocumentReq) (*string, error)
	UpdateDocument(req *d.UpdateDocumentReq) (bool, error)
	UpdateStatusById(req *d.UpdateDocumentStatusByIdReq) error
}

type IDocumentRepository interface {
	GetDocument(documentID int64) (*d.GetDocumentRes, error)
	CreateDocument(arg *models.CreateDocumentParams) (*d.CreateDocumentRes, error)
	ListDocumentByDirectory(directoryID int64) (*[]d.GetDocumentRes, error)
	DeleteDocument(arg *models.DeleteDocumentByIdParams) error
	DownloadDocumentById(documentID int64) (*d.GetDocumentDownloadRes, error)
	UpdateDocument(arg *models.UpdateDocumentNameByIdParams) error
	UpdateDocumentById(arg *models.UpdateDocumentByIdParams) error
	UpdateDocumentStatusById(*models.UpdateDocumentStatusByIdParams) error
	ExistsDocuments(documentID int64) (bool, error)
	GetEndpointExists(fileRute string) (bool, error)
	GetInstitutionByDocumentId(directoryId int64) (models.GetInstitutionNameByDirectoryIdRow, error)
}

type IDocumentHandler interface {
	Get(f *fiber.Ctx) error
	CreateDocument(f *fiber.Ctx) error
	DeleteDocument(f *fiber.Ctx) error
	DownloadDocumentById(f *fiber.Ctx) error
	UpdateDocument(c *fiber.Ctx) error
	UpdateStatusById(c *fiber.Ctx) error
	CreateVersion(f *fiber.Ctx) error
}
