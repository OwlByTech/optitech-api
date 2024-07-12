package interfaces

import (
	d "optitech/internal/dto/document"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IDocumentService interface {
	Get(req d.GetDocumentReq) (*d.GetDocumentRes, error)
	Create(arg *d.CreateDocumentReq) (*d.CreateDocumentRes, error)
}

type IDocumentRepository interface {
	GetDocument(documentID int64) (*d.GetDocumentRes, error)
	CreateDocument(arg *models.CreateDocumentParams) (*d.CreateDocumentRes, error)
}

type IDocumentHandler interface {
	Get(f *fiber.Ctx) error
	CreateDocument(f *fiber.Ctx) error
}
