package main

import (
	"database/sql"
	"gohaw/db"
	"gohaw/pages"
	"gohaw/static"
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
	db.DB, err = sql.Open("postgres", os.Getenv("MALSQL_DB"))
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", pages.Home)
	r.Handle("/*", static.Handler())
	r.Get("/anime/{id}", pages.Anime)

	port := ":3000"
	log.Printf("Running on %v\n", port)
	log.Panic(http.ListenAndServe(port, r))

}
