package pages

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func Filtered(title string) func(echo.Context) error {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.NoContent(400)
		}
		rows, err := DB.Query(`
		select  a.id, a.title, a.cover  from anime_infos ai  
		join animes a on a.id = ai.anime_id
		where ai.info_id = $1`, id)
		if err != nil {
			return err
		}
		var a []animeThumbnail
		for rows.Next() {
			var c animeThumbnail
			err := rows.Scan(&c.Id, &c.Title, &c.Cover)
			if err != nil {
				return err
			}
			a = append(a, c)
		}
		if isHx(c.Request()) {
			return c.Render(200, "Home", a)
		}
		return c.Render(200, "pageFiltered.html", a)
	}
}
