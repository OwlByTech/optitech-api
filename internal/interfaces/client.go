package interfaces

import (
	dto "optitech/internal/dto/client"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IClientService interface {
	Get(req dto.GetClientReq) (*dto.GetClientRes, error)
	GetPhoto(req dto.GetClientReq) (string, error)
	Create(req *dto.CreateClientReq) (*dto.CreateClientRes, error)
	Update(req *dto.UpdateClientReq) (bool, error)
	UpdateStatus(req *dto.UpdateClientStatusReq) (bool, error)
	UpdatePhoto(req *dto.UpdateClientPhotoReq) (bool, error)
	List() (*[]dto.GetClientRes, error)
	Delete(req dto.GetClientReq) (bool, error)
	Login(req *dto.LoginClientReq) (*dto.LoginClientRes, error)
	ResetPassword(req dto.ResetPasswordReq) (bool, error)
	ResetPasswordToken(req *dto.ResetPasswordTokenReq) (bool, error)
	ValidateResetPasswordToken(req dto.ValidateResetPasswordTokenReq) (bool, error)
}
type IClientRepository interface {
	GetClient(clientID int32) (*dto.GetClientRes, error)
	GetClientPhoto(clientID int32) (string, error)
	CreateClient(arg *models.CreateClientParams) (*dto.CreateClient, error)
	UpdateClient(arg *models.UpdateClientByIdParams) error
	UpdateStatusClient(arg *models.UpdateClientStatusByIdParams) error
	ListClient() (*[]dto.GetClientRes, error)
	DeleteClient(arg *models.DeleteClientByIdParams) error
	LoginClient(email string) (*models.Client, error)
	UpdatePhotoClient(arg *models.UpdateClientPhotoParams) error
}

type IClientHandler interface {
	GetSecure(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	GetPhoto(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	UpdateStatus(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
	ResetPasswordToken(c *fiber.Ctx) error
	ValidateResetPasswordToken(c *fiber.Ctx) error
	UpdatePhoto(c *fiber.Ctx) error
}
