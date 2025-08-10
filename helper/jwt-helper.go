package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/todsapon/go-reading-log/constants"
)

func GetUserIDFromCtx(c *fiber.Ctx) int {
	token := c.Locals(constants.CtxKeyJwt).(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIDFloat := claims["user_id"].(float64)
	return int(userIDFloat)
}
