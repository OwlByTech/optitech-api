package middleware

import (
	"encoding/json"
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
	unauthorized := c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})

	payloadClaims, err := security.JWTVerify(token, cfg.Env.JWTSecret)

	if err != nil {
		return unauthorized
	}

	var clientVerified cdto.ClientToken
	bytes, err := json.Marshal(payloadClaims.Claims)

	if err != nil {
		return unauthorized
	}

	if err := json.Unmarshal(bytes, &clientVerified); err != nil {
		return unauthorized
	}

	_, err = cm.ClientService.Get(cdto.GetClientReq{Id: int64(clientVerified.ID)})
	if err != nil {
		return unauthorized
	}

	return c.Next()
}
