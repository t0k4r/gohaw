package pages

import (
	"gohaw/db"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	animes, err := db.AnimesNow(50)
	if err != nil {
		fail(w, err)
	}
	if isHx(r) {
		render(w, "Filter", animes)
	} else {
		render(w, "pageHome.go.html", animes)
	}
}
