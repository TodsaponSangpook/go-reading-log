package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/todsapon/go-reading-log/config"
	"github.com/todsapon/go-reading-log/db"
	"github.com/todsapon/go-reading-log/router"
)

func main() {
	config.LoadEnv()
	db.Connect()
	defer db.DB.Close()

	app := fiber.New()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(config.GetPort()))
}
