package interfaces

import (
	dto "optitech/internal/dto/directory_tree"
	d "optitech/internal/dto/document"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IDocumentService interface {
	Get(req d.GetDocumentReq) (*d.GetDocumentRes, error)
	Create(arg *d.CreateDocumentReq) (*d.CreateDocumentRes, error)
	ListByDirectory(req dto.GetDirectoryTreeReq) (*[]d.GetDocumentRes, error)
	DeleteDocument(req d.GetDocumentReq) (bool, error)
}

type IDocumentRepository interface {
	GetDocument(documentID int64) (*d.GetDocumentRes, error)
	CreateDocument(arg *models.CreateDocumentParams) (*d.CreateDocumentRes, error)
	ListDocumentByDirectory(directoryID int32) (*[]d.GetDocumentRes, error)
	DeleteDocument(arg *models.DeleteDocumentByIdParams) error
}

type IDocumentHandler interface {
	Get(f *fiber.Ctx) error
	CreateDocument(f *fiber.Ctx) error
	DeleteDocument(f *fiber.Ctx) error
}
