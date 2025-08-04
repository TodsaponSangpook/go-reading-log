package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/todsapon/go-reading-log/db"
	"github.com/todsapon/go-reading-log/helper"
	"github.com/todsapon/go-reading-log/model"
)

func GetBooks(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromCtx(c)
	rows, err := db.DB.Query(
		context.Background(),
		"SELECT * FROM get_books_by_user($1)",
		userID,
	)

	if err != nil {
		return model.FailedResponse(c, fiber.StatusInternalServerError, "Database query failed")
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var b model.Book
		err := rows.Scan(
			&b.ID, &b.Title, &b.Author, &b.Category,
			&b.Status, &b.Rating, &b.Review,
			&b.StartedAt, &b.FinishedAt, &b.CreatedAt,
		)
		if err != nil {
			return model.FailedResponse(c, fiber.StatusInternalServerError, "Scan failed")
		}
		books = append(books, b)
	}

	return model.SuccessResponse(c, fiber.StatusOK, books, "")
}

// func CreateBook(c *fiber.Ctx) error {
// 	userID := helper.GetUserIDFromCtx(c)
// }
