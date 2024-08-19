package handler

import (
	"encoding/json"
	dto "optitech/internal/dto"
	ddto "optitech/internal/dto/document"
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
	req := &ddto.GetDocumentReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.documentService.Get(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDocument) CreateDocument(c *fiber.Ctx) error {

	data := c.Locals("institutionId")
	institutionId, ok := data.(int32)

	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}
	req := &ddto.CreateDocumentReq{
		InstitutionId: institutionId,
	}

	body := c.FormValue("data")
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	file, err := c.FormFile("file")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	req.File = file

	res, err := h.documentService.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDocument) DeleteDocument(c *fiber.Ctx) error {
	data := c.Locals("institutionId")
	institutionId, ok := data.(int32)
	data_asesor := c.Locals("asesorId")
	asesorID, ok_asesor := data_asesor.(int32)

	if !ok && !ok_asesor {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}
	params := c.AllParams()
	req := &ddto.GetDocumentReq{
		InstitutionId: institutionId,
		AsesorId:      asesorID,
	}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.documentService.DeleteDocument(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDocument) DownloadDocumentById(c *fiber.Ctx) error {
	data := c.Locals("institutionId")
	institutionId, ok := data.(int32)
	data_asesor := c.Locals("asesorId")
	asesorID, ok_asesor := data_asesor.(int32)

	if !ok && !ok_asesor {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}
	params := c.AllParams()
	req := &ddto.GetDocumentReq{
		InstitutionId: institutionId,
		AsesorId:      asesorID,
	}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.documentService.DownloadDocumentById(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDocument) UpdateDocument(c *fiber.Ctx) error {

	req := &ddto.UpdateDocumentReq{}
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.documentService.UpdateDocument(req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)

}
