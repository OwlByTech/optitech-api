package handler

import (
	"encoding/json"
	dto "optitech/internal/dto"
	clientDto "optitech/internal/dto/client"
	cdto "optitech/internal/dto/institution"
	"optitech/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handlerInstitution struct {
	institutionService interfaces.IInstitutionService
}

func NewHandlerInstitution(r interfaces.IInstitutionService) interfaces.IInstitutionHandler {
	return &handlerInstitution{
		institutionService: r,
	}
}
func (h *handlerInstitution) GetByClient(c *fiber.Ctx) error {
	data := c.Locals("clientId")
	clientId, ok := data.(int32)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}

	res_id, err := h.institutionService.GetByClient(clientDto.GetClientReq{Id: clientId})

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.institutionService.Get(cdto.GetInstitutionReq{Id: res_id})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(res)
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

func (h *handlerInstitution) GetLogo(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetInstitutionReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	res, err := h.institutionService.GetLogo(*req)

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
	data := c.Locals("clientId")
	clientId, ok := data.(int32)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}

	req := &cdto.CreateInstitutionReq{
		Clients: []int32{clientId},
	}
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}
	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	res, err := h.institutionService.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerInstitution) UpdateAsesor(c *fiber.Ctx) error {
	req := &cdto.UpdateAsesorInstitutionReq{}

	err := c.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = dto.ValidateDTO(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	r, err := h.institutionService.UpdateAsesor(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(r)

}

func (h *handlerInstitution) Update(c *fiber.Ctx) error {
	params_id := c.AllParams()
	req_id := &cdto.GetInstitutionReq{}
	if err := dto.ValidateParamsToDTO(params_id, req_id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	req := &cdto.UpdateInstitutionReq{
		InstitutionID: req_id.Id,
	}
	params := c.FormValue("data")
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	res, err := h.institutionService.Update(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerInstitution) UpdateLogo(c *fiber.Ctx) error {
	params_id := c.AllParams()
	req_id := &cdto.GetInstitutionReq{}
	if err := dto.ValidateParamsToDTO(params_id, req_id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	req := &cdto.UpdateLogoReq{
		InstitutionID: req_id.Id,
	}
	file, err := c.FormFile("logo")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	req.LogoFile = file
	res, err := h.institutionService.UpdateLogo(req)
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

func (h *handlerInstitution) CreateAllFormat(c *fiber.Ctx) error {
	req := &cdto.GetInstitutionReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	success, err := h.institutionService.CreateAllFormat(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"success": success,
	})
}
