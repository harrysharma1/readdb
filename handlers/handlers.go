package handlers

import "github.com/labstack/echo/v4"

func RegisterHandlers(e *echo.Echo) {
	// Root
	e.GET("/", handleRoot)
	// Books
	e.GET("/books", handleBooks)
}

func handleRoot(ctx echo.Context) error {
	return ctx.JSON(200, map[string]string{"message": "/"})
}
