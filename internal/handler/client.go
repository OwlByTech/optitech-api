package handler

import "github.com/gofiber/fiber/v2"


func GetClientHandler(c *fiber.Ctx) error {
  return c.SendString("Optitech API running with hot reload...")
}
