package handlers

import (
	"readdb/database"

	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func RegisterHandlers(e *echo.Echo, db *bolt.DB) {
	// Root
	e.GET("/", handleRoot(db))
	// Books
	e.GET("/books", GetBooksHandler(db))
	e.GET("/books/:id", GetBookByIdHandler(db))
	// Authors
	e.GET("/authors", GetAuthorsHandler(db))
	//Lists
	e.GET("/lists", GetListsHandler(db))
}

func handleRoot(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		books, err := database.GetAllBooks(db)
		if err != nil {
			return ctx.String(404, "books not found")
		}
		return ctx.Render(200, "index", map[string]interface{}{
			"Books": books,
		})
	}
}
