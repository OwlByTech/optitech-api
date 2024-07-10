package interfaces

import (
	dc "optitech/internal/dto/document_client"

	"github.com/gofiber/fiber/v2"
)

type IDocumentClientService interface {
	Get(req dc.GetDocumentClientReq) (*dc.GetDocumentClientRes, error)
}

type IDocumentClientRepository interface {
	GetDocumentClient(DocumentClientID int32) (*dc.GetDocumentClientRes, error)
}

type IDocumentClientHandler interface {
	Get(dc *fiber.Ctx) error
}
