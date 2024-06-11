package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/institution"
	"optitech/internal/interfaces"
)

type handlerInstitution struct {
	institutionService interfaces.IInstitutionService
}

func NewHandlerInstitution(r interfaces.IInstitutionService) interfaces.IHandler {
	return &handlerInstitution{
		institutionService: r,
	}
}

func (h *handlerInstitution) Get(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetInstitutionReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.institutionService.Get(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerInstitution) List(c *fiber.Ctx) error {

	res, err := h.institutionService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerInstitution) Create(c *fiber.Ctx) error {

	req := &cdto.CreateInstitutionReq{}
	params := c.FormValue("data")
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	file, err := c.FormFile("file")
	req.LogoFile = file
	res, err := h.institutionService.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerInstitution) Update(c *fiber.Ctx) error {
	params_id := c.AllParams()
	req_id := &cdto.GetInstitutionReq{}
	if err := dto.ValidateParamsToDTO(params_id, req_id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	req := &cdto.UpdateInstitutionReq{
		InstitutionID: req_id.InstitutionID,
	}
	params := c.FormValue("data")
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	file, err := c.FormFile("file")
	req.LogoFile = file
	res, err := h.institutionService.Update(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerInstitution) Delete(c *fiber.Ctx) error {

	params := c.AllParams()
	req := &cdto.GetInstitutionReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.institutionService.Delete(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)

}
