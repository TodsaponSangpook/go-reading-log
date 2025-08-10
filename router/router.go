package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/todsapon/go-reading-log/controller"
	"github.com/todsapon/go-reading-log/middleware"
)

func SetupRoutes(app *fiber.App) {
	userRoute := app.Group("/user")
	userRoute.Post("/register", middleware.ValidateBody[controller.RegisterRequest](), controller.Register)
	userRoute.Post("/login", middleware.ValidateBody[controller.LoginRequest](), controller.Login)
	userRoute.Get("/profile", middleware.JWTMiddleware(), controller.GetProfile)

	bookRoute := app.Group("/books", middleware.JWTMiddleware())
	bookRoute.Post("", middleware.ValidateBody[controller.CreateBookRequest](), controller.CreateBook)
	bookRoute.Patch("/:id", middleware.ValidateBody[controller.UpdateBookRequest](), controller.UpdateBook)
	bookRoute.Patch("/:id/status", middleware.ValidateBody[controller.UpdateBookStatusRequest](), controller.UpdateBookStatus)
	bookRoute.Get("", controller.GetBooks)
}
