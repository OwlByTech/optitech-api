package handler

import (
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/institution_client"
	"optitech/internal/interfaces"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type handlerInstitutionClient struct {
	institutionClientService interfaces.IInstitutionClientService
}

func NewHandlerInstitutionClient(r interfaces.IInstitutionClientService) interfaces.IInstitutionClientHandler {
	return &handlerInstitutionClient{
		institutionClientService: r,
	}
}

func (h *handlerInstitutionClient) Update(c *fiber.Ctx) error {
	req := &cdto.UpdateInstitutionClientReq{}

	err := c.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = dto.ValidateDTO(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.institutionClientService.Update(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerInstitutionClient) GetClient(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetInstitutionClientReq{}

	if id, ok := params["id"]; ok {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid ID format")
		}
		req.InstitutionId = int32(idInt)
	}

	res, err := h.institutionClientService.GetByInstitutionId(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
