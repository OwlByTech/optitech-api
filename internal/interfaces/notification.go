package interfaces

import (
	dto "optitech/internal/dto/notification"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type INotificationService interface {
	Create(req *dto.CreateNorificationReq) (*dto.CreateNotificationRes, error)
	Get(req dto.GetNotificationReq) (*dto.GetNotificationRes, error)
	List() (*[]dto.GetNotificationRes, error)
	Update(req *dto.UpdateNotificationVisualizedReq) (bool, error)
}

type INotificationRepositoy interface {
	CeateNotification(arg *models.CreateNoficationParams) (*dto.CreateNotificationRes, error)
	GetNotification(req *models.GetNotificationParams) (*dto.GetNotificationRes, error)
	ListNotifications() (*[]dto.GetNotificationRes, error)
	UpdateNotificationVisualized(arg *models.UpdateNotificationVisualizedParams) error
}

type INotificationHandler interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}
