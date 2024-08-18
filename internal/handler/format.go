package handler

import (
	"encoding/json"
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

	data := f.Locals("asesorId")
	asesorID, ok_asesor := data.(int32)

	if !ok_asesor {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}
	req := &fdto.CreateFormatReq{}
	body := f.FormValue("data")

	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	file, err := f.FormFile("file")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req.FormatFile = file
	req.AsesorId = asesorID
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

func (h *handlerFormat) ListById(c *fiber.Ctx) error {

	data := c.Locals("asesorId")
	asesorID, ok_asesor := data.(int32)

	if !ok_asesor {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}
	req := &fdto.ListFormatsReq{
		AsesorId: asesorID,
	}
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}
	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	res, err := h.formatService.ListById(req)
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

	params := c.AllParams()
	req_id := fdto.GetFormatReq{}
	if err := dto.ValidateParamsToDTO(params, req_id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	formatId, err := strconv.Atoi(formatIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid format ID")
	}

	formatId32 := int32(formatId)

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
