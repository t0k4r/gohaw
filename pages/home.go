package pages

import (
	"gohaw/db"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	animes, err := db.AnimesNow(20)
	if err != nil {
		return err
	}
	if isHx(c.Request()) {
		return c.Render(200, "Home", animes)
	}
	return c.Render(200, "pageHome.html", animes)
}
