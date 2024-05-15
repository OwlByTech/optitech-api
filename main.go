package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Optitech API running with hot reload...")
	})

	app.Listen(":" + os.Getenv("PORT"))
}
