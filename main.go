package main

import (
	"database/sql"
	"gohaw/db"
	"gohaw/pages"
	"gohaw/static"
	"gohaw/templates"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	db.DB, err = sql.Open("postgres", os.Getenv("PG_CONN"))
	if err != nil {
		log.Panic(err)
	}
	pages.Templ, err = template.ParseFS(templates.Files, "*.go.html")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", pages.Home)
	r.Handle("/*", http.FileServer(http.FS(static.Files)))
	r.Get("/types", pages.Filters("types"))
	r.Get("/types/{id}", pages.Filter("types"))
	r.Get("/themes", pages.Filters("themes"))
	r.Get("/themes/{id}", pages.Filter("themes"))
	r.Get("/genres", pages.Filters("genres"))
	r.Get("/genres/{id}", pages.Filter("genres"))
	r.Get("/studios", pages.Filters("studios"))
	r.Get("/studios/{id}", pages.Filter("studios"))
	r.Get("/seasons", pages.Filters("seasons"))
	r.Get("/seasons/{id}", pages.Filter("seasons"))
	r.Get("/anime/{id}", pages.Anime)

	port := ":3000"
	log.Printf("Running on %v\n", port)
	log.Panic(http.ListenAndServe(port, r))

}
