package handler

import (
	dto "optitech/internal/dto"
	fdto "optitech/internal/dto/document"
	"optitech/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handlerDocument struct {
	documentService interfaces.IDocumentService
}

func NewHandlerDocument(d interfaces.IDocumentService) interfaces.IDocumentHandler {
	return &handlerDocument{
		documentService: d,
	}
}

func (h *handlerDocument) Get(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &fdto.GetDocumentReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.documentService.Get(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
