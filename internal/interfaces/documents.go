package interfaces

import (
	d "optitech/internal/dto/document"

	"github.com/gofiber/fiber/v2"
)

type IDocumentService interface {
	Get(req d.GetDocumentReq) (*d.GetDocumentRes, error)
}

type IDocumentRepository interface {
	GetDocument(documentID int64) (*d.GetDocumentRes, error)
}

type IDocumentHandler interface {
	Get(f *fiber.Ctx) error
}
