package handlers

import (
	"readdb/database"
	"strconv"

	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func GetBooksHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		books, err := database.GetAllBooks(db)
		if err != nil {
			return ctx.JSON(500, map[string]string{
				"error": err.Error(),
			})
		}
		if len(books) > 0 {
			return ctx.JSON(200, books)
		}
		return ctx.JSON(200, "")
	}
}

func GetBookByIdHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return ctx.String(400, "invalid book id")
		}
		book, err := database.GetBookById(db, uint(idInt))
		if err != nil {
			return ctx.String(404, "books not found")
		}
		return ctx.Render(200, "book", map[string]interface{}{
			"Book": book,
		})

	}
}
