package main

import (
	"io"
	"log"
	"readdb/database"
	"readdb/handlers"
	"readdb/helper"
	"strings"
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
	handlers.RegisterHandlers(e, db)
	// Template
	e.Static("/static", "static")
	funcMap := template.FuncMap{
		"toUpper":  strings.ToUpper,
		"contains": strings.Contains,
		"replace":  strings.Replace,
		"truncate": helper.Truncate,
	}
	renderer := &Template{
		templates: template.Must(template.New("t").Funcs(funcMap).ParseGlob("views/*.html")),
	}
	e.Renderer = renderer
	e.Debug = true
	e.Logger.Fatal(e.Start(":6969"))
}
