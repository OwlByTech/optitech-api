package interfaces

import (
	dto "optitech/internal/dto/format"
	f "optitech/internal/dto/format"
	models "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IFormatService interface {
	Get(req f.GetFormatReq) (*f.GetFormatRes, error)
	Create(req *f.CreateFormatReq) (*f.CreateFormatRes, error)
	List() (*[]dto.GetFormatRes, error)
}

type IFormatRepository interface {
	GetFormat(formatID int32) (*f.GetFormatRes, error)
	CreateFormat(arg *models.CreateFormatParams) (*f.CreateFormatRes, error)
	List() (*[]f.GetFormatRes, error)
}

type IFormatHandler interface {
	Get(f *fiber.Ctx) error
	Create(f *fiber.Ctx) error
	List(c *fiber.Ctx) error
}
