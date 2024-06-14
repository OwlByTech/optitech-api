package handler

import (
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/client"
	service "optitech/internal/service/client"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetClientHandler(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetClientReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		// TODO: create a error handling structure for bad and good message with
		// status, for example nestjs error handling in json form
		// you should implement the previous as a midddleware for errors
		// follow the next reference https://docs.gofiber.io/guide/error-handling/
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := service.GetClientService(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func CreateClientHandler(c *fiber.Ctx) error {
	req := &cdto.CreateClientReq{}

	err := c.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = dto.ValidateDTO(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	r, err := service.CreateClientService(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(r)
}

func UpdateClientHandler(c *fiber.Ctx) error {
	clientIDStr := c.Params("id")
	clientID, err := strconv.ParseInt(clientIDStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Client ID")
	}

	req := &cdto.UpdateClientReq{
		ClientID: clientID,
	}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := service.UpdateClientService(*req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "Client updated successfully",
	})
}
