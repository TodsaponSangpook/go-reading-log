package helper

import (
	"fmt"
	"time"

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

func GetUserIDFromClaims(claims jwt.MapClaims) (int, error) {
	uid, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id not found or invalid")
	}
	return int(uid), nil
}

func GetTokenTypeFromClaims(claims jwt.MapClaims) (string, error) {
	typ, ok := claims["typ"].(string)
	if !ok {
		return "", fmt.Errorf("token type not found")
	}
	return typ, nil
}

func IsTokenExpiredFromClaims(claims jwt.MapClaims) bool {
	exp, ok := claims["exp"].(float64)
	if !ok {
		return true
	}
	return time.Now().Unix() >= int64(exp)
}
