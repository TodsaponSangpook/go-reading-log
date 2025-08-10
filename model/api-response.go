package model

import "github.com/gofiber/fiber/v2"

type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func SuccessResponse[T any](c *fiber.Ctx, status int, data T, message string) error {
	if message == "" {
		message = "Success"
	}
	res := APIResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
	return c.Status(status).JSON(res)
}

func FailedResponse(c *fiber.Ctx, status int, message string) error {
	res := APIResponse[any]{
		Success: false,
		Message: message,
		Data:    nil,
	}
	return c.Status(status).JSON(res)
}
