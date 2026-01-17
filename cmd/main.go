package main

import (
	"readdb/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]string{"message": "/"})
	})
	handlers.RegisterHandlers(e)
	e.Logger.Fatal(e.Start(":6969"))
}
