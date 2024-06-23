package interfaces

import (
	dto "optitech/internal/dto/asesor"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IAsesorService interface {
	Get(req dto.GetAsesorReq) (*dto.GetAsesorRes, error)
	Create(req *dto.CreateAsesorReq) (*dto.CreateAsesorRes, error)
	Update(req *dto.UpdateAsesorReq) (bool, error)
	List() (*[]dto.GetAsesorRes, error)
	Delete(req dto.GetAsesorReq) (bool, error)
}
type IAsesorRepository interface {
	GetAsesor(asesorID int32) (*dto.GetAsesorRes, error)
	CreateAsesor(arg *models.CreateAsesorParams) (*dto.CreateAsesorRes, error)
	UpdateAsesor(arg *models.UpdateAsesorByIdParams) error
	ListAsesor() (*[]dto.GetAsesorRes, error)
	DeleteAsesor(arg *models.DeleteAsesorByIdParams) error
}

type IAsesorHandler interface {
	GetSecure(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
