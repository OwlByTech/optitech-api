package handler

import (
	"github.com/gofiber/fiber/v2"
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/services"
	"optitech/internal/interfaces"
)

type handlerService struct {
	serviceService interfaces.IService
}

func NewHandlerService(r interfaces.IService) interfaces.IServiceHandler {
	return &handlerService{
		serviceService: r,
	}
}

func (h *handlerService) Get(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetServiceReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.serviceService.Get(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerService) List(c *fiber.Ctx) error {

	res, err := h.serviceService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
