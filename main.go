package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/todsapon/go-reading-log/config"
	"github.com/todsapon/go-reading-log/db"
)

func main() {
	config.LoadEnv()
	db.Connect()
	defer db.DB.Close()

	app := fiber.New()

	app.Listen(config.GetPort())
}
