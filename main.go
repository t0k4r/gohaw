package main

import (
	"database/sql"
	"gohaw/db"
	"gohaw/pages"
	"html/template"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	db.DB, err = sql.Open("postgres", os.Getenv("PG_CONN"))
	if err != nil {
		log.Panic(err)
	}

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", pages.Home)
	e.GET("/anime/:id", pages.Anime)
	e.GET("/types", pages.TypesList)
	e.GET("/genres", pages.FilterList("genres"))
	e.GET("/genres/:id", pages.Filtered("genres"))
	e.GET("/themes", pages.FilterList("themes"))
	e.GET("/themes/:id", pages.Filtered("themes"))
	e.GET("/studios", pages.FilterList("studios"))
	e.GET("/studios/:id", pages.Filtered("studios"))
	e.GET("/seasons", pages.SeasonsList)
	e.StaticFS("/", os.DirFS("public/static"))
	err = e.Start(":3000")
	if err != nil {
		log.Panic(err)
	}
}
