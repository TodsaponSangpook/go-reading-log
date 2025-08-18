package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/todsapon/go-reading-log/config"
	"github.com/todsapon/go-reading-log/constants"
	"github.com/todsapon/go-reading-log/helper"
	"github.com/todsapon/go-reading-log/model"
)

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.GetJwtSecret()),
		ContextKey: constants.CtxKeyJwt,
		SuccessHandler: func(c *fiber.Ctx) error {
			token, ok := c.Locals(constants.CtxKeyJwt).(*jwt.Token)
			if !ok || token == nil {
				return model.FailedResponse(c, fiber.StatusUnauthorized, "JWT token missing")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return model.FailedResponse(c, fiber.StatusUnauthorized, "Invalid JWT claims")
			}

			if typ, _ := claims[helper.ClaimType].(string); typ != helper.TokenTypeAccess {
				return model.FailedResponse(c, fiber.StatusUnauthorized, "Invalid token type")
			}

			userID, ok := claims[helper.ClaimUserID].(float64)
			if !ok {
				return model.FailedResponse(c, fiber.StatusUnauthorized, "user_id missing")
			}
			c.Locals(helper.ClaimUserID, userID)
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return model.FailedResponse(c, fiber.StatusUnauthorized, "JWT error: "+err.Error())
		},
	})
}
