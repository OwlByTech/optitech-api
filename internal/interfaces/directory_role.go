package interfaces

import (
	dto "optitech/internal/dto/directory_role"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IDirectoryRoleService interface {
	Create(req *dto.CreateDirectoryRoleReq) (*dto.CreateDirectoryRoleRes, error)
}

type IDirectoryRoleRepository interface {
	CreateDirectoryRole(arg *models.CreateDirectoryRoleParams) (*dto.CreateDirectoryRoleRes, error)
}

type IDirectoryRoleHandler interface {
	Create(c *fiber.Ctx) error
}
