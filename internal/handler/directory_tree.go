package handler

import (
	dto "optitech/internal/dto"
	ddto "optitech/internal/dto/directory_tree"
	"optitech/internal/interfaces"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
	log.Info("InstitutionID", institutionId)

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
	params := c.AllParams()
	req := &ddto.GetDirectoryTreeReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.directoryTreeService.Delete(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
