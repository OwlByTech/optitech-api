package handler

import (
	"optitech/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handlerDirectoryInstitution struct {
	directoryInstitutionService interfaces.IDirectoryInstitutionService
}

func NewHandlerDirectoryInstitution(d interfaces.IDirectoryInstitutionService) interfaces.IDirectoryInstitutionHandler {
	return &handlerDirectoryInstitution{
		directoryInstitutionService: d,
	}
}

func (h *handlerDirectoryInstitution) ListDirectoryInstitution(c *fiber.Ctx) error {

	res, err := h.directoryInstitutionService.ListDirectoryInstitution()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
