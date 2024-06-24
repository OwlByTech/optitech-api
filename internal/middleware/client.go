package middleware

import (
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
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}

	_, err = cm.ClientService.Get(cdto.GetClientReq{Id: clientVerified.ID})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}

	c.Locals("clientId", clientVerified.ID)

	return c.Next()
}
