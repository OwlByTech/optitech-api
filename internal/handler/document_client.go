package handler

import (
	dto "optitech/internal/dto"
	ddto "optitech/internal/dto/document_client"
	"optitech/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handlerDocumentClient struct {
	documentClientService interfaces.IDocumentClientService
}

func NewhandlerDocumentClient(f interfaces.IDocumentClientService) interfaces.IDocumentClientHandler {
	return &handlerDocumentClient{
		documentClientService: f,
	}
}

func (h *handlerDocumentClient) Get(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &ddto.GetDocumentClientReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.documentClientService.Get(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
