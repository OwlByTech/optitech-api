package handler

import (
	ndto "optitech/internal/dto/notification"
	"optitech/internal/interfaces"

	dto "optitech/internal/dto"

	"github.com/gofiber/fiber/v2"
)

type handlerNotification struct {
	notificationService interfaces.INotificationService
}

func NewHandlerNotification(r interfaces.INotificationService) interfaces.INotificationHandler {
	return &handlerNotification{
		notificationService: r,
	}
}

func (h *handlerNotification) Create(c *fiber.Ctx) error {
	req := &ndto.CreateNorificationReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Entrada inválida"+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.notificationService.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerNotification) Get(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &ndto.GetNotificationReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.notificationService.Get(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerNotification) List(c *fiber.Ctx) error {
	res, err := h.notificationService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handlerNotification) Update(c *fiber.Ctx) error {
	params_id := c.AllParams()
	req_id := &ndto.GetNotificationReq{}

	if err := dto.ValidateParamsToDTO(params_id, req_id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	visualized := true

	req := &ndto.UpdateNotificationVisualizedReq{
		NotificationID: req_id.ID,
		Visualized:     visualized,
	}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unvalid entry: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	success, err := h.notificationService.Update(req)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(success)
}
