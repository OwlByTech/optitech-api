package interfaces

import (
	dto "optitech/internal/dto/institution"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IInstitutionService interface {
	Get(req dto.GetInstitutionReq) (*dto.GetInstitutionRes, error)
	Create(req *dto.CreateInstitutionReq) (*dto.CreateInstitutionRes, error)
	Update(req *dto.UpdateInstitutionReq) (bool, error)
	List() ([]*dto.GetInstitutionRes, error)
	Delete(req dto.GetInstitutionReq) (bool, error)
}
type IInstitutionRepository interface {
	GetInstitution(institutionID int64) (*dto.GetInstitutionRes, error)
	CreateInstitution(arg *models.CreateInstitutionParams) (*dto.CreateInstitutionRes, error)
	UpdateInstitution(arg *models.UpdateInstitutionParams) error
	ListInstitutions() ([]*dto.GetInstitutionRes, error)
	DeleteInstitution(arg *models.DeleteInstitutionParams) error
}

type IHandler interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
