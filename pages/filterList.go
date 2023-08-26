package pages

import (
	"gohaw/db"

	"github.com/labstack/echo/v4"
)

func FilterList(title string) func(echo.Context) error {
	return func(c echo.Context) error {
		fl, err := db.Filtes(title)
		if err != nil {
			return err
		}
		if isHx(c.Request()) {
			return c.Render(200, "FilterList", fl)
		}
		return c.Render(200, "pageFliterList.html", fl)
	}
}

func TypesList(c echo.Context) error {
	fl, err := db.Types()
	if err != nil {
		return err
	}
	if isHx(c.Request()) {
		return c.Render(200, "FilterList", fl)
	}
	return c.Render(200, "pageFliterList.html", fl)
}

func SeasonsList(c echo.Context) error {
	fl, err := db.Seasons()
	if err != nil {
		return err
	}
	if isHx(c.Request()) {
		return c.Render(200, "FilterList", fl)
	}
	return c.Render(200, "pageFliterList.html", fl)
}
