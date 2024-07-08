package handler

import (
	dto "optitech/internal/dto"
	drdto "optitech/internal/dto/directory_role"
	"optitech/internal/interfaces"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type handlerDirectoryRole struct {
	directoryRoleService interfaces.IDirectoryRoleService
}

func NewHandlerDirectoryRole(r interfaces.IDirectoryRoleService) interfaces.IDirectoryRoleHandler {
	return &handlerDirectoryRole{
		directoryRoleService: r,
	}
}

func (h *handlerDirectoryRole) Create(c *fiber.Ctx) error {
	req := &drdto.CreateDirectoryRoleReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.directoryRoleService.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDirectoryRole) Get(c *fiber.Ctx) error {
	params := c.AllParams()

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id parameter")
	}

	req := &drdto.GetDirectoryRoleReq{
		UserId: id,
	}

	res, err := h.directoryRoleService.Get(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerDirectoryRole) List(c *fiber.Ctx) error {
	res, err := h.directoryRoleService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
