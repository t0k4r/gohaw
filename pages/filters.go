package pages

import (
	"gohaw/db"
	"net/http"
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
		}
		if isHx(r) {
			render(w, "Filters", filter)
		} else {
			render(w, "pageFilters.go.html", filter)
		}

	}
}
