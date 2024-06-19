package handler

import (
	"github.com/gofiber/fiber/v2"
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/client"
	"optitech/internal/interfaces"
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
	req := &cdto.CreateClientReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.clientService.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerClient) Update(c *fiber.Ctx) error {
	params_id := c.AllParams()
	req_id := &cdto.GetClientReq{}
	if err := dto.ValidateParamsToDTO(params_id, req_id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req := &cdto.UpdateClientReq{
		ClientID: req_id.Id,
	}

	res, err := h.clientService.Update(req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerClient) List(c *fiber.Ctx) error {
	res, err := h.clientService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerClient) Delete(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetClientReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.clientService.Delete(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerClient) Login(c *fiber.Ctx) error {
	req := &cdto.LoginClientReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.clientService.Login(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)

}

func (h *handlerClient) ResetPassword(c *fiber.Ctx) error {
	req := &cdto.ResetPasswordReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.clientService.ResetPassword(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
