package interfaces

import (
	dto "optitech/internal/dto/directory_tree"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IDirectoryService interface {
	Get(req dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error)
	Create(req *dto.CreateDirectoryTreeReq) (*dto.CreateDirectoryTreeRes, error)
}

type IDirectoryRepository interface {
	GetDirectory(directoryID int64) (*dto.GetDirectoryTreeRes, error)
	CreateDirectory(arg *models.CreateDirectoryTreeParams) (*dto.CreateDirectoryTreeRes, error)
}

type IDirectoryHandler interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}
