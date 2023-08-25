package pages

import (
	"gohaw/db"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
)

type anime struct {
	db.Anime
	Infos    []db.Infos
	Episodes []db.Episode
}

func Anime(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(400)
	}
	var a anime
	dba, err := db.AnimeFromId(id)
	if dba == nil {
		return c.NoContent(400)
	}
	if err != nil {
		return err
	}
	a.Anime = *dba

	g := new(errgroup.Group)
	g.Go(appendInfo(&a, "genres"))
	g.Go(appendInfo(&a, "themes"))
	g.Go(appendInfo(&a, "demographics"))
	g.Go(appendInfo(&a, "studios"))
	g.Go(appendInfo(&a, "producers"))
	g.Go(appendEpisodes(&a))
	err = g.Wait()
	if err != nil {
		return err
	}
	if isHx(c.Request()) {
		return c.Render(200, "Anime", a)
	}
	return c.Render(200, "pageAnime.html", a)
}

func appendInfo(a *anime, infoType string) func() error {
	return func() error {
		inf, err := db.InfosOfTypeFromAnimeId(a.Id, infoType)
		if err != nil {
			return err
		}
		a.Infos = append(a.Infos, *inf)
		return nil
	}
}
func appendEpisodes(a *anime) func() error {
	return func() error {
		episodes, err := db.EpisodesFromAnimeId(a.Id)
		if err != nil {
			return err
		}
		a.Episodes = episodes
		return nil
	}
}
