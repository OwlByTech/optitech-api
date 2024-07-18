package handler

import (
	"github.com/gofiber/fiber/v2"
	dto "optitech/internal/dto"
	ddto "optitech/internal/dto/directory_tree"
	"optitech/internal/interfaces"
	"strconv"
)

type handlerDirectoryTree struct {
	directoryTreeService interfaces.IDirectoryService
}

func NewHandlerDirectoryTree(r interfaces.IDirectoryService) interfaces.IDirectoryHandler {
	return &handlerDirectoryTree{
		directoryTreeService: r,
	}
}

func (h *handlerDirectoryTree) Get(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &ddto.GetDirectoryTreeReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.directoryTreeService.Get(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDirectoryTree) GetRoute(c *fiber.Ctx) error {
	data := c.Locals("institutionId")
	institutionId, ok := data.(int32)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}
	params := c.AllParams()
	req := &ddto.GetDirectoryTreeReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	req.InstitutionID = institutionId
	_, res, err := h.directoryTreeService.GetRoute(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDirectoryTree) Create(c *fiber.Ctx) error {
	req := &ddto.CreateDirectoryTreeReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Entrada inválida: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.directoryTreeService.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDirectoryTree) List(c *fiber.Ctx) error {
	res, err := h.directoryTreeService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDirectoryTree) ListByParent(c *fiber.Ctx) error {
	data := c.Locals("institutionId")
	institutionId, ok := data.(int32)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}

	params := c.AllParams()
	req := &ddto.GetDirectoryTreeReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	req.InstitutionID = institutionId
	res, err := h.directoryTreeService.ListByParent(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDirectoryTree) ListByChild(c *fiber.Ctx) error {
	data := c.Locals("institutionId")
	institutionId, ok := data.(int32)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}
	params := c.AllParams()
	req := &ddto.GetDirectoryTreeReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	req.InstitutionID = institutionId
	res, err := h.directoryTreeService.ListByChild(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDirectoryTree) Delete(c *fiber.Ctx) error {
	data := c.Locals("institutionId")
	institutionId, ok := data.(int32)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}
	params := c.AllParams()
	req := &ddto.GetDirectoryTreeReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	req.InstitutionID = institutionId
	res, err := h.directoryTreeService.Delete(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDirectoryTree) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID inválido")
	}

	req_id := &ddto.GetDirectoryTreeReq{
		Id: id,
	}

	if err := dto.ValidateDTO(req_id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req := &ddto.UpdateDirectoryTreeReq{
		DirectoryId: req_id.Id,
	}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unvalid entry: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	success, err := h.directoryTreeService.Update(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(success)
}
