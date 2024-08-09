package interfaces

import (
	dto "optitech/internal/dto/notification"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type INotificationService interface {
	Create(req *dto.CreateNorificationReq) (*dto.CreateNotificationRes, error)
}

type INotificationRepositoy interface {
	CeateNotification(arg *models.CreateNoficationParams) (*dto.CreateNotificationRes, error)
}

type INotificationHandler interface {
	Create(c *fiber.Ctx) error
}
