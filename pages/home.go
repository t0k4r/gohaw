package pages

import (
	"gohaw/db"
	"gohaw/views"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	animes, err := db.AnimesNow(50)
	if err != nil {
		fail(w, err)
		return
	}
	if isHx(r) {
		render(w, views.Home(animes))
	} else {
		render(w, views.Page("gohaw - Home", views.Home(animes)))
	}
}
