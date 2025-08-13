package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/todsapon/go-reading-log/config"
	"github.com/todsapon/go-reading-log/constants"
	"github.com/todsapon/go-reading-log/db"
	"github.com/todsapon/go-reading-log/helper"
	"github.com/todsapon/go-reading-log/model"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

func Register(c *fiber.Ctx) error {
	req := c.Locals(constants.CtxKeyBody).(RegisterRequest)

	// Generate hashed password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.FailedResponse(c, fiber.StatusInternalServerError, "Failed to hash password.")
	}

	// Insert user into database
	_, err = db.DB.Exec(
		context.Background(),
		"CALL register_user($1, $2, $3)",
		req.Email, string(hashed), req.Name,
	)

	if err != nil {
		return model.FailedResponse(c, fiber.StatusBadRequest, "Email is already in use.")
	}

	return model.SuccessResponse[any](c, fiber.StatusCreated, nil, "")
}

func Login(c *fiber.Ctx) error {
	req := c.Locals(constants.CtxKeyBody).(LoginRequest)

	var user model.User
	err := db.DB.QueryRow(
		context.Background(),
		"SELECT * FROM login_get_user($1)",
		req.Email,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name)

	if err != nil {
		return model.FailedResponse(c, fiber.StatusUnauthorized, "User not found.")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		return model.FailedResponse(c, fiber.StatusUnauthorized, "Invalid credentials.")
	}

	accessToken, err := helper.NewAccessToken(user.ID)
	if err != nil {
		return model.FailedResponse(c, fiber.StatusInternalServerError, "Failed to generate access token.")
	}

	refreshToken, exp, err := helper.NewRefreshToken(user.ID)
	if err != nil {
		return model.FailedResponse(c, fiber.StatusInternalServerError, "Failed to generate refresh token.")
	}

	_, err = db.DB.Exec(
		context.Background(),
		"CALL insert_refresh_token($1, $2, $3)",
		user.ID, refreshToken, exp,
	)
	if err != nil {
		return model.FailedResponse(c, fiber.StatusInternalServerError, "Failed to save refresh token.")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   false, // Enable on production.
		SameSite: "Lax",
		Expires:  exp,
		Path:     "/",
	})

	data := fiber.Map{
		"token": accessToken,
		"exp":   time.Now().Add(helper.AccessTTL).Unix(),
	}
	return model.SuccessResponse[any](c, fiber.StatusOK, data, "")
}

func RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return model.FailedResponse(c, fiber.StatusUnauthorized, "Missing refresh token")
	}

	// 1) Validate the JWT signature, type, and expiration from the refresh token
	tok, err := jwt.ParseWithClaims(refreshToken, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtRefreshSecret()), nil
	})
	if err != nil || !tok.Valid {
		// Token is invalid or signature verification failed
		return model.FailedResponse(c, fiber.StatusUnauthorized, "Invalid refresh token")
	}

	claims := tok.Claims.(jwt.MapClaims)
	if typ, _ := claims["typ"].(string); typ != "refresh" {
		// The "typ" claim must be "refresh" to ensure it's not an access token
		return model.FailedResponse(c, fiber.StatusUnauthorized, "Invalid token type")
	}

	// Check expiration from JWT claims (extra safety in case DB is tampered)
	jwtExpUnix := int64(claims["exp"].(float64))
	if time.Now().Unix() >= jwtExpUnix {
		return model.FailedResponse(c, fiber.StatusUnauthorized, "Refresh token expired")
	}
	userID := int(claims["user_id"].(float64))

	var refreshTokenExpires time.Time
	err = db.DB.QueryRow(
		context.Background(),
		"SELECT * FROM get_refresh_token($1, $2)",
		userID, refreshToken,
	).Scan(&refreshTokenExpires)
	if err != nil {
		return model.FailedResponse(c, fiber.StatusUnauthorized, "Refresh token not found")
	}
	if time.Now().After(refreshTokenExpires) {
		return model.FailedResponse(c, fiber.StatusUnauthorized, "Refresh token expired")
	}

	// 3) Generate a new access token for the user
	accessToken, err := helper.NewAccessToken(userID)
	if err != nil {
		return model.FailedResponse(c, fiber.StatusInternalServerError, "Failed to generate access token.")
	}

	// Return the new access token and its expiration time to the client
	data := fiber.Map{
		"token": accessToken,
		"exp":   time.Now().Add(helper.AccessTTL).Unix(),
	}
	return model.SuccessResponse[any](c, fiber.StatusOK, data, "")
}

func GetProfile(c *fiber.Ctx) error {
	token := c.Locals(constants.CtxKeyJwt).(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))

	var user model.User
	err := db.DB.QueryRow(
		context.Background(),
		"SELECT * FROM get_user_profile($1)",
		userID,
	).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt)

	if err != nil {
		return model.FailedResponse(c, fiber.StatusNotFound, "User not found.")
	}

	return model.SuccessResponse(c, fiber.StatusOK, user, "")
}
