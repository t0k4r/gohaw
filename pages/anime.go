package pages

import (
	"gohaw/db"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Anime(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fail(w, err)
		return
	}
	anime, err := db.AnimeFromId(id)
	if err != nil {
		fail(w, err)
		return
	}
	if isHx(r) {
		render(w, "Anime", anime)
	} else {
		render(w, "pageAnime.go.html", anime)
	}
}
