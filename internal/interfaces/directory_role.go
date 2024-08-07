package interfaces

import (
	dto "optitech/internal/dto/directory_role"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IDirectoryRoleService interface {
	Create(req *dto.CreateDirectoryRoleReq) (*dto.CreateDirectoryRoleRes, error)
	Get(req dto.GetDirectoryRoleReq) (*dto.GetDirectoryRoleRes, error)
	List() (*[]dto.GetDirectoryRoleRes, error)
	Update(req *dto.UpdateDirectoryRoleReq) (bool, error)
	Delete(req dto.GetDirectoryRoleReq) (bool, error)
}

type IDirectoryRoleRepository interface {
	CreateDirectoryRole(arg *models.CreateDirectoryRoleParams) (*dto.CreateDirectoryRoleRes, error)
	GetDirectoryRole(userID int64) (*dto.GetDirectoryRoleRes, error)
	ListDirectoryRole() (*[]dto.GetDirectoryRoleRes, error)
	UpdateDirectoryRole(arg *models.UpdateDirectoryRoleParams) error
	DeleteDirectoryRole(arg *models.DeleteDirectoryRoleByIdParams) error
}

type IDirectoryRoleHandler interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
