package handlers

import (
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func RegisterHandlers(e *echo.Echo, db *bolt.DB) {
	// Root
	e.GET("/", handleRoot)
	// Books
	e.GET("/books", GetBooksHandler(db))
	// Authors
	e.GET("/authors", GetAuthorsHandler(db))
}

func handleRoot(ctx echo.Context) error {
	return ctx.JSON(200, map[string]string{"message": "/"})
}
