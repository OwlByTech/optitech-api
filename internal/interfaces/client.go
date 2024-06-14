package interfaces

import (
	dto "optitech/internal/dto/client"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IClientService interface {
	Get(req dto.GetClientReq) (*dto.GetClientRes, error)
	Create(req *dto.CreateClientReq) (*dto.CreateClientRes, error)
	Update(req *dto.UpdateClientReq) (bool, error)
	List() (*[]dto.GetClientRes, error)
	Delete(req dto.GetClientReq) (bool, error)
}
type IClientRepository interface {
	GetClient(institutionID int64) (*dto.GetClientRes, error)
	CreateClient(arg *models.CreateInstitutionParams) (*dto.CreateClientRes, error)
	UpdateClient(arg *models.UpdateClientByIdParams) error
	ListClient() (*[]dto.GetClientRes, error)
	DeleteClient(arg *models.DeleteClientByIdParams) error
}

type IClientHandler interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
