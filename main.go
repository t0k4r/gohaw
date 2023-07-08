package main

import (
	"database/sql"
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

	pages.DB, err = sql.Open("postgres", os.Getenv("PG_CONN"))
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
	e.GET("/types", pages.Types)
	e.GET("/genres", pages.FilterList("genres"))
	e.GET("/themes", pages.FilterList("themes"))
	e.GET("/studios", pages.FilterList("studios"))
	e.GET("/seasons", pages.Seasons)
	e.StaticFS("/", os.DirFS("public/static"))
	err = e.Start(":3000")
	if err != nil {
		log.Panic(err)
	}
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Get("/", pages.Home())
	// r.Get("/style.css", func(w http.ResponseWriter, r *http.Request) {
	// 	styles, err := os.Open("templates/style.css")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	buf, err := io.ReadAll(styles)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	w.Header().Set("content-type", "text/css")
	// 	w.Write(buf)
	// })
	// r.Get("/genres", pages.Genres())
	// r.Get("/themes", pages.Themes())
	// r.Get("/studios", pages.Studios())

	// http.ListenAndServe(":3000", r)
}
