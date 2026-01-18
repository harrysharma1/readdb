package handlers

import (
	"readdb/database"

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
		return ctx.JSON(200, books)
	}
}
