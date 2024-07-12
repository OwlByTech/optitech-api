package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
)

func (s *Server) RoutesFile() {
	s.app.Get("api/assets/:name", func(c *fiber.Ctx) error {
		params_id := c.Params("name")
		err := filesystem.SendFile(c, http.Dir("./uploads"), params_id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("File not found")
		}
		return nil
	})
}
