package interfaces

import (
	cdto "optitech/internal/dto/client"
	dto "optitech/internal/dto/institution"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IInstitutionService interface {
	GetByClient(req cdto.GetClientReq) (*dto.GetInstitutionRes, error)
	Get(req dto.GetInstitutionReq) (*dto.GetInstitutionRes, error)
	Create(req *dto.CreateInstitutionReq) (*dto.CreateInstitutionRes, error)
	Update(req *dto.UpdateInstitutionReq) (bool, error)
	UpdateAsesor(req *dto.UpdateAsesorInstitutionReq) (bool, error)
	UpdateLogo(req *dto.UpdateLogoReq) (bool, error)
	List() (*[]dto.GetInstitutionRes, error)
	Delete(req dto.GetInstitutionReq) (bool, error)
}
type IInstitutionRepository interface {
	GetInstitutionByClient(ClientID int32) (int32, error)
	GetInstitution(institutionID int32) (*dto.GetInstitutionRes, error)
	CreateInstitution(arg *models.CreateInstitutionParams) (int32, error)
	UpdateInstitution(arg *models.UpdateInstitutionParams) error
	ListInstitutions() (*[]dto.GetInstitutionRes, error)
	DeleteInstitution(arg *models.DeleteInstitutionParams) error
	UpdateAsesorInstitution(arg *models.UpdateAsesorInstitutionParams) error
	UpdateLogoInstitution(arg *models.UpdateLogoInstitutionParams) error
}

type IInstitutionHandler interface {
	GetByClient(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	UpdateLogo(c *fiber.Ctx) error
	UpdateAsesor(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
