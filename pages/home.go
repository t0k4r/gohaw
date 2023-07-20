package pages

import (
	"time"

	"github.com/labstack/echo/v4"
)

type animeThumbnail struct {
	Id    int
	Cover string
	Title string
}

func Home(c echo.Context) error {
	rows, err := DB.Query(`
	select a.id, a.title, a.cover from animes a
	join seasons s on a.season_id = s.id
	where s.value < $1
	order by s.value desc
	limit 50`, time.Now())
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
	return c.Render(200, "pageHome.html", a)
}
