package handlers

import (
	"readdb/database"
	"strconv"

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
		if len(authors) > 0 {
			return ctx.JSON(200, authors)
		}
		return ctx.JSON(200, "")
	}
}

func GetAuthorByIdHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return ctx.String(400, "invalid author id")
		}
		author, err := database.GetAuthorById(db, uint(idInt))
		if err != nil {
			return ctx.String(404, "authors not found")
		}
		books, err := database.GetAllBooksByAuthorID(db, uint(idInt))
		if err != nil {
			return ctx.String(404, "books not found")
		}
		return ctx.Render(200, "author", map[string]interface{}{
			"Author": author,
			"Books":  books,
		})
	}
}
