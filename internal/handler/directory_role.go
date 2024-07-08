package handler

import (
	dto "optitech/internal/dto"
	drdto "optitech/internal/dto/directory_role"
	"optitech/internal/interfaces"

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
