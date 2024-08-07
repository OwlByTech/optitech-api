package handler

import (
	"log"
	dto "optitech/internal/dto"
	fdto "optitech/internal/dto/format"
	"optitech/internal/interfaces"
	"strconv"

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
	req := &fdto.CreateFormatReq{}

	if err := f.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Entrada inválida: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.formatService.Create(req)
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
	formatIdStr := c.Params("Id")
	if formatIdStr == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing format ID in URL")
	}

	formatId, err := strconv.Atoi(formatIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid format ID")
	}

	formatId32 := int32(formatId)


	params := c.AllParams()
	req := &fdto.GetFormatReq{}

	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req.Id = formatId32
	res, err := h.formatService.Delete(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerFormat) Update(c *fiber.Ctx) error {
	formatIdStr := c.Params("Id")
	if formatIdStr == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing format ID in URL")
	}

	formatId, err := strconv.Atoi(formatIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid format ID")
	}

	formatId32 := int32(formatId)

	log.Printf("formatId type: %T, value: %v\n", formatId32, formatId32)

	req := &fdto.UpdateFormatReq{
		FormatID: formatId32,
	}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid entry: "+err.Error())
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
