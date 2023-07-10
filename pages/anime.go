package pages

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type anime struct {
	Id          int
	Title       string
	Description string
	MalUrl      string
	Cover       string
}

func Anime(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(400)
	}
	row := DB.QueryRow(`
	select a.id, a.title, a.description, a.mal_url, a.cover_url
	from animes a
	where a.id = $1`, id)
	if row.Err() != nil {
		return row.Err()
	}
	var a anime
	err = row.Scan(&a.Id, &a.Title, &a.Description, &a.MalUrl, &a.Cover)
	if err != nil {
		return err
	}
	if isHx(c.Request()) {
		return c.Render(200, "Anime", a)
	}
	return c.Render(200, "pageAnime.html", a)
}
