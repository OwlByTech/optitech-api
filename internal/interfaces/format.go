package interfaces

import (
	f "optitech/internal/dto/format"

	"github.com/gofiber/fiber/v2"
)

type IFormatService interface {
	Get(req f.GetFormatReq) (*f.GetFormatRes, error)
}

type IFormatRepository interface {
	GetFormat(formatID int32) (*f.GetFormatRes, error)
}

type IFormatHandler interface {
	Get(f *fiber.Ctx) error
}
