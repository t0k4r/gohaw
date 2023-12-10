package pages

import (
	"fmt"
	"gohaw/db"
	"gohaw/views"
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
	anime, err := db.FatAnimeFromId(id)
	if err != nil {
		fail(w, err)
		return
	}
	if anime != nil {
		if isHx(r) {
			render(w, views.Anime(anime))
		} else {
			render(w, views.Page(fmt.Sprintf("gohaw - %v", anime.Title), views.Anime(anime)))
		}
	}

}
