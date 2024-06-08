package handler

import (
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/services"
	"optitech/internal/service/services"

	"github.com/gofiber/fiber/v2"
)

func GetServiceHandler(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetServiceReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := service.GetService(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func ListServicesHandler(c *fiber.Ctx) error {

	res, err := service.ListServices()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func CreateServiceHandler(c *fiber.Ctx) error {

	req := &cdto.CreateServiceReq{}
	err := c.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = dto.ValidateDTO(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	res, err := service.CreateService(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
