package handler

import (
	"github.com/gofiber/fiber/v2"
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/services"
	"optitech/internal/interfaces"
)

type handler_service struct {
	serviceService interfaces.IService
}

func NewHandlerService(r interfaces.IService) interfaces.IHandler {
	return &handler_service{
		serviceService: r,
	}
}

func (h *handler_service) Get(c *fiber.Ctx) error {
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

func (h *handler_service) List(c *fiber.Ctx) error {

	res, err := h.serviceService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handler_service) Create(c *fiber.Ctx) error {

	req := &cdto.CreateServiceReq{}
	err := c.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = dto.ValidateDTO(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	res, err := h.serviceService.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
func (h *handler_service) Update(c *fiber.Ctx) error {

	req := &cdto.CreateServiceReq{}
	err := c.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = dto.ValidateDTO(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	res, err := h.serviceService.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handler_service) Delete(c *fiber.Ctx) error {

	params := c.AllParams()
	req := &cdto.GetServiceReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.serviceService.Delete(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)

}
