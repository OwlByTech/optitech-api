package handler

import (
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/client"
	"optitech/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handlerClient struct {
	clientService interfaces.IClientService
}

func NewHandlerClient(r interfaces.IClientService) interfaces.IClientHandler {
	return &handlerClient{
		clientService: r,
	}
}

func (h *handlerClient) Get(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetClientReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.clientService.Get(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerClient) Create(c *fiber.Ctx) error {
	return nil
}

func (h *handlerClient) Update(c *fiber.Ctx) error {
	return nil
}

func (h *handlerClient) List(c *fiber.Ctx) error {
	return nil
}

func (h *handlerClient) Delete(c *fiber.Ctx) error {
	return nil
}
