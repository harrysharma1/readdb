package handlers

import (
	"readdb/database"

	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func GetAuthorsHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authors, err := database.GetAllAuthors(db)
		if err != nil {
			return ctx.JSON(500, map[string]string{
				"error": err.Error(),
			})
		}
		return ctx.JSON(200, authors)
	}
}
