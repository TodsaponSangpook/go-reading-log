package controller

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/todsapon/go-reading-log/constants"
	"github.com/todsapon/go-reading-log/db"
	"github.com/todsapon/go-reading-log/helper"
	"github.com/todsapon/go-reading-log/model"
)

type CreateBookRequest struct {
	Title    string `json:"title" validate:"required"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Status   string `json:"status"`
	Rating   int    `json:"rating"`
	Review   string `json:"review"`
}

/*
UpdateBookRequest is used to update only provided fields.
Pointer types allow distinguishing between omitted and empty values.
*/
type UpdateBookRequest struct {
	Title    *string `json:"title"`
	Author   *string `json:"author"`
	Category *string `json:"category"`
	Rating   *int    `json:"rating"`
	Review   *string `json:"review"`
}

type UpdateBookStatusRequest struct {
	Status *string `json:"status" validate:"required"`
}

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
		var startedAt sql.NullTime
		var finishedAt sql.NullTime
		var b model.Book

		err := rows.Scan(
			&b.ID, &b.Title, &b.Author, &b.Category,
			&b.Status, &b.Rating, &b.Review,
			&startedAt, &finishedAt, &b.CreatedAt,
		)
		if err != nil {
			log.Println(err)
			return model.FailedResponse(c, fiber.StatusInternalServerError, "Scan failed")
		}

		if startedAt.Valid {
			b.StartedAt = &startedAt.Time
		}
		if finishedAt.Valid {
			b.FinishedAt = &finishedAt.Time
		}
		books = append(books, b)
	}

	return model.SuccessResponse(c, fiber.StatusOK, books, "")
}

func CreateBook(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromCtx(c)
	req := c.Locals(constants.CtxKeyBody).(CreateBookRequest)

	_, err := db.DB.Exec(
		context.Background(),
		"CALL create_book($1, $2, $3, $4, $5, $6, $7)",
		userID, req.Title, req.Author, req.Category, req.Status, req.Rating, req.Review,
	)

	if err != nil {
		return model.FailedResponse(c, fiber.StatusBadRequest, "Book already exists.")
	}

	return model.SuccessResponse[any](c, fiber.StatusCreated, nil, "")
}

func UpdateBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.FailedResponse(c, fiber.StatusBadRequest, "Invalid book ID.")
	}

	userID := helper.GetUserIDFromCtx(c)
	req := c.Locals(constants.CtxKeyBody).(UpdateBookRequest)

	_, err = db.DB.Exec(
		context.Background(),
		"CALL update_book($1, $2, $3, $4, $5, $6, $7)",
		bookID, userID, req.Title, req.Author, req.Category, req.Rating, req.Review,
	)

	if err != nil {
		return model.FailedResponse(c, fiber.StatusInternalServerError, "Update book failed.")
	}

	return model.SuccessResponse[any](c, fiber.StatusOK, nil, "")
}

func UpdateBookStatus(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.FailedResponse(c, fiber.StatusBadRequest, "Invalid book ID.")
	}

	userID := helper.GetUserIDFromCtx(c)
	req := c.Locals(constants.CtxKeyBody).(UpdateBookStatusRequest)

	_, err = db.DB.Exec(
		context.Background(),
		"CALL update_book_status($1, $2, $3)",
		bookID, userID, req.Status,
	)

	if err != nil {
		return model.FailedResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse[any](c, fiber.StatusOK, nil, "")
}
