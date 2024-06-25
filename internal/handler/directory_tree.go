package handler

import (
	"optitech/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handlerDirectoryTree struct {
	directoryTreeService interfaces.IDirectoryService
}

func NewHnadlerDirectoryTree(r interfaces.IDirectoryService) interfaces.IDirectoryHandler {
	return &handlerDirectoryTree{
		directoryTreeService: r,
	}
}

func (h *handlerDirectoryTree) Get(c *fiber.Ctx) error {
	return nil
}

func (h *handlerDirectoryTree) Create(c *fiber.Ctx) error {
	return nil
}
