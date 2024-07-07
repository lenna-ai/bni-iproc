package jwthelpers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func MeJwt(c *fiber.Ctx, data *map[string]any)  {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	*data = claims["user"].(map[string]any)
}


func ExpJwt(c *fiber.Ctx,timeUnix int64) time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
    tm := time.Unix(timeUnix, 0).In(loc)
	return tm
}