package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/todsapon/go-reading-log/db"
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
		"INSERT INTO users (email, password_hash, name) VALUES ($1, $2, $3)",
		req.Email, string(hashed), req.Name,
	)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Email already used")
	}

	return c.SendStatus(fiber.StatusCreated)
}
