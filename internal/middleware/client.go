package middleware

import (
	adto "optitech/internal/dto/asesor"
	cdto "optitech/internal/dto/client"

	"optitech/internal/interfaces"
	"optitech/internal/security"
	"strings"

	cfg "optitech/internal/config"

	"github.com/gofiber/fiber/v2"
)

type ClientMiddleware struct {
	ClientService interfaces.IClientService
}
type InstitutionMiddleware struct {
	InstitutionService interfaces.IInstitutionService
}
type AsesorMiddleware struct {
	AsesorService interfaces.IAsesorService
}

type DirectoryMiddleware struct {
	InstitutionService interfaces.IInstitutionService
	AsesorService      interfaces.IAsesorService
}

func (cm ClientMiddleware) ClientJWT(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format",
		})
	}

	token := splitToken[1]
	var clientVerified cdto.ClientToken
	err := security.JWTGetPayload(token, cfg.Env.JWTSecret, &clientVerified)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "StatusUnauthorized"})
	}

	_, err = cm.ClientService.Get(cdto.GetClientReq{Id: clientVerified.ID})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "StatusUnauthorized"})
	}

	c.Locals("clientId", clientVerified.ID)

	return c.Next()
}
func (in InstitutionMiddleware) InstitutionJWT(c *fiber.Ctx) error {
	data := c.Locals("clientId")
	clientId, ok := data.(int32)

	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}

	res, err := in.InstitutionService.GetByClient(cdto.GetClientReq{Id: clientId})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	c.Locals("institutionId", res)
	return c.Next()
}

func (in DirectoryMiddleware) DirectoryJWT(c *fiber.Ctx) error {
	data := c.Locals("clientId")
	clientId, ok := data.(int32)

	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}

	if institution, _ := in.InstitutionService.GetByClient(cdto.GetClientReq{Id: clientId}); institution > 0 {
		c.Locals("institutionId", institution)
		return c.Next()

	}
	if asesor, _ := in.AsesorService.Get(adto.GetAsesorReq{Id: clientId}); asesor != nil {

		c.Locals("asesorId", asesor.Id)
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Status Unauthorized"})
}
func (in AsesorMiddleware) AsesorJWT(c *fiber.Ctx) error {
	data := c.Locals("clientId")
	clientId, ok := data.(int32)

	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}

	asesor, err := in.AsesorService.Get(adto.GetAsesorReq{Id: clientId})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	c.Locals("asesorId", asesor.Id)
	return c.Next()

}
