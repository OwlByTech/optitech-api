package interfaces

import (
	dto "optitech/internal/dto/services"

	"github.com/gofiber/fiber/v2"
)

type IService interface {
	Get(req dto.GetServiceReq) (*dto.GetServiceRes, error)
	List() (*[]dto.GetServiceRes, error)
}
type IServiceRepository interface {
	GetService(institutionID int32) (*dto.GetServiceRes, error)
	ListServices() (*[]dto.GetServiceRes, error)
}

type IServiceHandler interface {
	Get(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
}
