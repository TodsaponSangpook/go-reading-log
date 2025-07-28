package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/todsapon/go-reading-log/controller"
)

func SetupRoutes(app *fiber.App) {
	userRoute := app.Group("/user")
	userRoute.Get("/register", controller.Register)
	userRoute.Get("/login", controller.Login)
}
