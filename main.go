package main

import (
	"database/sql"
	"gohaw/pages"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	db, err := sql.Open("postgres", os.Getenv("PG_CONN"))
	if err != nil {
		log.Panic(err)
	}
	pages.DB = db

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", pages.Home())
	r.Get("/style.css", func(w http.ResponseWriter, r *http.Request) {
		styles, err := os.Open("templates/style.css")
		if err != nil {
			panic(err)
		}
		buf, err := io.ReadAll(styles)
		if err != nil {
			panic(err)
		}
		w.Header().Set("content-type", "text/css")
		w.Write(buf)
	})
	r.Get("/genres", pages.Genres())
	r.Get("/themes", pages.Themes())
	r.Get("/studios", pages.Studios())

	http.ListenAndServe(":3000", r)
}
