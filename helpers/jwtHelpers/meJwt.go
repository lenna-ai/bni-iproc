package jwthelpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func MeJwt(c *fiber.Ctx, data *map[string]any)  {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	*data = claims["data"].(map[string]any)
}
