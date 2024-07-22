package interfaces

import (
	dto "optitech/internal/dto/directory_tree"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IDirectoryService interface {
	Get(req dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error)
	GetRoute(req dto.GetDirectoryTreeReq) (*[]int64, *[]dto.GetDirectoryTreeRes, error)
	Create(req *dto.CreateDirectoryTreeReq) (*dto.CreateDirectoryTreeRes, error)
	List() (*[]dto.GetDirectoryTreeRes, error)
	ListByParent(req dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error)
	ListByChild(req dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error)
	Delete(req dto.GetDirectoryTreeReq) (bool, error)
	Update(req *dto.UpdateDirectoryTreeReq) (bool, error)
}

type IDirectoryRepository interface {
	GetDirectory(req *models.GetDirectoryTreeParams) (*dto.GetDirectoryTreeRes, error)
	GetDirectoryParentInstitution(institutionId int32) (*dto.GetDirectoryTreeRes, error)
	CreateDirectory(arg *models.CreateDirectoryTreeParams) (*dto.CreateDirectoryTreeRes, error)
	ListDirectory() (*[]dto.GetDirectoryTreeRes, error)
	DeleteDirectory(arg *models.DeleteDirectoryTreeByIdParams) error
	ListDirectoryByParent(parentId int64, institutionId int32) ([]*dto.GetDirectoryTreeRes, error)
	ListDirectoryHierarchy(childId int64, institutionId int32) (*[]dto.GetDirectoryTreeRes, error)
	UpdateDirectoryTree(arg *models.UpdateDirectoryTreeByIdParams) error
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
