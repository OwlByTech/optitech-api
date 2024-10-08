package interfaces

import (
	dto "optitech/internal/dto/directory_tree"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IDirectoryService interface {
	Get(req *dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error)
	GetRoute(req *dto.GetDirectoryTreeReq) (*[]int64, *[]dto.GetDirectoryTreeRes, error)
	Create(req *dto.CreateDirectoryTreeReq) (*dto.CreateDirectoryTreeRes, error)
	List() (*[]dto.GetDirectoryTreeRes, error)
	ListByParent(req *dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error)
	ListByChild(req *dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error)
	Delete(req *dto.GetDirectoryTreeReq) (bool, error)
	Update(req *dto.UpdateDirectoryTreeReq) (bool, error)
	GetIdByParent(req *dto.GetDirectoryTreeReq) (*int64, error)
}

type IDirectoryRepository interface {
	GetDirectory(req *dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error)
	GetDirectoryParent(req *dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error)
	CreateDirectory(arg *models.CreateDirectoryTreeParams) (*dto.CreateDirectoryTreeRes, error)
	ListDirectory() (*[]dto.GetDirectoryTreeRes, error)
	DeleteDirectory(arg *models.DeleteDirectoryTreeByIdParams) error
	ListDirectoryByParent(*dto.GetDirectoryTreeReq) ([]*dto.GetDirectoryTreeRes, error)
	ListDirectoryHierarchy(*dto.GetDirectoryTreeReq) (*[]dto.GetDirectoryTreeRes, error)
	UpdateDirectoryTree(arg *models.UpdateDirectoryTreeByIdParams) error
	GetDirectoryIdByParent(req *dto.GetDirectoryTreeReq) (*int64, error)
}

type IDirectoryHandler interface {
	Get(c *fiber.Ctx) error
	GetRoute(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	ListByParent(c *fiber.Ctx) error
	ListByChild(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}
