package main

import (
	"log"
	"readdb/database"
	"readdb/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	database.SeedAuthor(db)
	database.SeedBook(db)
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]string{"message": "/"})
	})
	handlers.RegisterHandlers(e, db)
	e.Logger.Fatal(e.Start(":6969"))
}
