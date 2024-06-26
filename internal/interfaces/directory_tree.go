package interfaces

import (
	dto "optitech/internal/dto/directory_tree"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IDirectoryService interface {
	Get(req dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error)
	Create(req *dto.CreateDirectoryTreeReq) (*dto.CreateDirectoryTreeRes, error)
	List() (*[]dto.GetDirectoryTreeRes, error)
	Delete(req dto.GetDirectoryTreeReq) (bool, error)
}

type IDirectoryRepository interface {
	GetDirectory(directoryID int64) (*dto.GetDirectoryTreeRes, error)
	CreateDirectory(arg *models.CreateDirectoryTreeParams) (*dto.CreateDirectoryTreeRes, error)
	ListDirectory() (*[]dto.GetDirectoryTreeRes, error)
	DeleteDirectory(arg *models.DeleteDirectoryTreeByIdParams) error
}

type IDirectoryHandler interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
