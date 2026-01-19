package main

import (
	"io"
	"log"
	"readdb/database"
	"readdb/handlers"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	// DB
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	database.SeedAuthor(db)
	database.SeedBook(db)
	database.SeedList(db)
	// Router
	e := echo.New()
	renderer := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer
	e.Static("/static", "static")
	handlers.RegisterHandlers(e, db)
	e.Logger.Fatal(e.Start(":6969"))
}
