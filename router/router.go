package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/todsapon/go-reading-log/controller"
)

func SetupRoutes(app *fiber.App) {
	userRoute := app.Group("/user")
	userRoute.Get("test", func(c *fiber.Ctx) error {
		log.Println("TEST")
		return nil
	})
	userRoute.Get("/register", controller.Register)
}
