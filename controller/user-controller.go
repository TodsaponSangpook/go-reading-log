package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/todsapon/go-reading-log/config"
	"github.com/todsapon/go-reading-log/db"
	"github.com/todsapon/go-reading-log/model"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	// Generate hashed password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}

	// Insert user into database
	_, err = db.DB.Exec(
		context.Background(),
		"CALL register_user($1, $2, $3)",
		req.Email, string(hashed), req.Name,
	)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Email already used")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func Login(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	var user model.User
	err := db.DB.QueryRow(
		context.Background(),
		"SELECT * FROM login_get_user($1)",
		req.Email,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	// generate token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(config.GetJwtSecret()))

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(fiber.Map{"token": signed})
}
