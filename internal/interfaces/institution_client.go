package interfaces

import (
	dto_client "optitech/internal/dto/client"
	dto "optitech/internal/dto/institution_client"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IInstitutionClientService interface {
	List(InstitutionID int32) (*[]dto_client.GetClientRes, error)
	Create(req *[]models.CreateInstitutionClientParams) error
	Exists(req *models.ExistsInstitutionClientParams) bool
	Recover(req *models.RecoverInstitutionClientParams) error
	Update(req dto.UpdateInstitutionClientReq) (bool, error)
	DeleteById(req *dto.GetInstitutionClientReq) error
	DeleteByInstitution(InstitutionID int32) error
	GetByInstitutionId(req dto.GetInstitutionClientReq) (*dto.GetInstitutionClientRes, error)
}
type IInstitutionClientRepository interface {
	ListInstitutionClient(InstitutionID int32) (*[]dto_client.GetClientRes, error)
	CreateInstitutionClient(arg *[]models.CreateInstitutionClientParams) error
	ExistsInstitutionClient(arg *models.ExistsInstitutionClientParams) bool
	RecoverInstitutionClient(arg *models.RecoverInstitutionClientParams) error
	DeleteInstitutionClientById(arg *models.DeleteInstitutionByClientParams) error
	DeleteInstitutionClientByInstitution(arg *models.DeleteInstitutionClientParams) error
	GetClientByInstitutionId(institutionID int32) (*dto.GetInstitutionClientRes, error)
}
type IInstitutionClientHandler interface {
	Update(c *fiber.Ctx) error
	GetClient(c *fiber.Ctx) error
}
