package handler

import (
	"encoding/json"
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/institution"
	"optitech/internal/service/institution"

	"github.com/gofiber/fiber/v2"
)

func GetInstitutionHandler(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetInstitutionReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := service.GetInstitutionService(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func ListInstitutionsHandler(c *fiber.Ctx) error {

	res, err := service.ListInstitutionsService()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func CreateInstitutionHandler(c *fiber.Ctx) error {

	req := &cdto.CreateInstitutionReq{}
	params := c.FormValue("data")
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	file, err := c.FormFile("file")
	req.LogoFile = file
	res, err := service.CreateInstitutionService(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func UpdateInstitutionHandler(c *fiber.Ctx) error {

	req := &cdto.UpdateInstitutionReq{}
	params := c.FormValue("data")
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	file, err := c.FormFile("file")
	req.LogoFile = file
	res, err := service.UpdateInstitutionService(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func DeleteInstitutionHandler(c *fiber.Ctx) error {

	params := c.AllParams()
	req := &cdto.GetInstitutionReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := service.DeleteInstitutionService(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)

}
