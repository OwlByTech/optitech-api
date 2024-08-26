package handler

import (
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/asesor"
	"optitech/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type handlerAsesor struct {
	asesorService interfaces.IAsesorService
}

func NewHandlerAsesor(r interfaces.IAsesorService) interfaces.IAsesorHandler {
	return &handlerAsesor{
		asesorService: r,
	}
}

func (h *handlerAsesor) Get(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetAsesorReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.asesorService.Get(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerAsesor) Create(c *fiber.Ctx) error {
	req := &cdto.CreateAsesorReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.asesorService.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerAsesor) Update(c *fiber.Ctx) error {
	params_id := c.AllParams()
	req_id := &cdto.GetAsesorReq{}
	if err := dto.ValidateParamsToDTO(params_id, req_id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req := &cdto.UpdateAsesorReq{
		AsesorID: req_id.Id,
	}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.asesorService.Update(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerAsesor) List(c *fiber.Ctx) error {
	res, err := h.asesorService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerAsesor) Delete(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetAsesorReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.asesorService.Delete(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
