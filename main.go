package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/todsapon/go-reading-log/db"
)

func main() {
	db.Connect()
	defer db.Pool.Close()

	app := fiber.New()

	app.Listen(":8080")
}
