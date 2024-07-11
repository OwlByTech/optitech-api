package interfaces

import (
	dc "optitech/internal/dto/document_client"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IDocumentClientService interface {
	GetDocumentClient(req dc.GetDocumentClientReq) (*dc.GetDocumentClientRes, error)
	CreateDocumentClient(arg *dc.CreateDocumentClientReq) (*dc.CreateDocumentClientRes, error)
}

type IDocumentClientRepository interface {
	GetDocumentClient(DocumentClientID int32) (*dc.GetDocumentClientRes, error)
	CreateDocumentClient(arg *models.CreateDocumentClientParams) (*dc.CreateDocumentClientRes, error)
}

type IDocumentClientHandler interface {
	GetDocumentClient(f *fiber.Ctx) error
	CreateDocumentClient(f *fiber.Ctx) error
}
