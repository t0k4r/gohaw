package pages

import (
	"gohaw/db"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type pageAnime struct {
	*db.Anime
	Episodes []db.Episode
}

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
	episodes, err := db.EpisodesFromAnimeId(id)
	if err != nil {
		fail(w, err)
		return
	}
	page := pageAnime{Anime: anime, Episodes: episodes}
	if isHx(r) {
		render(w, "Anime", page)
	} else {
		render(w, "pageAnime.go.html", page)
	}
}
