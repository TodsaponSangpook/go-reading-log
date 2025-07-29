package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/todsapon/go-reading-log/config"
	"github.com/todsapon/go-reading-log/constants"
)

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.GetJwtSecret()),
		ContextKey: constants.CtxKeyJwt,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println(err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		},
	})
}
