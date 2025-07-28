package middleware

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/todsapon/go-reading-log/constants"
)

var validate = validator.New()

func ValidateBody[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body T

		if err := c.BodyParser(&body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}

		// validate struct
		if err := validate.Struct(body); err != nil {
			errs := err.(validator.ValidationErrors)

			errorMap := make(map[string]string)
			for _, e := range errs {
				errorMap[strings.ToLower(e.Field())] = e.Error()
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": errorMap,
			})
		}

		c.Locals(constants.CtxKeyBody, body)
		return c.Next()
	}
}
