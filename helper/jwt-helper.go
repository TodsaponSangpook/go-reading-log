package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GetUserIDFromClaims(claims jwt.MapClaims) (int, error) {
	uid, ok := claims[ClaimUserID].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id not found or invalid")
	}
	return int(uid), nil
}

func GetTokenTypeFromClaims(claims jwt.MapClaims) (string, error) {
	typ, ok := claims[ClaimType].(string)
	if !ok {
		return "", fmt.Errorf("token type not found")
	}
	return typ, nil
}

func IsTokenExpiredFromClaims(claims jwt.MapClaims) bool {
	exp, ok := claims[ClaimExp].(float64)
	if !ok {
		return true
	}
	return time.Now().Unix() >= int64(exp)
}
