package interfaces

import (
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IDirectoryRoleService interface {
	Create(req *[]models.CreateDirectoryRoleParams) error
}

type IDirectoryRoleRepository interface {
	CreateDirectoryRole(arg *[]models.CreateDirectoryRoleParams) error
}

type IDirectoryRoleHandler interface {
	Create(c *fiber.Ctx) error
}
