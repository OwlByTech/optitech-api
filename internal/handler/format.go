package handler

import (
	dto "optitech/internal/dto"
	fdto "optitech/internal/dto/format"
	"optitech/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handlerFormat struct {
	formatService interfaces.IFormatService
}

func NewHandlerFormat(f interfaces.IFormatService) interfaces.IFormatHandler {
	return &handlerFormat{
		formatService: f,
	}
}

func (h *handlerFormat) Get(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &fdto.GetFormatReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.formatService.Get(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerFormat) Create(f *fiber.Ctx) error {
	req := fdto.CreateFormatReq{}

	if err := f.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Entrada inválida: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.formatService.Create(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return f.JSON(res)
}

func (h *handlerFormat) List(c *fiber.Ctx) error {
	res, err := h.formatService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerFormat) Delete(c *fiber.Ctx) error {
	data := c.Locals("formatId")
	formatId, ok := data.(int32)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}
	params := c.AllParams()
	req := &fdto.GetFormatReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	req.Id = formatId
	res, err := h.formatService.Delete(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerFormat) Update(c *fiber.Ctx) error {
	data := c.Locals("formatId")
	formatId, ok := data.(int32)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}
	params_id := c.AllParams()
	req_id := &fdto.GetFormatReq{}

	if err := dto.ValidateParamsToDTO(params_id, req_id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req := &fdto.UpdateFormatReq{
		FormatID: formatId,
	}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unvalid entry: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	success, err := h.formatService.Update(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(success)
}
