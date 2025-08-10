package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/todsapon/go-reading-log/config"
	"github.com/todsapon/go-reading-log/db"
	"github.com/todsapon/go-reading-log/router"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	config.LoadEnv()
	db.Connect()

	app := fiber.New()
	router.SetupRoutes(app)

	adaptor.FiberApp(app).ServeHTTP(w, r)
}
