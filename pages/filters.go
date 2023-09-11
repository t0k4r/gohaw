package pages

import (
	"gohaw/db"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Filters(filter string) http.HandlerFunc {
	var get func() (*db.Filters, error)
	switch filter {
	case "types":
		get = db.Types
	case "seasons":
		get = db.Seasons
	default:
		get = func() (*db.Filters, error) { return db.Filtes(filter) }
	}
	return func(w http.ResponseWriter, r *http.Request) {
		filter, err := get()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if isHx(r) {
			render(w, "Filters", filter)
		} else {
			render(w, "pageFilters.go.html", filter)
		}

	}
}

func Filter(filter string) http.HandlerFunc {
	var get func(int) ([]db.Anime, error)
	switch filter {
	case "types":
		get = db.AnimesFromTypeId
	case "seasons":
		get = db.AnimesFromSeasonId
	default:
		get = db.AnimesFromInfoId
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			fail(w, err)
			return
		}
		animes, err := get(id)
		if err != nil {
			fail(w, err)
			return
		}
		if isHx(r) {
			render(w, "Filter", animes)
		} else {
			render(w, "pageFilter.go.html", animes)

		}
	}
}
