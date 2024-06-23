package handler

import (
	"github.com/gofiber/fiber/v2"
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/institution_client"
	"optitech/internal/interfaces"
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
