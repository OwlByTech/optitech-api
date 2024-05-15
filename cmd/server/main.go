package main

import (
	"log"
	"optitech/initialize"
	"os"

	"optitech/internal/repository"
	sq "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := initialize.Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}

	repository.Queries = *sq.New(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Optitech API running with hot reload...")
	})

	app.Listen(":" + os.Getenv("PORT"))
}
