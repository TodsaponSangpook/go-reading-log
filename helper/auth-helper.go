package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/todsapon/go-reading-log/config"
)

const (
	AccessTTL  = 1 * time.Minute
	RefreshTTL = 7 * 24 * time.Hour

	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
	ClaimUserID      = "user_id"
	ClaimIat         = "iat"
	ClaimExp         = "exp"
	ClaimType        = "type"
)

func Hash(s string) string {
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}

func NewAccessToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		ClaimUserID: userID,
		ClaimIat:    time.Now().Unix(),
		ClaimExp:    time.Now().Add(AccessTTL).Unix(),
		ClaimType:   TokenTypeAccess,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetJwtSecret()))
}

func NewRefreshToken(userID int) (token string, exp time.Time, err error) {
	claims := jwt.MapClaims{
		ClaimUserID: userID,
		ClaimIat:    time.Now().Unix(),
		ClaimExp:    time.Now().Add(RefreshTTL).Unix(),
		ClaimType:   TokenTypeRefresh,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := refreshToken.SignedString([]byte(config.GetJwtRefreshSecret()))
	if err != nil {
		return "", time.Time{}, err
	}

	exp = time.Now().Add(RefreshTTL)
	return signed, exp, nil
}

func ParseToken(tokenStr string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims format")
	}
	return claims, nil
}
