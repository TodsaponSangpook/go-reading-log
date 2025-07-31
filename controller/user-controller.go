package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/todsapon/go-reading-log/config"
	"github.com/todsapon/go-reading-log/constants"
	"github.com/todsapon/go-reading-log/db"
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
		return model.FailedResponse(c, fiber.StatusInternalServerError, "Failed to hash password")
	}

	// Insert user into database
	_, err = db.DB.Exec(
		context.Background(),
		"CALL register_user($1, $2, $3)",
		req.Email, string(hashed), req.Name,
	)

	if err != nil {
		return model.FailedResponse(c, fiber.StatusBadRequest, "Email already used")
	}

	return model.SuccessResponse[any](c, fiber.StatusCreated, "Success", nil)
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
		return model.FailedResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		return model.FailedResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	// generate token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(config.GetJwtSecret()))

	if err != nil {
		return model.FailedResponse(c, fiber.StatusInternalServerError, "Failed to generate token")
	}

	data := fiber.Map{
		"token": signed,
	}
	return model.SuccessResponse[any](c, fiber.StatusCreated, "Success", data)
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
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.JSON(user)
}
