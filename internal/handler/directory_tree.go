package handler

import (
	dto "optitech/internal/dto"
	ddto "optitech/internal/dto/directory_tree"
	"optitech/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handlerDirectoryTree struct {
	directoryTreeService interfaces.IDirectoryService
}

func NewHnadlerDirectoryTree(r interfaces.IDirectoryService) interfaces.IDirectoryHandler {
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

func (h *handlerDirectoryTree) Create(c *fiber.Ctx) error {
	req := &ddto.CreateDirectoryTreeReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
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
