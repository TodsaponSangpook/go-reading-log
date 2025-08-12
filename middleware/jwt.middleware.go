package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/todsapon/go-reading-log/config"
	"github.com/todsapon/go-reading-log/constants"
	"github.com/todsapon/go-reading-log/model"
)

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.GetJwtSecret()),
		ContextKey: constants.CtxKeyJwt,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return model.FailedResponse(c, fiber.StatusUnauthorized, "JWT error: "+err.Error())
		},
	})
}
